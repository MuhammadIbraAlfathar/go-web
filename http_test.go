package go_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func HelloHandle(write http.ResponseWriter, request *http.Request) {
	fmt.Fprint(write, "Hello World")
}

func TestHttp(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	HelloHandle(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	stringBody := string(body)
	fmt.Println(stringBody)
}
