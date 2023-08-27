package log

import "fmt"

// 日志模块

func Info(text string) {
	fmt.Printf("\033[0;34m%s\033[0m\n", "[INFO]:"+text)

}

func Warn(text string) {
	fmt.Printf("\033[0;33m%s\033[0m\n", "[WARN]:"+text)
}

func Success(text string) {
	fmt.Printf("\033[0;32m%s\033[0m\n", "[OK]:"+text)
}

func Err(text string, err error) {
	fmt.Printf("\033[1;31m%s\033[0m\n", "[ERR]:"+text)
	fmt.Printf("\033[1;31m%s\033[0m\n", "\t[ERR]:"+err.Error())
}
