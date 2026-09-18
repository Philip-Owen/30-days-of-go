package main

import (
	"fmt"
	"strings"
)

var rawInput string = `Ethernet0,up,100000
Ethernet4,down,100000

Ethernet5,up
Ethernet8,Up ,25000
Ethernet12,up,10000
Ethernet16, down,10000`

func normalizeStatusInfo(status string) string {
	return strings.TrimSpace(strings.ToLower(status))
}

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

func main()  {
	// Split raw input at newline
	var interfaceData = strings.Split(rawInput, "\n")

	// Create variables for holding up/down counters
	var interfacesCounter int
	var interfacesUp int
	var interfacesDown int

	fmt.Println("Interface Summary Report")
	fmt.Println("------------------------")
	for i := 0; i < len(interfaceData); i++ {

		// Check to make sure line is not empty
		if len(interfaceData[i]) > 0 {

			var singleInterface = strings.Split(interfaceData[i], ",")

			// Expected slice length is 3, if it's not declare the line as malformed
			if len(singleInterface) == 3 {
				interfacesCounter++
				var statusNormalized string = normalizeStatusInfo(singleInterface[1])
				fmt.Printf("%s   %s   %s\n", singleInterface[0], statusNormalized, speedConverter(singleInterface[2]))

				if statusNormalized == "up" {
					interfacesUp++
				} else if statusNormalized == "down" {
					interfacesDown++
				}
			} else {
				fmt.Printf("** Skipping malformed line. Does not contain the correct amount of fields: '%s' **\n", interfaceData[i])
			}
		}
	}

	// Add space between interface data and counters
	fmt.Println()
	fmt.Printf("Number of Interfaces checked: %d\n", interfacesCounter)
	fmt.Printf("Interfaces up: %d\n", interfacesUp)
	fmt.Printf("Interfaces down: %d\n", interfacesDown)
}