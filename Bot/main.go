package main

import (
	"context"
	p "dogegambling/DataBase/Postgresql"
	r "dogegambling/DataBase/Redis"
	gateway "dogegambling/Gateway"
	"dogegambling/config"
	"dogegambling/handlers"
	"dogegambling/webhook"
	"flag"
	"fmt"
	"log"
	"os"

	"gopkg.in/telebot.v3"
	"gopkg.in/yaml.v3"
)

func main() {

	Bot := LoadConfigAndServeHandlers()
	go webhook.Serve(Bot)
	log.Println("Bot Started ....")
	Bot.Start()
}
func LoadConfigAndServeHandlers() *telebot.Bot {
	cfgPath, err := ParseFlags()
	if err != nil {
		log.Fatal(err)
	}
	cfg, err := NewConfig(cfgPath)
	if err != nil {
		log.Fatal(err)
	}

	handler := NewHandler(cfg)
	handler.Init()
	gateway.Init(cfg.BlockIO.Token, cfg.BlockIO.Pin)
	return handler.Bot
}

func NewHandler(c handlers.Config) handlers.Handler {
	db := handlers.Handler{
		RDB: r.InitRedisdb(c.Redis.User, c.Redis.Pass, c.Redis.Server, c.Redis.Port, c.Redis.DB),
		CTX: context.Background(),
		DB:  p.InitPostgredb(c.Postgresql.User, c.Postgresql.Pass, c.Postgresql.Server, c.Postgresql.Port, c.Postgresql.DB),
		Bot: config.InitBot(c.Bot.Token, c.Bot.Username, c.Bot.Gift, c.Bot.WithdrawChannelID, c.Bot.TransactionChannelID, c.Bot.Admins),
	}
	return db

}
func NewConfig(configPath string) (handlers.Config, error) {
	config := handlers.Config{}
	file, err := os.Open(configPath)
	if err != nil {
		return config, err
	}
	defer file.Close()

	d := yaml.NewDecoder(file)

	if err := d.Decode(&config); err != nil {
		return config, err
	}
	return config, nil
}

func ValidateConfigPath(path string) error {
	s, err := os.Stat(path)
	if err != nil {
		return err
	}
	if s.IsDir() {
		return fmt.Errorf("'%s' is a directory, not a normal file", path)
	}
	return nil
}
func ParseFlags() (string, error) {
	var configPath string
	flag.StringVar(&configPath, "config", "./config.yml", "path to config file")
	flag.Parse()

	if err := ValidateConfigPath(configPath); err != nil {
		return "", err
	}
	return configPath, nil
}
