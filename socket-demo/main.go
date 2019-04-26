package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
)

/**
* https://www.jb51.net/article/92565.htm
 */
func main()  {
	fmt.Println("111");

	conn,err:=net.DialIP("ip4:icmp", &net.IPAddr{IP: net.ParseIP("0.0.0.0")}, &net.IPAddr{IP: net.ParseIP("127.0.0.1")});
	checkError(err)

	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)

	result, err := readFully(conn)
	checkError(err)

	fmt.Println(string(result))

	os.Exit(0)
}

func readFully(conn net.Conn) ([]byte, error) {
	defer conn.Close()

	result := bytes.NewBuffer(nil)
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		result.Write(buf[0:n])
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
	}
	return result.Bytes(), nil
}


func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}