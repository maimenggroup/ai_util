package ai_util

import (
	"testing"
	"fmt"
)

func TestParseConfigPath(t *testing.T) {
	confFile := "../conf/offline/crawler.json"
	confDir, confBase, confType, err := ParseConfigPath(confFile)
	if err != nil {
		t.Error("TestParseConfigPath failed")
	} else {
		t.Logf("dir[%s], base[%s], type[%s]", confDir, confBase, confType)
	}
}

func TestGetRandomString(t *testing.T) {
	st := GetRandomString(26)
	Cookie := fmt.Sprintf("PHPSESSID=%s", st)
	t.Log(Cookie)
}