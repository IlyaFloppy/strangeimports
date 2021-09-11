package main

import (
	"obstest/genericpkg"

	// these imports fix the error
	_ "fmt"
	_ "sync"
)

func main() {
	obj := genericpkg.NewWrapperWithLock("this file does import sync")
	obj.PrintWithLock()
}
