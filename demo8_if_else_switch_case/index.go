package main

import (
	"fmt"
)

func main() {
	// someValue := 10
	// if someValue == 10 {
	// 	fmt.Println("== 10")
	// }else{
	// 	fmt.Println("!= 10")
	// }

	// if someValue > 10 || someValue < 2 {
	// 	fmt.Println(" someValue == 10 || someValue <2 ")
	// } else {
	// 	fmt.Println("not")
	// }

	if result := doSomething(); result == "ok" {
		fmt.Println("OK")
	} else {
		fmt.Println("Nok")
	}

	fnSwitchCase();
}

func doSomething() string {
	return "Nok"
}

func fnSwitchCase() {
	index := 3
	switch index {
	case 0:
		fmt.Println("0")
		break
	case 1:
		fmt.Println("1")
		break
	case 2:
		fmt.Println("2")
		break
	default:
		fmt.Println("something else")
		break
	}
}
