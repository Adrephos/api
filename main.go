package main

import (
	"fmt"
	"os"
	"github.com/Adrephos/api/routes"
)

func main() {
	router := routes.SetupRouter()

	fmt.Println("Server running on port " + os.Getenv("PORT"))
	router.Run()
}
