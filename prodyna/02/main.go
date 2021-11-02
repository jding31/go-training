package main

import (
	"fmt"
	pt "github.com/jding31/go-training/prodyna/02/pointer"
	tp "github.com/jding31/go-training/prodyna/02/typetest"
)

func main() {
	var s1 = "string"
	s2 := "string"
	var s3 string
	s3 = "test"
	var n1 = 5
	n2 := 5
	var n3 int // initialized with 0
	var n4 = 0xFFFF
	n5 := 65321
	fmt.Println(s1)
	fmt.Println(s1 == s2)
	fmt.Println(s3)
	fmt.Println(&s1)
	fmt.Println(&s2)
	fmt.Println(n1 == n2)
	fmt.Println(n3)
	fmt.Println(n4)
	fmt.Printf("%x\n", n5)

	fmt.Println("-----------")
	pt.PtTest()
	fmt.Println("-----------")
	tp.BooleanTest()
	tp.StringTest()
	tp.NumberTest()
	fmt.Println(tp.Pi)
	tp.IotaTest()
	tp.TypeTest()
}



