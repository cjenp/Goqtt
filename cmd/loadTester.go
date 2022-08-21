package cmd

import (
	"fmt"
	"sync"

	"github.com/cjenp/Goqtt/Goqtt"
	Mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/spf13/cobra"
)

var deviceDetailsPath string
var delayBetweenMessages int
var runFor int

func init() {
	ltesterCmd.Flags().IntVar(&delayBetweenMessages, "messageDely", 5, "Delay between two messages")
	ltesterCmd.Flags().IntVar(&runFor, "runFor", 600, "For how many seconds to run the mqtt client")
	ltesterCmd.Flags().StringVar(&deviceDetailsPath, "path", "devices.json", "Path to a JSON file with device details")
	ltesterCmd.MarkFlagRequired("path")
	rootCmd.AddCommand(ltesterCmd)
}

var ltesterCmd = &cobra.Command{
	Use:   "ltester",
	Short: "Run multiple mqtt client that send messages on a time interval",
	Long:  "Run multiple mqtt client that send messages on a time interval",
	Run: func(cmd *cobra.Command, args []string) {
		runLoadTest()
	},
}

func runLoadTest() {
	result := Goqtt.ReadFile(deviceDetailsPath)
	devices := Goqtt.LoadDeviceData(result)

	fmt.Println("Starting load tester")
	var waitGroup sync.WaitGroup
	for _, device := range devices {
		waitGroup.Add(1)
		go setupAndRunDevice(device, &waitGroup)
	}
	waitGroup.Wait()
	fmt.Println("Load test finished")
}

func setupAndRunDevice(device Goqtt.DeviceData, waitGroup *sync.WaitGroup) {
	opts := Mqtt.NewClientOptions()
	opts.AddBroker(broker)
	opts.SetClientID(device.ClientId)
	opts.SetUsername(device.Username)
	opts.SetPassword(device.Password)
	opts.SetCleanSession(true)

	cli := Goqtt.NewGoqttClient(topic, qos, retainMessages, *opts)
	cli.Connect(opts)
	cli.RunErratic(delayBetweenMessages, runFor)
	cli.Disconnect()
	waitGroup.Done()
}
