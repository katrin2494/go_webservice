package service

import (
	"fmt"
	"go_webservice/infrastructure"
	"go_webservice/user"
	tele "gopkg.in/telebot.v3"
	"log"
	"time"
)

// Обработчик бота команда start
func startHandler(c tele.Context) error {
	newUser := user.User{
		Name:       c.Sender().Username,
		TelegramId: c.Sender().ID,
		FirstName:  c.Sender().FirstName,
		LastName:   c.Sender().LastName,
		ChatId:     c.Chat().ID,
	}

	err := user.NewDataHandler(infrastructure.NewSqlHandler()).Create(newUser)

	if err != nil {
		return c.Send("Не смогли вас узнать попробуйте позже")
	}

	return c.Send(
		"Привет, " + c.Sender().Username + ", тут ты можешь посмотреть свои объявления.",
	)
}

// Обработчик бота команда adverts
func advertsHandler(c tele.Context) error {
	userDh := user.NewDataHandler(infrastructure.NewSqlHandler())
	existUser := userDh.FindOne(c.Sender().ID)

	if existUser == nil {
		return c.Send("Не смогли вас узнать попробуйте позже")
	}

	adverts, err := userDh.GetAdverts(existUser)

	if err != nil {
		return c.Send("Не смогли найти объявления")
	}

	message := "Список твоих объявлений:\n"
	if len(adverts) > 0 {
		for i := range adverts {
			advert := &adverts[i]
			advertTitle := fmt.Sprintf("%d) %s %s %d г. %.0f тг.", advert.ID, advert.CarMark, advert.CarModel, advert.Year, advert.Price)
			message += advertTitle
		}
	} else {
		message += "пуст"
	}

	return c.Send(message)
}

// Обработчик бота команда newAdvert
func newAdvertHandler(c tele.Context) error {
	userDh := user.NewDataHandler(infrastructure.NewSqlHandler())
	existUser := userDh.FindOne(c.Sender().ID)

	if existUser == nil {
		c.Send("Не смогли вас узнать попробуйте позже")
	}

	c.Send("Для создания объявления введите данные объявления в формате: марка, модель, цена, год")

	c.Bot().Handle(tele.OnText, func(c tele.Context) error {
		return c.Send("Мы сохранили ваш ответ " + c.Message().Text)
	})

	return nil
}

// Создаем коннект к телеграм боту
func InitBot(token string) {
	pref := tele.Settings{
		Token:  token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)

	if err != nil {
		log.Println(err)
	}

	//Реализуем роутинг
	b.Handle("/start", startHandler)
	b.Handle("/adverts", advertsHandler)
	b.Handle("/newAdvert", newAdvertHandler)

	b.Start()
}
