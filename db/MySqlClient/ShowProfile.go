package MySqlClient

import (
	client "Authentication-Service/db"
	"Authentication-Service/models"
	"fmt"
)


func ShowProfile(email string) (*models.DomainEmployee, string) {
	db := client.SqlClient()
	query := "SELECT FullName, FatherName, Age, Email, Password, Gender, NIC FROM employeeSignup WHERE Email='" + email + "'"
	var employee models.DomainEmployee
	result, _ := db.Query(query)

	defer result.Close()
	for result.Next() {
		err := result.Scan(&employee.FullName, &employee.FatherName, &employee.Age, &employee.Email, &employee.Password, &employee.Gender, &employee.NIC)
		if err != nil {
			fmt.Println("Error: ", err)
			return nil, fmt.Sprintf("%s", err)
		}
	}
	return &employee, ""
}