package mq

import (
	"github.com/go-stomp/stomp"
	"log"
)

var config = struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}{
	"192.168.8.246", //PROD
	"61613",
	"admin",
	"admin",
}

// 使用IP和端口连接到ActiveMQ服务器
// 返回ActiveMQ连接对象
func GetConnection() (stompConn *stomp.Conn) {
	stompConn, err := stomp.Dial("tcp", config.Host+":"+config.Port)
	if err != nil {
		log.Fatal("connect to active_mq server service, error: " + err.Error())
	}
	//log.Println("connect to active_mq server success: ")
	return stompConn
}

// 将消息发送到ActiveMQ中
func SendMessage(body []byte, destination string, conn *stomp.Conn) {
	err := conn.Send(destination, "text/plain", body)
	if err != nil {
		//log.Println("active mq message send erorr: " + err.Error())
	}
	//log.Println(fmt.Sprintf("send sucessful! destination: %s, msg body: %s.", destination, body))
}
