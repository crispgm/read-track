// package main serves read track API
package main

import (
	"embed"
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

//go:embed static/* templates/*
var resources embed.FS

func main() {
	flag.StringVar(&path, "conf", defaultPath, "Conf Path")
	flag.Parse()

	var appl app.Application
	var err error
	err = appl.Init(path)
	if err != nil {
		panic(err)
	}
	conf := appl.Conf()
	// auto migrate
	appl.AutoMigrate()
	// register routers
	r := gin.Default()
	appl.LoadRoutes(r, conf, &resources)
	// run
	r.Run(conf.HTTP.Port)
}
