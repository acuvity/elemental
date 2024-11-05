package genpython

import (
	"fmt"
	"os"

	"go.acuvity.ai/regolithe/spec"
)

// Config is used to guide the generator function
type Config struct {
	Public      bool
	SplitOutput bool
	LicenseFile string
	OutputDir   string
}

// GeneratorFunc will convert the given spec set into python types
func GeneratorFunc(sets []spec.SpecificationSet, cfg Config) error {

	if err := os.MkdirAll(cfg.OutputDir, 0750); err != nil && !os.IsExist(err) {
		return fmt.Errorf("'%s': error creating directory: %w", cfg.OutputDir, err)
	}

	var licenseBytes []byte
	var err error
	if cfg.LicenseFile != "" {
		licenseBytes, err = os.ReadFile(cfg.LicenseFile)
		if err != nil {
			return fmt.Errorf("error reading license file '%s': %w", cfg.LicenseFile, err)
		}
	}

	set := sets[0]
	converter := newConverter(set, cfg, string(licenseBytes))
	if err := converter.Do(); err != nil {
		return fmt.Errorf("error generating python types from spec set '%s': %w", set.Configuration().Name, err)
	}

	return nil
}
