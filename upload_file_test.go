package golangwebdasar

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

func UploadForm(writer http.ResponseWriter, request *http.Request) {
	err := myTemplates.ExecuteTemplate(writer, "upload_form.gohtml", nil)
	if err != nil {
		panic(err)
	}
}

func UploadFile(writer http.ResponseWriter, request *http.Request) {
	request.ParseMultipartForm(1 << 20) // max ukuran total request body 1mb, bukan total ukuran file yang diupload
	file, fileHeader, err := request.FormFile("file")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fileDestination, err := os.Create("./resources/" + fileHeader.Filename)
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(fileDestination, file)
	if err != nil {
		panic(err)
	}
	name := request.PostFormValue("name")
	myTemplates.ExecuteTemplate(writer, "upload_success.gohtml", map[string]interface{}{
		"Name": name,
		"File": "/static/" + fileHeader.Filename,
	})
}

//go:embed resources/samplefoto.PNG
var uploadFileTest []byte

func TestShowUploadForm(t *testing.T) {
	// test via web
	// mux := http.NewServeMux()
	// mux.HandleFunc("/", UploadForm)
	// mux.HandleFunc("/upload", UploadFile)
	// mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./resources"))))

	// server := http.Server{
	// 	Addr:    "localhost:8080",
	// 	Handler: mux,
	// }

	// err := server.ListenAndServe()
	// if err != nil {
	// 	panic(err)
	// }

	// test via unit test
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body) // body kosongan

	writer.WriteField("name", "Ini adalah foto yang saya upload")
	file, err := writer.CreateFormFile("file", "contohupload.png")
	if err != nil {
		panic(err)
	}
	upload, err := file.Write(uploadFileTest)
	if err != nil {
		panic(err)
	}
	fmt.Println("Status upload ", upload)
	writer.Close()

	request := httptest.NewRequest(http.MethodPost, Localhost + "/upload", body)
	request.Header.Set("Content-Type", writer.FormDataContentType()) // multipart/form-data | writer.FormDataContentType()
	recorder := httptest.NewRecorder()

	UploadFile(recorder, request)	

	bodyResponse, err := io.ReadAll(recorder.Result().Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bodyResponse))

}
