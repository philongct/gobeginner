package main

import (
	"fmt"
	"github.com/philongct/golessions/modules/remote"
	"github.com/philongct/golessions/modules/local"
)

func main() {
	fmt.Println(remote.Hello())
	fmt.Println(local.Introduce())
}