package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

const (
	WEB_SERVER_PORT = "WEB_SERVER_PORT"
	TG_BOT_TOKEN    = "TG_BOT_TOKEN"
	TG_BOT_CHAT_ID  = "TG_BOT_CHAT_ID"
	TG_BOT_DEBUG    = "TG_BOT_DEBUG"
	WEB_HOOK_SECRET = "WEB_HOOK_SECRET"
)

func getTokenConfig() string {
	token := os.Getenv(TG_BOT_TOKEN)
	if token == "" {
		log.Fatalf("config [%v] is empty\n", TG_BOT_TOKEN)
	}
	return token
}

func getChatIdConfig() int64 {
	chatId := os.Getenv(TG_BOT_CHAT_ID)
	if chatId == "" {
		log.Fatalf("config [%v] is empty\n", TG_BOT_CHAT_ID)
	}
	chatId1, _ := strconv.ParseInt(chatId, 10, 64)
	return chatId1
}

func getDebugConfig() bool {
	debug := os.Getenv(TG_BOT_DEBUG)
	if debug == "" {
		debug = "false"
	}
	debug1, _ := strconv.ParseBool(debug)
	return debug1
}

func getWebHookSecretConfig() string {
	v := os.Getenv(WEB_HOOK_SECRET)
	if v == "" {
		v = "feipugah9eu0"
	}
	return v
}

func getWebServerPort() string {
	v := os.Getenv(WEB_SERVER_PORT)
	if v == "" {
		return "80"
	}
	return v
}

func sendMessage(tgBot *TgBot, message string) {
	success, err := tgBot.sendMessage(message)
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("消息发送成功 %v", success)
}

func NewTgBotInstance() *TgBot {
	tgBot, err := NewTgBot(
		getTokenConfig(),
		getChatIdConfig(),
		getDebugConfig(),
	)
	if err != nil {
		log.Fatal(err)
	}
	return tgBot
}

var tgBot *TgBot

func main() {
	// 初始化的机器人实例
	tgBot = NewTgBotInstance()
	http.HandleFunc("/"+getWebHookSecretConfig(), requestHandle)
	log.Printf("服务启动成功,监听端口:%v\n", getWebServerPort())
	err := http.ListenAndServe(":"+getWebServerPort(), nil)
	if err != nil {
		return
	}
}

func requestHandle(response http.ResponseWriter, request *http.Request) {
	data, ioErr := io.ReadAll(request.Body)
	if ioErr != nil {
		handleResponse(response, ioErr)
		return
	}
	sendMessage(tgBot, string(data))
}

func handleResponse(response http.ResponseWriter, err error) {
	response.Header().Add("Content-Type", "application/json")
	if err == nil {
		_, _ = response.Write([]byte("{\"code\":\"0\",\"msg\":\"请求成功\"}"))
	} else {
		_, _ = response.Write([]byte("{\"code\":\"100\",\"msg\":\"" + err.Error() + "\"}"))
	}
}
