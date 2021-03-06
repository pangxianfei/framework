package queue

import (
	"github.com/pangxianfei/framework/helpers/zone"
	"github.com/pangxianfei/framework/queue/driver/nsq"
)

func Initialize() {
	//@todo使用不同的配置来新建不同的队列
	//@todo内存、nsq、rabbitmq
	setQueue(nsq.NewNsq("nsq"))
	initializeFailedProcessor()
}

var queue queuer

func Queue() queuer {
	return queue
}
func setQueue(q queuer) {
	queue = q
}

type queuer interface {
	producerer
	consumerer
	registerer
	SupportBroadCasting() bool // 对于在发出之前尚未启动的侦听器，它将不会接收事件。支持广播模式
	Close() (err error)
}

type registerer interface {
	Register(topicName string, channelName string) (err error)
	Unregister(topicName string, channelName string) (err error)
}

type producerer interface {
	Push(topicName string, channelName string, delay zone.Duration, body []byte) (err error)
}
type consumerer interface {
	//当handleerr不为零时，队列应发送REQ（重新队列），当为nil时，发送FIN（finish）
	Pop(topicName string, channelName string, handler func(hash string, body []byte) (handlerErr error), maxInFlight int) (err error)
}
