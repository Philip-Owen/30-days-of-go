package main

import (
	"fmt"
	"strings"
	"strconv"
)

// Raw input
var rawInput string = `Ethernet0,up,100000
Ethernet4,Down,100000
Ethernet8,up,25000

Ethernet12,   up,10000
Ethernet4,up,100000
Ethernet16,down,10000
Ethernet4,down,10000

Ethernet20,up,25000`

type SingleInterface struct {
	name string
	status string
	speed int
}

type Inventory struct {
	interfaces []SingleInterface
	interfacesMap map[string]SingleInterface
	totalValid int
	interfacesUp int
	interfacesDown int
	totalDuplicates int
	totalUnique int
	totalCapacity int
	duplicates map[string]SingleInterface
}

func dataNormalizer(input string) string {
	return strings.TrimSpace(input)
}

func speedConverter(speed int) string {
	return fmt.Sprintf("%dGb", (speed/1000))
}

// Check for conflicts between the duplicate interface records and unique interface records
func checkForRecordConflicts(duplicateInterfaces map[string]SingleInterface, uniqueInterfaces map[string]SingleInterface) {
	for _, dupRecord := range duplicateInterfaces {

		value, _ := uniqueInterfaces[dupRecord.name]

		if value != dupRecord {
			fmt.Printf("Conflict detected for %s\n", value.name)
			if value.speed != dupRecord.speed {
				fmt.Printf("     - Speed mismatch\n")
			}
			if value.status != dupRecord.status {
				fmt.Printf("     - Status mismatch\n")
			}
		}
	}
}

func parseLineToStruct(line string) (SingleInterface, bool) {
	// Split raw input
	rawInterface := strings.Split(line, ",")
	// Create SingleInterface struct instance
	var interfaceStruct SingleInterface
	// Create bool to track whether struct creation was correct
	var isSafeRecord bool = true

	// Check to make sure that rawInterface slice has correct length
	if len(rawInterface) == 3 {
		// Normalize raw input
		var name string = dataNormalizer(rawInterface[0])
		var status string = strings.ToLower(dataNormalizer(rawInterface[1]))
		speed, speedError := strconv.Atoi(dataNormalizer(rawInterface[2]))
		
		// If Atoi produces no error set values
		if speedError == nil {
			interfaceStruct.name = name
			interfaceStruct.status = status
			interfaceStruct.speed = speed
			
		} else {
			// if Atoi produces error set bool to false
			isSafeRecord = false
		}
	} else {
		// if rawInterface slice doesn't have correct length, set bool to false
		isSafeRecord = false
	}

	return interfaceStruct, isSafeRecord
}

func createInventory(rawInterfaces string) Inventory {
	parsedInterfaceData := strings.Split(rawInterfaces, "\n")
	var inventory Inventory
	inventory.duplicates = make(map[string]SingleInterface)
	inventory.interfacesMap = make(map[string]SingleInterface)


	for i := 0; i < len(parsedInterfaceData); i++ {
		// Check to make sure input is not a blank line
		if len(parsedInterfaceData[i]) > 0 {
			// Create interface struct and bool for validating struct creation
			singleInterface, isSafeRecord := parseLineToStruct(parsedInterfaceData[i])
			// Check to make sure struct was created properly
			if isSafeRecord {
				inventory.totalValid++
				// Check to see if duplicate record for interface name exists
				_, isDuplicate := inventory.interfacesMap[singleInterface.name]

				if isDuplicate {
					// Check to see if duplicate already exists with same interface name
					_, isDupofDup := inventory.duplicates[singleInterface.name]

					// If interface does not have a duplicate record, increment duplicate counter and add to duplicates map
					if !isDupofDup{
						inventory.totalDuplicates++	
						inventory.duplicates[singleInterface.name] = singleInterface
					}
				} else {
					// Increment unique interface counter, add current capacity to total, and add interface struct to inventory slice
					inventory.totalUnique++
					inventory.totalCapacity += singleInterface.speed
					inventory.interfacesMap[singleInterface.name] = singleInterface
					
					
					// Check if interface is up or down and increment appropriate counter
					if singleInterface.status == "up" {
						inventory.interfacesUp++
					} else if singleInterface.status == "down" {
						inventory.interfacesDown++
					}
				}
				inventory.interfaces = append(inventory.interfaces, singleInterface)
			}
		}
	}

	return inventory
}

func main()  {

	interfaceInventory := createInventory(rawInput)

	fmt.Println("Interface Inventory")
	fmt.Println("-------------------")

	for i := 0; i < len(interfaceInventory.interfaces); i++ {
		var interfaceName string = interfaceInventory.interfaces[i].name
		var interfaceStatus string = interfaceInventory.interfaces[i].status
		var interfaceSpeed string = speedConverter(interfaceInventory.interfaces[i].speed)

		fmt.Printf("%s %s %s\n", interfaceName, interfaceStatus, interfaceSpeed)
	}
	fmt.Println()
	fmt.Printf("Total valid interfaces: %d\n", interfaceInventory.totalValid)
	fmt.Printf("Total unique interface names: %d\n", interfaceInventory.totalUnique)
	fmt.Printf("Total duplicates found: %d\n", interfaceInventory.totalDuplicates)
	fmt.Printf("Total interfaces up: %d\n", interfaceInventory.interfacesUp)
	fmt.Printf("Total interfaces down: %d\n", interfaceInventory.interfacesDown)
	fmt.Printf("Total capacity: %s\n", speedConverter(interfaceInventory.totalCapacity))

	if interfaceInventory.totalDuplicates > 0 {
		fmt.Println()
		for _, dupRecord := range interfaceInventory.duplicates {
			fmt.Printf("Duplicates records found for interface: %s\n", dupRecord.name)
		}
		fmt.Println()
		checkForRecordConflicts(interfaceInventory.duplicates, interfaceInventory.interfacesMap)
	}
}
