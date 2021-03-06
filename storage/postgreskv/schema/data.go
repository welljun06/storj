// Code generated by go-bindata. DO NOT EDIT.
// sources:
// 2018092201_initial-tables.down.sql
// 2018092201_initial-tables.up.sql
package schema

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

var __2018092201_initialTablesDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x8f\xc1\x0a\x82\x40\x10\x86\xef\x3e\xc5\x1e\x0d\x7a\x03\x4f\x5a\x5b\x08\xa1\x22\x5b\xe1\x69\x18\xd7\x89\xa4\x5c\x65\x1c\x03\xdf\x3e\x08\x0a\x54\xa8\x4e\x03\x33\xff\x7c\x7c\xff\x36\x4f\x33\xb5\x3b\x26\x1b\x13\xa7\x89\xba\xd7\xbd\x40\x55\x33\x59\x69\x79\x04\xa6\x07\x71\x4f\x7e\x54\x18\x1d\xae\xd5\x74\xc4\x89\xd1\x7b\x9d\xaf\x02\xef\x1b\xe3\xaf\x5f\x53\x64\x5a\x75\x28\x57\x40\x57\x41\x43\x82\x73\xa8\x6d\x9b\xae\x75\xe4\x04\x6a\x67\x99\x1a\x72\xe2\xff\xf0\x28\x47\x21\x9c\xc7\x17\x29\xe1\xc1\x59\x14\x02\xbc\x08\xf1\x8c\xb9\x80\x9f\x62\x7d\x7e\x89\x56\x28\x08\x1d\x93\xc8\xf8\xee\x10\x46\x07\xfd\xb9\x4d\x96\xe5\x60\x6f\x24\x7d\xe0\x3d\x03\x00\x00\xff\xff\x96\xd8\x40\x70\x6d\x01\x00\x00")

func _2018092201_initialTablesDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__2018092201_initialTablesDownSql,
		"2018092201_initial-tables.down.sql",
	)
}

func _2018092201_initialTablesDownSql() (*asset, error) {
	bytes, err := _2018092201_initialTablesDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "2018092201_initial-tables.down.sql", size: 365, mode: os.FileMode(420), modTime: time.Unix(1539025200, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __2018092201_initialTablesUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x58\x6d\x6f\xdb\xba\x15\xfe\xee\x5f\x71\x3e\x14\x90\x85\x2b\x27\xe9\xb0\x02\x5b\x52\x17\x70\x1d\x3a\xd1\xea\xc8\x85\xac\x34\x08\x70\x01\x83\x92\x68\x9b\xab\x4c\x6a\x22\xd5\xc4\xc3\xfd\xf1\x03\xa9\x37\x4a\x96\x9d\xdc\xad\x28\x0a\xac\xfa\x12\x47\x3a\x7c\xce\xe1\x73\x0e\xcf\x0b\xa7\x3e\x9a\x04\x08\x82\xc9\xc7\x39\x82\x30\x8f\xbe\x12\x29\x60\x38\x00\x80\xf2\x3f\x86\x77\x04\x3e\x3e\x06\x68\xa2\x5f\xaa\xe7\xb3\xef\xde\x4d\xfc\x47\xf8\x84\x1e\x1d\xfd\x32\x26\x09\xdd\x81\xeb\x05\xb5\x88\xb7\x08\xc0\xbb\x9f\xcf\xeb\x17\xd3\x5b\x34\xfd\x04\xc3\x42\xf2\x03\x5c\xc0\xc4\xbb\x2e\xd7\xbd\x87\xbf\xbc\x7b\x67\x0f\xec\xab\xc1\x60\x34\x82\x9c\x49\x9a\x80\xdc\x12\xf8\x44\xf6\x5f\x70\x92\x93\xa5\xe4\x19\x01\xca\x24\xc9\xd6\x38\x22\x20\xf2\x34\xe5\x99\x14\x90\x62\x21\x28\xdb\x68\xe1\xc2\x58\x10\x24\xc5\x19\x96\x24\xd9\x3b\xc0\xb3\x06\x4f\x48\x9e\xfd\x13\x70\x24\x73\x9c\x24\xfb\x06\x22\xda\x62\xb6\xa9\x30\xb4\x39\x54\x92\x4c\xbd\xce\x70\xa4\x7e\xa5\x24\x2b\xb1\x1d\x90\x5b\x2a\x14\x64\x9c\xef\x76\x7b\xc8\xf8\x13\x88\x2d\xcf\x93\x18\x44\xbe\x5e\xd3\x88\xc0\x9a\x67\x40\xbe\x91\x6c\x2f\xb7\x94\x6d\xce\x06\xae\xb7\x44\x7e\xa0\x88\x59\x34\xdc\x36\xb4\x3a\x85\x46\x1b\xbe\x4c\xe6\xf7\x68\x09\x43\xcb\xba\xbc\xd4\x54\x3b\x80\x45\x44\xe9\xd0\x3a\xb7\x6c\xc5\xcb\xa0\xe5\xa6\x14\xcb\x6d\x8c\x25\x6e\xf9\xa9\xe3\xa3\x03\x07\xf8\x68\x86\x7c\xe4\x4d\xd1\xb2\xcf\x14\xbb\x70\xe4\x3a\x4f\x12\x85\xfe\x12\x58\xe9\xcd\x5a\xfc\xfd\x07\xb0\xac\x12\x63\x47\x24\xd6\xd6\xf5\x63\x38\x83\x41\x27\x86\x2a\x3b\x9c\x5a\x7d\x11\x0c\xe5\x9e\xbf\xb8\xe8\xa1\xde\xf2\x2a\xcd\x88\x94\x7b\x98\x2c\x35\xca\x12\xcd\xd1\x34\x00\xc2\x22\x1e\x93\x1a\xc6\x22\x22\xc2\x29\xb1\x6c\x98\x54\x9b\x75\x6a\x3b\x00\x2a\xf1\x4a\x5b\x67\x41\xfd\xba\x67\x49\xb5\xb7\xce\x92\xea\x75\xb9\x62\xe6\x2f\xee\x6a\x8b\x95\xf7\x46\x23\xd8\xd0\x6f\x84\x01\xd6\xaf\x01\x0b\xd8\xd1\xcd\x56\x42\xa8\x62\x26\x67\x31\x50\xa6\x23\xb0\xf6\xac\xc4\x61\x42\x1c\x90\x59\xce\x22\x2c\x09\x50\x09\x78\xad\xe2\x51\x49\x31\xf2\x2c\x8d\x60\x95\x1c\x42\xa2\x74\x14\x50\x58\x02\xcf\x4a\x69\x4b\xff\x49\xb9\xb0\x1c\xa0\x6b\xc0\x6c\x7f\x36\x18\x8d\x94\x2c\x7a\xc6\xbb\x34\x21\xe2\xb2\xfc\x5f\x3d\x95\xb6\x95\x5e\x65\x46\xa3\xf1\x34\x81\xe9\xc0\x5b\x1b\x60\xf4\x01\xa0\x96\x3c\x86\xb4\xe6\xfc\x10\xac\x17\xc9\x90\x3c\x01\x76\x7e\x80\x76\x0c\xec\xfc\x35\x68\x21\xce\xce\x43\xfc\xef\xbe\xd3\xe7\xc0\x5f\xbf\x2b\xda\x3b\x13\x4d\x49\xfe\xcf\x88\x7f\xfb\xee\x88\x7f\xef\x22\x1a\x92\xa7\x40\x0b\xd9\xef\x45\xe3\x49\xb4\x16\x8d\x26\x5c\x95\x35\x66\xf7\xde\x34\x70\x17\x5e\x17\x3a\x6c\xd2\x9b\xd3\xd4\x2d\x74\x83\x7c\x07\xaa\xc3\x52\xbd\xb1\xc1\x47\xc1\xbd\xef\x2d\x0b\x79\x75\xd2\xdf\xbc\x19\x5c\xa3\xe9\x7c\xe2\x23\x7d\xd6\x55\xe2\x2f\x3e\x5e\x35\x85\xd0\x40\xb8\x1a\x7c\x44\x37\xae\xd7\xc8\x5e\x8e\x41\xe4\xa1\x90\x19\x65\x9b\xd2\x16\x9d\x2d\x2a\xd5\x76\x07\xe7\x72\x0c\x29\x17\x54\x52\xce\x86\x82\xc8\x55\xb8\x97\x64\x68\x41\x43\xca\x45\x5d\x44\x5c\x4f\xab\x28\x11\xdc\x59\x03\xa2\xca\x6d\x70\x8b\x3c\xa3\x16\xa8\x6d\xf5\x5b\xf2\x16\x66\x0b\x1f\x86\x35\x17\xbf\x35\x38\x23\x78\x6b\x97\xf0\xc8\xbb\x06\x77\x56\xfc\x2e\xd1\x34\xc6\xd5\x00\x79\xd7\x57\x83\x37\x6f\x60\x3e\xf1\x6e\xee\x27\x37\x08\xac\x34\x49\x37\xe2\x5f\x89\x05\xee\xdd\xdd\x7d\x51\xc1\x96\x81\xef\x4e\x03\xa3\xb0\xd5\xee\x52\x1b\xc4\x2b\xca\xa2\x8c\xec\x08\x93\xc3\xb0\xa0\xb7\xdf\x15\x0d\xb9\x0f\xb7\xae\xea\x5f\x8a\x42\x54\x92\xa3\x5b\x8c\x4d\x45\x5a\xe8\x00\x8f\x24\x91\xab\x84\xb0\x8d\xdc\x0e\x43\x5b\xef\x07\xc6\xaa\xfb\x80\xf9\x62\xf1\xb9\xa6\x27\xec\xb8\xc9\x24\xa6\x0f\xa3\xa1\x44\xc1\xd4\xfc\x87\x30\x36\x8c\xe9\x73\x80\xaa\x87\x47\xf9\x14\xa7\x2d\x77\x5e\xde\xda\x6f\xda\xb6\x3f\xe9\x90\xd1\x08\x6e\xfe\xab\x5a\xf5\x44\x95\x74\x11\x2d\x44\x15\x1f\x9c\xa6\x84\xc5\x24\x56\x95\x47\x6e\xb1\xd4\x2b\x14\xfe\x16\x0b\x55\x8a\x14\xc5\x23\x92\x68\x3f\x0b\x07\x32\x22\xf3\x8c\x35\x05\x2e\xe5\x42\xd0\x30\x29\x14\x15\x00\xab\x48\xb5\x5b\x2b\x65\x4a\x69\x84\xd6\x0d\x43\xf1\x95\xa6\xa9\x6a\xe3\xf8\x37\xa2\xbb\x3e\x05\x9f\x72\x49\x98\xa4\x38\x69\x29\xb2\x4f\x55\xc0\x88\xef\x52\xce\x08\x93\x46\x04\x5a\xe7\xbd\x5d\x59\x99\x7f\x2e\xba\xa9\x4c\x3d\x43\xc6\x75\x07\x08\x11\x66\xca\xd8\x90\xc8\x27\x42\x18\x58\xe7\x16\x60\x16\xab\x55\xc0\xe5\x56\x97\x73\xcc\x94\x79\x6a\x8f\x02\x72\x16\xab\xa2\x7d\x6e\xd9\x2f\xda\xd4\x4d\xf6\x7d\xb6\x55\x42\xa6\x8d\xaf\x42\x0d\x39\xdf\x9d\x00\xee\x08\xc2\x1f\x7f\x00\xb2\x7e\xff\xfd\xf9\xe2\x42\x1b\xfe\x0a\x1a\x5a\xcb\x35\x23\xe6\x2b\x8d\xa4\x3c\xcc\x78\xb6\xc3\x49\x79\xe4\x79\x16\x13\x75\x1e\x5f\xe2\x46\x9b\x32\x43\xb3\xd9\x6c\xd6\xec\x41\x4d\x18\x85\xf1\xc5\xe7\xd9\xa9\x92\xd1\x87\x7a\xbc\x6e\xf4\xa7\x26\xa3\x3b\x9d\x4e\x96\x08\x1e\x6e\x91\x67\x9c\xd8\xa2\xef\x6c\x9f\x5a\xdd\xf9\x56\x49\x49\xab\x30\x3b\x50\xfd\xa8\x24\x62\xe0\x15\x76\xe8\x0c\x76\x20\xda\x59\x74\x90\x57\x8b\x3e\xfb\xf8\x2a\x34\x5f\x22\x23\x07\x9d\x34\xb8\x62\x44\x25\x9c\x7e\x48\xe4\x5d\x1f\x7c\xd0\x1a\x0a\x5e\x8d\x08\xaa\xfc\xd2\x5e\xdc\xc9\x60\x2f\x95\x93\xe0\xf1\x73\x31\x26\xad\x30\x8b\x57\xaa\x3d\x57\x6e\x19\xf6\xcc\x38\xbd\x33\x8b\x31\x7d\xd4\x41\x91\x50\x21\x57\x31\xcd\x48\x24\x79\xb6\x1f\x9a\x43\x97\x03\x31\xcd\xcc\xf8\x10\x12\x67\x72\x85\xcb\xcf\x30\x2e\x06\x1f\xd0\x0d\xfb\x4a\xf2\x2a\x72\xca\x0f\xf6\xa0\x0a\xa0\x25\x0a\x16\xb3\x43\xbb\xcb\x70\x7a\x70\x83\x5b\xf0\xd1\xf4\xde\x5f\xba\x5f\x50\x4d\x10\x65\x69\x2e\x45\xb3\xbf\xea\x39\x88\xbe\xca\xc8\x6e\x55\xd2\x46\x14\xde\xa8\x44\x54\x39\x9a\x2c\x21\xc3\x6c\x43\x56\x09\x7f\x72\xfa\xbc\xfa\xa7\x90\xbb\xf1\x57\x2e\xb1\xdb\xaa\xb6\x74\xb3\xed\xd5\xd5\x0a\xbc\x7a\xed\x6f\xf0\x56\xad\x6d\xce\xab\x26\xbe\x17\x20\x3c\x2b\x42\x74\xb2\x2c\x62\xf5\x88\x90\x71\xe7\x51\xcf\x8e\x1d\x49\xdd\x0c\x54\x23\x74\xd8\xfe\xf8\x70\x8b\x7c\x64\x5e\x9c\x8c\xbb\x18\x76\xa3\x38\xa6\x42\x52\x16\x49\x35\xd3\xae\xe9\x33\x0c\xab\x56\x35\x2e\x76\x77\xcc\xa3\xc3\xf2\x6f\xa7\xb3\x4d\xe3\xb3\x66\xa2\xa5\xc5\x76\xd5\x8f\x0e\x3b\xc7\xce\x7c\x6b\x68\x85\x34\xee\x17\x2b\x36\x38\xa4\x67\x75\x68\x80\xbb\x2c\xdc\xbc\xf0\xc1\x30\x01\x3e\x80\x21\x74\x34\xd1\xa8\x16\xad\x46\x53\xde\x3f\x06\xf7\x1e\x4c\xa9\xd3\x78\xf5\xf1\x53\xf1\x78\x60\xd7\xb8\x3e\x9e\x27\x51\xd2\xb8\x0c\x06\x18\x03\x3d\xeb\x8d\x84\xf2\x59\xf8\xd7\xc8\x87\x8f\x8f\xa6\x9a\x7e\xc9\xb9\x7b\xe7\x06\x87\x39\x52\x53\x5f\x1e\x63\xda\xfa\x76\xef\xa9\xd4\x33\x31\x6e\x5e\xe0\x67\x8a\x83\x0e\xaf\x7d\x75\x33\x5d\x3f\x9f\xb5\xe2\xba\xb6\xe9\x87\xc6\xc4\x8f\xf6\x66\xf7\x6c\xa7\xeb\x67\xa7\xdf\xc5\x15\x95\x5d\x9e\xf4\xa6\xbb\xf7\x6e\xb6\xd9\x57\x1c\x2e\x39\x76\x7b\x95\xc6\x67\x7d\x57\x54\x3d\x46\xc2\x1c\xcd\x02\x58\xdc\x07\xc8\x87\x7f\x2c\x5c\xcf\x8c\x04\x58\x78\x3d\x3a\xc7\x87\x44\xbd\x76\x47\xed\xe8\x1e\x8d\xf4\x05\x2b\x70\x46\xf4\x65\x15\x95\x40\x9e\xa9\x50\xa3\x41\x84\x99\xa5\x07\x90\x14\x67\x12\xf8\xfa\x30\x75\xf2\x4c\xc9\x3f\xe9\xdb\xd8\x08\xe7\x82\x40\x2e\x40\x72\x50\x83\x81\x1e\x0a\x00\x27\x49\xa5\x46\xe4\x61\xc4\x99\x1a\x0e\x84\x02\xd3\xc3\x4c\x81\xf3\x44\xac\x8c\x40\xc2\xf9\x57\xd5\xb6\xae\x79\x66\x3b\x20\x38\x3c\xa9\x41\x23\xfa\xaa\x54\x70\x06\x5b\x92\x91\x96\x1f\xcc\x03\x67\x50\x7d\x70\x94\x9a\x08\xe8\x39\x3f\x63\x33\x5f\x0e\xda\xa1\xd8\xa6\xb7\x88\xbb\xaa\x99\xe8\xeb\x8c\x96\xba\x2d\x7a\xb1\x89\x59\x65\xe4\x1b\xc9\x04\xf9\xd5\xcc\xfc\x6a\x66\x7e\x35\x33\x27\x9b\x99\xf1\x4f\xda\xcd\xbc\xff\x21\xdd\x0c\x5c\xa3\xe5\xf4\xff\xa3\xa5\xf9\xce\xd1\xd0\x76\xfa\x41\x49\xfc\xe9\x5c\xf6\xab\x6f\x79\xcd\x8e\x0e\x0a\xb3\xc1\xf6\xeb\xab\xf3\x7f\x02\x00\x00\xff\xff\xa9\xf1\x2d\xb9\x7b\x1f\x00\x00")

func _2018092201_initialTablesUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__2018092201_initialTablesUpSql,
		"2018092201_initial-tables.up.sql",
	)
}

func _2018092201_initialTablesUpSql() (*asset, error) {
	bytes, err := _2018092201_initialTablesUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "2018092201_initial-tables.up.sql", size: 8059, mode: os.FileMode(420), modTime: time.Unix(1539025200, 0)}
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
	"2018092201_initial-tables.down.sql": _2018092201_initialTablesDownSql,
	"2018092201_initial-tables.up.sql":   _2018092201_initialTablesUpSql,
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
	"2018092201_initial-tables.down.sql": &bintree{_2018092201_initialTablesDownSql, map[string]*bintree{}},
	"2018092201_initial-tables.up.sql":   &bintree{_2018092201_initialTablesUpSql, map[string]*bintree{}},
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
