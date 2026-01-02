package main

import (
	"fmt"
)

func main() {
	fmt.Println("НЕ СУТУЛЬСЯ, пожалуйста")

	go func() {
		fmt.Println("hello from goroutine")
	}()
}
