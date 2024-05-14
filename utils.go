package art

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func (a *ArtClient) doRequest(ctx context.Context, body io.Reader, method, path string) (*http.Response, error) {
	if path == "" {
		path = "/"
	}
	if !strings.HasPrefix(path, "/") {
		path = fmt.Sprintf("/%s", path)
	}
	url := fmt.Sprintf("%s%s", a.BaseURL, path)
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	res, err := a.Client.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func checkApiError(code int, status string) error {
	if code < http.StatusOK || code >= http.StatusBadRequest {
		return fmt.Errorf("error code: %d\nerror message: %s", code, status)
	}
	return nil
}
