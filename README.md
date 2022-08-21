# Goqtt

Goqtt is a simple wrapper for the https://github.com/eclipse/paho.mqtt.golang MQTT Client library.
It can run in two different modes:
 - Interactive client<br/>
      &nbsp;Example: `go run main.go iclient` 
 - Load testing       
      &nbsp;Example: `go run main.go ltester --path devices.json --messageDely 10 --runFor 60`
 
