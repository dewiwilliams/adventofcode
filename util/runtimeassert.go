package util

import (
	"log"
	"runtime/debug"
)

func RuntimeAssert(condition bool) {
	if !condition {
		debug.PrintStack()
		log.Fatalln("Assertion failed")
	}
}
