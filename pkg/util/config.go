package util

import (
	"os"

	gcfg "gopkg.in/gcfg.v1"
)

func ReadModuleConfig(cfg interface{}, path string, module string) bool {

	fname := os.Getenv("GOPATH") + path + module + ".ini"
	err := gcfg.ReadFileInto(cfg, fname)
	if err == nil {
		return true
	}
	return false
}
