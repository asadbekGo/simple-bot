package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/asadbekGo/telegram-message-service/api"
	"github.com/asadbekGo/telegram-message-service/config"
	"github.com/asadbekGo/telegram-message-service/pkg/bot"
	"github.com/asadbekGo/telegram-message-service/pkg/db"
	"github.com/asadbekGo/telegram-message-service/pkg/logger"
	"github.com/asadbekGo/telegram-message-service/storage"
)

func main() {
	// Loading config
	cfg := config.Load()
	err := config.CreateFile()
	if err != nil {
		log.Println(err)
	}

	// Loading logger
	log := logger.New(cfg.Logger.LogLevel, "simple-bot-server")
	defer func(l logger.Logger) {
		err := logger.Cleanup(l)
		if err != nil {
			log.Fatal("failed cleanup logger", logger.Error(err))
		}
	}(log)

	// Logger postgres config
	log.Info("main: sqlxConfig",
		logger.String("host", cfg.Postgres.PostgresqlHost),
		logger.String("port", cfg.Postgres.PostgresqlPort),
		logger.String("database", cfg.Postgres.PostgresqlDbname),
	)

	// Creating postgresql connection
	connDB, err := db.ConnectionToPostgresDB(&cfg)
	if err != nil {
		log.Fatal("connection to postgres error")
	}

	// Logger telegram bot config
	log.Info("main: telegram bot",
		logger.String("token", cfg.TelegramBot.Token),
	)

	// New connection telegram bot
	tgBot, err := bot.NewConnectionBot(&cfg)
	if err != nil {
		log.Fatal("Error telegram bot connection: %v", logger.Error(err))
	}

	newUpdate := tgbotapi.NewUpdate(0)
	newUpdate.Timeout = 60

	serviceStorage := storage.NewStorage(connDB)

	updates := tgBot.GetUpdatesChan(newUpdate)

	api.NewRouter(api.Option{
		Conf:    cfg,
		Log:     log,
		Bot:     tgBot,
		Updates: updates,
		Storage: serviceStorage,
	})
}
