package helper

import (
	"log"
	"runtime"
)

// print file/line/function_name
func preOutput() {
	_, file, line, _ := runtime.Caller(2)
	println("")
	log.Printf("output file : %v (line : %v)", file, line)
	//log.Printf("output fuction : %v \n", runtime.FuncForPC(pc).Name())
}

// print
func Output(value ...interface{}) {
	preOutput()
	log.Printf("%+v", value...)
}

// print
func Outputf(format string, value ...interface{}) {
	preOutput()
	log.Printf(format+"\n", value...)
}
