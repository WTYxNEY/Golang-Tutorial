package main

import (
	"fmt"
)

func main() {
	msg := "some massage";
	var msgPointer * string = &msg; //pointer เก็บค่าไม่ได้ เก็บได้แต่ตำแหน่ง ถูกใช้สำหรับ performance tuning
	fmt.Println(msg)
	fmt.Println(*msgPointer)//ได้ตำแหน่งฐาน 16 ใส่ * เพื่อดึง value มาแสดง

	changeMassage(&msg, "New Massage 1")
	fmt.Println(msg)

	changeMassage(msgPointer, "New Massage 2")
	fmt.Println(msg)
	
	changeMassage(msgPointer, "New Massage 3")
	fmt.Println(*msgPointer)
}

func changeMassage(aPointer *string, newMassage string){// ส่งตำแหน่งเข้ามายัง function
	*aPointer = newMassage
}