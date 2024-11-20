package main

import (
	"task-tracker-server/config"
	"task-tracker-server/internal/app"
)

//	@title			Task tracker server
//	@version		1.0
//	@description	Task tracker server app.

//	@contact.name	neketli
//	@contact.email	neketli.dev@gmail.com

//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization

// @BasePath	/api/v1
func main() {
	app.Run(config.New())
}
