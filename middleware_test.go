package golangwebdasar

import (
	"fmt"
	"net/http"
	"testing"
)

type LogMiddleware struct {
	Handler http.Handler
}

func (middleware *LogMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Sebelum handler dieksekusi")
	middleware.Handler.ServeHTTP(writer, request)
	fmt.Println("Sesudah handler dieksekusi")
}

type ErrorHandler struct {
	Handler http.Handler	
}

func (errorHandler *ErrorHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	defer func ()  {
		err := recover()
		if err != nil {
			fmt.Println("Terjadi Error")
			writer.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(writer, "Error : %s", err)
		}
	}()
	errorHandler.Handler.ServeHTTP(writer, request)
}

func TestMiddleware(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func (writer http.ResponseWriter, request *http.Request)  {
		fmt.Fprint(writer, "Hello Middleware")
	})
	mux.HandleFunc("/panic", func (writer http.ResponseWriter, request *http.Request)  {
		// fmt.Fprint(writer, "Panic")
		panic("UPS ERROR")
	})

	logMiddleware := new(LogMiddleware)
	logMiddleware.Handler = mux

	errorHandler := new(ErrorHandler)
	errorHandler.Handler = logMiddleware

	server := http.Server{
		Addr: "localhost:8080",
		Handler: errorHandler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}