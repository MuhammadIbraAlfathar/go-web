package go_web

import (
	"fmt"
	"net/http"
	"testing"
)

type LogMiddleware struct {
	Handle http.Handler
}

type ErrorHandler struct {
	Handler http.Handler
}

func (middleware *LogMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Before execute handler")
	middleware.Handle.ServeHTTP(writer, request)
	fmt.Println("After execute handler")
}

func (errorHandler *ErrorHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Terjadi error")
			writer.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(writer, "Error: %s", err)
		}
	}()

	errorHandler.Handler.ServeHTTP(writer, request)
}

func TestMiddleware(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Execute Handler")
		fmt.Fprint(writer, "hello middleware")
	})
	mux.HandleFunc("/panic", func(writer http.ResponseWriter, request *http.Request) {
		panic("test")
	})

	logMiddleware := LogMiddleware{
		Handle: mux,
	}

	errorHandler := &ErrorHandler{
		Handler: &logMiddleware,
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: errorHandler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
