package collector

import "MacOS_Stat_Checker/internal/capture"

type ChargingStatus string

const (
	StatusDischarging          ChargingStatus = "discharging"
	StatusCharging             ChargingStatus = "charging"
	StatusFullyCharged         ChargingStatus = "fullyCharged"
	StatusPluggedInNotCharging ChargingStatus = "pluggedInNotCharging"
)

type BatteryStats struct {
	LevelPercent           float64
	CycleCount             int
	ChargingStatus         ChargingStatus
	TimeToDischargeMinutes float64
	AvgTimeToFullMinutes   float64
	AmperageMilliamps      int
	VoltageMillivolts      int
	TemperatureDegC        float64
	TopProcesses           []ProcessBatteryStats
}

type ProcessBatteryStats struct {
	Name         string
	UsagePercent float64
}

// @brief 	captures the relevant battery stats using Golang (calls into C funcs)
// @return	sends back the Golang-native BatteryStats and an error (if failed)
func CollectBattery() (BatteryStats, error) {

	battery_raw, err := capture.CollectBatteryRaw()
	if err != nil {
		return BatteryStats{}, err
	}

	// perform calcs on battery_raw values

	return BatteryStats{}, nil // TBC
}
