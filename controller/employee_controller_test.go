package controller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	loggerNew "golang-assessment/logger"
	"golang-assessment/models"
	repository "golang-assessment/respository"
	"golang-assessment/services"
)

func setupTestLogger(t *testing.T) {
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
func TestCreateEmployee(t *testing.T) {
	// Setup
	setupTestLogger(t)
	db := setupTestDB(t)
	repo := repository.NewEmployeeRepository(db)
	service := services.NewEmployeeService(repo) // Create a real service instance
	controller := NewEmployeeController(service)

	// Test CreateEmployee
	t.Run("TestCreateEmployee", func(t *testing.T) {
		// Prepare request data
		employee := models.Employee{Name: "John Doe", Position: "Software Engineer", Salary: 50000}
		jsonStr, _ := json.Marshal(employee)
		req, _ := http.NewRequest("POST", "/employees", strings.NewReader(string(jsonStr)))
		req.Header.Set("Content-Type", "application/json")

		// Create a response recorder
		rr := httptest.NewRecorder()

		// Create a test context with the request and response recorder
		router := gin.Default()
		router.POST("/employees", controller.CreateEmployee)
		router.ServeHTTP(rr, req)

		// Assert response status code
		assert.Equal(t, http.StatusCreated, rr.Code)

		// Assert response body
		var createdEmployee models.Employee
		err := json.Unmarshal(rr.Body.Bytes(), &createdEmployee)
		assert.Nil(t, err)
		assert.Equal(t, employee.Name, createdEmployee.Name)
	})

}

func TestGetEmployeeByID(t *testing.T) {
	// Setup
	setupTestLogger(t)
	db := setupTestDB(t)
	repo := repository.NewEmployeeRepository(db)
	service := services.NewEmployeeService(repo)
	controller := NewEmployeeController(service)

	// Test case: Valid employee ID
	t.Run("TestGetEmployeeByID_ValidID", func(t *testing.T) {
		// Prepare request
		req, err := http.NewRequest("GET", "/employees/6", nil)
		assert.Nil(t, err)
		// Create a response recorder
		rr := httptest.NewRecorder()

		// Create a test context with the request and response recorder
		router := gin.Default()
		router.GET("/employees/:id", controller.GetEmployeeByID)
		router.ServeHTTP(rr, req)

		// Assert response status code
		assert.Equal(t, http.StatusOK, rr.Code)

		// Assert response body
		var actualEmployee models.Employee
		err = json.Unmarshal(rr.Body.Bytes(), &actualEmployee)
		assert.Nil(t, err)
		assert.Equal(t, 6, actualEmployee.ID) // Assuming the ID is returned in the response
	})

	t.Run("TestGetEmployeeByID_InvalidID", func(t *testing.T) {
		// Prepare request
		req, err := http.NewRequest("GET", "/employees/0", nil)
		assert.Nil(t, err)
		// Create a response recorder
		rr := httptest.NewRecorder()

		// Create a test context with the request and response recorder
		router := gin.Default()
		router.GET("/employees/:id", controller.GetEmployeeByID)
		router.ServeHTTP(rr, req)

		// Assert response status code
		assert.Equal(t, http.StatusNotFound, rr.Code)

	})
}

func TestUpdateEmployee(t *testing.T) {
	// Setup
	setupTestLogger(t)
	db := setupTestDB(t)
	repo := repository.NewEmployeeRepository(db)
	service := services.NewEmployeeService(repo)
	controller := NewEmployeeController(service)
	// Test case: Valid employee update
	t.Run("TestUpdateEmployee_ValidData", func(t *testing.T) {
		// Stub service method to return a hardcoded updated employee
		expectedEmployee := &models.Employee{ID: 6, Name: "Updated Name", Position: "Updated Position", Salary: 60000}

		// Prepare request data
		updatedEmployee := models.Employee{Name: "Updated Name", Position: "Updated Position", Salary: 60000}
		jsonStr, _ := json.Marshal(updatedEmployee)
		req, _ := http.NewRequest("PUT", "/employees/6", strings.NewReader(string(jsonStr)))
		req.Header.Set("Content-Type", "application/json")

		// Create a response recorder
		rr := httptest.NewRecorder()

		// Create a test context with the request and response recorder
		router := gin.Default()
		router.PUT("/employees/:id", controller.UpdateEmployee)
		router.ServeHTTP(rr, req)

		// Assert response status code
		assert.Equal(t, http.StatusOK, rr.Code)

		// Assert response body
		var actualEmployee models.Employee
		err := json.Unmarshal(rr.Body.Bytes(), &actualEmployee)
		assert.Nil(t, err)
		assert.Equal(t, expectedEmployee, &actualEmployee)
	})

	// Test case: Invalid request data
	t.Run("TestUpdateEmployee_InvalidData", func(t *testing.T) {
		// Prepare request data with invalid JSON
		req, _ := http.NewRequest("PUT", "/employees/1", strings.NewReader("invalid JSON"))
		req.Header.Set("Content-Type", "application/json")

		// Create a response recorder
		rr := httptest.NewRecorder()

		// Create a test context with the request and response recorder
		router := gin.Default()
		router.PUT("/employees/:id", controller.UpdateEmployee)
		router.ServeHTTP(rr, req)

		// Assert response status code
		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})
}

func TestDeleteEmployee(t *testing.T) {
	// Setup
	setupTestLogger(t)
	db := setupTestDB(t)
	repo := repository.NewEmployeeRepository(db)
	service := services.NewEmployeeService(repo)
	controller := NewEmployeeController(service)

	// Test case: Valid employee deletion
	t.Run("TestDeleteEmployee_ValidID", func(t *testing.T) {
		// Prepare request
		req, _ := http.NewRequest("DELETE", "/employees/11", nil)

		// Create a response recorder
		rr := httptest.NewRecorder()

		// Create a test context with the request and response recorder
		router := gin.Default()
		router.DELETE("/employees/:id", controller.DeleteEmployee)
		router.ServeHTTP(rr, req)

		// Assert response status code
		assert.Equal(t, http.StatusOK, rr.Code)

		// Assert response body
		expectedBody := gin.H{"data": "Successfully deleted the employee"}
		assertResponseBody(t, rr.Body.Bytes(), expectedBody)
	})

	// Test case: Invalid employee ID
	t.Run("TestDeleteEmployee_InvalidID", func(t *testing.T) {

		// Prepare request
		req, _ := http.NewRequest("DELETE", "/employees/invalid_id", nil)

		// Create a response recorder
		rr := httptest.NewRecorder()

		// Create a test context with the request and response recorder
		router := gin.Default()
		router.DELETE("/employees/:id", controller.DeleteEmployee)
		router.ServeHTTP(rr, req)

		// Assert response status code
		assert.Equal(t, http.StatusNotFound, rr.Code)

		// Assert response body
		expectedBody := gin.H{"error": "Employee not found"}
		assertResponseBody(t, rr.Body.Bytes(), expectedBody)
	})
}

// Helper function to assert response body
func assertResponseBody(t *testing.T, body []byte, expected gin.H) {
	var actual gin.H
	err := json.Unmarshal(body, &actual)
	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func TestListEmployees(t *testing.T) {
	// Setup
	setupTestLogger(t)
	db := setupTestDB(t)
	repo := repository.NewEmployeeRepository(db)
	service := services.NewEmployeeService(repo)
	controller := NewEmployeeController(service)

	// Test case: Valid list of employees
	t.Run("TestListEmployees_ValidData", func(t *testing.T) {
		// Prepare request
		req, _ := http.NewRequest("GET", "/employees", nil)
		req.URL.RawQuery = "page=1&limit=10" // Simulate query parameters

		// Create a response recorder
		rr := httptest.NewRecorder()

		// Create a test context with the request and response recorder
		router := gin.Default()
		router.GET("/employees", controller.ListEmployees)
		router.ServeHTTP(rr, req)

		// Assert response status code
		assert.Equal(t, http.StatusOK, rr.Code)

		// Assert response body
		var actualEmployees []models.Employee
		err := json.Unmarshal(rr.Body.Bytes(), &actualEmployees)
		assert.Nil(t, err)
		assert.NotNil(t, actualEmployees)
	})
}
