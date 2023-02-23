package main

import "sync"

// Object 对象池模式
type Object struct {
	name string
}

func NewObject() interface{} {
	return &Object{name: "ob"}
}
func NewPool() *sync.Pool {
	pool := &sync.Pool{
		New: NewObject,
	}
	return pool
}
