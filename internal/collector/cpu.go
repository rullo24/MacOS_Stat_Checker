package collector

import "MacOS_Stat_Checker/internal/capture"

type CpuStaticStats struct {
	Name                string
	NumPerformanceCores int
	NumEfficiencyCores  int
}

type CpuDynamicStats struct {
	UserTicks   []uint64
	SystemTicks []uint64
	IdleTicks   []uint64
	NiceTicks   []uint64
}

func CollectCpuStaticStats() (CpuStaticStats, error) {
	cpuStaticRaw, err := capture.CollectCpuStaticRaw()
	if err != nil {
		return CpuStaticStats{}, err
	}

	return CpuStaticStats{
		Name:                cpuStaticRaw.Name,
		NumPerformanceCores: cpuStaticRaw.PerfCoreCount,
		NumEfficiencyCores:  cpuStaticRaw.EfficiencyCoreCount,
	}, nil
}

func CollectCpuDynamicStats(coreCount int) (CpuDynamicStats, error) {
	cpuDynamicRaw, err := capture.CollectCpuDynamicRaw(coreCount)
	if err != nil {
		return CpuDynamicStats{}, err
	}

	return CpuDynamicStats{
		UserTicks:   cpuDynamicRaw.UserTicks,
		SystemTicks: cpuDynamicRaw.SystemTicks,
		IdleTicks:   cpuDynamicRaw.IdleTicks,
		NiceTicks:   cpuDynamicRaw.NiceTicks,
	}, nil
}
