//go:build darwin
// +build darwin

package system

func KeyboardMonitor(fc MonitorWorkerFunc) {
	panic("notSupport")
}
func MouseMonitor(fc MonitorWorkerFunc) {
	panic("notSupport")

}

func Input(event Event) error {
	panic("notSupport")
}
