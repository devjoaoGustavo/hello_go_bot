package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const getUpdatesPath = "getUpdates"

type MessageEntity struct {
	Type   string `json:"type"`
	OffSet int    `json:"off_set"`
	Length int    `json:"length"`
	Url    string `json:"url"`
	User   User   `json:"user"`
}
type Chat struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Type      string `json:"type"`
}
type From struct {
	Id           int    `json:"id"`
	IsBot        bool   `json:"is_bot"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Username     string `json:"username"`
	LanguageCode string `json:"language_code"`
}
type Audio struct {
	FileId    string `json:"file_id"`
	Duration  int    `json:"duration"`
	Performer string `json:"performer"`
	Title     string `json:"title"`
	MimeType  string `json:"mime_type"`
	FileSise  int    `json:"file_size"`
}
type Document struct {
	FileId   string    `json:"file_id"`
	Thumb    PhotoSize `json:"thumb"`
	FileName string    `json:"file_name"`
	MimeType string    `json:"mime_type"`
	FileSize int       `json:"file_size"`
}
type Animation struct {
	FileId   string    `json:"file_id"`
	Thumb    PhotoSize `json:"thumb"`
	FileName string    `json:"file_name"`
	MimeType string    `json:"mime_type"`
	FileSize int       `json:"file_size"`
}
type Game struct {
	Title        string          `json:"title"`
	Description  string          `json:"description"`
	Photo        []PhotoSize     `json:"photo"`
	Text         string          `json:"text"`
	TextEntities []MessageEntity `json:"text_entities"`
	Animation    Animation       `json:"animation"`
}
type PhotoSize struct {
	FileId   string `json:"file_id"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
	FileSize int    `json:"file_size"`
}
type MaskPosition struct {
	Point  string  `json:"point"`
	XShift float64 `json:"x_shift"`
	YShift float64 `json:"y_shift"`
	Scale  float64 `json:"scale"`
}
type Sticker struct {
	FileId       string       `json:"file_id"`
	Width        int          `json:"width"`
	Height       int          `json:"height"`
	Thumb        PhotoSize    `json:"thumb"`
	Emoji        string       `json:"emoji"`
	SetName      string       `json:"set_name"`
	MaskPosition MaskPosition `json:"mask_position"`
	FileSize     int          `json:"file_size"`
}
type Video struct {
	FileId   string    `json:"file_id"`
	Width    int       `json:"width"`
	Height   int       `json:"height"`
	Duration int       `json:"duration"`
	Thumb    PhotoSize `json:"thumb"`
	MimeType string    `json:"mime_type"`
	FileSize int       `json:"file_size"`
}
type Voice struct {
	FileId   string `json:"file_id"`
	Duration int    `json:"duration"`
	MimeType string `json:"mime_type"`
	FileSize int    `json:"file_size"`
}
type VideoNote struct {
	FileId   string    `json:"file_id"`
	Length   int       `json:"length"`
	Duration int       `json:"duration"`
	Thumb    PhotoSize `json:"thumb"`
	FileSize int       `json:"file_size"`
}
type Contact struct {
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	UserId      int    `json:"user_id"`
}
type Venue struct {
	Location     Location `json:"location"`
	Title        string   `json:"title"`
	Address      string   `json:"address"`
	FoursquareId string   `json:"foursquare_id`
}
type Invoice struct {
	Title          string `json:"title"`
	Description    string `json:"description"`
	StartParameter string `json:"start_parameter"`
	Currency       string `json:"currency"`
	TotalAmount    int    `json:"total_amount"`
}
type SuccessfulPayment struct {
	Currency                string    `json:"currency"`
	TotalAmount             int       `json:"total_amount"`
	InvoicePayload          string    `json:"invoice_payload"`
	ShippingOptionId        string    `json:"shipping_option_id"`
	OrderInfo               OrderInfo `json:"order_info"`
	TelegramPaymentChargeId string    `json:"telegram_payment_charge_id"`
	ProviderPaymentChargeId string    `json:"provider_payment_charge_id"`
}
type Message struct {
	MessageId             int               `json:"message_id"`
	From                  From              `json:"from"`
	Date                  int               `json:"date"`
	Chat                  Chat              `json:"chat"`
	ForwardFrom           User              `json:"forward_from"`
	ForwardFromChat       Chat              `json:"forward_from_chat"`
	ForwardFromMessageId  int               `json:"forward_from_message_id"`
	ForwardSignature      string            `json:"forward_signature"`
	ForwardDate           int               `json:"forward_date"`
	ReplyToMessage        *Message          `json:"reply_to_message"`
	EditMessage           int               `json:"edit_message"`
	MediaGroupId          string            `json:"media_group_id"`
	AuthorSignature       string            `json:"author_signature"`
	Text                  string            `json:"text"`
	Entities              []MessageEntity   `json:"entities"`
	Audio                 Audio             `json:"audio"`
	Document              Document          `json:"document"`
	Game                  Game              `json:"game"`
	Photo                 []PhotoSize       `json:"photo"`
	Sticker               Sticker           `json:"sticker"`
	Video                 Video             `json:"video"`
	Voice                 Voice             `json:"voice"`
	VideoNote             VideoNote         `json:"video_note"`
	Caption               string            `json:"caption"`
	Contact               Contact           `json:"contact"`
	Location              Location          `json:"location"`
	Venue                 Venue             `json:"venue"`
	NewChatMembers        []User            `json:"new_chat_members"`
	LeftChatMember        User              `json:"left_chat_member`
	NewChatTitle          string            `json:"new_chat_title"`
	NewChatPhoto          []PhotoSize       `json:"new_chat_photo"`
	DeleteChatPhoto       bool              `json:"delete_chat_photo"`
	GroupChatCreated      bool              `json:'group_chat_created"`
	SupergroupChatCreated bool              `json:"supergroup_chat_created"`
	ChannelChatCreated    bool              `json:"channel_chat_created"`
	MigrateToChatId       int               `json:"migrate_to_chat_id"`
	MigrateFromChatId     int               `json:"migrate_from_chat_id"`
	PinnedMessage         *Message          `json:"pinned_message"`
	Invoice               Invoice           `json:"invoice"`
	SuccessfulPayment     SuccessfulPayment `json:"successful_payment"`
}
type Location struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}
type InlineQuery struct {
	Id       string   `json:"id"`
	From     User     `json:"from"`
	Location Location `json:"location"`
	Query    string   `json:"query"`
	OffSet   string   `json:"off_set"`
}
type ChosenInlineResult struct {
	ResultId        string   `json:"result_id"`
	From            User     `json:"from"`
	Location        Location `json:"location"`
	InlineMessageId string   `json:"inline_message_id"`
	Query           string   `json:"query"`
}
type CallbackQuery struct {
	Id              string  `json:"id"`
	From            User    `json:"from"`
	Message         Message `json:"message"`
	InlineMessageId string  `json:"inline_message_id"`
	ChatInstance    string  `json:"chat_instance"`
	Data            string  `json:"data"`
	GameShortName   string  `json:"game_short_name"`
}
type LabeledPrice struct {
	Label  string `json:"label"`
	Amount int    `json:"amount"`
}
type ShippingAddress struct {
	CountryCode string `json:"country_code"`
	State       string `json:"state"`
	City        string `json:"city"`
	StreetLine1 string `json:"street_line1`
	StreetLine2 string `json:"street_line2`
	PostCode    string `json:"post_code"`
}
type ShippingQuery struct {
	Id              string          `json:"id"`
	From            User            `json:"from"`
	InvoicePayload  string          `json:"invoice_payload"`
	ShippingAddress ShippingAddress `json:"shipping_address`
}
type OrderInfo struct {
	Name            string          `json:"name"`
	PhoneNumber     string          `json:"phone_number"`
	Email           string          `json:"email"`
	ShippingAddress ShippingAddress `json:"shipping_address`
}
type PreCheckoutQuery struct {
	Id               string    `json:"id"`
	From             User      `json:"from"`
	Currency         string    `json:"currency"`
	TotalAmount      int       `json:"total_amount"`
	InvoicePayload   string    `json:"invoice_payload"`
	ShippingOptionId string    `json:"shipping_option_id"`
	OrderInfo        OrderInfo `json:"order_info"`
}
type Update struct {
	UpdateId           int                `json:"update_id"`
	Message            Message            `json:"message"`
	EditedMessage      Message            `json:"updated_message"`
	ChannelPost        Message            `json:"channel_post"`
	EditedChannelPost  Message            `json:"edited_channel_post"`
	InlineQuery        InlineQuery        `json:"inline_query"`
	ChosenInlineResult ChosenInlineResult `json:"chosen_inline_result"`
	CallbackQuery      CallbackQuery      `json:"callback_query"`
	ShippingQuery      ShippingQuery      `json:"shipping_query"`
	PreCheckoutQuery   PreCheckoutQuery   `json:"pre_checkout_query"`
}
type GetUpdates struct {
	Ok     bool     `json:"ok"`
	Result []Update `json:"result"`
}

func showUpdates(config Config) {
	var updates GetUpdates
	url := fmt.Sprintf("%s/bot%s/%s", config.HelloGoBot.Telegram.Api.EndPoint, config.HelloGoBot.Telegram.Api.Token, getUpdatesPath)
	resp, err0 := http.Get(url)
	if err0 != nil {
		s := err0.Error()
		fmt.Printf("%q\n", s)
		return
	}
	defer resp.Body.Close()
	body, err1 := ioutil.ReadAll(resp.Body)
	if err1 != nil {
		s := err1.Error()
		fmt.Printf("%q\n", s)
		return
	}
	err2 := json.Unmarshal(body, &updates)
	if err2 != nil {
		s := err2.Error()
		fmt.Printf("%q\n", s)
		return
	}
	fmt.Println("Esta é a resposta do método getUpdates")
	fmt.Println("Ok:", updates.Ok)
	fmt.Println("result:", updates.Result)

	for _, update := range updates.Result {
		fmt.Println("UpdateID:", update.UpdateId)
		fmt.Println("Message:", update.Message)
		fmt.Println("Message MessageId:", update.Message.MessageId)

		fmt.Println("Message From:", update.Message.From)
		fmt.Println("Message From Id:", update.Message.From.Id)
		fmt.Println("Message From IsBot:", update.Message.From.IsBot)
		fmt.Println("Message From FirstName:", update.Message.From.FirstName)
		fmt.Println("Message From LastName:", update.Message.From.LastName)
		fmt.Println("Message From Username:", update.Message.From.Username)
		fmt.Println("Message From LanguageCode:", update.Message.From.LanguageCode)

		fmt.Println("Message Document", update.Message.Document)
		fmt.Println("Message Document FileId", update.Message.Document.FileId)
		fmt.Println("Message Document Thumb", update.Message.Document.Thumb)
		fmt.Println("Message Document Thumb FileId", update.Message.Document.Thumb.FileId)
		fmt.Println("Message Document Thumb Width", update.Message.Document.Thumb.Width)
		fmt.Println("Message Document Thumb Height", update.Message.Document.Thumb.Height)
		fmt.Println("Message Document Thumb FileSize", update.Message.Document.Thumb.FileSize)
		fmt.Println("Message Document FileName", update.Message.Document.FileName)
		fmt.Println("Message Document MimeType", update.Message.Document.MimeType)
		fmt.Println("Message Document FileSize", update.Message.Document.FileSize)

		fmt.Println("Message Chat:", update.Message.Chat)
		fmt.Println("Message Chat Id:", update.Message.Chat.Id)
		fmt.Println("Message Chat FirstName:", update.Message.Chat.FirstName)
		fmt.Println("Message Chat LastName:", update.Message.Chat.LastName)
		fmt.Println("Message Chat Username:", update.Message.Chat.Username)
		fmt.Println("Message Chat Type:", update.Message.Chat.Type)

		fmt.Println("Message Date:", update.Message.Date)
		fmt.Println("Message Text:", update.Message.Text)
		fmt.Println("Message Entities:", update.Message.Entities)
		for _, entity := range update.Message.Entities {
			fmt.Println("Message Entities OffSet:", entity.OffSet)
			fmt.Println("Message Entities Length:", entity.Length)
			fmt.Println("Message Entities Type:", entity.Type)
		}
	}
}
