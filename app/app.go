package app

import (
	stdlog "log"

	applog "github.com/ShoshinNikita/test-go-submodules/log"
	lgr "github.com/go-pkgz/lgr"
)

func Start() {
	stdlog.Printf("Hello World")
	applog.Printf("Hello World")
	lgr.Printf("Hello World")
}
