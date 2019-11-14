package tg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChatID(t *testing.T) {
	tests := []struct {
		chatID *ChatID
		result []byte
	}{
		{chatID: &ChatID{ID: 123}, result: []byte("123")},
		{chatID: &ChatID{ID: 123, Username: "@hello"}, result: []byte("123")},
		{chatID: &ChatID{Username: "@hello"}, result: []byte("@hello")},
	}
	for _, test := range tests {
		v, e := test.chatID.MarshalJSON()
		assert.Equal(t, v, test.result)
		assert.Nil(t, e)
	}
}
