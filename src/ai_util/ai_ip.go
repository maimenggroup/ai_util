package ai_util

import (
	"errors"
	"io/ioutil"
	"net"
	"time"
)

func HLocalIp() (net.IP, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return net.IPv4zero, err
	}

	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP, nil
			}
		}
	}
	return net.IPv4zero, errors.New("unknown error")
}

func HPublicIp() (net.IP, error) {
	conn, err := net.DialTimeout("tcp", "ns1.dnspod.net:6666", 3*time.Second)
	if err != nil {
		return net.IPv4zero, err
	}
	buf, err := ioutil.ReadAll(conn)
	if err != nil {
		return net.IPv4zero, err
	}
	return net.ParseIP(string(buf)), nil
}
