package tg_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/websuslik/unibot/tg"
	"testing"
)

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
