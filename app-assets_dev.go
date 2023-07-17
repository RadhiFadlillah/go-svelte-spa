//go:build dev

package main

import (
	"io/fs"
	"net/http"
)

type devAssets struct {
	http.Dir
}

func (a *devAssets) Open(name string) (fs.File, error) {
	return a.Dir.Open(name)
}

var assets = &devAssets{http.Dir("build/public-dev")}
