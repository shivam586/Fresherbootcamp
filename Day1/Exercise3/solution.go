package main

import (
	"fmt"
)

type EmployeeSalary interface {
	CalculateSalary() int
}

type FullTime struct {
	empId    int
	basicpay int
}

type Contract struct {
	empId    int
	basicpay int
}

type Freelancer struct {
	empId       int
	ratePerHour int
	totalHours  int
}


func (p FullTime) CalculateSalary() int {
	return p.basicpay*28
}


func (c Contract) CalculateSalary() int {
	return c.basicpay*28
}


func (f Freelancer) CalculateSalary() int {
	return f.ratePerHour * f.totalHours*28
}

func main() {
	emp1 := FullTime{
		empId:    1,
		basicpay: 500,
	}
	emp2 := FullTime{
		empId:    2,
		basicpay: 500,
	}
	emp3 := Contract{
		empId:    3,
		basicpay: 100,
	}
	emp4 := Freelancer{
		empId:       4,
		ratePerHour: 10,
		totalHours:  9,
	}
	emp5 := Freelancer{
		empId:       5,
		ratePerHour: 10,
		totalHours:  12,
	}
    var emp EmployeeSalary = emp1
    var empint EmployeeSalary = emp2
	var empint2 EmployeeSalary = emp3
	var empint3 EmployeeSalary = emp4
	var empint4 EmployeeSalary = emp5
    fmt.Println("Salary of Fulltime Employee with emplyeeid1:-",emp.CalculateSalary())
    fmt.Println("Salary of Fulltime Employee with emplyeeid2:-",empint.CalculateSalary())
	fmt.Println("Salary of Contract Employee with emplyeeid3:-",empint2.CalculateSalary())
	fmt.Println("Salary of Freelancer Employee with emplyeeid4",empint3.CalculateSalary())
	fmt.Println("Salary of Freelancer Employee with emplyeeid5",empint4.CalculateSalary())

}