package tools

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Response 结构体封装 HTTP 响应信息
type Response struct {
	Body       []byte
	StatusCode int
	Header     http.Header
	URL        string
	Duration   time.Duration
	Raw        *http.Response
}

// SendRequest 发送 HTTP 请求（统一入口）
func SendRequest(target, method, bodyData string, headers map[string]string, proxy string) (*Response, error) {
	start := time.Now()

	// 创建 HTTP 客户端
	client := &http.Client{
		Timeout: 10 * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			// 不跟随重定向
			return http.ErrUseLastResponse
		},
	}

	// 设置代理（如果有的话）
	if proxy != "" {
		pu, err := url.Parse(proxy)
		if err != nil {
			return nil, fmt.Errorf("代理URL格式错误: %v", err)
		}
		client.Transport = &http.Transport{
			Proxy: http.ProxyURL(pu),
		}
	}

	// 构建请求体
	var reqBody io.Reader = nil
	if bodyData != "" {
		reqBody = strings.NewReader(bodyData)
	}

	// 创建请求
	req, err := http.NewRequest(method, target, reqBody)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}

	// 设置默认 User-Agent
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/138.0.0.0 Safari/537.36")

	// 设置自定义请求头
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %v", err)
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应失败: %v", err)
	}

	// 构造返回值
	response := &Response{
		Body:       body,
		StatusCode: resp.StatusCode,
		Header:     resp.Header,
		URL:        resp.Request.URL.String(),
		Duration:   time.Since(start),
		Raw:        resp,
	}

	return response, nil
}
