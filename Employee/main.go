package main

import "fmt"

// Employee interface defines a method to get employee details
type Employee interface {
	GetDetails() string
	GetID() uint64 // Добавим метод для получения ID, чтобы использовать его в качестве ключа
}

// Full-time employee
type FullTimeEmployee struct {
	ID     uint64
	Name   string
	Salary uint32
}

func (f FullTimeEmployee) GetDetails() string {
	return fmt.Sprintf("ID: %d, Name: %s, Salary: %d Tenge", f.ID, f.Name, f.Salary)
}

func (f FullTimeEmployee) GetID() uint64 {
	return f.ID
}

// Part-time employee
type PartTimeEmployee struct {
	ID          uint64
	Name        string
	HourlyRate  uint64
	HoursWorked float32
}

func (p PartTimeEmployee) GetDetails() string {
	return fmt.Sprintf("ID: %d, Name: %s, Hourly Rate: %d Tenge, Hours Worked: %.1f", p.ID, p.Name, p.HourlyRate, p.HoursWorked)
}

func (p PartTimeEmployee) GetID() uint64 {
	return p.ID
}

// Company structure
type Company struct {
	employees map[uint64]Employee // Используем uint64 в качестве ключа
}

func NewCompany() *Company {
	return &Company{
		employees: make(map[uint64]Employee),
	}
}

// Adding employee to company
func (c *Company) AddEmployee(emp Employee) {
	c.employees[emp.GetID()] = emp
}

// List of all employees
func (c *Company) ListEmployees() {
	fmt.Println("Список сотрудников:")
	for _, emp := range c.employees {
		fmt.Println(emp.GetDetails())
	}
}

func main() {
	c := NewCompany()

	c.AddEmployee(FullTimeEmployee{ID: 1, Name: "Alice", Salary: 5000})
	c.AddEmployee(PartTimeEmployee{ID: 2, Name: "Bob", HourlyRate: 20, HoursWorked: 100})

	c.ListEmployees()
}
