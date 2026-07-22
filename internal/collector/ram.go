package collector

type RamStats struct {
	AppMemoryInUseMb        int
	WiredMemoryInUseMb      int
	CompressedMemoryInUseMb int
	FreeMemoryMb            int
	TopProcesses            []ProcessRamStats
}

type ProcessRamStats struct {
	Name         string
	UsagePercent float64
}
