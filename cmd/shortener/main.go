package main

import (
	"github.com/usmonzodasomon/shortener/internal/config"
	"github.com/usmonzodasomon/shortener/pkg/logger"
)

func main() {
	cfg := config.MustLoad()
	logger := logger.Logger(*cfg)
	logger.Info("Hello, world!")
	logger.Error("Error occured!", "error", "some error")
	//fmt.Println(cfg)
}
