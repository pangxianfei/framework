package job

import (
	"github.com/pangxianfei/framework/helpers/log"
	"github.com/pangxianfei/framework/queue"
)


// tmaic do the broadcast it self, for compatible with the queue driver which doesn't support topic broadcasting
func topicName(j jobber) string {
	return "tmaic-" + j.Name()
}
func channelName(j jobber) string {
	return j.Name()
}
func RegisterQueue() {
	for _, j := range jobMap {
		log.Debug(topicName(j))
		if err := queue.Queue().Register(topicName(j), channelName(j)); err != nil {
			panic(err)
		}
	}
}
