package golangwebdasar

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)


func SimpleHtml(writer http.ResponseWriter, request *http.Request) {
	templateText := `<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Document 1</title>
		<link rel="stylesheet" href="style.css">
	</head>
	<body>
		{{.}}
	</body>
	</html>`
	t := template.Must(template.New("Simple").Parse(templateText))
	t.ExecuteTemplate(writer, "Simple", "Hai Fauzan")
}

func TestSimpleHtml(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, Localhost, nil)
	recorder := httptest.NewRecorder()

	SimpleHtml(recorder, request)

	// jika ingin menampilkan di terminal
	// body, _ := io.ReadAll(recorder.Result().Body)
	// fmt.Println(string(body))

	// jika ingin menampilkan ke web
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(SimpleHtml),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func SimpleHTMLFile(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/simple.gohtml"))
	t.ExecuteTemplate(writer, "simple.gohtml", "Hello HTML File")
}

func TestSimpleHtmlFile(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, Localhost, nil)
	recorder := httptest.NewRecorder()

	SimpleHTMLFile(recorder, request)

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

func TemplateDirectory(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseGlob("./templates/*.gohtml"))
	t.ExecuteTemplate(writer, "simple.gohtml", "Hello HTML File with ParseGLOB")
}

func TestSimpleHtmlTemplateDirectory(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, Localhost, nil)
	recorder := httptest.NewRecorder()

	TemplateDirectory(recorder, request)

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

//go:embed templates/*.gohtml
var Templates embed.FS

func TemplateEmbed(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFS(Templates, "templates/*.gohtml"))
	t.ExecuteTemplate(writer, "simple.gohtml", "Hello HTML template with embed")
}

func TestSimpleHtmlTemplateWithEmbed(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, Localhost, nil)
	recorder := httptest.NewRecorder()

	TemplateEmbed(recorder, request)

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

func TemplateDataInterface(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/profile.gohtml"))

	t.ExecuteTemplate(writer, "profile.gohtml", map[string]interface{}{
		"Title": "Ini adalah profile saya",
		"Name": "Fauzan Nurhidayat",
		"Address": map[string]interface{}{
			"Country": "Indonesia",
		},
	})
}


func TestSimpleHtmlTemplateDataInterface(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, Localhost, nil)
	recorder := httptest.NewRecorder()

	TemplateDataInterface(recorder, request)

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

type Address struct {
	Country string
}

type Profile struct {
	Title string
	Name string
	Address
}

func TemplateDataStruct(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/profile.gohtml"))

	t.ExecuteTemplate(writer, "profile.gohtml", Profile{
		Title: "Ini adalah profile saya menggunakan struct",
		Name: "Fauzan Nurhidayat",
		Address: Address{
			Country: "Indonesia",
		},
	})
}

func TestSimpleHtmlTemplateDataStruct(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, Localhost, nil)
	recorder := httptest.NewRecorder()

	TemplateDataStruct(recorder, request)

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

func TemplateDataStructIfElse(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/ifelse.gohtml"))

	t.ExecuteTemplate(writer, "ifelse.gohtml", Profile{
		Title: "Ini adalah profile saya menggunakan struct",
		Name: "Fauzan Nurhidayat",
		// Address: Address{
		// 	Country: "Indonesia",
		// },
	})
}

func TestSimpleHtmlTemplateDataStructIfElse(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, Localhost, nil)
	recorder := httptest.NewRecorder()

	TemplateDataStructIfElse(recorder, request)

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

type Hobbies struct{
	Title string
	Name []string
}

func TemplateDataStructRange(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/range.gohtml"))

	t.ExecuteTemplate(writer, "range.gohtml", Hobbies{
		Title: "Daftar Hobi",
		Name: []string{
			"makan", "minum", "tidur",
		},
	})
}

func TestSimpleHtmlTemplateDataStructRange(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, Localhost, nil)
	recorder := httptest.NewRecorder()

	TemplateDataStructRange(recorder, request)

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

func TemplateLayout(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles(
		"./templates/header.gohtml",
		"./templates/footer.gohtml",
		"./templates/layout.gohtml",
		))

	t.ExecuteTemplate(writer, "layout", map[string]interface{}{
		"Name": "Fauzan",
		"Title": "Ini adalah template layout",
	})
}

func TestTemplateLayout(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, Localhost, nil)
	recorder := httptest.NewRecorder()

	TemplateLayout(recorder, request)

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

// mengakses function di template

func (profile Profile) SayHello(name string) string {
	return "Hello " + name + ", My name is " + profile.Name
}

func TemplateFunction(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.New("FUNCTION").Parse(`{{ .SayHello "Guest" }}`))
	t.ExecuteTemplate(writer, "FUNCTION", Profile{
		Name: "Fauzan",
	})
}

func TestTemplateFunction(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, Localhost, nil)
	recorder := httptest.NewRecorder()

	TemplateFunction(recorder, request)

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

// menambahkan function sendiri ke global function
func TemplateFunctionMap(writer http.ResponseWriter, request *http.Request) {
	t := template.New("add-function")
	// t = t.Funcs(map[string]interface{}{
	// 	"upper": func (value string) string  {
	// 		return strings.ToUpper(value)
	// 	},
	// })
	// t = template.Must(t.Parse(`{{ upper .Name }}`))
	// jika ingin function pipelines 
	t = t.Funcs(map[string]interface{}{
		"sayHello": func (name string) string {
			return "Hello " + name
		},
		"upper": func (value string) string  {
			return strings.ToUpper(value)
		},
	})
	// hasil dari function sayHello() akan dikirim ke function upper()
	t = template.Must(t.Parse(`{{ sayHello .Name | upper}}`))
	t.ExecuteTemplate(writer, "add-function", Profile{
		Name: "Fauzan Nurhidayat NEWw",
	})
}

func TestTemplateGlobalFunction(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, Localhost, nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionMap(recorder, request)

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