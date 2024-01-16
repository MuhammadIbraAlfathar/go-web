package go_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type Address struct {
	Street string
	City   string
}

type Page struct {
	Title   string
	Name    string
	Address Address
}

func TemplateDataStruct(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./template/name.gohtml"))

	err := t.ExecuteTemplate(w, "name.gohtml", Page{
		Title: "Sebuah Template data struct",
		Name:  "Muhammad Ibra Alfathar",
		Address: Address{
			Street: "Jl. Kenangan 10",
		},
	})
	if err != nil {
		panic(err)
	}
}

func TemplateDataMap(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./template/name.gohtml"))

	data := map[string]interface{}{
		"Title": "Template Data",
		"Name":  "Muhammad Ibra Alfathar",
	}

	err := t.ExecuteTemplate(w, "name.gohtml", data)
	if err != nil {
		panic(err)
	}
}

func TestTemplateDataMap(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateDataMap(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))
}

func TestTemplateDataStruct(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateDataStruct(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))
}
