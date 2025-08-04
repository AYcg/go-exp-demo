package tools

import (
	"net/http"
	"net/url"
	"sync"
	"time"
)

// 全局变量存储代理
var (
	globalProxy string
	proxyMutex  sync.RWMutex
)

// SetProxy 设置全局代理
func SetProxy(proxy string) {
	proxyMutex.Lock()
	defer proxyMutex.Unlock()
	globalProxy = proxy
}

// GetProxy 获取当前代理设置
func GetProxy() string {
	proxyMutex.RLock()
	defer proxyMutex.RUnlock()
	return globalProxy
}

// TestProxy 测试代理是否可用
func TestProxy(proxy string) (bool, string) {
	if proxy == "" {
		return false, "代理地址为空"
	}

	// 验证代理URL格式
	proxyURL, err := url.Parse(proxy)
	if err != nil {
		return false, "代理URL格式错误: " + err.Error()
	}

	// 创建带超时的HTTP客户端
	client := &http.Client{
		Timeout: 5 * time.Second,
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		},
	}

	// 测试地址
	testURL := "http://baidu.com"
	req, err := http.NewRequest("GET", testURL, nil)
	if err != nil {
		return false, "创建请求失败: " + err.Error()
	}

	// 发送测试请求
	resp, err := client.Do(req)
	if err != nil {
		return false, "代理连接失败: " + err.Error()
	}
	defer resp.Body.Close()

	// 状态码检查
	if resp.StatusCode >= 400 {
		return false, "代理返回错误状态码: " + resp.Status
	}

	return true, "代理测试成功"
}
