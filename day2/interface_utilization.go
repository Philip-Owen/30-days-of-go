package main

import (
	"fmt"
)

func calculateUtilization(interfaceCapacity float32, currentUtilization float32) float32 {
	return (currentUtilization / interfaceCapacity) * 100
}

func main()  {
	// Interface utilization data representation
	var interfaces = [5]string{"Ethernet0", "Ethernet4", "Ethernet8", "Ethernet12", "Ethernet16"}
	var capacity = [5]float32{100, 100, 25, 10, 10}
	var utilization = [5]float32{42, 91, 18, 2, 9.7}

	// Set baseline thresholds
	// Normal below 70%, Warning 70% - 90%, Critical > 90%
	const normalThreshold float32 = 70
	const criticalThreshold float32 = 90

	// Track highest utilization
	var highestUtilizationPercentage float32
	var highestUtilizationInterface string

	// Track status counts
	var normalCount int
	var warningCount int
	var criticalCount int


	fmt.Println("Interface Utilization Report")
	fmt.Println("----------------------------")
	for i := 0; i < len(interfaces); i++ {
		var thresholdStatus string
		var utilizationPercentage float32 = calculateUtilization(capacity[i], utilization[i])

		// Set threshold status based on utilization compared to baseline threshold
		if utilizationPercentage < normalThreshold {
			thresholdStatus = "NORMAL"
			normalCount++
		} else if utilizationPercentage < criticalThreshold {
			thresholdStatus = "WARNING"
			warningCount++
		} else {
			thresholdStatus = "CRITICAL"
			criticalCount++
		}


		fmt.Printf("%s %.1f %% %s \n", interfaces[i], utilizationPercentage, thresholdStatus)

		// Compare utilization to find the highest
		if utilizationPercentage > highestUtilizationPercentage {
			highestUtilizationPercentage = utilizationPercentage
			highestUtilizationInterface = interfaces[i]
		}
	}
	fmt.Println()
	fmt.Printf("Number of interfaces checked: %d\n", len(interfaces))
	fmt.Printf("%d NORMAL\n", normalCount)
	fmt.Printf("%d WARNING\n", warningCount)
	fmt.Printf("%d CRITICAL\n", criticalCount)
	fmt.Println()
	fmt.Printf("The interface with the highest utilization is %s at %.1f %%\n", highestUtilizationInterface, highestUtilizationPercentage)
}