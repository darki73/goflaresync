package entities

// ZoneAccount is the definition of an account in a zone.
type ZoneAccount struct {
	// ID is the account ID.
	ID string `json:"id"`
	// Name is the account name.
	Name string `json:"name"`
}

// ZoneMeta is the definition of the metadata of a zone.
type ZoneMeta struct {
	// Step is the step of the zone.
	Step int `json:"step"`
	// CDNOnly is a flag that indicates if the zone is CDN only.
	CDNOnly bool `json:"cdn_only"`
	// DNSOnly is a flag that indicates if the zone is DNS only.
	DNSOnly bool `json:"dns_only"`
	// CustomCertificateQuota is the quota of custom certificates.
	CustomCertificateQuota int `json:"custom_certificate_quota"`
	// FoundationDNS is a flag that indicates if the zone is foundation DNS.
	FoundationDNS bool `json:"foundation_dns"`
	// PageRuleQuota is the quota of page rules.
	PageRuleQuota int `json:"page_rule_quota"`
	// PhishingDetected is a flag that indicates if the zone is detected as phishing.
	PhishingDetected bool `json:"phishing_detected"`
}

// ZoneOwner is the definition of the owner of a zone.
type ZoneOwner struct {
	// ID is the owner ID.
	ID string `json:"id"`
	// Type is the owner type.
	Type string `json:"type"`
	// Name is the owner name.
	Name string `json:"name"`
}

// Zone is the definition of a zone.
type Zone struct {
	// Account is the account of the zone.
	Account *ZoneAccount `json:"account"`
	// ActivatedOn is the date when the zone was activated.
	ActivatedOn string `json:"activated_on"`
	// CreatedOn is the date when the zone was created.
	CreatedOn string `json:"created_on"`
	// DevelopmentMode is a flag that indicates if the zone is on development mode.
	DevelopmentMode int `json:"development_mode"`
	// ID is the zone ID.
	ID string `json:"id"`
	// Meta is the metadata of the zone.
	Meta *ZoneMeta `json:"meta"`
	// ModifiedOn is the date when the zone was modified.
	ModifiedOn string `json:"modified_on"`
	// Name is the zone name.
	Name string `json:"name"`
	// OriginalDnsHost is the original DNS host.
	OriginalDnsHost string `json:"original_dns_host"`
	// OriginalNameServers is the original name servers.
	OriginalNameServers []string `json:"original_name_servers"`
	// OriginalRegistrar is the original registrar.
	OriginalRegistrar string `json:"original_registrar"`
	// Owner is the owner of the zone.
	Owner *ZoneOwner `json:"owner"`
	// VanityNameServers is the vanity name servers.
	VanityNameServers []string `json:"vanity_name_servers"`
}

// ZoneListResponse is the definition of the response of the API call that returns a list of zones.
type ZoneListResponse struct {
	// Errors is a list of API errors returned by the API.
	Errors []*Error `json:"errors"`
	// Messages is a list of informational messages returned by the API.
	Messages []*Message `json:"messages"`
	// ResultInfo contains information about the result of the API call.
	ResultInfo *ResultInfo `json:"result_info"`
	// Success is a flag that indicates if the API call was successful.
	Success bool `json:"success"`
	// Result is the result of the API call.
	Result []*Zone `json:"result"`
}
