package go_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func RequestHeader(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("content-type")
	_, err := fmt.Fprint(w, contentType)
	if err != nil {
		panic(err)
	}
}

func ResponseHeader(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("x-powered-by", "ibra")
	fmt.Fprint(w, "OK")
}

func TestHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/", nil)
	request.Header.Add("content-type", "application/json")
	record := httptest.NewRecorder()

	RequestHeader(record, request)

	response := record.Result()

	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func TestResponseHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	request.Header.Add("content-type", "application/json")
	record := httptest.NewRecorder()

	ResponseHeader(record, request)

	response := record.Result()

	body, _ := io.ReadAll(response.Body)

	responseHeader := record.Header().Get("x-powered-by")

	fmt.Println(string(body))
	fmt.Println(responseHeader)
}
