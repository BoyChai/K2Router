package control

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func Detection(ch chan string, wg *sync.WaitGroup) {
	for ip := range ch {
		if isPortOpen(ip, 80) {
			fmt.Println(ip)
			IpPool = append(IpPool, ip)
		}
		wg.Done()
	}
}
func isPortOpen(host string, port int) bool {
	// 构建目标地址
	target := fmt.Sprintf("%s:%d", host, port)

	// 尝试连接目标地址的80端口，设置超时时间为2秒
	conn, err := net.DialTimeout("tcp", target, 1*time.Second)
	if err != nil {
		return false // 连接失败，端口不开放
	}

	defer conn.Close()
	return true // 连接成功，端口开放
}

// IncIP 增加IP地址
func IncIP(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}
