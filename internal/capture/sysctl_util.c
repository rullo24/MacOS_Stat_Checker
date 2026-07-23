// apple includes
#include <sys/sysctl.h>

// user includes
#include "sysctl_util.h"

// @brief   Read an int-valued sysctl by name into p_out
// @return  0 on success; -1 on failure
int get_sysctl_int(const char* name, int *p_out) {
    size_t size = sizeof(*p_out); // sizeof(int)
    return sysctlbyname(name, p_out, &size, NULL, 0) == 0 ? 0 : -1;
}
