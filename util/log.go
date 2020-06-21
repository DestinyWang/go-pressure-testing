package util

import "runtime"

var LogFormatter = "FuncName=[%s]"

func RunFuncName() string {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	return runtime.FuncForPC(pc[0]).Name()
}
