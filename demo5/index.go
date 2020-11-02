package main

import (
	"fmt"
)

func main() {
	fn1()
	fn2("Goodmorning")
	fn3("Goodmorning", 1)
	const a, b = 2, 3
	fmt.Printf("%d+%d=%d\n", a, b, sum(a, b)) //มี return

	var x, y int = swap(5, 6) //shrotform => x,y := swap(a,b)
	fmt.Printf("%d, %d => %d,%d\n", 5, 6, x, y)

	_c,_d,title := swapV2(10, 20)// ไม่ใส่ : เพราะไม่ได้ประกาศตัวแปรใหม่ แค่แทนค่าเฉยๆ
	fmt.Printf("%d, %d => %d,%d,%s\n", 10, 20, _c, _d,title)
}

func fn1() {
	fmt.Println("CodeNey")
}
func fn2(msg string) {
	fmt.Println(msg)
}
func fn3(title string, version int) {
	fmt.Print(title)
	fmt.Println(version)
}

func sum(a int, b int) int {
	return a + b
}

func swap(a int, b int) (int, int) { //input type & return type
	return b, a
}

func swapV2(a, b int) (x, y int,title string) { //input type & return type
	y = a
	x = b
	title = "Result"
	return x, y,title
}
