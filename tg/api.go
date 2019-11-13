package tg

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

// https://core.telegram.org/bots/api Bot API 4.4
type API struct {
	Token string
}

type MethodArgs interface {
	GetRequestArgs() (*RequestArgs, error)
}

type RequestArgs struct {
	Body    *bytes.Buffer
	Headers map[string]string
}

func (api *API) buildRequest(method string, args MethodArgs) (*http.Request, error) {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/%s", api.Token, method)
	requestArgs, err := args.GetRequestArgs()
	if err != nil {
		return nil, err
	}
	request, _ := http.NewRequest("POST", url, requestArgs.Body)
	for key, value := range requestArgs.Headers {
		request.Header.Set(key, value)
	}
	return request, nil
}

func (api *API) sendRequest(request *http.Request, timeout time.Duration) ([]byte, error) {
	client := &http.Client{Timeout: timeout}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	result, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	if err = response.Body.Close(); err != nil {
		return nil, err
	}
	return result, nil
}

func (api *API) parseResponseBody(body []byte) (map[string]*json.RawMessage, error) {
	var result map[string]*json.RawMessage
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (api *API) checkIfSuccess(result map[string]*json.RawMessage) error {
	var ok bool
	if err := json.Unmarshal(*result["ok"], &ok); err != nil {
		return err
	}
	if ok == false {
		var errorMessage string
		if err := json.Unmarshal(*result["description"], &errorMessage); err != nil {
			return err
		}
		return errors.New(errorMessage)
	}
	return nil
}

func (api *API) execute(method string, args MethodArgs, timeout time.Duration) (*json.RawMessage, error) {
	request, err := api.buildRequest(method, args)
	if err != nil {
		return nil, NewBuildRequestError(err.Error())
	}
	body, err := api.sendRequest(request, timeout)
	if err != nil {
		return nil, NewSendRequestError(err.Error())
	}
	result, err := api.parseResponseBody(body)
	if err != nil {
		return nil, NewParseResponseBodyError(err.Error())
	}
	if err = api.checkIfSuccess(result); err != nil {
		return nil, NewAPIError(err.Error())
	}
	return result["result"], nil
}

type ChatID struct {
	ID       int
	Username string
}

func (cid *ChatID) MarshalJSON() ([]byte, error) {
	var value string
	if cid.ID != 0 {
		value = strconv.Itoa(cid.ID)
	} else {
		value = cid.Username
	}
	return []byte(value), nil
}

const (
	AllowedUpdateMessage            = "message"
	AllowedUpdateEditedMessage      = "edited_message"
	AllowedUpdateChannelPost        = "channel_post"
	AllowedUpdateEditedChannelPost  = "edited_channel_post"
	AllowedUpdateInlineQuery        = "inline_query"
	AllowedUpdateChosenInlineResult = "chosen_inline_result"
	AllowedUpdateCallbackQuery      = "callback_query"
	AllowedUpdateShippingQuery      = "shipping_query"
	AllowedUpdatePreCheckoutQuery   = "pre_checkout_query"
	AllowedUpdatePoll               = "poll"
)

type GetUpdatesArgs struct {
	Offset         int           `json:"offset,omitempty"`
	Limit          int           `json:"limit,omitempty"`
	Timeout        time.Duration `json:"timeout,omitempty"`
	AllowedUpdates []string      `json:"allowed_updates,omitempty"`
}

func (p *GetUpdatesArgs) GetRequestArgs() (*RequestArgs, error) {
	return buildJSONRequestArgs(p)
}

// https://core.telegram.org/bots/api#getupdates
func (api *API) GetUpdates(args *GetUpdatesArgs) (*[]*Update, error) {
	var update *[]*Update
	method := "getUpdates"
	var timeout time.Duration
	if args.Timeout > 0 {
		timeout = args.Timeout*time.Second - 1
	} else {
		timeout = 5 * time.Second
	}
	response, err := api.execute(method, args, timeout)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(*response, &update); err != nil {
		return nil, err
	}
	return update, nil
}

type SetWebhookArgs struct {
	URL               string     `json:"url"`
	MaxConnections    int        `json:"max_connections,omitempty"`
	AllowedUpdates    []string   `json:"allowed_updates,omitempty"`
	CertificateAsFile *InputFile `json:"-"`
}

func (p *SetWebhookArgs) GetRequestArgs() (*RequestArgs, error) {
	if p.CertificateAsFile.isAllSet() {
		args := marshallToMap(p)
		return buildMultipartRequestArgs(args, []*InputFile{p.CertificateAsFile})
	}
	return buildJSONRequestArgs(p)
}

// https://core.telegram.org/bots/api#setwebhook
func (api *API) SetWebhook(args *SetWebhookArgs) (*bool, error) {
	var success *bool
	method := "setWebhook"
	timeout := 5 * time.Second
	response, err := api.execute(method, args, timeout)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(*response, &success); err != nil {
		return nil, err
	}
	return success, nil
}

type DeleteWebhookArgs struct {
}

func (p *DeleteWebhookArgs) GetRequestArgs() (*RequestArgs, error) {
	return buildJSONRequestArgs(p)
}

// https://core.telegram.org/bots/api#deletewebhook
func (api *API) DeleteWebhook(args *DeleteWebhookArgs) (*bool, error) {
	var success *bool
	method := "deleteWebhook"
	timeout := 5 * time.Second
	response, err := api.execute(method, args, timeout)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(*response, &success); err != nil {
		return nil, err
	}
	return success, nil
}

type GetWebhookInfoArgs struct {
}

func (p *GetWebhookInfoArgs) GetRequestArgs() (*RequestArgs, error) {
	return buildJSONRequestArgs(p)
}

// https://core.telegram.org/bots/api#getwebhookinfo
func (api *API) GetWebhookInfo(args *GetWebhookInfoArgs) (*WebhookInfo, error) {
	var webhookInfo *WebhookInfo
	method := "getWebhookInfo"
	timeout := 5 * time.Second
	response, err := api.execute(method, args, timeout)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(*response, &webhookInfo); err != nil {
		return nil, err
	}
	return webhookInfo, nil
}

type GetMeArgs struct {
}

func (p *GetMeArgs) GetRequestArgs() (*RequestArgs, error) {
	return buildJSONRequestArgs(p)
}

// https://core.telegram.org/bots/api#getme
func (api *API) GetMe(args *GetMeArgs) (*User, error) {
	var user *User
	method := "getMe"
	timeout := 5 * time.Second
	response, err := api.execute(method, args, timeout)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(*response, &user); err != nil {
		return nil, err
	}
	return user, nil
}

const (
	ParseModeMarkdown = "Markdown"
	ParseModeHTML     = "HTML"
)

type SendMessageArgs struct {
	ChatID                *ChatID     `json:"chat_id"`
	Text                  string      `json:"text"`
	ParseMode             string      `json:"parse_mode,omitempty"`
	DisableWebPagePreview bool        `json:"disable_web_page_preview,omitempty"`
	DisableNotification   bool        `json:"disable_notification,omitempty"`
	ReplyToMessageID      int         `json:"reply_to_message_id,omitempty"`
	ReplyMarkup           interface{} `json:"reply_markup,omitempty"` // InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply
}

func (p *SendMessageArgs) GetRequestArgs() (*RequestArgs, error) {
	return buildJSONRequestArgs(p)
}

// https://core.telegram.org/bots/api#sendmessage
func (api *API) SendMessage(args *SendMessageArgs) (*Message, error) {
	var message *Message
	method := "sendMessage"
	timeout := 5 * time.Second
	response, err := api.execute(method, args, timeout)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(*response, &message); err != nil {
		return nil, err
	}
	return message, nil
}

type ForwardMessageArgs struct {
	ChatID              *ChatID `json:"chat_id"`
	FromChatID          *ChatID `json:"from_chat_id"`
	DisableNotification bool    `json:"disable_notification,omitempty"`
	MessageID           int     `json:"message_id"`
}

func (p *ForwardMessageArgs) GetRequestArgs() (*RequestArgs, error) {
	return buildJSONRequestArgs(p)
}

// https://core.telegram.org/bots/api#forwardmessage
func (api *API) ForwardMessage(args *ForwardMessageArgs) (*Message, error) {
	var message *Message
	method := "forwardMessage"
	timeout := 5 * time.Second
	response, err := api.execute(method, args, timeout)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(*response, &message); err != nil {
		return nil, err
	}
	return message, nil
}

type SendPhotoArgs struct {
	ChatID              *ChatID     `json:"chat_id"`
	Photo               string      `json:"photo"`
	Caption             string      `json:"caption,omitempty"`
	ParseMode           string      `json:"parse_mode,omitempty"`
	DisableNotification bool        `json:"disable_notification,omitempty"`
	ReplyToMessageID    int         `json:"reply_to_message_id,omitempty"`
	ReplyMarkup         interface{} `json:"reply_markup,omitempty"` // InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply
	PhotoAsFile         *InputFile  `json:"-"`
}

func (p *SendPhotoArgs) GetRequestArgs() (*RequestArgs, error) {
	if p.PhotoAsFile.isAllSet() {
		args := marshallToMap(p)
		return buildMultipartRequestArgs(args, []*InputFile{p.PhotoAsFile})
	}
	return buildJSONRequestArgs(p)

}

// https://core.telegram.org/bots/api#sendphoto
func (api *API) SendPhoto(args *SendPhotoArgs) (*Message, error) {
	var message *Message
	method := "sendPhoto"
	timeout := 5 * time.Second
	response, err := api.execute(method, args, timeout)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(*response, &message); err != nil {
		return nil, err
	}
	return message, nil
}

type SendAudioArgs struct {
	ChatID              *ChatID     `json:"chat_id"`
	Audio               string      `json:"audio"`
	Caption             string      `json:"caption,omitempty"`
	ParseMode           string      `json:"parse_mode,omitempty"`
	Duration            int         `json:"duration,omitempty"`
	Performer           string      `json:"performer,omitempty"`
	Title               string      `json:"title,omitempty"`
	Thumb               string      `json:"thumb,omitempty"`
	DisableNotification bool        `json:"disable_notification,omitempty"`
	ReplyToMessageID    int         `json:"reply_to_message_id,omitempty"`
	ReplyMarkup         interface{} `json:"reply_markup,omitempty"` // InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply
	AudioAsFile         *InputFile  `json:"-"`
	ThumbAsFile         *InputFile  `json:"-"`
}

func (p *SendAudioArgs) GetRequestArgs() (*RequestArgs, error) {
	if p.AudioAsFile.isAllSet() || p.ThumbAsFile.isAllSet() {
		args := marshallToMap(p)
		var files []*InputFile
		if p.AudioAsFile.isAllSet() {
			files = append(files, p.AudioAsFile)
		}
		if p.ThumbAsFile.isAllSet() {
			files = append(files, p.ThumbAsFile)
		}
		return buildMultipartRequestArgs(args, files)
	}
	return buildJSONRequestArgs(p)
}

// https://core.telegram.org/bots/api#sendaudio
func (api *API) SendAudio(args *SendAudioArgs) (*Message, error) {
	var message *Message
	method := "sendAudio"
	timeout := 5 * time.Second
	response, err := api.execute(method, args, timeout)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(*response, &message); err != nil {
		return nil, err
	}
	return message, nil
}

type SendDocumentArgs struct {
	ChatID              *ChatID     `json:"chat_id"`
	Document            string      `json:"document"`
	Thumb               string      `json:"thumb,omitempty"`
	Caption             string      `json:"caption,omitempty"`
	ParseMode           string      `json:"parse_mode,omitempty"`
	DisableNotification bool        `json:"disable_notification,omitempty"`
	ReplyToMessageID    int         `json:"reply_to_message_id,omitempty"`
	ReplyMarkup         interface{} `json:"reply_markup,omitempty"` // InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply
	DocumentAsFile      *InputFile  `json:"-"`
	ThumbAsFile         *InputFile  `json:"-"`
}

func (p *SendDocumentArgs) GetRequestArgs() (*RequestArgs, error) {
	if p.DocumentAsFile.isAllSet() || p.ThumbAsFile.isAllSet() {
		args := marshallToMap(p)
		var files []*InputFile
		if p.DocumentAsFile.isAllSet() {
			files = append(files, p.DocumentAsFile)
		}
		if p.ThumbAsFile.isAllSet() {
			files = append(files, p.ThumbAsFile)
		}
		return buildMultipartRequestArgs(args, files)
	}
	return buildJSONRequestArgs(p)
}

// https://core.telegram.org/bots/api#senddocument
func (api *API) SendDocument(args *SendDocumentArgs) (*Message, error) {
	var message *Message
	method := "sendDocument"
	timeout := 5 * time.Second
	response, err := api.execute(method, args, timeout)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(*response, &message); err != nil {
		return nil, err
	}
	return message, nil
}

type SendVideoArgs struct {
	ChatID              *ChatID     `json:"chat_id"`
	Video               string      `json:"video"`
	Duration            int         `json:"duration,omitempty"`
	Width               int         `json:"width,omitempty"`
	Height              int         `json:"height,omitempty"`
	Thumb               string      `json:"thumb,omitempty"`
	Caption             string      `json:"caption,omitempty"`
	ParseMode           string      `json:"parse_mode,omitempty"`
	SupportsStreaming   bool        `json:"supports_streaming,omitempty"`
	DisableNotification bool        `json:"disable_notification,omitempty"`
	ReplyToMessageID    int         `json:"reply_to_message_id,omitempty"`
	ReplyMarkup         interface{} `json:"reply_markup,omitempty"` // InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply
	VideoAsFile         *InputFile  `json:"-"`
	ThumbAsFile         *InputFile  `json:"-"`
}

func (p *SendVideoArgs) GetRequestArgs() (*RequestArgs, error) {
	if p.VideoAsFile.isAllSet() || p.ThumbAsFile.isAllSet() {
		args := marshallToMap(p)
		var files []*InputFile
		if p.VideoAsFile.isAllSet() {
			files = append(files, p.VideoAsFile)
		}
		if p.ThumbAsFile.isAllSet() {
			files = append(files, p.ThumbAsFile)
		}
		return buildMultipartRequestArgs(args, files)
	}
	return buildJSONRequestArgs(p)
}

// https://core.telegram.org/bots/api#sendvideo
func (api *API) SendVideo(args *SendVideoArgs) (*Message, error) {
	var message *Message
	method := "sendVideo"
	timeout := 5 * time.Second
	response, err := api.execute(method, args, timeout)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(*response, &message); err != nil {
		return nil, err
	}
	return message, nil
}

type SendAnimationArgs struct {
	ChatID              *ChatID     `json:"chat_id"`
	Animation           string      `json:"animation"`
	Duration            int         `json:"duration,omitempty"`
	Width               int         `json:"width,omitempty"`
	Height              int         `json:"height,omitempty"`
	Thumb               string      `json:"thumb,omitempty"`
	Caption             string      `json:"caption,omitempty"`
	ParseMode           string      `json:"parse_mode,omitempty"`
	DisableNotification bool        `json:"disable_notification,omitempty"`
	ReplyToMessageID    int         `json:"reply_to_message_id,omitempty"`
	ReplyMarkup         interface{} `json:"reply_markup,omitempty"` // InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply
	AnimationAsFile     *InputFile  `json:"-"`
	ThumbAsFile         *InputFile  `json:"-"`
}

func (p *SendAnimationArgs) GetRequestArgs() (*RequestArgs, error) {
	if p.AnimationAsFile.isAllSet() || p.ThumbAsFile.isAllSet() {
		args := marshallToMap(p)
		var files []*InputFile
		if p.AnimationAsFile.isAllSet() {
			files = append(files, p.AnimationAsFile)
		}
		if p.ThumbAsFile.isAllSet() {
			files = append(files, p.ThumbAsFile)
		}
		return buildMultipartRequestArgs(args, files)
	}
	return buildJSONRequestArgs(p)
}

// https://core.telegram.org/bots/api#sendanimation
func (api *API) SendAnimation(args *SendAnimationArgs) (*Message, error) {
	var message *Message
	method := "sendAnimation"
	timeout := 5 * time.Second
	response, err := api.execute(method, args, timeout)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(*response, &message); err != nil {
		return nil, err
	}
	return message, nil
}

type SendVoiceArgs struct {
	ChatID              *ChatID     `json:"chat_id"`
	Voice               string      `json:"voice"`
	Caption             string      `json:"caption,omitempty"`
	ParseMode           string      `json:"parse_mode,omitempty"`
	Duration            int         `json:"duration,omitempty"`
	DisableNotification bool        `json:"disable_notification,omitempty"`
	ReplyToMessageID    int         `json:"reply_to_message_id,omitempty"`
	ReplyMarkup         interface{} `json:"reply_markup,omitempty"` // InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply
	VoiceAsFile         *InputFile  `json:"-"`
}

func (p *SendVoiceArgs) GetRequestArgs() (*RequestArgs, error) {
	if p.VoiceAsFile.isAllSet() {
		args := marshallToMap(p)
		return buildMultipartRequestArgs(args, []*InputFile{p.VoiceAsFile})
	}
	return buildJSONRequestArgs(p)
}

// https://core.telegram.org/bots/api#sendvoice
func (api *API) SendVoice(args *SendVoiceArgs) (*Message, error) {
	var message *Message
	method := "sendVoice"
	timeout := 5 * time.Second
	response, err := api.execute(method, args, timeout)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(*response, &message); err != nil {
		return nil, err
	}
	return message, nil
}

type SendVideoNoteArgs struct {
	ChatID              *ChatID     `json:"chat_id"`
	VideoNote           InputFile   `json:"video_note"`
	Duration            int         `json:"duration,omitempty"`
	Length              int         `json:"length,omitempty"`
	Thumb               InputFile   `json:"thumb,omitempty"`
	DisableNotification bool        `json:"disable_notification,omitempty"`
	ReplyToMessageID    int         `json:"reply_to_message_id,omitempty"`
	ReplyMarkup         interface{} `json:"reply_markup,omitempty"` // InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply
	VideoNoteAsFile     *InputFile  `json:"-"`
	ThumbAsFile         *InputFile  `json:"-"`
}

func (p *SendVideoNoteArgs) GetRequestArgs() (*RequestArgs, error) {
	if p.VideoNoteAsFile.isAllSet() || p.ThumbAsFile.isAllSet() {
		args := marshallToMap(p)
		var files []*InputFile
		if p.VideoNoteAsFile.isAllSet() {
			files = append(files, p.VideoNoteAsFile)
		}
		if p.ThumbAsFile.isAllSet() {
			files = append(files, p.ThumbAsFile)
		}
		return buildMultipartRequestArgs(args, files)
	}
	return buildJSONRequestArgs(p)
}

// https://core.telegram.org/bots/api#sendvideoNote
func (api *API) SendVideoNote(args *SendVideoNoteArgs) (*Message, error) {
	var message *Message
	method := "sendVideoNote"
	timeout := 5 * time.Second
	response, err := api.execute(method, args, timeout)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(*response, &message); err != nil {
		return nil, err
	}
	return message, nil
}

type SendMediaGroupArgs struct {
	ChatID              *ChatID      `json:"chat_id"`
	Media               []InputMedia `json:"media"` // InputMediaPhoto and InputMediaVideo
	DisableNotification bool         `json:"disable_notification,omitempty"`
	ReplyToMessageID    int          `json:"reply_to_message_id,omitempty"`
}

func (p *SendMediaGroupArgs) GetRequestArgs() (*RequestArgs, error) {
	var files []*InputFile
	for _, media := range p.Media {
		for _, file := range media.getMedia() {
			if file.isAllSet() {
				files = append(files, file)
			}
		}
	}
	if len(files) > 0 {
		args := marshallToMap(p)
		return buildMultipartRequestArgs(args, files)
	}
	return buildJSONRequestArgs(p)
}

// https://core.telegram.org/bots/api#sendmediagroup
func (api *API) SendMediaGroup(args *SendMediaGroupArgs) (*[]*Message, error) {
	var messages *[]*Message
	method := "sendMediaGroup"
	timeout := 5 * time.Second
	response, err := api.execute(method, args, timeout)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(*response, &messages); err != nil {
		return nil, err
	}
	return messages, nil
}

type SendLocationArgs struct {
	ChatID              *ChatID     `json:"chat_id"`
	Latitude            float64     `json:"latitude"`
	Longitude           float64     `json:"longitude"`
	LivePeriod          int         `json:"live_period,omitempty"`
	DisableNotification bool        `json:"disable_notification,omitempty"`
	ReplyToMessageID    int         `json:"reply_to_message_id,omitempty"`
	ReplyMarkup         interface{} `json:"reply_markup,omitempty"` // InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply
}

func (p *SendLocationArgs) GetRequestArgs() (*RequestArgs, error) {
	return buildJSONRequestArgs(p)
}

// https://core.telegram.org/bots/api#sendlocation
func (api *API) SendLocation(args *SendLocationArgs) (*Message, error) {
	var message *Message
	method := "sendLocation"
	timeout := 5 * time.Second
	response, err := api.execute(method, args, timeout)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(*response, &message); err != nil {
		return nil, err
	}
	return message, nil
}

type EditMessageLiveLocationArgs struct {
	ChatID          *ChatID               `json:"chat_id,omitempty"`
	MessageID       int                   `json:"message_id,omitempty"`
	InlineMessageID string                `json:"inline_message_id,omitempty"`
	Latitude        float64               `json:"latitude"`
	Longitude       float64               `json:"longitude"`
	ReplyMarkup     *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

func (p *EditMessageLiveLocationArgs) GetRequestArgs() (*RequestArgs, error) {
	return buildJSONRequestArgs(p)
}

// https://core.telegram.org/bots/api#editmessagelivelocation
func (api *API) EditMessageLiveLocation(args *EditMessageLiveLocationArgs) (*OptionalMessage, error) {
	method := "editMessageLiveLocation"
	timeout := 5 * time.Second
	response, err := api.execute(method, args, timeout)
	if err != nil {
		return nil, err
	}
	return buildOptionalMessage(response)
}

type StopMessageLiveLocationArgs struct {
	ChatID          *ChatID               `json:"chat_id,omitempty"`
	MessageID       int                   `json:"message_id,omitempty"`
	InlineMessageID string                `json:"inline_message_id,omitempty"`
	ReplyMarkup     *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

func (p *StopMessageLiveLocationArgs) GetRequestArgs() (*RequestArgs, error) {
	return buildJSONRequestArgs(p)
}

// https://core.telegram.org/bots/api#stopmessagelivelocation
func (api *API) StopMessageLiveLocation(args *StopMessageLiveLocationArgs) (*OptionalMessage, error) {
	method := "stopMessageLiveLocation"
	timeout := 5 * time.Second
	response, err := api.execute(method, args, timeout)
	if err != nil {
		return nil, err
	}
	return buildOptionalMessage(response)
}

type SendVenueArgs struct {
	ChatID              *ChatID     `json:"chat_id"`
	Latitude            float64     `json:"latitude"`
	Longitude           float64     `json:"longitude"`
	Title               string      `json:"title"`
	Address             string      `json:"address"`
	FoursquareID        string      `json:"foursquare_id,omitempty"`
	FoursquareType      string      `json:"foursquare_type,omitempty"`
	DisableNotification bool        `json:"disable_notification,omitempty"`
	ReplyToMessageID    int         `json:"reply_to_message_id,omitempty"`
	ReplyMarkup         interface{} `json:"reply_markup,omitempty"` // InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply
}

func (p *SendVenueArgs) GetRequestArgs() (*RequestArgs, error) {
	return buildJSONRequestArgs(p)
}

// https://core.telegram.org/bots/api#sendvenue
func (api *API) SendVenue(args *SendVenueArgs) (*Message, error) {
	var message *Message
	method := "sendVenue"
	timeout := 5 * time.Second
	response, err := api.execute(method, args, timeout)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(*response, &message); err != nil {
		return nil, err
	}
	return message, nil
}

type SendContactArgs struct {
	ChatID              *ChatID     `json:"chat_id"`
	PhoneNumber         string      `json:"phone_number"`
	FirstName           string      `json:"first_name"`
	LastName            string      `json:"last_name,omitempty"`
	Vcard               string      `json:"vcard,omitempty"`
	DisableNotification bool        `json:"disable_notification,omitempty"`
	ReplyToMessageID    int         `json:"reply_to_message_id,omitempty"`
	ReplyMarkup         interface{} `json:"reply_markup,omitempty"` // InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply
}

func (p *SendContactArgs) GetRequestArgs() (*RequestArgs, error) {
	return buildJSONRequestArgs(p)
}

// https://core.telegram.org/bots/api#sendcontact
func (api *API) SendContact(args *SendContactArgs) (*Message, error) {
	var message *Message
	method := "sendContact"
	timeout := 5 * time.Second
	response, err := api.execute(method, args, timeout)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(*response, &message); err != nil {
		return nil, err
	}
	return message, nil
}

type SendPollArgs struct {
	ChatID              *ChatID     `json:"chat_id"`
	Question            string      `json:"question"`
	Options             []string    `json:"options"`
	DisableNotification bool        `json:"disable_notification,omitempty"`
	ReplyToMessageID    int         `json:"reply_to_message_id,omitempty"`
	ReplyMarkup         interface{} `json:"reply_markup,omitempty"` // InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply
}

func (p *SendPollArgs) GetRequestArgs() (*RequestArgs, error) {
	return buildJSONRequestArgs(p)
}

// https://core.telegram.org/bots/api#sendpoll
func (api *API) SendPoll(args *SendPollArgs) (*Message, error) {
	var message *Message
	method := "sendPoll"
	timeout := 5 * time.Second
	response, err := api.execute(method, args, timeout)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(*response, &message); err != nil {
		return nil, err
	}
	return message, nil
}

const (
	ChatActionTyping          = "typing"
	ChatActionUploadPhoto     = "upload_photo"
	ChatActionRecordVideo     = "record_video"
	ChatActionUploadVideo     = "upload_video"
	ChatActionRecordAudio     = "record_audio"
	ChatActionUploadAudio     = "upload_audio"
	ChatActionUploadDocument  = "upload_document"
	ChatActionFindLocation    = "find_location"
	ChatActionRecordVideoNote = "record_video_note"
	ChatActionUploadVideoNote = "upload_video_note"
)

type SendChatActionArgs struct {
	ChatID *ChatID `json:"chat_id"`
	Action string  `json:"action"`
}

func (p *SendChatActionArgs) GetRequestArgs() (*RequestArgs, error) {
	return buildJSONRequestArgs(p)
}

// https://core.telegram.org/bots/api#sendchataction
func (api *API) SendChatAction(args *SendChatActionArgs) (*bool, error) {
	var success *bool
	method := "sendChatAction"
	timeout := 5 * time.Second
	response, err := api.execute(method, args, timeout)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(*response, &success); err != nil {
		return nil, err
	}
	return success, nil
}

type GetUserProfilePhotosArgs struct {
	UserID int `json:"user_id"`
	Offset int `json:"offset,omitempty"`
	Limit  int `json:"limit,omitempty"`
}

func (p *GetUserProfilePhotosArgs) GetRequestArgs() (*RequestArgs, error) {
	return buildJSONRequestArgs(p)
}

// https://core.telegram.org/bots/api#getuserprofilephotos
func (api *API) GetUserProfilePhotos(args *GetUserProfilePhotosArgs) (*UserProfilePhotos, error) {
	var userProfilePhotos *UserProfilePhotos
	method := "getUserProfilePhotos"
	timeout := 5 * time.Second
	response, err := api.execute(method, args, timeout)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(*response, &userProfilePhotos); err != nil {
		return nil, err
	}
	return userProfilePhotos, nil
}

type GetFileArgs struct {
	FileID string `json:"file_id"`
}

func (p *GetFileArgs) GetRequestArgs() (*RequestArgs, error) {
	return buildJSONRequestArgs(p)
}

// https://core.telegram.org/bots/api#getfile
func (api *API) GetFile(args *GetFileArgs) (*File, error) {
	var file *File
	method := "getFile"
	timeout := 5 * time.Second
	response, err := api.execute(method, args, timeout)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(*response, &file); err != nil {
		return nil, err
	}
	return file, nil
}

type KickChatMemberArgs struct {
	ChatID    *ChatID `json:"chat_id"`
	UserID    int     `json:"user_id"`
	UntilDate int     `json:"until_date,omitempty"`
}

func (p *KickChatMemberArgs) GetRequestArgs() (*RequestArgs, error) {
	return buildJSONRequestArgs(p)
}

// https://core.telegram.org/bots/api#kickchatmember
func (api *API) KickChatMember(args *KickChatMemberArgs) (*bool, error) {
	var success *bool
	method := "kickChatMember"
	timeout := 5 * time.Second
	response, err := api.execute(method, args, timeout)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(*response, &success); err != nil {
		return nil, err
	}
	return success, nil
}

type UnbanChatMemberArgs struct {
	ChatID *ChatID `json:"chat_id"`
	UserID int     `json:"user_id"`
}

func (p *UnbanChatMemberArgs) GetRequestArgs() (*RequestArgs, error) {
	return buildJSONRequestArgs(p)
}

// https://core.telegram.org/bots/api#unbanchatmember
func (api *API) UnbanChatMember(args *UnbanChatMemberArgs) (*bool, error) {
	var success *bool
	method := "unbanChatMember"
	timeout := 5 * time.Second
	response, err := api.execute(method, args, timeout)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(*response, &success); err != nil {
		return nil, err
	}
	return success, nil
}

type RestrictChatMemberArgs struct {
	ChatID      *ChatID          `json:"chat_id"`
	UserID      int              `json:"user_id"`
	Permissions *ChatPermissions `json:"permissions"`
	UntilDate   int              `json:"until_date,omitempty"`
}

func (p *RestrictChatMemberArgs) GetRequestArgs() (*RequestArgs, error) {
	return buildJSONRequestArgs(p)
}

// https://core.telegram.org/bots/api#restrictchatmember
func (api *API) RestrictChatMember(args *RestrictChatMemberArgs) (*bool, error) {
	var success *bool
	method := "restrictChatMember"
	timeout := 5 * time.Second
	response, err := api.execute(method, args, timeout)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(*response, &success); err != nil {
		return nil, err
	}
	return success, nil
}

type PromoteChatMemberArgs struct {
	ChatID             *ChatID `json:"chat_id"`
	UserID             int     `json:"user_id"`
	CanChangeInfo      bool    `json:"can_change_info,omitempty"`
	CanPostMessages    bool    `json:"can_post_messages,omitempty"`
	CanEditMessages    bool    `json:"can_edit_messages,omitempty"`
	CanDeleteMessages  bool    `json:"can_delete_messages,omitempty"`
	CanInviteUsers     bool    `json:"can_invite_users,omitempty"`
	CanRestrictMembers bool    `json:"can_restrict_members,omitempty"`
	CanPinMessages     bool    `json:"can_pin_messages,omitempty"`
	CanPromoteMembers  bool    `json:"can_promote_members,omitempty"`
}

func (p *PromoteChatMemberArgs) GetRequestArgs() (*RequestArgs, error) {
	return buildJSONRequestArgs(p)
}

// https://core.telegram.org/bots/api#promotechatmember
func (api *API) PromoteChatMember(args *PromoteChatMemberArgs) (*bool, error) {
	var success *bool
	method := "promoteChatMember"
	timeout := 5 * time.Second
	response, err := api.execute(method, args, timeout)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(*response, &success); err != nil {
		return nil, err
	}
	return success, nil
}

type SetChatPermissionsArgs struct {
	ChatID      *ChatID         `json:"chat_id"`
	Permissions ChatPermissions `json:"permissions"`
}

func (p *SetChatPermissionsArgs) GetRequestArgs() (*RequestArgs, error) {
	return buildJSONRequestArgs(p)
}

// https://core.telegram.org/bots/api#setchatpermissions
func (api *API) SetChatPermissions(args *SetChatPermissionsArgs) (*bool, error) {
	var success *bool
	method := "setChatPermissions"
	timeout := 5 * time.Second
	response, err := api.execute(method, args, timeout)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(*response, &success); err != nil {
		return nil, err
	}
	return success, nil
}

type ExportChatInviteLinkArgs struct {
	ChatID *ChatID `json:"chat_id"`
}

func (p *ExportChatInviteLinkArgs) GetRequestArgs() (*RequestArgs, error) {
	return buildJSONRequestArgs(p)
}

// https://core.telegram.org/bots/api#exportchatinvitelink
func (api *API) ExportChatInviteLink(args *ExportChatInviteLinkArgs) (*string, error) {
	var inviteLink *string
	method := "exportChatInviteLink"
	timeout := 5 * time.Second
	response, err := api.execute(method, args, timeout)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(*response, &inviteLink); err != nil {
		return nil, err
	}
	return inviteLink, nil
}

type SetChatPhotoArgs struct {
	ChatID      *ChatID    `json:"chat_id"`
	Photo       string     `json:"photo"`
	PhotoAsFile *InputFile `json:"-"`
}

func (p *SetChatPhotoArgs) GetRequestArgs() (*RequestArgs, error) {
	if p.PhotoAsFile.isAllSet() {
		args := marshallToMap(p)
		return buildMultipartRequestArgs(args, []*InputFile{p.PhotoAsFile})
	}
	return buildJSONRequestArgs(p)
}

// https://core.telegram.org/bots/api#setchatphoto
func (api *API) SetChatPhoto(args *SetChatPhotoArgs) (*bool, error) {
	var success *bool
	method := "setChatPhoto"
	timeout := 5 * time.Second
	response, err := api.execute(method, args, timeout)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(*response, &success); err != nil {
		return nil, err
	}
	return success, nil
}

type DeleteChatPhotoArgs struct {
	ChatID *ChatID `json:"chat_id"`
}

func (p *DeleteChatPhotoArgs) GetRequestArgs() (*RequestArgs, error) {
	return buildJSONRequestArgs(p)
}

// https://core.telegram.org/bots/api#deletechatPhoto
func (api *API) DeleteChatPhoto(args *DeleteChatPhotoArgs) (*bool, error) {
	var success *bool
	method := "deleteChatPhoto"
	timeout := 5 * time.Second
	response, err := api.execute(method, args, timeout)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(*response, &success); err != nil {
		return nil, err
	}
	return success, nil
}

type SetChatTitleArgs struct {
	ChatID *ChatID `json:"chat_id"`
	Title  string  `json:"title"`
}

func (p *SetChatTitleArgs) GetRequestArgs() (*RequestArgs, error) {
	return buildJSONRequestArgs(p)
}

// https://core.telegram.org/bots/api#setchattitle
func (api *API) SetChatTitle(args *SetChatTitleArgs) (*bool, error) {
	var success *bool
	method := "setChatTitle"
	timeout := 5 * time.Second
	response, err := api.execute(method, args, timeout)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(*response, &success); err != nil {
		return nil, err
	}
	return success, nil
}

type SetChatDescriptionArgs struct {
	ChatID      *ChatID `json:"chat_id"`
	Description string  `json:"description,omitempty"`
}

func (p *SetChatDescriptionArgs) GetRequestArgs() (*RequestArgs, error) {
	return buildJSONRequestArgs(p)
}

// https://core.telegram.org/bots/api#setchatdescription
func (api *API) SetChatDescription(args *SetChatDescriptionArgs) (*bool, error) {
	var success *bool
	method := "setChatDescription"
	timeout := 5 * time.Second
	response, err := api.execute(method, args, timeout)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(*response, &success); err != nil {
		return nil, err
	}
	return success, nil
}

type PinChatMessageArgs struct {
	ChatID              *ChatID `json:"chat_id"`
	MessageID           int     `json:"message_id"`
	DisableNotification bool    `json:"disable_notification,omitempty"`
}

func (p *PinChatMessageArgs) GetRequestArgs() (*RequestArgs, error) {
	return buildJSONRequestArgs(p)
}

// https://core.telegram.org/bots/api#pinchatmessage
func (api *API) PinChatMessage(args *PinChatMessageArgs) (*bool, error) {
	var success *bool
	method := "pinChatMessage"
	timeout := 5 * time.Second
	response, err := api.execute(method, args, timeout)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(*response, &success); err != nil {
		return nil, err
	}
	return success, nil
}

type UnpinChatMessageArgs struct {
	ChatID *ChatID `json:"chat_id"`
}

func (p *UnpinChatMessageArgs) GetRequestArgs() (*RequestArgs, error) {
	return buildJSONRequestArgs(p)
}

// https://core.telegram.org/bots/api#unpinchatmessage
func (api *API) UnpinChatMessage(args *UnpinChatMessageArgs) (*bool, error) {
	var success *bool
	method := "unpinChatMessage"
	timeout := 5 * time.Second
	response, err := api.execute(method, args, timeout)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(*response, &success); err != nil {
		return nil, err
	}
	return success, nil
}

type LeaveChatArgs struct {
	ChatID *ChatID `json:"chat_id"`
}

func (p *LeaveChatArgs) GetRequestArgs() (*RequestArgs, error) {
	return buildJSONRequestArgs(p)
}

// https://core.telegram.org/bots/api#leavechat
func (api *API) LeaveChat(args *LeaveChatArgs) (*bool, error) {
	var success *bool
	method := "leaveChat"
	timeout := 5 * time.Second
	response, err := api.execute(method, args, timeout)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(*response, &success); err != nil {
		return nil, err
	}
	return success, nil
}

type GetChatArgs struct {
	ChatID *ChatID `json:"chat_id"`
}

func (p *GetChatArgs) GetRequestArgs() (*RequestArgs, error) {
	return buildJSONRequestArgs(p)
}

// https://core.telegram.org/bots/api#getchat
func (api *API) GetChat(args *GetChatArgs) (*Chat, error) {
	var chat *Chat
	method := "getChat"
	timeout := 5 * time.Second
	response, err := api.execute(method, args, timeout)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(*response, &chat); err != nil {
		return nil, err
	}
	return chat, nil
}

type GetChatAdministratorsArgs struct {
	ChatID *ChatID `json:"chat_id"`
}

func (p *GetChatAdministratorsArgs) GetRequestArgs() (*RequestArgs, error) {
	return buildJSONRequestArgs(p)
}

// https://core.telegram.org/bots/api#getchatadministrators
func (api *API) GetChatAdministrators(args *GetChatAdministratorsArgs) (*[]*ChatMember, error) {
	var chatMembers *[]*ChatMember
	method := "getChatAdministrators"
	timeout := 5 * time.Second
	response, err := api.execute(method, args, timeout)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(*response, &chatMembers); err != nil {
		return nil, err
	}
	return chatMembers, nil
}

type GetChatMembersCountArgs struct {
	ChatID *ChatID `json:"chat_id"`
}

func (p *GetChatMembersCountArgs) GetRequestArgs() (*RequestArgs, error) {
	return buildJSONRequestArgs(p)
}

// https://core.telegram.org/bots/api#getchatmemberscount
func (api *API) GetChatMembersCount(args *GetChatMembersCountArgs) (*int, error) {
	var count *int
	method := "getChatMembersCount"
	timeout := 5 * time.Second
	response, err := api.execute(method, args, timeout)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(*response, &count); err != nil {
		return nil, err
	}
	return count, nil
}

type GetChatMemberArgs struct {
	ChatID *ChatID `json:"chat_id"`
	UserID int     `json:"user_id"`
}

func (p *GetChatMemberArgs) GetRequestArgs() (*RequestArgs, error) {
	return buildJSONRequestArgs(p)
}

// https://core.telegram.org/bots/api#getchatmember
func (api *API) GetChatMember(args *GetChatMemberArgs) (*ChatMember, error) {
	var chatMember *ChatMember
	method := "getChatMember"
	timeout := 5 * time.Second
	response, err := api.execute(method, args, timeout)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(*response, &chatMember); err != nil {
		return nil, err
	}
	return chatMember, nil
}

type SetChatStickerSetArgs struct {
	ChatID         *ChatID `json:"chat_id"`
	StickerSetName string  `json:"sticker_set_name"`
}

func (p *SetChatStickerSetArgs) GetRequestArgs() (*RequestArgs, error) {
	return buildJSONRequestArgs(p)
}

// https://core.telegram.org/bots/api#setchatstickerset
func (api *API) SetChatStickerSet(args *SetChatStickerSetArgs) (*bool, error) {
	var success *bool
	method := "setChatStickerSet"
	timeout := 5 * time.Second
	response, err := api.execute(method, args, timeout)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(*response, &success); err != nil {
		return nil, err
	}
	return success, nil
}

type DeleteChatStickerSetArgs struct {
	ChatID *ChatID `json:"chat_id"`
}

func (p *DeleteChatStickerSetArgs) GetRequestArgs() (*RequestArgs, error) {
	return buildJSONRequestArgs(p)
}

// https://core.telegram.org/bots/api#deletechatstickerset
func (api *API) DeleteChatStickerSet(args *DeleteChatStickerSetArgs) (*bool, error) {
	var success *bool
	method := "deleteChatStickerSet"
	timeout := 5 * time.Second
	response, err := api.execute(method, args, timeout)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(*response, &success); err != nil {
		return nil, err
	}
	return success, nil
}

type AnswerCallbackQueryArgs struct {
	CallbackQueryID string `json:"callback_query_id"`
	Text            string `json:"text,omitempty"`
	ShowAlert       bool   `json:"show_alert,omitempty"`
	URL             string `json:"url,omitempty"`
	CacheTime       int    `json:"cache_time,omitempty"`
}

func (p *AnswerCallbackQueryArgs) GetRequestArgs() (*RequestArgs, error) {
	return buildJSONRequestArgs(p)
}

// https://core.telegram.org/bots/api#answercallbackquery
func (api *API) AnswerCallbackQuery(args *AnswerCallbackQueryArgs) (*bool, error) {
	var success *bool
	method := "answerCallbackQuery"
	timeout := 5 * time.Second
	response, err := api.execute(method, args, timeout)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(*response, &success); err != nil {
		return nil, err
	}
	return success, nil
}

type EditMessageTextArgs struct {
	ChatID                *ChatID               `json:"chat_id,omitempty"`
	MessageID             int                   `json:"message_id,omitempty"`
	InlineMessageID       string                `json:"inline_message_id,omitempty"`
	Text                  string                `json:"text"`
	ParseMode             string                `json:"parse_mode,omitempty"`
	DisableWebPagePreview bool                  `json:"disable_web_page_preview,omitempty"`
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

func (p *EditMessageTextArgs) GetRequestArgs() (*RequestArgs, error) {
	return buildJSONRequestArgs(p)
}

// https://core.telegram.org/bots/api#editmessagetext
func (api *API) EditMessageText(args *EditMessageTextArgs) (*OptionalMessage, error) {
	method := "editMessageText"
	timeout := 5 * time.Second
	response, err := api.execute(method, args, timeout)
	if err != nil {
		return nil, err
	}
	return buildOptionalMessage(response)
}

type EditMessageCaptionArgs struct {
	ChatID          *ChatID               `json:"chat_id,omitempty"`
	MessageID       int                   `json:"message_id,omitempty"`
	InlineMessageID string                `json:"inline_message_id,omitempty"`
	Caption         string                `json:"caption,omitempty"`
	ParseMode       string                `json:"parse_mode,omitempty"`
	ReplyMarkup     *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

func (p *EditMessageCaptionArgs) GetRequestArgs() (*RequestArgs, error) {
	return buildJSONRequestArgs(p)
}

// https://core.telegram.org/bots/api#editmessagecaption
func (api *API) EditMessageCaption(args *EditMessageCaptionArgs) (*OptionalMessage, error) {
	method := "editMessageCaption"
	timeout := 5 * time.Second
	response, err := api.execute(method, args, timeout)
	if err != nil {
		return nil, err
	}
	return buildOptionalMessage(response)
}

type EditMessageMediaArgs struct {
	ChatID          *ChatID               `json:"chat_id,omitempty"`
	MessageID       int                   `json:"message_id,omitempty"`
	InlineMessageID string                `json:"inline_message_id,omitempty"`
	Media           *InputMedia           `json:"media"`
	ReplyMarkup     *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

func (p *EditMessageMediaArgs) GetRequestArgs() (*RequestArgs, error) {
	var files []*InputFile
	media := *p.Media
	for _, file := range media.getMedia() {
		if file.isAllSet() {
			files = append(files, file)
		}
	}
	if len(files) > 0 {
		args := marshallToMap(p)
		return buildMultipartRequestArgs(args, files)
	}
	return buildJSONRequestArgs(p)
}

// https://core.telegram.org/bots/api#editmessagemedia
func (api *API) EditMessageMedia(args *EditMessageMediaArgs) (*OptionalMessage, error) {
	method := "editMessageMedia"
	timeout := 5 * time.Second
	response, err := api.execute(method, args, timeout)
	if err != nil {
		return nil, err
	}
	return buildOptionalMessage(response)
}

type EditMessageReplyMarkupArgs struct {
	ChatID          *ChatID               `json:"chat_id,omitempty"`
	MessageID       int                   `json:"message_id,omitempty"`
	InlineMessageID string                `json:"inline_message_id,omitempty"`
	ReplyMarkup     *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

func (p *EditMessageReplyMarkupArgs) GetRequestArgs() (*RequestArgs, error) {
	return buildJSONRequestArgs(p)
}

// https://core.telegram.org/bots/api#editmessagereplymarkup
func (api *API) EditMessageReplyMarkup(args *EditMessageReplyMarkupArgs) (*OptionalMessage, error) {
	method := "editMessageReplyMarkup"
	timeout := 5 * time.Second
	response, err := api.execute(method, args, timeout)
	if err != nil {
		return nil, err
	}
	return buildOptionalMessage(response)
}

type StopPollArgs struct {
	ChatID      *ChatID               `json:"chat_id"`
	MessageID   int                   `json:"message_id"`
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

func (p *StopPollArgs) GetRequestArgs() (*RequestArgs, error) {
	return buildJSONRequestArgs(p)
}

// https://core.telegram.org/bots/api#stoppoll
func (api *API) StopPoll(args *StopPollArgs) (*Poll, error) {
	var poll *Poll
	method := "stopPoll"
	timeout := 5 * time.Second
	response, err := api.execute(method, args, timeout)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(*response, &poll); err != nil {
		return nil, err
	}
	return poll, nil
}

type DeleteMessageArgs struct {
	ChatID    *ChatID `json:"chat_id"`
	MessageID int     `json:"message_id"`
}

func (p *DeleteMessageArgs) GetRequestArgs() (*RequestArgs, error) {
	return buildJSONRequestArgs(p)
}

// https://core.telegram.org/bots/api#deletemessage
func (api *API) DeleteMessage(args *DeleteMessageArgs) (*bool, error) {
	var success *bool
	method := "deleteMessage"
	timeout := 5 * time.Second
	response, err := api.execute(method, args, timeout)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(*response, &success); err != nil {
		return nil, err
	}
	return success, nil
}

type SendStickerArgs struct {
	ChatID              *ChatID     `json:"chat_id"`
	Sticker             string      `json:"sticker"`
	DisableNotification bool        `json:"disable_notification,omitempty"`
	ReplyToMessageID    int         `json:"reply_to_message_id,omitempty"`
	ReplyMarkup         interface{} `json:"reply_markup,omitempty"` // InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply
	StickerAsFile       *InputFile  `json:"-"`
}

func (p *SendStickerArgs) GetRequestArgs() (*RequestArgs, error) {
	if p.StickerAsFile.isAllSet() {
		args := marshallToMap(p)
		return buildMultipartRequestArgs(args, []*InputFile{p.StickerAsFile})
	}
	return buildJSONRequestArgs(p)
}

// https://core.telegram.org/bots/api#sendsticker
func (api *API) SendSticker(args *SendStickerArgs) (*Message, error) {
	var message *Message
	method := "sendSticker"
	timeout := 5 * time.Second
	response, err := api.execute(method, args, timeout)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(*response, &message); err != nil {
		return nil, err
	}
	return message, nil
}

type GetStickerSetArgs struct {
	Name string `json:"name"`
}

func (p *GetStickerSetArgs) GetRequestArgs() (*RequestArgs, error) {
	return buildJSONRequestArgs(p)
}

// https://core.telegram.org/bots/api#getstickerset
func (api *API) GetStickerSet(args *GetStickerSetArgs) (*StickerSet, error) {
	var stickerSet *StickerSet
	method := "getStickerSet"
	timeout := 5 * time.Second
	response, err := api.execute(method, args, timeout)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(*response, &stickerSet); err != nil {
		return nil, err
	}
	return stickerSet, nil
}

type UploadStickerFileArgs struct {
	UserID           int        `json:"user_id"`
	PngSticker       string     `json:"png_sticker"`
	PngStickerAsFile *InputFile `json:"-"`
}

func (p *UploadStickerFileArgs) GetRequestArgs() (*RequestArgs, error) {
	if p.PngStickerAsFile.isAllSet() {
		args := marshallToMap(p)
		return buildMultipartRequestArgs(args, []*InputFile{p.PngStickerAsFile})
	}
	return buildJSONRequestArgs(p)
}

// https://core.telegram.org/bots/api#uploadstickerfile
func (api *API) UploadStickerFile(args *UploadStickerFileArgs) (*File, error) {
	var file *File
	method := "uploadStickerFile"
	timeout := 5 * time.Second
	response, err := api.execute(method, args, timeout)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(*response, &file); err != nil {
		return nil, err
	}
	return file, nil
}

type CreateNewStickerSetArgs struct {
	UserID           int           `json:"user_id"`
	Name             string        `json:"name"`
	Title            string        `json:"title"`
	PngSticker       string        `json:"png_sticker"`
	Emojis           string        `json:"emojis"`
	ContainsMasks    bool          `json:"contains_masks,omitempty"`
	MaskPosition     *MaskPosition `json:"mask_position,omitempty"`
	PngStickerAsFile *InputFile    `json:"-"`
}

func (p *CreateNewStickerSetArgs) GetRequestArgs() (*RequestArgs, error) {
	if p.PngStickerAsFile.isAllSet() {
		args := marshallToMap(p)
		return buildMultipartRequestArgs(args, []*InputFile{p.PngStickerAsFile})
	}
	return buildJSONRequestArgs(p)
}

// https://core.telegram.org/bots/api#createnewstickerset
func (api *API) CreateNewStickerSet(args *CreateNewStickerSetArgs) (*bool, error) {
	var success *bool
	method := "createNewStickerSet"
	timeout := 5 * time.Second
	response, err := api.execute(method, args, timeout)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(*response, &success); err != nil {
		return nil, err
	}
	return success, nil
}

type AddStickerToSetArgs struct {
	UserID           int           `json:"user_id"`
	Name             string        `json:"name"`
	PngSticker       string        `json:"png_sticker"`
	Emojis           string        `json:"emojis"`
	MaskPosition     *MaskPosition `json:"mask_position,omitempty"`
	PngStickerAsFile *InputFile    `json:"-"`
}

func (p *AddStickerToSetArgs) GetRequestArgs() (*RequestArgs, error) {
	if p.PngStickerAsFile.isAllSet() {
		args := marshallToMap(p)
		return buildMultipartRequestArgs(args, []*InputFile{p.PngStickerAsFile})
	}
	return buildJSONRequestArgs(p)
}

// https://core.telegram.org/bots/api#addstickertoset
func (api *API) AddStickerToSet(args *AddStickerToSetArgs) (*bool, error) {
	var success *bool
	method := "addStickerToSet"
	timeout := 5 * time.Second
	response, err := api.execute(method, args, timeout)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(*response, &success); err != nil {
		return nil, err
	}
	return success, nil
}

type SetStickerPositionInSetArgs struct {
	Sticker  string `json:"sticker"`
	Position int    `json:"position"`
}

func (p *SetStickerPositionInSetArgs) GetRequestArgs() (*RequestArgs, error) {
	return buildJSONRequestArgs(p)
}

// https://core.telegram.org/bots/api#setstickerpositioninset
func (api *API) SetStickerPositionInSet(args *SetStickerPositionInSetArgs) (*bool, error) {
	var success *bool
	method := "setStickerPositionInSet"
	timeout := 5 * time.Second
	response, err := api.execute(method, args, timeout)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(*response, &success); err != nil {
		return nil, err
	}
	return success, nil
}

type DeleteStickerFromSetArgs struct {
	Sticker string `json:"sticker"`
}

func (p *DeleteStickerFromSetArgs) GetRequestArgs() (*RequestArgs, error) {
	return buildJSONRequestArgs(p)
}

// https://core.telegram.org/bots/api#deletestickerfromset
func (api *API) DeleteStickerFromSet(args *DeleteStickerFromSetArgs) (*bool, error) {
	var success *bool
	method := "deleteStickerFromSet"
	timeout := 5 * time.Second
	response, err := api.execute(method, args, timeout)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(*response, &success); err != nil {
		return nil, err
	}
	return success, nil
}

type AnswerInlineQueryArgs struct {
	InlineQueryID     string      `json:"inline_query_id"`
	Results           interface{} `json:"results"` // InlineQueryResult
	CacheTime         int         `json:"cache_time,omitempty"`
	IsPersonal        bool        `json:"is_personal,omitempty"`
	NextOffset        string      `json:"next_offset,omitempty"`
	SwitchPmText      string      `json:"switch_pm_text,omitempty"`
	SwitchPmParameter string      `json:"switch_pm_parameter,omitempty"`
}

func (p *AnswerInlineQueryArgs) GetRequestArgs() (*RequestArgs, error) {
	return buildJSONRequestArgs(p)
}

// https://core.telegram.org/bots/api#answerinlinequery
func (api *API) AnswerInlineQuery(args *AnswerInlineQueryArgs) (*bool, error) {
	var success *bool
	method := "answerInlineQuery"
	timeout := 5 * time.Second
	response, err := api.execute(method, args, timeout)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(*response, &success); err != nil {
		return nil, err
	}
	return success, nil
}

type SendInvoiceArgs struct {
	ChatID                    *ChatID               `json:"chat_id"`
	Title                     string                `json:"title"`
	Description               string                `json:"description"`
	Payload                   string                `json:"payload"`
	ProviderToken             string                `json:"provider_token"`
	StartParameter            string                `json:"start_parameter"`
	Currency                  string                `json:"currency"`
	Prices                    []*LabeledPrice       `json:"prices"`
	ProviderData              string                `json:"provider_data,omitempty"`
	PhotoURL                  string                `json:"photo_url,omitempty"`
	PhotoSize                 int                   `json:"photo_size,omitempty"`
	PhotoWidth                int                   `json:"photo_width,omitempty"`
	PhotoHeight               int                   `json:"photo_height,omitempty"`
	NeedName                  bool                  `json:"need_name,omitempty"`
	NeedPhoneNumber           bool                  `json:"need_phone_number,omitempty"`
	NeedEmail                 bool                  `json:"need_email,omitempty"`
	NeedShippingAddress       bool                  `json:"need_shipping_address,omitempty"`
	SendPhoneNumberToProvider bool                  `json:"send_phone_number_to_provider,omitempty"`
	SendEmailToProvider       bool                  `json:"send_email_to_provider,omitempty"`
	IsFlexible                bool                  `json:"is_flexible,omitempty"`
	DisableNotification       bool                  `json:"disable_notification,omitempty"`
	ReplyToMessageID          int                   `json:"reply_to_message_id,omitempty"`
	ReplyMarkup               *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

func (p *SendInvoiceArgs) GetRequestArgs() (*RequestArgs, error) {
	return buildJSONRequestArgs(p)
}

// https://core.telegram.org/bots/api#sendinvoice
func (api *API) SendInvoice(args *SendInvoiceArgs) (*Message, error) {
	var message *Message
	method := "sendInvoice"
	timeout := 5 * time.Second
	response, err := api.execute(method, args, timeout)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(*response, &message); err != nil {
		return nil, err
	}
	return message, nil
}

type AnswerShippingQueryArgs struct {
	ShippingQueryID string            `json:"shipping_query_id"`
	Ok              bool              `json:"ok"`
	ShippingOptions []*ShippingOption `json:"shipping_options,omitempty"`
	ErrorMessage    string            `json:"error_message,omitempty"`
}

func (p *AnswerShippingQueryArgs) GetRequestArgs() (*RequestArgs, error) {
	return buildJSONRequestArgs(p)
}

// https://core.telegram.org/bots/api#answershippingquery
func (api *API) AnswerShippingQuery(args *AnswerShippingQueryArgs) (*bool, error) {
	var success *bool
	method := "answerShippingQuery"
	timeout := 5 * time.Second
	response, err := api.execute(method, args, timeout)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(*response, &success); err != nil {
		return nil, err
	}
	return success, nil
}

type AnswerPreCheckoutQueryArgs struct {
	PreCheckoutQueryID string `json:"pre_checkout_query_id"`
	Ok                 bool   `json:"ok"`
	ErrorMessage       string `json:"error_message,omitempty"`
}

func (p *AnswerPreCheckoutQueryArgs) GetRequestArgs() (*RequestArgs, error) {
	return buildJSONRequestArgs(p)
}

// https://core.telegram.org/bots/api#answerprecheckoutquery
func (api *API) AnswerPreCheckoutQuery(args *AnswerPreCheckoutQueryArgs) (*bool, error) {
	var success *bool
	method := "answerPreCheckoutQuery"
	timeout := 5 * time.Second
	response, err := api.execute(method, args, timeout)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(*response, &success); err != nil {
		return nil, err
	}
	return success, nil
}

type SetPassportDataErrorsArgs struct {
	UserID int           `json:"user_id"`
	Errors []interface{} `json:"errors"` // PassportElementError
}

func (p *SetPassportDataErrorsArgs) GetRequestArgs() (*RequestArgs, error) {
	return buildJSONRequestArgs(p)
}

// https://core.telegram.org/bots/api#setpassportdataerrors
func (api *API) SetPassportDataErrors(args *SetPassportDataErrorsArgs) (*bool, error) {
	var success *bool
	method := "setPassportDataErrors"
	timeout := 5 * time.Second
	response, err := api.execute(method, args, timeout)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(*response, &success); err != nil {
		return nil, err
	}
	return success, nil
}

type SendGameArgs struct {
	ChatID              *ChatID               `json:"chat_id"`
	GameShortName       string                `json:"game_short_name"`
	DisableNotification bool                  `json:"disable_notification,omitempty"`
	ReplyToMessageID    int                   `json:"reply_to_message_id,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

func (p *SendGameArgs) GetRequestArgs() (*RequestArgs, error) {
	return buildJSONRequestArgs(p)
}

// https://core.telegram.org/bots/api#sendgame
func (api *API) SendGame(args *SendGameArgs) (*Message, error) {
	var message *Message
	method := "sendGame"
	timeout := 5 * time.Second
	response, err := api.execute(method, args, timeout)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(*response, &message); err != nil {
		return nil, err
	}
	return message, nil
}

type SetGameScoreArgs struct {
	UserID             int     `json:"user_id"`
	Score              int     `json:"score"`
	Force              bool    `json:"force,omitempty"`
	DisableEditMessage bool    `json:"disable_edit_message,omitempty"`
	ChatID             *ChatID `json:"chat_id,omitempty"`
	MessageID          int     `json:"message_id,omitempty"`
	InlineMessageID    string  `json:"inline_message_id,omitempty"`
}

func (p *SetGameScoreArgs) GetRequestArgs() (*RequestArgs, error) {
	return buildJSONRequestArgs(p)
}

// https://core.telegram.org/bots/api#setgamescore
func (api *API) SetGameScore(args *SetGameScoreArgs) (*OptionalMessage, error) {
	method := "setGameScore"
	timeout := 5 * time.Second
	response, err := api.execute(method, args, timeout)
	if err != nil {
		return nil, err
	}
	return buildOptionalMessage(response)
}

type GetGameHighScoresArgs struct {
	UserID          int     `json:"user_id"`
	ChatID          *ChatID `json:"chat_id,omitempty"`
	MessageID       int     `json:"message_id,omitempty"`
	InlineMessageID string  `json:"inline_message_id,omitempty"`
}

func (p *GetGameHighScoresArgs) GetRequestArgs() (*RequestArgs, error) {
	return buildJSONRequestArgs(p)
}

// https://core.telegram.org/bots/api#getgamehighscores
func (api *API) GetGameHighScores(args *GetGameHighScoresArgs) (*[]*GameHighScore, error) {
	var gameHighScore *[]*GameHighScore
	method := "getGameHighScores"
	timeout := 5 * time.Second
	response, err := api.execute(method, args, timeout)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(*response, &gameHighScore); err != nil {
		return nil, err
	}
	return gameHighScore, nil
}
