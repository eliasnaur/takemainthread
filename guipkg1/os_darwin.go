package guipkg1

/*
#include <Foundation/Foundation.h>

__attribute__ ((visibility ("hidden"))) void gio_wakeupMainThread(void);

static bool isMainThread() {
	return [NSThread isMainThread];
}
*/
import "C"

var mainFuncs = make(chan func(), 1)

func isMainThread() bool {
	return bool(C.isMainThread())
}

// runOnMain runs the function on the main thread.
func runOnMain(f func()) {
	if isMainThread() {
		f()
		return
	}
	go func() {
		mainFuncs <- f
		C.gio_wakeupMainThread()
	}()
}

//export gio_dispatchMainFuncs
func gio_dispatchMainFuncs() {
	for {
		select {
		case f := <-mainFuncs:
			f()
		default:
			return
		}
	}
}
