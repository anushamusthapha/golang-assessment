package services_test

import (
	"golang-assessment/models"
	repository "golang-assessment/respository"

	loggerNew "golang-assessment/logger"
	"golang-assessment/services"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupTestLogger() {
	// Initialize a buffer to capture log output
	logger := logrus.New()
	// Replace the global logger with the mock logger
	loggerNew.Log = logger
}
func setupTestDB(t *testing.T) *gorm.DB {
	// Connect to your test database
	dsn := "host=localhost user=postgres password=root123 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("error opening database: %v", err)
	}
	return db
}

func TestEmployeeService_CreateEmployee(t *testing.T) {
	// Setup
	setupTestLogger()
	db := setupTestDB(t)
	repo := repository.NewEmployeeRepository(db)
	service := services.NewEmployeeService(repo)

	// Test case: Valid employee creation
	t.Run("TestCreateEmployee_ValidData", func(t *testing.T) {
		expectedEmployee := models.Employee{Name: "John Doe", Position: "Software Engineer", Salary: 50000}

		createdEmployee := service.CreateEmployee(expectedEmployee.Name, expectedEmployee.Position, expectedEmployee.Salary)

		// Assert the created employee
		assert.Equal(t, expectedEmployee.Name, createdEmployee.Name)
		assert.Equal(t, expectedEmployee.Position, createdEmployee.Position)
		assert.Equal(t, expectedEmployee.Salary, createdEmployee.Salary)
	})
}

func TestEmployeeService_GetEmployeeByID(t *testing.T) {
	// Setup
	setupTestLogger()
	db := setupTestDB(t)
	repo := repository.NewEmployeeRepository(db)
	service := services.NewEmployeeService(repo)

	// Test case: Valid employee ID
	t.Run("TestGetEmployeeByID_ValidID", func(t *testing.T) {
		expectedEmployee := models.Employee{ID: 15, Name: "John Doe", Position: "Developer", Salary: 60000}

		employee, err := service.GetEmployeeByID(15)

		// Assert the retrieved employee
		assert.Nil(t, err)
		assert.Equal(t, expectedEmployee.ID, employee.ID)
		assert.Equal(t, expectedEmployee.Name, employee.Name)
		assert.Equal(t, expectedEmployee.Position, employee.Position)
		assert.Equal(t, expectedEmployee.Salary, employee.Salary)
	})

	// Test case: Invalid employee ID
	t.Run("TestGetEmployeeByID_InvalidID", func(t *testing.T) {
		_, err := service.GetEmployeeByID(100)

		// Assert that an error occurred due to invalid ID
		assert.NotNil(t, err)
	})
}

func TestEmployeeService_UpdateEmployee(t *testing.T) {
	setupTestLogger()
	db := setupTestDB(t)
	repo := repository.NewEmployeeRepository(db)
	service := services.NewEmployeeService(repo)
	// Test case: Valid employee update
	t.Run("TestUpdateEmployee_ValidData", func(t *testing.T) {
		expectedEmployee := models.Employee{ID: 16, Name: "Updated Name", Position: "Updated Position", Salary: 60000}

		updatedEmployee, err := service.UpdateEmployee(16, expectedEmployee.Name, expectedEmployee.Position, expectedEmployee.Salary)

		// Assert the updated employee
		assert.Nil(t, err)
		assert.Equal(t, expectedEmployee.ID, updatedEmployee.ID)
		assert.Equal(t, expectedEmployee.Name, updatedEmployee.Name)
		assert.Equal(t, expectedEmployee.Position, updatedEmployee.Position)
		assert.Equal(t, expectedEmployee.Salary, updatedEmployee.Salary)
	})

	// Test case: Error updating employee
	t.Run("TestUpdateEmployee_Error", func(t *testing.T) {
		updatedEmployee, err := service.UpdateEmployee(90000, "Jane Doe", "Manager", 60000)

		// Assert that an error occurred during update
		assert.NotNil(t, err)
		assert.Equal(t, models.Employee{}, updatedEmployee)
	})
}

func TestEmployeeService_DeleteEmployee(t *testing.T) {
	setupTestLogger()
	db := setupTestDB(t)
	repo := repository.NewEmployeeRepository(db)
	service := services.NewEmployeeService(repo)

	// Test case: Valid employee deletion
	t.Run("TestDeleteEmployee_ValidID", func(t *testing.T) {
		err := service.DeleteEmployee(8)

		// Assert no error occurred during deletion
		assert.Nil(t, err)
	})

	// Test case: Error deleting employee
	t.Run("TestDeleteEmployee_Error", func(t *testing.T) {

		err := service.DeleteEmployee(1)

		// Assert that an error occurred during deletion
		assert.NotNil(t, err)
	})
}

func TestEmployeeService_ListEmployees(t *testing.T) {
	setupTestLogger()
	db := setupTestDB(t)
	repo := repository.NewEmployeeRepository(db)
	service := services.NewEmployeeService(repo)

	// Test case: Valid list of employees
	t.Run("TestListEmployees_ValidData", func(t *testing.T) {
		expectedEmployees := []models.Employee{
			{ID: 1, Name: "John Doe", Position: "Software Engineer", Salary: 50000},
			{ID: 2, Name: "Jane Doe", Position: "Manager", Salary: 60000},
			// Add more expected employees if needed
		}

		employees, err := service.ListEmployees(1, 10)

		// Assert the list of employees
		assert.Nil(t, err)
		assert.NotNil(t, expectedEmployees, employees)

	})
}
