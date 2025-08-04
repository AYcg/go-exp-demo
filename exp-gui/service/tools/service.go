package tools

// Service 是一个包装器，用于将 config 包的方法暴露给 Wails
type Service struct{}

// SetGlobalHeaders 调用 config 包的函数
func (s *Service) SetGlobalHeaders(headers map[string]string) error {
	SetGlobalHeaders(headers)
	return nil
}

// GetGlobalHeaders 调用 config 包的函数
func (s *Service) GetGlobalHeaders() map[string]string {
	return GetGlobalHeaders()
}

// SetProxy 调用 config 包的函数
func (s *Service) SetProxy(proxy string) error {
	SetProxy(proxy)
	return nil
}

// GetProxy 调用 config 包的函数
func (s *Service) GetProxy() string {
	return GetProxy()
}

// TestProxy 调用 config 包的函数
func (s *Service) TestProxy(proxy string) (bool, string) {
	return TestProxy(proxy)
}
