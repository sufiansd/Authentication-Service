package service

import (
	"Authentication-Service/constants"
	"Authentication-Service/db/MySqlClient"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"strings"
)

func ValidateHeader(tokenString string) (bool, string){
	claims := jwt.MapClaims{}
	tokenString = strings.Split(tokenString, " ")[1]
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return constants.MySignKey, nil
	})
	if err != nil{
		return false, fmt.Sprintf("%s", err)
	}
	authorized := claims["authorized"]
	userID :=  claims["userID"]
	convertedUserID  := fmt.Sprintf("%f", userID)
	expiration := claims["exp"]

	if authorized == false{
		return false, "Unauthorized"
	}
	if expiration == 0{
		return false, "Token Expired"
	}
	exists, responseString := MySqlClient.Exists(convertedUserID)
	if exists{
		return true, "User Valid"
	}else{
		//fmt.Println(responseString)
		return false, responseString
	}
}
