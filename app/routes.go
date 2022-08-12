package app

import (
	"net/http"

	"bugvalidate/app/actions/home"
	"bugvalidate/app/middleware"
	"bugvalidate/public"

	"github.com/gobuffalo/buffalo"
)

// SetRoutes for the application
func setRoutes(root *buffalo.App) {
	root.Use(middleware.RequestID)
	root.Use(middleware.Database)
	root.Use(middleware.ParameterLogger)
	root.Use(middleware.CSRF)

	root.GET("/", home.Index).Name("homeIndexPath")
	root.POST("/create-1", home.Create1).Name("homeCreate1Path")
	root.POST("/create-2", home.Create2).Name("homeCreate2Path")
	root.ServeFiles("/", http.FS(public.FS()))
}
