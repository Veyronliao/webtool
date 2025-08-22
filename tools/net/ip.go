package net

import (
	"fmt"
	"net"
	"net/http"
)

func GetRemoteClientIp(r *http.Request) string {
	remoteIp := r.RemoteAddr
	if ip := r.Header.Get("X-Real-IP"); ip != "" {
		remoteIp = ip
	} else if ip = r.Header.Get("X-Forwarded-For"); ip != "" {
		remoteIp = ip
	} else {
		remoteIp, _, _ = net.SplitHostPort(remoteIp)
	}
	//本地ip
	if remoteIp == "::1" {
		remoteIp = "127.0.0.1"
	}
	return remoteIp
}

func GetPreferredIPv4() (net.IP, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 || iface.Flags&net.FlagLoopback != 0 {
			continue
		}
		addrs, err := iface.Addrs()
		if err != nil {
			continue
		}
		for _, addr := range addrs {
			ipNet, ok := addr.(*net.IPNet)
			if ok && ipNet.IP.To4() != nil && !ipNet.IP.IsLinkLocalUnicast() {
				return ipNet.IP, nil
			}
		}
	}
	return nil, fmt.Errorf("no suitable IPv4 address found")
}
func GetLocalIPv4s() ([]net.IP, error) {
	var ips []net.IP
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}
	for _, addr := range addrs {
		ipNet, ok := addr.(*net.IPNet)
		if ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
			ips = append(ips, ipNet.IP)
		}
	}
	return ips, nil
}
