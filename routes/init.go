package routes

import (
	"core/app/entities"
	"core/config"
	"core/middleware"
	"github.com/gin-gonic/contrib/gzip"
	"github.com/gin-gonic/gin"
	"master/resource"
	"master/resource/helper"
)

var (
	authEntity entities.AuthInterface
)

type Router struct {
}

func GetEngine(r *resource.Resource) *gin.Engine {
	// Setup resource
	h := &helper.Helper{
		Resource: r,
	}
	authEntity = entities.NewAuthEntity(h)

	// Set up gin
	gin.SetMode(config.AppMode)
	app := gin.Default()
	app.Use(gzip.Gzip(gzip.DefaultCompression))
	app.Use(middleware.CORS())

	// Setup router
	router := &Router{}
	ApplyAuthRoutes(app, router)
	app.Use(middleware.AuthInit())

	return app
}

func CurrentUserID(c *gin.Context) int {
	userIDSRaw, exist := c.Get("userID")
	if !exist {
		return -1
	}
	userID := userIDSRaw.(int)
	return userID
}
