package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {

	token := os.Getenv("TELEGRAM_BOT_TOKEN")

	if token == "" {
		log.Fatal("Токен бота не задан! Задайте переменную TELEGRAM_BOT_TOKEN или укажите токен в коде")
	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic("Ошибка создания бота: ", err)
	}

	bot.Debug = false // Измените на true для отладки
	log.Printf("Бот %s успешно запущен", bot.Self.UserName)

	// Настройка получения обновлений
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	// Основной цикл обработки сообщений
	for update := range updates {
		// Игнорируем любые не-текстовые сообщения
		if update.Message == nil || update.Message.Text == "" {
			continue
		}

		// Обработка команд
		if update.Message.IsCommand() {
			cmd := update.Message.Command()
			if cmd == "start" {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Привет! Я эхо-бот. Просто напиши мне что-нибудь")
				bot.Send(msg)
			}
			continue
		}

		// ---- ОСНОВНАЯ ЛОГИКА БОТА ----
		// Эхо-ответ с оригинальным текстом
		echoText := update.Message.Text

		// Создаем ответное сообщение
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, echoText)
		msg.ReplyToMessageID = update.Message.MessageID

		// Отправляем сообщение
		if _, err := bot.Send(msg); err != nil {
			log.Printf("Ошибка отправки: %v", err)
		}
	}
}
