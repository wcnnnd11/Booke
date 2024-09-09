package main

import (
	"GVB_server/utils/random"
	"fmt"
)

func main() {
	s := random.RandString(16)
	fmt.Println(s)
}
