package main

import (
	"github.com/alserov/translator/internal/app"
	"github.com/alserov/translator/internal/config"
)

func main() {
	cfg := config.MustLoad()

	application := app.NewApp(cfg)
	application.MustStart()
}
