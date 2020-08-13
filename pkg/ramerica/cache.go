package ramerica

import (
	"log"
	"os"
	"time"

	"github.com/golang/glog"
	"github.com/steveyen/gkvlite"
)

// cachePath is the location of the cache
var cachePath = os.ExpandEnv("${HOME}/.campwiz2.cache")
var searchPageExpiry = time.Duration(6*3600) * time.Second

// store is a gkvlite Store
var store = getCacheStore()
var collection = store.SetCollection("cache", nil)

// Returns a gkvlite collection
func getCacheStore() *gkvlite.Store {
	glog.Infof("Opening cache store: %s", cachePath)
	f, err := os.OpenFile(cachePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		log.Fatal(err)
	}
	s, err := gkvlite.NewStore(f)
	if err != nil {
		log.Fatal(err)
	}
	return s
}
