package main

import (
	"github.com/gin-gonic/gin"
	"github.com/RodrigoMattosoSilveira/rstpl/internal/utils"

)

func main() {
	r := gin.Default()

	// Serve static assets if you have them (optional) 
	r.Static("/static", "./static")

	// Routes
	r.GET("/", func(c *gin.Context) {
		utils.Render(c, "home.html", gin.H{
			"Title":   "Home",
			"ShowNav": true,
		})
	})

	r.GET("/about", func(c *gin.Context) {
		utils.Render (c, "about.html", gin.H{
			"Title":   "Home",
			"ShowNav": true,
		})
	})

	r.GET("/welcome", func(c *gin.Context) {
		utils.Render(c, "welcome.html", buildPipeline())
	})

	r.GET("/bemvindo", func(c *gin.Context) {
		utils.Render(c, "bemvindo.html", buildPipeline())
	})

	r.GET("/login", func(c *gin.Context) {
		utils.Render(c, "login.html", buildPipeline())
	})

	r.GET("/logon", func(c *gin.Context) {
		utils.Render(c, "logon.html", buildPipeline())
	})

	r.GET("/welcome_login", func(c *gin.Context) {
		var partials = []utils.TmplPartial{
			{Name: "layout", Fn: "layout.html",  Prefix: `{{ define "layout" }}`, FullName: "", FileStr: ""},
			{Name: "bottom",    Fn: "welcome.html", Prefix: `{{ define "bottom" }}`,    FullName: "", FileStr: ""},
			{Name: "top", Fn: "cc.tmpl",      Prefix: `{{ define "top" }}`, FullName: "", FileStr: ""},		
		}

		// Call our custom renderer.
		// The name "layout.tmpl" tells the template engine which template definition to execute first.
		data := gin.H{
			"Tenant": "MC",
			"Host":   "Madone Logistics",
		}
		utils.RenderPage(c, partials, data)
	})
	
	r.Run(":8080")
}
/*
 * An attempt to consilidate data for template rendering
 */
func buildPipeline() gin.H {
	return gin.H{
		"Tenant": "MC",
		"Host": "Madrone Logistics",
	}
}