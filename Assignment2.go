package main

import (
	"fmt"
)

type Employee struct {
	EmployeeName   string
	EmployeeAge    int
	EmployeeSalary int
}

type Department struct {
	DepartmentName string
	EmployeeList   map[string]*Employee
}

// adding Employee
func (d *Department) AddEmployee(e Employee) {
	//initialize map
	if d.EmployeeList == nil {
		d.EmployeeList = make(map[string]*Employee)
	}
	d.EmployeeList[e.EmployeeName] = &e
}

// Removing Employee
func (d *Department) RemoveEmployee(employeeName string) {
	delete(d.EmployeeList, employeeName)
}

// Calculating Avg Salary
func (d *Department) CalculateAvgSalary() int {
	if len(d.EmployeeList) == 0 {
		return 0
	}
	var totalSalary int
	for _, em := range d.EmployeeList {
		totalSalary += em.EmployeeSalary
	}
	return totalSalary / int(len(d.EmployeeList))

}

//Give raise

func (d *Department) GiveRaise(employeeName string, percentage int) {
	if emp, e := d.EmployeeList[employeeName]; e {
		emp.EmployeeSalary += emp.EmployeeSalary * (percentage / 100)
	}
}

func (d *Department) ListEmployee() {
	for _, employe := range d.EmployeeList {
		fmt.Println(employe.EmployeeName)
	}
}

func main() {

	deptIt := Department{DepartmentName: "IT"}

	Employee1 := Employee{EmployeeName: "Sharanu", EmployeeAge: 22, EmployeeSalary: 11000}
	Employee2 := Employee{EmployeeName: "Ketan", EmployeeAge: 23, EmployeeSalary: 12000}
	Employee3 := Employee{EmployeeName: "Shiv", EmployeeAge: 34, EmployeeSalary: 15000}

	deptIt.AddEmployee(Employee1)
	deptIt.AddEmployee(Employee2)
	deptIt.AddEmployee(Employee3)

	deptIt.ListEmployee()

	deptIt.RemoveEmployee("Sharanu")
	fmt.Println("Removed_name Sharanu...........")
	deptIt.ListEmployee()

	fmt.Println("Average Salary: ", deptIt.CalculateAvgSalary())

	deptIt.GiveRaise("Ketan", 10)

	

}

