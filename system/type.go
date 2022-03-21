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
	// EvKey 按键事件
	EvKey
)
const (
	KeyDown = 0
	KeyUp   = 128
)
