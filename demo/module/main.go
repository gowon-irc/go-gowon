package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/gowon-irc/go-gowon"
	"github.com/gowon-irc/gowon/pkg/message"
	"github.com/jessevdk/go-flags"
)

type Options struct {
	Prefix string `short:"P" long:"prefix" env:"GOWON_PREFIX" default:"." description:"prefix for commands"`
	Broker string `short:"b" long:"broker" env:"GOWON_BROKER" default:"localhost:1883" description:"mqtt broker"`
}

const mqttConnectRetryInternal = 5 * time.Second

func testHandler(m message.Message) (string, error) {
	return "testing", nil
}

func testHandler2(m message.Message) (string, error) {
	return "testing2", nil
}

func regexHandler(m message.Message) (string, error) {
	return fmt.Sprintf("{green}echoing: {red}%s{clear}", m.Msg), nil
}

func main() {
	opts := Options{}
	if _, err := flags.Parse(&opts); err != nil {
		log.Fatal(err)
	}

	mqttOpts := mqtt.NewClientOptions().AddBroker(fmt.Sprintf("tcp://%s", opts.Broker))
	mqttOpts.SetClientID("gowon_module")
	mqttOpts.SetConnectRetry(true)
	mqttOpts.SetConnectRetryInterval(mqttConnectRetryInternal)

	c := mqtt.NewClient(mqttOpts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	mr := gowon.NewMessageRouter()
	mr.AddCommand("test", testHandler)
	mr.AddCommand("test2", testHandler2)
	mr.AddRegex(".*hello.*", regexHandler)
	mr.Subscribe(c, "module")

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	<-sigs
}
