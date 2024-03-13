package golangwebdasar

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ResponseCode(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")
	if name == "" {
		writer.WriteHeader(400)
		fmt.Fprint(writer, "name is empty")
	} else {
		writer.WriteHeader(201)
		fmt.Fprintf(writer, "Halo %s", name)
	}
}

func TestResponseCode(t *testing.T) {
	request := httptest.NewRequest("GET", Localhost + "/home?name=Fauzan", nil)
	recorder := httptest.NewRecorder()

	ResponseCode(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
	fmt.Println("Status code : ", response.StatusCode)
	fmt.Println("Status : ", response.Status)
}