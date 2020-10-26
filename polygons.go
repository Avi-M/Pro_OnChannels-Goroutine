package main

import "fmt"

type Polygons interface {
	Perimeter()
}
type Object interface {
	NoOfSides()
}
type Pentagon int

func (p Pentagon) Perimeter() {
	fmt.Println("perimeter of Pentagon", 5*p)

}
func (p Pentagon) NoOfSides() {
	fmt.Println("pentagon has 5 sides")

}
func main() {
	var p Polygons
	p = Pentagon(5)
	p.Perimeter()
	var o Pentagon = p.(Pentagon)
	o.NoOfSides()

	var obj Object = Pentagon(5)
	obj.NoOfSides()
	var pent Pentagon = obj.(Pentagon)
	pent.Perimeter()
}
