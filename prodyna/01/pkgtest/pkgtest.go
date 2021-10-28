package pkgtest

import (
	f "fmt"
)

// init function will be automatically executed
func init() {
	f.Println("Initial Package Test")
}

func PackageTest() {
	f.Println("Do Package Test")
}
