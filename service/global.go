package gotest

import (
	"github.com/joaosoft/go-log/service"
)

var global = make(map[string]interface{})
var log = golog.NewLogDefault("go-test", golog.InfoLevel)

func init() {
	global[path_key] = defaultPath
}
