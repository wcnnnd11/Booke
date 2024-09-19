package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4}
	Reverse(s)
	fmt.Println(s)

}

// Reverse 切片反转
func Reverse[T any](slice []T) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}
