package test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTaskSendHostMetrics(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter,  r *http.Request) {
		a, _ := ioutil.ReadAll(r.Body)
		fmt.Println(string(a))
	}))
	defer ts.Close()

	fmt.Println("url:", ts.URL)
	//hub.sendHostMetrics(ts.URL)
}
