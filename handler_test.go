package golangwebdasar

import (
	"fmt"
	"net/http"
	"testing"
)


func TestServerWithHandler(t *testing.T) {
	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World Cuy")
	}

	server := http.Server{
		Addr: "localhost:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestServeMux(t *testing.T) {
	mux := http.NewServeMux()
	// path yang paling panjang akan dieksekusi terlebih dahulu
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "ini adalah halaman index")
	})
	mux.HandleFunc("/home/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "ini adalah halaman home")
	})
	mux.HandleFunc("/home/sweet/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "ini adalah halaman home/sweet")
	})
	mux.HandleFunc("/store/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "ini adalah halaman store")
	})

	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestRequest(t *testing.T) {
	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, r.Method)
		fmt.Fprint(w, r.RequestURI)
	}
	
	server := http.Server{
		Addr: "localhost:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}