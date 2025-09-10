package main

import (
	"flag"

	"github.com/eliasnaur/takemainthread/guipkg1"
	"github.com/eliasnaur/takemainthread/guipkg2"
	"github.com/eliasnaur/takemainthread/mainthread"
)

var first = flag.String("first", "pkg1", "'pkg1', 'pkg2'")

// This would be func main() if #70089 were implemented.
func realMain() {
	flag.Parse()
	switch *first {
	case "pkg1":
		guipkg1.NewWindow()
		guipkg2.NewWindow()
	case "pkg2":
		guipkg2.NewWindow()
		guipkg1.NewWindow()
	}
}

func main() {
	// This goroutine and the call to GiveMainThread would not be needed if
	// golang.org/issue/70089 were implemented.
	go func() {
		realMain()
	}()
	mainthread.GiveMainThread()
}
