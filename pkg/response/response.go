package response

var (
	StatusSuccess = "success"
	StatusError   = "error"
)

type Body struct {
	Code    int                 `json:"code"`
	Status  string              `json:"status"`
	Message string              `json:"message"`
	Errors  map[string][]string `json:"errors,omitempty"`
	Data    any                 `json:"data,omitempty"`
}
