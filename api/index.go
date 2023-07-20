package handler

import (
	"log"
	"net/http"

	"github.com/eduardylopes/image-capture-go/screenshot"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	elementID := r.URL.Query().Get("elementId")

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if url == "" {
		http.Error(w, "Missing 'url' parameter", http.StatusBadRequest)
		return
	}

	if elementID == "" {
		http.Error(w, "Missing 'elementId' parameter", http.StatusBadRequest)
	}

	buf := screenshot.GetScreenshot(url, elementID)

	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Cache-Control", "public, immutable, no-transform, s-maxage=31536000, max-age=31536000")

	if _, err := w.Write(buf); err != nil {
		log.Fatal("Failed to write screenshot to the response")
	}
}
