// Code generated by go-bindata.
// sources:
// manifest/data/Dockerfile.node
// manifest/data/Dockerfile.rails
// manifest/data/Dockerfile.ruby
// manifest/data/Dockerfile.unknown
// DO NOT EDIT!

package manifest

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path/filepath"
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
	name string
	size int64
	mode os.FileMode
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

var _dataDockerfileNode = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x72\x0b\xf2\xf7\x55\xc8\xcb\x4f\x49\xb5\x32\xd0\x33\x34\xe0\xe2\x72\x8d\x08\xf0\x0f\x76\x55\x30\x36\x30\x30\xe0\x72\xf5\x0b\x53\x08\xf0\x0f\x0a\x81\xf0\xb8\xc2\xfd\x83\xbc\x5d\x3c\x83\x14\xf4\x13\x0b\x0a\xb8\xb8\x9c\xfd\x03\x22\x15\x0a\x12\x93\xb3\x13\xd3\x53\xf5\xb2\x8a\xf3\xf3\xc0\xe2\xfa\xc8\x22\x5c\x41\xa1\x7e\x0a\x79\x05\xb9\x0a\x99\x79\xc5\x25\x89\x39\x39\x50\x4d\x7a\x30\x13\x7c\x5d\x14\xa2\x95\x80\xf2\x4a\x3a\x0a\x4a\x40\x05\x45\x25\x4a\xb1\x5c\x80\x00\x00\x00\xff\xff\x54\x98\x55\x86\x90\x00\x00\x00")

func dataDockerfileNodeBytes() ([]byte, error) {
	return bindataRead(
		_dataDockerfileNode,
		"data/Dockerfile.node",
	)
}

func dataDockerfileNode() (*asset, error) {
	bytes, err := dataDockerfileNodeBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/Dockerfile.node", size: 144, mode: os.FileMode(420), modTime: time.Unix(1437445390, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _dataDockerfileRails = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x54\x8e\xcf\xce\x82\x30\x10\xc4\xef\xfb\x14\x9b\x9e\x3f\x4a\xf3\x79\xf3\x2a\x68\x8c\x81\x92\xfa\x3f\xc6\x43\x91\x6a\x88\x15\x9a\x52\x4c\x78\x7b\xa1\xe2\x81\x4c\xd2\x74\x7e\x93\x9d\xdd\xa5\xe0\x09\xda\x36\xef\xe6\xff\xb4\x17\x80\xd8\xa7\x28\x8d\x0b\x1e\xca\x61\x6b\x0a\xe9\xd4\x04\x05\x1d\x96\x55\xe3\xa4\xd6\x58\xd5\x85\x02\x88\x4f\x19\xdf\xc6\x38\x63\x8c\x41\x9c\x1e\x30\xe3\x62\xf7\x75\x70\xe4\x62\x13\xad\x05\x86\xd2\x18\x80\x05\xcf\xce\xb8\x52\xaf\x7b\xa9\x95\x47\xe1\x68\x26\x09\xd5\xf5\xed\x39\x89\x3d\xf1\x37\xe4\x6d\x55\xf4\xb3\xe3\xfe\xb1\x91\xfe\xea\x93\x08\x2f\xc4\xca\x52\x37\xe4\x0f\x49\xa3\xec\x5b\xd9\xe1\x17\xe4\xc3\xcb\xa8\x17\xb9\xc2\x27\x00\x00\xff\xff\xbe\x5d\x14\xf6\xf2\x00\x00\x00")

func dataDockerfileRailsBytes() ([]byte, error) {
	return bindataRead(
		_dataDockerfileRails,
		"data/Dockerfile.rails",
	)
}

func dataDockerfileRails() (*asset, error) {
	bytes, err := dataDockerfileRailsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/Dockerfile.rails", size: 242, mode: os.FileMode(420), modTime: time.Unix(1437445397, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _dataDockerfileRuby = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x72\x0b\xf2\xf7\x55\x28\x2a\x4d\xaa\xb4\x32\xd2\x03\x42\x2e\x2e\xd7\x88\x00\xff\x60\x57\x05\x63\x03\x03\x03\x2e\x57\xbf\x30\x85\x00\xff\xa0\x10\x08\x8f\x2b\xdc\x3f\xc8\xdb\xc5\x33\x48\x41\x3f\xb1\xa0\x80\x8b\xcb\xd9\x3f\x20\x52\xc1\x3d\x35\x37\x2d\x33\x27\x15\x2c\xa4\x0f\xe5\xa0\xc8\xe8\xe5\xe4\x27\x67\xa3\x48\x83\x45\xb8\x82\x42\xfd\x14\x92\x4a\xf3\x52\x80\x7a\x33\xf3\x8a\x4b\x12\x73\x72\xa0\x26\xea\x41\x8c\x07\x04\x00\x00\xff\xff\xd7\x6f\x55\x75\x98\x00\x00\x00")

func dataDockerfileRubyBytes() ([]byte, error) {
	return bindataRead(
		_dataDockerfileRuby,
		"data/Dockerfile.ruby",
	)
}

func dataDockerfileRuby() (*asset, error) {
	bytes, err := dataDockerfileRubyBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/Dockerfile.ruby", size: 152, mode: os.FileMode(420), modTime: time.Unix(1437175004, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _dataDockerfileUnknown = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x24\xc8\x41\xaa\xc2\x30\x14\x05\xd0\x79\x56\x71\xc9\xb8\xfc\xbf\x88\x06\x41\xa4\x44\x32\x11\x11\x07\x31\xbd\xda\xd0\xf0\x12\xea\x0b\xe2\xee\x15\x9c\x1d\xce\x2e\xf8\x09\xfd\xd6\x45\xbb\x31\x27\x1f\x0e\x6e\x1f\xf0\x1f\x5b\x33\xa3\x3f\x9e\xf1\xf7\xb3\x19\x27\x87\x8b\x65\x5a\xaa\x1d\x60\xbf\x55\x72\x8a\x9a\xab\x40\xdf\x8d\xe8\xb2\x4a\x7d\xc9\x80\x56\x18\x9f\x04\xe7\xac\xd0\x85\x78\x50\xb8\x45\xe5\x0c\x57\xd3\xca\xed\x9e\x0b\xed\xd5\x7c\x02\x00\x00\xff\xff\x1b\x6d\x99\x86\x76\x00\x00\x00")

func dataDockerfileUnknownBytes() ([]byte, error) {
	return bindataRead(
		_dataDockerfileUnknown,
		"data/Dockerfile.unknown",
	)
}

func dataDockerfileUnknown() (*asset, error) {
	bytes, err := dataDockerfileUnknownBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/Dockerfile.unknown", size: 118, mode: os.FileMode(420), modTime: time.Unix(1437445390, 0)}
	a := &asset{bytes: bytes, info:  info}
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
	if (err != nil) {
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
	"data/Dockerfile.node": dataDockerfileNode,
	"data/Dockerfile.rails": dataDockerfileRails,
	"data/Dockerfile.ruby": dataDockerfileRuby,
	"data/Dockerfile.unknown": dataDockerfileUnknown,
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
	Func func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"data": &bintree{nil, map[string]*bintree{
		"Dockerfile.node": &bintree{dataDockerfileNode, map[string]*bintree{
		}},
		"Dockerfile.rails": &bintree{dataDockerfileRails, map[string]*bintree{
		}},
		"Dockerfile.ruby": &bintree{dataDockerfileRuby, map[string]*bintree{
		}},
		"Dockerfile.unknown": &bintree{dataDockerfileUnknown, map[string]*bintree{
		}},
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
