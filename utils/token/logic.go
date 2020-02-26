package token

import (
	"bytes"
	"futong_server_agent_go/utils"
	"io"
	"io/ioutil"
	"net/http"
)

func collectData(req *http.Request) (data []byte, err error) {
	u := req.URL
	s := u.Path
	if u.RawQuery != "" {
		s += "?"
		s += u.RawQuery
	}
	s += "\n"

	data = []byte(s)
	if incBody(req) {
		s2, rErr := bytesFromRequest(req)
		if rErr != nil {
			err = rErr
			return
		}
		req.Body = ioutil.NopCloser(bytes.NewReader(s2))
		data = append(data, s2...)
	}
	return
}

func incBody(req *http.Request) bool {
	return req.Body != nil && req.Header.Get("Content-Type") == utils.CONTENT_TYPE_JSON
}

// BytesFromRequest 读取http.Request.Body的内容到slice中
func bytesFromRequest(r *http.Request) (b []byte, err error) {
	if r.ContentLength == 0 {
		return
	}
	if r.ContentLength > 0 {
		b = make([]byte, int(r.ContentLength))
		_, err = io.ReadFull(r.Body, b)
		return
	}
	return ioutil.ReadAll(r.Body)
}
