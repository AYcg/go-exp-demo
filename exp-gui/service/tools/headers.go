package tools

import "sync"

// 全局变量存储请求头
var (
	globalHeaders = make(map[string]string)
	headerMutex   sync.RWMutex
)

// SetGlobalHeaders 设置全局请求头
func SetGlobalHeaders(headers map[string]string) {
	headerMutex.Lock()
	defer headerMutex.Unlock()

	// 清空并设置
	globalHeaders = make(map[string]string)
	for k, v := range headers {
		globalHeaders[k] = v
	}
}

// GetGlobalHeaders 获取全局请求头（返回副本）
func GetGlobalHeaders() map[string]string {
	headerMutex.RLock()
	defer headerMutex.RUnlock()

	headersCopy := make(map[string]string)
	for k, v := range globalHeaders {
		headersCopy[k] = v
	}
	return headersCopy
}
