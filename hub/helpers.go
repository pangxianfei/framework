package hub

import (
	"errors"

	"github.com/golang/protobuf/proto"

	"github.com/pangxianfei/framework/queue"
)

func Emit(e Eventer) (errs map[ListenerName]error) {
	// 对于在发出之前尚未启动的侦听器，它将不会接收事件。支持广播模式
	if queue.Queue().SupportBroadCasting() {
		eventListenerList := eventListener(e)
		if len(eventListenerList) <= 0 {
			errs = make(map[ListenerName]error)
			errs["nil"] = errors.New("listener doesnt't exist")
			return errs
		}

		l := eventListenerList[0]
		if err := queue.NewProducer(topicName(e, l, queue.Queue().SupportBroadCasting), channelName(l), e.paramData(), l.Retries(), l.Delay()).Push(); err != nil {
			errs = make(map[ListenerName]error)
			errs[channelName(l)] = err
		}
		return errs
	}

	// 推送到多个侦听器
	for _, l := range eventListener(e) {
		if err := queue.NewProducer(topicName(e, l, queue.Queue().SupportBroadCasting), channelName(l), e.paramData(), l.Retries(), l.Delay()).Push(); err != nil {
			if errs == nil {
				errs = make(map[ListenerName]error)
			}

			errs[channelName(l)] = err
		}
	}
	return errs
}

func On(listenerName ListenerName) {

	listener := listenerMap[listenerName]
	if listener == nil {
		panic(errors.New("listener " + listenerName + " doesn't exist"))
	}

	for _, e := range listener.Subscribe() {
		err := queue.NewConsumer(topicName(e, listener, queue.Queue().SupportBroadCasting), channelName(listener), e.ParamProto(), func(paramPtr proto.Message) error {
			if err := listener.Construct(paramPtr); err != nil {
				return err
			}

			if err := listener.Handle(); err != nil {
				return err
			}

			return nil
		}).Pop()
		if err != nil {
			panic(err)
		}
	}
}
