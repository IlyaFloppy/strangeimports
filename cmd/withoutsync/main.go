package main

import (
	"obstest/genericpkg"
)

func main() {
	obj := genericpkg.NewWrapperWithLock("this file does not import sync")
	obj.PrintWithLock()
}
