package dome9

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const (
	libraryVersion = "0.1.0"
	defaultBaseURL = "https://api.dome9.com/"
	userAgent      = "godome9/" + libraryVersion
	mediaType      = "application/json"
)

// Credentials for Dome9 API.
type Credentials struct {
	// API key.
	KeyID string

	// API secret.
	KeySecret string
}

// Client for Dome9 v2 API.
type Client struct {
	// API Credentials.
	Credentials *Credentials

	// HTTP client.
	client *http.Client

	// API Base URL.
	BaseURL *url.URL

	// HTTP User agent.
	UserAgent string

	// Services used for communicating with the API
	AzureCloudAccounts AzureCloudAccountsService
}

// NewClient returns a new Dome9 API client.
func NewClient(httpClient *http.Client, credentials *Credentials) (*Client, error) {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	if credentials == nil {
		return nil, fmt.Errorf("Credentials must be provided")
	}

	baseURL, err := url.Parse(defaultBaseURL)
	if err != nil {
		return nil, err
	}

	c := &Client{client: httpClient, Credentials: credentials, BaseURL: baseURL, UserAgent: userAgent}
	c.AzureCloudAccounts = &AzureCloudAccountsServiceOp{client: c}
	return c, nil
}

// ClientOpt are options for New.
type ClientOpt func(*Client) error

// New returns a new Dome 9 API client instance.
func New(httpClient *http.Client, credentials *Credentials, opts ...ClientOpt) (*Client, error) {
	c, err := NewClient(httpClient, credentials)
	if err != nil {
		return nil, err
	}
	for _, opt := range opts {
		if err := opt(c); err != nil {
			return nil, err
		}
	}

	return c, nil
}

// SetBaseURL is a client option for setting the base URL.
func SetBaseURL(bu string) ClientOpt {
	return func(c *Client) error {
		u, err := url.Parse(bu)
		if err != nil {
			return err
		}

		c.BaseURL = u
		return nil
	}
}

// SetUserAgent is a client option for setting the user agent.
func SetUserAgent(ua string) ClientOpt {
	return func(c *Client) error {
		c.UserAgent = fmt.Sprintf("%s %s", ua, c.UserAgent)
		return nil
	}
}

// NewRequest creates an API request. A relative URL can be provided in urlStr, which will be resolved to the
// BaseURL of the Client. Relative URLS should always be specified without a preceding slash. If specified, the
// value pointed to by body is JSON encoded and included in as the request body.
func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)

	buf := new(bytes.Buffer)
	if body != nil {
		err = json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(c.Credentials.KeyID, c.Credentials.KeySecret)

	req.Header.Add("Content-Type", mediaType)
	req.Header.Add("Accept", mediaType)
	req.Header.Add("User-Agent", c.UserAgent)

	return req, nil
}

// Do sends an API request and returns the API response. The API response is JSON decoded and stored in the value
// pointed to by v, or returned as an error if an API error has occurred. If v implements the io.Writer interface,
// the raw response will be written to v, without attempting to decode it.
func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() {
		if rerr := resp.Body.Close(); err == nil {
			err = rerr
		}
	}()

	err = CheckResponse(resp)
	if err != nil {
		return resp, err
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			_, err = io.Copy(w, resp.Body)
			if err != nil {
				return nil, err
			}
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
			if err != nil {
				return nil, err
			}
		}
	}

	return resp, err
}

// CheckResponse checks the API response for errors, and returns them if present.
// A response is considered an error if it has a status code outside the 200 range.
func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; c >= 200 && c <= 299 {
		return nil
	}
	return fmt.Errorf("Response status is: %v", r.StatusCode)
}
