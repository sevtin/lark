package utils

import (
	"runtime"
	"strings"
)

func FuncName() string {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	function := runtime.FuncForPC(pc[0]).Name()
	names := strings.Split(function, ".")
	return names[len(names)-1]
}
