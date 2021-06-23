package domain

// Response ...
type Response struct {
	Status     int                    `json:"status"`
	Message    string                 `json:"message"`
	Validation map[string]interface{} `json:"validation"`
	Data       interface{}            `json:"data"`
}
