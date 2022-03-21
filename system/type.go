package system

type (
	Type  uint8
	Code  int32
	Value int32
	Event struct {
		Type
		Code
		Value
	}
)

const (
	// EvMouseMove 鼠标移动
	EvMouseMove Type = iota
	// EvMouseWheelMove 滚轮移动
	EvMouseWheelMove
	// EvMouseKey 鼠标按键事件
	EvMouseKey
	// EvKeyboardKey 键盘按键事件
	EvKeyboardKey
)
const (
	KeyDown = 0
	KeyUp   = 128
)
