package tg_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/websuslik/unibot/tg"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
	"time"
)

type HTTPClientMock struct {
	mock.Mock
}

func (m *HTTPClientMock) Do(req *http.Request, timeout time.Duration) (*http.Response, error) {
	args := m.Called(req)
	return args.Get(0).(*http.Response), args.Error(1)
}

func TestChatID(t *testing.T) {
	tests := []struct {
		chatID *tg.ChatID
		result []byte
	}{
		{chatID: &tg.ChatID{ID: 123}, result: []byte("123")},
		{chatID: &tg.ChatID{ID: 123, Username: "@hello"}, result: []byte("123")},
		{chatID: &tg.ChatID{Username: "@hello"}, result: []byte("@hello")},
	}
	for _, test := range tests {
		v, e := test.chatID.MarshalJSON()
		assert.Equal(t, v, test.result)
		assert.Nil(t, e)
	}
}
func setMockResponseBody(m *HTTPClientMock, body string) {
	m.On("Do", mock.Anything).Return(&http.Response{
		Body: ioutil.NopCloser(strings.NewReader(body))}, nil)
}
func TestGetUpdates(t *testing.T) {
	args := &tg.GetUpdatesArgs{}
	req, _ := args.GetRequestArgs()
	assert.Equal(t, req.Headers, map[string]string{"Content-Type": "application/json"})

	client := new(HTTPClientMock)
	setMockResponseBody(client, "{\"ok\": true, \"result\": [{}, {}]}")

	api := &tg.API{
		Token:  "TOKEN",
		Client: client,
	}
	res, err := api.GetUpdates(args)
	assert.Nil(t, err)
	assert.Equal(t, len(*res), 2)
}
