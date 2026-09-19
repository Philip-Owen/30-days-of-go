package main

import (
	"fmt"
	"strings"
	"strconv"
)

// Mock the raw data
var rawInput string = `Ethernet0,up,100000
Ethernet4,down,100000
Ethernet8,up,25000
Ethernet9,
Ethernet12,up,10000
Ethernet13,up,1000

Ethernet16, down,10000`

// Define interface struct
// example
// name: Ethernet0
// status: up
// speed: 10000
// malformed: false
// raw: Ethernet0,up,10000
type Interface struct {
	name string
	status string
	speed int
}

// Convert raw speed to be more readable
func speedConverter(speed int) string {
	return fmt.Sprintf("%dGb", (speed/1000))
}

// Strips any extra whitespace and normalizes string case for comparison
func normalizeStatusInfo(status string) string {
	return strings.TrimSpace(strings.ToLower(status))
}

// Convert the raw input into a slice of structs
func convertRawInterfacesToStructs(interfaceData []string) []Interface {
	var interfacesInventory []Interface

	for i := 0; i < len(interfaceData); i++ {
		// Check to make sure raw line is not empty
		if len(interfaceData[i]) > 0 {
			// Split raw line to gather interface data
			singleInterfaceRaw := strings.Split(interfaceData[i], ",")

			// Check whether the converted slice has the correct length
			if len(singleInterfaceRaw) == 3 {
				var interfaceName string = singleInterfaceRaw[0]
				var interfaceStatus string = singleInterfaceRaw[1]
				interfaceSpeed, interfaceSpeedError := strconv.Atoi(singleInterfaceRaw[2])

				if interfaceSpeedError == nil {
					interfacesInventory = append(interfacesInventory, Interface{interfaceName, interfaceStatus, interfaceSpeed})
				}
			}
		}
	}
	return interfacesInventory
}

func main()  {

	// Take raw input and split into slice of strings for use in converter function
	parsedInterfaceData := strings.Split(rawInput, "\n")
	
	// convert parsedInterfaceData into a slice of Interface structs
	inventory := convertRawInterfacesToStructs(parsedInterfaceData)

	// Define counters variables
	var interfacesCounter int
	var interfacesUp int
	var interfacesDown int
	var capacityCounter int
	var downInterfacesList []string
	var downInterfaces string = "Interfaces down: "


	fmt.Println("Interface Inventory Report")
	fmt.Println("--------------------------")
	for i := 0; i < len(inventory); i++ {
		// Get data from item in inventory to be used in Interface struct
		var name string = inventory[i].name
		var status string = inventory[i].status
		var speed int = inventory[i].speed

		interfacesCounter++
		capacityCounter += speed
		var statusNormalized = normalizeStatusInfo(status)

		if statusNormalized == "up" {
			interfacesUp++
		} else if statusNormalized == "down" {
			interfacesDown++
			downInterfacesList = append(downInterfacesList, name)
		}

		fmt.Printf("%s %s %s\n", name, status, speedConverter(speed))
	}
	
	// Convert Mbps to Gbps
	var totalCapacity string = speedConverter(capacityCounter)

	// Print counters information
	fmt.Println()
	fmt.Printf("Total of valid Interfaces: %d\n", interfacesCounter)	
	fmt.Printf("Interfaces up: %d\n", interfacesUp)
	fmt.Printf("Interfaces down: %d\n", interfacesDown)
	fmt.Printf("Total bandwidth: %s\n", totalCapacity)

	// Check if there are down interfaces. if 1 or more down interfaces print list
	if interfacesDown > 0 {
		for i := 0; i < len(downInterfacesList); i++ {
			if i != (len(downInterfacesList) - 1) {
				downInterfaces += fmt.Sprintf("%s, ", downInterfacesList[i])
			} else {
				downInterfaces += fmt.Sprintf("%s", downInterfacesList[i])
			}
		}
		fmt.Println(downInterfaces)
	}
}