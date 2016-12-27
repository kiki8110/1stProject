// Example
package main

import (
	"fmt"
)

func main() {
	var b1 byte = 49
	var b2 byte = 061
	var b3 byte = 0x31
	var b4 byte = '1'
	fmt.Println(b1)
	fmt.Println(b2)
	fmt.Println(b3)
	fmt.Println(b4)

	var r1 rune = '한'
	var r2 rune = '\ud55c'
	fmt.Println(r1)
	fmt.Println(r2)

}
