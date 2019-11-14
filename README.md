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
	"time"
)

func main() {
	api := tg.API{Token: "API_TOKEN"}
	offset := 0
	for {
		args := &tg.GetUpdatesArgs{
			Timeout: 30,
			AllowedUpdates: []string{tg.AllowedUpdateMessage},
			Offset: offset,
		}
		if updates, err := api.GetUpdates(args); err != nil {
			log.Println(err)
		} else {
			for _, update := range *updates {
				messageArgs := &tg.SendMessageArgs{
					ChatID: &tg.ChatID{ID: update.Message.Chat.ID},
					Text: update.Message.Text,
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