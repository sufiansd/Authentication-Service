package main

import (
	"Authentication-Service/gen/restapi"
	"Authentication-Service/gen/restapi/operations"
	"Authentication-Service/handlers"
	"Authentication-Service/service"
	"flag"
	"fmt"
	"github.com/go-openapi/loads"
	"log"
)

var portFlag = flag.Int("Port", 3002, "Port to run this server on.")

func main(){
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil{
		log.Fatal(err)
	}

	api := operations.NewAuthenticationServiceAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	flag.Parsed()
	server.Port = *portFlag

	api.EmployeeSignupHandler = handlers.SignupEmployeeHandler()
	api.LoginHandler = handlers.LoginHandler()
	api.ValidateTokenHandler = handlers.ValidateTokenHandler()
	api.ShowProfileHandler = handlers.ShowProfileHandler()
	api.BearerAuth = func(token string) (interface{}, error) {
		fmt.Println(token)
		authorized, _ := service.ValidateToken(token)
		if authorized {
			fmt.Println(authorized)
			fmt.Println("Authorized!")
			return api, nil
		} else {
			fmt.Println("Unauthorized")
			return nil, nil
		}
	}


	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
