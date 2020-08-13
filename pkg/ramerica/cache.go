package ramerica

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/peterbourgon/diskv"
)

var searchPageExpiry = time.Duration(6*3600) * time.Second

// store is a gkvlite Store
var store = getCacheStore()

// Returns a gkvlite collection
func getCacheStore() *diskv.Diskv {
	log.Printf("Opening cache store")
	flatTransform := func(s string) []string { return []string{} }

	// Initialize a new diskv store, rooted at "my-data-dir", with a 1MB cache.
	d := diskv.New(diskv.Options{
		BasePath:     "data2",
		Transform:    flatTransform,
		CacheSizeMax: 1024 * 1024,
	})

	return d
}

func md5sum(s string) string {
	h := md5.New()
	io.WriteString(h, s)
	return fmt.Sprintf("%x", h.Sum(nil))
}
