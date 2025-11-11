package main

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	"path/filepath"
)
// Helper: renders a view with layout and partials
func render (c *gin.Context, tmpl string, data gin.H) {
	var layout string
	var layoutStr string
	route := c.FullPath()

	switch route {
	case "/":
		layout = "layout.html"
		layoutStr = "layout"
	case "/about":
		layout = "layout.html"
		layoutStr = "layout"
	case "/welcome":
		layout = "body.html"
		layoutStr = "body"
	case "/bemvindo":
		layout = "body.html"
		layoutStr = "body"
	case "/login":
		layout = "body.html"
		layoutStr = "body"
	case "/logon":
		layout = "body.html"
		layoutStr = "body"
	default:
		layout = "layout.html"
		layoutStr = "layout"
	}
	files := []string{
		filepath.Join("templates", layout),
		filepath.Join("templates", tmpl),
	}
	t := template.Must(template.ParseFiles(files...))
	c.Status(http.StatusOK)
	t.ExecuteTemplate(c.Writer, layoutStr, data)
}
func main() {
	r := gin.Default()

	// Serve static assets if you have them (optional) 
	r.Static("/static", "./static")

	// Routes
	r.GET("/", func(c *gin.Context) {
		render(c, "home.html", gin.H{
			"Title":   "Home",
			"ShowNav": true,
		})
	})

	r.GET("/about", func(c *gin.Context) {
		render (c, "about.html", gin.H{
			"Title":   "Home",
			"ShowNav": true,
		})
	})

	r.GET("/welcome", func(c *gin.Context) {
		render(c, "welcome.html", buildPipeline())
	})

	r.GET("/bemvindo", func(c *gin.Context) {
		render(c, "bemvindo.html", buildPipeline())
	})

	r.GET("/login", func(c *gin.Context) {
		render(c, "login.html", buildPipeline())
	})

	r.GET("/logon", func(c *gin.Context) {
		render(c, "logon.html", buildPipeline())
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