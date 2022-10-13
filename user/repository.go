package user

import (
	"go_webservice/infrastructure"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name       string `json:"name"`
	TelegramId int64  `json:"telegram_id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	ChatId     int64  `json:"chat_id"`
	Adverts    []Advert
}

type Advert struct {
	gorm.Model
	UserID   string  `json:"user_id"`
	CarMark  string  `json:"car_mark"`
	CarModel string  `json:"car_model"`
	Price    float64 `json:"price"`
	Year     int64   `json:"year"`
}

type DataHandler struct {
	sqlHandler *infrastructure.SqlHandler
}

// Создает объект обработчик
func NewDataHandler(sqlHandler *infrastructure.SqlHandler) (dh *DataHandler) {
	return &DataHandler{sqlHandler: sqlHandler}
}

// Создает нового пользователя
func (handler *DataHandler) Create(user User) error {
	sqlHandler := handler.sqlHandler.Db
	existUser := User{}
	result := sqlHandler.First(&existUser, User{TelegramId: user.TelegramId})

	if result.RowsAffected == 0 {
		createResult := sqlHandler.Create(&user)

		return createResult.Error
	}

	return nil
}

// Ищет одного пользователя в БД по логину телеграм
func (handler *DataHandler) FindOne(telegramId int64) *User {
	sqlHandler := handler.sqlHandler.Db
	user := User{}
	result := sqlHandler.Where("telegram_id = ?", telegramId).First(&user)

	if result.RowsAffected == 0 {
		return nil
	}

	return &user
}

// Ищет всех пользователей в БД по логину телеграм
func (handler *DataHandler) Find(telegramIds []string) []User {
	var users []User
	sqlHandler := handler.sqlHandler.Db
	result := sqlHandler.Where("name IN ?", telegramIds).Find(&users)

	if result.RowsAffected == 0 {
		return nil
	}

	return users
}

// Возращает все объявления пользователя
func (handler *DataHandler) GetAdverts(user *User) ([]Advert, error) {
	sqlHandler := handler.sqlHandler.Db

	err := sqlHandler.Model(user).Preload("Adverts").Find(&user).Error

	return user.Adverts, err
}
