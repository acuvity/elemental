package genpython

import (
	"bufio"
	"bytes"
	"fmt"
	"path"
	"strings"

	"go.acuvity.ai/regolithe/spec"
)

func licenseHeaderString(license string) string {
	var lines []string
	sc := bufio.NewScanner(strings.NewReader(license))
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	var b strings.Builder
	for _, l := range lines {
		fmt.Fprintf(&b, "# %s\n", l)
	}
	b.WriteString("#\n#\n")
	return b.String()
}

func (c *converter) writeInit() error {

	tmpl, err := makeTemplate(initTemplateFile)
	if err != nil {
		return err
	}

	var buf bytes.Buffer

	if err = tmpl.Execute(
		&buf,
		struct {
			License string
			Set     spec.SpecificationSet
		}{
			License: c.license,
			Set:     c.inSpecSet,
		}); err != nil {
		return fmt.Errorf("unable to generate __init__.py: %s", err)
	}

	if err := writeFile(path.Join(c.outDir, "__init__.py"), buf.Bytes()); err != nil {
		return fmt.Errorf("unable to write __init__.py file: %s", err)
	}

	return nil
}

func (c *converter) writeBase() error {

	tmpl, err := makeTemplate(baseTemplateFile)
	if err != nil {
		return err
	}

	var buf bytes.Buffer

	if err = tmpl.Execute(
		&buf,
		struct {
			License string
		}{
			License: c.license,
		}); err != nil {
		return fmt.Errorf("unable to generate base.py: %s", err)
	}

	if err := writeFile(path.Join(c.outDir, "base.py"), buf.Bytes()); err != nil {
		return fmt.Errorf("unable to write base.py file: %s", err)
	}

	return nil
}

func (c *converter) processSpec(s spec.Specification, name string) error {

	tmpl, err := makeTemplate(modelTemplateFile)
	if err != nil {
		return err
	}

	model := s.Model()

	if s.Model().Private && c.public {
		return nil
	}

	var buf bytes.Buffer

	if err = tmpl.Execute(
		&buf,
		struct {
			License    string
			Set        spec.SpecificationSet
			Spec       spec.Specification
			Model      *spec.Model
			PublicMode bool
		}{
			License:    c.license,
			PublicMode: c.public,
			Set:        c.inSpecSet,
			Spec:       s,
			Model:      model,
		}); err != nil {
		return fmt.Errorf("unable to generate model '%s': %s", name, err)
	}

	if err := writeFile(path.Join(c.outDir, name+".py"), buf.Bytes()); err != nil {
		return fmt.Errorf("unable to write file for spec: %s", name)
	}

	return nil
}
