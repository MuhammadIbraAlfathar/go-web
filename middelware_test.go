package go_web

import (
	"fmt"
	"net/http"
	"testing"
)

type LogMiddleware struct {
	Handle http.Handler
}

func (middleware *LogMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Before execute handler")
	middleware.Handle.ServeHTTP(writer, request)
	fmt.Println("After execute handler")
}

func TestMiddleware(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Execute Handler")
		fmt.Fprint(writer, "hello middleware")
	})

	logMiddleware := LogMiddleware{
		Handle: mux,
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: &logMiddleware,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
