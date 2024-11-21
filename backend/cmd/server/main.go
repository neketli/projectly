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

//	@securityDefinitions.apikey	Bearer
//	@in							header
//	@name						Authorization

// @BasePath	/api/v1
// @Host		localhost:8083
func main() {
	app.Run(config.New())
}
