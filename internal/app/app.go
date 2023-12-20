package app

import (
	"context"
	"errors"
	"fmt"
	"github.com/alserov/translator/internal/config"
	"github.com/alserov/translator/internal/controller"
	"github.com/alserov/translator/internal/controller/handlers"
	"github.com/alserov/translator/internal/db/clickhouse"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type app struct {
	port int
	dsn  string

	log *slog.Logger

	rTimeout time.Duration
	wTimeout time.Duration
}

type App interface {
	MustStart()
}

func NewApp(cfg *config.Config) App {
	return &app{
		port:     cfg.Port,
		dsn:      cfg.Dsn,
		rTimeout: cfg.ReadTimeout,
		wTimeout: cfg.WriteTimeout,
	}
}

func (a *app) MustStart() {
	db := clickhouse.MustConnect(a.dsn)
	repo := clickhouse.NewRepository(db)

	authHandler := handlers.NewAuthHandler(repo)
	translatorHandler := handlers.NewTranslatorHandler()
	controller.NewRouter(&controller.Handlers{
		Auth:       authHandler,
		Translator: translatorHandler,
	})

	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", a.port),
		ReadTimeout:  a.rTimeout,
		WriteTimeout: a.wTimeout,
	}

	chStop := make(chan os.Signal)
	signal.Notify(chStop, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := s.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			panic("failed to start server: " + err.Error())
		}
	}()

	fmt.Println("app is running")
	select {
	case <-chStop:
		if err := s.Shutdown(context.Background()); err != nil {
			panic("failed to shutdown server: " + err.Error())
		}
	}
}
