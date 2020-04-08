package main 

import(
	"fmt"
	"net"
)

func login(userId int, userPwd string) (err error){
	//net包的Dial函数连接服务端
	conn, err := net.Dial("tcp", "127.0.0.1:8888"）
	if err != nil {
		fmt.Println(err)
		return
	}
	
}