package web

import "github.com/awesome-goose/contracts"

type Platform struct {
	config *Config
}

func NewPlatform(options ...Option) *Platform {
	config := &Config{}

	for _, option := range options {
		option(config)
	}

	return &Platform{}
}

func (p *Platform) Boot(container contracts.Container) (contracts.App, error) {
	app := NewApp(p.config)

	return app, nil
}
