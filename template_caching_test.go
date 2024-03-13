package golangwebdasar

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Templates diambil dari embed yang ada di file template_test.go
var myTemplates = template.Must(template.ParseFS(Templates, "templates/*.gohtml"))

func TemplateCaching(writer http.ResponseWriter, request *http.Request) {
	myTemplates.ExecuteTemplate(writer, "simple.gohtml", "Hello HTML template with Caching")
}

func TestTemplateCaching(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, Localhost, nil)
	recorder := httptest.NewRecorder()

	TemplateCaching(recorder, request)

	// jika ingin menampilkan di terminal
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))

	// jika ingin menampilkan ke web
	// server := http.Server{
	// 	Addr:    "localhost:8080",
	// 	Handler: http.HandlerFunc(SimpleHtml),
	// }

	// err := server.ListenAndServe()
	// if err != nil {
	// 	panic(err)
	// }
}