package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type router struct {
	Gin *gin.Engine
}

func (r *router) setupRoutes() error {
	r.loadTemplates("./templates/*")
	r.Gin.Static("/assets", "./assets")

	r.Gin.GET("/", func(c *gin.Context) {
		r.renderTemplate(
			c,
			http.StatusOK,
			"home.html",
			gin.H{},
		)
	})
	r.Gin.NoRoute(func(c *gin.Context) {
		r.renderTemplate(
			c,
			http.StatusNotFound,
			"404.html",
			gin.H{},
		)
	})

	return nil
}

func (r *router) loadTemplates(pattern string) {
	r.Gin.LoadHTMLGlob(pattern)
}

func (r *router) renderTemplate(c *gin.Context, code int, name string, obj interface{}) {
	c.HTML(code, name, obj)
}
