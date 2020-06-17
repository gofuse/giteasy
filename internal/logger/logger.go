package logger

import (
	"fmt"
	"log"
)

var (
	nocolor = "%v"
	green   = "\033[0;32m%v\033[0m"
	yellow  = "\033[0;33m%v\033[0m"
	red     = "\033[0;31m%v\033[0m"
)

func Println(args ...interface{}) {
	log.Println(fmt.Sprintf(nocolor, fmt.Sprint(args...)))
}

func Info(args ...interface{}) {
	log.Println(fmt.Sprintf(green, fmt.Sprint(args...)))
}

func Debug(args ...interface{}) {
	log.Println(fmt.Sprintf(nocolor, fmt.Sprint(args...)))
}

func Warn(args ...interface{}) {
	log.Println(fmt.Sprintf(yellow, fmt.Sprint(args...)))
}

func Error(args ...interface{}) {
	log.Println(fmt.Sprintf(red, fmt.Sprint(args...)))
}

func Errorf(args ...interface{}) string {
	return fmt.Sprintf(red, fmt.Sprint(args...))
}

func Fatal(args ...interface{}) {
	log.Fatal(fmt.Sprintf(red, fmt.Sprint(args...)))
}
