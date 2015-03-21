package static

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/zenazn/goji/web"
)

var faviconPath = "favicon.ico"

func Static(directory string) func(c *web.C, h http.Handler) http.Handler {
	return func(c *web.C, h http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			if r.Method == "GET" {
				originalPath := r.URL.Path[1:]

				if strings.HasPrefix(originalPath, directory) || originalPath == faviconPath {
					if path, err := filepath.Abs(originalPath); err == nil {
						if fi, err := os.Stat(path); err == nil && path != "" {
							if !fi.IsDir() {
								w.Header().Add("Cache-Control", "public, max-age=31536000")
								http.ServeFile(w, r, path)
								return
							}
						}
					}
				}
			}

			h.ServeHTTP(w, r)
			return
		}
		return http.HandlerFunc(fn)
	}
}
