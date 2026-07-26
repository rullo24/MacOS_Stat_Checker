// apple includes
#include <mach/mach.h>              // mach_task_self(), core Mach types (kern_result_t)
#include <mach/mach_host.h>         // mach_host_self(), host_processor_info()
#include <mach/processor_info.h>    // processor_cpu_load_info_t, PROCESSOR_CPU_LOAD_INFO
#include <mach/vm_map.h>            // vm_deallocate()

// user includes
#include "cpu_darwin.h"
#include "sysctl_util.h"

// @brief   Captures the name of the CPU + core topology
// @return  0 on success; -1 on failure
int collect_cpu_static_raw(cpu_static_raw_t *p_out) {
    // collecting CPU name string
    size_t name_size = CPU_NAME_BUFFER_SIZE;
    int name_res = sysctlbyname("machdep.cpu.brand_string", p_out->name, &name_size, NULL, 0);

    // collecting core counts
    int perf_res = get_sysctl_int("hw.perflevel0.physicalcpu", &p_out->perf_core_count);
    int eff_res = get_sysctl_int("hw.perflevel1.physicalcpu", &p_out->efficiency_core_count);

    if (name_res != 0 || perf_res != 0 || eff_res != 0) {
        return -1; // failed
    }

    return 0; // success
}

// @brief   Captures current CPU tick counts
// @return  0 on success; -1 on failure
int collect_cpu_dynamic_raw(cpu_dynamic_raw_t *p_out) {
    natural_t core_count = 0;
    processor_info_array_t info_array;
    mach_msg_type_number_t info_count;

    // capture CPU tick info
    kern_return_t kr = host_processor_info(
        mach_host_self(),
        PROCESSOR_CPU_LOAD_INFO,
        &core_count,
        &info_array,
        &info_count
    );

    // capture CPU tick info success check
    if (kr != KERN_SUCCESS) {
        return -1;
    }

    // clamp - shouldn't happen on real hardware
    if (core_count > CPU_RAW_MAX_CORES) {
        core_count = CPU_RAW_MAX_CORES;
    }

    // cycle over each core in CPU
    processor_cpu_load_info_t cpu_load = (processor_cpu_load_info_t)info_array;
    for (natural_t i = 0; i < core_count; i++) {
        p_out->user_ticks[i]    = cpu_load[i].cpu_ticks[CPU_STATE_USER];
        p_out->system_ticks[i]  = cpu_load[i].cpu_ticks[CPU_STATE_SYSTEM];
        p_out->idle_ticks[i]    = cpu_load[i].cpu_ticks[CPU_STATE_IDLE];
        p_out->nice_ticks[i]    = cpu_load[i].cpu_ticks[CPU_STATE_NICE];
    }

    // release call for info_array
    vm_deallocate(
        mach_task_self(),
        (vm_address_t)info_array,
        info_count * sizeof(*info_array)
    );

    return 0;
}
