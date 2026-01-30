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

func (d *Department) GiveRaise(EmployeeName string, percentage int) {
	if emp, exists := d.EmployeeList[EmployeeName]; exists {
		raiseAmount := float64(emp.EmployeeSalary) * (float64(percentage) / 100.0)
		emp.EmployeeSalary += int(raiseAmount)
		d.EmployeeList[EmployeeName] = emp
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

	fmt.Println("--------------------------------------------------")

	deptIt.ListEmployee()

	fmt.Println("------------------------------------------------------------")

	deptIt.RemoveEmployee("Sharanu")
	fmt.Println("Removed_name Sharanu...........")
	deptIt.ListEmployee()

	fmt.Println("-----------------------------------------------------------------")

	fmt.Println("Average Salary: ", deptIt.CalculateAvgSalary())

	fmt.Println("--------------------------------------------------------------")

	initialSalary := Employee1.EmployeeSalary
	empName := Employee1.EmployeeName
	deptIt.EmployeeList[empName] = &Employee{
		EmployeeName:   empName,
		EmployeeSalary: initialSalary,
	}

	fmt.Printf("Before raise for %s: Salary is %d\n", empName, deptIt.EmployeeList[empName].EmployeeSalary)

	raisePercent := 10
	deptIt.GiveRaise(empName, raisePercent)

	finalSalary := deptIt.EmployeeList[empName].EmployeeSalary

	fmt.Printf("After %d%% raise for %s: Salary is %d\n", raisePercent, empName, finalSalary)

	

}




