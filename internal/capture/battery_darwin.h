#pragma once

// std lib includes
#include <stdint.h>

typedef struct {
    int current_capacity;
    int max_capacity;
    int design_capacity;
    int nominal_capacity;
    int cycle_count;
    int time_remaining;      // minutes, 65535 = N/A
    int avg_time_to_full;    // minutes, 65535 = N/A
    int32_t amperage_ma;     // signed: negative = discharging
    int voltage_mv;
    int temperature_centic;  // centi-Celsius, divide by 100 for Celsius
    int is_charging;
    int is_plugged_in;
    int is_fully_charged;
} battery_raw_t;

int collect_battery_raw(battery_raw_t *out);
