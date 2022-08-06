import { Socket } from "net"

export default interface Timeout {
    PONG: Pong;
    TIME: number;
}
interface Pong {
    socket: Socket;
    code: string;
}