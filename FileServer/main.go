package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RealIP)

	mediaDir := "/home/milind/Media"
	if _, err := os.Stat(mediaDir); os.IsNotExist(err) {
		log.Fatalf("Directory %s does not exist", mediaDir)
	}

	// Define allowed directories
	allowedDirs := []string{
		"Movies",
		"TV",
		"Books",
		"Martin Garrix Sets",
		"Music",
	}

	fileServer := http.FileServer(http.Dir(mediaDir))

	r.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		reqPath := chi.URLParam(r, "*")

		// Check if the request is for an allowed directory
		allowed := false
		for _, dir := range allowedDirs {
			if strings.HasPrefix(reqPath, dir+"/") || reqPath == dir {
				allowed = true
				break
			}
		}

		if !allowed {
			http.Error(w, "Access denied", http.StatusForbidden)
			return
		}

		fullPath := filepath.Join(mediaDir, reqPath)
		if _, err := os.Stat(fullPath); os.IsNotExist(err) {
			http.Error(w, "File not found", http.StatusNotFound)
			return
		}

		r.URL.Path = reqPath
		fileServer.ServeHTTP(w, r)
	})

	// Add root handler to show available directories
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write([]byte("<h1>Available Media Directories</h1><ul>"))
		for _, dir := range allowedDirs {
			w.Write([]byte(fmt.Sprintf(`<li><a href="/%s">%s</a></li>`, dir, dir)))
		}
		w.Write([]byte("</ul>"))
	})

	port := ":8080"
	fmt.Printf("Starting file server on port %s serving directories: %v from %s\n", 
		port, allowedDirs, mediaDir)
	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatal(err)
	}
}
