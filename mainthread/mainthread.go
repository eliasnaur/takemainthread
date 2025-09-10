package mainthread

import "runtime"

func init() {
	// Lock main thread, and receive it in
	// GiveMainThread.
	runtime.LockOSThread()
}

var mainFuncs = make(chan func())
var done = make(chan struct{})

func GiveMainThread() {
	for {
		f := <-mainFuncs
		f()
		done <- struct{}{}
	}
}

func Take(f func()) bool {
	select {
	case mainFuncs <- f:
		<-done
		return true
	default:
		return false
	}
}
