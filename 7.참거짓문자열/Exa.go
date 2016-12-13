// Exa
package main

import (
	"fmt"
)

func main() {
	var b1 bool = true

	fmt.Println(b1)

	b1 = 2 > 3 // 비교 연산의 결과
	fmt.Println(b1)

	b1 = true && false // 논리연산의 결과
	fmt.Println(b1)

	// 문자열
	var s1 string = "Hello, test"
	var s2 string = "안녕, 테스트"

	// 백쿼터(~자판에 있는 ')로 여러줄 표현
	var s3 = `안녕
	test는 테스트`

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)

}
