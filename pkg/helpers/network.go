package helpers

import (
	"github.com/darki73/goflaresync/pkg/configuration"
	"github.com/darki73/goflaresync/pkg/log"
	"io"
	"net/http"
)

// GetExternalAddress returns the current public IP address.
func GetExternalAddress() (string, error) {
	response, err := http.Get(configuration.GetConfiguration().GetWatcher().GetAddressSource())
	if err != nil {
		return "", err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalf("error closing response body: %s", err)
		}
	}(response.Body)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
