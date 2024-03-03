package main

import "segments-api/internal/app"

const configsDir = "configs"

func main() {
	app.Run(configsDir)
}
