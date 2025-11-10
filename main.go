package main

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	"path/filepath"
)

func main() {
	r := gin.Default()

	// Serve static assets if you have them (optional)
	r.Static("/static", "./static")
	// Helper: renders a view with layout and partials
	render := func(c *gin.Context, tmpl string, data gin.H) {
		files := []string{
			filepath.Join("templates", "layout.html"),
			filepath.Join("templates", "header.html"),
			filepath.Join("templates", "sidebar.html"),
			filepath.Join("templates", "footer.html"),
			filepath.Join("templates", tmpl),
		}
		t := template.Must(template.ParseFiles(files...))
		c.Status(http.StatusOK)
		t.ExecuteTemplate(c.Writer, "layout", data)
	}
	render_ := func(c *gin.Context, tmpl string, data gin.H) {
		files := []string{
			filepath.Join("templates", "body.html"),
			filepath.Join("templates", tmpl),
		}
		t := template.Must(template.ParseFiles(files...))
		c.Status(http.StatusOK)
		t.ExecuteTemplate(c.Writer, "body", data)
	}
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

	r.GET("/login", func(c *gin.Context) {
		render(c, "login.html", gin.H{
			"Title":   "Login",
			"ShowNav": false,
		})
	})

	r.GET("/welcome", func(c *gin.Context) {
		render_(c, "welcome.html", gin.H{
			"Title":   "Welcome",
			"ShowNav": false,
		})
	})
	r.GET("/bemvindo", func(c *gin.Context) {
		render_(c, "bemvindo.html", gin.H{
			"Title":   "Bem Vindo",
			"ShowNav": false,
		})
	})

	r.Run(":8080")
}

// Helper: Parse all templates with layout support
func loadTemplates(templatesDir string) *template.Template {
	tmpl := template.Must(template.ParseGlob(templatesDir + "/*.html"))
	return tmpl
}
