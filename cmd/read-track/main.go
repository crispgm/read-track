// package main serves read track API
package main

import (
	"github.com/gin-gonic/gin"

	"github.com/crispgm/read-track/internal/app"
	"github.com/crispgm/read-track/internal/infra"
)

func main() {
	// load conf
	conf, err := infra.LoadConf("./")
	if err != nil {
		panic(err)
	}

	// register routers
	r := gin.Default()
	app.LoadRoutes(r, conf)

	// run
	r.Run(conf.HTTP.Port)
}
