package utils

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"
)

var (
	skipLength int = 0
)

func init() {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		os.Exit(2)
	}
	skipLength = strings.Index(filename, "/utils")
	fmt.Printf("skipLength: %d\n", skipLength)
	if skipLength < 0 {
		skipLength = 0
	}
}

// Log prints a log message with the current time, source file, line number, and the function name that called Log.
// The format is:
//
//	<time> <sourcefile>:<line> <functionname> <message>
//
// The source file path is relative to the current working directory.
func Log(s string) {
	time := time.Now().Format("2006-01-02 15:04:05.000")
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc)
	frame, _ := frames.Next()
	file := "." + frame.File[skipLength:]
	fmt.Printf("%s %s:%d %s %s\n", time, file, frame.Line, frame.Function, s)
}

// Fatal logs an error message with the current time, source file, line number, and the function name that called Fatal.
// The format is:
//
//	<time> <sourcefile>:<line> <functionname> <error message>
//
// The source file path is relative to the current working directory.

func Fatal(e error) {
	time := time.Now().Format("2006-01-02 15:04:05.000")
	pc := make([]uintptr, 20)
	runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc)
	frame, _ := frames.Next()
	file := "." + frame.File[skipLength:]
	fmt.Printf("%s %s:%d %s %s\n", time, file, frame.Line, frame.Function, e.Error())
}
