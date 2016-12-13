// Exa
package main

import (
	"fmt"
)

const errbound = 1e-14

func main() {
	var f1 float32 = 0.1 // 고정소수점
	var f2 float32 = .1
	var f3 float32 = 1e-1 // 부동소수점
	fmt.Println(f1)
	fmt.Println(f2)
	fmt.Println(f3)

	var fs1 = 0.0
	fs1 += 0.1
	fmt.Println(fs1)
	fs1 += 0.1
	fmt.Println(fs1)
	fs1 += 0.1
	fmt.Println(fs1)
	fs1 += 0.1
	fmt.Println(fs1)
	fs1 += 0.1
	fmt.Println(fs1)
	fs1 += 0.1
	fmt.Println(fs1)
	fs1 += 0.1
	fmt.Println(fs1)
	fs1 += 0.1
	fmt.Println(fs1)
	fs1 += 0.1
	fmt.Println(fs1)
	fs1 += 0.1
	fmt.Println(fs1)
	fmt.Println(fs1 == 1.0)
	fmt.Println((fs1 - 1.0) <= errbound)
}
