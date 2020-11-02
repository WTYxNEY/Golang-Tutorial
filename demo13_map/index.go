package main

import (
	"fmt"
)

func main() {
	var numbers = map[string]int{"one": 1, "two": 2, "three": 3} //[key of type]key of value => key ซ้ำไม่ได้
	fmt.Println(numbers["two"])

	var colors = make(map[string]string)//make map to dynamic
	colors["red"] = "#f00";
	colors["green"] = "#f0";
	colors["blue"] = "#00f";
	fmt.Println("",colors)
	fmt.Println("",colors["green"])

	var courses = make(map[string]map[string]int)
	courses["Andriod"] = make(map[string]int)//ต้อง make ค่ามาก่อนเสมอ
	courses["Andriod"]["price"] = 200
	courses["Andriod"]["code"] = 1234

	courses["IOS"] = make(map[string]int)
	courses["IOS"]["price"] = 100
	courses["IOS"]["code"] = 4567
	fmt.Println(courses)
	fmt.Println(courses["IOS"]["code"])
}
