package main

import (
	"fmt"
)

func main() {
	var p1 product
	p1.name = "Arduino"
	p1.price = 100
	p1.stock = 20
	show(p1)
	update(&p1)

	show(p1) //p1 after updated

	// p1 = p1.setDiscount(1)
	p1 = p1.clear()
	show(p1)
}

type product struct { //สร้าง struct ใหม่ที่มีหลาย type เก็บอยู่ ด้วย type ที่เราสร้าง
	name  string
	price int
	stock int
}

func (p product) clear() product { // func (ตัวแปร struct) funcName() returnValue/*optional*/ {}
	p.price = 0
	p.stock = 0
	return p
}
func (p product) setDiscount(d int) product {
	p.price = p.price - d
	return p
}

func show(p product) {
	fmt.Println(p)
}

func update(p *product) {
	p.price = p.price + 100
	p.stock = 10
	// fmt.Println(*p)
}
