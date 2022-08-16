package response

type ServerResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
