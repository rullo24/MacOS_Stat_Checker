#pragma once

// std lib includes
#include <stdint.h>

// apple includes
#include <CoreFoundation/CoreFoundation.h>

int cf_get_int(CFDictionaryRef dict, CFStringRef key, int *out);
int cf_get_int32(CFDictionaryRef dict, CFStringRef key, int32_t *out);
int cf_get_bool(CFDictionaryRef dict, CFStringRef key, int *out);
