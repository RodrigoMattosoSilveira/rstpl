package render

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"
)

/*
 * Assembles a string consisting of 
   1. {{ define <<name>> }}, where name is the name of a block of templates
   2. {{ template <<template>> . }}, where tempplate is the name of partial template being included in this block of templates
   3. {{ end }}
 */
func GenerateCombo(name string, templates []string) string {
    b := &strings.Builder{}
    fmt.Fprintf(b, `{{define "%s"}}`, name)
    for _, t := range templates {
        fmt.Fprintf(b, `{{template "%s" .}}`, t)
    }
    b.WriteString(`{{end}}`)

    return b.String()
}
func LoadTemplates(tmplRoot string, tmplExt string) *template.Template {
    root := template.New("")

    // Walk views/ recursively
    filepath.Walk(tmplRoot, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        if info.IsDir() {
            return nil
        }
        if strings.HasSuffix(path, "." + tmplExt) {
            _, err := root.ParseFiles(path)
            if err != nil {
                panic(err)
            }
        }
        return nil
    })

    return root
}