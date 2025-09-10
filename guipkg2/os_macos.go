package guipkg2

/*
#cgo CFLAGS: -Werror -Wno-deprecated-declarations -fobjc-arc -x objective-c
#cgo LDFLAGS: -framework AppKit -framework QuartzCore

#include <AppKit/AppKit.h>

extern void guipkg2_init(void);
extern void guipkg2_main(void);
extern void guipkg2_runOnMain(uintptr_t handle);
extern void guipkg2_createWindow(CGFloat width, CGFloat height);

*/
import "C"
import (
	"runtime/cgo"

	"github.com/eliasnaur/takemainthread/mainthread"
)

func init() {
	// Register launching finished listener.
	C.guipkg2_init()
}

var launched = make(chan struct{})

func NewWindow() {
	go mainthread.Take(func() {
		C.guipkg2_main()
	})
	// Wait for launching finished.
	<-launched
	runOnMain(func() {
		C.guipkg2_createWindow(800, 600)
	})
}

//export guipkg2_onFinishLaunching
func guipkg2_onFinishLaunching() {
	close(launched)
}

// runOnMain runs the function on the main thread.
func runOnMain(f func()) {
	C.guipkg2_runOnMain(C.uintptr_t(cgo.NewHandle(f)))
}

//export guipkg2_runFunc
func guipkg2_runFunc(h C.uintptr_t) {
	handle := cgo.Handle(h)
	defer handle.Delete()
	f := handle.Value().(func())
	f()
}
