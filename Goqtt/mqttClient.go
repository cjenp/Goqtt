// This package is a wrapper for the "github.com/eclipse/paho.mqtt.golang" Mqtt client.
// Its main purpose is to simulate a mqtt device that has unpredicted behaviour.
package Goqtt

import (
	"math/rand"
	"time"

	Mqtt "github.com/eclipse/paho.mqtt.golang"
)

type GoqttWrapper struct {
	topic          string
	qos            byte
	retainMessages bool

	Mqtt.ClientOptions
	Mqtt.Client
}

func NewGoqttClient(topic string, qos byte, retainMessages bool, clientOptions Mqtt.ClientOptions) *GoqttWrapper {
	return &GoqttWrapper{topic, qos, retainMessages, clientOptions, Mqtt.NewClient(&clientOptions)}
}

// Connect the mqtt client
func (goqttWrapper *GoqttWrapper) Connect(options *Mqtt.ClientOptions) {

	if token := goqttWrapper.Client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
}

// Publish a message to the mqtt topic
func (goqttWrapper *GoqttWrapper) Publish(message string) {

	if !goqttWrapper.Client.IsConnectionOpen() {
		panic("Client not connected!")
	}

	token := goqttWrapper.Client.Publish(goqttWrapper.topic, goqttWrapper.qos, goqttWrapper.retainMessages, message)
	token.Wait()
}

// Disconnect the mqtt client
func (goqttWrapper *GoqttWrapper) Disconnect() {

	if !goqttWrapper.Client.IsConnectionOpen() {
		return
	}

	goqttWrapper.Client.Disconnect(250)
}

// Start publishing messages over mqtt. There is a delay between each message contolled by the "delayBetweenMessages" argument.
// Before each message is published there is a 10% change we will simulate a "reconnect" event.
// The function runtime is controller by the "runFor" argument.
func (goqttWrapper *GoqttWrapper) RunErratic(delayBetweenMessages int, runFor int) {

	for start := time.Now(); time.Since(start) < time.Duration(runFor)*time.Second; {

		if rand.Float32() < 0.1 {
			goqttWrapper.Disconnect()
			time.Sleep(time.Second * time.Duration(rand.Intn(5)+1))
			goqttWrapper.Connect(&goqttWrapper.ClientOptions)
			continue
		}

		goqttWrapper.Publish("{'1':'12', '2':'55', '3':'32'}")
		time.Sleep(time.Duration(delayBetweenMessages) * time.Second)
	}

}
