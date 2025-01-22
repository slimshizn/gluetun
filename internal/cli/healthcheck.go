package cli

import (
	"context"
	"net"
	"net/http"
	"time"

	"github.com/qdm12/gluetun/internal/configuration/settings"
	"github.com/qdm12/gluetun/internal/healthcheck"
	"github.com/qdm12/gosettings/reader"
)

func (c *CLI) HealthCheck(ctx context.Context, reader *reader.Reader, _ Warner) (err error) {
	// Extract the health server port from the configuration.
	var config settings.Health
	err = config.Read(reader)
	if err != nil {
		return err
	}

	config.SetDefaults()

	err = config.Validate()
	if err != nil {
		return err
	}

	_, port, err := net.SplitHostPort(config.ServerAddress)
	if err != nil {
		return err
	}

	const timeout = 10 * time.Second
	httpClient := &http.Client{Timeout: timeout}
	client := healthcheck.NewClient(httpClient)
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	url := "http://127.0.0.1:" + port
	return client.Check(ctx, url)
}
