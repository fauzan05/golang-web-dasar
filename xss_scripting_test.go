package golangwebdasar

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateAutoEscape(writer http.ResponseWriter, request *http.Request) {
	// agar tag htmlnya tidak di escape, maka gunakan method template.HTML, template.CSS, atau template.JS
	myTemplates.ExecuteTemplate(writer, "post.gohtml", map[string]interface{}{
		"Title": "Contoh template auto escape",
		"Body": template.HTML("<p>Ini adalah isi dari body</p>"),
	})
}

func TestTemplateXss(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, Localhost, nil)
	recorder := httptest.NewRecorder()

	TemplateAutoEscape(recorder, request)

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