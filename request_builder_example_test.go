package httplib_test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/cmstar/go-httplib"
)

func ExampleRequestBuilder() {
	// Create a test server.
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("===request===")
		fmt.Println(r.RequestURI)

		for k, v := range r.Header {
			if strings.HasPrefix(k, "X-Custom") {
				fmt.Printf("%s: %s\n", k, strings.Join(v, ", "))
			}
		}

		body, _ := io.ReadAll(r.Body)
		fmt.Println(string(body))

		w.Write([]byte("It works!"))
	}))
	defer ts.Close()

	// Setup RequestBuilder. It supports fluent APIs.
	b := httplib.NewBuilder("POST", ts.URL).
		WithQueries(map[string]any{
			"q2": "v2",
			"q3": 3,
		}).
		WithHeader("x-custom-value", 112233).
		WithForm("f1", "vf1").
		WithForm("f2", "vf2")

	// Do request.
	resp, err := b.ReadString()

	fmt.Println("===response===")
	fmt.Println(resp, err)

	// Output:
	// ===request===
	// /?q2=v2&q3=3
	// X-Custom-Value: 112233
	// f1=vf1&f2=vf2
	// ===response===
	// It works! <nil>
}
