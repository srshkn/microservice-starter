package main

import (
	"fmt"
	"gateway/internal/config"
)

func main() {
	cfg, _, err := config.New()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(cfg.Server.Host)
}
