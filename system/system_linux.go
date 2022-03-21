//go:build !windows || !darwin
// +build !windows !darwin

package system

func KeyboardMonitor(fc MonitorWorkerFunc) {

}
func MouseMonitor(fc MonitorWorkerFunc) {
}

func Input(event Event) error {
	return nil
}
