package main

import (
	"fmt"
	"net"
)

//处理套接字
func process(conn net.Conn) {

	//创建一个切片用以接收套接字
	buf := make([]byte, 1024)
	//Conn接口下Read方法两个返回值，一个长度 int，一个err类型
	n, err := conn.Read(buf)
	if err != nil {
		return
	}
	//类型转换，把buf转成string类型
	fmt.Println(string(buf[:n]))

}

func main() {
	//net包的Listen方法，返回一个Listner接口和一个错误
	//Listener接口里有三个方法，Addr(),Accpet(),Close()
	listen, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("服务器开始监听端口")
	for {
		//Listner接口的accpet方法
		//返回Conn接口

		conn, err := listen.Accept()
		if err != nil {
			fmt.Println(err)
			//return
		}
		fmt.Printf("服务器监听来自%s的请求: \n", conn.RemoteAddr().String())
		defer listen.Close()
		//开启协程处理套接字
		go process(conn)

	}
}
