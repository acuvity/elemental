package genpython

import (
	"embed"
	"fmt"
	"os"
	"path"
	"strings"
	"text/template"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

//go:embed templates
var f embed.FS

type templatePath string

const (
	initTemplateFile  templatePath = "templates/init.py.tmpl"
	baseTemplateFile  templatePath = "templates/base.py.tmpl"
	modelTemplateFile templatePath = "templates/model.py.tmpl"
)

var functions = template.FuncMap{
	"upper":      strings.ToUpper,
	"lower":      strings.ToLower,
	"capitalize": cases.Title(language.Und, cases.NoLower).String,
	"join":       strings.Join,
	"hasPrefix":  strings.HasPrefix,
}

func makeTemplate(p templatePath) (*template.Template, error) {

	data, err := f.ReadFile(string(p))
	if err != nil {
		return nil, fmt.Errorf("unable to read template file %s: %w. Available entries in templates are: %v", p, err, availableEntries())
	}

	return template.New(path.Base(string(p))).Funcs(functions).Parse(string(data))
}

func writeFile(path string, data []byte) error {

	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("unable to write file %s: %w", path, err)
	}

	// #nosec G307
	defer f.Close() // nolint: errcheck
	if _, err := f.Write(data); err != nil {
		return fmt.Errorf("unable to write file %s: %w", path, err)
	}

	return nil
}

func availableEntries() []string {
	entries, err := f.ReadDir("templates")
	if err != nil {
		panic(err)
	}

	var names []string
	for _, entry := range entries {
		names = append(names, entry.Name())
	}

	return names
}
