package go_web

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func UploadForm(w http.ResponseWriter, r *http.Request) {
	err := myTemplate.ExecuteTemplate(w, "form.gohtml", nil)
	if err != nil {
		return
	}
}

func Upload(w http.ResponseWriter, r *http.Request) {
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		panic(err)
	}

	fileDestination, err := os.Create("./resources/" + fileHeader.Filename)

	if err != nil {
		panic(err)
	}

	_, err = io.Copy(fileDestination, file)

	if err != nil {
		panic(err)
	}

	name := r.PostFormValue("name")

	err = myTemplate.ExecuteTemplate(w, "uploadsuccess.gohtml", map[string]interface{}{
		"Name": name,
		"File": "/static/" + fileHeader.Filename,
	})
	if err != nil {
		panic(err)
	}
}

func TestUploadForm(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", UploadForm)
	mux.HandleFunc("/upload", Upload)
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./resources"))))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

//go:embed resources/test.jpeg
var uploadFileTest []byte

func TestUploadFile(t *testing.T) {

	body := new(bytes.Buffer)

	writer := multipart.NewWriter(body)
	file, _ := writer.CreateFormFile("file", "contoh.png")
	_, _ = file.Write(uploadFileTest)
	_ = writer.Close()

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080", body)
	request.Header.Set("Content-Type", writer.FormDataContentType())
	recorder := httptest.NewRecorder()

	Upload(recorder, request)

	bodyResponse, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(bodyResponse))
}
