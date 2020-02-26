package token

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"fmt"
	"futong_server_agent_go/utils/conf"
	"net/http"
	"strings"
)

var DefaultCredential = DefaultNewCredentials()

type Credentials struct {
	AccessKey string
	SecretKey []byte
}

func NewCredentials(accessKey, secretKey string) *Credentials {
	return &Credentials{accessKey, []byte(secretKey)}
}

func DefaultNewCredentials() *Credentials {
	//accessKey := beego.AppConfig.DefaultString("token::accessKey", "futong")
	//secretKey := beego.AppConfig.DefaultString("token::secretKey", "futongcloud")
	accessKey := conf.Config.MustValue("auth", "accessKey")
	secretKey := conf.Config.MustValue("auth", "secretKey")

	return NewCredentials(accessKey, secretKey)
}

// SignRequest 对request相关数据进行签名
func (ath *Credentials) SignRequest(req *http.Request) (token string, err error) {
	data, err := collectData(req)
	if err != nil {
		return
	}
	token = ath.Sign(data)
	return
}

// Sign 对数据进行签名
func (ath *Credentials) Sign(data []byte) (token string) {
	h := hmac.New(sha1.New, ath.SecretKey)
	h.Write(data)

	sign := base64.URLEncoding.EncodeToString(h.Sum(nil))
	return fmt.Sprintf("%s:%s", ath.AccessKey, sign)
}

// Verify 验证request凭证
func (ath *Credentials) Verify(req *http.Request) error {
	auth := req.Header.Get("Authorization")
	if auth == "" {
		return errors.New("token.Verify: token is empty")
	}

	tokenList := strings.Split(auth, " ")
	if len(tokenList) != 2 || strings.ToLower(tokenList[0]) != "bearer" {
		return errors.New("token.Verify: token format error")
	}

	token, err := ath.SignRequest(req)
	if err != nil {
		return err
	}

	if auth != "Bearer "+token {
		return errors.New("token.Verify: token is invalid")
	}

	return nil
}
