package main

import (
	"demo15/employee"
) //สร้าง mod เพื่อให้อ้างถึงได้package ได้

func main() {//มองว่า Init คือ constructor
	e := employee.Init(
		"Lek",
		"Adolf",
		30,
		20,
	)
	e.LeavesRemaining()

	e = employee.Init(
		"Kan",
		"Adolf",
		30,
		20,
	)
	e.LeavesRemaining()
	//return เฉพาะค่าแรกที่ถูกสร้างเสมอด้วย pointer
}
