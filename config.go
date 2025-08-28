package web

import (
	"github.com/awesome-goose/framework/contracts"
)

type Config struct {
	Name        string
	Version     string
	Author      string
	Description string

	Routes  contracts.Routes
	Host    string
	Port    int
	Timeout int
	Hooks   map[contracts.HookType]contracts.HookFn
}

type Option func(*Config)

// WithName sets the name of the .
func WithName(name string) Option {
	return func(Config *Config) {
		Config.Name = name
	}
}

// WithVersion sets the version of the .
func WithVersion(version string) Option {
	return func(Config *Config) {
		Config.Version = version
	}
}

// WithAuthor sets the author of the .
func WithAuthor(author string) Option {
	return func(Config *Config) {
		Config.Author = author
	}
}

// WithDescription sets the description of the .
func WithDescription(description string) Option {
	return func(Config *Config) {
		Config.Description = description
	}
}

func WithRoutes(routes contracts.Routes) Option {
	return func(Config *Config) {
		Config.Routes = routes
	}
}

func WithHost(host string) Option {
	return func(Config *Config) {
		Config.Host = host
	}
}

func WithPort(port int) Option {
	return func(Config *Config) {
		Config.Port = port
	}
}

func WithTimeout(timeout int) Option {
	return func(Config *Config) {
		Config.Timeout = timeout
	}
}

func WithHooks(hooks map[contracts.HookType]contracts.HookFn) Option {
	return func(Config *Config) {
		if Config.Hooks == nil {
			Config.Hooks = make(map[contracts.HookType]contracts.HookFn)
		}

		for hookType, hookFn := range hooks {
			Config.Hooks[hookType] = hookFn
		}
	}
}

func WithHook(hookType contracts.HookType, hookFn contracts.HookFn) Option {
	return func(Config *Config) {
		if Config.Hooks == nil {
			Config.Hooks = make(map[contracts.HookType]contracts.HookFn)
		}

		Config.Hooks[hookType] = hookFn
	}
}
