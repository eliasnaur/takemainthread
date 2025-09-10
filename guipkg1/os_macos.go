package guipkg1

/*
#cgo CFLAGS: -Werror -Wno-deprecated-declarations -fobjc-arc -x objective-c
#cgo LDFLAGS: -framework AppKit -framework QuartzCore

#include <AppKit/AppKit.h>

__attribute__ ((visibility ("hidden"))) void gio_main(void);
__attribute__ ((visibility ("hidden"))) void gio_createWindow(CGFloat width, CGFloat height);

*/
import "C"

var launched = make(chan struct{})

//export gio_onFinishLaunching
func gio_onFinishLaunching() {
	close(launched)
}

func NewWindow() {
	<-launched
	res := make(chan struct{})
	runOnMain(func() {
		defer close(res)
		C.gio_createWindow(800, 600)
	})
	<-res
}

func Main() {
	if !isMainThread() {
		panic("app.Main must run on the main goroutine")
	}
	C.gio_main()
}
