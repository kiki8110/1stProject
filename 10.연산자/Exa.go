// Exa 연산자
package main

import (
	"fmt"
)

func main() {
	// 부호 연산자
	var i int = -3
	fmt.Println("i:", i)
	fmt.Println("+i:", +i)
	fmt.Println("-i:", -i)

	// 사직 연산자
	var i1 int = 2
	var i2 int = 3

	fmt.Println(i1, "+", i2, "=", i1+i2)
	fmt.Println(i1, "-", i2, "=", i1-i2)
	fmt.Println(i1, "*", i2, "=", i1*i2)

	var f1 float32 = 0.2
	var f2 float32 = 0.3
	fmt.Println(f1, "+", f2, "=", f1+f2)
	fmt.Println(f1, "-", f2, "=", f1-f2)
	fmt.Println(f1, "*", f2, "=", f1*f2)

	var c1 complex64 = 2 + 3i
	var c2 complex64 = 1 + 2i
	fmt.Println(c1, "+", c2, "=", c1+c2)
	fmt.Println(c1, "-", c2, "=", c1-c2)
	fmt.Println(c1, "*", c2, "=", c1*c2)

	var s1 string = "hello,"
	var s2 string = " test"
	fmt.Println(s1, "+", s2, "=", s1+s2)

	// error
	//	var ii1 int = 3
	//	var ff1 float32 = 1.0
	//	fmt.Println(ii1 + ff1)

	var ii1 int = 8
	var ii2 int = 3
	fmt.Println(ii1, "/", ii2, "=", ii1/ii2)

	var ff1 float32 = 3.2
	var ff2 float32 = 2.1
	fmt.Println(ff1, "/", ff2, "=", ff1/ff2)

	var test int = 1010
	fmt.Println(test, "+", ii1, "-", ii2, "*", ii1, "=", test+ii1-ii2*ii1)
}
