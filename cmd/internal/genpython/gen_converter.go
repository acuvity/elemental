package genpython

import (
	"fmt"

	"go.acuvity.ai/regolithe/spec"
	"golang.org/x/sync/errgroup"
)

type converter struct {
	license        string
	public         bool
	splitOutput    bool
	inSpecSet      spec.SpecificationSet
	outDir         string
	resourceToRest map[string]string
}

func newConverter(inSpecSet spec.SpecificationSet, cfg Config, license string) *converter {

	c := &converter{
		license:        licenseHeaderString(license),
		public:         cfg.Public,
		splitOutput:    cfg.SplitOutput,
		outDir:         cfg.OutputDir,
		inSpecSet:      inSpecSet,
		resourceToRest: make(map[string]string),
	}

	for _, spec := range inSpecSet.Specifications() {
		model := spec.Model()
		c.resourceToRest[model.ResourceName] = model.RestName
	}

	return c
}

func (c *converter) Do() error {

	var g errgroup.Group

	// write the __init__.py file
	g.Go(func() error {
		if err := c.writeInit(); err != nil {
			return fmt.Errorf("unable to write __init__.py file: %w", err)
		}
		return nil
	})

	// write the base.py file
	g.Go(func() error {
		if err := c.writeBase(); err != nil {
			return fmt.Errorf("unable to write base.py file: %w", err)
		}
		return nil
	})

	// convert all spec files
	for _, iterSpec := range c.inSpecSet.Specifications() {
		func(s spec.Specification, restName string) {
			g.Go(func() error {
				if err := c.processSpec(s, restName); err != nil {
					return fmt.Errorf("unable to process spec for '%s': %w", restName, err)
				}
				return nil
			})
		}(iterSpec, iterSpec.Model().RestName)
	}

	return g.Wait()
}
