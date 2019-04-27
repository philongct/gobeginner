package remote

import (
	"rsc.io/sampler"
	"strings"
)

func SayHi() string {
	return strings.Replace(sampler.Hello(), "Hello", "Hi", -1)
}