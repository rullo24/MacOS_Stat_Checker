package collector

import "MacOS_Stat_Checker/internal/capture"

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
	TimeToDischargeMinutes float64
	AvgTimeToFullMinutes   float64
	AmperageMilliamps      int
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
	default:
	}

	return BatteryStats{
		LevelPercent:           (float64(battery_raw.CurrentCapacity) / float64(battery_raw.MaxCapacity)) * 100.0,
		CycleCount:             battery_raw.CycleCount,
		ChargingStatus:         l_charging_status,
		TimeToDischargeMinutes: float64(battery_raw.TimeRemaining),
		AvgTimeToFullMinutes:   float64(battery_raw.AvgTimeToFull),
		AmperageMilliamps:      int(battery_raw.AmperageMilliamps),
		VoltageMillivolts:      battery_raw.VoltageMillivolts,
		TemperatureDegC:        float64(battery_raw.TemperatureCentiC) / 100,
	}, nil // TBC
}
