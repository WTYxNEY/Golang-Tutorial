package main

import (
	"demo15/employee"
) //สร้าง mod เพื่อให้อ้างถึงได้package ได้

func main() {//มองว่า new คือ constructor
	e := employee.New(
		"Sam",
		"Adolf",
		30,
		20,
	)

	e.LeavesRemaining()
}
