package go_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type MyPage struct {
	Name string
}

type User struct {
	Username string
}

func (myPage MyPage) SayHello(name string) string {
	return "Hello " + name + ", My Name is " + myPage.Name
}
func TemplateFunction(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.New("FUNCTION").Parse(`{{.SayHello "Budi"}}`))

	err := t.ExecuteTemplate(w, "FUNCTION", MyPage{
		Name: "Ibra",
	})
	if err != nil {
		panic(err)
	}
}

func TemplateFunctionGlobal(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.New("FUNCTION_GLOBE").Parse(`{{len .Username}}`))

	err := t.ExecuteTemplate(writer, "FUNCTION_GLOBE", User{
		Username: "Joko Susilo",
	})

	if err != nil {
		panic(err)
	}
}

func TemplateFunctionCreateGlobal(w http.ResponseWriter, r *http.Request) {
	t := template.New("FUNCTION")

	t = t.Funcs(map[string]interface{}{
		"upper": func(value string) string {
			return strings.ToUpper(value)
		},
	})

	t = template.Must(t.Parse(`{{ upper .Name }}'`))

	err := t.ExecuteTemplate(w, "FUNCTION", MyPage{
		Name: "Ibra",
	})
	if err != nil {
		panic(err)
	}

}

func TemplateFunctionPipeline(w http.ResponseWriter, r *http.Request) {
	t := template.New("PIPELINE")

	t = t.Funcs(map[string]interface{}{
		"sayHello": func(value string) string {
			return "Hello" + value
		},
		"upper": func(value string) string {
			return strings.ToUpper(value)
		},
	})

	t = template.Must(t.Parse(`{{ sayHello .Name | upper }}`))
	err := t.ExecuteTemplate(w, "PIPELINE", MyPage{
		Name: "Ibra",
	})
	if err != nil {
		panic(err)
	}
}

func TestTemplateFunctionGlobal(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionGlobal(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))
}

func TestTemplateFunction(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunction(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))
}

func TestTemplateFunctionCreateGlobal(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionCreateGlobal(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))
}

func TestTemplateFunctionPipeline(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionPipeline(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))
}
