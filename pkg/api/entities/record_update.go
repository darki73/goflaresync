package entities

// RecordUpdateRequest is the definition of the request of the record update API.
type RecordUpdateRequest struct {
	// Content is the content of the record.
	Content string `json:"content"`
	// Name is the name of the record.
	Name string `json:"name"`
	// Proxied is a flag that indicates if the record is proxied.
	Proxied bool `json:"proxied"`
	// Type is the type of the record.
	Type string `json:"type"`
	// Comment is the comment of the record.
	Comment string `json:"comment"`
	// TTL is the TTL of the record.
	TTL int `json:"ttl"`
	// Tags is the list of tags of the record.
	Tags []string `json:"tags"`
}

// RecordUpdateResponse is the definition of the response of the record update API.
type RecordUpdateResponse struct {
	// Errors is the list of errors.
	Errors []*Error `json:"errors"`
	// Messages is the list of messages.
	Messages []*Message `json:"messages"`
	// Result is the result of the API call.
	Result *Record `json:"result"`
	// Success is a flag that indicates if the API call was successful.
	Success bool `json:"success"`
}
