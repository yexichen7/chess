/**
* @Author: yexichen
* @Date:2023/7/11-19:23
* @Desc
**/

package api

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net"
	"net/http"
	"sync"
	"time"
)

const (
	DelayTime         = time.Second / 20 //延迟
	ReceiveDataLength = 589              //接受数据长度
)

type PositionStruct struct {
	SdPlayer    int      `json:"SdPlayer"`    // 轮到谁走，0=红方，1=黑方
	UcpcSquares [256]int `json:"UcpcSquares"` // 棋盘上的棋子
	RoomId      string   `json:"RoomId"`      //棋手ID
}

var wg sync.WaitGroup //进程锁
var up = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
} //websocket协议升级结构体

var MsgCh = make(chan PositionStruct, 500)  //信息广播通道
var mplock sync.Mutex                       //map锁
var mlock sync.Mutex                        //map锁
var room = make(map[*websocket.Conn]string) //房间map
var roomcount = make(map[string]int)        //房间人数

type Server struct {
	Ip   string
	Port int
} //服务器

func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:   ip,
		Port: port,
	}
	return server
}

// Start 启动
func (this *Server) Start() {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		fmt.Println("net listen err:", err)
	}
	defer listener.Close()
}

// WsBroadcast 广播websocket
func WsBroadcast() {
	for {

	}
}
