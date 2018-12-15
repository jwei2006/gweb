package sys

import (
	"fmt"
	"gweb/config"
	"log"
	"os"
)

var _debug *log.Logger
var _info *log.Logger
var _warn *log.Logger
var _error *log.Logger

func InitLogger() {
	if config.Conf.IsDebug {
		_debug = log.New(os.Stdout, "[debug] ", log.LstdFlags|log.Lshortfile)
		_info = log.New(os.Stdout, "[info] ", log.LstdFlags|log.Lshortfile)
		_warn = log.New(os.Stdout, "[warn] ", log.LstdFlags|log.Lshortfile)
		_error = log.New(os.Stderr, "[error] ", log.LstdFlags|log.Lshortfile)
		return
	}
	file, e := os.Create(config.Conf.LogInfoPath)
	if e != nil {
		fmt.Println(e)
		log.Panicln(e)
	} else {
		fmt.Println("log info创建成功")
		log.Println("log info创建成功")
	}
	errFile, e1 := os.Create(config.Conf.LogErrorPath)
	if e1 != nil {
		fmt.Println(e1)
		log.Panicln(e)
	} else {
		fmt.Println("log err创建成功")
		log.Println("log err创建成功")
	}
	_debug = log.New(file, "[debug] ", log.LstdFlags|log.Lshortfile)
	_info = log.New(file, "[info] ", log.LstdFlags|log.Lshortfile)
	_warn = log.New(file, "[warn] ", log.LstdFlags|log.Lshortfile)
	_error = log.New(errFile, "[error] ", log.LstdFlags|log.Lshortfile)
}

func Debug() *log.Logger {
	return _debug
}
func Info() *log.Logger {
	return _info
}
func Warn() *log.Logger {
	return _warn
}
func Error() *log.Logger {
	return _error
}
