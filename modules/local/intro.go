package local

import "runtime"

func Introduce() string {
	return "I'm " + runtime.Version()
}