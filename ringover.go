package ringover

import (
  "bytes"
  "encoding/json"
  "errors"
  "io"
  "io/ioutil"
  "net/http"
  "net/url"
  "strings"
)

const (
  defaultRestEndpointURL     = "https://public-api.ringover.com/v2"
  defaultUSRestEndpointURL   = "https://public-api-us.ringover.com/v2"
  defaultRestEndpointVersion = "v2"
  defaultHeaderName          = "Authorization"
  acceptedContentType        = "application/json"
)

type ClientConfig struct {
  HttpClient          *http.Client
  RestEndpointURL     string
  RestEndpointVersion string
}

type auth struct {
  HeaderName string
  ApiKey     string
}

type Client struct {
  config *ClientConfig
  client *http.Client
  auth   *auth

  BaseURL *url.URL
}

// New returns a new APi Client
func New() *Client {
  config := ClientConfig{}

  config.HttpClient = http.DefaultClient
  config.RestEndpointURL = defaultRestEndpointURL
  config.RestEndpointVersion = defaultRestEndpointVersion

  // Create client
  baseURL, _ := url.Parse(config.RestEndpointURL)

  client := &Client{config: &config, client: config.HttpClient, auth: &auth{}, BaseURL: baseURL}

  return client
}

// Authenticate saves authenitcation parameters for user
func (client *Client) Authenticate(apiKey string) {
  parts := strings.Split(apiKey, "_")

  if len(parts) > 0 && parts[0] == "US" {
    client.config.RestEndpointURL = defaultUSRestEndpointURL
  }

  client.auth.HeaderName = defaultHeaderName
  client.auth.ApiKey = apiKey
}

// NewRequest creates an API request
func (client *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
  rel, err := url.Parse(client.config.RestEndpointVersion + urlStr)
  if err != nil {
    return nil, err
  }

  url := client.BaseURL.ResolveReference(rel)

  var buf io.ReadWriter
  if body != nil {
    buf = new(bytes.Buffer)
    err := json.NewEncoder(buf).Encode(body)
    if err != nil {
      return nil, err
    }
  }

  req, err := http.NewRequest(method, url.String(), buf)
  if err != nil {
    return nil, err
  }

  req.Header.Add(client.auth.HeaderName, client.auth.ApiKey)
  req.Header.Add("Accept", acceptedContentType)
  req.Header.Add("Content-type", acceptedContentType)

  return req, nil
}

// Do sends an API Request
func (client *Client) Do(req *http.Request) ([]byte, *http.Response, error) {
  response, err := client.client.Do(req)

  if err != nil {
    return nil, response, err
  }

  if response.StatusCode >= 500 {
    return nil, response, errors.New("Server error: " + response.Status)
  } else if response.StatusCode >= 400 {
    return nil, response, errors.New("Error: " + response.Status)
  }

  defer response.Body.Close()

  data, err := ioutil.ReadAll(response.Body)

  if err != nil {
    return nil, response, err
  }

  return data, response, nil
}
