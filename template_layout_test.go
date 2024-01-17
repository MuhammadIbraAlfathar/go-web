package go_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateLayout(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./template/header.gohtml", "./template/layout.gohtml", "./template/footer.gohtml"))

	err := t.ExecuteTemplate(w, "layout", Page{
		Title: "Sebuah Template data struct",
		Name:  "Muhammad Ibra",
	})
	if err != nil {
		panic(err)
	}
}

func TestTemplateLayout(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateLayout(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))
}
