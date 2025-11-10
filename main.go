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

	// Registration page
	r.GET("/register", func(c *gin.Context) {
		render(c, "register.html", gin.H{
			"Title":   "Register",
			"ShowNav": false,
		})
	})
	r.GET("/logon", func(c *gin.Context) {
		render(c, "logon.html", gin.H{
			"Title":   "Register",
			"ShowNav": false,
		})
	})

	r.GET("/welcome", func(c *gin.Context) {
		render_(c, "welcome.html", buildPipeline())
	})

	r.GET("/bemvindo", func(c *gin.Context) {
		render_(c, "bemvindo.html", buildPipeline())
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