package go_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateAutoEscape(w http.ResponseWriter, r *http.Request) {
	err := myTemplate.ExecuteTemplate(w, "post.gohtml", map[string]interface{}{
		"Name":  "ibra",
		"Title": "test",
		"Body":  "<p>Test body auto escape</p>",
	})
	if err != nil {
		return
	}
}

func TemplateAutoEscapeDisabled(w http.ResponseWriter, r *http.Request) {
	err := myTemplate.ExecuteTemplate(w, "post.gohtml", map[string]interface{}{
		"Name":  "ibra",
		"Title": "test",
		"Body":  template.HTML("<p>Test body auto escape</p>"),
	})
	if err != nil {
		return
	}
}

func TemplateXSS(w http.ResponseWriter, r *http.Request) {
	err := myTemplate.ExecuteTemplate(w, "post.gohtml", map[string]interface{}{
		"Name":  "ibra",
		"Title": "test",
		"Body":  template.HTML(r.URL.Query().Get("body")),
	})
	if err != nil {
		return
	}
}

func TestTemplateAutoEscape(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateAutoEscape(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))
}

func TestTemplateAutoDisabled(*testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateAutoEscapeDisabled(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))
}

func TestTemplateXSS(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/?body=<p>alert</p>", nil)
	recorder := httptest.NewRecorder()

	TemplateXSS(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))
}

func TestTemplateAutoEscapeServer(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(TemplateAutoEscapeDisabled),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestTemplateXSSServer(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(TemplateXSS),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
