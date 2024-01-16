package go_web

import (
	_ "embed"
	"fmt"
	"net/http"
	"testing"
)

func ServeFile(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		http.ServeFile(w, r, "./resources/notfound.html")
	} else {
		http.ServeFile(w, r, "./resources/ok.html")
	}
}

//go:embed resources/notfound.html
var resourcesNotFound string

//go:embed resources/ok.html
var resourcesOk string

func ServeFileEmbed(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if name == "" {
		fmt.Fprint(w, resourcesNotFound)
	} else {
		fmt.Fprint(w, resourcesOk)
	}
}

func TestServeFile(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(ServeFile),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestEmbedServeFile(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(ServeFileEmbed),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
