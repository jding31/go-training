package main

/*
	go mod init xxxx/yyyy
		|- main.go -> package "main"
 		|- aaaa
			|- helper.go -> package "aaaa"
		|- bbbb
			|- cccc
				|- util.go -> package "cccc"
		|- dddd
			|- prtfunc.go -> package "dddd"
			|- eeee
				|- prtfunc2.go -> package "dddd"
	If you want to import aaaa package in main.go, you need to import like xxxx/yyyy/aaaa
	If you want to import cccc package in main.go, you need to import like xxxx/yyyy/bbbb/cccc
	Otherwise you will get package is not in GOROOT error message

	prtfunc.go and prtfunc2 has the same package declaration, but in different path
	You can not import two packages with the same name:
	xxxx/yyyy/dddd
	xxxx/yyyy/dddd/eeee
	You should give them different name
	d xxxx/yyyy/dddd
	e xxxx/yyyy/dddd/eeee
 */
import (
	f "fmt"
	"github.com/jding31/go-training/prodyna/01/pkgtest"
	p "github.com/jding31/go-training/prodyna/01/pkgtest/test3"
	_ "github.com/jding31/go-training/prodyna/01/random"

	// "prodyna/01/pkgtest"
)

func main() {
	f.Println("Hello 世界")
	pkgtest.PackageTest()
	pkgtest.PackageTest2()
	p.PackageTest3()
}
