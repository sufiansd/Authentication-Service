package MySqlClient

import (
	"Authentication-Service/db"
)

func Exists(userID string) (bool, string){
	db := 	client.SqlClient()
	checkQuery := "SELECT * FROM employeeSignup WHERE ID = '"+userID+"'"
	result, err := db.Query(checkQuery)
	if err!= nil{
		return false, "Error with database"
	}
	if result.Next(){
		return true, "User exists"
	}

	return false, "User does not exist"
}
