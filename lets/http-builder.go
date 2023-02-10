package lets

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// Type for saving header value.
type httpHeader struct {
	Name  string
	Value string
}

// Type for saving oy builder params.
type HttpBuilder struct {
	url      string
	client   *http.Client
	headers  []*httpHeader
	response *http.Response
}

// Set http client with default configuration.
func (h *HttpBuilder) Default() {
	h.client = &http.Client{
		Timeout: time.Duration(5) * time.Second,
	}
}

// Manual set http builder.
func (h *HttpBuilder) SetClient(client *http.Client) {
	h.client = client
}

// Set End Point
func (h *HttpBuilder) SetUrl(url string) {
	LogI("HttpBuilder: set endPoint to \"%s\"", url)

	h.url = url
}

// Setting up header name and value.
func (h *HttpBuilder) AddHeader(name string, value string) {
	LogI("HttpBuilder: add header \"%s\":\"%s\"", name, value)

	for _, header := range h.headers {
		if header.Name == name {
			header.Value = value
			return
		}
	}

	h.headers = append(h.headers, &httpHeader{
		Name:  name,
		Value: value,
	})
}

// Automatically post data to oy.
func (h *HttpBuilder) Post(endPoint string, body interface{}) (fullUrl, response string, err error) {
	fullUrl = fmt.Sprintf("%s%s", h.url, endPoint)
	payloadString := ToJson(body)

	LogI("HttpBuilder: POST \"%s\"\n%s", fullUrl, payloadString)

	payload := strings.NewReader(payloadString)
	req, err := http.NewRequest(http.MethodPost, fullUrl, payload)
	if err != nil {
		return
	}

	// Header Setup
	for _, header := range h.headers {
		LogI("HttpBuilder: SetHeader: %s: %s", header.Name, header.Value)
		req.Header.Add(header.Name, header.Value)
	}

	h.response, err = h.client.Do(req)
	if err != nil {
		return
	}
	defer h.response.Body.Close()

	resBody, err := io.ReadAll(h.response.Body)
	if err != nil {
		return
	}

	response = string(resBody)
	LogI("HttpBuilder: Response: %v \n%s", h.response.StatusCode, response)
	return
}

func (h *HttpBuilder) GetRequest() *http.Response {
	return h.response
}

func (h *HttpBuilder) GetResponse() *http.Response {
	return h.response
}
