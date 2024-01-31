package syncapi

import (
	"context"
	"encoding/json"
	"net/http"
)

func getRequest[T any](ctx context.Context, cli *http.Client, apikey string, url string) (res T, err error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return
	}

	req.Header.Set(apiKeyHeader, apikey)

	resp, err := cli.Do(req)
	if err != nil {
		return
	}

	err = json.NewDecoder(resp.Body).Decode(&res)

	return
}

func postRequest[T any](ctx context.Context, cli *http.Client, apikey string, url string) (res T, err error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return
	}

	req.Header.Set(apiKeyHeader, apikey)

	resp, err := cli.Do(req)
	if err != nil {
		return
	}

	err = json.NewDecoder(resp.Body).Decode(&res)

	return
}
