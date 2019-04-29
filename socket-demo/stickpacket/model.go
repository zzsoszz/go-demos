package main

import (
	"log"
)

type Header struct {
	Version [2]byte // 协议版本
	Length  int     // 数据部分长度
}


type Packet struct {
	Total uint16
	Length uint16
	Header Header
	Data   []byte
}

func A() {
	log.Println("AAAA")
}
func B() {
	log.Println("BBB")
}
