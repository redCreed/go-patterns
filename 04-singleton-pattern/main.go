package main

import "log"

func main() {
	s := New()
	s["name"] = "lee"
	//验证唯一性 只会初始化一次
	s1 := New()
	if s1["name"] != "lee" {
		log.Fatal("singleton pattern error")
	}
	//change name
	s1["name"] = "anne"
	if s["name"] != "anne" {
		log.Fatal("singleton pattern error")
	}
}
