package main

import (
	"EntryTask/Server"
	"fmt"
	"net"
)

func main(){
		fmt.Printf("Start server..")
		listen, err :=net.Listen("tcp","127.0.0.1:5000")
		if err!=nil{
			fmt.Println("Listen failed,err:",err)
			return
		}
		for{
			conn, err:=listen.Accept()
			if err!=nil{
				fmt.Println("Accept failed, err:",err)
				continue
			}
			go process(conn)
		}
	}

	func process(conn net.Conn){
		defer conn.Close()
		for{
			buf :=make([]byte,512)
			n, err :=conn.Read(buf)
			if err!=nil{
				fmt.Println("read err:", err)
				return
			}
			//fmt.Printf(string(buf[0:n]))
			Server.CheckInfo(string(buf[0:n]))
		}
	}