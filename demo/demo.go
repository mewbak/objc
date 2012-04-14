package main

import (
	. "github.com/mkrautz/objc/AppKit"
	. "github.com/mkrautz/objc/Foundation"
	"runtime"
)

func init() {
	defer runtime.LockOSThread()
}

func main() {
	pool := NewNSAutoreleasePool()
	defer pool.Release()

	app := NSSharedApplication()
	_ = app

	mask := NSTitledWindowMask | NSClosableWindowMask | NSMiniaturizableWindowMask | NSResizableWindowMask
	window := NewNSWindow(NSRectMake(0, 0, 500, 500), mask, NSBackingStoreBuffered, false)
	window.AutoRelease()
	window.Display()
	window.MakeKeyAndOrderFront(window.Object)

	loop := NSRunLoopMain()
	loop.Run()
}
