package golangwebdasar

import (
	"embed"
	"io/fs"
	"net/http"
	"testing"
)


func TestFileServer(t *testing.T) {
	directory := http.Dir("./resources")
	fileServer := http.FileServer(directory)

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

//go:embed resources
var resources embed.FS
func TestFileServerWithGoEmbed(t *testing.T) {
	directory, _ := fs.Sub(resources, "resources") // masuk ke dalam folder lagi agar ketika mengakses file index.html, di URLnya tidak static/resources/index.html
	// if err != nil {
	// 	panic(err)
	// }
	fileServer := http.FileServer(http.FS(directory))

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}