package main

import "fmt"

type Employee interface {
	printName(name string)
	printSalary(basic int, tax int) int
}

//user defind type
type Emp int

func (e Emp) printName(name string) {

	fmt.Println("employeeId", e)
	fmt.Println("employeeName", name)
}
func (e Emp) printSalary(basic int, tax int) int {
	var sal = basic * tax / 100
	return basic - sal

}
func main() {
	var e1 Employee
	e1 = Emp(1)
	e1.printName("Priya Banarji")
	fmt.Println(e1.printSalary(10000, 10))

}
