package aliyun_igraph_go_sdk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Client struct {
	Endpoint   string
	UserName   string
	PassWord   string
	Src        string
	httpClient *http.Client
}

func NewClient(endpoint string, userName string, passWord string, src string) *Client {
	return &Client{
		Endpoint:   endpoint,
		UserName:   userName,
		PassWord:   passWord,
		Src:        src,
		httpClient: defaultHttpClient,
	}
}

// WithRequestTimeout with custom timeout for a request
func (c *Client) WithRequestTimeout(timeout time.Duration) *Client {
	if c.httpClient == defaultHttpClient {
		c.httpClient = &http.Client{
			Timeout: timeout,
		}
	} else {
		c.httpClient.Timeout = timeout
	}
	return c
}

func (c *Client) Read(readRequest ReadRequest) (*Response, error) {
	vErr := readRequest.Validate()
	if vErr != nil {
		return nil, vErr
	}

	readRequest.AddQueryParam("src", c.Src)
	buildUri := readRequest.BuildUri()
	uri := buildUri.RequestURI()
	headers := map[string]string{}

	httpResp, err := request(c, "GET", uri, headers, nil)
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	buf, ioErr := ioutil.ReadAll(httpResp.Body)
	if ioErr != nil {
		return nil, NewBadResponseError(ioErr.Error(), httpResp.Header, httpResp.StatusCode)
	}

	readResult := ReadResult{}
	if jErr := json.Unmarshal(buf, &readResult); jErr != nil {
		fmt.Println(jErr)
		return nil, NewBadResponseError("Illegal readResult:"+string(buf), httpResp.Header, httpResp.StatusCode)
	}

	var resp *Response
	if len(readResult.ErrorInfo) == 0 {
		result := readResult.Result
		resp = NewResponse(result)
	} else {
		return nil, NewBadResponseError(fmt.Sprintf("Failed to read, message:%v",
			readResult.ErrorInfo), httpResp.Header, httpResp.StatusCode)
	}
	return resp, nil
}

func (c *Client) Write(writeRequest WriteRequest) (*Response, error) {
	vErr := writeRequest.Validate()
	if vErr != nil {
		return nil, vErr
	}
	buildUri := writeRequest.BuildUri()
	uri := buildUri.RequestURI()
	headers := map[string]string{}

	httpResp, err := request(c, "GET", uri, headers, nil)
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	buf, ioErr := ioutil.ReadAll(httpResp.Body)
	if ioErr != nil {
		return nil, NewBadResponseError(ioErr.Error(), httpResp.Header, httpResp.StatusCode)
	}
	writeResult := WriteResult{}
	if jErr := json.Unmarshal(buf, &writeResult); jErr != nil {
		fmt.Println(jErr)
		return nil, NewBadResponseError("Illegal writeResult:"+string(buf), httpResp.Header, httpResp.StatusCode)
	}

	switch writeResult.Errno {
	case 0:
		return NewResponse([]Result{}), nil
	case 1:
		return nil, NewBadResponseError(fmt.Sprintf("Failed to write, illegal reqeust body, errorCode[%v], resp:[%v]",
			writeResult.Errno, string(buf)), httpResp.Header, httpResp.StatusCode)
	default:
		return nil, NewBadResponseError(fmt.Sprintf("Failed to write, errorCode[%v], resp:[%v]",
			writeResult.Errno, string(buf)), httpResp.Header, httpResp.StatusCode)
	}

}
