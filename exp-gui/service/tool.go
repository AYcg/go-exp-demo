package service

import (
	"changeme/service/exp"
	"changeme/service/tools"
	"errors"
	"strings"
	"sync"
)

type ToolService struct {
	vulnStatus map[string]bool
	mu         sync.RWMutex
}

func NewToolService() *ToolService {
	return &ToolService{
		vulnStatus: make(map[string]bool),
	}
}

// 基本信息说明
func (t *ToolService) GetInfo() string {
	var builder strings.Builder

	// 静态警告信息部分
	builder.WriteString(`
    <style>
      .custom-p {
        margin: 4px 0;
        color: red;
        text-indent: 3.5em;
      }
    </style>
    <div>----------------------------------------------------------------</div>
    <p class="custom-p">本工具仅提供给安全测试人员进行安全自查使用</p>
    <p class="custom-p">用户滥用造成的一切后果与作者无关</p>
    <p class="custom-p">使用者请务必遵守当地法律</p>
    <p class="custom-p">本程序不得用于商业用途，仅限学习交流</p>
    <div>----------------------------------------------------------------</div>
`)

	// 添加支持的漏洞列表
	builder.WriteString(`<div><p class="custom-p">支持的漏洞列表：</p><ul style="margin-left: 3em;">`)

	for _, poc := range t.GetPocList() {
		if poc.Value != "all" {
			builder.WriteString(`<li>` + poc.Label + `</li>`)
		}
	}
	builder.WriteString(`</ul></div>`)

	return builder.String()
}

// 获取支持的漏洞列表
type PocInfo struct {
	Value string `json:"value"`
	Label string `json:"label"`
}

func (t *ToolService) GetPocList() []PocInfo {
	var list []PocInfo
	list = append(list, PocInfo{Value: "all", Label: "全部检测"})
	for name := range exp.ListAll() {
		list = append(list, PocInfo{Value: name, Label: name + " 漏洞"})
	}
	return list
}

// 设置漏洞是否存在状态
func (t *ToolService) SetVulnerable(pocType string, vulnerable bool) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.vulnStatus[pocType] = vulnerable
}

// 查询漏洞是否已确认存在
func (t *ToolService) IsVulnerable(pocType string) bool {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.vulnStatus[pocType]
}

// 控制器主入口：漏洞检测与利用
func (t *ToolService) POC(url string, pocType string, action string, payload string, filename string, expType string) (string, error) {
	if url == "" {
		return "", errors.New("URL不能为空")
	}

	headers := tools.GetGlobalHeaders()
	proxy := tools.GetProxy()

	switch action {
	case "check":
		// 全部检测
		if pocType == "all" {
			var results []string
			for name, entry := range exp.ListAll() {
				if entry.Check != nil {
					exists, msg := entry.Check(url, headers, proxy)
					t.SetVulnerable(name, exists)
					results = append(results, msg)
				} else {
					results = append(results, "[-] "+name+" 不支持检测")
				}
			}
			return strings.Join(results, "<br>"), nil
		}

		// 单个漏洞检测
		entry, ok := exp.GetExploit(pocType)
		if !ok || entry.Check == nil {
			return "", errors.New("未找到对应检测模块: " + pocType)
		}
		exists, msg := entry.Check(url, headers, proxy)
		t.SetVulnerable(pocType, exists)
		return msg, nil

	case "exploit":
		if !t.IsVulnerable(pocType) {
			return "[-] 请先进行漏洞检测", nil
		}

		entry, ok := exp.GetExploit(pocType)
		if !ok {
			return "", errors.New("未找到漏洞模块: " + pocType)
		}

		if expType == "cmd" && entry.Exploit != nil {
			return entry.Exploit(url, headers, proxy, payload), nil
		}

		if expType == "upload" && entry.Upload != nil {
			return entry.Upload(url, headers, proxy, payload, filename), nil
		}

		return "[-] 当前漏洞不支持该利用方式", nil

	default:
		return "", errors.New("不支持的操作类型: " + action)
	}
}
