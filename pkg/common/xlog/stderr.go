package xlog

import (
	"runtime"
)

func panicRedirect(path string) (err error) {
	if runtime.GOOS == "windows" {
		return
	}
	return
	//var (
	//	file *os.File
	//)
	//file, err = os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
	//if err != nil {
	//	return
	//}
	//err = syscall.Dup2(int(file.Fd()), int(os.Stderr.Fd()))
	//if err != nil {
	//	return
	//}
	//runtime.SetFinalizer(file, func(fd *os.File) {
	//	fd.Close()
	//})
	return
}
