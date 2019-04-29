package main

//
//import (
//	"bufio"
//	"bytes"
//	"fmt"
//	"io"
//	"net"
//	"os"
//	"strings"
//)
//
//
//
///**
//* https://www.jb51.net/article/92565.htm
// */
//func main()  {
//	//test1();
//	//test2();
//	//test3();
//	test4();
//}
//
//
//
///**
// * advance 取下一个token需要前进几步(下一个token起始位置)
// * token  当前token
// */
////return i + 1, data[0:i], nil
//func findNextToken(data []byte,atEOF bool) (advance int,token []byte,err error) {
//	if atEOF && len(data) == 0 {
//		return 0, nil, nil
//	}
//	i:=bytes.IndexByte(data,'\n');
//	if i>=0{
//		return i+1,data[0:i],nil;
//	}
//	if atEOF {
//		return len(data),data, nil
//	}
//	return 0, nil,nil;
//}
//
//
//
//func test4()  {
//	input := "foo\n  foo\nfoo";
//	scanner := bufio.NewScanner(strings.NewReader(input))
//	split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
//		 return findNextToken(data,atEOF);
//	}
//	//bufio.ScanLines
//	//bufio.ScanWords
//	scanner.Split(split)
//	for scanner.Scan() {
//		fmt.Printf("%s\n", scanner.Text())
//	}
//}
//
//
//
//func test3()  {
//
//	proverbs := []string{
//		"Channels orchestrate mutexes serialize",
//		"Cgo is not Go",
//		"Errors are values",
//		"Don't panic",
//	}
//	var writer bytes.Buffer
//
//	for _, p := range proverbs {
//		n, err := writer.Write([]byte(p))
//		if err != nil {
//			fmt.Println(err)
//			os.Exit(1)
//		}
//		if n != len(p) {
//			fmt.Println("failed to write data")
//			os.Exit(1)
//		}
//	}
//
//	fmt.Println(writer.String())
//
//}
//func test1()  {
//	reader := strings.NewReader("Clear is better than clever")
//	p := make([]byte, 4)
//	for {
//		n, err := reader.Read(p)
//		if err == io.EOF {
//			break
//		}
//		fmt.Println(string(p[:n]))
//	}
//}
//
//
//
//func test2()  {
//	fmt.Println("111");
//
//	conn,err:=net.DialIP("ip4:icmp", &net.IPAddr{IP: net.ParseIP("0.0.0.0")}, &net.IPAddr{IP: net.ParseIP("192.168.1.110")});
//	checkError(err)
//
//	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
//	checkError(err)
//
//	result, err := readFully(conn)
//	checkError(err)
//
//	fmt.Println(string(result))
//
//	os.Exit(0)
//}
//func readFully(conn net.Conn) ([]byte, error) {
//	defer conn.Close()
//
//	result := bytes.NewBuffer(nil)
//	var buf [512]byte
//	for {
//		n, err := conn.Read(buf[0:])
//		result.Write(buf[0:n])
//		if err != nil {
//			if err == io.EOF {
//				break
//			}
//			return nil, err
//		}
//	}
//	return result.Bytes(), nil
//}
//
//
//func checkError(err error) {
//	if err != nil {
//		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
//		os.Exit(1)
//	}
//}