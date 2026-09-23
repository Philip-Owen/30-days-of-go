package main

import (
	"fmt"
	"maps"
	"slices"
	"strconv"
	"strings"
)

var desiredState string = `Ethernet0,up,100000
Ethernet4,up,100000
Ethernet8,up,25000
Ethernet12,down,10000
Ethernet16,up,10000`

var actualState string = `Ethernet0,up,100000
Ethernet4,down,100000
Ethernet8,up,10000
Ethernet12,down,10000
Ethernet20,up,25000`

type State struct {
	desired map[string]SingleInterface
	current map[string]SingleInterface
}

type SingleInterface struct {
	name string
	status string
	speed int
}

type ComplianceRecord struct {
	name string
	message string
}

type ComplianceStatus struct {
	compliant int
	statusMismatch int
	speedMismatch int
	multipleMismatch int
	missing int
	unexpected int
}

type Compliance struct {
	results map[int]ComplianceRecord
	complianceStatus ComplianceStatus
}

func normalizeInput(input string) string {
	return strings.TrimSpace(input)
}

func speedConverter(speed int) int {
	return (speed/1000)
}

func getNumberFromName(name string) int {
	getEthernetNumber := strings.Split(name, "Ethernet")
	indexNumber, _ := strconv.Atoi(getEthernetNumber[1])

	return indexNumber
}

func getCompliancePercentage(status ComplianceStatus) float64 {
	desired := status.compliant + status.speedMismatch + status.statusMismatch + status.multipleMismatch + status.missing
	desiredFloat := float64(desired)
	compliant := float64(status.compliant)

	total := (compliant/desiredFloat) * 100

	return total
}

func parseLineToStruct(line string) SingleInterface {
	parsedInterface := strings.Split(line, ",")

	name := normalizeInput(parsedInterface[0])
	status := normalizeInput(strings.ToLower(parsedInterface[1]))
	speed, _ := strconv.Atoi(normalizeInput(parsedInterface[2]))

	return SingleInterface{name, status, speed}
}

func buildStateMap(data string) map[string]SingleInterface {
	parsedInput := strings.Split(data, "\n")
	stateMap := make(map[string]SingleInterface)

	for i := 0; i < len(parsedInput); i++ {
		stateRecord := parseLineToStruct(parsedInput[i])
		stateMap[stateRecord.name] = stateRecord
	}
	return stateMap
}

func complianceRecordBuilder(key string, message string) ComplianceRecord {
	return ComplianceRecord{key, message}
}

func checkCompliance(state State) Compliance {
	var compliance = Compliance{}
	compliance.results = make(map[int]ComplianceRecord)
	
	for key, value := range state.current {
		// Get interface number and convert to int to use as key to keep output sorted consistently
		indexNumber := getNumberFromName(key)
		
		_, exists := state.desired[key]

		if exists {
			if value == state.desired[key] {				
				compliance.complianceStatus.compliant++
				compliance.results[indexNumber] = complianceRecordBuilder(key, "COMPLIANT")
			} else {

				if value.status != state.desired[key].status && value.speed != state.desired[key].speed {
					compliance.complianceStatus.multipleMismatch++
					categoryMessage := fmt.Sprintf("MULTIPLE_MISMATCH expected=%s actual=%s expected=%d actual=%d", state.desired[key].status, value.status, speedConverter(state.desired[key].speed), speedConverter(value.speed))
					compliance.results[indexNumber] = complianceRecordBuilder(key, categoryMessage)
				} else {
					if value.status != state.desired[key].status {
						compliance.complianceStatus.statusMismatch++
						categoryMessage := fmt.Sprintf("STATUS_MISMATCH expected=%s actual=%s", state.desired[key].status, value.status)
						compliance.results[indexNumber] = complianceRecordBuilder(key, categoryMessage)
					}
					if value.speed != state.desired[key].speed {
						compliance.complianceStatus.speedMismatch++
						categoryMessage := fmt.Sprintf("SPEED_MISMATCH expected=%dGb actual=%dGb", speedConverter(state.desired[key].speed), speedConverter(value.speed))
						compliance.results[indexNumber] = complianceRecordBuilder(key, categoryMessage)
					}
				}
			}

		} else {
			compliance.complianceStatus.unexpected++
			compliance.results[indexNumber] = complianceRecordBuilder(key, "UNEXPECTED")
		}
	}

	for key := range state.desired {
		// Get interface number and convert to int to use as key to keep output sorted consistently
		indexNumber := getNumberFromName(key)
		
		_, exists := state.current[key]

		if !exists {
			compliance.complianceStatus.missing++
			compliance.results[indexNumber] = complianceRecordBuilder(key, "MISSING")
		}
	}

	return compliance
}

func compliancePresentationBuilder(compliance Compliance)  {

	keys := slices.Collect(maps.Keys(compliance.results))

	slices.Sort(keys)

	fmt.Println("Interface Compliance Report")
	fmt.Println("---------------------------")
	for i := 0; i < len(keys); i++ {
		fmt.Printf("%s %s\n", compliance.results[keys[i]].name, compliance.results[keys[i]].message)
	}
	fmt.Println()
	fmt.Printf("%d COMPLIANT\n", compliance.complianceStatus.compliant)
	fmt.Printf("%d STATUS_MISMATCH\n", compliance.complianceStatus.statusMismatch)
	fmt.Printf("%d SPEED_MISMATCH\n", compliance.complianceStatus.speedMismatch)
	fmt.Printf("%d MULTIPLE_MISMATCH\n", compliance.complianceStatus.multipleMismatch )
	fmt.Printf("%d MISSING\n", compliance.complianceStatus.missing)
	fmt.Printf("%d UNEXPECTED\n", compliance.complianceStatus.unexpected)
	fmt.Println()
	if len(compliance.results) == compliance.complianceStatus.compliant {
		fmt.Println("COMPLIANCE CHECK: PASS")
	} else {
		fmt.Println("COMPLIANCE CHECK: FAIL")
	}

	fmt.Printf("Compliance: %.1f%%\n", getCompliancePercentage(compliance.complianceStatus))
}


func main()  {
	state := State{}
	state.desired = buildStateMap(desiredState)
	state.current = buildStateMap(actualState)

	compliance := checkCompliance(state)
	compliancePresentationBuilder(compliance)
}