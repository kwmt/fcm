package fcm

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestSend(t *testing.T) {
	const dummyKey = "dummyKey"

	tc := &testClient{
		resp: &http.Response{
			Status:     "200 OK",
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewReader([]byte(`{"message_id": 123456}`))),
		},
		err: nil,
	}
	c := NewClient(dummyKey)
	c.httpClient = tc

	msg := NewMessage("news").SetPriority(High)
	resp, err := c.Send(msg)

	if auth := tc.req.Header.Get("Authorization"); auth != "Bearer "+dummyKey {
		t.Errorf("key: %s, expect %s", auth, dummyKey)
	}

	if err != nil {
		t.Error(err)
	}
	if resp.MessageId != 123456 {
		t.Errorf("message id: got %d, expect %d", resp.MessageId, 123456)
	}

}

func TestGetRegistrationTokenInfo(t *testing.T) {
	const dummyKey = "dummyKey"

	tc := &testClient{
		resp: &http.Response{
			Status:     "200 OK",
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewReader([]byte(`{"applicationVersion":"1.0","application":"com.example","scope":"*","authorizedEntity":"1234","platform":"IOS"}`))),
		},
		err: nil,
	}
	c := NewClient(dummyKey)
	c.httpClient = tc

	resp, err := c.GetRegistrationTokenInfo("regToken")

	if err != nil {
		t.Error(err)
	}
	if resp.Application != "com.example" {
		t.Errorf("application: got %s, expect %s", resp.Application, "com.example")
	}
	if resp.Platform != "IOS" {
		t.Errorf("platform: got %s, expect %s", resp.Platform, "IOS")
	}

}

type testClient struct {
	req  *http.Request
	resp *http.Response
	err  error
}

func (c *testClient) Do(req *http.Request) (*http.Response, error) {
	c.req = req
	return c.resp, c.err
}
