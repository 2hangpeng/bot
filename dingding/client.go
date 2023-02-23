package dingding

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const (
	dingTalkHost    = "oapi.dingtalk.com"
	dingTalkPath    = "robot/send"
	defaultTimeout  = time.Second * 3
	contentTypeJSON = "application/json"
	charsetUTF8     = "UTF-8"
)

type Client struct {
	AccessToken string
	Secret      string
	Timeout     time.Duration
}

func NewClient(accessToken string, secret string) Client {
	return Client{
		AccessToken: accessToken,
		Secret:      secret,
		Timeout:     defaultTimeout,
	}
}

type Response struct {
	ErrMsg  string `json:"errmsg"`
	ErrCode int64  `json:"errcode"`
}

func (c *Client) Send(message Message) (*Response, error) {
	if len(c.AccessToken) == 0 {
		return nil, fmt.Errorf("accessToken is empty")
	}
	if message == nil {
		return nil, fmt.Errorf("message is nil")
	}
	reqBytes, err := message.ToBytes()
	if err != nil {
		return nil, fmt.Errorf("message to bytes fail, err: %s, message: %+v", err.Error(), message)
	}
	u := url.URL{
		Scheme: "https",
		Host:   dingTalkHost,
		Path:   dingTalkPath,
	}
	q := u.Query()
	q.Set("access_token", c.AccessToken)
	if len(c.Secret) > 0 {
		timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)
		q.Set("timestamp", timestamp)
		sign, err := c.sign(timestamp)
		if err != nil {
			return nil, fmt.Errorf("get sign fail, err: %s, message: %+v", err.Error(), message)
		}
		q.Set("sign", sign)
	}
	u.RawQuery = q.Encode()
	webHook := u.String()
	req, err := http.NewRequest(http.MethodPost, webHook, bytes.NewReader(reqBytes))
	if err != nil {
		return nil, fmt.Errorf("create request fail, webHook: %s, message: %+v", webHook, message)
	}
	req.Header.Add("Accept-Charset", charsetUTF8)
	req.Header.Add("Content-Type", contentTypeJSON)
	client := &http.Client{
		Transport:     nil,
		CheckRedirect: nil,
		Jar:           nil,
		Timeout:       c.Timeout,
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request fail, err: %s", err.Error())
	}
	defer resp.Body.Close()
	resultBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response body fail, err: %s", err.Error())
	}
	var response Response
	err = jsoniter.Unmarshal(resultBytes, &response)
	if err != nil {
		return nil, fmt.Errorf("unmarshal response body fail, body: %s", string(resultBytes))
	}
	return nil, nil
}

func (c *Client) sign(timestamp string) (string, error) {
	stringToSign := fmt.Sprintf("%s\n%s", timestamp, c.Secret)
	h := hmac.New(sha256.New, []byte(c.Secret))
	if _, err := io.WriteString(h, stringToSign); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(h.Sum(nil)), nil
}
