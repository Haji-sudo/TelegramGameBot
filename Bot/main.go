package main

import (
	"context"
	p "dogegambling/DataBase/Postgresql"
	r "dogegambling/DataBase/Redis"
	gateway "dogegambling/Gateway"
	"dogegambling/config"
	"dogegambling/handlers"
	"flag"
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

var ( //Color For Console
	colorGreen  = "\033[32m"
	colorCyan   = "\033[36m"
	colorReset  = "\033[0m"
	colorPurple = "\033[35m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
)

func main() {
	LoadConfigAndServeHandlers()
	fmt.Println(string(colorCyan), "\n\t Bot Started ....", string(colorReset))
	config.Bot.Start()

}

func LoadConfigAndServeHandlers() {
	cfgPath, err := ParseFlags()
	if err != nil {
		log.Fatal(err)
	}
	cfg, err := NewConfig(cfgPath)
	fmt.Println(string(colorGreen), "\n\t The config has been loaded .")
	// time.Sleep(time.Second * 1)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(colorPurple), "\n\t The Handlers Setting up .. ")
	// time.Sleep(time.Second * 1)

	handler := NewHandler(cfg)
	handler.Init()
	fmt.Println(string(colorBlue), "\n\t The Handlers Launched ...")

	fmt.Println(string(colorYellow), "\n\t The Gateway Setting up .... ")
	// time.Sleep(time.Second * 1)
	gateway.Init(cfg.BlockIO.Token, cfg.BlockIO.Pin, cfg.BlockIO.Webhook)
}

func NewHandler(c handlers.Config) handlers.Handler {
	db := handlers.Handler{
		RDB: r.InitRedisdb(c.Redis.User, c.Redis.Pass, c.Redis.Server, c.Redis.Port, c.Redis.DB),
		CTX: context.Background(),
		DB:  p.InitPostgredb(c.Postgresql.User, c.Postgresql.Pass, c.Postgresql.Server, c.Postgresql.Port, c.Postgresql.DB),
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
