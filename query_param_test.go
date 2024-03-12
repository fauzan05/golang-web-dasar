package golangwebdasar

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)


func SayHelloGuest(w http.ResponseWriter, r *http.Request) {
	first_name := r.URL.Query().Get("first_name")
	last_name := r.URL.Query().Get("last_name")
	if first_name == "" && last_name == "" {
		fmt.Fprint(w, "Hello")
	} else {
		fmt.Fprintf(w, "Hello %s %s", first_name, last_name)
	}
}

func TestQueryParameter(t *testing.T) {
	host := "http://localhost:8080"
	request := httptest.NewRequest(http.MethodGet, host + "/?first_name=Fauzan&last_name=Nurhidayat", nil)
	recorder := httptest.NewRecorder()

	SayHelloGuest(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)

	fmt.Println(bodyString)
}

func MultipleParameterValues(writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	names := query["name"]
	fmt.Fprintln(writer, strings.Join(names, " | "))
}

func TestMultipleParameterValues(t *testing.T) {
	host := "http://localhost:8080"
	request := httptest.NewRequest(http.MethodGet, host + "/?name=Fauzan&name=Susi&name=Rudi", nil)
	recorder := httptest.NewRecorder()

	MultipleParameterValues(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)

	fmt.Println(bodyString)
}