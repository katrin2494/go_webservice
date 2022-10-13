package newsletter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go_webservice/infrastructure"
	"go_webservice/user"
	tele "gopkg.in/telebot.v3"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Newsletter struct {
	Message string `json:"message"`
	Users   string `json:"users"`
}

func Create(botToken string, message string, telegramIds string) (totalSent int) {
	userDh := user.NewDataHandler(infrastructure.NewSqlHandler())
	telegramIdsSlice := strings.Split(telegramIds, ",")
	users := userDh.Find(telegramIdsSlice)
	client := http.Client{
		Timeout: time.Duration(30 * time.Second),
	}

	totalSent = 0

	for _, value := range users {
		requestData := map[string]string{
			"chat_id": strconv.FormatInt(value.ChatId, 10),
			"text":    message,
		}

		requestBody, _ := json.Marshal(requestData)
		botUrl := fmt.Sprintf("%s/bot%s", tele.DefaultApiURL, botToken)
		resp, err := client.Post(botUrl+"/sendMessage", "application/json", bytes.NewBuffer(requestBody))

		if err != nil {
			log.Println(err)
			continue
		}

		if resp.StatusCode == 200 {
			totalSent++
		}

	}
	return totalSent
}
