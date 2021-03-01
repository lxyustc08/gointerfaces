package main

import (
	"github.com/lxyustc08/gointerfaces/gocmdflags"
	"github.com/lxyustc08/gointerfaces/tokenxml"
)

//var period = flag.Duration("period", 1*time.Second, "sleep period")

var temp = gocmdflags.CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	tokenxml.TestXMLDecoder()
}
