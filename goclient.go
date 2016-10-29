package fcm

import (
	"fmt"
	"net/http"
	//"net/url"
	"io/ioutil"
	//"strings"
	"bytes"
	"encoding/json"
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

type DownstreamHttpMessage struct {
	To               string        `json:"to,omitempty"`
	RegistrationIds  []string      `json:"registration_ids,omitempty"`
	Priority         string        `json:"priority,omitempty"`
	ContentAvailable bool          `json:"content_available,omitempty"`
	TimeToLive       int64         `json:"time_to_live,omitempty"`
	DryRun           bool          `json:"dry_run,omitempty"`
	Notification     *Notification `json:"notification,omitempty"`
	//	Data             Data         `json:"data,omitempty"`
}

type Notification struct {
	Title string `json:"title,omitempty"`
	Body  string `json:"body,omitempty"`
	Sound string `json:"sound,omitempty"`
	Badge int64  `json:"badge,omitempty"`
	Icon  string `json:"icon,omitempty"`
	Tag   string `json:"tag,omitempty"`
}

// TODO
// type Data struct {
// 	data map[string]interface{}
// }

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

func New(key string) *client {
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
	reader := bytes.NewReader(body)

	// for debug
	fmt.Println(bytes.NewBuffer(body).String())

	req, err := http.NewRequest("POST", "https://fcm.googleapis.com/fcm/send", reader)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", c.serverKey)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer resp.Body.Close()

	return parse(resp)
}

func parse(resp *http.Response) (*Response, error) {
	var res Response
	b, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		fmt.Println(string(b))
	}
	if err := json.Unmarshal(b, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
