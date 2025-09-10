package main

import (
	"github.com/eliasnaur/takemainthread/guipkg1"
	"github.com/eliasnaur/takemainthread/guipkg2"
	"github.com/eliasnaur/takemainthread/mainthread"
)

func main() {
	// This goroutine and the call to
	// GiveMainThread would not be needed if
	// golang.org/issue/70089 were implemented.
	go func() {
		realMain()
	}()
	mainthread.GiveMainThread()
}

// This would be func main() if #70089 were implemented.
func realMain() {
	guipkg1.NewWindow()
	guipkg2.NewWindow()
}
