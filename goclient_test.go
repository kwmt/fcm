package fcm

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestSend(t *testing.T) {

	tc := &testClient{
		resp: &http.Response{
			Status:     "200 OK",
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewReader([]byte(`{"message_id": 123456}`))),
		},
		err: nil,
	}
	c := New("key")
	c.httpClient = tc

	msg := new(DownstreamHttpMessage)
	msg.To = "/topics/all"
	msg.Priority = High
	resp, err := c.Send(msg)
	if err != nil {
		t.Error(err)
	}
	if resp.MessageId != 123456 {
		t.Errorf("body: got %s, expect %s", resp.MessageId, 123456)
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
