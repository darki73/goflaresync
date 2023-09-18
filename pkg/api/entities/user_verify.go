package entities

// TokenVerify is the definition of the token verify API.
type TokenVerify struct {
	// ExpiresOn is the date when the token expires.
	ExpiresOn string `json:"expires_on"`
	// ID is the token ID.
	ID string `json:"id"`
	// NotBefore is the date before which token is not valid.
	NotBefore string `json:"not_before"`
	// Status is the status of the token.
	Status string `json:"status"`
}

// TokenVerifyResponse is the definition of the response of the token verify API.
type TokenVerifyResponse struct {
	// Errors is the list of errors.
	Errors []*Error `json:"errors"`
	// Messages is the list of messages.
	Messages []*Message `json:"messages"`
	// Result is the result of the API call.
	Result *TokenVerify `json:"result"`
	// Success is a flag that indicates if the API call was successful.
	Success bool `json:"success"`
}
