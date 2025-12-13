package main

import (
	"log"

	"github.com/RodrigoMattosoSilveira/rstpl/internal/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	// Serve static assets if you have them (optional)
	app.Static("/static", "./static")

	// Routes

	app.Get("/welcome/:id", func(c *fiber.Ctx) error {
		id, err :=c.ParamsInt("id")
		if err != nil {
			log.Printf("Invalid id: %d", id)
			c.Redirect("/login")
		}
		var partials = []utils.TmplPartial {
			{Name: "layout", Fn: "layout.html", },  // Must be the first one
			{Name: "top", Fn: "cc.tmpl", },
			{Name: "bottom",    Fn: "welcome.html", },
		}

		// Call our custom renderer.
		// The name "layout.tmpl" tells the template engine which template definition to execute first.
		return utils.RenderPage(c, partials, buildPipeline())
	}).Name("welcome")

	app.Get("/login", func(c *fiber.Ctx) error {
		var partials = []utils.TmplPartial {
			{Name: "layout", Fn: "layout.html", }, // Must be the first one
			{Name: "top", Fn: "login.html", },
		}

		// Call our custom renderer.
		// The name "layout.tmpl" tells the template engine which template definition to execute first.
		return utils.RenderPage(c, partials, buildPipeline())
	})

	app.Post("/login", func(c *fiber.Ctx) error {
		log.Println("Login attempt:", c.FormValue("username"))
		// In a real app, you'd validate credentials here.
		return c.Redirect("/welcome/1")
	})
	log.Fatal(app.Listen(":3000"))
}

/*
 * An attempt to consilidate data for template rendering
 */
func buildPipeline() fiber.Map {
	return fiber.Map{
		"Tenant": "MC",
		"Host":   "Madrone Logistics",
	}
}
