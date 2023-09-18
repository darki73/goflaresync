package entities

// RecordMeta is the definition of the metadata of a record.
type RecordMeta struct {
	// AutoAdded is a flag that indicates if the record was auto added.
	AutoAdded bool `json:"auto_added"`
	// Source is the source of the record.
	Source string `json:"source"`
}

// Record is the definition of a record.
type Record struct {
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
	// CreatedOn is the date when the record was created.
	CreatedOn string `json:"created_on"`
	// ID is the record ID.
	ID string `json:"id"`
	// Locked is a flag that indicates if the record is locked.
	Locked bool `json:"locked"`
	// Meta is the metadata of the record.
	Meta *RecordMeta `json:"meta"`
	// ModifiedOn is the date when the record was modified.
	ModifiedOn string `json:"modified_on"`
	// Proxiable is a flag that indicates if the record is proxiable.
	Proxiable bool `json:"proxiable"`
	// Tags is the list of tags of the record.
	Tags []string `json:"tags"`
	// TTL is the TTL of the record.
	TTL int `json:"ttl"`
	// ZoneID is the zone ID of the record.
	ZoneID string `json:"zone_id"`
	// ZoneName is the zone name of the record.
	ZoneName string `json:"zone_name"`
}

// RecordListResponse is the definition of the response of the record list API.
type RecordListResponse struct {
	// Errors is the list of errors.
	Errors []*Error `json:"errors"`
	// Messages is the list of messages.
	Messages []*Message `json:"messages"`
	// ResultInfo is the result information.
	ResultInfo *ResultInfo `json:"result_info"`
	// Success is a flag that indicates if the API call was successful.
	Success bool `json:"success"`
	// Result is the result of the API call.
	Result []*Record `json:"result"`
}
