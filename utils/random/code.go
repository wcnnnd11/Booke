package random

import (
	"fmt"
	"math/rand"
	"time"
)

//var stringCode = "abcdefghijklmnopqrstuvwxyz"

func Code(length int) string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%4v", rand.Intn(10000))
}
