package gowon

import (
	"encoding/json"
	"log"
	"regexp"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type messageHandler func(Message) (string, error)

type MessageRouter struct {
	Commands map[string]messageHandler
	Regexes  map[string]messageHandler
}

func NewMessageRouter() *MessageRouter {
	return &MessageRouter{
		Commands: make(map[string]messageHandler),
		Regexes:  make(map[string]messageHandler),
	}
}

func (mr MessageRouter) AddCommand(command string, mh messageHandler) {
	mr.Commands[command] = mh
}

func (mr MessageRouter) AddRegex(regex string, mh messageHandler) {
	mr.Regexes[regex] = mh
}

func (mr MessageRouter) Route(msg Message) (string, error) {
	handler, prs := mr.Commands[msg.Command]
	if prs {
		out, err := handler(msg)

		return out, err
	}

	for r, handler := range mr.Regexes {
		match, _ := regexp.MatchString(r, msg.Msg)
		if match {
			out, err := handler(msg)

			return out, err
		}
	}

	return "", nil
}

func (mr MessageRouter) SubscribeChannel(opts *mqtt.ClientOptions, module string, inTopic string, outTopic string) {
	oldOnConnect := opts.OnConnect

	opts.OnConnect = func(client mqtt.Client) {
		if oldOnConnect != nil {
			oldOnConnect(client)
		}

		client.Subscribe(inTopic, 0, func(client mqtt.Client, msg mqtt.Message) {

			ms, err := CreateMessageStruct(msg.Payload())
			if err != nil {
				log.Print(err)

				return
			}

			out, err := mr.Route(ms)
			if err != nil {
				log.Print(err)

				return
			}

			if out == "" {
				return
			}

			ms.Module = module
			ms.Msg = out
			mb, err := json.Marshal(ms)
			if err != nil {
				log.Print(err)

				return
			}
			client.Publish(outTopic, 0, false, mb)
		})

		log.Printf("Subscription to %s complete", inTopic)
	}
}

func (mr MessageRouter) Subscribe(opts *mqtt.ClientOptions, module string) {
	mr.SubscribeChannel(opts, module, "/gowon/input", "/gowon/output")
}

func (mr MessageRouter) SubscribeMiddleware(opts *mqtt.ClientOptions, module string) {
	mr.SubscribeChannel(opts, module, "/gowon/output", "/gowon/output")
}
