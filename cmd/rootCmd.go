package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var broker string
var retainMessages bool
var topic string
var qos byte

var rootCmd = &cobra.Command{
	Use:   "Goqtt.exe [command]",
	Short: "Goqtt is a mqttt client writen in Go, that can operate in 2 different modes.",
	Long: `Goqtt is a mqttt client writen in Go, that can operate in 2 different modes:
				  iclient - Single interactive client to send messages
				  ltester - Load tester that runs multiple mqtts client at once`,
}

func Execute() {

	rootCmd.PersistentFlags().StringVarP(&broker, "broker", "b", "localhost:1883", "Mqtt broker address (and port)")
	rootCmd.PersistentFlags().StringVarP(&topic, "topic", "t", "ExampleTopic", "Topic to publish messages to")
	rootCmd.PersistentFlags().BoolVarP(&retainMessages, "retain", "r", false, "Retain messages")
	rootCmd.PersistentFlags().Uint8VarP(&qos, "qos", "q", 0, "Mqtt QoS")

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
