package main

import (
	"fmt"
)

func main() {
	var numbers1 = make([]int, 0, 5) //3 len, 5 cap => ถ้าแอดค่าเกินค่าเพิ่ม cap ไป 1 เท่าตัว
	numbers1 = append(numbers1, 1)
	numbers1 = append(numbers1, 2)
	showSlice(numbers1)

	var numbers2 []int//ตัวแปรแบบ slice(array ที่ไม่ระบุ sizeถ)
	numbers2 = append(numbers2, 1)
	numbers2 = append(numbers2, 2)
	showSlice(numbers2)
}

func showSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x) //len ปัจจุุบันมีข้อมูลเท่าไรใน array, cap ดูว่า array เราจุได้เท่าไร
}
