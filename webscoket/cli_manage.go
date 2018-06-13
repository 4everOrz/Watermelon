package webscoket

import (
	"container/list"
	"time"

	"github.com/gorilla/websocket"
)

type Subscription struct {
	Archive []Event //缓存数组
	New     <-chan Event
}

func newEvent(msg []byte) Event {
	return Event{int(time.Now().Unix()), msg}
}

func Join(user string, ws *websocket.Conn) {
	subscribe <- Subscriber{Name: user, Conn: ws}
}

func Leave(user string) {
	unsubscribe <- user
}

type Subscriber struct {
	Name string
	Conn *websocket.Conn
}

type Event struct {
	Timestamp int
	Content   []byte
}

var (
	subscribe   = make(chan Subscriber, 20)
	unsubscribe = make(chan string, 20)
	publish     = make(chan Event, 20)
	subscribers = list.New()
)

func init() {
	go manage()
}
func manage() {
	for {
		select {
		case sub := <-subscribe: //读取到 new client 加入
			if !isUserExist(subscribers, sub.Name) {
				subscribers.PushBack(sub) //new client 放入队列尾
			}
		case event := <-publish: //有数据传入
			broadcast(event) //向各个客户端广播数据
		case unsub := <-unsubscribe:
			for sub := subscribers.Front(); sub != nil; sub = sub.Next() {
				if sub.Value.(Subscriber).Name == unsub {
					subscribers.Remove(sub)
					ws := sub.Value.(Subscriber).Conn
					if ws != nil {
						ws.Close()
					}
					break
				}
			}
		}
	}
}

//判断用户是否已经在队列中
func isUserExist(subscribers *list.List, user string) bool {
	for sub := subscribers.Front(); sub != nil; sub = sub.Next() {
		if sub.Value.(Subscriber).Name == user {
			return true
		}
	}
	return false
}
