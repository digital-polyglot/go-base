package main

import "testing"
import "go.bug.st/serial"
import "github.com/rs/zerolog"
import "github.com/rs/zerolog/log"
import "github.com/stretchr/testify/assert"

func TestPorts(t *testing.T) {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	ports, err := serial.GetPortsList()
	assert.Nil(t, err)
	assert.NotEqual(t, len(ports), 0, "No serial ports found!")
	for _, port := range ports {
		log.Info().Str("port", port).Msg("")
	}
}
