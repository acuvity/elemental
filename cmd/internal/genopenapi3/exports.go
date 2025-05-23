package genopenapi3

import (
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"

	"go.acuvity.ai/regolithe/spec"
)

// Config is used to guide the generator function
type Config struct {
	Public      bool
	SplitOutput bool
	OutputDir   string
}

// GeneratorFunc will convert the given spec set into an openapi3 document
func GeneratorFunc(sets []spec.SpecificationSet, cfg Config) error {

	outFolder := path.Join(cfg.OutputDir, "openapi3")
	if err := os.MkdirAll(outFolder, 0750); err != nil && !os.IsExist(err) {
		return fmt.Errorf("'%s': error creating directory: %w", outFolder, err)
	}

	newFileFunc := func(name string) (io.WriteCloser, error) {
		filename := filepath.Join(outFolder, name)
		file, err := os.Create(filename)
		if err != nil {
			return nil, fmt.Errorf("'%s': error creating file: %w", filename, err)
		}
		return file, nil
	}

	set := sets[0]
	converter := newConverter(set, cfg)
	if err := converter.Do(newFileFunc); err != nil {
		return fmt.Errorf("error generating openapi3 document from spec set '%s': %w", set.Configuration().Name, err)
	}

	return nil
}
