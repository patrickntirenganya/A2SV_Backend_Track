package main

import (
	"fmt"

	"library_management/controllers"
	"library_management/services"
)

func main() {
	fmt.Println("Welcome to the Library Management System!")

	library := services.NewLibrary()

	controller := controllers.NewLibraryController(library)

	controller.Start()
}
