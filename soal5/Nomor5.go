package main

import "fmt"

type Employee interface {
	setSalary(salary int)
	getSalary() int
	setGrade(grade string)
	getGrade() string
	label()
}

type Engineer struct {
	salary int
	grade  string
}

func (e *Engineer) setSalary(salary int) {
	e.salary = salary
}

func (e *Engineer) getSalary() int {
	return e.salary
}

func (e *Engineer) setGrade(grade string) {
	e.grade = grade
}

func (e *Engineer) getGrade() string {
	return e.grade
}

func (e *Engineer) label() {
	fmt.Println("Employee's data:")
}

type Manager struct {
	salary int
	grade  string
}

func (m *Manager) setSalary(salary int) {
	m.salary = salary
}

func (m *Manager) getSalary() int {
	return m.salary
}

func (m *Manager) setGrade(grade string) {
	m.grade = grade
}

func (m *Manager) getGrade() string {
	return m.grade
}

func (m *Manager) label() {
	fmt.Println("Employee's data:")
}

func main() {
	engineer := &Engineer{}
	engineer.setSalary(50000)
	engineer.setGrade("A")

	engineer.label()
	fmt.Println("Salary:", engineer.getSalary())
	fmt.Println("Grade:", engineer.getGrade())

	manager := &Manager{}
	manager.setSalary(70000)
	manager.setGrade("B")

	manager.label()
	fmt.Println("Salary:", manager.getSalary())
	fmt.Println("Grade:", manager.getGrade())
}
