package utils

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
}

type Context string
type Layer string

const (
	Catalog Context = "catalog"
	Shop    Context = "shop"

	Api    Layer = "api"
	Domain Layer = "domain"
	Data   Layer = "data"
)

func LoggerWith(context Context, layer Layer) zerolog.Logger {
	return log.With().Str("context", string(context)).Str("layer", string(layer)).Logger()
}
