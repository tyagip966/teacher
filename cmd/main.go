package main

import (
	"log"
	"teacher/utils"
)

func main() {
	container := utils.Container{Injected: false}
	err := container.TriggerDI()
	if err != nil {
		log.Fatal("Error starting Server")
	}
}
