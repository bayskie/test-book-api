package formatter

type APIResponse struct {
	Data  any `json:"data"`
	Error any `json:"error"`
}

func NewResponse(data any, err error) APIResponse {
	if err != nil {
		return APIResponse{Data: nil, Error: err.Error()}
	}
	return APIResponse{Data: data, Error: nil}
}
