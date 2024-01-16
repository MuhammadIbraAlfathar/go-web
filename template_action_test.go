package go_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateActionIf(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./template/if.gohtml"))

	err := t.ExecuteTemplate(w, "if.gohtml", Page{
		Title: "Sebuah Template data struct",
	})
	if err != nil {
		panic(err)
	}
}

func TemplateActionOperator(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./template/comparator.gohtml"))

	err := t.ExecuteTemplate(w, "comparator.gohtml", map[string]interface{}{
		"Title":      "Template action comparator",
		"FinalValue": 10,
	})
	if err != nil {
		panic(err)
	}
}

func TemplateActionRange(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./template/range.gohtml"))

	err := t.ExecuteTemplate(w, "range.gohtml", map[string]interface{}{
		"Title": "Template action comparator",
		"Hobbies": []string{
			"Game", "Read", "Code",
		},
	})
	if err != nil {
		panic(err)
	}
}

func TemplateActionWith(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./template/with.gohtml"))

	err := t.ExecuteTemplate(w, "with.gohtml", Page{
		Title: "Sebuah Template data struct",
		Address: Address{
			Street: "Jl. Kebagusan",
			City:   "Jakarta Timur",
		},
	})
	if err != nil {
		panic(err)
	}
}

func TestTemplateActionIf(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionIf(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))
}

func TestTemplateActionOperator(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionOperator(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))
}

func TestTemplateActionRange(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionRange(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))
}

func TestTemplateActionWith(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionWith(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))
}
