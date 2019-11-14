# Unibot

Unibot is a Telegram Bot API implementation written in Golang for an educational purpose.

**I do not recommend using its in production code.**

## Usage example

Just a simple echobot. Write some text message to a Bot and it will send the same message back to you. 

This example is using long polling to receive updates.
```go
package main

import (
	"github.com/websuslik/unibot/tg"
	"log"
	"net/http"
	"time"
)

type Client struct {
}

func (m Client) Do(req *http.Request, timeout time.Duration) (*http.Response, error) {
	client := &http.Client{Timeout: timeout}
	return client.Do(req)
}

func main() {
	api := tg.API{
		Token:  "873254919:AAGqsPS3GXmKi3l-JMqHHUPfFbfZLuHFO6E",
		Client: new(Client),
	}
	offset := 0
	for {
		args := &tg.GetUpdatesArgs{
			Timeout:        30,
			AllowedUpdates: []string{tg.AllowedUpdateMessage},
			Offset:         offset,
		}
		if updates, err := api.GetUpdates(args); err != nil {
			log.Println(err)
		} else {
			for _, update := range *updates {
				messageArgs := &tg.SendMessageArgs{
					ChatID: &tg.ChatID{ID: update.Message.Chat.ID},
					Text:   update.Message.Text,
				}
				if _, err := api.SendMessage(messageArgs); err != nil {
					log.Println(err)
				}
				offset = update.UpdateID + 1
			}
		}
	}
}


```