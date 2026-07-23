package capture

// C includes and defs
/*
#cgo LDFLAGS: -framework IOKit -framework CoreFoundation
#include "battery_darwin.h"
*/

import (
	"C"
	"errors"
)

type BatteryRaw struct {
	CurrentCapacity   int
	MaxCapacity       int
	DesignCapacity    int
	NominalCapacity   int
	CycleCount        int
	TimeRemaining     int
	AvgTimeToFull     int
	AmperageMA        int32
	VoltageMV         int
	TemperatureCentiC int
	IsCharging        bool
	IsPluggedIn       bool
	IsFullyCharged    bool
}

func CollectBatteryRaw() (BatteryRaw, error) {
	var raw C.battery_raw_t // pulled from battery_darwin.h

	// capturing battery stats in raw
	if C.collect_battery_raw(&raw) != 0 { // failure
		return BatteryRaw{}, errors.New("no battery present")
	}

	return BatteryRaw{
		CurrentCapacity:   int(raw.current_capacity),
		MaxCapacity:       int(raw.max_capacity),
		DesignCapacity:    int(raw.design_capacity),
		NominalCapacity:   int(raw.nominal_capacity),
		CycleCount:        int(raw.cycle_count),
		TimeRemaining:     int(raw.time_remaining),
		AvgTimeToFull:     int(raw.avg_time_to_full),
		AmperageMA:        int32(raw.amperage_ma),
		VoltageMV:         int(raw.voltage_mv),
		TemperatureCentiC: int(raw.temperature_centic),
		IsCharging:        raw.is_charging != 0,
		IsPluggedIn:       raw.is_plugged_in != 0,
		IsFullyCharged:    raw.is_fully_charged != 0,
	}, nil

}
