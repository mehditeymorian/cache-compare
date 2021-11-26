package response

type ValueResponse struct {
	Value string `json:"value"`
}

func NewValueResponse(value string) *ValueResponse {
	return &ValueResponse{Value: value}
}
