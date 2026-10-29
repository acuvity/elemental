package genopenapi3

import (
	"errors"
	"io"
	"os"
	"path/filepath"
	"testing"

	"go.acuvity.ai/regolithe/spec"
)

func TestConverter_Do__error_bad_externalType_mapping(t *testing.T) {
	t.Parallel()

	specDir := t.TempDir()

	badTypeMapping := replaceTrailingTabsWithDoubleSpaceForYAML(`
		'[]byte':
			openapi3:
				type: malformed-json }
	`)

	rawSpec := replaceTrailingTabsWithDoubleSpaceForYAML(`
		model:
			rest_name: test
			resource_name: tests
			entity_name: Test
			package: None
			group: N/A
			description: dummy.
		attributes:
			v1:
			- name: someField
				description: useful description.
				type: external
				subtype: '[]byte'
				exposed: true
	`)

	for filename, content := range map[string]string{
		"regolithe.ini": regolitheINI,
		"_type.mapping": badTypeMapping,
		"test.spec":     rawSpec,
	} {
		filename = filepath.Join(specDir, filename)
		if err := os.WriteFile(filename, []byte(content), os.ModePerm); err != nil {
			t.Fatalf("error writing temporary file '%s': %v", filename, err)
		}
	}

	spec, err := spec.LoadSpecificationSet(specDir, nil, nil, "openapi3")
	if err != nil {
		t.Fatalf("error parsing spec set from test data: %v", err)
	}

	converter := newConverter(spec, Config{})
	if err := converter.Do(nil); !errors.Is(err, errUnmarshalingExternalType) {
		t.Fatalf("unexpected error\nwant: %v\n got: %v", errUnmarshalingExternalType, err)
	}
}

func TestConverter_Do__error_writer(t *testing.T) {

	specDir := t.TempDir()

	rawSpec := replaceTrailingTabsWithDoubleSpaceForYAML(`
		model:
			rest_name: test
			resource_name: tests
			entity_name: Test
			package: None
			group: N/A
			description: dummy.
	`)

	for filename, content := range map[string]string{
		"regolithe.ini": regolitheINI,
		"_type.mapping": typeMapping,
		"test.spec":     rawSpec,
	} {
		filename = filepath.Join(specDir, filename)
		if err := os.WriteFile(filename, []byte(content), os.ModePerm); err != nil {
			t.Fatalf("error writing temporary file '%s': %v", filename, err)
		}
	}

	spec, err := spec.LoadSpecificationSet(specDir, nil, nil, "openapi3")
	if err != nil {
		t.Fatalf("error parsing spec set from test data: %v", err)
	}

	simulatedErr1 := errors.New("simulated error 1")
	fw := &fakeWriter{wrErr: simulatedErr1}
	writerFactory := func(string) (io.WriteCloser, error) { return fw, nil }
	converter := newConverter(spec, Config{})
	if err := converter.Do(writerFactory); !errors.Is(err, simulatedErr1) {
		t.Fatalf("unexpected error\nwant: %v\n got: %v", simulatedErr1, err)
	}

	simulatedErr2 := errors.New("simulated error 2")
	writerFactory = func(string) (io.WriteCloser, error) { return nil, simulatedErr2 }
	converter = newConverter(spec, Config{})
	if err := converter.Do(writerFactory); !errors.Is(err, simulatedErr2) {
		t.Fatalf("unexpected error\nwant: %v\n got: %v", simulatedErr2, err)
	}
}

func TestParseExamples__panics(t *testing.T) {
	t.Parallel()

	cases := map[string]any{
		"not a map":              []any{"nope"},
		"example not a map":      map[string]any{"ex": "nope"},
		"missing value":          map[string]any{"ex": map[string]any{"summary": "no value"}},
		"non-string summary":     map[string]any{"ex": map[string]any{"value": 1, "summary": 42}},
		"non-string description": map[string]any{"ex": map[string]any{"value": 1, "description": true}},
	}

	for name, entry := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			defer func() {
				if recover() == nil {
					t.Fatalf("expected parseExamples to panic")
				}
			}()

			parseExamples("test", entry)
		})
	}
}

func TestGetResponseMapExtension__malformedExamples(t *testing.T) {
	t.Parallel()

	defer func() {
		if recover() == nil {
			t.Fatalf("expected getResponseMapExtension to panic")
		}
	}()

	getResponseMapExtension(map[string]any{
		responseMapExtensionName: map[string]any{
			"200": map[string]any{"examples": "nope"},
		},
	})
}
