package models

// ErrorMessage error message response
type (
	ErrorMessage struct {
		UUID      string      `json:"uuid"`
		ID        int64       `json:"id"`
		SLUG      string      `json:"slug"`
		Timestamp int64       `json:"timestamp"`
		Updated   int64       `json:"updated"`
		Code      interface{} `json:"code"`
		Group     string      `json:"group"`
		ErrorKey  string      `json:"error_key"`
		Message   Lang        `json:"message"`
		Detail    Lang        `json:"detail"`
	}

	Lang struct {
		EN string `json:"en"`
		TH string `json:"th"`
	}
)
