package main

import (
	"fmt"
	"os"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

var f MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("Topic: %s\n", msg.Topic())
	fmt.Printf("Message: %s\n", msg.Payload())
}

func main() {

	opts := MQTT.NewClientOptions().AddBroker("localhost:1883")
	opts.SetClientID("subscriber")
	opts.SetDefaultPublishHandler(f)

	c := MQTT.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	if token := c.Subscribe("TUGAS 4", 0, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	time.Sleep(60 * time.Second)

	// Unsubscribe
	if token := c.Unsubscribe("TUGAS 4"); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	c.Disconnect(250)
}
