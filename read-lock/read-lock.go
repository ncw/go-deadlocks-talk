package main

import "sync"

var rw sync.RWMutex

func RO() {
	rw.RLock()
	defer rw.RUnlock()
	// Read only stuff
	MoreRO()
}

func MoreRO() {
	rw.RLock()
	defer rw.RUnlock()
	// Read only stuff
}

func RW() {
	rw.Lock()
	defer rw.Unlock()
	// Read Write stuff
}

func main() {
	go func() {
		for {
			RO()
		}
	}()
	go func() {
		for {
			RW()
		}
	}()
	select {}
}
