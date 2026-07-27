package capture

/*
#cgo LDFLAGS: -framework IOKit -framework CoreFoundation
#include "cpu_darwin.h"
*/
import "C"
import "errors"

type CpuStaticRaw struct {
	Name                string
	PerfCoreCount       int
	EfficiencyCoreCount int
}

type CpuDynRaw struct {
	UserTicks   []uint64
	SystemTicks []uint64
	IdleTicks   []uint64
	NiceTicks   []uint64
}

func CollectCpuStaticRaw() (CpuStaticRaw, error) {
	var raw_static C.cpu_static_raw_t
	if C.collect_cpu_static_raw(&raw_static) != 0 {
		return CpuStaticRaw{}, errors.New("couldn't collect static CPU info")
	}

	return CpuStaticRaw{
		Name:                C.GoString(&raw_static.name[0]),
		PerfCoreCount:       int(raw_static.perf_core_count),
		EfficiencyCoreCount: int(raw_static.efficiency_core_count),
	}, nil

}

func CollectCpuDynamicRaw(coreCount int) (CpuDynRaw, error) {
	var raw_dyn C.cpu_dynamic_raw_t
	if C.collect_cpu_dynamic_raw(&raw_dyn) != 0 {
		return CpuDynRaw{}, errors.New("couldn't collect dynamic CPU info")
	}

	// create arrays for core holding
	var l_userTicks []uint64 = make([]uint64, coreCount)
	var l_systemTicks []uint64 = make([]uint64, coreCount)
	var l_idleTicks []uint64 = make([]uint64, coreCount)
	var l_niceTicks []uint64 = make([]uint64, coreCount)

	// iterate over each core and collect values
	for i := 0; i < coreCount; i++ {
		l_userTicks[i] = uint64(raw_dyn.user_ticks[i])
		l_systemTicks[i] = uint64(raw_dyn.system_ticks[i])
		l_idleTicks[i] = uint64(raw_dyn.idle_ticks[i])
		l_niceTicks[i] = uint64(raw_dyn.nice_ticks[i])
	}

	return CpuDynRaw{
		UserTicks:   l_userTicks,
		SystemTicks: l_systemTicks,
		IdleTicks:   l_idleTicks,
		NiceTicks:   l_niceTicks,
	}, nil

}
