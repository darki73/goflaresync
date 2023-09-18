package entities

// Error is the definition of the error in the response
type Error struct {
	// Code is the code of the error
	Code int `json:"code"`
	// Message is the message of the error
	Message string `json:"message"`
}

// Message is the definition of the message in the response
type Message struct {
	// Code is the code of the message
	Code int `json:"code"`
	// Message is the message of the message
	Message string `json:"message"`
}

// ResultInfo is the definition of the result info in the response
type ResultInfo struct {
	// Count is the number of results returned in the response
	Count int `json:"count"`
	// Page is the current page number
	Page int `json:"page"`
	// PerPage is the number of results per page returned in the response
	PerPage int `json:"per_page"`
	// TotalCount is the total number of results returned in the response
	TotalCount int `json:"total_count"`
}
