package main

import (
	"MacOS_Stat_Checker/internal/collector"
	"fmt"
	"time"
)

func main() {

	// running battery collector
	batteryChan := make(chan collector.BatteryStats)
	stopChan := make(chan bool)
	go collector.RunSampler(
		time.Second,
		collector.CollectBatteryStats,
		batteryChan,
		stopChan,
	)

	// inf iterate
	var batteryValues collector.BatteryStats
	for {
		batteryValues = <-batteryChan
		fmt.Println(batteryValues)
	}
}
