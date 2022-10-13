package service

import (
	"encoding/json"
	"fmt"
	"go_webservice/newsletter"
	"log"
	"net/http"
)

type Server struct {
	botToken string
}

// Обработчик урла /_health
func heathHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Server is alive")
}

// Обработчик урла /sendNewsletter
func (server *Server) sendNewsletterHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(w, fmt.Sprintf("expect POST, got %v", req.Method), http.StatusMethodNotAllowed)
		return
	}

	var body newsletter.Newsletter
	err := json.NewDecoder(req.Body).Decode(&body)

	if err != nil {
		log.Println(err)
	}

	totalSent := newsletter.Create(server.botToken, body.Message, body.Users)

	fmt.Fprintf(w, "Было отправлено %d сообщений", totalSent)
}

// Создаем новый сервер для получения команд от админки
func InitServer(host string, port string, botToken string) {
	mux := http.NewServeMux()
	server := Server{botToken: botToken}

	//Реализуем роутинг
	mux.HandleFunc("/_health", heathHandler)
	mux.HandleFunc("/sendNewsletter", server.sendNewsletterHandler)

	if err := http.ListenAndServe(fmt.Sprintf("%s:%v", host, port), mux); err != nil {
		log.Println("Ошибка запуска сервера")
	}
}
