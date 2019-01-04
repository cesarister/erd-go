// Code generated by go-bindata.
// sources:
// templates/dot.tmpl
// templates/dot_relations.tmpl
// templates/dot_tables.tmpl
// DO NOT EDIT!

package main

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

var _templatesDotTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x92\x5d\xcb\xd3\x30\x14\x80\xef\xfb\x2b\x42\xae\xbb\x98\x55\x87\x8a\xcb\xc0\x0b\x85\x81\x6e\xa2\xbb\xf2\x03\x49\x9b\xb3\x36\x9a\x25\x25\x39\x63\x60\xe8\x7f\x97\xb6\x1b\x4d\xe7\xf0\xed\x45\x39\x3c\xe7\x23\x4f\x4f\x13\xe3\x82\x28\x38\x6a\x0b\x84\x2a\x87\x94\x2c\xba\x2e\xab\xbd\x6c\x1b\x12\x33\x42\x08\x19\xe3\x6f\x43\xdc\x3f\x7d\x83\x3e\x12\x76\xd0\x68\x60\x7c\xbf\x45\xf4\xba\x3c\x23\x04\x66\x64\x09\x66\x98\x71\xab\x1f\x88\x58\xaf\xdf\xef\x77\x07\xf2\x69\xbf\xdd\x1d\x16\x5f\xb6\x5f\xdf\x09\x5a\x70\xba\x89\xf1\x7f\x73\xba\x6e\xfd\xac\x6f\xdb\x6c\xf2\xf9\xb8\x5f\xe7\x80\xc2\xdc\x41\xe3\x2a\x81\xf9\xcc\x13\xac\x9a\xb9\x58\xa7\x20\x40\x2b\x38\x5b\x4d\x85\x5e\xda\xdf\xff\xc0\x56\x2a\x41\x39\x2b\x72\xce\x0a\x3a\xe1\x93\xf4\xb5\xb6\x7d\x86\x27\xb4\x02\x5b\x81\x45\x2f\x11\x04\xfa\x33\x4c\x99\xd0\x1a\x6d\x21\x08\x3a\x06\x74\x7e\xaa\xd2\x5e\x7c\xf8\x3c\xa0\x1f\x6f\xb2\x9b\x60\xb2\xeb\x71\x77\xf4\xfb\x2e\x69\x3c\x3a\x8b\x41\xff\x01\xb1\x7c\xf1\xd0\xeb\x65\xce\x19\x5f\x25\x0d\x2d\xd8\x8b\x56\xd8\x88\x25\xe3\x89\x59\x23\x5b\x10\x1f\x3d\x54\xce\xab\x54\x01\x54\x9d\x2a\xf4\x8e\xa5\xc3\xe6\x91\x40\x31\x41\xe9\xbd\xbb\x0c\x94\xb3\xd7\x4f\x9d\x3d\x7c\x96\xb4\xb5\x01\xf1\xbc\xb8\xc3\x4a\x07\x94\xb6\x02\xb1\x64\xaf\x52\xad\x18\x11\x4e\xad\x91\x38\xde\xd3\x9f\x1e\x8c\x44\xed\x6c\xa0\x84\x5d\xff\xf0\x7d\x09\xca\xd2\xc0\x35\xdf\x65\x31\x82\x55\x5d\xf7\x37\x00\x00\xff\xff\xec\x09\x89\xb9\xf0\x02\x00\x00")

func templatesDotTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesDotTmpl,
		"templates/dot.tmpl",
	)
}

func templatesDotTmpl() (*asset, error) {
	bytes, err := templatesDotTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/dot.tmpl", size: 752, mode: os.FileMode(420), modTime: time.Unix(1546592531, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesDot_relationsTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x92\x51\x4b\xfb\x30\x14\xc5\xdf\xfb\x29\x2e\x7d\xfa\xef\xef\x1a\xdd\xb3\xdd\x40\x04\x9f\xa4\xc2\xd8\x9b\x88\xdc\x9a\xdb\x2d\x10\x13\x4c\x23\x43\x42\xbe\xbb\xa4\x76\xb1\x9a\x74\xf8\x76\x39\xdc\xfc\xce\xe5\x9c\x38\xc7\xa9\x13\x8a\xa0\xe4\xda\x3e\x1b\x92\x68\x85\x56\x7d\xe9\x7d\xe1\x9c\x41\xb5\x27\x60\xdb\x93\xea\x7d\x01\xe0\x1c\xbb\xa7\xce\xee\xb0\x95\xd4\xe0\x2b\x79\x0f\x55\x15\xd4\xad\xd8\x1f\x7e\xc8\x8f\x05\x40\xd8\xaf\x40\x74\xf0\x8f\xde\xe0\x6b\xe5\x16\x0d\x17\x0a\xa5\xb0\x1f\x50\xfe\x2f\x17\x50\x0d\x5c\x00\x34\x46\x1f\x0f\x84\x7c\xad\x5f\x8c\x3e\x2e\xc3\x28\xb1\x25\xb9\xae\xeb\xbb\x87\x66\xb7\xb9\x62\xac\xa9\x2f\x87\x71\xb3\x8c\x70\x92\x3d\x9d\x71\xb8\x28\x17\x79\xbe\x25\x4a\x2d\x56\xf3\x16\xe9\x99\x4a\x2b\xd2\x59\xcc\x29\x8e\xc9\x25\xde\x67\xb8\x8a\x47\xec\x98\x53\x4c\xfb\xc6\x5a\x23\xda\x77\x4b\x3d\x1b\xd0\x71\x31\x31\x9a\x79\xf0\x47\xc3\x21\xb6\xd0\xe8\xf9\x5e\x2c\x0a\x39\xf6\x12\xc6\xd9\x5e\xf2\xb5\x24\xfc\x5f\xad\x7c\xd3\x43\x9c\x89\xc1\x6a\xd6\x20\xbd\x31\x96\x92\x50\xc6\x9f\x9b\xed\x24\x93\xd0\xd3\x75\x31\x15\x26\xf3\x67\x00\x00\x00\xff\xff\x93\x21\x06\x17\x35\x03\x00\x00")

func templatesDot_relationsTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesDot_relationsTmpl,
		"templates/dot_relations.tmpl",
	)
}

func templatesDot_relationsTmpl() (*asset, error) {
	bytes, err := templatesDot_relationsTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/dot_relations.tmpl", size: 821, mode: os.FileMode(420), modTime: time.Unix(1546592531, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesDot_tablesTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x93\x5f\x8b\x9b\x40\x14\xc5\xdf\xf3\x29\x2e\x97\xa5\x4f\x89\x66\xdb\xb4\x0f\xdd\x19\xc1\x7f\xd9\x15\xac\x06\x77\x68\xa1\xa5\x14\x8d\x93\x20\x9d\x6a\xd1\xd9\xc2\x32\xf5\xbb\x17\x27\x63\x1a\xd3\xa4\xb0\x79\xc9\xdc\x93\x3b\xf7\x9e\xfc\x8e\x2a\x55\xf2\x5d\x55\x73\xc0\xb2\x91\xdf\x64\x5e\x08\xde\x61\xdf\xcf\x94\x6a\xf3\x7a\xcf\xe1\x46\x7e\x9f\xc3\x8d\x84\xf7\x14\x2c\xa6\x7f\xed\xfb\x19\x80\x52\x16\xab\xa4\xe0\x7d\x0f\x5f\x44\x5e\x70\x41\x09\x61\xae\x17\x87\x33\xd0\x1f\x2f\xcd\x82\x30\xa3\xb8\x44\x23\xf8\x61\x1c\x6f\xdc\x20\x88\x92\xfb\x33\xf5\x71\xe3\xfa\x07\xd5\x7a\x3b\xea\x9f\xa2\x80\x3d\x50\xbc\x7d\xb3\x1a\x15\x37\x8e\xee\x13\x8a\x7e\x98\xb0\x30\x1b\x45\xc7\x7c\x13\x96\x8d\xc7\xa1\x08\xce\xba\xe1\xa3\xa9\xbd\x94\xb1\xf4\x03\x9e\x8e\x77\xc8\x3a\x4d\x18\x6c\xd2\x28\x61\x8b\xc7\xe8\x73\x48\xf1\x76\x85\xb0\x76\xfd\x90\xe2\x03\x17\xbf\xb8\xac\xb6\x39\x14\x8d\x28\xd1\x21\x9e\xf3\xf7\x8f\x13\xdb\x73\x88\x3d\xdc\x76\x88\xcd\x82\xa3\x17\x7b\x34\x43\x6c\x4d\xe4\x50\x28\xb5\x80\x6a\x07\x96\xdf\x88\xa7\x1f\x75\x07\x0b\x8d\x11\xe0\xf7\xa1\xf5\xbf\xec\x8c\xfb\x38\x5c\xb3\x17\xe0\x5c\x5d\x80\x39\x9a\x1c\xdc\x98\x7c\x87\x78\xb7\x3a\x5e\x63\xcd\x18\xbb\x4e\x55\xfb\xb8\xc4\xed\x35\x4e\xf0\x68\x34\xc7\x01\x53\x00\xae\x94\x6d\x55\x3c\x49\xde\x59\xfa\xf1\x39\xf2\x30\xdb\xf4\xf0\x43\x08\x6e\x5b\xe5\x02\x22\x99\x8b\x6a\x8b\xd3\x85\x4b\x04\x3f\x8d\xd3\x8c\xe2\xbe\xe5\xcf\xef\x96\xe8\xbc\xaa\x8b\xee\xe7\x9d\x52\x57\xd6\x5c\xb6\xc5\xeb\x72\xb2\xff\x72\x9c\xc7\x5e\xd3\xf9\x4f\xbc\x66\xcc\x34\x6f\xfd\xce\x9c\xd8\x28\xf6\xdb\x46\x34\xad\x99\x31\xdf\x55\x42\x68\x81\xe2\x80\xee\x5a\x33\xce\x75\x7b\x27\x9f\x05\xa7\xc3\x1d\x5e\x9e\xaf\xd5\xf5\xd7\xbb\xd9\xa9\x74\x7a\xfe\x13\x00\x00\xff\xff\xa8\x98\x28\xae\xe9\x03\x00\x00")

func templatesDot_tablesTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesDot_tablesTmpl,
		"templates/dot_tables.tmpl",
	)
}

func templatesDot_tablesTmpl() (*asset, error) {
	bytes, err := templatesDot_tablesTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/dot_tables.tmpl", size: 1001, mode: os.FileMode(420), modTime: time.Unix(1546592531, 0)}
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
	"templates/dot.tmpl": templatesDotTmpl,
	"templates/dot_relations.tmpl": templatesDot_relationsTmpl,
	"templates/dot_tables.tmpl": templatesDot_tablesTmpl,
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
		"dot.tmpl": &bintree{templatesDotTmpl, map[string]*bintree{}},
		"dot_relations.tmpl": &bintree{templatesDot_relationsTmpl, map[string]*bintree{}},
		"dot_tables.tmpl": &bintree{templatesDot_tablesTmpl, map[string]*bintree{}},
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
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
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

