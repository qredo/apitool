package webui

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/qredo/apitool/defs"

	"github.com/pkg/errors"

	"github.com/gin-gonic/gin"
)

type webSignRequest struct {
	Method string `json:"method"`
	URL    string `json:"url"`
	Body   []byte `json:"body"`
	ApiKey string `json:"api_key"`
	Secret string `json:"secret"`
}

type WebSignResponse struct {
	APIKey    string `json:"api_key"`
	Timestamp string `json:"timestamp"`
	Signature string `json:"signature"`
	Response  string `json:"response"`
}

func Sign(c *gin.Context) {
	req, err := processRequest(c)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}

	c.JSON(200, &WebSignResponse{
		APIKey:    req.ApiKey,
		Timestamp: req.Timestamp,
		Signature: req.Signature,
	})
}

func Send(c *gin.Context) {
	req, err := processRequest(c)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}

	resp, err := req.Send()

	var prettyJSON bytes.Buffer
	err = json.Indent(&prettyJSON, []byte(resp), "", "\t")
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}

	c.JSON(200, &WebSignResponse{
		APIKey:    req.ApiKey,
		Timestamp: req.Timestamp,
		Signature: req.Signature,
		Response:  prettyJSON.String(),
	})
}

func processRequest(c *gin.Context) (*defs.Request, error) {
	wsRequest := &webSignRequest{}
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return nil, errors.New("body")
	}
	err = json.Unmarshal(body, wsRequest)
	if err != nil {
		return nil, err
	}

	req := &defs.Request{
		Timestamp: fmt.Sprintf("%v", time.Now().Unix()),
		Method:    wsRequest.Method,
		URL:       strings.TrimSpace(wsRequest.URL),
		ApiKey:    strings.TrimSpace(wsRequest.ApiKey),
		Body:      wsRequest.Body,
	}

	if err = req.Validate(); err != nil {
		return nil, errors.Wrap(err, "validation")
	}

	secret, err := base64.URLEncoding.DecodeString(strings.TrimSpace(wsRequest.Secret))
	if err != nil {
		return nil, errors.Wrap(err, "secret")
	}

	req.Sign(secret)

	return req, nil
}
