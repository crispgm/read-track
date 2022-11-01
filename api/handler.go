// Package api serves in Vercel
package api

import (
	"net/http"

	"github.com/crispgm/read-track/internal/app"
	"github.com/crispgm/read-track/internal/infra"
	"github.com/gin-gonic/gin"
)

// Handler Vercel faas function
func Handler(w http.ResponseWriter, r *http.Request) {
	conf, err := infra.LoadConf("./")
	if err != nil {
		panic(err)
	}
	e := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	app.LoadRoutes(e, conf)
	e.ServeHTTP(w, r)
}
