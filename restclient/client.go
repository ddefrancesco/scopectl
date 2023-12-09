package restclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type ScopeBodyRequest struct {
	Body string `json:"body"`
}
type ScopeResponse struct {
	Code     int    `json:"code"`
	Response string `json:"response"`
	Cmd      string `json:"cmd"`
}

type ScopeErr struct {
	Err            int    `json:"error_code"`
	ErrDescription string `json:"error_description"`
	ScopeFunction  string `json:"scope_function"`
	Cmd            string `json:"cmd"`
}

type RequestPath struct {
	Command string
	Items   map[string]string
}

type EtxRestClient struct {
	BaseURL     string
	Method      string
	PathParams  RequestPath
	RequestBody ScopeBodyRequest
	httpclient  *http.Client
}

func NewClient(baseURL string, method string, pathParams RequestPath, requestBody ScopeBodyRequest) *EtxRestClient {
	return &EtxRestClient{
		BaseURL:     baseURL,
		Method:      method,
		PathParams:  pathParams,
		RequestBody: requestBody,
		httpclient:  &http.Client{},
	}
}

func (c *EtxRestClient) doRequest(headers map[string]string, body io.Reader) (*http.Response, error) {

	var buf bytes.Buffer
	for _, value := range c.PathParams.Items {
		//buf.WriteString("/")
		buf.WriteString(value)

	}

	pprms := c.BaseURL + "/" + c.PathParams.Command
	log.Println("Calling resource: " + pprms)
	req, err := http.NewRequest(c.Method, pprms, body)
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := c.httpclient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *EtxRestClient) decodeJSON(resp *http.Response, v interface{}) error {
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(v)
}

func (c *EtxRestClient) encodeJSON(v interface{}) (io.Reader, error) {
	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(v)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (c *EtxRestClient) GetPost() (*ScopeResponse, error) {
	body, err := c.encodeJSON(c.RequestBody)
	if err != nil {
		return nil, err
	}
	resp, err := c.doRequest(nil, body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusAccepted {
		return nil, fmt.Errorf("error fetching request: %s", resp.Status)
	}

	var post ScopeResponse
	if err := c.decodeJSON(resp, &post); err != nil {
		return nil, err
	}

	return &post, nil
}
