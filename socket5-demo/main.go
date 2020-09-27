package main

import (
	"io"
	"log"
	"net"
)

func main()  {
	l,err:=net.Listen("tcp","localhost:4242");
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn,err :=l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go proxy(conn);
	}
}

func proxy(conn net.Conn){
	defer conn.Close()
	upstream,err:=net.Dial("tcp","localhost:3000")
	if err!=nil{
		log.Print(err)
		return;
	}
	defer upstream.Close()
	go io.Copy(upstream,conn)
	io.Copy(conn,upstream)
}


//func copyStream(dst io.Writer, src io.Reader){
//	defer conn.Close()
//	for {
//		var buf [128]byte
//		conn.SetReadDeadline(time.Now().Add(5*time.Second))
//		n,err:=conn.Read(buf[:])
//		if err != nil{
//			log.Print(err)
//			return
//		}
//		os.Stderr.Write(buf[:n])
//	}
//	n,err:=io.Copy(os.Stderr,conn);
//	log.Printf("Completed connection with n=%d,err=%v",n,err)
//}

//func handleConnection(conn net.Conn){
//	defer conn.Close()
//	for {
//		var buf [128]byte
//		conn.SetReadDeadline(time.Now().Add(5*time.Second))
//		n,err:=conn.Read(buf[:])
//		if err != nil{
//			log.Print(err)
//			return
//		}
//		os.Stderr.Write(buf[:n])
//	}
//	n,err:=io.Copy(os.Stderr,conn);
//	log.Printf("Completed connection with n=%d,err=%v",n,err)
//}


//func handleConnection(conn net.Conn){
//	n,err:=io.Copy(os.Stderr,conn);
//	log.Printf("Completed connection with n=%d,err=%v",n,err)
//}
