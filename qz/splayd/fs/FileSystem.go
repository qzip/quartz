package fs

import (
	"io/fs"
	"net/http"
)

/*
implements
  SplayFileInfo   io.fs.FileInfo
  SplayFileSystem   net.http.FileSystem
  File    net.http.File
*/

// SplayFileSystem implements https://pkg.go.dev/net/http#FileSystem
type SplayFileSystem struct {
	Store     SplayedStore
	ClassName string
}

func (sfs *SplayFileSystem) Open(name string) (file http.File, err error) {

	return
}

// File implements net.http.File
type File struct {
}

// Close implements https://pkg.go.dev/io#Closer
func (f *File) Close() error {

	return nil
}

// Read implements https://pkg.go.dev/io#Reader
func (f *File) Read(p []byte) (n int, err error) {

	return
}

// Seek implements https://pkg.go.dev/io#Seeker
func (f *File) Seek(offset int64, whence int) (int64, error) {

	return 0, nil
}

func (f *File) Readdir(count int) (fi []fs.FileInfo, err error) {

	return
}
func (f *File) Stat() (fi fs.FileInfo, err error) {

	return
}
