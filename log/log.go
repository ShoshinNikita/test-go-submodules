package log

import "log"

func Printf(format string, args ...interface{}) {
	log.Printf(format, args...)
}
