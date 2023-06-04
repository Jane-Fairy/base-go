package main

import (
	"runtime"
	"sync"
	"time"
)

func liveLockDemo() {
	runtime.GOMAXPROCS(3)
	sync.NewCond(&sync.Mutex{})
	go func() {
		for range time.Tick(1 * time.Second) {

		}
	}()
}
