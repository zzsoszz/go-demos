package  main

import (
	"encoding/gob"
	"fmt"
	"net"
)




func StartClient()  {

	tcpAddr, err := net.ResolveTCPAddr("tcp4", "192.168.1.110:2020");
	if err == nil {
		con,err:=net.DialTCP("tcp",nil,tcpAddr);

		if err ==nil {

			//var b []byte=make([]byte,1024);
			//con.Read(b);
			//fmt.Println(b)

			var packet=&Packet{};
			dec := gob.NewDecoder(con)

			err = dec.Decode(packet)
			fmt.Println(string(packet.Data));

		}else{
			fmt.Println(err);
		}


	}else{
		fmt.Println(err);
	}

}