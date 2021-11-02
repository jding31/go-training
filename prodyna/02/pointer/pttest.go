package pointer

import "fmt"

func PtTest() {
	var p1 *string
	var p2 *string
	var p3 *string
	var p4 **string
	var s1 = "string 1"
	s2 := "string 2"
	p1 = &s1
	p2 = &s2
	p4 = &p2
	strPointerPrint(p1)
	strPointerPrint(p2)
	fmt.Println(p3)
	fmt.Println(p4)
	fmt.Println(*p4)
	fmt.Println(**p4)

	var p5 = new(string)
	p5 = p2
	strPointerPrint(p5)

}

func strPointerPrint(sp *string) {
	fmt.Println(sp)
	fmt.Println(*sp)
}
