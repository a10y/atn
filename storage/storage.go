// Storage system for atn, where we keep track of uploaded messages as well as public
// keys.
package storage

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path"
)

// MessageType is an enumeration of the types of messages we expect to store in the
// system.
type MessageType uint8

const (
	MessageText MessageType = iota
	MessageKey
)

// Storage subsystem.
type Storage struct {
	root string // Root directory on the filesystem where atn stores data
}

// Default root of the atn storage system
// TODO: figure out how to make this a constant again. Probably need to use a function.
func DefaultPath() string {
	return path.Join(os.Getenv("HOME"), ".atn")
}

// Directories that hold the keys and the messages
const keysDir string = "keys"
const textDir string = "txt"

func ensurePath(pathname string) error {
	info, err := os.Stat(pathname)
	if err == nil {
		if !info.IsDir() {
			return fmt.Errorf("Storage root %v must be directory")
		}
	} else if os.IsNotExist(err) {
		if err = os.MkdirAll(pathname, 0700); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("os.Stat(%v) died with error %v", pathname, err)
	}
	return nil
}

// Create new Storage system with default root directory.
func NewDefault() *Storage {
	return New(DefaultPath())
}

// Creates a new Storage system rooted at the specified path.
func New(root string) *Storage {
	return &Storage{
		root: root,
	}
}

// Initialize the storage system. Creates root directory and subdirectories if not
// already present. Returns any error encounted by the filesystem.
func (s *Storage) Init() error {
	if err := ensurePath(s.root); err != nil {
		return err
	}

	// Create subdirectories
	if err := ensurePath(path.Join(s.root, keysDir)); err != nil {
		return err
	}
	if err := ensurePath(path.Join(s.root, textDir)); err != nil {
		return err
	}

	return nil
}

// Get the string representation of the SHA256 digest of data.
func digestBytes(data []byte) string {
	sum := sha256.Sum256(data)
	// Return hex string representing the sum
	return hex.EncodeToString(sum[:])
}

// Add a new message to our repository.
// TODO: verify message contents to ensure it hasn't been tampered with, using the public
// 	 keys.
func (s *Storage) AddMessage(data []byte) (string, error) {
	// Construct a content addressable scheme from the data
	digest := digestBytes(data)

	dir1 := digest[:2]
	dir2 := digest[2:4]
	dirPath := path.Join(s.root, textDir, dir1, dir2)

	if err := ensurePath(dirPath); err != nil {
		return "", err
	}

	// Write the file from the data
	filePath := path.Join(dirPath, digest)
	file, err := os.Create(filePath)
	if err != nil {
		return "", err
	}

	// Write out to file
	if _, err := io.Copy(file, bytes.NewReader(data)); err != nil {
		return "", err
	}

	return digest, nil
}
