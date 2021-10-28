package main

/*
	go mod init xxxx/yyyy
		|- main.go -> package "main"
 		|- aaaa
			|- helper.go -> package "aaaa"
		|- bbbb
			|- cccc
				|- util.go -> package "cccc"
	If you want to import aaaa package in main.go, you need to import like xxxx/yyyy/aaaa
	If you want to import cccc package in main.go, you need to import like xxxx/yyyy/bbbb/cccc
	Otherwise you will get package is not in GOROOT error message
 */
import (
	f "fmt"
	"github.com/jding31/go-training/prodyna/01/pkgtest"
	_ "github.com/jding31/go-training/prodyna/01/random"
	// "prodyna/01/pkgtest"
)

func main() {
	f.Println("Hello 世界")
	pkgtest.PackageTest()
}
