package message

const (
	LoginMesType = "LoginMes"
)

type Message struct {
	Type string //消息类型，为string
	Data string //消息数据，为string
}

type LoginMes struct {
	UserID   int    //用户ID
	UserPwd  string //密码
	UserName string //用户名
}

type LoginResMes struct {
	Code  int    //返回状态码
	Error string //返回错误信息

}
