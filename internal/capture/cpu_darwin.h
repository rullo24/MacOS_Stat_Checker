#pragma once

// std lib includes
#include <stdint.h>

#define CPU_RAW_MAX_CORES       64
#define CPU_NAME_BUFFER_SIZE    1024

typedef struct {
    char name[CPU_NAME_BUFFER_SIZE];            // generous buffer (unsure of true, max possible size)
    int perf_core_count;                        // P-cores, from hw.perflevel0.physicalcpu
    int efficiency_core_count;                  // E-cores, from hw.perflevel1.physicalcpu
} cpu_static_raw_t;

typedef struct {
    uint64_t user_ticks[CPU_RAW_MAX_CORES];     // user_ticks for each CPU core
    uint64_t system_ticks[CPU_RAW_MAX_CORES];   // system_ticks for each CPU core
    uint64_t idle_ticks[CPU_RAW_MAX_CORES];     // idle_ticks for each CPU core
    uint64_t nice_ticks[CPU_RAW_MAX_CORES];     // nice_ticks for each CPU core
} cpu_dynamic_raw_t;
