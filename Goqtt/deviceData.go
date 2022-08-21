package Goqtt

import (
	"bufio"
	"encoding/json"
	"os"
	"strings"
)

type DeviceData struct {
	ClientId string
	Username string
	Password string
}

func LoadDeviceData(jsonData string) []DeviceData {

	var deviceData []DeviceData
	json.Unmarshal([]byte(jsonData), &deviceData)
	return deviceData
}

func ReadFile(filePath string) string {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var sb strings.Builder
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sb.WriteString(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return sb.String()
}
