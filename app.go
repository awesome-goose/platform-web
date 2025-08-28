package web

import (
	"fmt"
	"net/http"

	"github.com/awesome-goose/framework/contracts"
)

type App struct {
	config *Config
	fn     func(c contracts.Context) error
}

func NewApp(config *Config) *App {
	return &App{config: config}
}

func (a *App) Run(fn func(c contracts.Context) error) error {
	a.fn = fn
	server := &http.Server{Addr: fmt.Sprintf("%s:%d", a.config.Host, a.config.Port), Handler: a}

	return server.ListenAndServe()
}

func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	context := NewContext(w, r)
	err := a.fn(context)
	if err != nil {
		// write to w http.ResponseWriter
		return
	}
}
