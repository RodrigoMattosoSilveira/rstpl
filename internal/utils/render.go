package utils

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"

	"github.com/gin-gonic/gin"
)

// Helper: renders a view with layout and partials
// templateCache avoids re-parsing templates repeatedly
var templateCache = struct {
	mu   sync.RWMutex
	data map[string]*template.Template
}{
	data: make(map[string]*template.Template),
}

func Render(c *gin.Context, partial string, data gin.H) {
	var layout string
	var layoutName string
	route := c.FullPath()

	switch route {
	case "/":
		layout = "layout.html"
		layoutName = "layout"
	case "/about":
		layout = "layout.html"
		layoutName = "layout"
	case "/welcome":
		layout = "body.html"
		layoutName = "body"
	case "/bemvindo":
		layout = "body.html"
		layoutName = "body"
	case "/login":
		layout = "body.html"
		layoutName = "body"
	case "/logon":
		layout = "body.html"
		layoutName = "body"
	default:
		layout = "layout.html"
		layoutName = "layout"
	}
	// Key for cache
	key := layout + "|" + partial

	// Try cached template
	templateCache.mu.RLock()
	t, ok := templateCache.data[key]
	templateCache.mu.RUnlock()

	if !ok {
		files := []string{
			filepath.Join("templates", layout),
			filepath.Join("templates", partial),
		}
		t = template.Must(template.ParseFiles(files...))
		templateCache.mu.Lock()
		templateCache.data[key] = t
		templateCache.mu.Unlock()
	}

	// Execute template using its defined name (not filename)
	c.Status(http.StatusOK)
	if err := t.ExecuteTemplate(c.Writer, layoutName, data); err != nil {
		c.String(http.StatusInternalServerError, "template error: %v", err)
	}
}

type TmplPartial struct {
	Prefix string
	Fn string
	FullName string
	FileStr string
	Name string
}

func RenderPage(c *gin.Context, partials []TmplPartial, data gin.H) {

	projectRoot, err := FindProjectRoot()
	if err != nil {
		log.Printf("ERROR: Failed to find project root: %v", err)
		c.AbortWithStatus(500)
		return
	}

	var partialsStr []string
	for _, partial := range partials {
		partial.FullName = filepath.Join(projectRoot, "templates", partial.Fn)
		partialsStr = append(partialsStr,  ReadTemplateFile(partial))
	}

	tmpl := template.New("layout")
	for _, part := range partialsStr {
		tmpl, err = tmpl.Parse(part)
		if err != nil {
			log.Fatal(err)
		}
	}

	// 4. Execute the template.
	err = tmpl.Execute(c.Writer, data)
	if err != nil {
		log.Printf("ERROR: Failed to execute template '%s': %v", "layout", err)
		c.AbortWithStatus(500)
	}
}

func ReadTemplateFile(tmpl TmplPartial) string {
	// Read the file into a byte slice, then convert to string
	content, err := os.ReadFile(tmpl.FullName)
	if err != nil {
		log.Fatal(err)
	}
	// templateStr := tmpl.Prefix + string(content) + "\n" + "{{ end }}"
	templateStr := "\n" + tmpl.Prefix + "\n" + string(content)+ "\n"  + "{{ end }}"
	log.Println(templateStr)
	return templateStr
}
