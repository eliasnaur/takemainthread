package guipkg2

/*
#cgo CFLAGS: -Werror -Wno-deprecated-declarations -fobjc-arc -x objective-c
#cgo LDFLAGS: -framework AppKit -framework QuartzCore

#include <AppKit/AppKit.h>

extern void gio_init(void);
extern void gio_main(void);
extern void gio_runOnMain(uintptr_t handle);
extern void gio_createWindow(CGFloat width, CGFloat height);

*/
import "C"
import "runtime/cgo"

func init() {
	// Register launching finished listener.
	C.gio_init()
}

var launched = make(chan struct{})

func NewWindow() {
	// Wait for launching finished.
	<-launched
	runOnMain(func() {
		C.gio_createWindow(800, 600)
	})
}

func Main() {
	C.gio_main()
}

//export gio_onFinishLaunching
func gio_onFinishLaunching() {
	close(launched)
}

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
