package genopenapi3

import (
	"errors"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"go.aporeto.io/regolithe/spec"
)

func TestConverter_Do__error_bad_externalType_mapping(t *testing.T) {
	t.Parallel()

	specDir, err := ioutil.TempDir("", t.Name()+"_*")
	if err != nil {
		t.Fatalf("error creating temporary directory for test function: %v", err)
	}
	t.Cleanup(func() {
		if err := os.RemoveAll(specDir); err != nil {
			// no need to fail the test; it is just a temporary dir that
			// the OS will eventually destroy, but let's log the error
			t.Logf("error removing temporary test directory: %v", err)
		}
	})

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
		if err := ioutil.WriteFile(filename, []byte(content), os.ModePerm); err != nil {
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

	specDir, err := ioutil.TempDir("", t.Name()+"_*")
	if err != nil {
		t.Fatalf("error creating temporary directory for test function: %v", err)
	}
	t.Cleanup(func() {
		if err := os.RemoveAll(specDir); err != nil {
			// no need to fail the test; it is just a temporary dir that
			// the OS will eventually destroy, but let's log the error
			t.Logf("error removing temporary test directory: %v", err)
		}
	})

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
		if err := ioutil.WriteFile(filename, []byte(content), os.ModePerm); err != nil {
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
