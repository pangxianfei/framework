package hub

import "github.com/pangxianfei/framework/queue"

// 自己做广播，与不支持主题广播的队列驱动程序兼容
func topicName(e Eventer, l Listener, supportBroadCasting func() bool) string {
	if supportBroadCasting() {
		return "event-" + EventName(e)
	}
	return "event-" + EventName(e) + "-" + channelName(l)
}


func channelName(l Listener) string {
	return l.Name()
}


func RegisterQueue() {
	for e, llist := range hub {
		for _, l := range llist {
			if err := queue.Queue().Register(topicName(event(e), l, queue.Queue().SupportBroadCasting), channelName(l)); err != nil {
				panic(err)
			}
		}
	}
}
