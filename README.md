# Abstract
Why can't Go, html/template, and Fiber collaborate to compose templates?

I want to render a `template` using a `base template` and one of two `partial templates`, based on the context. Something like:

```text
  Patial A                         Base                          Template
|────────────|                 |───────────────|               |───────────────|
|            |                 |               |               |               |
|  AAAAAAAA  |  componse with  | {{ .Target }} |   to yield    |   AAAAAAAA    | 
|            |                 |               |               |               |
|────────────|                 |───────────────|               |───────────────|


  Patial A                         Base                          Template
|────────────|                 |───────────────|               |───────────────|
|            |                 |               |               |               |
|  BBBBBBBB  |  componse with  | {{ .Target }} |   to yield    |   BBBBBBBB    | 
|            |                 |               |               |               |
|────────────|                 |───────────────|               |───────────────|

```

# Implementation
I constructed an experiement using the following folder structure, templates, and logic
## Folder Structure
templates/
│
├── main.go
└── view/
    ├── base.html
    ├── partailA.html
    ├── partailB.html

## Templates
I used the html/template `define` and `block` and `layout` actions:

```html
{{ define "base" .}}
  <!-- This is the base template -->
<div>
	<!-- insert particl template here -->
	{{ block "partial" . }}{{ end }}
</div>
```
{{ end }}

These are the partial templates:
```html
**{{ define "partial" }}**
  <!-- This is a partical template -->
<div>
	<p>Partial Template A</p>
</div>
{{ end }}
```

```html
{{ define "partial"}}
  <!-- This is another template -->
<div>
	<p>Partial Template B</p>
</div>
{{ end }}
```

## Logic

```go
package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

const (
	ADMIN         = "admin"
)

var LayoutDir string = "views/layouts"
var bootstrap *template.Template

func main() {
	var err error
	
	// Create a new engine
	engine := html.New("./views", ".html")

	// Pass the engine to the Views
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Static("/", "./public")

	app.Get("/patialA", func(c *fiber.Ctx) error {
		// Render index
		return c.Render("partialA", fiber.Map{}, "base")
	})

	app.Get("/patialB", func(c *fiber.Ctx) error {
		// Render index
		return c.Render("partialB", fiber.Map{}, "base")
	})
	log.Fatal(app.Listen(":3000"))
}
```

After trying all kinds of suggestions, including from ChatGPT, I gave up on the Go, html/template, and Fiber stack, switched to Go, html/template, and Gin, and made it work. 

Regardless, I'm still curious about whether there is a way to make it work with Fiber.