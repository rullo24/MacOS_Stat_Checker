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

}
