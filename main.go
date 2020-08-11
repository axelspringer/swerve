package main

import (
	"log"

	"github.com/axelspringer/swerve/app"
)

func main() {
	application := app.NewApplication()
	if err := application.Setup(); err != nil {
		log.Fatal(err)
		return
	}
	application.Run()
}
