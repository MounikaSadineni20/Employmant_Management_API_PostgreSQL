package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	database "vscode/employeeapidatabase/database"
	entity "vscode/employeeapidatabase/entities"

	"github.com/gin-gonic/gin"
)

func GetAllEmployeeRecords(context *gin.Context) {
	var employees []entity.Employee
	db := database.ConfigureDB()
	rows, err := db.Query("select * from employees")
	if err != nil {
		return
	}
	for rows.Next() {
		var id int
		var name string
		var email string
		var experience int
		var gender string
		var prevEmployer string
		_ = rows.Scan(&id, &name, &email, &experience, &gender, &prevEmployer)
		employees = append(employees, entity.Employee{Id: id, Name: name, Email: email, Experience: experience, Gender: gender, PrevEmployer: prevEmployer})
	}
	if len(employees) == 0 {
		context.IndentedJSON(http.StatusOK, "No Employee data available")
		return
	}
	context.IndentedJSON(http.StatusOK, employees)
}

func GetByEmployeeName(name string) ([]entity.Employee, error) {
	var employees []entity.Employee
	db := database.ConfigureDB()
	rows, err := db.Query("SELECT * from Employees WHERE name = $1", name)
	if err != nil {
		return nil, errors.New("After Error :No Employee data found with the provided Employee Name")
	}
	count := 0
	for rows.Next() {
		var id int
		var name string
		var email string
		var experience int
		var gender string
		var prevEmployer string
		count++
		_ = rows.Scan(&id, &name, &email, &experience, &gender, &prevEmployer)
		employees = append(employees, entity.Employee{Id: id, Name: name, Email: email, Experience: experience, Gender: gender, PrevEmployer: prevEmployer})
	}
	if count == 0 {
		return nil, errors.New("After Scanning :No Employee data found with the provided Empoyee Name")
	}
	return employees, nil
}

func GetEmployeeByName(context *gin.Context) {
	employeeName := context.Param("name")
	filteredEmployees, err := GetByEmployeeName(employeeName)
	fmt.Println(filteredEmployees)
	fmt.Println(employeeName)
	fmt.Println(err)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "No Employee Found with the provided Empoyee Name"})
		return
	}
	context.IndentedJSON(http.StatusOK, filteredEmployees)
}

func GetByEmployeeId(id string) ([]entity.Employee, error) {
	var employees []entity.Employee
	db := database.ConfigureDB()
	rows, err := db.Query("SELECT * from Employees WHERE id = $1", id)
	if err != nil {
		return nil, errors.New("No Employee data found with the provided Employee Id")
	}
	count := 0
	for rows.Next() {
		var id int
		var name string
		var email string
		var experience int
		var gender string
		var prevEmployer string
		_ = rows.Scan(&id, &name, &email, &experience, &gender, &prevEmployer)
		count++
		employees = append(employees, entity.Employee{Id: id, Name: name, Email: email, Experience: experience, Gender: gender, PrevEmployer: prevEmployer})
	}
	if count == 0 {
		return nil, errors.New("No Employee data found with the provided Empoyee Name")
	}
	return employees, nil
}

func GetEmoployeeById(context *gin.Context) {
	employeeId := context.Param("id")
	filteredEmployees, err := GetByEmployeeId(employeeId)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	context.IndentedJSON(http.StatusOK, filteredEmployees)
}

func AddEmployeeRecord(context *gin.Context) {
	var newEmployee entity.Employee
	db := database.ConfigureDB()
	if err := context.BindJSON(&newEmployee); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Some thing went wrong while Adding a Employee. Please Make sure all fields are present"})
		return
	}
	checking := IsValid(newEmployee)
	if len(checking) != 0 {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": checking})
		return
	}
	_, err := db.Exec("Insert into Employees(id, name, email, experience, gender, prevemployer) values($1, $2, $3, $4, $5, $6)",
		newEmployee.Id, newEmployee.Name, newEmployee.Email, newEmployee.Experience, newEmployee.Gender, newEmployee.PrevEmployer)
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Some thing went wrong while Adding a Employee"})
	}
	context.IndentedJSON(http.StatusCreated, newEmployee)
}

func DeleteById(id int) (string, error) {
	db := database.ConfigureDB()
	rows, err := db.Query("SELECT * from Employees WHERE id = $1", id)
	if err != nil || !rows.Next() {
		return "", errors.New("No Employee data found by Empoyee Id")
	}
	_, err1 := db.Exec("delete from employees where id = $1", id)
	if err1 != nil {
		return "", err1
	}
	return "Deleted Successfully", nil
}

func DeleteEmployeeById(context *gin.Context) {
	empId, _ := strconv.Atoi(context.Param("id"))
	msg, err := DeleteById(int(empId))
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, err.Error())
		return
	}
	context.IndentedJSON(http.StatusOK, msg)
}
