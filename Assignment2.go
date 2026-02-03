/* Task 2:-
Develop a system to manage employees and their departments.
Each employee has a name, age, and salary. Each department
has a name, a list of employees, and a method to calculate
the average salary of its employees. Additionally, create
methods to add and remove employees from departments and to give a raise to an employee.*/


package main

import (
	"fmt"
)

type Employee struct {
	Id     int
	Name   string
	Age    int
	Salary int
}

type Department struct {
	DepartmentName string
	EmployeeList   map[int]*Employee
}

// adding Employee
func (d *Department) AddEmployee(e Employee) {
	//initialize map
	if d.EmployeeList == nil {
		d.EmployeeList = make(map[int]*Employee)
	}
	d.EmployeeList[e.Id] = &e
}

// Removing Employee
func (d *Department) RemoveEmployee(Id int) {
	delete(d.EmployeeList, Id)
}

// Calculating Avg Salary
func (d *Department) CalculateAvgSalary() int {
	if len(d.EmployeeList) == 0 {
		return 0
	}
	var totalSalary int
	for _, em := range d.EmployeeList {
		totalSalary += em.Salary
	}
	return totalSalary / int(len(d.EmployeeList))

}

// Give raise
func (d *Department) GiveRaise(EmployeeId int, percentage int) {
	if emp, exists := d.EmployeeList[EmployeeId]; exists {
		raiseAmount := float64(emp.Salary) * (float64(percentage) / 100.0)
		emp.Salary += int(raiseAmount)
		d.EmployeeList[EmployeeId] = emp
	}
}

// Displaying list of eployee
func (d *Department) ListEmployee() {
	for _, employe := range d.EmployeeList {
		fmt.Println("ID: ", employe.Id, "Name: ", employe.Name, "||", "Age: ", employe.Age, "||", "Salary: ", employe.Salary)
	}
}

func main() {

	deptIt := Department{DepartmentName: "IT"}
	fmt.Println("Department: ", deptIt.DepartmentName)

	Employee1 := Employee{Id: 102, Name: "Sharanu", Age: 22, Salary: 11000}
	Employee2 := Employee{Id: 103, Name: "Ketan", Age: 23, Salary: 12000}
	Employee3 := Employee{Id: 104, Name: "Shiv", Age: 34, Salary: 15000}

	deptIt.AddEmployee(Employee1)
	deptIt.AddEmployee(Employee2)
	deptIt.AddEmployee(Employee3)

	deptIt.ListEmployee()

	fmt.Println("------------------------------------------------------------")

	deptIt.RemoveEmployee(Employee2.Id)
	fmt.Printf("Removed_Employee_id %d from Department IT.............\n", Employee2.Id)
	deptIt.ListEmployee()

	fmt.Println("-----------------------------------------------------------------")

	fmt.Println("Average Salary of Employees: ", deptIt.CalculateAvgSalary())

	fmt.Println("--------------------------------------------------------------")

	initialSalary := Employee1.Salary
	empId := Employee1.Id
	deptIt.EmployeeList[empId] = &Employee{
		Id:     empId,
		Salary: initialSalary,
	}

	fmt.Printf("Before raise for Employee Id:%d Salary is %d\n", empId, deptIt.EmployeeList[empId].Salary)

	raisePercent := 10
	deptIt.GiveRaise(empId, raisePercent)

	finalSalary := deptIt.EmployeeList[empId].Salary

	fmt.Printf("After %d%% raise for Employee Id: %d: Salary is %d \n", raisePercent, empId, finalSalary)

}



/*
Output :-

Department:  IT
ID:  102 Name:  Sharanu || Age:  22 || Salary:  11000
ID:  103 Name:  Ketan || Age:  23 || Salary:  12000
ID:  104 Name:  Shiv || Age:  34 || Salary:  15000
------------------------------------------------------------     
Removed_Employee_id 103 from Department IT.............
ID:  102 Name:  Sharanu || Age:  22 || Salary:  11000
ID:  104 Name:  Shiv || Age:  34 || Salary:  15000
-----------------------------------------------------------------
Average Salary of Employees:  13000
--------------------------------------------------------------
Before raise for Employee Id:102 Salary is 11000
After 10% raise for Employee Id: 102: Salary is 12100

*/


	






