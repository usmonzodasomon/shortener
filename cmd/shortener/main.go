package main

import (
	"fmt"
	"github.com/usmonzodasomon/shortener/internal/config"
)

func main() {
	cfg := config.MustLoad()
	fmt.Println(cfg)
}
