package middleware

import (
	"io/fs"
	"net/http"
)

func StaticFiles(contentFS fs.FS) http.Handler {
	return http.StripPrefix("/static/",
		http.FileServer(http.FS(contentFS)),
	)
}
