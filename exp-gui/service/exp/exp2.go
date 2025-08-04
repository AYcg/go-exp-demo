package exp

import (
	"changeme/service/tools"
	"fmt"
	"strings"
)

func init() {
	RegisterExploit("demo2", ExploitEntry{
		Check:   Demo2Check,
		Exploit: Demo2Exploi_cmd,
		Upload:  Demo2Exploit_upload,
	})
}

// Check 检测漏洞
func Demo2Check(target string, headers map[string]string, proxy string) (bool, string) {
	resp, err := tools.SendRequest(target, "GET", "", headers, proxy)
	if err != nil {
		return false, fmt.Sprintf("[!] 请求失败: %v", err)
	}

	if strings.Contains(string(resp.Body), "refresh") {

		return true, "[+] 检测到refresh字段，可能存在漏洞"
	}

	return false, "[-] demo2漏洞未发现"
}

// Exploit 命令执行利用漏洞
func Demo2Exploi_cmd(target string, headers map[string]string, proxy string, payload string) string {
	resp, err := tools.SendRequest(target, "POST", payload, headers, proxy)
	if err != nil {
		return fmt.Sprintf("[!] 利用失败: %v", err)
	}

	// 这里写你的实际利用逻辑
	//if strings.Contains(string(resp.Body), "Found") {
	//	return "[+] Ruoyi漏洞利用成功"
	//}
	if resp.StatusCode == 302 {
		return "[+] demo2命令执行漏洞利用成功"
	}

	return "[-] demo2漏洞利用失败"
}

// Exploit 文件上传利用漏洞
func Demo2Exploit_upload(target string, headers map[string]string, proxy string, payload string, filename string) string {
	body := "filename:" + filename + "&filecontext:" + payload
	resp, err := tools.SendRequest(target, "POST", body, headers, proxy)
	if err != nil {
		return fmt.Sprintf("[!] 利用失败: %v", err)
	}

	// 这里写你的实际利用逻辑
	//if strings.Contains(string(resp.Body), "Found") {
	//	return "[+] Ruoyi漏洞利用成功"
	//}
	if resp.StatusCode == 302 {
		return "[+] demo2漏洞利用成功"
	}

	return "[-] demo2漏洞利用失败"
}
