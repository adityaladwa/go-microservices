package main

import (
	"fmt"
	"github.com/adityaladwa/go-microservices/accountservice/services"
)

var appName = "accountservice"

func main(){
	fmt.Printf("Starting %v\n",appName)
	service.StartWebServer("6767")
}
