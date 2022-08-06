export const bits = {
    EXIT: ~0x0,
    HEARTBEAT: 0x0,
    LOGIN: 0x1,
    LOGOUT: 0x2,
    ADD_WORKER: 0x3,
    REMOVE_WORKER: 0x4,
    CHAT: 0x5,
    BARITONE: 0x6,
    MOD_COMMAND: 0x7,
} as const;