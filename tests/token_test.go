package test

import (
	"fmt"
	"futong_server_agent_go/hub"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestVerifyToken(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		accessKey := "futong"
		secretKey := "futongcloud"
		mac := qbox.NewMac(accessKey, secretKey)


		suc, err := mac.VerifyCallback(r)
		if !suc || err != nil {
			t.Fatal("verify token fail", err)
		}
	}))
	defer ts.Close()

	fmt.Println("url:", ts.URL)
	_, err := hub.HttpPost(ts.URL+"/", nil)
	if err != nil {
		t.Fatal(err)
	}
}
