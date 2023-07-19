package backend

import (
	"fmt"
	"io"
	"mime"
	"net/http"
	"path/filepath"
	"strconv"
)

func (s *Server) serveAssets(w http.ResponseWriter, r *http.Request) {
	// If error ever occured, return HTTP error
	var err error
	defer func() { markHttpError(w, err) }()

	// Get request header
	reqEtag := r.Header.Get("If-None-Match")
	reqLastModified := r.Header.Get("If-Modified-Since")

	// Open file
	filePath := r.URL.Path
	if filePath == "/" {
		filePath = "index.html"
	}

	src, err := s.Assets.Open(filePath)
	if err != nil {
		return
	}
	defer src.Close()

	// Get file statistic
	info, err := src.Stat()
	if err != nil {
		return
	}

	// Check if file is modified
	etag := fmt.Sprintf(`W/"%x-%x"`, info.ModTime().Unix(), info.Size())
	lastModified := info.ModTime().UTC().Format("Mon, 02 Jan 2006 15:04:05 GMT")
	if reqEtag == etag || reqLastModified == lastModified {
		w.WriteHeader(http.StatusNotModified)
		return
	}

	// Get content type
	ext := filepath.Ext(filePath)
	mimeType := mime.TypeByExtension(ext)

	// Write response header
	w.Header().Set("ETag", etag)
	w.Header().Set("Last-Modified", lastModified)
	w.Header().Set("Content-Length", strconv.FormatInt(info.Size(), 10))

	if mimeType != "" {
		w.Header().Set("Content-Type", mimeType)
		w.Header().Set("X-Content-Type-Options", "nosniff")
	}

	// Serve file
	io.Copy(w, src)
}
