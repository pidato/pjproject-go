//
// Created by CMO on 3/10/20.
//

#ifndef PJ_EXT_CALLBACKS_H
#define PJ_EXT_CALLBACKS_H

struct PiCallbacks {
    void (*PiEncoder_onFrame)(void *frame, unsigned long long prevExternNanos);
};

static struct PiCallbacks PI_CALLBACKS;

#endif //PJ_EXT_CALLBACKS_H
