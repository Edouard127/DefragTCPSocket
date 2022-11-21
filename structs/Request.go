package structs

type Player struct {
	Player [16]byte
}

type LoginRequest struct {
	Player
	Server [32]byte
	LAN    [1]byte
}

type LogoutRequest struct {
	Player
}

type AddWorkerRequest struct{}

type RemoveWorkerRequest struct {
	Player
}

type InfoRequest struct {
	Player
}

type RotateRequest struct {
	Player
	Yaw   [32]byte
	Pitch [32]byte
}

type JobRequest struct {
	Player
	Job [1]byte
}

type ErrorRequest struct {
	Player
	Error []byte
}

type ChatRequest struct {
	Player
	Message []byte
}

type BaritoneRequest struct {
	Player
	Command []byte
}

type LambdaRequest struct {
	Player
	Command []byte
}

type HighwayRequest struct {
	Player
	Command []byte
}

type ScreenshotRequest struct {
	Player
}
