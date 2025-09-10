package guipkg1

/*
#cgo CFLAGS: -Werror -Wno-deprecated-declarations -fobjc-arc -x objective-c
#cgo LDFLAGS: -framework AppKit -framework QuartzCore

#include <AppKit/AppKit.h>

extern void guipkg1_init(void);
extern void guipkg1_main(void);
extern void guipkg1_runOnMain(uintptr_t handle);
extern void guipkg1_createWindow(CGFloat width, CGFloat height);

*/
import "C"
import (
	"runtime/cgo"

	"github.com/eliasnaur/takemainthread/mainthread"
)

func init() {
	// Register launching finished listener.
	C.guipkg1_init()
}

var launched = make(chan struct{})

func NewWindow() {
	go mainthread.Take(func() {
		C.guipkg1_main()
	})
	// Wait for launching finished.
	<-launched
	runOnMain(func() {
		C.guipkg1_createWindow(800, 600)
	})
}

//export guipkg1_onFinishLaunching
func guipkg1_onFinishLaunching() {
	close(launched)
}

// runOnMain runs the function on the main thread.
func runOnMain(f func()) {
	C.guipkg1_runOnMain(C.uintptr_t(cgo.NewHandle(f)))
}

//export guipkg1_runFunc
func guipkg1_runFunc(h C.uintptr_t) {
	handle := cgo.Handle(h)
	defer handle.Delete()
	f := handle.Value().(func())
	f()
}
