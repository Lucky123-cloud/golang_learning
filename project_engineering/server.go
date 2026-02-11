package main

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/99designs/gqlgen/graphql/handler/apollofederatedtracingv1/logger"
)

var (
	PortOutOfRange = errors.New("config: error occured, port out of range")
	PortNotSet     = errors.New("config: received empty port")
)

type Config struct {
	port            string
	ReadTimeOut     time.Duration
	ReadHeadTimeOut time.Duration
	WriteTimeOut    time.Duration
	maxBytes        int64
	idolTimeout     time.Duration
	shutDownTimeout time.Duration
}

type Server struct {
	log   logger.Logger
	srv   *http.Server
	cfg   Config
	route []routeInfo
}

type routeInfo struct {
	method      string
	path        string
	handler     http.HandlerFunc
	maxBodySize int64
}

// New create an instance of a Server
func New(
	ctx context.Context,
	log *logger.Logger,
	cfg Config,
) (*Server, error) {
	err := prepareConfig(&cfg)

	if err != nil {
		return nil, err
	}
}

// Validates server config and sets defaults
func prepareConfig(cfg *Config) error {
	err := ValidatePort(cfg)
	if err != nil {
		return err
	}

	if cfg.maxBytes <= 0 {
		cfg.maxBytes = 10 * 1024 * 1024
	}
}

func ValidatePort(cfg *Config) error {
	if cfg.port == "" {

	}
}
