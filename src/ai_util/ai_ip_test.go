package ai_util

import (
	"testing"
)

func TestHLocalIp(t *testing.T) {
	if ip, err := HLocalIp(); err != nil {
		t.Errorf("hlocal ip failed. ip: [%v], err: [%v].", ip, err)
	} else {
		t.Log(ip)
	}
}

func TestHPublicIp(t *testing.T) {
	if ip, err := HPublicIp(); err != nil {
		t.Errorf("hlocal ip failed. ip: [%v], err: [%v].", ip, err)
	} else {
		t.Log(ip)
	}
}
