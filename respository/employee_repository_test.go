package repository

import (
	"testing"

	loggerNew "golang-assessment/logger"
	"golang-assessment/models"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	// Connect to your test database
	dsn := "host=localhost user=postgres password=root123 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("error opening database: %v", err)
	}
	return db
}
func setupTestLogger(t *testing.T) {
	// Initialize a buffer to capture log output
	logger := logrus.New()
	// Replace the global logger with the mock logger
	loggerNew.Log = logger
}

func TestEmployeeRepository(t *testing.T) {
	// Setup
	setupTestLogger(t)
	db := setupTestDB(t)
	repo := NewEmployeeRepository(db)

	t.Run("TestCreateEmployee", func(t *testing.T) {
		// Test CreateEmployee function
		employee := &models.Employee{Name: "John Doe", Position: "Software Engineer", Salary: 50000}
		repo.CreateEmployee(employee)
	})

	t.Run("TestGetEmployeeByID", func(t *testing.T) {
		// Test GetEmployeeByID function
		id := 11
		employee, err := repo.GetEmployeeByID(id)
		expectedResponse := models.Employee(models.Employee{ID: 11, Name: "John Doe", Position: "Developer", Salary: 60000})
		assert.Nil(t, err)
		assert.Equal(t, employee, expectedResponse)
	})

	t.Run("TestUpdateEmployee", func(t *testing.T) {
		// Test UpdateEmployee function
		employee := &models.Employee{ID: 11, Name: "John Doe", Position: "Developer", Salary: 60000}
		err := repo.UpdateEmployee(employee)
		assert.Nil(t, err)
	})

	t.Run("TestDeleteEmployee", func(t *testing.T) {
		// Test DeleteEmployee function
		id := 11
		err := repo.DeleteEmployee(id)
		assert.Nil(t, err)
	})

}
