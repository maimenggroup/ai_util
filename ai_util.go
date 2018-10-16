package ai_util

import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"github.com/cihub/seelog"
	"time"
	"math/rand"
)

func WaitSignalToStop() {
	stop := make(chan os.Signal)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	s := <-stop
	seelog.Infof("get signal[%v]. syncer will exit", s)
}

func ParseConfigPath(conffile string) (string, string, string, error) {
	confdir := filepath.Dir(conffile)
	confbase := filepath.Base(conffile)
	nametype := strings.Split(confbase, ".")
	if len(nametype) != 2 {
		return "", "", "", errors.New(fmt.Sprintf("confbase has no suffix type. confbase: [%v]", confbase))
	}
	confname := nametype[0]
	conftype := nametype[1]

	return confdir, confname, conftype, nil
}

func IsFileExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func GetRandomString(length int) string{
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

