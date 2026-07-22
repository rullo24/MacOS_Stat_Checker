package collector

const (
	CHARGER_NOT_CONNECTED = 1
	CHARGER_CONNECTED     = 2
)

type BatteryStats struct {
	LevelPercent      float64
	TimeToDischargeMs float64
	HealthPercent     float64
	AmperageMilliamps int
	Voltage           float64
	TemperatureDegC   float64
	ChargingStatus    int
}

type ProcessBatteryStats struct {
	Name         string
	UsagePercent float64
}
