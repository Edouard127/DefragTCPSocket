package structs

type LoginResponse struct {
	Player
	Success [1]byte
}

type AddWorkerResponse struct {
	Player
}

type RemoveWorkerResponse struct {
	Player
}

type InfoResponse struct {
	Player
	Health [1]byte
	Food   [1]byte
	X      [32]byte
	Y      [32]byte
	Z      [32]byte
}

type JobResponse struct {
	Player
	Job    [1]byte
	Status [1]byte
	Goal   []byte
}

type ErrorResponse struct {
	Player
	Error []byte
}

type ChatResponse struct {
	Player
}

type BaritoneResponse struct {
	Player
	Succeed [1]byte
}

type LambdaResponse struct {
	Player
	Succeed [1]byte
}

type HighwayResponse struct {
	Player
	Succeed [1]byte
}

type ScreenshotResponse struct {
	Player
	Data []byte
}
