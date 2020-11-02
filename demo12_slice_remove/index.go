package main

import (
	"fmt"
)

func main() {
 var numbers = []int{1,2,3,4,5,6,7,8}
 showSlice(numbers)

 // leading remove
 numbers = numbers[1 : len(numbers)]//[start, end] start ไม่เอาที่ตัว , ให้เหลือถึงตัวไหน
 showSlice(numbers)
//  numbers = numbers[1 : len(numbers)]//[start, end]
//  showSlice(numbers)

 //trailing remove
 numbers = numbers[2 : len(numbers)-1]//[0, len()-1] ไม่เอากี่ตัว , ให้เหลือถึงตัวไหน
 showSlice(numbers)

 //remove specific index
 numbers = removeIndex(numbers, 2)
 showSlice(numbers)
}

func showSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x) //len ปัจจุุบันมีข้อมูลเท่าไรใน array, cap ดูว่า array เราจุได้เท่าไร
}

func removeIndex(s []int, index int) []int{
	return append(s[:index], s[index+1:]...)//return ค่าใหม่ที่ไม่รวม index ที่ต้องการ
}