package collector

import "MacOS_Stat_Checker/internal/capture"

type CpuStaticStats struct {
	Name                string
	NumPerformanceCores int
	NumEfficiencyCores  int
}

type CpuDynamicStats struct {
	SystemUsagePercent []float64
	UserUsagePercent   []float64
	IdleUsagePercent   []float64
	NiceUsagePercent   []float64
}

func CollectCpuStaticStats() (CpuStaticStats, error) {
	cpuStaticRaw, err := capture.CollectCpuStaticRaw()
	if err != nil {
		return CpuStaticStats{}, err
	}

	return CpuStaticStats{
		Name:                cpuStaticRaw.name,
		NumPerformanceCores: cpuStaticRaw.perfCoreCount,
		NumEfficiencyCores:  cpuStaticRaw.efficiencyCoreCount,
	}, nil
}

func CollectCpuDynamicStats(coreCount int) (CpuDynamicStats, error) {
	cpuDynamicRaw, err := capture.CollectCpuDynamicRaw(coreCount)
	if err != nil {
		return CpuDynamicStats{}, err
	}

	return
}
