export const bits = {
    EXIT: ~0x0,
    INIT: 0x0,
    CONNECT: 0x1,
    ADD_WORKER: 0x2,
    LOGIN: 0x3,
    LOGOUT: 0x4,
    CHAT: 0x5,
    BARITONE: 0x6,
    MOD_COMMAND: 0x7,
} as const;