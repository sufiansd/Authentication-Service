package service

import (
	"Authentication-Service/constants"
	"Authentication-Service/db/MySqlClient"
	"Authentication-Service/models"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"strings"
)

func ShowProfile(email string) (*models.DomainEmployee, string) {
	employee, errorString := MySqlClient.ShowProfile(email)
	employee.Password = "Hidden"
	return employee, errorString
}

func ValidateToken(tokenString string) (bool, string){
	fmt.Println("Welcome to service layer",tokenString)
	claims := jwt.MapClaims{}
	tokenString = strings.Split(tokenString, " ")[1]
	fmt.Println("TOKEN IS: "+tokenString)
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return constants.MySignKey, nil
	})
	if err != nil{
		return false, fmt.Sprintf("%s", err)
	}
	authorized := claims["authorized"]
	//userID :=  claims["userID"]
	//convertedUserID  := fmt.Sprintf("%f", userID)
	expiration := claims["exp"]

	if authorized == true && expiration != 0{
		return true, "Authorized"
	}else{
		return false, "Unauthorized"
	}
}


