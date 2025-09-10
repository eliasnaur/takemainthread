package guipkg1

/*
#include <Foundation/Foundation.h>

void gio_runOnMain(uintptr_t handle);

*/
import "C"
import "runtime/cgo"

// runOnMain runs the function on the main thread.
func runOnMain(f func()) {
	C.gio_runOnMain(C.uintptr_t(cgo.NewHandle(f)))
}

//export gio_runFunc
func gio_runFunc(h C.uintptr_t) {
	handle := cgo.Handle(h)
	defer handle.Delete()
	f := handle.Value().(func())
	f()
}
