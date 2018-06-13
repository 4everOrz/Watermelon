package webscoket

//借助工具测试，浏览器打开 http://www.blue-zero.com/WebSocket/
//输入地址  ws://127.0.0.1:8080/v1/ws?user=1
//点击连接即可在发送框发送消息
//如果成功的话你可以在接收区域看到服务器回复了你同样的消息内容
import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:   1024,
	WriteBufferSize:  1024,
	HandshakeTimeout: 5 * time.Second, // 取消ws跨域校验
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WebscoketJoin(c *gin.Context) {
	user := c.GetString("username")
	ws, err := wsupgrader.Upgrade(c.Writer, c.Request, nil)
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(c.Writer, "Not a websocket handshake", 400)
		return
	} else if err != nil {
		log.Println("Cannot setup WebSocket connection:", err)
		return
	}

	Join(user, ws) //新增客户端连接
	defer Leave(user)
	for {
		_, msg, err := ws.ReadMessage() //读取客户端消息
		if err != nil {
			log.Println("reading message failed")
		}
		publish <- newEvent(msg)
	}
}

func broadcast(event Event) {
	msg := event.Content
	for sub := subscribers.Front(); sub != nil; sub = sub.Next() {
		ws := sub.Value.(Subscriber).Conn
		if ws != nil {
			err := ws.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				unsubscribe <- sub.Value.(Subscriber).Name
			}
		}
	}
}
