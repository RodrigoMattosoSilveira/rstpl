package utils

import (
	"html/template"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/gofiber/fiber/v2"
)

// Helper: renders a view with layout and partials
// templateCache avoids re-parsing templates repeatedly
var templateCache = struct {
	mu   sync.RWMutex
	data map[string]*template.Template
}{
	data: make(map[string]*template.Template),
}

type TmplPartial struct {
	Prefix   string
	Fn       string
	FullName string
	FileStr  string
	Name     string
}

func RenderPage(c *fiber.Ctx, partials []TmplPartial, data fiber.Map) error {

	projectRoot, err := FindProjectRoot()
	if err != nil {
		log.Printf("ERROR: Failed to find project root: %v", err)		
		// c.AbortWithStatus(500)
		return err
	}

	// TODO add caching mechanism here

	// 3. Parse the templates.
	var partialsStr []string
	partialsStrKey := ""
	for _, partial := range partials {
		partial.FullName = filepath.Join(projectRoot, "templates", partial.Fn)
		partialsStr = append(partialsStr, ReadTemplateFile(partial))
		partialsStrKey += partial.FullName
	}

	tmpl := template.New("layout")
	for _, part := range partialsStr {
		tmpl, err = tmpl.Parse(part)
		if err != nil {
			log.Fatal(err)
		}
	}

	// 4. Execute the template.
	// Send it to console to debug
	// err = tmpl.ExecuteTemplate(os.Stdout, "layout", data)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// Execute template into the response body
	// Use c.Response().BodyWriter() so template writes directly to Fiber response
	c.Type("html", "utf-8")
	return tmpl.Execute(c.Response().BodyWriter(), data)
}

func ReadTemplateFile(tmpl TmplPartial) string {
	// Read the file into a byte slice, then convert to string
	content, err := os.ReadFile(tmpl.FullName)
	if err != nil {
		log.Fatal(err)
	}
	// templateStr := tmpl.Prefix + string(content) + "\n" + "{{ end }}"
	// templateStr := "\n" + tmpl.Prefix + "\n" + string(content)+ "\n"  + "{{ end }}"
	templateStr := derivePrefix(tmpl.Name) + string(content) + "\n" + "{{ end }}"
	log.Println(templateStr)
	return templateStr
}

func derivePrefix(name string) string {
	// Prefix: `{{ define "bottom" }}`
	return "\n" + `{{ define "` + name + `" }}` + "\n"
}
