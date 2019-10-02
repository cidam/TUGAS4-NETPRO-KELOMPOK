package main

import (
	"fmt"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

var f MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("Topic: %s\n", msg.Topic())
	fmt.Printf("Message: %s\n", msg.Payload())
}

func main() {

	opts := MQTT.NewClientOptions().AddBroker("localhost:1883")
	opts.SetClientID("publisher")
	opts.SetDefaultPublishHandler(f)

	c := MQTT.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	msg := fmt.Sprintf("Tugas 4 Network Programming")
	token := c.Publish("TUGAS 4", 0, false, msg)
	token.Wait()

	c.Disconnect(250)
}
