package tg

// https://core.telegram.org/bots/api#update
type Update struct {
	UpdateID           int                 `json:"update_id"`
	Message            *Message            `json:"message,omitempty"`
	EditedMessage      *Message            `json:"edited_message,omitempty"`
	ChannelPost        *Message            `json:"channel_post,omitempty"`
	EditedChannelPost  *Message            `json:"edited_channel_post,omitempty"`
	InlineQuery        *InlineQuery        `json:"inline_query,omitempty"`
	ChosenInlineResult *ChosenInlineResult `json:"chosen_inline_result,omitempty"`
	CallbackQuery      *CallbackQuery      `json:"callback_query,omitempty"`
	ShippingQuery      *ShippingQuery      `json:"shipping_query,omitempty"`
	PreCheckoutQuery   *PreCheckoutQuery   `json:"pre_checkout_query,omitempty"`
	Poll               *Poll               `json:"poll,omitempty"`
}

// https://core.telegram.org/bots/api#webhookinfo
type WebhookInfo struct {
	URL                  string   `json:"url"`
	HasCustomCertificate bool     `json:"has_custom_certificate"`
	PendingUpdateCount   int      `json:"pending_update_count"`
	LastErrorDate        int      `json:"last_error_date,omitempty"`
	LastErrorMessage     string   `json:"last_error_message,omitempty"`
	MaxConnections       int      `json:"max_connections,omitempty"`
	AllowedUpdates       []string `json:"allowed_updates,omitempty"`
}

// https://core.telegram.org/bots/api#user
type User struct {
	ID           int    `json:"id"`
	IsBot        bool   `json:"is_bot"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name,omitempty"`
	Username     string `json:"username,omitempty"`
	LanguageCode string `json:"language_code,omitempty"`
}

// https://core.telegram.org/bots/api#chat
type Chat struct {
	ID               int              `json:"id"`
	Type             string           `json:"type"`
	Title            string           `json:"title,omitempty"`
	Username         string           `json:"username,omitempty"`
	FirstName        string           `json:"first_name,omitempty"`
	LastName         string           `json:"last_name,omitempty"`
	Photo            *ChatPhoto       `json:"photo,omitempty"`
	Description      string           `json:"description,omitempty"`
	InviteLink       string           `json:"invite_link,omitempty"`
	PinnedMessage    *Message         `json:"pinned_message,omitempty"`
	Permissions      *ChatPermissions `json:"permissions,omitempty"`
	StickerSetName   string           `json:"sticker_set_name,omitempty"`
	CanSetStickerSet bool             `json:"can_set_sticker_set,omitempty"`
}

// https://core.telegram.org/bots/api#message
type Message struct {
	MessageID             int                   `json:"message_id"`
	From                  *User                 `json:"from,omitempty"`
	Date                  int                   `json:"date"`
	Chat                  *Chat                 `json:"chat"`
	ForwardFrom           *User                 `json:"forward_from,omitempty"`
	ForwardFromChat       *Chat                 `json:"forward_from_chat,omitempty"`
	ForwardFromMessageID  int                   `json:"forward_from_message_id,omitempty"`
	ForwardSignature      string                `json:"forward_signature,omitempty"`
	ForwardSenderName     string                `json:"forward_sender_name,omitempty"`
	ForwardDate           int                   `json:"forward_date,omitempty"`
	ReplyToMessage        *Message              `json:"reply_to_message,omitempty"`
	EditDate              int                   `json:"edit_date,omitempty"`
	MediaGroupID          string                `json:"media_group_id,omitempty"`
	AuthorSignature       string                `json:"author_signature,omitempty"`
	Text                  string                `json:"text,omitempty"`
	Entities              []*MessageEntity      `json:"entities,omitempty"`
	CaptionEntities       []*MessageEntity      `json:"caption_entities,omitempty"`
	Audio                 *Audio                `json:"audio,omitempty"`
	Document              *Document             `json:"document,omitempty"`
	Animation             *Animation            `json:"animation,omitempty"`
	Game                  *Game                 `json:"game,omitempty"`
	Photo                 []*PhotoSize          `json:"photo,omitempty"`
	Sticker               *Sticker              `json:"sticker,omitempty"`
	Video                 *Video                `json:"video,omitempty"`
	Voice                 *Voice                `json:"voice,omitempty"`
	VideoNote             *VideoNote            `json:"video_note,omitempty"`
	Caption               string                `json:"caption,omitempty"`
	Contact               *Contact              `json:"contact,omitempty"`
	Location              *Location             `json:"location,omitempty"`
	Venue                 *Venue                `json:"venue,omitempty"`
	Poll                  *Poll                 `json:"poll,omitempty"`
	NewChatMembers        []*User               `json:"new_chat_members,omitempty"`
	LeftChatMember        *User                 `json:"left_chat_member,omitempty"`
	NewChatTitle          string                `json:"new_chat_title,omitempty"`
	NewChatPhoto          []*PhotoSize          `json:"new_chat_photo,omitempty"`
	DeleteChatPhoto       bool                  `json:"delete_chat_photo,omitempty"`
	GroupChatCreated      bool                  `json:"group_chat_created,omitempty"`
	SupergroupChatCreated bool                  `json:"supergroup_chat_created,omitempty"`
	ChannelChatCreated    bool                  `json:"channel_chat_created,omitempty"`
	MigrateToChatID       int                   `json:"migrate_to_chat_id,omitempty"`
	MigrateFromChatID     int                   `json:"migrate_from_chat_id,omitempty"`
	PinnedMessage         *Message              `json:"pinned_message,omitempty"`
	Invoice               *Invoice              `json:"invoice,omitempty"`
	SuccessfulPayment     *SuccessfulPayment    `json:"successful_payment,omitempty"`
	ConnectedWebsite      string                `json:"connected_website,omitempty"`
	PassportData          *PassportData         `json:"passport_data,omitempty"`
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// https://core.telegram.org/bots/api#messageentity
type MessageEntity struct {
	Type   string `json:"type"`
	Offset int    `json:"offset"`
	Length int    `json:"length"`
	URL    string `json:"url,omitempty"`
	User   *User  `json:"user,omitempty"`
}

// https://core.telegram.org/bots/api#photosize
type PhotoSize struct {
	FileID   string `json:"file_id"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
	FileSize int    `json:"file_size,omitempty"`
}

// https://core.telegram.org/bots/api#audio
type Audio struct {
	FileID    string     `json:"file_id"`
	Duration  int        `json:"duration"`
	Performer string     `json:"performer,omitempty"`
	Title     string     `json:"title,omitempty"`
	MimeType  string     `json:"mime_type,omitempty"`
	FileSize  int        `json:"file_size,omitempty"`
	Thumb     *PhotoSize `json:"thumb,omitempty"`
}

// https://core.telegram.org/bots/api#document
type Document struct {
	FileID   string     `json:"file_id"`
	Thumb    *PhotoSize `json:"thumb,omitempty"`
	FileName string     `json:"file_name,omitempty"`
	MimeType string     `json:"mime_type,omitempty"`
	FileSize int        `json:"file_size,omitempty"`
}

// https://core.telegram.org/bots/api#video
type Video struct {
	FileID   string     `json:"file_id"`
	Width    int        `json:"width"`
	Height   int        `json:"height"`
	Duration int        `json:"duration"`
	Thumb    *PhotoSize `json:"thumb,omitempty"`
	MimeType string     `json:"mime_type,omitempty"`
	FileSize int        `json:"file_size,omitempty"`
}

// https://core.telegram.org/bots/api#animation
type Animation struct {
	FileID   string     `json:"file_id"`
	Width    int        `json:"width"`
	Height   int        `json:"height"`
	Duration int        `json:"duration"`
	Thumb    *PhotoSize `json:"thumb,omitempty"`
	FileName string     `json:"file_name,omitempty"`
	MimeType string     `json:"mime_type,omitempty"`
	FileSize int        `json:"file_size,omitempty"`
}

// https://core.telegram.org/bots/api#voice
type Voice struct {
	FileID   string `json:"file_id"`
	Duration int    `json:"duration"`
	MimeType string `json:"mime_type,omitempty"`
	FileSize int    `json:"file_size,omitempty"`
}

// https://core.telegram.org/bots/api#videonote
type VideoNote struct {
	FileID   string     `json:"file_id"`
	Length   int        `json:"length"`
	Duration int        `json:"duration"`
	Thumb    *PhotoSize `json:"thumb,omitempty"`
	FileSize int        `json:"file_size,omitempty"`
}

// https://core.telegram.org/bots/api#contact
type Contact struct {
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name,omitempty"`
	UserID      int    `json:"user_id,omitempty"`
	Vcard       string `json:"vcard,omitempty"`
}

// https://core.telegram.org/bots/api#location
type Location struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

// https://core.telegram.org/bots/api#venue
type Venue struct {
	Location       *Location `json:"location"`
	Title          string    `json:"title"`
	Address        string    `json:"address"`
	FoursquareID   string    `json:"foursquare_id,omitempty"`
	FoursquareType string    `json:"foursquare_type,omitempty"`
}

// https://core.telegram.org/bots/api#polloption
type PollOption struct {
	Text       string `json:"text"`
	VoterCount int    `json:"voter_count"`
}

// https://core.telegram.org/bots/api#poll
type Poll struct {
	ID       string        `json:"id"`
	Question string        `json:"question"`
	Options  []*PollOption `json:"options"`
	IsClosed bool          `json:"is_closed"`
}

// https://core.telegram.org/bots/api#userprofilephotos
type UserProfilePhotos struct {
	TotalCount int            `json:"total_count"`
	Photos     [][]*PhotoSize `json:"photos"`
}

// https://core.telegram.org/bots/api#file
type File struct {
	FileID   string `json:"file_id"`
	FileSize int    `json:"file_size,omitempty"`
	FilePath string `json:"file_path,omitempty"`
}

// https://core.telegram.org/bots/api#replykeyboardmarkup
type ReplyKeyboardMarkup struct {
	Keyboard        [][]*KeyboardButton `json:"keyboard"`
	ResizeKeyboard  bool                `json:"resize_keyboard,omitempty"`
	OneTimeKeyboard bool                `json:"one_time_keyboard,omitempty"`
	Selective       bool                `json:"selective,omitempty"`
}

// https://core.telegram.org/bots/api#keyboardbutton
type KeyboardButton struct {
	Text            string `json:"text"`
	RequestContact  bool   `json:"request_contact,omitempty"`
	RequestLocation bool   `json:"request_location,omitempty"`
}

// https://core.telegram.org/bots/api#replykeyboardremove
type ReplyKeyboardRemove struct {
	RemoveKeyboard bool `json:"remove_keyboard"`
	Selective      bool `json:"selective,omitempty"`
}

// https://core.telegram.org/bots/api#inlinekeyboardmarkup
type InlineKeyboardMarkup struct {
	InlineKeyboard [][]*InlineKeyboardButton `json:"inline_keyboard"`
}

// https://core.telegram.org/bots/api#inlinekeyboardbutton
type InlineKeyboardButton struct {
	Text                         string        `json:"text"`
	URL                          string        `json:"url,omitempty"`
	LoginURL                     *LoginURL     `json:"login_url,omitempty"`
	CallbackData                 string        `json:"callback_data,omitempty"`
	SwitchInlineQuery            string        `json:"switch_inline_query,omitempty"`
	SwitchInlineQueryCurrentChat string        `json:"switch_inline_query_current_chat,omitempty"`
	CallbackGame                 *CallbackGame `json:"callback_game,omitempty"`
	Pay                          bool          `json:"pay,omitempty"`
}

// https://core.telegram.org/bots/api#loginurl
type LoginURL struct {
	URL                string `json:"url"`
	ForwardText        string `json:"forward_text,omitempty"`
	BotUsername        string `json:"bot_username,omitempty"`
	RequestWriteAccess bool   `json:"request_write_access,omitempty"`
}

// https://core.telegram.org/bots/api#callbackquery
type CallbackQuery struct {
	ID              string   `json:"id"`
	From            *User    `json:"from"`
	Message         *Message `json:"message,omitempty"`
	InlineMessageID string   `json:"inline_message_id,omitempty"`
	ChatInstance    string   `json:"chat_instance"`
	Data            string   `json:"data,omitempty"`
	GameShortName   string   `json:"game_short_name,omitempty"`
}

// https://core.telegram.org/bots/api#forcereply
type ForceReply struct {
	ForceReply bool `json:"force_reply"`
	Selective  bool `json:"selective,omitempty"`
}

// https://core.telegram.org/bots/api#chatphoto
type ChatPhoto struct {
	SmallFileID string `json:"small_file_id"`
	BigFileID   string `json:"big_file_id"`
}

// https://core.telegram.org/bots/api#chatmember
type ChatMember struct {
	User                  *User  `json:"user"`
	Status                string `json:"status"`
	UntilDate             int    `json:"until_date,omitempty"`
	CanBeEdited           bool   `json:"can_be_edited,omitempty"`
	CanPostMessages       bool   `json:"can_post_messages,omitempty"`
	CanEditMessages       bool   `json:"can_edit_messages,omitempty"`
	CanDeleteMessages     bool   `json:"can_delete_messages,omitempty"`
	CanRestrictMembers    bool   `json:"can_restrict_members,omitempty"`
	CanPromoteMembers     bool   `json:"can_promote_members,omitempty"`
	CanChangeInfo         bool   `json:"can_change_info,omitempty"`
	CanInviteUsers        bool   `json:"can_invite_users,omitempty"`
	CanPinMessages        bool   `json:"can_pin_messages,omitempty"`
	IsMember              bool   `json:"is_member,omitempty"`
	CanSendMessages       bool   `json:"can_send_messages,omitempty"`
	CanSendMediaMessages  bool   `json:"can_send_media_messages,omitempty"`
	CanSendPolls          bool   `json:"can_send_polls,omitempty"`
	CanSendOtherMessages  bool   `json:"can_send_other_messages,omitempty"`
	CanAddWebPagePreviews bool   `json:"can_add_web_page_previews,omitempty"`
}

// https://core.telegram.org/bots/api#chatpermissions
type ChatPermissions struct {
	CanSendMessages       bool `json:"can_send_messages,omitempty"`
	CanSendMediaMessages  bool `json:"can_send_media_messages,omitempty"`
	CanSendPolls          bool `json:"can_send_polls,omitempty"`
	CanSendOtherMessages  bool `json:"can_send_other_messages,omitempty"`
	CanAddWebPagePreviews bool `json:"can_add_web_page_previews,omitempty"`
	CanChangeInfo         bool `json:"can_change_info,omitempty"`
	CanInviteUsers        bool `json:"can_invite_users,omitempty"`
	CanPinMessages        bool `json:"can_pin_messages,omitempty"`
}

// https://core.telegram.org/bots/api#responseparameters
type ResponseParameters struct {
	MigrateToChatID int `json:"migrate_to_chat_id,omitempty"`
	RetryAfter      int `json:"retry_after,omitempty"`
}

const (
	InputMediaTypePhoto     = "photo"
	InputMediaTypeVideo     = "video"
	InputMediaTypeDocument  = "document"
	InputMediaTypeAudio     = "audio"
	InputMediaTypeAnimation = "animation"
)

type InputMedia interface {
	getMedia() []*InputFile
}

// https://core.telegram.org/bots/api#inputmediaphoto
type InputMediaPhoto struct {
	Type        string     `json:"type"`
	Media       string     `json:"media"`
	Caption     string     `json:"caption,omitempty"`
	ParseMode   string     `json:"parse_mode,omitempty"`
	MediaAsFile *InputFile `json:"-"`
}

func (m *InputMediaPhoto) getMedia() []*InputFile {
	return []*InputFile{m.MediaAsFile}
}

// https://core.telegram.org/bots/api#inputmediavideo
type InputMediaVideo struct {
	Type              string     `json:"type"`
	Media             string     `json:"media"`
	Thumb             string     `json:"thumb,omitempty"`
	Caption           string     `json:"caption,omitempty"`
	ParseMode         string     `json:"parse_mode,omitempty"`
	Width             int        `json:"width,omitempty"`
	Height            int        `json:"height,omitempty"`
	Duration          int        `json:"duration,omitempty"`
	SupportsStreaming bool       `json:"supports_streaming,omitempty"`
	MediaAsFile       *InputFile `json:"-"`
	ThumbAsFile       *InputFile `json:"-"`
}

func (m *InputMediaVideo) getMedia() []*InputFile {
	return []*InputFile{m.MediaAsFile, m.ThumbAsFile}
}

// https://core.telegram.org/bots/api#inputmediaanimation
type InputMediaAnimation struct {
	Type        string     `json:"type"`
	Media       string     `json:"media"`
	Thumb       string     `json:"thumb,omitempty"`
	Caption     string     `json:"caption,omitempty"`
	ParseMode   string     `json:"parse_mode,omitempty"`
	Width       int        `json:"width,omitempty"`
	Height      int        `json:"height,omitempty"`
	Duration    int        `json:"duration,omitempty"`
	MediaAsFile *InputFile `json:"-"`
	ThumbAsFile *InputFile `json:"-"`
}

func (m *InputMediaAnimation) getMedia() []*InputFile {
	return []*InputFile{m.MediaAsFile, m.ThumbAsFile}
}

// https://core.telegram.org/bots/api#inputmediaaudio
type InputMediaAudio struct {
	Type        string     `json:"type"`
	Media       string     `json:"media"`
	Thumb       string     `json:"thumb,omitempty"`
	Caption     string     `json:"caption,omitempty"`
	ParseMode   string     `json:"parse_mode,omitempty"`
	Duration    int        `json:"duration,omitempty"`
	Performer   string     `json:"performer,omitempty"`
	Title       string     `json:"title,omitempty"`
	MediaAsFile *InputFile `json:"-"`
	ThumbAsFile *InputFile `json:"-"`
}

func (m *InputMediaAudio) getMedia() []*InputFile {
	return []*InputFile{m.MediaAsFile, m.ThumbAsFile}
}

// https://core.telegram.org/bots/api#inputmediadocument
type InputMediaDocument struct {
	Type        string     `json:"type"`
	Media       string     `json:"media"`
	Thumb       string     `json:"thumb,omitempty"`
	Caption     string     `json:"caption,omitempty"`
	ParseMode   string     `json:"parse_mode,omitempty"`
	MediaAsFile *InputFile `json:"-"`
	ThumbAsFile *InputFile `json:"-"`
}

func (m *InputMediaDocument) getMedia() []*InputFile {
	return []*InputFile{m.MediaAsFile, m.ThumbAsFile}
}

// https://core.telegram.org/bots/api#sticker
type Sticker struct {
	FileID       string        `json:"file_id"`
	Width        int           `json:"width"`
	Height       int           `json:"height"`
	IsAnimated   bool          `json:"is_animated"`
	Thumb        *PhotoSize    `json:"thumb,omitempty"`
	Emoji        string        `json:"emoji,omitempty"`
	SetName      string        `json:"set_name,omitempty"`
	MaskPosition *MaskPosition `json:"mask_position,omitempty"`
	FileSize     int           `json:"file_size,omitempty"`
}

// https://core.telegram.org/bots/api#stickerset
type StickerSet struct {
	Name          string     `json:"name"`
	Title         string     `json:"title"`
	IsAnimated    bool       `json:"is_animated"`
	ContainsMasks bool       `json:"contains_masks"`
	Stickers      []*Sticker `json:"stickers"`
}

// https://core.telegram.org/bots/api#maskposition
type MaskPosition struct {
	Point  string  `json:"point"`
	XShift float64 `json:"x_shift"`
	YShift float64 `json:"y_shift"`
	Scale  float64 `json:"scale"`
}

// https://core.telegram.org/bots/api#inlinequery
type InlineQuery struct {
	ID       string    `json:"id"`
	From     *User     `json:"from"`
	Location *Location `json:"location,omitempty"`
	Query    string    `json:"query"`
	Offset   string    `json:"offset"`
}

const (
	InlineQueryResultTypePhoto    = "photo"
	InlineQueryResultTypeGif      = "gif"
	InlineQueryResultTypeMPEG4Gif = "mpeg4_gif"
	InlineQueryResultTypeVideo    = "video"
	InlineQueryResultTypeAudio    = "audio"
	InlineQueryResultTypeVoice    = "voice"
	InlineQueryResultTypeDocument = "document"
	InlineQueryResultTypeLocation = "location"
	InlineQueryResultTypeVenue    = "venue"
	InlineQueryResultTypeContact  = "contact"
	InlineQueryResultTypeGame     = "game"
)

// https://core.telegram.org/bots/api#inlinequeryresultarticle
type InlineQueryResultArticle struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	Title               string                `json:"title"`
	InputMessageContent interface{}           `json:"input_message_content"` // InputTextMessageContent or InputLocationMessageContent or InputVenueMessageContent or InputContactMessageContent
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	URL                 string                `json:"url,omitempty"`
	HideURL             bool                  `json:"hide_url,omitempty"`
	Description         string                `json:"description,omitempty"`
	ThumbURL            string                `json:"thumb_url,omitempty"`
	ThumbWidth          int                   `json:"thumb_width,omitempty"`
	ThumbHeight         int                   `json:"thumb_height,omitempty"`
}

// https://core.telegram.org/bots/api#inlinequeryresultphoto
type InlineQueryResultPhoto struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	PhotoURL            string                `json:"photo_url"`
	ThumbURL            string                `json:"thumb_url"`
	PhotoWidth          int                   `json:"photo_width,omitempty"`
	PhotoHeight         int                   `json:"photo_height,omitempty"`
	Title               string                `json:"title,omitempty"`
	Description         string                `json:"description,omitempty"`
	Caption             string                `json:"caption,omitempty"`
	ParseMode           string                `json:"parse_mode,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent interface{}           `json:"input_message_content,omitempty"` // InputTextMessageContent or InputLocationMessageContent or InputVenueMessageContent or InputContactMessageContent
}

// https://core.telegram.org/bots/api#inlinequeryresultgif
type InlineQueryResultGif struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	GifURL              string                `json:"gif_url"`
	GifWidth            int                   `json:"gif_width,omitempty"`
	GifHeight           int                   `json:"gif_height,omitempty"`
	GifDuration         int                   `json:"gif_duration,omitempty"`
	ThumbURL            string                `json:"thumb_url"`
	Title               string                `json:"title,omitempty"`
	Caption             string                `json:"caption,omitempty"`
	ParseMode           string                `json:"parse_mode,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent interface{}           `json:"input_message_content,omitempty"` // InputTextMessageContent or InputLocationMessageContent or InputVenueMessageContent or InputContactMessageContent
}

// https://core.telegram.org/bots/api#inlinequeryresultmpeg4gif
type InlineQueryResultMPEG4Gif struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	MPEG4URL            string                `json:"mpeg4_url"`
	MPEG4Width          int                   `json:"mpeg4_width,omitempty"`
	MPEG4Height         int                   `json:"mpeg4_height,omitempty"`
	MPEG4Duration       int                   `json:"mpeg4_duration,omitempty"`
	ThumbURL            string                `json:"thumb_url"`
	Title               string                `json:"title,omitempty"`
	Caption             string                `json:"caption,omitempty"`
	ParseMode           string                `json:"parse_mode,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent interface{}           `json:"input_message_content,omitempty"` // InputTextMessageContent or InputLocationMessageContent or InputVenueMessageContent or InputContactMessageContent
}

// https://core.telegram.org/bots/api#inlinequeryresultvideo
type InlineQueryResultVideo struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	VideoURL            string                `json:"video_url"`
	MimeType            string                `json:"mime_type"`
	ThumbURL            string                `json:"thumb_url"`
	Title               string                `json:"title"`
	Caption             string                `json:"caption,omitempty"`
	ParseMode           string                `json:"parse_mode,omitempty"`
	VideoWidth          int                   `json:"video_width,omitempty"`
	VideoHeight         int                   `json:"video_height,omitempty"`
	VideoDuration       int                   `json:"video_duration,omitempty"`
	Description         string                `json:"description,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent interface{}           `json:"input_message_content,omitempty"` // InputTextMessageContent or InputLocationMessageContent or InputVenueMessageContent or InputContactMessageContent
}

// https://core.telegram.org/bots/api#inlinequeryresultaudio
type InlineQueryResultAudio struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	AudioURL            string                `json:"audio_url"`
	Title               string                `json:"title"`
	Caption             string                `json:"caption,omitempty"`
	ParseMode           string                `json:"parse_mode,omitempty"`
	Performer           string                `json:"performer,omitempty"`
	AudioDuration       int                   `json:"audio_duration,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent interface{}           `json:"input_message_content,omitempty"` // InputTextMessageContent or InputLocationMessageContent or InputVenueMessageContent or InputContactMessageContent
}

// https://core.telegram.org/bots/api#inlinequeryresultvoice
type InlineQueryResultVoice struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	VoiceURL            string                `json:"voice_url"`
	Title               string                `json:"title"`
	Caption             string                `json:"caption,omitempty"`
	ParseMode           string                `json:"parse_mode,omitempty"`
	VoiceDuration       int                   `json:"voice_duration,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent interface{}           `json:"input_message_content,omitempty"` // InputTextMessageContent or InputLocationMessageContent or InputVenueMessageContent or InputContactMessageContent
}

// https://core.telegram.org/bots/api#inlinequeryresultdocument
type InlineQueryResultDocument struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	Title               string                `json:"title"`
	Caption             string                `json:"caption,omitempty"`
	ParseMode           string                `json:"parse_mode,omitempty"`
	DocumentURL         string                `json:"document_url"`
	MimeType            string                `json:"mime_type"`
	Description         string                `json:"description,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent interface{}           `json:"input_message_content,omitempty"` // InputTextMessageContent or InputLocationMessageContent or InputVenueMessageContent or InputContactMessageContent
	ThumbURL            string                `json:"thumb_url,omitempty"`
	ThumbWidth          int                   `json:"thumb_width,omitempty"`
	ThumbHeight         int                   `json:"thumb_height,omitempty"`
}

// https://core.telegram.org/bots/api#inlinequeryresultlocation
type InlineQueryResultLocation struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	Latitude            float64               `json:"latitude"`
	Longitude           float64               `json:"longitude"`
	Title               string                `json:"title"`
	LivePeriod          int                   `json:"live_period,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent interface{}           `json:"input_message_content,omitempty"` // InputTextMessageContent or InputLocationMessageContent or InputVenueMessageContent or InputContactMessageContent
	ThumbURL            string                `json:"thumb_url,omitempty"`
	ThumbWidth          int                   `json:"thumb_width,omitempty"`
	ThumbHeight         int                   `json:"thumb_height,omitempty"`
}

// https://core.telegram.org/bots/api#inlinequeryresultvenue
type InlineQueryResultVenue struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	Latitude            float64               `json:"latitude"`
	Longitude           float64               `json:"longitude"`
	Title               string                `json:"title"`
	Address             string                `json:"address"`
	FoursquareID        string                `json:"foursquare_id,omitempty"`
	FoursquareType      string                `json:"foursquare_type,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent interface{}           `json:"input_message_content,omitempty"` // InputTextMessageContent or InputLocationMessageContent or InputVenueMessageContent or InputContactMessageContent
	ThumbURL            string                `json:"thumb_url,omitempty"`
	ThumbWidth          int                   `json:"thumb_width,omitempty"`
	ThumbHeight         int                   `json:"thumb_height,omitempty"`
}

// https://core.telegram.org/bots/api#inlinequeryresultcontact
type InlineQueryResultContact struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	PhoneNumber         string                `json:"phone_number"`
	FirstName           string                `json:"first_name"`
	LastName            string                `json:"last_name,omitempty"`
	Vcard               string                `json:"vcard,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent interface{}           `json:"input_message_content,omitempty"` // InputTextMessageContent or InputLocationMessageContent or InputVenueMessageContent or InputContactMessageContent
	ThumbURL            string                `json:"thumb_url,omitempty"`
	ThumbWidth          int                   `json:"thumb_width,omitempty"`
	ThumbHeight         int                   `json:"thumb_height,omitempty"`
}

// https://core.telegram.org/bots/api#inlinequeryresultgame
type InlineQueryResultGame struct {
	Type          string                `json:"type"`
	ID            string                `json:"id"`
	GameShortName string                `json:"game_short_name"`
	ReplyMarkup   *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// https://core.telegram.org/bots/api#inlinequeryresultcachedphoto
type InlineQueryResultCachedPhoto struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	PhotoFileID         string                `json:"photo_file_id"`
	Title               string                `json:"title,omitempty"`
	Description         string                `json:"description,omitempty"`
	Caption             string                `json:"caption,omitempty"`
	ParseMode           string                `json:"parse_mode,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent interface{}           `json:"input_message_content,omitempty"` // InputTextMessageContent or InputLocationMessageContent or InputVenueMessageContent or InputContactMessageContent
}

// https://core.telegram.org/bots/api#inlinequeryresultcachedgif
type InlineQueryResultCachedGif struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	GifFileID           string                `json:"gif_file_id"`
	Title               string                `json:"title,omitempty"`
	Caption             string                `json:"caption,omitempty"`
	ParseMode           string                `json:"parse_mode,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent interface{}           `json:"input_message_content,omitempty"` // InputTextMessageContent or InputLocationMessageContent or InputVenueMessageContent or InputContactMessageContent
}

// https://core.telegram.org/bots/api#inlinequeryresultcachedmpeg4gif
type InlineQueryResultCachedMPEG4Gif struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	MPEG4FileID         string                `json:"mpeg4_file_id"`
	Title               string                `json:"title,omitempty"`
	Caption             string                `json:"caption,omitempty"`
	ParseMode           string                `json:"parse_mode,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent interface{}           `json:"input_message_content,omitempty"` // InputTextMessageContent or InputLocationMessageContent or InputVenueMessageContent or InputContactMessageContent
}

// https://core.telegram.org/bots/api#inlinequeryresultcachedsticker
type InlineQueryResultCachedSticker struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	StickerFileID       string                `json:"sticker_file_id"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent interface{}           `json:"input_message_content,omitempty"` // InputTextMessageContent or InputLocationMessageContent or InputVenueMessageContent or InputContactMessageContent
}

// https://core.telegram.org/bots/api#inlinequeryresultcacheddocument
type InlineQueryResultCachedDocument struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	Title               string                `json:"title"`
	DocumentFileID      string                `json:"document_file_id"`
	Description         string                `json:"description,omitempty"`
	Caption             string                `json:"caption,omitempty"`
	ParseMode           string                `json:"parse_mode,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent interface{}           `json:"input_message_content,omitempty"` // InputTextMessageContent or InputLocationMessageContent or InputVenueMessageContent or InputContactMessageContent
}

// https://core.telegram.org/bots/api#inlinequeryresultcachedvideo
type InlineQueryResultCachedVideo struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	VideoFileID         string                `json:"video_file_id"`
	Title               string                `json:"title"`
	Description         string                `json:"description,omitempty"`
	Caption             string                `json:"caption,omitempty"`
	ParseMode           string                `json:"parse_mode,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent interface{}           `json:"input_message_content,omitempty"` // InputTextMessageContent or InputLocationMessageContent or InputVenueMessageContent or InputContactMessageContent
}

// https://core.telegram.org/bots/api#inlinequeryresultcachedvoice
type InlineQueryResultCachedVoice struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	VoiceFileID         string                `json:"voice_file_id"`
	Title               string                `json:"title"`
	Caption             string                `json:"caption,omitempty"`
	ParseMode           string                `json:"parse_mode,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent interface{}           `json:"input_message_content,omitempty"` // InputTextMessageContent or InputLocationMessageContent or InputVenueMessageContent or InputContactMessageContent
}

// https://core.telegram.org/bots/api#inlinequeryresultcachedaudio
type InlineQueryResultCachedAudio struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	AudioFileID         string                `json:"audio_file_id"`
	Caption             string                `json:"caption,omitempty"`
	ParseMode           string                `json:"parse_mode,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent interface{}           `json:"input_message_content,omitempty"` // InputTextMessageContent or InputLocationMessageContent or InputVenueMessageContent or InputContactMessageContent
}

// https://core.telegram.org/bots/api#inputmessagecontent
type InputMessageContent struct {
}

// https://core.telegram.org/bots/api#inputtextmessagecontent
type InputTextMessageContent struct {
	MessageText           string `json:"message_text"`
	ParseMode             string `json:"parse_mode,omitempty"`
	DisableWebPagePreview bool   `json:"disable_web_page_preview,omitempty"`
}

// https://core.telegram.org/bots/api#inputlocationmessagecontent
type InputLocationMessageContent struct {
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
	LivePeriod int     `json:"live_period,omitempty"`
}

// https://core.telegram.org/bots/api#inputvenuemessagecontent
type InputVenueMessageContent struct {
	Latitude       float64 `json:"latitude"`
	Longitude      float64 `json:"longitude"`
	Title          string  `json:"title"`
	Address        string  `json:"address"`
	FoursquareID   string  `json:"foursquare_id,omitempty"`
	FoursquareType string  `json:"foursquare_type,omitempty"`
}

// https://core.telegram.org/bots/api#inputcontactmessagecontent
type InputContactMessageContent struct {
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name,omitempty"`
	Vcard       string `json:"vcard,omitempty"`
}

// https://core.telegram.org/bots/api#choseninlineresult
type ChosenInlineResult struct {
	ResultID        string    `json:"result_id"`
	From            *User     `json:"from"`
	Location        *Location `json:"location,omitempty"`
	InlineMessageID string    `json:"inline_message_id,omitempty"`
	Query           string    `json:"query"`
}

// https://core.telegram.org/bots/api#labeledprice
type LabeledPrice struct {
	Label  string `json:"label"`
	Amount int    `json:"amount"`
}

// https://core.telegram.org/bots/api#invoice
type Invoice struct {
	Title          string `json:"title"`
	Description    string `json:"description"`
	StartParameter string `json:"start_parameter"`
	Currency       string `json:"currency"`
	TotalAmount    int    `json:"total_amount"`
}

// https://core.telegram.org/bots/api#shippingaddress
type ShippingAddress struct {
	CountryCode string `json:"country_code"`
	State       string `json:"state"`
	City        string `json:"city"`
	StreetLine1 string `json:"street_line1"`
	StreetLine2 string `json:"street_line2"`
	PostCode    string `json:"post_code"`
}

// https://core.telegram.org/bots/api#orderinfo
type OrderInfo struct {
	Name            string           `json:"name,omitempty"`
	PhoneNumber     string           `json:"phone_number,omitempty"`
	Email           string           `json:"email,omitempty"`
	ShippingAddress *ShippingAddress `json:"shipping_address,omitempty"`
}

// https://core.telegram.org/bots/api#shippingoption
type ShippingOption struct {
	ID     string          `json:"id"`
	Title  string          `json:"title"`
	Prices []*LabeledPrice `json:"prices"`
}

// https://core.telegram.org/bots/api#successfulpayment
type SuccessfulPayment struct {
	Currency                string     `json:"currency"`
	TotalAmount             int        `json:"total_amount"`
	InvoicePayload          string     `json:"invoice_payload"`
	ShippingOptionID        string     `json:"shipping_option_id,omitempty"`
	OrderInfo               *OrderInfo `json:"order_info,omitempty"`
	TelegramPaymentChargeID string     `json:"telegram_payment_charge_id"`
	ProviderPaymentChargeID string     `json:"provider_payment_charge_id"`
}

// https://core.telegram.org/bots/api#shippingquery
type ShippingQuery struct {
	ID              string           `json:"id"`
	From            *User            `json:"from"`
	InvoicePayload  string           `json:"invoice_payload"`
	ShippingAddress *ShippingAddress `json:"shipping_address"`
}

// https://core.telegram.org/bots/api#precheckoutquery
type PreCheckoutQuery struct {
	ID               string     `json:"id"`
	From             *User      `json:"from"`
	Currency         string     `json:"currency"`
	TotalAmount      int        `json:"total_amount"`
	InvoicePayload   string     `json:"invoice_payload"`
	ShippingOptionID string     `json:"shipping_option_id,omitempty"`
	OrderInfo        *OrderInfo `json:"order_info,omitempty"`
}

// https://core.telegram.org/bots/api#passportdata
type PassportData struct {
	Data        []*EncryptedPassportElement `json:"data"`
	Credentials *EncryptedCredentials       `json:"credentials"`
}

// https://core.telegram.org/bots/api#passportfile
type PassportFile struct {
	FileID   string `json:"file_id"`
	FileSize int    `json:"file_size"`
	FileDate int    `json:"file_date"`
}

// https://core.telegram.org/bots/api#encryptedpassportelement
type EncryptedPassportElement struct {
	Type        string          `json:"type"`
	Data        string          `json:"data,omitempty"`
	PhoneNumber string          `json:"phone_number,omitempty"`
	Email       string          `json:"email,omitempty"`
	Files       []*PassportFile `json:"files,omitempty"`
	FrontSide   *PassportFile   `json:"front_side,omitempty"`
	ReverseSide *PassportFile   `json:"reverse_side,omitempty"`
	Selfie      *PassportFile   `json:"selfie,omitempty"`
	Translation []*PassportFile `json:"translation,omitempty"`
	Hash        string          `json:"hash"`
}

// https://core.telegram.org/bots/api#encryptedcredentials
type EncryptedCredentials struct {
	Data   string `json:"data"`
	Hash   string `json:"hash"`
	Secret string `json:"secret"`
}

const (
	PassportElementErrorTypePersonalDetails  = "personal_details"
	PassportElementErrorTypePassport         = "passport"
	PassportElementErrorTypeDriverLicense    = "driver_license"
	PassportElementErrorTypeIdentityCard     = "identity_card"
	PassportElementErrorTypeInternalPassport = "internal_passport"
	PassportElementErrorTypeAddress          = "address"
)

const (
	PassportElementErrorSourceData              = "data"
	PassportElementErrorSourceFront_side        = "front_side"
	PassportElementErrorSourceReverse_side      = "reverse_side"
	PassportElementErrorSourceSelfie            = "selfie"
	PassportElementErrorSourceFile              = "file"
	PassportElementErrorSourceFiles             = "files"
	PassportElementErrorSourceTranslation_file  = "translation_file"
	PassportElementErrorSourceTranslation_files = "translation_files"
	PassportElementErrorSourceUnspecified       = "unspecified"
)

// https://core.telegram.org/bots/api#passportelementerrordatafield
type PassportElementErrorDataField struct {
	Source    string `json:"source"`
	Type      string `json:"type"`
	FieldName string `json:"field_name"`
	DataHash  string `json:"data_hash"`
	Message   string `json:"message"`
}

// https://core.telegram.org/bots/api#passportelementerrorfrontside
type PassportElementErrorFrontSide struct {
	Source   string `json:"source"`
	Type     string `json:"type"`
	FileHash string `json:"file_hash"`
	Message  string `json:"message"`
}

// https://core.telegram.org/bots/api#passportelementerrorreverseside
type PassportElementErrorReverseSide struct {
	Source   string `json:"source"`
	Type     string `json:"type"`
	FileHash string `json:"file_hash"`
	Message  string `json:"message"`
}

// https://core.telegram.org/bots/api#passportelementerrorselfie
type PassportElementErrorSelfie struct {
	Source   string `json:"source"`
	Type     string `json:"type"`
	FileHash string `json:"file_hash"`
	Message  string `json:"message"`
}

// https://core.telegram.org/bots/api#passportelementerrorfile
type PassportElementErrorFile struct {
	Source   string `json:"source"`
	Type     string `json:"type"`
	FileHash string `json:"file_hash"`
	Message  string `json:"message"`
}

// https://core.telegram.org/bots/api#passportelementerrorfiles
type PassportElementErrorFiles struct {
	Source     string   `json:"source"`
	Type       string   `json:"type"`
	FileHashes []string `json:"file_hashes"`
	Message    string   `json:"message"`
}

// https://core.telegram.org/bots/api#passportelementerrortranslationfile
type PassportElementErrorTranslationFile struct {
	Source   string `json:"source"`
	Type     string `json:"type"`
	FileHash string `json:"file_hash"`
	Message  string `json:"message"`
}

// https://core.telegram.org/bots/api#passportelementerrortranslationfiles
type PassportElementErrorTranslationFiles struct {
	Source     string   `json:"source"`
	Type       string   `json:"type"`
	FileHashes []string `json:"file_hashes"`
	Message    string   `json:"message"`
}

// https://core.telegram.org/bots/api#passportelementerrorunspecified
type PassportElementErrorUnspecified struct {
	Source      string `json:"source"`
	Type        string `json:"type"`
	ElementHash string `json:"element_hash"`
	Message     string `json:"message"`
}

// https://core.telegram.org/bots/api#game
type Game struct {
	Title        string           `json:"title"`
	Description  string           `json:"description"`
	Photo        []*PhotoSize     `json:"photo"`
	Text         string           `json:"text,omitempty"`
	TextEntities []*MessageEntity `json:"text_entities,omitempty"`
	Animation    *Animation       `json:"animation,omitempty"`
}

// https://core.telegram.org/bots/api#callbackgame
type CallbackGame struct {
}

// https://core.telegram.org/bots/api#gamehighscore
type GameHighScore struct {
	Position int   `json:"position"`
	User     *User `json:"user"`
	Score    int   `json:"score"`
}

type OptionalMessage struct {
	Successful bool
	Message    *Message
}

type InputFile struct {
	FileName string
	Name     string
}

func (f *InputFile) isAllSet() bool {
	if f != nil && f.Name != "" && f.FileName != "" {
		return true
	}
	return false
}
