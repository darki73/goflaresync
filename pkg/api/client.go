package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/darki73/goflaresync/pkg/api/entities"
	"github.com/darki73/goflaresync/pkg/configuration"
	"github.com/darki73/goflaresync/pkg/log"
	"github.com/darki73/goflaresync/pkg/version"
	"io"
	"net/http"
	"strings"
)

// API is the definition of the Cloudflare API.
type API struct {
	// email is the email of the Cloudflare API.
	email string
	// token is the token of the Cloudflare API.
	token string
	// userAgent is the user agent of the Cloudflare API.
	userAgent string
	// baseURL is the base URL of the Cloudflare API.
	baseURL string
	// authenticated is a flag that indicates if the Cloudflare API is authenticated.
	authenticated bool
}

// NewClient returns a new Cloudflare API client.
func NewClient() (*API, error) {
	userAgent := fmt.Sprintf(
		"goflaresync/%s-%s",
		version.GetVersion(),
		version.GetCommit(),
	)

	config := configuration.GetConfiguration().GetCredentials()

	api := &API{
		email:         config.GetEmail(),
		token:         config.GetToken(),
		userAgent:     userAgent,
		baseURL:       "https://api.cloudflare.com/client/v4",
		authenticated: false,
	}

	if err := api.Authenticate(); err != nil {
		return nil, err
	}

	return api, nil
}

// Authenticate authenticates the Cloudflare API.
func (api *API) Authenticate() error {
	log.DebugWithFields(
		"aAttempting to authenticate user with provided credentials",
		log.FieldsMap{
			"source": "api",
			"action": "authenticate",
		},
	)
	body, err := api.query(http.MethodGet, "user/tokens/verify", nil)
	if err != nil {
		return err
	}

	response := &entities.TokenVerifyResponse{}

	if err := json.Unmarshal(body, response); err != nil {
		return err
	}

	if !response.Success {
		log.DebugWithFields(
			"failed to authenticate user due to invalid credentials",
			log.FieldsMap{
				"source": "api",
				"action": "authenticate",
			},
		)
		return ErrInvalidCredentials
	}

	if response.Result.Status != "active" {
		log.DebugWithFields(
			"successfully authenticated user but token is not active",
			log.FieldsMap{
				"source": "api",
				"action": "authenticate",
			},
		)
		return ErrTokenInactive
	}

	api.authenticated = true

	log.DebugWithFields(
		"successfully authenticated user",
		log.FieldsMap{
			"source": "api",
			"action": "authenticate",
		},
	)

	return nil
}

// ListZones returns a list of zones.
func (api *API) ListZones() (*entities.ZoneListResponse, error) {
	if !api.isAuthenticated() {
		return nil, ErrNotAuthenticated
	}

	body, err := api.query(http.MethodGet, "zones", nil)
	if err != nil {
		return nil, err
	}

	response := &entities.ZoneListResponse{}

	if err := json.Unmarshal(body, response); err != nil {
		return nil, err
	}

	return response, nil
}

// ListRecords returns a list of records.
func (api *API) ListRecords(zone *entities.Zone) (*entities.RecordListResponse, error) {
	if !api.isAuthenticated() {
		return nil, ErrNotAuthenticated
	}

	body, err := api.query(http.MethodGet, fmt.Sprintf("zones/%s/dns_records", zone.ID), nil)
	if err != nil {
		return nil, err
	}

	response := &entities.RecordListResponse{}

	if err := json.Unmarshal(body, response); err != nil {
		return nil, err
	}

	return response, nil
}

// UpdateRecord updates a record.
func (api *API) UpdateRecord(zone *entities.Zone, record *entities.Record) (*entities.RecordUpdateResponse, error) {
	if !api.isAuthenticated() {
		return nil, ErrNotAuthenticated
	}

	body, err := api.query(http.MethodPut, fmt.Sprintf("zones/%s/dns_records/%s", zone.ID, record.ID), entities.RecordUpdateRequest{
		Content: record.Content,
		Name:    record.Name,
		Proxied: record.Proxied,
		Type:    record.Type,
		Comment: record.Comment,
		TTL:     record.TTL,
		Tags:    record.Tags,
	})

	if err != nil {
		return nil, err
	}

	response := &entities.RecordUpdateResponse{}

	if err := json.Unmarshal(body, response); err != nil {
		return nil, err
	}

	return response, nil
}

// getToken returns the token of the Cloudflare API.
func (api *API) getToken() string {
	return api.token
}

// getEmail returns the email of the Cloudflare API.
func (api *API) getEmail() string {
	return api.email
}

// getUserAgent returns the user agent of the Cloudflare API.
func (api *API) getUserAgent() string {
	return api.userAgent
}

// getBaseURL returns the base URL of the Cloudflare API.
func (api *API) getBaseURL() string {
	return api.baseURL
}

// getHeaders returns the headers of the Cloudflare API.
func (api *API) getHeaders() map[string]string {
	return map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", api.getToken()),
		"X-Auth-Email":  api.getEmail(),
		"Content-Type":  "application/json",
		"User-Agent":    api.getUserAgent(),
	}
}

// isAuthenticated returns a flag that indicates if the Cloudflare API is authenticated.
func (api *API) isAuthenticated() bool {
	return api.authenticated
}

// query queries the Cloudflare API.
func (api *API) query(method string, path string, query interface{}) ([]byte, error) {
	log.Tracef("[API] Query method called with the following arguments: method=%s, path=%s, query=%v", method, path, query)
	if strings.HasPrefix(path, "/") {
		path = strings.TrimPrefix(path, "/")
	}

	url := fmt.Sprintf("%s/%s", api.getBaseURL(), path)

	var request *http.Request
	var err error

	switch method {
	case http.MethodGet:
		if query != nil {
			queryString, err := json.Marshal(query)
			if err != nil {
				return nil, err
			}
			url = fmt.Sprintf("%s?%s", url, string(queryString))
		}
		request, err = http.NewRequest(method, url, nil)
	case http.MethodPost, http.MethodPut:
		jsonData, err := json.Marshal(query)
		if err != nil {
			return nil, err
		}
		request, err = http.NewRequest(method, url, bytes.NewBuffer(jsonData))
	default:
		return nil, fmt.Errorf("invalid method: %s", method)
	}

	if err != nil {
		return nil, err
	}

	for key, value := range api.getHeaders() {
		request.Header.Set(key, value)
	}

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned non-200 status code: %d", response.StatusCode)
	}

	return io.ReadAll(response.Body)
}
