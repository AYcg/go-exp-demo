package exp

import (
	"changeme/service/tools"
	"fmt"
	"strings"
)

func init() {
	RegisterExploit("demo1", ExploitEntry{
		Check:   Demo1Check,
		Exploit: Demo1Exploi_cmd,
		Upload:  Demo1Exploit_upload,
	})
}

// Check 检测漏洞
func Demo1Check(target string, headers map[string]string, proxy string) (bool, string) {
	resp, err := tools.SendRequest(target, "GET", "", headers, proxy)
	if err != nil {
		return false, fmt.Sprintf("[!] 请求失败: %v", err)
	}

	if strings.Contains(string(resp.Body), "uid") {

		return true, "<p style=\"color:red;display:inline;\">[+] 可能存在demo1漏洞</p>"
	}

	return false, "[-] demo1漏洞未发现"
}

// Exploit 命令执行利用漏洞
func Demo1Exploi_cmd(target string, headers map[string]string, proxy string, payload string) string {
	resp, err := tools.SendRequest(target, "POST", payload, headers, proxy)
	if err != nil {
		return fmt.Sprintf("[!] 利用失败: %v", err)
	}

	// 这里写你的实际利用逻辑
	//if strings.Contains(string(resp.Body), "Found") {
	//	return "[+] Ruoyi漏洞利用成功"
	//}
	if resp.StatusCode == 302 {
		return "[+] demo1命令执行漏洞利用成功"
	}

	return "[-] demo1漏洞利用失败"
}

// Exploit 文件上传利用漏洞
func Demo1Exploit_upload(target string, headers map[string]string, proxy string, payload string, filename string) string {
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
		return "[+] demo1漏洞利用成功"
	}

	return "[-] demo1漏洞利用失败"
}
