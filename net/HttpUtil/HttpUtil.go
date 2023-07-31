package HttpUtil

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io"
	"net/http"
)

type Method string

const (
	GET     Method = "GET"
	POST    Method = "POST"
	PUT     Method = "PUT"
	DELETE  Method = "DELETE"
	PATCH   Method = "PATCH"
	OPTIONS Method = "OPTIONS"
)

// 构建请求
type Response struct {
	StatusCode int
	Header     map[string]string
	Body       string
}
type httpBuild struct {
	method Method
	url    string
	header map[string]string
	param  map[string]string
	body   map[string]interface{}
}

func NewClient() *httpBuild {
	return &httpBuild{}
}

func (h *httpBuild) Url(url string) *httpBuild {
	h.url = url
	return h
}

func (h *httpBuild) Method(method Method) *httpBuild {
	h.method = method
	return h
}

func (h *httpBuild) Param(param map[string]string) *httpBuild {
	h.param = param
	return h
}

func (h *httpBuild) Header(header map[string]string) *httpBuild {
	h.header = header
	return h
}

func (h *httpBuild) Body(body map[string]interface{}) *httpBuild {
	h.body = body
	return h
}

func (h *httpBuild) Send() (*Response, error) {
	return send(h.method, h.url, h.header, h.body, h.param)
}

func send(method Method, url string, header map[string]string, body map[string]interface{}, param map[string]string) (*Response, error) {
	byteB, err := json.Marshal(body)
	req, err := &http.Request{}, nil
	if body == nil {
		req, err = http.NewRequest(string(method), url, nil)
	} else {
		req, err = http.NewRequest(string(method), url, bytes.NewBuffer(byteB))
	}

	// 跳过https ssl
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: transport}

	// 设置请求参数
	if param != nil {
		q := req.URL.Query()
		for k, v := range param {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}
	if err != nil {
		return nil, err
	}
	// 设置请求头
	if header != nil {
		for k, v := range header {
			req.Header.Set(k, v)
		}
	}
	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	// 读取响应
	bodyR, errR := io.ReadAll(resp.Body)
	if errR != nil {
		return nil, errR
	}
	headers := make(map[string]string)
	for key, values := range resp.Header {
		headers[key] = values[0]
	}
	return &Response{
		StatusCode: resp.StatusCode,
		Header:     headers,
		Body:       string(bodyR),
	}, nil
}
