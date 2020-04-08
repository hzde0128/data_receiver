package httpex

import (
	"io"
	"io/ioutil"
	"net/http"
	neturl "net/url"
)

type Client struct {
	http.Client
}

func (c *Client) GetEx(url string, params map[string]string, headers ...Header) (*Response, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	if params != nil {
		query := request.URL.Query()
		for k, v := range params {
			query.Add(k, v)
		}

		request.URL.RawQuery = query.Encode()
	}

	if headers != nil {
		for k, v := range headers[0] {
			request.Header.Set(k, v)
		}
	}

	resp, err := c.Do(request)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	binaryBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	responseEx := newResponse(resp, binaryBody)

	return responseEx, nil
}

func (c *Client) PostEx(url string, body io.Reader, params map[string]string, headers ...Header) (*Response, error) {
	request, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}

	if params != nil {
		query := request.URL.Query()
		for k, v := range params {
			query.Add(k, v)
		}

		request.URL.RawQuery = query.Encode()
	}

	if headers != nil {
		for k, v := range headers[0] {
			request.Header.Set(k, v)
		}
	}

	resp, err := c.Do(request)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	binaryBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	responseEx := newResponse(resp, binaryBody)

	return responseEx, nil
}

func (c *Client) PostFormEx(url string, body map[string]string) (*Response, error) {
	var data neturl.Values = make(neturl.Values)

	for k, v := range body {
		data[k] = append(make([]string, 0), v)
	}

	resp, err := c.PostForm(url, data)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	binaryBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	responseEx := newResponse(resp, binaryBody)

	return responseEx, nil
}
