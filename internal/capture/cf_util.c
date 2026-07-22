// user includes
#include "cf_util.h"

// @brief   Read an int value for `key` from `dict` into `out`.
// @return  0 on success, -1 if missing or wrong type
int cf_get_int(CFDictionaryRef dict, CFStringRef key, int *out) {
    CFNumberRef n = CFDictionaryGetValue(dict, key);
    if (n == NULL) {
        return -1;
    }
    return CFNumberGetValue(n, kCFNumberIntType, out) ? 0 : -1;
}

// @brief   Read a signed int32 value for `key` from `dict` into `out`.
// @return  0 on success, -1 if missing or wrong type.
int cf_get_int32(CFDictionaryRef dict, CFStringRef key, int32_t *out) {
    CFNumberRef n = CFDictionaryGetValue(dict, key);
    if (n == NULL) {
        return -1;
    }
    return CFNumberGetValue(n, kCFNumberSInt32Type, out) ? 0 : -1;
}

// @brief   Read a boolean value for `key` from `dict` into `out` (0/1).
// @return  0 on success, -1 if missing or wrong type.
int cf_get_bool(CFDictionaryRef dict, CFStringRef key, int *out) {
    CFBooleanRef b = CFDictionaryGetValue(dict, key);
    if (b == NULL) {
        return -1;
    }
    *out = CFBooleanGetValue(b) ? 1 : 0;
    return 0;
}
