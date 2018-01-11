// Code generated by go-bindata.
// sources:
// templates/README.md
// templates/identities_registry.gotpl
// templates/model.gotpl
// templates/relationships_registry.gotpl
// DO NOT EDIT!

package static

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _templatesReadmeMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x04\xc0\xcb\x11\x02\x21\x10\x45\xd1\x7d\x47\xf1\x2c\x53\x22\x81\xc6\xbe\x0a\xc5\x47\x6a\x60\x33\xd9\xcf\x79\x2b\x31\x56\xf7\xc3\x36\x4b\x85\x8d\xbe\xb5\xa3\x09\xa1\xf3\x57\x46\x8c\x4c\x04\xa1\x3a\x75\x0a\xca\x75\xfa\x75\xbf\xcc\x24\x69\x78\x43\xcb\x3f\xcd\x7f\xd8\x13\x00\x00\xff\xff\xaa\x97\xff\x85\x4d\x00\x00\x00")

func templatesReadmeMdBytes() ([]byte, error) {
	return bindataRead(
		_templatesReadmeMd,
		"templates/README.md",
	)
}

func templatesReadmeMd() (*asset, error) {
	bytes, err := templatesReadmeMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/README.md", size: 77, mode: os.FileMode(420), modTime: time.Unix(1515709395, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesIdentities_registryGotpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x56\xcf\x6b\xe3\x3a\x10\xbe\xfb\xaf\x18\x42\x79\x38\xd0\x38\x97\xc7\x3b\x14\x7a\x28\xe5\x15\x02\xdb\x52\x5a\xd8\x4b\xe9\x41\x75\xc6\xee\xb0\xb2\x14\xa4\x49\xbb\x45\xf8\x7f\x5f\xfc\x43\xb2\x93\x38\x5e\xef\xa6\xbb\x39\xd9\xf2\xcc\x37\x9f\xe6\xfb\x46\xca\x46\xa4\xdf\x44\x8e\xe0\x1c\x24\x8f\xc8\xc9\xb5\x56\x19\xe5\x5b\x23\x98\xb4\x4a\xee\x44\x81\x50\x96\x51\x44\xc5\x46\x1b\x86\x59\x4e\xfc\xba\x7d\x49\x52\x5d\x2c\xc5\x46\x1b\x64\xbd\x20\x95\x2e\x51\x62\x81\x8a\x85\x9c\x45\x51\xb6\x55\x29\x90\x22\x8e\xe7\xe0\x22\x00\xa8\xb0\x8d\x50\x39\x36\x15\x1e\x37\x98\x52\x46\x69\x5d\xc1\x56\xe8\x55\x4c\x40\x48\x1e\x30\x27\xcb\x68\x56\x6b\x54\x4c\xfc\x11\x3b\x97\xfc\x5f\x3f\x55\x6c\xca\xd2\xaf\xcf\x9d\x03\x54\xeb\x0a\xa0\x8c\xa2\xe5\x12\x6e\xf5\x1a\xe5\x57\x34\x96\xb4\x02\x83\xbc\x35\xca\x02\xbf\x22\xa4\x5b\x63\x50\x31\xbc\xb5\xdf\x74\x56\x2f\x17\x55\x7c\xd2\xf0\xed\xe7\xc6\x73\xc8\xa4\x16\xfc\xdf\xbf\xe0\x5a\x9c\xd0\x9e\xab\xfb\xd5\x4a\x65\x3a\xf1\x65\xca\x12\x9a\xe2\x0d\xab\x8c\xc4\x8b\xc4\x1b\x1d\xc8\x07\x1e\x02\x14\xbe\x03\x29\xcb\x42\xa5\xe8\x29\xf4\xb3\x20\xd3\xa6\x5e\xcc\xe9\x0d\x15\x90\x47\x50\xa2\xc0\x96\xe5\x91\x22\x71\x88\xb5\x6c\x48\xe5\xf3\x5e\x37\xfb\x29\x2e\x8a\x00\xec\x3b\x71\xfa\xda\xc1\x4f\x97\x08\x20\x15\xb6\x71\x4a\xa7\x07\x74\x82\xd4\x6e\xb9\x68\x43\xc1\x77\xee\x0e\xdf\x0f\x32\xe2\x79\x5b\x74\xe1\x15\xac\x5e\xd7\x98\x89\xad\x64\x8f\xd0\xe6\x2b\x92\x11\x80\xd7\x78\xaf\x03\xd7\x82\x31\xd7\xe6\x84\x36\xa7\x1e\xe1\x78\x9b\x7d\x91\x38\xc4\x4e\x6f\x73\x48\xf9\xc4\x36\x7b\x3e\x7f\xb8\xd5\xd7\x5a\x31\x2a\xfe\x75\x63\x8b\xa1\xd4\x09\xee\x1e\x2f\x38\x66\xf2\x81\xcc\x93\xbd\xee\x1c\x50\x06\x4a\x33\x24\x2b\xfb\xa0\x35\xc3\xe2\xa4\x39\xf8\x67\x37\xfe\x5e\x6e\x8d\x90\x50\x96\x5f\xc8\xb2\xeb\x8a\xee\xc8\xf4\x69\xaa\x4d\x98\x93\x29\xaa\x0d\x0d\xcb\x78\xc1\xb1\x99\xf9\x89\x6a\xbf\x35\x3a\x27\xa8\x76\x74\xac\xfe\x92\x72\x57\x52\xb6\x5c\x08\x6d\x27\x94\x94\x80\xdf\xc9\x32\xa9\xdc\x1b\x99\xd0\xb6\xcd\xdf\xc9\x89\xe7\xf0\xf4\xbc\x7f\x24\xd5\xae\x8f\x42\xc9\xa1\x00\x17\x18\x4f\x6b\xf0\x70\xff\xce\x0f\xf7\x5d\xef\xec\x4d\x18\x10\x92\x84\x45\x7b\x2b\x36\x70\x09\x85\xd8\x3c\x35\x66\x18\x26\x3b\x91\x4a\x15\x74\x86\x1d\x93\x8b\xcb\x1d\x62\x8b\x10\xd6\x42\x9d\xd1\x39\x9c\xd5\x44\xea\xd0\xab\x86\x92\x47\x9b\x39\xe7\xbf\x96\xe5\xec\xa2\x4a\xeb\x83\x8f\x6e\xb3\xf7\xd2\xbf\xa4\xf8\xe3\xc6\xe8\xa2\xae\xb3\xf3\x6f\x24\x6c\x55\x58\xab\x53\x12\x8c\x6b\x60\xdd\x9b\xb2\x9a\xc6\xce\x55\xd4\x83\x8a\x1b\x92\xc7\x6e\xa0\x3d\xb9\xbb\xc6\x3f\xd5\x8f\xcf\xc1\x6b\xf5\xfa\xe0\x79\x2e\x65\xcd\xa5\x4d\x3d\x72\x6e\x07\xff\xed\xe3\x74\xc7\xf4\x21\xb5\xca\x9f\x0d\x71\x18\x3d\x9e\x27\xa8\x3f\x3e\xcb\x7b\x73\xe6\xab\xba\xa9\x6e\x38\x74\xc4\x79\x58\xdf\x9b\xec\xea\x37\x30\xf2\x65\x4f\x84\x6a\xcc\xcb\xe8\x47\x00\x00\x00\xff\xff\x3a\xcb\x38\xd5\x69\x0b\x00\x00")

func templatesIdentities_registryGotplBytes() ([]byte, error) {
	return bindataRead(
		_templatesIdentities_registryGotpl,
		"templates/identities_registry.gotpl",
	)
}

func templatesIdentities_registryGotpl() (*asset, error) {
	bytes, err := templatesIdentities_registryGotplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/identities_registry.gotpl", size: 2921, mode: os.FileMode(420), modTime: time.Unix(1515695604, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesModelGotpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5a\xdd\x6f\xdb\x38\x12\x7f\xf7\x5f\xc1\x13\xba\xbb\xf6\xc1\x51\xfa\xec\xbd\x2c\x50\xa4\xe9\x22\xb8\x76\x5b\xd4\xbd\xde\x43\x51\x6c\x18\x69\xec\xb0\x95\x48\x95\xa2\xd2\xe4\x0c\xfd\xef\x07\x7e\x8a\xd4\x87\x25\x67\xd3\x6d\x1f\xfc\x12\xc4\x43\xf2\x37\x9f\x1c\xce\x90\x2a\x70\xf2\x19\x6f\x01\xed\x76\x28\x5e\x83\x88\xcf\x19\xdd\x90\x6d\xc5\xb1\x20\x8c\xc6\x7f\xe0\x1c\x50\x5d\xcf\x66\x24\x2f\x18\x17\x68\x3e\x43\x08\xa1\x68\x93\x8b\x48\xff\x57\xde\xd3\x24\x9a\xe9\xff\xb7\x44\xdc\x54\xd7\x71\xc2\xf2\x53\x5c\x30\x0e\x82\x9d\x10\x9a\x9c\x42\x06\x39\x50\x81\x33\xbd\x64\xb7\x43\x1c\xd3\x2d\xa0\x78\x5d\x40\x12\xbf\xbb\x2f\xe0\x0d\x67\xb7\x24\x05\x5e\xa2\x93\xba\xd6\x58\x52\x1c\x54\xd7\x6e\x09\xd0\x54\x0d\x2e\x66\xb3\x99\x43\x78\x42\x96\xe8\x09\xd0\x2a\x47\xab\x33\x14\x5f\xd0\x2a\x2f\xa5\xb0\xa7\xa7\x72\x85\x1a\x50\xf0\xa8\xae\x11\x87\x82\x43\x09\x54\x94\x48\xdc\x00\x2a\x58\x59\x92\xeb\x0c\xd0\x2d\xce\x2a\x28\xd1\x86\x71\x84\x85\xe0\xe4\xba\x12\xa0\xb8\xeb\xe5\xcf\x2c\xcd\x18\x22\x8a\x67\x42\x22\x76\xf0\x4b\xc1\x09\xdd\xce\x66\x09\xa3\xa5\x35\xd3\x6e\x77\x62\x05\xa5\x38\x87\x25\x7a\xa2\xb8\x49\x61\xf5\xe2\xf7\x9a\xb9\xd1\xd9\x88\x4d\x35\xa7\xb6\xc4\x7a\xa9\x9c\xa0\xff\xab\xeb\xd8\xda\xc6\x2d\xe9\x48\x75\xa6\x55\xb1\x2b\x02\x6b\x2a\x63\x36\xff\x1b\xab\x69\xa7\x5c\x50\x41\xc4\xbd\xd1\xf9\x32\x05\xf5\xb3\x2d\x91\xa3\xb3\x8d\xfa\xcd\xae\x3f\x41\x22\xe2\xd9\x2d\xe6\x63\x48\x67\xc8\x05\x45\xec\x88\x3b\x25\x9d\x9c\xba\x42\x2e\x06\x14\xc8\x5b\x28\x85\xa4\xd7\x75\xb4\x54\x93\xce\xb1\x80\x2d\xe3\xf7\xab\x70\x12\xab\x78\xe2\x3c\xb5\x9c\xd5\x3a\x56\xc8\x06\x51\x26\xcc\xac\xcb\xf2\x2d\x63\xa2\x89\x92\xb6\x94\x6f\xb2\x8a\xe3\x0c\xd5\xf5\x4b\x52\x0a\x5f\x63\x8c\x32\x49\x61\x9b\xbd\xab\x5c\x74\xec\xc7\xfd\xf0\xf1\x9f\x3d\x73\x8c\x17\xce\x19\x15\x40\x85\x67\x76\x51\x71\xaa\x6d\x4e\x7a\x6d\x5e\x22\x42\xd5\x4f\x29\x62\x3c\xdb\x54\x34\x41\x73\x36\x2a\xc6\xa2\xcd\x6a\xbe\xe8\xf7\x8b\xb2\xb9\x16\xa3\x0f\xb4\x71\xec\xcc\x6a\x50\x34\x62\x63\x54\x30\x42\x05\x70\x24\x18\xc2\x28\x91\x63\x52\xd6\x31\xe9\x0e\xd3\xa3\x08\x85\x0f\x14\xdb\x10\x2c\xb7\xba\xd1\x43\x09\xb0\x3a\x43\xb8\x28\x80\xa6\xf3\x31\xf0\x5d\xbd\x44\x2c\x8e\xe3\x85\x6f\x84\x9f\x25\x88\x51\xf6\x99\xc2\x31\x70\x65\xe0\x15\xc1\xd4\x4f\x8c\x28\x7c\xd5\x7c\x8d\xdb\x1e\x53\x77\xcd\x7f\x6e\x79\xc6\x71\xdc\xf6\xa1\xd6\x7f\xa2\x79\x58\x25\x1e\x6c\x1d\x99\x45\xff\x5c\x4a\xf5\x25\x84\x4e\x7d\x56\x2e\xbd\xbb\x2d\x07\xc7\x80\x55\x42\x2d\x88\xe7\xfd\x3b\x62\xa1\x91\xeb\x20\x06\x59\x25\x8c\xf1\xd5\x6e\x4a\x18\xbd\x05\x2e\x7c\xdb\xab\x58\xa3\x9d\x68\xd6\xaa\x96\x87\x9a\x58\xfe\xed\xd9\x1b\x1e\x5a\xcb\x7a\x7b\x66\xee\x6a\xdf\x54\x44\x40\xee\xd9\x6a\xaf\x95\xe4\xdc\xfd\xd6\x78\x0e\x1b\x5c\x65\xe2\x35\x4f\x81\x07\x69\x23\xd5\x03\x88\xc9\x11\x42\xb7\x68\x43\x20\x4b\x4b\x1b\x8e\x89\x0e\x87\x43\x4c\xe2\xb3\x9a\x2f\xd0\x87\x8f\xfa\x00\x6c\x25\x0b\x4b\x6e\xd4\x6a\x9d\xfe\xaf\x8d\x40\xee\x9c\x6d\x4a\x00\x77\x04\x34\x09\xdd\x03\x31\xa7\x96\xb6\x86\xd6\xfe\x3d\xf0\x92\x30\x1a\x28\x7e\x6b\x68\x0f\x57\xd4\xa0\xce\x17\x88\x50\xe3\x66\x3f\x13\x82\x88\x9f\xbd\xb9\xbc\xa4\x1b\x16\x5b\xfe\xb5\x12\xc8\x3b\x59\x87\x8f\xd6\xf6\x91\x9a\xb3\x14\x32\x29\x2c\x46\x9d\xc3\x6f\xdf\xe9\x62\x2a\x90\x2a\x91\x12\xee\x76\xbe\x85\x5b\x96\xdd\xed\x50\x8e\x3f\x83\xa4\xaa\x02\x6b\x26\x8b\x14\x2b\xa8\xb4\xe6\x2b\x29\x82\x55\x45\xaa\x7c\xf5\xa9\x64\x74\x15\x9d\x44\xe8\x5a\xfd\xf3\xa7\x12\xd2\x18\x36\xba\xd2\xab\x64\x19\x18\xbf\xaa\x04\xdc\x19\x67\xfc\x01\x5f\x07\x55\xb6\xe7\x82\xcc\x8a\x7d\xbb\x5e\xca\xa2\x1c\x34\x08\x32\x5f\x0c\x2d\x6c\x05\xe0\xcf\xfd\xb3\x9a\x78\xf4\xd5\x5d\xed\x71\xe9\xb2\x1b\xc1\xb6\xa8\x53\x25\x9f\xb4\x3d\xe3\xaa\x0e\x0d\xcd\x7e\x49\x89\x20\x38\x23\xff\xf3\xcb\xdb\x56\xe1\xa6\x18\x07\x38\x2d\x86\x3d\xd1\xde\x5b\x1e\x0c\x95\x64\x36\xdc\x07\x8c\xb6\x40\x7f\xb1\x02\x40\x9d\x12\xc0\xe6\xbd\x56\x22\xd2\x95\x68\x20\xdd\x2f\x25\xaa\x28\xf9\x52\xd9\xe2\x46\xae\x99\x2c\xb1\x9c\x3c\x5f\xa0\x30\xf9\xe8\x9a\xcf\xd4\x7b\x8d\x1c\xd6\xf8\x36\x6d\xc6\x0e\xba\x99\x24\x4f\x46\x79\x96\x40\x6a\x03\xd1\x16\xf2\x90\x95\xd0\x86\x88\xa2\x66\x58\x7b\x48\x2b\xbf\x06\xe1\xf1\x2d\x41\x3c\xb6\xf2\x01\x83\x39\x49\x8d\x01\x16\x13\x2d\x30\x4d\x75\x74\x86\x48\xda\xaf\x60\x5f\xaa\xbd\xc1\x3c\x4d\x58\x0a\x69\x3b\xe9\xaa\x7c\x31\x41\xa9\x9e\x4c\x3b\x31\xd7\xee\x3b\xfa\x6c\xd5\x3e\x70\x04\x8e\xc8\xa5\x04\xfb\xd1\xce\x3a\xdf\xbb\xcf\xa1\x4c\x38\x29\x84\x31\x86\xb4\x04\x4b\xc2\xb3\x9f\x25\x95\xda\xd0\x6a\x8e\xac\x3b\x9a\xf0\x1b\x77\xca\x73\x96\xf4\xec\xaf\x13\x29\x00\x7c\xe9\x91\x21\xfa\x40\x59\xca\x92\x8f\x51\x7b\xaf\x28\xf2\x5a\x77\xc9\xc1\x9e\x0a\xa7\x5d\x39\x41\x42\xcd\xae\xba\x81\xe8\x9d\xb0\x4a\x8d\xfd\x3b\x46\x71\xee\xa8\x62\xb8\x6e\x72\x11\xaf\x0b\x4e\xa8\xd8\xcc\xa3\x7f\xfd\x54\xae\x7e\x2a\x7f\x8b\x64\x65\xdb\xe4\x45\xe5\x9a\x86\xa4\x13\xcf\xc2\xb8\x63\xe0\xc0\xd5\xe7\xad\xf4\xd5\xef\x20\x64\x07\xa4\x3d\xf4\x3b\x08\x29\x66\x67\xbf\xf9\x5e\xeb\x9d\x60\x36\x14\x87\x04\xc8\x6d\x3b\x51\x3c\xe9\xd5\x7b\x80\xd7\x7c\x11\x72\xb0\x17\x06\xa1\x59\x74\x9e\xe8\xa4\xc4\xa0\xba\xb1\xc1\xe8\x2b\xb8\x1e\x50\xd0\x65\xc2\x2d\xb9\x05\xfa\x68\x3a\x0e\xb0\x9b\x7b\x1b\xaa\x57\x5b\x97\x2e\xfb\xf5\x44\x67\xc8\x43\x08\x02\x2e\xbc\x36\x79\x8f\x33\x92\x62\xa1\x12\x3c\x49\x41\xab\x98\x54\x9c\x03\x15\x88\xd0\x0d\xe3\xb9\xde\x7c\xa5\x60\x1c\x52\x99\xde\x74\x63\xa8\x8f\xfb\x8a\xc3\x94\xec\x68\x98\xc8\x13\x9a\x73\xc6\xad\xec\xea\x47\x19\x36\x1d\x17\x8a\xb6\xb3\x1b\xeb\x4b\x45\x38\xa4\x17\xfb\x26\xf6\xdd\xcc\x85\x81\xdc\x1c\x2b\xea\x32\xe5\x1d\xc7\xb4\x24\x52\xc1\x60\x2c\xbe\xb8\x2b\x58\x09\x4d\x39\x69\xc8\x6f\x8d\x10\xe1\x6c\x99\x41\x94\x2f\x22\xbd\x2b\x23\x3b\x2c\xc7\x38\x0f\x65\xb5\x06\xb0\x50\x66\x47\x87\x69\x73\xc0\x95\x8b\x5f\x15\xde\x3f\xce\x10\x25\x99\xd7\x64\xb5\x6c\xe3\xfa\xad\x90\xbe\x94\x8b\x6d\xe7\x15\xe6\xe4\x5e\x65\x04\xc9\xe1\x20\x55\xde\x91\x1c\x7e\x44\x45\xe0\x4e\x00\xa7\x38\x3b\x48\x99\x0b\xb3\xe8\x3b\x2b\x34\xa8\x5e\xfc\x2c\xcb\xd8\x57\x48\xcf\x6f\x18\x49\x9a\xd8\xde\xa7\x9a\x0e\xb5\x4b\xaa\xee\x01\x5a\x6a\xe9\xc8\x9d\x0f\x68\xb7\x6c\xca\x02\xb9\xee\x13\x23\xb4\x23\xc0\x55\xb4\x44\xd1\x95\x44\xab\x97\x2a\xe3\x3c\xab\x04\xdb\x02\x05\x8e\x85\xda\x31\x43\x36\x82\x96\x6d\xe0\x00\x27\x37\x42\x60\x3e\xc9\x06\x6f\xb0\x4c\xf0\x74\x9a\x57\x97\xfa\x1c\x6f\xf1\xb8\xd2\xea\x79\xb9\xe0\x5b\x68\xb6\x15\x28\x7e\x85\xef\x5e\x02\xdd\x8a\x1b\xf4\x74\x8a\x6e\xaf\xf0\x1d\xc9\xab\x5c\x2f\x99\xaa\xa1\xa4\x36\x7c\x24\x65\x83\xb3\x12\xbe\x99\x4a\x84\x1e\xa4\x12\xa1\x0f\x54\xc9\xf1\xf9\xf6\x2a\xe1\x3b\xf5\x02\x82\x9e\xc6\x4f\x87\x0e\x86\x4d\xc6\xb0\x98\x94\x7f\x8c\x13\x5f\xc8\x05\x07\xfa\xf0\xbd\x79\x1f\x79\x3c\x7d\x4d\x61\x3b\x55\xe8\x4b\x3a\x59\x64\x42\xc5\xbc\x25\xf6\xe2\xb1\xfd\x34\x16\x88\x8f\xe9\x35\x1d\xa7\x87\x7b\xcd\x4a\xf1\x0d\xbc\x36\x51\xe6\x87\x38\xad\x91\xfa\xef\x73\xda\x8f\x5a\x7c\xfd\x95\xf4\xf1\x5d\x8b\xad\xc7\x11\xfc\x3b\x16\x57\x8f\x14\x5b\xbd\x83\x3d\x53\xc9\x06\x65\x40\x5b\xa5\xdb\x02\xfd\x86\x9e\x3a\x99\x4c\xc3\x19\x4e\xf1\xdf\x3a\x0c\x06\xf4\xac\x75\xab\xa1\xb3\xca\xde\x3d\x90\xcc\xde\xcc\x15\x90\x90\x0d\x49\x54\x3f\xf6\x82\x71\xd7\xe3\x04\xcd\xb7\xa3\x06\xd3\xdd\xdd\x89\xee\x5e\x9b\xcf\x05\xd4\x2d\xee\x67\xb8\xb7\x5d\xdc\xf0\x25\xc4\x10\xf7\xb9\x82\xb0\x97\x78\x8d\xfb\x07\x04\x51\xad\x1f\xd9\xa0\xdb\x25\x62\x9f\x65\xc0\xf4\x33\x6c\x1a\xb8\x57\xb8\xf8\x20\x59\x7c\xfc\x55\x2e\xd8\xf9\xb6\xb9\x9d\x19\x6b\x9d\x9e\xa2\xff\x02\x4a\x58\x95\xa5\xaa\xc5\xdb\x10\x9a\x22\x22\x96\xa8\x64\x28\x03\xf1\x4b\x89\x92\x1b\x48\x3e\x23\x66\x9e\x9a\xd9\x57\xe0\x28\xc1\x25\x20\x42\x53\xb8\x83\x14\x95\x05\x24\x28\xc7\xc5\x6c\xe4\xaa\xf8\xa5\x5c\x7a\x8e\x4b\xe8\x11\xd0\x3e\xaa\xf6\x2a\x5e\x06\x5e\xda\x54\x59\xe6\x79\xa1\x0c\x67\xe6\xb8\x18\xf5\xc7\x00\x97\xf9\x42\xae\xfe\xa0\xdd\xf1\x71\x9a\x37\xf6\x2a\x1c\xe8\x39\xdb\xf7\xad\x45\x30\xb3\xf3\x3a\x84\x0b\xf5\x36\xe4\x54\x96\x01\xd9\x8f\xb3\xef\x2b\x8c\x90\xc7\xd9\x21\xba\x2a\x55\x9b\xcf\x5a\xec\xbd\x26\xd9\x12\x8a\xb3\x35\x93\xa9\xa7\x7b\x6b\x10\xf5\xa5\xa6\x68\x35\x1a\xe6\x76\x7f\x9b\xeb\xc6\xde\x8e\x05\x21\x9f\xba\xea\xb6\x1d\xb2\xeb\xf0\x70\xbc\x24\xe5\x2d\x55\x7d\xd8\x6a\xa0\x3f\x3c\xa9\xeb\x83\x7a\xb8\xa6\x90\x70\xcb\x6a\x97\x10\x97\x6d\x9d\x5a\xad\x9e\x95\xcc\x27\xaf\x7a\x9b\xc2\x41\xad\x1c\xf4\x2d\x26\x19\xbe\x26\x19\x11\xf7\x1e\xb2\x47\x35\x5f\xcf\xb4\x26\x46\xa3\xc8\xe7\x37\x98\x52\xc8\x9a\x01\x43\x30\x78\xcd\xf0\x04\x28\x0e\xca\xd7\xaf\x69\xe6\x09\xe9\x53\xb5\xf6\xad\x79\xa3\xb8\xc1\x73\x80\x1b\xf5\xa9\x1a\xb7\x35\x6f\x2a\xae\x2d\xe2\xdc\xe8\xd0\xd1\x7e\xd2\xe6\xad\x56\xba\x50\x6b\xc1\x49\x89\xbc\xf7\x42\x4d\x31\xc1\x44\x49\xd6\x1b\x45\x6a\x38\xe4\x4e\xab\xfc\x21\x9c\xdb\x5b\x54\xa2\xed\x76\x28\xc1\x05\x11\x4a\x20\x34\xd7\x49\x3f\x58\xbc\x78\xa0\x94\xb6\xe8\x3c\x44\xce\x93\xba\x8e\x3c\xaf\x59\xe1\xa3\x89\xfc\x0f\x64\xd5\xc3\x29\x60\x84\xfa\x39\xb5\xc3\x66\x20\x8e\x0a\x0e\x49\xb8\xe9\x1b\x9a\x8d\x4d\x6f\xce\x84\xc8\x0c\x1e\x82\x2c\xa4\x23\x9a\xdc\xd8\x7a\x55\x19\x85\xf5\xee\x73\xf5\x80\x21\x68\x11\x9b\xd1\x51\xa0\x17\x24\x13\xc0\xd5\x97\x46\x6e\xac\xa1\x69\xb8\x60\xce\x38\x22\xe3\x40\xb6\xf4\xdf\xe0\x65\x8e\x86\x66\x10\xfd\x39\x53\x10\x73\x2c\x02\xb4\x1c\x0b\x93\xd6\xdc\xe0\x78\x56\x6b\x5e\x7b\x34\x5d\xff\xd6\x12\xb9\xb1\x51\x14\xef\xe1\xd6\x8d\x35\x34\x8d\x16\xcc\x19\x47\x94\x85\x9a\x07\x26\x7f\x1a\x1c\x33\x32\x0a\xe1\xdf\x78\x99\x21\x47\x5a\x75\xef\xc4\x26\xc0\xb5\x72\xa9\xa5\xac\x3a\x97\x33\xa3\x58\xde\xcd\x95\x05\xb3\xa4\x55\xf7\x6e\x6b\x02\x5c\x5b\x34\x43\x59\x75\x6e\x20\xc6\xb0\xfc\x8f\x09\xec\x27\xae\xbd\xef\xbd\x83\x08\xea\x84\x0a\x77\x8f\x23\x69\x79\xfc\x19\xa3\x70\x6f\x38\xc9\x31\xbf\x0f\xf6\x4e\x43\xd3\x80\xc1\x9c\x51\xc4\xb7\x80\xd3\xf0\x14\xb7\x94\x95\xb9\xf5\x75\xe3\x13\xb0\xc2\x0b\x0b\x89\xa5\x29\xab\xf6\x0d\xf2\x28\xd6\xba\xb5\x17\xd7\xde\x5e\x5c\x4f\xde\x8b\x6b\xfd\x72\xd7\xa0\xa8\xdf\x06\xc5\x8e\x8d\xa3\x54\xd7\xe6\xe9\xd1\xc2\x68\x82\xfd\x8c\xd9\x0d\x8f\xc7\x43\xe7\xfd\x0d\x21\x47\xd2\x62\xf9\x33\xc6\xe1\x02\xb1\x3c\x99\x26\x0b\xf4\x1f\xfd\x19\x8b\xa3\xeb\xdf\x52\x14\x33\x34\x28\x85\xa2\x77\x3e\x31\x39\xa4\x75\xfc\x46\x8d\xd2\x00\xb3\xbf\xa7\x63\x32\xa5\xd6\xb1\x5d\x42\xc7\x76\xe9\xd8\x2e\x1d\xdb\xa5\x63\xbb\x74\x6c\x97\x8e\xed\xd2\xb1\x5d\x3a\xb6\x4b\xc7\x76\xe9\xd8\x2e\x1d\xdb\x25\xaf\x5d\xfa\x7f\x00\x00\x00\xff\xff\xf9\x7f\x94\x44\xc3\x3d\x00\x00")

func templatesModelGotplBytes() ([]byte, error) {
	return bindataRead(
		_templatesModelGotpl,
		"templates/model.gotpl",
	)
}

func templatesModelGotpl() (*asset, error) {
	bytes, err := templatesModelGotplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/model.gotpl", size: 15811, mode: os.FileMode(420), modTime: time.Unix(1515713855, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesRelationships_registryGotpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x94\x41\x8b\xdb\x30\x14\x84\xef\xfa\x15\x83\x59\x4a\x02\x59\xbb\xe7\x40\x0e\x65\x0b\xcb\x1e\xb6\x94\x2c\x3d\x85\x1c\x14\xf9\xc5\x16\x2b\x4b\xae\xfc\x9c\x12\x84\xfe\x7b\x89\x9d\xec\xc6\xad\x4b\x4a\xb3\xb4\x3d\xea\x59\xef\x9b\x19\x46\xb8\x96\xea\x59\x16\x84\x10\x90\x3e\x11\xa7\x77\xce\x6e\x75\xd1\x7a\xc9\xda\xd9\xf4\x93\xac\x08\x31\x0a\xa1\xab\xda\x79\x46\x52\x68\x2e\xdb\x4d\xaa\x5c\x95\xc9\xda\x79\x62\x77\xab\xad\xca\xc8\x50\x45\x96\xa5\x49\x84\x50\xce\x36\x0c\xeb\x72\xa7\x9e\xd8\x6b\x5b\x60\x81\x64\xd5\x9d\xd7\x09\xb2\x0c\xd6\x19\x6d\x79\x8e\x9d\xf4\xaa\x24\xf5\x3c\xcb\x49\xe6\xca\xe5\x24\x84\xd8\x49\x0f\x4f\xa6\x53\x6f\x4a\x5d\x37\x4b\x2a\x74\xc3\x7e\x8f\x17\x89\x74\x39\xf6\x5d\x88\x2c\xc3\xe0\x0b\x3c\x71\xeb\x6d\x03\x2e\x09\x95\xcb\xc9\x0c\xc9\xa9\xd8\xb6\x56\x0d\x77\x26\xd3\x4b\x3a\x08\x42\xe0\x88\x1e\x77\x2a\xa2\xe8\xc9\xda\x6a\x9e\x4c\x4f\x0b\x63\xac\xc5\x25\xb5\x10\x0f\xcb\x21\xc0\x4b\x5b\x10\x6e\xc8\xb2\xe6\xfd\xa1\x95\x19\x6e\x4e\x4c\xcc\x17\x7d\x77\xc3\xf8\x31\xfe\x4a\x77\x15\xc2\x19\x29\xc6\x87\xbc\x3f\xad\xb1\xc0\xbb\x71\x47\x41\x00\x07\x23\xb7\xd0\x5b\x14\x8c\x89\x21\xfb\xea\x20\xfd\x60\x8c\xfb\xd6\xdc\x79\x92\x4c\x53\xbc\xef\xb5\x81\xf3\xf1\x1c\x95\xac\x57\x4d\xf7\x22\xd6\x1b\xe7\x0c\x7a\x66\x4f\x3d\xe6\xab\xa5\x27\xcb\x87\x40\xaf\xec\x7b\x62\x24\xaa\x63\x24\x27\x30\x90\x84\xf0\x72\x3d\xc6\x04\x73\xb0\x6f\x69\x76\x86\x24\x9b\x9f\xae\xc7\x99\x18\x99\x5e\x48\xf3\xa5\xce\xc7\xd2\xf4\xe3\xeb\xd2\xb4\x1d\xe3\x0d\xd2\xe8\x2d\xe8\xeb\x19\xfd\xd1\xe5\x84\xa4\xa2\x6a\x43\x3e\x19\xfa\xfe\x2c\x59\x95\x7f\x6c\xbb\x63\xbc\xa1\xf1\x1f\x6a\xf8\xfd\x56\x3e\x92\xa1\x91\x56\xfa\xf1\x75\xad\xe4\x1d\xe3\xaf\xbe\xb1\x7b\xe2\x9f\xa2\x2c\x89\xbd\xa6\xdd\x95\x61\x0a\xe2\x2b\x93\x0c\xdd\x3c\x4a\xbb\xff\x3f\x1c\x3d\xd8\xad\xfb\xd7\x4e\x06\xd3\xd8\xff\xa2\x8f\xe7\x28\xbe\x07\x00\x00\xff\xff\x1b\x1a\x93\x85\x53\x07\x00\x00")

func templatesRelationships_registryGotplBytes() ([]byte, error) {
	return bindataRead(
		_templatesRelationships_registryGotpl,
		"templates/relationships_registry.gotpl",
	)
}

func templatesRelationships_registryGotpl() (*asset, error) {
	bytes, err := templatesRelationships_registryGotplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/relationships_registry.gotpl", size: 1875, mode: os.FileMode(420), modTime: time.Unix(1515701556, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"templates/README.md": templatesReadmeMd,
	"templates/identities_registry.gotpl": templatesIdentities_registryGotpl,
	"templates/model.gotpl": templatesModelGotpl,
	"templates/relationships_registry.gotpl": templatesRelationships_registryGotpl,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"templates": &bintree{nil, map[string]*bintree{
		"README.md": &bintree{templatesReadmeMd, map[string]*bintree{}},
		"identities_registry.gotpl": &bintree{templatesIdentities_registryGotpl, map[string]*bintree{}},
		"model.gotpl": &bintree{templatesModelGotpl, map[string]*bintree{}},
		"relationships_registry.gotpl": &bintree{templatesRelationships_registryGotpl, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

