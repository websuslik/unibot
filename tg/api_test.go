package tg_test

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/websuslik/unibot/tg"
	"testing"
	"time"
)

var commonUser = map[string]interface{}{
	"id":         123,
	"is_bot":     false,
	"first_name": "Yuri",
}

var commonChat = map[string]interface{}{
	"id":   123,
	"type": "private",
}

var commonMessage = map[string]interface{}{
	"message_id": 123,
	"from":       commonUser,
	"date":       123,
	"chat":       commonChat,
	"text":       "Hello, World!",
}

var commonTrueResponse = map[string]interface{}{
	"ok":     true,
	"result": true,
}

type HttpClientMock struct {
	mock.Mock
}

func (m *HttpClientMock) Do(url string, args *tg.RequestArgs, timeout time.Duration) ([]byte, error) {
	callArgs := m.Called(url, args, timeout)
	return callArgs.Get(0).([]byte), callArgs.Error(1)
}

func setUpMock(method string, bodyData map[string]interface{}) (*HttpClientMock, *tg.API) {
	m := new(HttpClientMock)
	body, _ := json.Marshal(bodyData)
	url := fmt.Sprintf("https://api.telegram.org/botTOKEN/%s", method)
	m.On("Do", url, mock.Anything, tg.Timeout*time.Second).Return(body, nil)
	api := &tg.API{Token: "TOKEN", Client: m}
	return m, api
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

func TestGetUpdates(t *testing.T) {
	m, api := setUpMock("getUpdates", map[string]interface{}{
		"ok": true,
		"result": []interface{}{
			map[string]interface{}{
				"update_id": 123,
				"message":   commonMessage,
			},
		},
	})
	args := &tg.GetUpdatesArgs{}
	res, err := api.GetUpdates(args)
	m.AssertExpectations(t)
	assert.Nil(t, err)
	assert.Equal(t, len(*res), 1)
	for _, value := range *res {
		assert.Equal(t, value.UpdateID, 123)
		assert.Equal(t, value.Message.Text, "Hello, World!")
	}
}

func TestSetWebhook(t *testing.T) {
	m, api := setUpMock("setWebhook", commonTrueResponse)
	args := &tg.SetWebhookArgs{}
	res, err := api.SetWebhook(args)
	m.AssertExpectations(t)
	assert.Nil(t, err)
	assert.Equal(t, *res, true)
}
