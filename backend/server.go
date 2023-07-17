package backend

import (
	"fmt"
	"io"
	"io/fs"
	"mime"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/zerolog/log"
)

// Server handles routing for the app
type Server struct {
	Assets fs.FS
}

// Serve serves app in specified port.
func (s *Server) Serve(port int) error {
	// Create router
	router := httprouter.New()
	router.GET("/*filepath", s.serveAssets)

	// Catch unhandled error
	router.PanicHandler = func(w http.ResponseWriter, r *http.Request, arg interface{}) {
		http.Error(w, fmt.Sprintf("unrecovered error: %v", arg), 500)
	}

	// Create server
	url := fmt.Sprintf(":%d", port)
	svr := &http.Server{
		Addr:         url,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: time.Minute,
	}

	// Serve app
	log.Printf("serve app in %s", url)
	return svr.ListenAndServe()
}

func (s *Server) serveAssets(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// If error ever occured, return HTTP error
	var err error
	defer func() {
		if err != nil {
			http.Error(w, err.Error(), 500)
		}
	}()

	// Get request header
	reqEtag := r.Header.Get("If-None-Match")
	reqLastModified := r.Header.Get("If-Modified-Since")

	// Open file
	filePath := ps.ByName("filepath")
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
