package control

import (
	"bytes"
	"fmt"
	"fyne.io/fyne/v2/widget"
	"net/http"
	"net/url"
	"strings"
	"sync"
)

func SetRouter(ipAddress, adminPass, name2g, pass2g, name5g, pass5g string, text *widget.Label, wg *sync.WaitGroup) {
	defer wg.Done()
	text.SetText("正在初始化路由器...")
	fmt.Println("初始化")
	Get0(ipAddress)

	text.SetText("正在设置上网方式")
	fmt.Println("设置上网方式")
	formData := "connectionType=DHCP&autodetect=0&pppoeUser=+&pppoePass=&staticIp=&staticNetmask=&staticGateway=&staticPriDns="

	req, _ := http.NewRequest("POST", "http://"+ipAddress+"/cgi-bin/luci/admin/quickguide/internet_setting", bytes.NewBufferString(formData))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Referer", "http://"+ipAddress+"/cgi-bin/luci/admin/quickguide/internet_setting")
	req.Header.Set("Origin", "http://"+ipAddress)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36")

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("发送请求错误:", err)
		text.SetText("设置失败")
		return
	}
	defer resp.Body.Close()
	// 读取响应体
	if err != nil {
		fmt.Println("读取响应体错误:", err)
		text.SetText("设置失败")
		return

	}

	text.SetText("正在设置账号密码")
	fmt.Println("开始设置账号密码")
	Get2(ipAddress, name2g, pass2g, name5g, pass5g, adminPass)

	text.SetText("设置成功！")
}

func Get2(ip, name2g, pass2g, name5g, pass5g, pass string) {
	Url := "http://" + ip + "/cgi-bin/luci/admin/quickguide/wireless_setting"

	// 构建表单数据
	formData := url.Values{
		"savevalidate": {"1"},
		"username":     {"admin"},
		"ssid":         {name2g},
		"key":          {pass2g},
		"inic_ssid":    {name5g},
		"inic_key":     {pass5g},
		"password":     {pass},
	}

	// 创建 POST 请求
	req, _ := http.NewRequest("POST", Url, strings.NewReader(formData.Encode()))
	// 设置请求头
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Origin", "http://"+ip+"")
	req.Header.Set("Referer", "http://"+ip+"/cgi-bin/luci/admin/quickguide/wireless_setting")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36")

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("发送请求错误:", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应体
	if err != nil {
		fmt.Println("读取响应体错误:", err)
		return
	}

}

func Get0(ip string) {
	Url := "http://" + ip + "/cgi-bin/luci/admin/quickguide/wireless_setting"

	// 构建表单数据
	formData := url.Values{
		"userprotocol": {"1"},
		"configclick":  {"click"},
	}

	// 创建 POST 请求
	req, err := http.NewRequest("POST", Url, strings.NewReader(formData.Encode()))
	if err != nil {
		fmt.Println("创建请求错误:", err)
		return
	}

	// 设置请求头
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Origin", "http://"+ip)
	req.Header.Set("Referer", "http://"+ip+"/cgi-bin/luci/")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36")

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("发送请求错误:", err)
		return
	}
	defer resp.Body.Close()

	// 输出响应
	fmt.Println("第一步完成")
}
