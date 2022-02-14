package defs

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"io/ioutil"
	"net/http"
)

type Request struct {
	Method    string
	URL       string
	Body      []byte
	ApiKey    string
	Timestamp string
	Signature string
}

func (r *Request) Validate() error {
	if r.Method == "" {
		return errors.New("method")
	}
	if r.URL == "" {
		return errors.New("url")
	}
	if r.ApiKey == "" {
		return errors.New("api key")
	}
	if r.Timestamp == "" {
		return errors.New("timestamp")
	}

	return nil
}

func (r *Request) Sign(secret []byte) {
	mac := hmac.New(sha256.New, secret)
	mac.Write([]byte(r.Timestamp))
	mac.Write([]byte(r.Method))
	mac.Write([]byte(r.URL))
	mac.Write(r.Body)
	sum := mac.Sum(nil)
	r.Signature = base64.RawURLEncoding.EncodeToString(sum)
}

func (r *Request) Send() (string, error) {

	body := bytes.NewReader(r.Body)

	httpReq, err := http.NewRequest(r.Method, r.URL, body)
	if err != nil {
		return "", err
	}

	httpReq.Header.Set("qredo-api-key", r.ApiKey)
	httpReq.Header.Set("qredo-api-ts", r.Timestamp)
	httpReq.Header.Set("qredo-api-sig", r.Signature)

	cl := http.Client{}
	resp, err := cl.Do(httpReq)
	if err != nil {
		return "", err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(data), nil
}
