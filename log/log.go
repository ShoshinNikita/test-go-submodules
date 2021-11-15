package log

import "log"

func Printf(format string, args ...interface{}) {
	log.Printf("log:"+format, args...)
}
