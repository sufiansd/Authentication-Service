package handlers

import (
	"Authentication-Service/gen/restapi/operations"
	"Authentication-Service/service"
	"github.com/go-openapi/runtime/middleware"
)

type ShowProfile struct {}

func ShowProfileHandler() operations.ShowProfileHandler{
	return &ShowProfile{}
}

func (s ShowProfile) Handle(params operations.ShowProfileParams, i interface{}) middleware.Responder {


	employee, errorString := service.ShowProfile(params.Email)
	if errorString == ""{
		return operations.NewShowProfileOK().WithPayload(DomainModelToGenModel(employee))
	}else{
		return operations.NewShowProfileInternalServerError().WithPayload(errorString)
	}
}