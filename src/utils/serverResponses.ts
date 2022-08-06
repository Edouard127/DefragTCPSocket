export const responses = {
    DISCONNECT: "-1",
    UNAUTHORIZED: "0",
    OK: "1",
    TIMEOUT: "2",
    SERVER_CLOSED: "3",
    SERVER_ERROR: "4",
    BAD_COMMAND: "5",
    BAD_ARGUMENTS: "6",
    BAD_PASSWORD: "7",
    WORKER_NOT_FOUND: "8",
} as const;

export type ServerResponses = typeof responses[keyof typeof responses];
