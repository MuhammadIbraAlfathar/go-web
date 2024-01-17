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

//go:embed template
var templatess embed.FS

var myTemplate = template.Must(template.ParseFS(templates, "template/*.gohtml"))

func TemplateCaching(w http.ResponseWriter, r *http.Request) {
	err := myTemplate.ExecuteTemplate(w, "test.gohtml", "Hello World")
	if err != nil {
		panic(err)
	}
}

func TestTemplateCaching(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateCaching(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))
}
