package main

import (
	"log"
)

type Header struct {
	Version [2]byte // 协议版本
	Length  int32     // 数据部分长度
}

type Packet struct {
	Header Header
	Data   []byte
}

func A() {
	log.Println("AAAA")
}
func B() {
	log.Println("BBB")
}
