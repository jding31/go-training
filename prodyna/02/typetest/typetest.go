package typetest

import (
	"fmt"
	"unsafe"
)

func BooleanTest() {
	var b1 bool // Default is false
	var b2 bool
	fmt.Println(b1)
	fmt.Println(b2 == b1)

	var b3 = true
	b4 := false
	fmt.Println(!b3 == b4)

	var b5 *bool
	b6 := new(bool) // Default is false
	fmt.Println(b5)
	fmt.Println(b6)
	fmt.Println(*b6)
	fmt.Println(b5 == b6)
}

func StringTest() {
	var s1 = "abcde"
	fmt.Println("------------")
	for _, r := range s1 {
		fmt.Printf("%c", r)
	}
	fmt.Println("")
	s1 += " xyz"
	for _, r := range s1 {
		fmt.Print(string(r))
	}
	fmt.Println("")
}

func NumberTest() {
	var n1 int8 = 1
	var n2 int16 = 1
	// fmt.Println(n1 == n2) Invalid operation: n1 == n2 (mismatched types int8 and int16)
	fmt.Printf("n1: value %d; %T, %d \n", n1, n1, unsafe.Sizeof(n1))
	fmt.Printf("n2: value %d; %T, %d \n", n1, n2, unsafe.Sizeof(n2))

	var f1 float32 // default value is 0
	f2 := 3.141592653589793 // float64
	f3 := 3.14159265358979323846
	fmt.Printf("f1: value %f; %T, %d \n", f1, f1, unsafe.Sizeof(n1))
	fmt.Printf("f1: value %g; %T, %d \n", f1, f1, unsafe.Sizeof(n1))
	fmt.Printf("f2: value %f; %T, %d \n", f2, f2, unsafe.Sizeof(n2))
	fmt.Printf("f2: value %g; %T, %d \n", f2, f2, unsafe.Sizeof(n2))
	fmt.Println(f2) //3.141592653589793
	fmt.Println(f3) //3.141592653589793 loss of precision

}

const (
	Pi = 3.14159
)

const (
	Monday Day = iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

func IotaTest() {
	//var i1 = iota // iota can only be used in const block
	fmt.Printf("The value of Monday is %v\n", Monday)
	fmt.Printf("The value of Tuesday is %v\n", Tuesday)
	fmt.Printf("The value of Wednesday is %v\n", Wednesday)
	fmt.Printf("The value of Thursday is %v\n", Thursday)
	fmt.Printf("The value of Friday is %v\n", Friday)
	fmt.Printf("The value of Saturday is %v\n", Saturday)
	fmt.Printf("The value of Sunday is %v\n", Sunday)
}

func TypeTest() {
	fmt.Println(GetWeekDay(0))
	fmt.Println(GetWeekDay(1))
	fmt.Println(GetWeekDay(2))
	fmt.Println(GetWeekDay(3))
	fmt.Println(GetWeekDay(4))
	fmt.Println(GetWeekDay(5))
	fmt.Println(GetWeekDay(6))
	fmt.Println(GetWeekDay(99))

}

func GetWeekDay(d Day) string {
	return [...]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}[d%6]
}

type Day int
