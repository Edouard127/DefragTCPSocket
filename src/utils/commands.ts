export const bits = {
    EXIT: ~0x0,
    LOGIN: 0x0,
    LOGOUT: 0x1,
    ADD_WORKER: 0x2,
    REMOVE_WORKER: 0x3,
    CHAT: 0x4,
    BARITONE: 0x5,
    MOD_COMMAND: 0x6,
} as const;