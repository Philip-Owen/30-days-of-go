package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

var currentData string = `Ethernet0,up,100000
Ethernet4,up,100000
Ethernet8,down,25000
Ethernet12,up,10000
Ethernet16,up,10000`

// var newData string = `Ethernet0,up,100000
// Ethernet4,up,100000
// Ethernet8,down,25000
// Ethernet12,up,10000
// Ethernet16,up,10000`
var newData string = `Ethernet0,up,100000
Ethernet4,down,100000
Ethernet8,down,10000
Ethernet12,up,10000
Ethernet20,up,25000`

type StateTracker struct {
	current map[string]SingleInterface
	new map[string]SingleInterface
	recordAdded int 
	recordRemoved int
	recordChanged int
	recordUnchanged int
	stateChanged bool
	changes []ChangedRecord
}

type ChangedRecord struct {
	recordName string
	isChanged bool
	isStatusChanged bool
	isSpeedChanged bool
	isAdded bool
	isRemoved bool
	changedSpeed SpeedChanged
	changedStatus StatusChanged
}

type SpeedChanged struct {
	currentSpeed int
	newSpeed int
}

type StatusChanged struct {
	currentStatus string
	newStatus string
}

type SingleInterface struct {
	name string
	status string
	speed int
}

func speedConverter(speed int) int {
	return (speed/1000)
}

func trimWhitespace(input string) string {
	return strings.TrimSpace(input)
}

// Take raw line and convert to struct
func parseLineToStruct(rawLine string) SingleInterface {
	parsedInterface := strings.Split(rawLine, ",")

	name := trimWhitespace(parsedInterface[0])
	status := strings.ToLower(trimWhitespace(parsedInterface[1]))
	speed, _ := strconv.Atoi(trimWhitespace(parsedInterface[2]))

	return SingleInterface{name, status, speed}
}

// Take raw input and convert to map of struct
func buildStateMap(snapshot string) map[string]SingleInterface {
	parsedSnapshot := strings.Split(snapshot, "\n")
	stateMap := make(map[string]SingleInterface)

	for i := 0; i < len(parsedSnapshot); i++ {
		singleInterface := parseLineToStruct(parsedSnapshot[i])
		stateMap[singleInterface.name] = singleInterface
	}
	return stateMap
}

// Compare current state vs new state to find changed and removed records
func compareCurrentToNew(state *StateTracker)  {
	
	for key := range state.current {
		_, isInNewState := state.new[key]
		if isInNewState {
			if state.current[key] == state.new[key] {			
				state.recordUnchanged++
			} else {
				state.recordChanged++
				record := ChangedRecord{isChanged: true, recordName: key}

				if state.current[key].status != state.new[key].status {
					record.isStatusChanged = true
					status := StatusChanged{currentStatus: state.current[key].status, newStatus:  state.new[key].status}
					record.changedStatus = status
				}				

				if state.current[key].speed != state.new[key].speed {
					record.isSpeedChanged = true
					speed := SpeedChanged{currentSpeed: state.current[key].speed, newSpeed:  state.new[key].speed}
					record.changedSpeed = speed
				}
				
				state.changes = append(state.changes, record)
			}
		} else {
			state.recordRemoved++
			state.changes = append(state.changes, ChangedRecord{isRemoved: true, recordName: key})
		}
	}
}

// Compare new state vs current state to find added records
func compareNewtoCurrent(state *StateTracker)  {

	for key := range state.new {
		_, isInCurrentState := state.current[key]
		if !isInCurrentState {
			state.recordAdded++
			state.changes = append(state.changes, ChangedRecord{isAdded: true, recordName: key})
		}
	}
}

// Compare the two state maps to find changes
func compareStateMaps(state *StateTracker) {
	// Check if current state matches new state
	if reflect.DeepEqual(state.current, state.new) {
		fmt.Println("No changes detected")
	} else {
		state.stateChanged = true
		compareCurrentToNew(state)
		compareNewtoCurrent(state)
	}
}

func printChanges(changes []ChangedRecord)  {
	for i := 0; i < len(changes); i++ {
		if changes[i].isChanged {
			var changedMessage strings.Builder
			changedMessage.WriteString(fmt.Sprintf("%s changed:", changes[i].recordName))

			if changes[i].isStatusChanged {
				changedMessage.WriteString(fmt.Sprintf(" %s -> %s", changes[i].changedStatus.currentStatus, changes[i].changedStatus.newStatus))
			}

			if changes[i].isSpeedChanged {
				currentSpeed := speedConverter(changes[i].changedSpeed.currentSpeed)
				newSpeed := speedConverter(changes[i].changedSpeed.newSpeed)
				changedMessage.WriteString(fmt.Sprintf(" %dGb -> %dGb", currentSpeed, newSpeed))
			}
			fmt.Println(changedMessage.String())
		}

		if changes[i].isAdded {
			fmt.Printf("%s added\n", changes[i].recordName)
		}

		if changes[i].isRemoved {
			fmt.Printf("%s removed\n", changes[i].recordName)
		}
	}
}

func main()  {
	// Create StateTracker struct and build maps for the current and new state
	state := StateTracker{}
	state.stateChanged = false
	state.current = buildStateMap(currentData)
	state.new = buildStateMap(newData)

	// Create a pointer to StateTracker to reference in compareStateMaps function
	var statePtr *StateTracker
	statePtr = &state

	// Compare maps and create changed records
	compareStateMaps(statePtr)

	fmt.Println("Interface Change Report")
	fmt.Println("-----------------------")
	printChanges(state.changes)
	if state.stateChanged {
		fmt.Println()
		fmt.Println("Summary:")
		fmt.Printf("%d interfaces changed\n", state.recordChanged)
		fmt.Printf("%d added\n", state.recordAdded)
		fmt.Printf("%d removed\n", state.recordRemoved)
		fmt.Printf("%d unchanged\n", state.recordUnchanged)
	}
}