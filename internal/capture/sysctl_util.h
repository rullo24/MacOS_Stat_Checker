#pragma once

// apple includes
#include <sys/sysctl.h>

int get_sysctl_int(const char* name, int *p_out);
