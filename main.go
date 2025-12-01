package main

import (
	"log"

	"github.com/RodrigoMattosoSilveira/rstpl/internal/utils"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Serve static assets if you have them (optional)
	app.Static("/static", "./static")

	// Routes
	app.Get("/", func(c *fiber.Ctx) error {
		utils.Render(c, "home.html", fiber.Map{
			"Title":   "Home",
			"ShowNav": true,
		})
		return nil
	})

	app.Get("/about", func(c *fiber.Ctx) error {
		return utils.Render(c, "about.html", fiber.Map{
			"Title":   "Home",
			"ShowNav": true,
		})
	})

	app.Get("/welcome", func(c *fiber.Ctx) error {
		return utils.Render(c, "welcome.html", buildPipeline())
	})

	app.Get("/bemvindo", func(c *fiber.Ctx) error {
		return utils.Render(c, "bemvindo.html", buildPipeline())
	})

	app.Get("/login", func(c *fiber.Ctx) error {
		return utils.Render(c, "lofiber.Maptml", buildPipeline())
	})

	app.Get("/logon", func(c *fiber.Ctx) error {
		return utils.Render(c, "logon.html", buildPipeline())
	})

	app.Get("/welcome_login", func(c *fiber.Ctx) error {
		var partials = []utils.TmplPartial {
			{Name: "layout", Fn: "layout.html", Prefix: `{{ define "layout" }}`, FullName: "", FileStr: ""},
			{Name: "bottom", Fn: "welcome.html", Prefix: `{{ define "bottom" }}`, FullName: "", FileStr: ""},
			{Name: "top", Fn: "cc.tmpl", Prefix: `{{ define "top" }}`, FullName: "", FileStr: ""},
		}

		// Call our custom renderer.
		// The name "layout.tmpl" tells the template engine which template definition to execute first.
		data := fiber.Map{
			"Tenant": "MC",
			"Host":   "Madone Logistics",
		}
		return utils.RenderPage(c, partials, data)
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
