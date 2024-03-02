package main

//	@title			Database System Design Project
//	@version		1.0.0

//	@contact.name	Maintainer Polaris
//	@contact.email	polaris@fduhole.com

//	@license.name	Apache 2.0
//	@license.url	https://www.apache.org/licenses/LICENSE-2.0.html

//	@host
//	@BasePath	/api

import (
	"log"
	"src/bootstrap"
	_ "src/docs"
)

func main() {
	app, err := bootstrap.Init()
	if err != nil {
		log.Fatal(err)
		return
	}
	err = app.Listen(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
