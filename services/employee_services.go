package services

import (
	"golang-assessment/models"
	repository "golang-assessment/respository"
)

type EmployeeService struct {
	repository *repository.EmployeeRepository
}

func NewEmployeeService(repository *repository.EmployeeRepository) *EmployeeService {
	return &EmployeeService{repository: repository}
}

func (s *EmployeeService) CreateEmployee(name, position string, salary float64) models.Employee {
	employee := models.Employee{Name: name, Position: position, Salary: salary}
	s.repository.CreateEmployee(&employee)
	return employee
}

func (s *EmployeeService) GetEmployeeByID(id int) (models.Employee, error) {
	return s.repository.GetEmployeeByID(id)
}

func (s *EmployeeService) UpdateEmployee(id int, name, position string, salary float64) (models.Employee, error) {
	employee, err := s.repository.GetEmployeeByID(id)
	if err != nil {
		return models.Employee{}, err
	}
	employee.Name = name
	employee.Position = position
	employee.Salary = salary
	err = s.repository.UpdateEmployee(&employee)
	return employee, err
}

func (s *EmployeeService) DeleteEmployee(id int) error {
	return s.repository.DeleteEmployee(id)
}

func (s *EmployeeService) ListEmployees(page, limit int) ([]models.Employee, error) {
	offset := (page - 1) * limit
	return s.repository.ListEmployee(offset, limit)
}
