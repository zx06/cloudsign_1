package vercel_v1

import (
	"net/http"
	"github.com/rs/zerolog/log"
)

func PingHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("pong"))
	if err != nil {
		log.Ctx(r.Context()).Error().Err(err).Msg("failed to write response")
	}
	log.Ctx(r.Context()).Info().Msg("ping")
}
