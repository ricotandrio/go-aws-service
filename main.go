package main

import (
	"go-service-aws/configs"
	"go-service-aws/routers"
	"fmt"
)

func main() {
	configs.InitEnvironment()
	configs.ConnectDB()

	router := routers.SetupRouter()

	err := router.Run(":8080")
	
	if err != nil {
		panic(err)
	}
	
	fmt.Println("Server is running on http://localhost:8080")
}
