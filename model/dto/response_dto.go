package dto

type StatusResponse string

var (
	StatusSuccess StatusResponse = "SUCCESS"
	StatusFailed  StatusResponse = "FAILED"
)

type ResponseDTO[T any] struct {
	Code    int            `json:"code"`
	Status  StatusResponse `json:"status"`
	Message string         `json:"message"`
	Data    T              `json:"data,omitempty"`
	Error   any            `json:"error,omitempty"`
}
