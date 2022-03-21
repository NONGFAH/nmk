//go:build windows
// +build windows

package system

import (
	"fmt"
	"github.com/ivpusic/grpool"
	"github.com/nongfah/go-hook/pkg/keyboard"
	"github.com/nongfah/go-hook/pkg/mouse"
	"github.com/nongfah/go-hook/pkg/types"
	"github.com/nongfah/go-hook/pkg/win32"
	"os"
	"os/signal"
	"unsafe"
)

func mouseHookHandler(c chan<- types.MouseEvent) types.HOOKPROC {
	return func(code int32, wParam, lParam uintptr) uintptr {
		if lParam != 0 {
			event := types.MouseEvent{
				Message:        types.Message(wParam),
				MSLLHOOKSTRUCT: *(*types.MSLLHOOKSTRUCT)(unsafe.Pointer(lParam)),
			}
			c <- event
			// todo 出屏事件处理
			if event.X == 0 {
				lock()
			}
		}
		win32.CallNextHookEx(0, code, wParam, lParam)
		rwm.RLock()
		defer rwm.RUnlock()
		return uintptr(flag)
	}
}
func keyboardHookHandler(c chan<- types.KeyboardEvent) types.HOOKPROC {
	return func(code int32, wParam, lParam uintptr) uintptr {
		if lParam != 0 {
			c <- types.KeyboardEvent{
				Message:         types.Message(wParam),
				KBDLLHOOKSTRUCT: *(*types.KBDLLHOOKSTRUCT)(unsafe.Pointer(lParam)),
			}
		}
		win32.CallNextHookEx(0, code, wParam, lParam)
		rwm.RLock()
		defer rwm.RUnlock()
		return uintptr(flag)
	}
}

func MouseMonitor(fc MonitorWorkerFunc) {
	pool := grpool.NewPool(32, 16)
	defer pool.Release()
	mouseChan := make(chan types.MouseEvent, 128)
	err := mouse.Install(mouseHookHandler, mouseChan)
	if err != nil {
		return
	}
	defer func() {
		_ = mouse.Uninstall()
	}()
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	fmt.Println("start capturing mouse input")
	var event Event
	for {
		select {
		case <-signalChan:
			fmt.Println("Received shutdown signal")
			return
		case m := <-mouseChan:
			switch m.Message {
			case types.WM_MOUSEMOVE:
				event = Event{EvMouseMove, Code(m.X), Value(m.Y)}
				break
			case types.WM_MOUSEWHEEL:
				event = Event{EvMouseWheelMove, 0, Value(m.MouseData)}
				break
			case types.WM_LBUTTONDOWN:
				event = Event{EvKey, Code(types.VK_LBUTTON), KeyDown}
				break
			case types.WM_RBUTTONDOWN:
				event = Event{EvKey, Code(types.VK_RBUTTON), KeyDown}
				break
			case types.WM_MBUTTONDOWN:
				event = Event{EvKey, Code(types.VK_MBUTTON), KeyDown}
				break
			case types.WM_LBUTTONUP:
				event = Event{EvKey, Code(types.VK_LBUTTON), KeyUp}
				break
			case types.WM_RBUTTONUP:
				event = Event{EvKey, Code(types.VK_RBUTTON), KeyUp}
				break
			case types.WM_MBUTTONUP:
				event = Event{EvKey, Code(types.VK_MBUTTON), KeyUp}
				break
			// todo 增加鼠标侧键按钮
			//case types.WM_SBUTTONDOWN:
			//	event = Event{EvKey, Code(types.VK_XBUTTON1), KeyDown}
			//	break
			//case types.WM_SBUTTONUP:
			//	event = Event{EvKey, Code(types.VK_XBUTTON1), KeyUp}
			//	break
			default:
				continue
			}
			pool.JobQueue <- func() {
				fc(event)
			}
		}
	}
}
func KeyboardMonitor(fc MonitorWorkerFunc) {
	pool := grpool.NewPool(32, 16)
	defer pool.Release()
	keyboardChan := make(chan types.KeyboardEvent, 128)
	err := keyboard.Install(keyboardHookHandler, keyboardChan)
	if err != nil {
		return
	}
	defer func() {
		_ = keyboard.Uninstall()
	}()
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	fmt.Println("start capturing Keyboard input")

	var value Value
	for {
		select {
		case <-signalChan:
			fmt.Println("Received shutdown signal")
			return
		case k := <-keyboardChan:
			switch k.Message {
			case types.WM_KEYUP, types.WM_SYSKEYUP:
				value = KeyUp
				break
			case types.WM_KEYDOWN, types.WM_SYSKEYDOWN:
				value = KeyDown
				break
			default:
				continue
			}
			pool.JobQueue <- func() {
				fc(Event{EvKey, Code(k.ScanCode), value})
			}
		}
	}
}

func Input(event Event) error {
	switch event.Type {
	case EvMouseMove:
		param := types.MSLLHOOKSTRUCT{
			POINT: types.POINT{
				X: int32(event.Code),
				Y: int32(event.Value),
			},
			MouseData:   0,
			Flags:       0x0001,
			Time:        0,
			DWExtraInfo: 0,
		}
		err := mouse.Input(param)
		if err != nil {
			return err
		}
		break
	case EvMouseWheelMove:
		value := 120
		if event.Value < 0 {
			value = -120
		}
		param := types.MSLLHOOKSTRUCT{
			POINT: types.POINT{
				X: 0,
				Y: 0,
			},
			MouseData:   int32(value),
			Flags:       0x0800,
			Time:        0,
			DWExtraInfo: 0,
		}
		err := mouse.Input(param)
		if err != nil {
			return err
		}
		break
	case EvKey:
		switch event.Code {
		case Code(types.VK_LBUTTON), Code(types.VK_RBUTTON), Code(types.VK_MBUTTON):
			if err := mouseKeyPress(event.Code, event.Value == KeyDown); err != nil {
				return err
			}
			break
		case Code(types.VK_XBUTTON1), Code(types.VK_XBUTTON2):
			break
		default:
			vkCode, ok := scanCode2VkCodeMap[event.Code]
			if !ok {
				fmt.Println("failed to map scanCode to vkCode, unknown error")
				break
			}
			param := types.KBDLLHOOKSTRUCT{
				VKCode:      vkCode,
				ScanCode:    uint32(event.Code),
				Flags:       uint32(event.Value),
				Time:        0,
				DWExtraInfo: 0,
			}
			err := keyboard.Input(param)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func mouseKeyPress(code Code, down bool) (err error) {
	flag := 0
	if code == Code(types.VK_LBUTTON) {
		if down {
			flag = 0x0002
		} else {
			flag = 0x0004
		}
	}
	if code == Code(types.VK_RBUTTON) {
		if down {
			flag = 0x0008
		} else {
			flag = 0x0010
		}
	}
	if code == Code(types.VK_MBUTTON) {
		if down {
			flag = 0x0020
		} else {
			flag = 0x0040
		}
	}
	param := types.MSLLHOOKSTRUCT{
		POINT: types.POINT{
			X: 0,
			Y: 0,
		},
		MouseData:   0,
		Flags:       uint32(flag),
		Time:        0,
		DWExtraInfo: 0,
	}
	return mouse.Input(param)
}
