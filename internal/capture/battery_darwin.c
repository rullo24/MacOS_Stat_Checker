// apple includes
#include <CoreFoundation/CoreFoundation.h>
#include <IOKit/IOKitLib.h>

// user includes
#include "cf_util.h"
#include "battery_darwin.h"

// @brief   Collects raw battery register values from CoreFoundation
// @return  Returns 0 on success, non-zero if no battery service was found
int collect_battery_raw(battery_raw_t *out) {

    // look up the battery's IOKit service by class name
    io_service_t service = IOServiceGetMatchingService(
        kIOMainPortDefault,
        IOServiceMatching("AppleSmartBattery")
    );

    // pull all properties for this service into one dict.
    CFMutableDictionaryRef props = NULL;
    kern_return_t kr = IORegistryEntryCreateCFProperties(
        service,
        &props,
        kCFAllocatorDefault,
        0
    );
    IOObjectRelease(service);

    // check if properties can be captured from IOKit service
    if (kr != KERN_SUCCESS || props == NULL) {
        return -1;
    }

    // extract each field by key (get_<> calls require address of obj value)
    cf_get_int(props, CFSTR("CurrentCapacity"), &out->current_capacity);
    cf_get_int(props, CFSTR("MaxCapacity"), &out->max_capacity);
    cf_get_int(props, CFSTR("DesignCapacity"), &out->design_capacity);
    cf_get_int(props, CFSTR("NominalChargeCapacity"), &out->nominal_capacity);
    cf_get_int(props, CFSTR("CycleCount"), &out->cycle_count);
    cf_get_int(props, CFSTR("TimeRemaining"), &out->time_remaining);
    cf_get_int(props, CFSTR("AvgTimeToFull"), &out->avg_time_to_full);
    cf_get_int32(props, CFSTR("Amperage"), &out->amperage_ma);
    cf_get_int(props, CFSTR("Voltage"), &out->voltage_mv);
    cf_get_int(props, CFSTR("Temperature"), &out->temperature_centic);
    cf_get_bool(props, CFSTR("IsCharging"), &out->is_charging);
    cf_get_bool(props, CFSTR("ExternalConnected"), &out->is_plugged_in);
    cf_get_bool(props, CFSTR("FullyCharged"), &out->is_fully_charged);

    // free dictionary
    CFRelease(props);

    return 0;
}
