package collector

import (
	"MacOS_Stat_Checker/internal/capture"
	"math"
)

type ChargingStatus string

const (
	BatteryStatusDischarging          ChargingStatus = "discharging"
	BatteryStatusCharging             ChargingStatus = "charging"
	BatteryStatusPluggedInNotCharging ChargingStatus = "pluggedInNotCharging"
)

type BatteryStats struct {
	LevelPercent           float64
	CycleCount             int
	ChargingStatus         ChargingStatus
	TimeToDischargeMinutes int
	AvgTimeToFullMinutes   int
	AmperageMilliamps      int32
	VoltageMillivolts      int
	TemperatureDegC        float64
}

// @brief 	captures the relevant battery stats using Golang (calls into C funcs)
// @return	sends back the Golang-native BatteryStats and an error (if failed)
func CollectBatteryStats() (BatteryStats, error) {

	battery_raw, err := capture.CollectBatteryRaw()
	if err != nil {
		return BatteryStats{}, err
	}

	// perform calcs on battery_raw values
	var l_charging_status ChargingStatus
	switch {
	case battery_raw.IsCharging:
		l_charging_status = BatteryStatusCharging
	case battery_raw.IsPluggedIn:
		l_charging_status = BatteryStatusPluggedInNotCharging
	case !battery_raw.IsCharging:
		l_charging_status = BatteryStatusDischarging
	}

	// terniary ops
	var l_time_to_discharge int = battery_raw.TimeRemaining
	var l_avg_time_to_full int = battery_raw.AvgTimeToFull
	if l_time_to_discharge == math.MaxUint16 {
		l_time_to_discharge = 0 // sentinel N/A
	}
	if l_avg_time_to_full == math.MaxUint16 {
		l_avg_time_to_full = 0 // sentinel N/A
	}

	return BatteryStats{
		LevelPercent:           (float64(battery_raw.CurrentCapacity) / float64(battery_raw.MaxCapacity)) * 100.0,
		CycleCount:             battery_raw.CycleCount,
		ChargingStatus:         l_charging_status,
		TimeToDischargeMinutes: l_time_to_discharge,
		AvgTimeToFullMinutes:   l_avg_time_to_full,
		AmperageMilliamps:      battery_raw.AmperageMilliamps,
		VoltageMillivolts:      battery_raw.VoltageMillivolts,
		TemperatureDegC:        float64(battery_raw.TemperatureCentiC) / 100,
	}, nil
}
