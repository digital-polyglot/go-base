package main

import "os"
import "go.bug.st/serial"
import "github.com/rs/zerolog"
import "github.com/rs/zerolog/log"

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	ports, err := serial.GetPortsList()
	if err != nil {
		log.Error().Err(err)
	}
	if len(ports) == 0 {
		log.Warn().Msg("No serial ports found!")
	}
	for _, port := range ports {
		log.Info().Str("port", port).Msg("")
	}
}
