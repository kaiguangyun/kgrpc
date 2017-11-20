package debug

import (
	"log"
	"runtime"
)

// print file/line/function_name
func prePrint() {
	_, file, line, _ := runtime.Caller(2)
	println("")
	log.Printf("error file : %v (line : %v)", file, line)
}

// print error
func Println(value ...interface{}) {
	prePrint()
	log.Println(value...)
}

// print error
func Printf(format string, value ...interface{}) {
	prePrint()
	log.Printf(format+"\n", value...)
}

// print error and write log storage
func Info(value ...interface{}) {
	prePrint()
	log.Println(value...)
}

// print error and write log storage
func Infof(format string, value ...interface{}) {
	prePrint()
	log.Printf(format+"\n", value...)
}

// print error and write log storage
func Warning(value ...interface{}) {
	prePrint()
	log.Println(value...)
}

// print error and write log storage
func Warningf(format string, value ...interface{}) {
	prePrint()
	log.Printf(format+"\n", value...)
}

// print error and write log storage
func Danger(value ...interface{}) {
	prePrint()
	log.Println(value...)
}

// print error and write log storage
func Dangerf(format string, value ...interface{}) {
	prePrint()
	log.Printf(format+"\n", value...)
}

// print error and write log storage, then exit
func Fatal(value ...interface{}) {
	prePrint()
	log.Fatalln(value...)
}

// print error and write log storage, then exit
func Fatalf(format string, value ...interface{}) {
	prePrint()
	log.Fatalf(format+"\n", value...)
}
