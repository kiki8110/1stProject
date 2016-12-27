// Exa
package main

import (
	"fmt"
)

func main() {
	var i1, i2, i3 int = 2, 2, 3
	fmt.Println(i1, "==", i2, " 결과", i1 == i2)
	fmt.Println(i1, "==", i3, " 결과", i1 == i3)

	var f1, f2, f3 float32 = 0.1, 0.1, 0.2
	fmt.Println(f1, "==", f2, " 결과", f1 == f2)
	fmt.Println(f1, "==", f3, " 결과", f1 == f3)

	var s1, s2, s3 string = "abc", "abc", "ABC"
	fmt.Println(s1, "==", s2, " 결과", s1 == s2)
	fmt.Println(s1, "==", s3, " 결과", s1 == s3)

	var hs1, hs2, hs3 string = "강감찬", "강감찬", "홍길동"
	fmt.Println(hs1, "==", hs2, " 결과", hs1 == hs2)
	fmt.Println(hs1, "==", hs3, " 결과", hs1 == hs3)

	var ss1, ss2, ss3 string = "abc", "zb", "ABC"
	fmt.Println("a의 아스키 코드:", 'a')
	fmt.Println("z의 아스키 코드:", 'z')
	fmt.Println("A의 아스키 코드:", 'A')
	fmt.Println(ss1, "<", ss2, " 결과", ss1 < ss2)
	fmt.Println(ss1, "<", ss3, " 결과", ss1 < ss3)

	var hss1, hss2, hss3 string = "강감찬", "킹콩", "홍길동"
	fmt.Println(hss1, "<", hss2, " 결과", hss1 < hss2)
	fmt.Println(hss1, "<", hss3, " 결과", hss1 < hss3)

	// 실수 형식의 오차로 인해 개발자의 표현한 값과 실제 값이 달라서 버그 발생
	var ff1 float32 = 0.99999999999999999999999999991
	var ff2 float32 = 0.99999999999999999999999999992
	fmt.Println(ff1 == ff2)
	fmt.Println("ff1:", ff1, " ff2:", ff2)
}
