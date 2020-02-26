package hub

import (
	"futong_server_agent_go/utils/conf"
	"futong_server_agent_go/utils/token"
	"io/ioutil"
	"net/http"
	"strings"
)

func signRequest(r *http.Request) (string, error) {
	return token.DefaultCredential.SignRequest(r)
}

func request(method, url string, body []byte) ([]byte, error) {
	req, _ := http.NewRequest(method,url,strings.NewReader(string(body)))

	tokenStr, err := signRequest(req)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer " + tokenStr)

	userId := conf.Config.MustValue("auth", "userId")
	req.Header.Set("UserId", userId)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func HttpGet(url string) ([]byte, error) {
	return request(http.MethodGet, url, nil)
}

func HttpPost(url string,  body []byte) ([]byte, error) {
	return request(http.MethodPost, url, body)
}
