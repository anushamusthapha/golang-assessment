package controller

import (
	"golang-assessment/logger"
	"golang-assessment/models"
	"golang-assessment/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type EmployeeController struct {
	service *services.EmployeeService
}

func NewEmployeeController(service *services.EmployeeService) *EmployeeController {
	return &EmployeeController{service: service}
}

func (ctrl *EmployeeController) CreateEmployee(c *gin.Context) {
	var employee models.Employee
	if err := c.ShouldBindJSON(&employee); err != nil {
		logger.Log.Errorf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newEmployee := ctrl.service.CreateEmployee(employee.Name, employee.Position, employee.Salary)
	logger.Log.Infof("Created employee: %v", newEmployee)
	c.JSON(http.StatusCreated, newEmployee)
}

func (ctrl *EmployeeController) GetEmployeeByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Log.Errorf("Invalid ID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}
	employee, err := ctrl.service.GetEmployeeByID(id)
	if err != nil {
		logger.Log.Errorf("Error retrieving employee by ID %d: %v", id, err)
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	logger.Log.Infof("Retrieved employee: %v", employee)
	c.JSON(http.StatusOK, employee)
}

func (ctrl *EmployeeController) UpdateEmployee(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Log.Errorf("Invalid ID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}
	var employee models.Employee
	if err := c.ShouldBindJSON(&employee); err != nil {
		logger.Log.Errorf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedEmployee, err := ctrl.service.UpdateEmployee(id, employee.Name, employee.Position, employee.Salary)
	if err != nil {
		logger.Log.Errorf("Error updating employee by ID %d: %v", id, err)
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	logger.Log.Infof("Updated employee: %v", updatedEmployee)
	c.JSON(http.StatusOK, updatedEmployee)
}

func (ctrl *EmployeeController) DeleteEmployee(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Log.Errorf("Invalid ID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}
	if err := ctrl.service.DeleteEmployee(id); err != nil {
		logger.Log.Errorf("Error deleting employee by ID %d: %v", id, err)
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	logger.Log.Infof("Deleted employee with ID: %d", id)
	c.JSON(http.StatusOK, gin.H{"data": "Successfully deleted the employee"})
}

func (ctrl *EmployeeController) ListEmployees(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil || limit < 1 {
		limit = 10
	}
	employees, err := ctrl.service.ListEmployees(page, limit)
	if err != nil {
		logger.Log.Errorf("Error listing employees: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	logger.Log.Infof("Listed employees: %v", employees)
	c.JSON(http.StatusOK, employees)
}
