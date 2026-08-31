package main

import (
	"log"

	"Task_manager/data"
	"Task_manager/router"
)

func main() {
	taskService := data.NewTaskService()

	r := router.SetupRouter(taskService)

	log.Println("Task Management API running on http://localhost:8080")

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
