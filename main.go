// package main serves read track API
package main

import (
	"flag"

	"github.com/crispgm/read-track/internal/app"
	"github.com/gin-gonic/gin"
)

var (
	path string

	// default port
	port        = ":80"
	defaultPath = "./"
)

func main() {
	flag.StringVar(&path, "working-path", defaultPath, "Working Path")
	flag.Parse()

	var appl app.Application
	var err error
	err = appl.Init(path)
	if err != nil {
		panic(err)
	}
	// auto migrate
	err = appl.AutoMigrate()
	if err != nil {
		panic(err)
	}
	// register routers
	r := gin.Default()
	err = appl.LoadRoutes(r)
	if err != nil {
		panic(err)
	}
	// run
	r.Run(appl.Conf().HTTP.Port)
}
