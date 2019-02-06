package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

type (
	HttpClient struct {
		BaseURL    *url.URL
		httpClient *http.Client
	}
)

func NewHttpClient(httpClient *http.Client, url *url.URL) *HttpClient {
	return &HttpClient{
		httpClient: httpClient,
		BaseURL:    url,
	}
}

func (c *HttpClient) NewRequest(method, path string, body interface{}, queryValues map[string]string) (*http.Request, error) {

	rURL := &url.URL{Path: path} // Set url

	// Adding url query params
	query := rURL.Query()
	if queryValues != nil {
		for k, v := range queryValues {
			query.Add(k, url.QueryEscape(v))
		}
	}
	rURL.RawQuery = query.Encode()
	u := c.BaseURL.ResolveReference(rURL)

	fmt.Println("*jas*", u.String())

	// Adding rq body
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

func (c *HttpClient) Do(ctx context.Context, req *http.Request, v interface{}) (*http.Response, error) {

	// https://medium.com/@marcus.olsson/adding-context-and-options-to-your-go-client-package-244c4ad1231b
	req = req.WithContext(ctx)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}
		return nil, err
	}

	// https://groups.google.com/forum/#!topic/golang-nuts/4Rr8BYVKrAI
	defer func() {
		io.Copy(ioutil.Discard, resp.Body)
		defer resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusOK {
		responseData, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, errors.Errorf("some went wrong. status %d, resp: %s", resp.StatusCode, string(responseData))
	}

	err = json.NewDecoder(resp.Body).Decode(v)
	return resp, err
}
