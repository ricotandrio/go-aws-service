package main

import (
	"fmt"
	"go-service-aws/configs"
	"go-service-aws/routers"
)

func main() {
	configs.InitEnvironment()
	configs.ConnectDB()

	router := routers.SetupRouter()

	err := router.Run(":8080")

	if err != nil {
		panic(err)
	}
	
	fmt.Println("Server is running on port 8080")
}
