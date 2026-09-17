package main

import (
	"fmt"
)

// Convert: 
// Ethernet0  up    100000
// Ethernet4  down  100000
// Ethernet8  up    25000
// Ethernet12 up    10000
// 
// To:
// Interface Status Report
// -----------------------
// Ethernet0  : UP   - 100G
// Ethernet4  : DOWN - 100G
// Ethernet8  : UP   - 25G
// Ethernet12 : UP   - 10G

// 4 interfaces checked
// 3 up
// 1 down

func speedConverter(speed string) string {
	var convertedSpeed string

	switch speed {
	case "10000":
		convertedSpeed = "10Gb"
	case "25000":
		convertedSpeed = "25Gb"
	case "100000":
		convertedSpeed = "100Gb"
	}

	return convertedSpeed
}

func main() {

	var statusUp int
	var statusDown int

	var interfaces = [4]string{"Ethernet0", "Ethernet4", "Ethernet8", "Ethernet12"}
	var status = [4]string{"Up", "Down", "Up", "Up"}
	var speed = [4]string{"100000", "100000", "25000", "10000"}


	fmt.Println("Interface Status Report")
	fmt.Println("-----------------------")
	for i := 0; i < len(interfaces); i++ {
		if status[i] == "Up" {
			statusUp++
			fmt.Println(interfaces[i], ":", status[i], "-", speedConverter(speed[i]))
		} else {
			statusDown++
			fmt.Println(interfaces[i], ":", status[i], "-", speedConverter(speed[i]), "[WARNING]")
		}

	}
	fmt.Println()
	fmt.Println(len(interfaces), "Interfaces checked.")
	fmt.Println(statusUp, "up")
	fmt.Println(statusDown, "down")
}
