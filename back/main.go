package main

import (
	app "leisure_time/cmd/app"
	config "leisure_time/cmd/config"
)

func main() {
	config := config.DefaultAppConfig()

	app := app.NewApp(config)

	app.Setup()
	app.Run()
}
