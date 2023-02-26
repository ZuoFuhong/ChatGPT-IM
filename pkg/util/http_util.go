package util

import (
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

// HttpClient http 请求客户端
type HttpClient struct {
	once   *sync.Once
	client *http.Client
}

// DefaultClient 默认的 HTTP 请求客户端
var DefaultClient = &HttpClient{
	new(sync.Once),
	new(http.Client),
}

func init() {
	DefaultClient.create()
}

// DoReq 基于默认HTTP客户端发起HTTP请求
func (c *HttpClient) DoReq(method, url string, reqBody io.Reader, header map[string]string) (
	rspBody []byte, err error) {

	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		return
	}
	// 默认基于JSON请求类型，如果是其他类型，可以通过header参数重置
	req.Header.Set("Content-Type", "application/json;charset=utf8")
	for k, v := range header {
		req.Header.Set(k, v)
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	rspBody, err = io.ReadAll(resp.Body)
	log.Printf("DoReq rspheader:%v, rspbody:%s\n", resp.Header, string(rspBody))
	return
}

// create 创建一个默认的HTTP客户端
func (c *HttpClient) create() *http.Client {
	c.once.Do(func() {
		c.client = &http.Client{
			Transport: &http.Transport{
				Proxy:                 http.ProxyFromEnvironment,
				MaxIdleConns:          100,
				IdleConnTimeout:       90 * time.Second,
				TLSHandshakeTimeout:   5 * time.Second,
				ExpectContinueTimeout: 1 * time.Second,
				MaxIdleConnsPerHost:   100,
			},
			Timeout: 10 * time.Second,
		}
	})
	return c.client
}
