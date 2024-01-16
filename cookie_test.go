package go_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func AddCookie(w http.ResponseWriter, r *http.Request) {
	cookie := new(http.Cookie)
	cookie.Name = "test-cookie"
	cookie.Value = r.URL.Query().Get("name")
	cookie.Path = "/"

	http.SetCookie(w, cookie)
	fmt.Fprint(w, "Success add cookie")
}

func GetCookie(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("test-cookie")
	if err != nil {
		fmt.Fprint(w, "Cookie not found")
	} else {
		name := cookie.Value
		fmt.Fprintf(w, "Hello %s", name)
	}
}

func TestCookie(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/get-cookie", GetCookie)
	mux.HandleFunc("/set-cookie", AddCookie)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestAddCookie(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/?name=ibra&alfathar", nil)
	record := httptest.NewRecorder()

	AddCookie(record, request)

	cookies := record.Result().Cookies()

	for _, cookie := range cookies {
		fmt.Printf("Cookie %s : %s \n", cookie.Name, cookie.Value)
	}
}

func TestGetCookie(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	record := httptest.NewRecorder()

	cookie := new(http.Cookie)
	cookie.Name = "test-cookie"
	cookie.Value = "ibra"
	request.AddCookie(cookie)

	GetCookie(record, request)

	body, _ := io.ReadAll(record.Result().Body)
	fmt.Println(string(body))

}
