package main

import "fmt"

func main() {
	pool := NewPool()
	tt := pool.Get().(*Object)
	fmt.Println(tt.name)
}
