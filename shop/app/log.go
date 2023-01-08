package app

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
}

type Context int

const (
	Catalog Context = iota
	Shop
)

func (c Context) String() string {
	switch c {
	case Catalog:
		return "catalog"
	case Shop:
		return "shop"
	}
	return "unknown"
}

type Layer int

const (
	Api Layer = iota
	Domain
	Data
)

func (c Layer) String() string {
	switch c {
	case Api:
		return "api"
	case Domain:
		return "domain"
	case Data:
		return "data"
	}
	return "unknown"
}

func LoggerWith(context Context, layer Layer) zerolog.Logger {
	return log.With().Str("context", context.String()).Str("layer", layer.String()).Logger()
}
