package control

import (
	"fmt"
	"fyne.io/fyne/v2/widget"
	"net/http"
	"net/url"
	"strings"
	"sync"
)

const (
	//ipAddress        = "192.168.2.1"
	wirelessSetting  = "/cgi-bin/luci/admin/quickguide/wireless_setting"
	internetSetting  = "/cgi-bin/luci/admin/quickguide/internet_setting"
	defaultUserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36"
)

func SetRouter(ipAddress, adminPass, name2g, pass2g, name5g, pass5g string, text *widget.Label, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("初始化" + ipAddress)
	text.SetText("初始化中")
	err := getRequest(ipAddress, wirelessSetting, "POST", url.Values{"userprotocol": {"1"}, "configclick": {"click"}})
	if err != nil {
		fmt.Println(1)
		return
	}
	fmt.Println(ipAddress + ":设置上网方式")
	text.SetText("设置上网方式")
	err = postRequest(ipAddress, internetSetting, "POST", "application/x-www-form-urlencoded", url.Values{"connectionType": {"DHCP"}, "autodetect": {"0"}})
	if err != nil {
		fmt.Println(2)
		return
	}
	fmt.Println(ipAddress + ":开始设置账号密码")
	text.SetText("设置账号密码")
	err = postRequest(ipAddress, wirelessSetting, "POST", "application/x-www-form-urlencoded",
		url.Values{"savevalidate": {"1"}, "username": {"admin"}, "ssid": {name2g}, "key": {pass2g},
			"inic_ssid": {name5g}, "inic_key": {pass5g}, "password": {adminPass}})
	if err != nil {
		fmt.Println(3)
		return
	}
	fmt.Println("设置成功")
	text.SetText("设置成功")
}

func getRequest(ip, endpoint, method string, params url.Values) error {
	url := "http://" + ip + endpoint

	req, err := http.NewRequest(method, url, strings.NewReader(params.Encode()))
	if err != nil {
		fmt.Printf("创建 %s 请求错误: %v\n", method, err)
		return err
	}

	setupRequestHeaders(ip, req)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		FailPool = append(FailPool, ip)
		fmt.Printf(ip+":发送 %s 请求错误: %v\n", method, err)
		return err
	}
	defer resp.Body.Close()

	return nil
}

func postRequest(ip, endpoint, method, contentType string, data url.Values) error {
	url := "http://" + ip + endpoint

	req, err := http.NewRequest(method, url, strings.NewReader(data.Encode()))
	if err != nil {
		FailPool = append(FailPool, ip)
		fmt.Printf("创建 %s 请求错误: %v\n", method, err)
		return err
	}

	setupRequestHeaders(ip, req)
	req.Header.Set("Content-Type", contentType)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		FailPool = append(FailPool, ip)
		fmt.Printf("发送 %s 请求错误: %v\n", method, err)
		return err
	}
	defer resp.Body.Close()
	return nil
}

func setupRequestHeaders(ipAddress string, req *http.Request) {
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Origin", "http://"+ipAddress)
	req.Header.Set("Referer", "http://"+ipAddress+"/cgi-bin/luci/")
	req.Header.Set("User-Agent", defaultUserAgent)
}
