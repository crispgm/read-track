package app

import (
	"encoding/gob"
	"fmt"
	"net/http"

	"github.com/crispgm/read-track/internal/infra"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// LoadRoutes .
func (app Application) LoadRoutes(r *gin.Engine, conf *infra.Conf) error {
	gob.Register(map[string]interface{}{})
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("auth-session", store))

	auth, err := infra.InitAuthenticator(app.Conf().Auth0)
	if err != nil {
		return err
	}

	r.StaticFS("/public/static", http.Dir(fmt.Sprintf("%s/%s", app.path, "static")))
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

	r.GET("/login", Login(auth))
	r.GET("/callback", Callback(auth))
	r.GET("/logout", app.Logout)

	return nil
}

func (app Application) validTokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
