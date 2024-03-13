package golangwebdasar

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

const Localhost string = "http://localhost:8080"

func RequestHeader(writer http.ResponseWriter, request *http.Request) {
	contentType := request.Header.Get("content-type")
	fmt.Fprint(writer, contentType)
}

func ResponseHeader(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("X-Powered-By", "Fauzan-Nurhidayat")
	fmt.Fprint(writer, "OK")
}

func TestRequestHeader(t *testing.T) {
	request := httptest.NewRequest("GET", Localhost + "/home", nil)
	request.Header.Add("content-type", "makanan") // menambahkan data header pada request ke server
	recorder := httptest.NewRecorder()

	RequestHeader(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

func TestResponseHeader(t *testing.T) {
	request := httptest.NewRequest("GET", Localhost + "/home", nil)
	recorder := httptest.NewRecorder()

	ResponseHeader(recorder, request)

	poweredBy := recorder.Header().Get("X-Powered-By") // mengambil data header yang dikirim oleh server
	fmt.Println(poweredBy)
}