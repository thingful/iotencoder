// Code generated by go-bindata. DO NOT EDIT.
// sources:
// sql/20180525115614_create_device_table.down.sql
// sql/20180525115614_create_device_table.up.sql
// sql/20180526232618_add_streams_table.down.sql
// sql/20180526232618_add_streams_table.up.sql
package migrations

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

var __20180525115614_create_device_tableDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\xcb\xb1\x0e\xc2\x20\x10\x00\xd0\xbd\x5f\x71\xff\xd1\x09\xcb\x99\x90\x18\xda\x08\x03\x4e\x0c\x40\xcc\x2d\xde\xa5\x47\x8d\xfe\xbd\xab\x5a\xe3\xfe\x9e\x3d\xcf\x0b\x38\x6f\x31\x81\x3b\x02\x26\x17\x62\x80\xda\xee\x54\x9a\xe6\xce\x42\x25\x53\x7d\xc0\x64\xc2\x64\x2c\x8e\xc3\xf0\x37\x6c\xda\xd6\xbc\x51\xfd\x75\xa2\x39\x9c\x70\x7f\x76\xec\xb2\x7c\x28\x52\x61\xa5\x4e\x7c\xfb\x96\x98\x22\xfa\xe0\x66\xff\xc6\xe5\x5a\xd6\xa7\x74\x1e\x5f\x01\x00\x00\xff\xff\x13\x76\x2d\x32\xd8\x00\x00\x00")

func _20180525115614_create_device_tableDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__20180525115614_create_device_tableDownSql,
		"20180525115614_create_device_table.down.sql",
	)
}

func _20180525115614_create_device_tableDownSql() (*asset, error) {
	bytes, err := _20180525115614_create_device_tableDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "20180525115614_create_device_table.down.sql", size: 216, mode: os.FileMode(420), modTime: time.Unix(1527783892, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __20180525115614_create_device_tableUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x90\xc1\x6e\xf2\x30\x10\x84\xef\x79\x8a\xb9\x01\x12\x6f\xc0\xc9\xc0\xa2\xdf\xfa\x13\x27\x4d\x1c\x91\xf4\x12\xa5\xb1\x85\x2c\x2a\x1c\x39\x0e\x2a\x6f\x5f\x11\x15\xd5\x14\x54\xf5\xb6\x3b\x3b\x33\x5a\x7d\x9b\x9c\x98\x24\x50\x25\x49\x14\x3c\x15\xe0\x3b\x88\x54\x82\x2a\x5e\xc8\x02\xfd\xa1\x73\x97\xde\xdb\x55\x14\x7d\x39\x65\x9d\x11\x94\x19\x7a\x3b\x18\x6f\xec\x09\xac\x00\x89\x32\xc1\x7c\x66\x4e\xca\x5a\x37\x5b\x62\x66\x47\x3f\x8d\x8b\x20\xc7\xd6\x31\xfd\x68\x57\xfa\x6c\x3a\x3d\x60\x1e\x01\x46\xa1\xa0\x9c\xb3\x18\x59\xce\x13\x96\xd7\xf8\x4f\xf5\x32\x02\xde\x9c\x3d\x6a\x07\x49\x95\x9c\xb2\xa2\x8c\xe3\xab\xee\x6d\x6f\xba\x47\xb9\x77\xe6\xdc\x7a\xdd\x1c\xf5\x05\xeb\x5a\x12\xbb\xbb\x8e\x83\x76\xcd\x68\xd4\x63\xee\xdd\x9e\x0e\xc6\x8f\x4a\x63\x9b\x96\xd7\x57\xb3\x9c\x36\x7c\x42\x72\x67\x6b\xfd\x1f\x5c\x21\x9f\x70\x0e\x3d\x9d\xd3\xad\xd7\xaa\x69\x3d\x24\x4f\xa8\x90\x2c\xc9\xb0\xe7\xf2\xdf\xb4\xe2\x35\x15\x84\x2d\xed\x58\x19\x5f\x1f\xdd\xcf\x17\x51\x40\xb3\x14\xfc\xa5\x24\x70\xb1\xa5\xea\x39\xd4\x66\xe2\xd3\x18\xf5\x11\x01\xa9\xf8\x66\x3d\xe9\x41\xd5\x6f\x1d\x37\x5c\x4f\x6a\x6e\xa7\xc5\xea\x33\x00\x00\xff\xff\xe5\xac\x4c\x5c\x43\x02\x00\x00")

func _20180525115614_create_device_tableUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__20180525115614_create_device_tableUpSql,
		"20180525115614_create_device_table.up.sql",
	)
}

func _20180525115614_create_device_tableUpSql() (*asset, error) {
	bytes, err := _20180525115614_create_device_tableUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "20180525115614_create_device_table.up.sql", size: 579, mode: os.FileMode(420), modTime: time.Unix(1527783892, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __20180526232618_add_streams_tableDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\xf0\x74\x53\x70\x8d\xf0\x0c\x0e\x09\x56\x28\x2e\x29\x4a\x4d\xcc\x2d\x56\x70\x76\x0c\x76\x76\x74\x71\xb5\x06\x04\x00\x00\xff\xff\xa6\x34\x43\x4f\x25\x00\x00\x00")

func _20180526232618_add_streams_tableDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__20180526232618_add_streams_tableDownSql,
		"20180526232618_add_streams_table.down.sql",
	)
}

func _20180526232618_add_streams_tableDownSql() (*asset, error) {
	bytes, err := _20180526232618_add_streams_tableDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "20180526232618_add_streams_table.down.sql", size: 37, mode: os.FileMode(420), modTime: time.Unix(1527783892, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __20180526232618_add_streams_tableUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8f\xb1\x4e\xc3\x30\x14\x45\x77\x7f\xc5\x1d\x13\xa9\x7f\xd0\xc9\xb4\x37\x60\x91\x38\xc5\x7e\x51\x53\x16\x2b\xc4\x1e\x2c\x40\x42\x4d\x40\xf0\xf7\x28\x55\x09\x0c\x8c\x96\xcf\x79\x3a\x77\xe7\xa8\x85\x10\x7d\x53\x13\xa6\x82\x6d\x05\xec\x8d\x17\x8f\x69\x3e\xa7\xe1\x75\x42\xa1\x80\x1c\xe1\xe9\x8c\xae\x71\x70\xa6\xd1\xee\x84\x7b\x9e\x36\x0a\x88\xe9\x23\x8f\x29\xe4\x08\x63\x85\xb7\x74\x97\x0b\xb6\xab\x6b\x38\x56\x74\xb4\x3b\xfa\x2b\x35\x15\x39\x96\x8b\xf4\xf6\xfe\xf4\x92\xc7\xf0\x9c\xbe\x20\xec\x65\x55\x96\xbf\xf1\x9c\x86\x39\xc5\x30\xcc\x10\xd3\xd0\x8b\x6e\x0e\x38\x1a\xb9\xbb\x3c\xf1\xd8\x5a\x62\xcf\x4a\x77\xf5\xe2\x1d\x8b\x52\x95\x5b\xa5\xae\x33\x3a\x6b\x1e\x3a\xc2\xd8\x3d\xfb\xff\xd7\x84\x35\x38\xfc\x56\x84\x1c\x3f\x15\xd0\xda\x1f\xaa\x58\xa9\xcd\x9f\xd8\x72\xab\xbe\x03\x00\x00\xff\xff\x3c\x31\x63\xe0\x2f\x01\x00\x00")

func _20180526232618_add_streams_tableUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__20180526232618_add_streams_tableUpSql,
		"20180526232618_add_streams_table.up.sql",
	)
}

func _20180526232618_add_streams_tableUpSql() (*asset, error) {
	bytes, err := _20180526232618_add_streams_tableUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "20180526232618_add_streams_table.up.sql", size: 303, mode: os.FileMode(420), modTime: time.Unix(1528299124, 0)}
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
	"20180525115614_create_device_table.down.sql": _20180525115614_create_device_tableDownSql,
	"20180525115614_create_device_table.up.sql": _20180525115614_create_device_tableUpSql,
	"20180526232618_add_streams_table.down.sql": _20180526232618_add_streams_tableDownSql,
	"20180526232618_add_streams_table.up.sql": _20180526232618_add_streams_tableUpSql,
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
	"20180525115614_create_device_table.down.sql": &bintree{_20180525115614_create_device_tableDownSql, map[string]*bintree{}},
	"20180525115614_create_device_table.up.sql": &bintree{_20180525115614_create_device_tableUpSql, map[string]*bintree{}},
	"20180526232618_add_streams_table.down.sql": &bintree{_20180526232618_add_streams_tableDownSql, map[string]*bintree{}},
	"20180526232618_add_streams_table.up.sql": &bintree{_20180526232618_add_streams_tableUpSql, map[string]*bintree{}},
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

