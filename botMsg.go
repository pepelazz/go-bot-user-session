package userSession

import (
	"fmt"
	"github.com/pepelazz/go-bot-telebot"
	"strings"
	"strconv"
)

func (s *S) SetMsgFromBot(msg telebot.Message) {
	msg.Text = strings.TrimSpace(msg.Text)
	s.BotMsg.CurrentMsg = msg
	s.BotMsg.AnswerMessageType = BotMsgTypeText
}

func (s *S) MsgContain(arr ...string) (isTrue bool) {
	for _, txt := range arr {
		if strings.Contains(strings.ToLower(s.MsgText()), strings.ToLower(txt)) {
			return true
		}
	}
	return false
}

func (s *S) SendChatAction(action string) (err error) {
	return  Bot.SendChatAction(s, action)
}

func (s *S) SendMsg() (result *telebot.MsgResult, err error) {
	switch s.BotMsg.AnswerMessageType {
	case BotMsgTypePhoto:
		err = Bot.SendPhoto(s, &s.BotMsg.AnswerPhoto, s.BotMsg.SendOptions)
		Graylog.L().Infom(map[string]interface{}{"userId": s.IdInt(), "type": "answerPhoto"}, s.BotMsg.AnswerPhoto.Url)
	case BotMsgTypeSticker:
		err = Bot.SendSticker(s, &s.BotMsg.AnswerSticker, s.BotMsg.SendOptions)
		Graylog.L().Infom(map[string]interface{}{"userId": s.IdInt(), "type": "answerSticker"}, s.BotMsg.AnswerSticker.FileID)
	case BotMsgTypeDocument:
		err = Bot.SendDocument(s, &s.BotMsg.AnswerDocument, s.SendOptions())
		Graylog.L().Infom(map[string]interface{}{"userId": s.IdInt(), "type": "answerDocument"}, s.BotMsg.AnswerDocument.FileName)
	case BotMsgTypeText:
		msg := strings.Replace(s.AnswerMsg(), "\\n", "\n", -1)
		result, err = Bot.SendMessage(s, msg, s.SendOptions())
		Graylog.L().Infom(map[string]interface{}{"userId": s.IdInt(), "type": "answerTxt"}, msg)
	default:
		result, err = Bot.SendMessage(s, s.AnswerMsg(), nil)
	}
	if err != nil {
		Graylog.L().Errm(map[string]interface{}{"userId": s.IdInt(), "type": "answerTxt"},fmt.Sprintf("bot send message error %s", err))
	}
	s.ClearMsgText()
	return
}

func (s *S) CurrentMsg() telebot.Message {
	return s.BotMsg.CurrentMsg
}

func (s *S) MsgText() string {
	return s.BotMsg.CurrentMsg.Text
}

func (s *S) SetMsgText(text string) *S {
	s.BotMsg.CurrentMsg.Text = text
	return s
}

func (s *S) ClearMsgText() *S {
	s.BotMsg.CurrentMsg.Text = ""
	s.BotMsg.AnswerMessage.Text = ""
	return s
}

func (s *S) AnswerMsg() string {
	return s.BotMsg.AnswerMessage.Text
}

func (s *S) SetAnswerMsg(msg string) *S {
	s.BotMsg.AnswerMessage = telebot.Message{Text: msg}
	return s
}

func (s *S) SetAnswerMsgWithPhoto(msg string, path string, url string) *S {
	s.BotMsg.AnswerMessageType = BotMsgTypePhoto
	if len(path) > 0 {
		photo, _ := telebot.NewFile(path)
		s.BotMsg.AnswerPhoto = telebot.Photo{File:photo, Caption: msg}
		return s
	}
	if len(url) > 0 {
		s.BotMsg.AnswerPhoto = telebot.Photo{Url:url, Caption: msg}
		return s
	}
	s.BotMsg.AnswerMessageType = BotMsgTypeText
	s.BotMsg.AnswerMessage = telebot.Message{Text: msg}

	return s
}

func (s *S) SetAnswerWithDocument(path string) *S {
	s.BotMsg.AnswerMessageType = BotMsgTypeDocument
	file, _ := telebot.NewFile(path)
	s.BotMsg.AnswerDocument = telebot.Document{File:file}
	return s
}

func (s *S) SetAnswerWithSticker(stickerId string) *S {
	s.BotMsg.AnswerMessageType = BotMsgTypeSticker
	file := telebot.File{FileID: stickerId}
	s.BotMsg.AnswerSticker = telebot.Sticker{File:file}
	return s
}

func (s *S) SendOptions() *telebot.SendOptions {
	return s.BotMsg.SendOptions
}

func (s *S) IdInt() (id int) {
	id, _ =strconv.Atoi(s.Id)
	return
}

func (s *S) Kb() [][]string {
	return s.BotMsg.SendOptions.ReplyMarkup.CustomKeyboard
}

func (s *S) SetKb(kb interface{}) *S {
	switch v := kb.(type) {
	case [][]string:
		s.BotMsg.SendOptions.ReplyMarkup.CustomKeyboard = v
	case [][]telebot.KeyboardButton:
		s.BotMsg.SendOptions.ReplyMarkup.InlineKeyboard = v
	case nil:
		s.BotMsg.SendOptions.ReplyMarkup.CustomKeyboard = nil
		s.BotMsg.SendOptions.ReplyMarkup.HideCustomKeyboard = true
	default:
		fmt.Printf("UserSession.SetAnswerKb error: uknown keyboard type %s\n", kb)
	}
	//s.BotMsg.SendOptions.ReplyMarkup.HideCustomKeyboard = false
	s.BotMsg.SendOptions.ReplyMarkup.ResizeKeyboard = true
	s.SetKbOneTime(false)
	return s
}

func (s *S) MakeKb(btns ...string) *S {
	s.BotMsg.SendOptions.ReplyMarkup.CustomKeyboard = [][]string{btns}
	s.BotMsg.SendOptions.ReplyMarkup.HideCustomKeyboard = false
	s.BotMsg.SendOptions.ReplyMarkup.ResizeKeyboard = true
	return s
}

func (s *S) SetKbOneTime(v bool) *S {
	s.BotMsg.SendOptions.ReplyMarkup.OneTimeKeyboard = v
	return s
}

func (s *S) HideKb() *S {
	s.BotMsg.SendOptions.ReplyMarkup.HideCustomKeyboard = true
	s.BotMsg.SendOptions.ReplyMarkup.CustomKeyboard = nil
	s.SetKbOneTime(false)
	return s
}

func (s *S) ClearAllKb() *S {
	s.BotMsg.SendOptions = &telebot.SendOptions{ParseMode: telebot.ModeHTML}
	return s
}


