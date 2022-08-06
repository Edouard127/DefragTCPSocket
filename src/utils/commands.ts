export const bits = {
    EXIT: ~0x0,
    INIT: 0x0,
    CONNECT: 0x1,
    ADD_WORKER: 0x2,
    DISCONNECT: 0x3,
    CHAT: 0x4,
    BARITONE: 0x5,
    MOD_COMMAND: 0x6,
} as const;