package app

import (
	"encoding/gob"
	"fmt"
	"net/http"

	"github.com/gin-contrib/logger"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// LoadRoutes .
func (app Application) LoadRoutes(r *gin.Engine) error {
	// logger
	r.Use(logger.SetLogger())

	// auth
	gob.Register(map[string]interface{}{})
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("auth-session", store))

	// static
	r.StaticFS("/public/static", http.Dir(fmt.Sprintf("%s/%s", app.path, "static")))

	// index
	r.GET("/", app.Index)

	api := r.Group("/api")
	{
		api.GET("/add", app.Add)
	}

	page := r.Group("/page", app.IsAuthenticated)
	{
		page.GET("/list", app.List)
		page.GET("/setup", app.Setup)
		page.GET("/dashboard", app.Dashboard)
		page.GET("/export", app.Export)
	}

	r.GET("/login", Login(app.authenticator))
	r.GET("/callback", Callback(app.authenticator))
	r.GET("/logout", app.Logout)

	return nil
}
