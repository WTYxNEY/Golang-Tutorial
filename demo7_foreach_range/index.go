package main

import (
	"fmt"
)

func main() {
	courses := []string{"Andriod", "IOS", "React"}

	for index, item := range courses {
		fmt.Printf("%d. %s\n", index+1, item)
	}
	for _, item := range courses { // ถ้าไม่ใช้ก็ใส่ _
		fmt.Printf(" %s\n",  item)
	}
}
