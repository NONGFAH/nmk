package system

import "sync"

type MonitorWorkerFunc func(event Event)

var flag = 0
var rwm sync.RWMutex

func UnLock() {
	rwm.Lock()
	flag = 0
	rwm.Unlock()
}

func lock() {
	rwm.Lock()
	flag = 1
	rwm.Unlock()
}
