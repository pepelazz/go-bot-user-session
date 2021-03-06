package userSession

import (
	"github.com/pepelazz/go-bot-telebot"
	"strconv"
	"github.com/pepelazz/go-bot-common"
)

var (
	Bot *telebot.Bot
	Graylog *goBotCommon.GraylogType
)

func Init(bot *telebot.Bot, graylog *goBotCommon.GraylogType) {
	Bot = bot
	Graylog = graylog
}

type BotMsgType uint16

const (
	BotMsgTypeText = BotMsgType(1)
	BotMsgTypePhoto = BotMsgType(2)
	BotMsgTypeSticker = BotMsgType(3)
	BotMsgTypeGif = BotMsgType(4)
	BotMsgTypeDocument = BotMsgType(5)
)

type S struct {
	Id       string //userTgId
	BotMsg   struct {
			 CurrentMsg        telebot.Message
			 AnswerMessage     telebot.Message
			 AnswerPhoto       telebot.Photo
			 AnswerSticker     telebot.Sticker
			 AnswerDocument    telebot.Document
			 AnswerMessageType BotMsgType
			 SendOptions       *telebot.SendOptions
		 }
	Callback struct {
			 Res         *telebot.CallbackResponse
			 Req         *telebot.Callback
			 IsProcessed bool
		 }
}

// реализация интерфейса для SendMessage
func (s *S) Destination() string {
	return s.Id
}

func New(msg telebot.Message) (newSession *S, err error) {
	newSession = &S{
		Id: strconv.Itoa(msg.Sender.ID),
	}
	newSession.BotMsg.CurrentMsg = msg
	newSession.BotMsg.AnswerMessageType = BotMsgTypeText
	newSession.BotMsg.SendOptions = &telebot.SendOptions{
		ReplyMarkup: telebot.ReplyMarkup{
			ResizeKeyboard: true,
			CustomKeyboard: [][]string{},

		},
		ParseMode: telebot.ModeHTML,
	}
	return
}

func NewFromCb(cb telebot.Callback) (newSession *S, err error) {
	newSession = &S{
		Id: strconv.Itoa(cb.Sender.ID),
	}
	newSession.BotMsg.CurrentMsg = cb.Message
	newSession.BotMsg.AnswerMessageType = BotMsgTypeText
	newSession.BotMsg.SendOptions = &telebot.SendOptions{
		ReplyMarkup: telebot.ReplyMarkup{
			ResizeKeyboard: true,
			CustomKeyboard: [][]string{},

		},
		ParseMode: telebot.ModeHTML,
	}
	newSession.Callback.Req = &cb
	return
}
