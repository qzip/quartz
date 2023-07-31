/*
Copyright (c) 2019, QzIP Blockchain Technology LLP

All Rights Reserved

Author: Ashish Banerjee, <ashish@qzip.in>

*/

package fs

import (
	"crypto/sha256"
	"encoding/hex"
	"io/fs"
	"time"
)

// SplayFileHeader header for Splay FS
type SplayFileHeader struct {
	ContentHash   HashData `json:"contentHash"`
	ContentLength int64    `json:"contentLength"`
	ChunkSize     int64    `json:"chunkSize"`
}

// SplayMetadata Passed by the saver
type SplayMetadata struct {
	Path      string    `json:"path"`
	ClassName string    `json:"className"`
	AuthURI   string    `json:"authURI"`
	TmStamp   time.Time `json:"tmStamp"`
	UUID      string    `json:"uuid"`
}

// SplayChunk file fragments
type SplayChunk struct {
	ChunkHash HashData `json:"ChunkHash"`
	Size      int64    `json:"size"`
	Offset    int64    `json:"offset"`
}

// SplayFile splayed File splayed
type SplayFile struct {
	Metadata SplayMetadata   `json:"metadata"`
	Header   SplayFileHeader `json:"header"`
	Chunks   []SplayChunk    `json:"chunks"`
	FileInfo SplayFileInfo   `json:"fileinfo"`
}

type SplayFileInfo struct {
	Mod   fs.FileMode `json:"mode"` //https://pkg.go.dev/io/fs#FileMode
	Nam   string      `json:"name"`
	Siz   int64       `json:"size"`
	ModTm time.Time   `json:"modtime"`
}

// base name of the file
func (sfi *SplayFileInfo) Name() string {
	return sfi.Nam
}

// length in bytes for regular files; system-dependent for others
func (sfi *SplayFileInfo) Size() int64 {
	return sfi.Siz
}

// file mode bits
func (sfi *SplayFileInfo) Mode() fs.FileMode {
	return sfi.Mod
}

// modification time
func (sfi *SplayFileInfo) ModTime() time.Time {
	return sfi.ModTm
}

// abbreviation for Mode().IsDir()
func (sfi *SplayFileInfo) IsDir() bool {
	return sfi.Mod&fs.ModeDir != 0
}

func (sfi *SplayFileInfo) Sys() interface{} {
	return nil
}

// SplayedStore splay storage
type SplayedStore interface {
	ContentAddressedStore // lower level than splay fs
	SaveSplay(sfile *SplayFile) (hash HashData, err error)
	LoadSplayByUUID(uuid string) (file *SplayFile, err error)
	//LoadUUID a path may be associated with mulltiple class names
	// for example, a path may have json as well as binary file associations
	LoadUUID(path, className string) (uuid []string, err error)
}

// ContentAddressedStore chunck storage
type ContentAddressedStore interface {
	LoadCAS(hash HashData) (content []byte, err error)
	SaveCAS(content []byte) (hash HashData, err error)
}

// copied from qzip/blockchain

// HashData []bytearray serialized as Hex encoding
type HashData []byte

// MarshalText custom marshall
func (h HashData) MarshalText() ([]byte, error) {

	return []byte(hex.EncodeToString(h)), nil
}
func (h HashData) String() string {
	return hex.EncodeToString(h)
}

// UnmarshalText custom marshall
func (h *HashData) UnmarshalText(in []byte) (err error) {
	b, err := hex.DecodeString(string(in))
	*h = HashData(b)

	return
}

// NewHashData creates a newHash Data
func NewHashData(in []byte) *HashData {
	hx := HashData(in)
	return &hx
}

func HashContent(content []byte) HashData {
	return HashData(content)
}

// Hash returns SHA-256 hash (tendermint merkle tree uses SHA256)
func Hash(in []byte) []byte {
	hash := sha256.New()
	hash.Write(in)
	return hash.Sum(nil)
}
