package controllers

import (
	"net/url"
	entity "vscode/employeeapidatabase/entities"
)

func IsValid(employee entity.Employee) url.Values {
	errs := url.Values{}
	if employee.Name == "" {
		errs.Add("Name", "The name is required!")
	}
	if !(employee.Id >= 1) {
		errs.Add("Id", "The Id should not be in Negative")
	}
	if employee.Email == "" {
		errs.Add("Email", "The Email is required!")
	}
	if length := len(employee.Email); !(length > 5 && length < 30) {
		errs.Add("Email", "The Email should be in the range of 5-30!")
	}
	if employee.Gender == "" {
		errs.Add("Gender", "The Gender is required!")
	}
	if employee.Experience == 0 {
		errs.Add("Experience", "The Experience is required!")
	}
	if employee.PrevEmployer == "" {
		errs.Add("PrevEmployer", "The PrevEmployer is required!")
	}

	return errs
}
