package logs

import (
	"log"
	"reflect"
	"strings"
)

var V = false

func Logf(format string, a ...interface{}) () {
	if V {
		for i,e := range a {
			if reflect.TypeOf(a[i]).String() == "string" {
				a[i] = strings.Replace(e.(string), "\r\n", "\\n", -1)
				a[i] = strings.Replace(a[i].(string), "\n", "\\n", -1)
			}
		}
		log.Printf(format, a...)
	}
}

func SetLevel(lvl string) () {
	if lvl == "V" {
		V = true
	}
}