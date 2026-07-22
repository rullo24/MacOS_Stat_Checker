package collector

type CPUStats struct {
	Name               string
	TemperatureDegC    float64
	SystemUsagePercent float64
	UserUsagePercent   float64
	IdleUsagePercent   float64
	TopProcesses       []ProcessCPUStats
}

type ProcessCPUStats struct {
	ProcessName  string
	UsagePercent float64
}
