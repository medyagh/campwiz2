package ramerica

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"path/filepath"
	"time"

	"github.com/mitchellh/go-homedir"
	"github.com/peterbourgon/diskv"
)

var searchPageExpiry = time.Duration(6*3600) * time.Second

// store is a gkvlite Store
var store = setupKVStore()

// Returns a gkvlite collection
func setupKVStore() *diskv.Diskv {
	log.Printf("Opening cache store")
	flatTransform := func(s string) []string { return []string{} }

	h, err := homedir.Dir()
	if err != nil {
		log.Fatalf("failed to get home dir")
	}
	// Initialize a new diskv store, rooted at "homedir/campwiz2", with a 50MB cache.
	d := diskv.New(diskv.Options{
		BasePath:     filepath.Join(h, "campwiz2"),
		Transform:    flatTransform,
		CacheSizeMax: 50 * 1024 * 1024, // 50 MB
	})

	return d
}

func md5sum(s string) (string, error) {
	h := md5.New()
	_, err := io.WriteString(h, s)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
