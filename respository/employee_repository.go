package repository

import (
	"fmt"
	"golang-assessment/logger"
	"golang-assessment/models"
	"sync"

	"gorm.io/gorm"
)

type EmployeeRepository struct {
	db *gorm.DB
	mu sync.Mutex
}

func NewEmployeeRepository(db *gorm.DB) *EmployeeRepository {
	return &EmployeeRepository{db: db}
}

func (r *EmployeeRepository) CreateEmployee(employee *models.Employee) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if err := r.db.Create(employee).Error; err != nil {
		logger.Log.Errorf("Error creating employee: %v", err)
	} else {
		logger.Log.Infof("Employee created: %v", employee)
	}
}

func (r *EmployeeRepository) GetEmployeeByID(id int) (models.Employee, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	var employee models.Employee

	if err := r.db.First(&employee, id).Error; err != nil {
		logger.Log.Errorf("Error retreiving employee by ID %d:%v", id, err)
		return models.Employee{}, err
	}

	logger.Log.Infof("Retrieved employee:%v", employee)
	return employee, nil
}

func (r *EmployeeRepository) UpdateEmployee(employee *models.Employee) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if err := r.db.Save(employee).Error; err != nil {
		logger.Log.Errorf("Error updating employee :%v", err)
		return err
	}

	logger.Log.Infof("Employee updated : %v", employee)
	return nil
}

func (r *EmployeeRepository) DeleteEmployee(id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	result := r.db.Delete(&models.Employee{}, id)
	if result.Error != nil {
		logger.Log.Errorf("Error deleting employee by ID %d: %v", id, result.Error)
		return result.Error
	}
	if result.RowsAffected == 0 {
		// No rows were affected, indicating that the data with the provided ID is not present
		return fmt.Errorf("employee with ID %d not found", id)
	}

	logger.Log.Infof("Employee deleted with ID %d", id)
	return nil
}

func (r *EmployeeRepository) ListEmployee(offset, limit int) ([]models.Employee, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	var employee []models.Employee

	if err := r.db.Offset(offset).Limit(limit).Find(&employee).Error; err != nil {
		logger.Log.Errorf("Error listing employee: %v", err)
		return nil, err
	}

	logger.Log.Infof("Listed Employees :%v", employee)
	return employee, nil
}
