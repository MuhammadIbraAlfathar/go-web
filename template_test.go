package go_web

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SimpleHtml(w http.ResponseWriter, r *http.Request) {
	templateText := `<html><body>{{.}}</body></html>`

	t, err := template.New("SIMPLE").Parse(templateText)
	if err != nil {
		panic(err)
	}

	err = t.ExecuteTemplate(w, "SIMPLE", "Hello World")
	if err != nil {
		panic(err)
	}
}

func SimpleHtmlFile(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./template/test.gohtml")
	if err != nil {
		panic(err)
	}

	err = t.ExecuteTemplate(w, "test.gohtml", "Go HTML")
	if err != nil {
		panic(err)
	}
}

func TemplateDirectory(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseGlob("./template/*.gohtml")
	if err != nil {
		panic(err)
	}

	err = t.ExecuteTemplate(w, "test.gohtml", "Go HTML")
	if err != nil {
		panic(err)
	}
}

//go:embed template/*.gohtml
var templates embed.FS

func TemplateEmbed(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFS(templates, "template/*.gohtml")
	if err != nil {
		panic(err)
	}

	err = t.ExecuteTemplate(w, "test.gohtml", "Go HTML")
	if err != nil {
		panic(err)
	}
}

func TestSimpleHtml(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	SimpleHtml(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))
}

func TestSimpleHtmlFile(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	SimpleHtmlFile(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))
}

func TestTemplateDirectory(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateDirectory(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))
}

func TestTemplateEmbed(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateEmbed(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))
}
