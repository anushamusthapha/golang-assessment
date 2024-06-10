package routers

import (
	"golang-assessment/controller"
	repository "golang-assessment/respository"
	"golang-assessment/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	employeeRepo := repository.NewEmployeeRepository(db)
	employeeService := services.NewEmployeeService(employeeRepo)
	employeeController := controller.NewEmployeeController(employeeService)

	router := gin.Default()

	router.POST("/employees", employeeController.CreateEmployee)
	router.GET("/employees/:id", employeeController.GetEmployeeByID)
	router.PUT("/employees/:id", employeeController.UpdateEmployee)
	router.DELETE("/employees/:id", employeeController.DeleteEmployee)
	router.GET("/employees", employeeController.ListEmployees)

	return router
}
