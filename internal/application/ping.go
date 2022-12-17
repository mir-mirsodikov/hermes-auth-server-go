package application

type GetPingResponse struct {
	Message string `json:"message"`
}

func GetPing() GetPingResponse {
	return GetPingResponse{Message: "pong"}
}
