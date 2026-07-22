package collector

type ChargingStatus string

const (
	StatusDischarging          ChargingStatus = "discharging"
	StatusCharging             ChargingStatus = "charging"
	StatusFullyCharged         ChargingStatus = "fullyCharged"
	StatusPluggedInNotCharging ChargingStatus = "pluggedInNotCharging"
)

type BatteryStats struct {
	LevelPercent           float64
	ChargingStatus         ChargingStatus
	TimeToDischargeMinutes float64
	TimeToFullMinutes      float64
	CycleCount             int
	HealthPercent          float64
	AmperageMilliamps      int
	Voltage                float64
	TemperatureDegC        float64
}

type ProcessBatteryStats struct {
	Name         string
	UsagePercent float64
}
