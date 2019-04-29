package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

func Start()  {

	type Address struct {
		Country    string
	}

	type P struct {
		X, Y, Z int
		Name    string
		Address Address
	}

	type Q struct {
		X, Y *int32
		Name string
		Address Address
	}

	var network bytes.Buffer
	enc := gob.NewEncoder(&network)
	dec := gob.NewDecoder(&network)
	// Encode (send) the value.
	err := enc.Encode(P{3, 4, 5, "Pythagoras",Address{
		"china",
	}})
	err = enc.Encode(P{3, 4, 5, "Pythagoras",Address{
		"china",
	}})


	if err != nil {
		log.Fatal("encode error:", err)
	}
	// Decode (receive) the value.
	var q Q
	var q2 Q
	err = dec.Decode(&q)
	err = dec.Decode(&q2)
	if err != nil {
		log.Fatal("decode error:", err)
	}
	fmt.Println(q)
	fmt.Printf("%q: {%d,%d,%s}\n", q.Name, *q.X, *q.Y,q.Address)
	fmt.Printf("%q: {%d,%d,%s}\n", q2.Name, *q2.X, *q2.Y,q2.Address)


}





//
////var data=[]byte("hello");
////var version=[2]byte{0x00,0x00};
//var packet=Packet{
//	Total:uint16(11),
//	Length:uint16(13),
//	//Header:Header{
//	//	Version:version,
//	//	Length:len(data),
//	//},
//	//Data:data,
//};
//
///**
// *  json格式
// */
//jsondata,err:=json.Marshal(packet);
//if err ==nil{
//	fmt.Println(string(jsondata));
//}
//
////var buffer bytes.Buffer
////binary.Write(&buffer,binary.BigEndian,packet);
//////fmt.Printf("%b",data)
////
////fmt.Println(buffer.Bytes())
////fmt.Printf("%b",buffer.Bytes())
////fmt.Println("");






//
//type T struct {
//	A int64
//	B float64
//}
//
//func main() {
//	// Create a struct and write it.
//	t := T{A: 0xEEFFEEFF, B: 3.14}
//	buf := &bytes.Buffer{}
//	err := binary.Write(buf, binary.BigEndian, t)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println(buf.Bytes())
//
//	// Read into an empty struct.
//	t = T{}
//	err = binary.Read(buf, binary.BigEndian, &t)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Printf("%x %f", t.A, t.B)
//
//}
