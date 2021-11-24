package handlers

import (
	"Authentication-Service/gen/restapi/operations"
	"Authentication-Service/service"
	"fmt"
	"github.com/go-openapi/runtime/middleware"
)

type ValidateToken struct{}

func ValidateTokenHandler()operations.ValidateTokenHandler{
	return &ValidateToken{}
}

func (v ValidateToken) Handle(params operations.ValidateTokenParams) middleware.Responder {
	fmt.Println("Welcome to handler")
	isValid, err := service.ValidateHeader(params.HTTPRequest.Header.Get("Authorization"))
	if isValid{
		return operations.NewValidateTokenOK().WithPayload("200")
	}else{
		if err == "Unauthorized"{
			return operations.NewValidateTokenUnauthorized().WithPayload("401")
		}else if err == "Token is expired"{
			return operations.NewValidateTokenUnauthorized().WithPayload("401")
		}else{
			return operations.NewValidateTokenInternalServerError().WithPayload("500")
		}
	}
}
