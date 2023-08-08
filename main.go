package main

import (
	controller "vscode/employeeapidatabase/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/getAllEmployees", controller.GetAllEmployeeRecords)
	router.GET("/getEmployeeByEmployeeName/:name", controller.GetEmployeeByName)
	router.GET("/getEmployeebyEmloeeId/:id", controller.GetEmoployeeById)
	router.POST("/createEmployee", controller.AddEmployeeRecord)
	router.DELETE("/deleteEmployeeById/:id", controller.DeleteEmployeeById)
	router.Run("localhost:9191")
}
