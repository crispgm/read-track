// package main serves read track API
package main

import (
	"flag"

	"github.com/crispgm/read-track/internal/app"
	"github.com/crispgm/read-track/internal/infra"
	"github.com/gin-gonic/gin"
)

var (
	path string

	// default port
	port        = ":80"
	defaultPath = "./"
)

func main() {
	flag.StringVar(&path, "conf", defaultPath, "Conf Path")
	flag.Parse()

	var appl app.Application
	var err error
	err = appl.Init(path)
	if err != nil {
		panic(err)
	}
	// load conf
	conf, err := infra.LoadConf("./")
	if err != nil {
		panic(err)
	}
	// register routers
	r := gin.Default()
	app.LoadRoutes(r)

	// run
	r.Run(conf.HTTP.Port)
}
