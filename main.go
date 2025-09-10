package main

import (
	"runtime"

	"github.com/eliasnaur/takemainthread/guipkg1"
	"github.com/eliasnaur/takemainthread/guipkg2"
)

func init() {
	// Darwin requires that UI operations happen on the main thread only.
	runtime.LockOSThread()
}

func main() {
	go func() {
		guipkg1.NewWindow()
		guipkg2.NewWindow()
	}()
	guipkg1.Main()
}
