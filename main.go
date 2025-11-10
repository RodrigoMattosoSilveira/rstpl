package main

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Serve static assets if you have them (optional)
	r.Static("/static", "./static")

	// Load all templates
	r.SetHTMLTemplate(loadTemplates("./templates"))

	// Routes
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.html", gin.H{"Title": "Home"})
	})

	r.GET("/about", func(c *gin.Context) {
		c.HTML(http.StatusOK, "about.html", gin.H{"Title": "About"})
	})

	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{"Title": "Login"})
	})

	r.GET("/body", func(c *gin.Context) {
		c.HTML(http.StatusOK, "body.html", gin.H{"Title": "Login"})
	})

	r.GET("/welcome", func(c *gin.Context) {
		c.HTML(http.StatusOK, "welcome.html", gin.H{"Title": "Login"})
	})


	r.GET("/bemvindo", func(c *gin.Context) {
		c.HTML(http.StatusOK, "bemvindo.html", gin.H{"Title": "Login"})
	})

	r.Run(":8080")
}

// Helper: Parse all templates with layout support
func loadTemplates(templatesDir string) *template.Template {
	tmpl := template.Must(template.ParseGlob(templatesDir + "/*.html"))
	return tmpl
}
