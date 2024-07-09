package config

import (
	"log"
	"myapp/internal/client/http"
	"myapp/internal/repository/pg"
	"myapp/internal/service/tg_service"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Tg     tg_service.TgConfig
	Server http.SerConfig
	Db     pg.DBConfig
}

func Get() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("config Load .env err:", err)
	}
	var c Config

	c.Tg.TgEndp = os.Getenv("TG_ENDPOINT")
	c.Tg.Token = os.Getenv("BOT_TOKEN")
	c.Tg.ServerStatUrl = os.Getenv("SERVER_STAT_URL")
	c.Tg.BotId, _ = strconv.Atoi(os.Getenv("BOT_ID"))
	c.Tg.ServerUrl = os.Getenv("SERVER_URL")

	c.Server.Port = os.Getenv("APP_PORT")
	c.Db.User = os.Getenv("PG_USER")
	c.Db.Password = os.Getenv("PG_PASSWORD")
	c.Db.Database = os.Getenv("PG_DATABASE")
	c.Db.Host = os.Getenv("PG_HOST")
	c.Db.Port = os.Getenv("PG_PORT")

	/////////////////////////////////////////////////////////////////
	// c.TG_ENDPOINT = "https://api.telegram.org/bot%s/%s"
	// c.TOKEN       = ""
	// c.PORT        = ""
	// c.PG_USER     = ""
	// c.PG_PASSWORD = ""
	// c.PG_DATABASE = ""
	// c.PG_HOST     = ""

	return &c
}