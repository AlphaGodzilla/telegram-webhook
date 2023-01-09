package main

import (
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

type TgBot struct {
	token  string
	chatId int64
	debug  bool
	bot    *tgbotapi.BotAPI
}

func NewTgBot(token string, chatId int64, debug bool) (*TgBot, error) {
	botApi, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}
	botApi.Debug = debug
	log.Printf("机器人账户认证成功 %s", botApi.Self.UserName)
	return &TgBot{
		token:  token,
		debug:  debug,
		chatId: chatId,
		bot:    botApi,
	}, nil
}

func (tb *TgBot) sendMessage(text string) (bool, error) {
	msg := tgbotapi.NewMessage(tb.chatId, text)
	msg1, err := tb.bot.Send(msg)
	if err != nil {
		return false, err
	}
	msg1Json, err := json.Marshal(msg1)
	if err != nil {
		return false, err
	}
	log.Printf("成功发送消息: %s", string(msg1Json))
	return true, nil
}
