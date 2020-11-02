package main

import (
	"fmt"
)

var count int = 0

func main() {
	fmt.Println("Begin")

	var tmp1 int = 0
	tmp1 = 10
	var tmp2 string = "hello"
	var tmp3 bool = true
	// const tmp4 int = 0;

	//implicit Declaration => ไม่ต้องใส่ type
	tmp5 := 0
	tmp6 := "Watunyu"

	fmt.Println(tmp1);
	fmt.Println(tmp2);
	fmt.Println(tmp3);
	fmt.Println(tmp5);
	fmt.Println(tmp6);
	run();
	
}

func run()  {
	count++
	fmt.Println(count);
}