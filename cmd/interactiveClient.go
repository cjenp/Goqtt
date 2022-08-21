package cmd

import (
	"fmt"

	"github.com/cjenp/Goqtt/Goqtt"
	Mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/spf13/cobra"
)

var clientId string
var username string
var password string

func init() {
	interactiveClientCommand.Flags().StringVar(&clientId, "id", "device1", "Mqtt client ClientID")
	interactiveClientCommand.Flags().StringVarP(&username, "username", "u", "device1", "Mqtt client username")
	interactiveClientCommand.Flags().StringVarP(&password, "password", "p", "password", "Mqtt client password")
	rootCmd.AddCommand(interactiveClientCommand)
}

var interactiveClientCommand = &cobra.Command{
	Use:   "iclient",
	Short: "Run a single mqtt client, that enables manual input of messages to send",
	Long:  "Run a single mqtt client, that enables manual input of messages to send",
	Run: func(cmd *cobra.Command, args []string) {
		runDevice()
	},
}

func runDevice() {
	opts := Mqtt.NewClientOptions()
	opts.AddBroker(broker)
	opts.SetClientID(clientId)
	opts.SetUsername(username)
	opts.SetPassword(password)
	opts.SetCleanSession(true)

	cli := Goqtt.NewGoqttClient(topic, qos, retainMessages, *opts)
	cli.Connect(opts)
	fmt.Println("Client connected. Type 'exit' to quit or type message to be sent!")
	for {
		fmt.Print("Enter message: ")
		var inputMessage string
		fmt.Scanln(&inputMessage)

		if inputMessage == "exit" {
			break
		}

		cli.Publish(inputMessage)
	}
	cli.Disconnect()
	fmt.Println("Client disconnected.")
}
