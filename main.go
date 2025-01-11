package main

import (
	"fmt"
	"net"
	"net/http"
)

func getIP(r *http.Request) string {
	// 从 RemoteAddr 中提取 IP 地址
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return ""
	}
	return ip
}

func handler(w http.ResponseWriter, r *http.Request) {
	// 获取客户端 IP
	ip := getIP(r)

	// 判断是 IPv4 还是 IPv6
	parsedIP := net.ParseIP(ip)
	if parsedIP == nil {
		http.Error(w, "Invalid IP address", http.StatusBadRequest)
		return
	}

	// 返回 IP 地址
	w.Header().Set("Content-Type", "text/plain")
	if _, err := fmt.Fprintf(w, "%s", ip); err != nil {
		fmt.Println("Error write ip")
	}
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server is running on :6655...")
	_ = http.ListenAndServe(":6655", nil)
}
