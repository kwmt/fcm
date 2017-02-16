package fcm

import (
	"fmt"
	"net/http"
	//"net/url"
	"io/ioutil"
	//"strings"
	"bytes"
	"encoding/json"
	"io"
	"time"
)

type client struct {
	httpClient httpRunner
	serverKey  string
}

type httpRunner interface {
	Do(*http.Request) (*http.Response, error)
}

const (
	High   = "high"
	Normal = "normal"
)

type Result struct {
	MessageId      string `json:"message_id,omitempty"`
	RegistrationId string `json:"registration_id,omitempty"`
	Error          string `json:"error,omitempty"`
}

type Response struct {
	MessageId    int64    `json:"message_id,omitempty"`
	MulticastId  int64    `json:"multicast_id,omitempty"`
	Success      int64    `json:"success,omitempty"`
	Failure      int64    `json:"failure,omitempty"`
	CanonicalIds int64    `json:"canonical_ids,omitempty"`
	Results      []Result `json:"results,omitempty"`
}

type ResponseInfo struct {
	ApplicationVersion string `json:"applicationVersion,omitempty"`
	Application        string `json:"application,omitempty"`
	Scope              string `json:"scope,omitempty"`
	AuthorizedEntity   string `json:"authorizedEntity,omitempty"`
	Platform           string `json:"platform,omitempty"`
}

func NewClient(key string) *client {
	c := &client{
		httpClient: &http.Client{Timeout: time.Duration(30) * time.Second},
		serverKey:  "key=" + key,
	}
	return c
}

func (c *client) Send(message *DownstreamHttpMessage) (*Response, error) {
	body, err := json.Marshal(&message)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	// for debug
	// fmt.Println(bytes.NewBuffer(body).String())
	resp, err := c.request("POST", "https://fcm.googleapis.com/fcm/send", bytes.NewReader(body))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer resp.Body.Close()

	return toResponse(resp)
}

func (c *client) GetRegistrationTokenInfo(regToken string) (*ResponseInfo, error) {
	url := "https://iid.googleapis.com/iid/info/" + regToken
	resp, err := c.request("GET", url, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return toResponseInfo(resp)
}

func (c *client) request(httpMethod string, url string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(httpMethod, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", c.serverKey)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return resp, nil
}

func toResponse(resp *http.Response) (*Response, error) {
	var res Response
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Println(string(b))

	if err := json.Unmarshal(b, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// TODO: 共通化
func toResponseInfo(resp *http.Response) (*ResponseInfo, error) {
	var res ResponseInfo
	b, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		fmt.Println(string(b))
	}
	if err := json.Unmarshal(b, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
