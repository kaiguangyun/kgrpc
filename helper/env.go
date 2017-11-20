package helper

import (
	"github.com/kaiguangyun/kgrpc/debug"
	"os"
)

// Get Env Or Service Config
func GetEnv(key string) (value string) {
	value = os.Getenv(key)

	if value == "" {
		debug.Infof("GetEnv(%v) error : %v is empty or not exist", key, key)
	}

	return value
}
