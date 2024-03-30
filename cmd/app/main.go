package main

import (
	_ "segments-api/docs"
	"segments-api/internal/app"
)

const configsDir = "configs"

// @title Segments API
// @version 1.0
// @description Segments API server.

// @contact.name Evgenii Suvorov
// @contact.email eo.suvorov@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
func main() {
	app.Run(configsDir)
}
