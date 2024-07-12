package main

import (
	"log"
	app "pr1/internal/rest/app"
	"pr1/internal/rest/config"
)

func main() {

	if err := app.Run(config.NewConfig()); err != nil {
		log.Println("error app.Run(): ", err)
	}
}
