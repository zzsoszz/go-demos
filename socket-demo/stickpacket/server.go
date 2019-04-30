package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"net"
)

/**
 *	https://medium.com/@kpbird/golang-serialize-struct-using-gob-part-2-f6134dd4f22c
 */
func Start()  {

	//test6();
	test7();
	

}

func test7()  {
	con,err:=net.Listen("tcp","0.0.0.0:2020");
	if err ==nil {
		fmt.Println("listen on 2020")
		for{
			clientCon,err:=con.Accept();
			fmt.Println("connected client:"+con.Addr().String())
			if(err==nil){
				go dealClient(clientCon);
			}
		}
	}else{
		fmt.Println(err);
	}
}



func dealClient(clientCon net.Conn){
	var data=[]byte("hello");
	var version=[2]byte{0x00,0x00};
	var packet=Packet{
		Header:Header{
			Version:version,
			Length:int32(len(data)),
		},
		Data:data,
	};


	var network bytes.Buffer
	enc := gob.NewEncoder(&network)
	err := enc.Encode(&packet);
	if err ==nil {
		fmt.Println(network.Bytes());
		clientCon.Write(network.Bytes());
		clientCon.Write(network.Bytes());
		clientCon.Write(network.Bytes());
		clientCon.Close();
		fmt.Println("write finesh!")
	}

}


func test6()  {

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
