package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/hamba/cmd"
	"github.com/msales/kage"
	"github.com/msales/kage/server"
	"github.com/urfave/cli/v2"
)

func runServer(c *cli.Context) error {
	ctx, err := cmd.NewContext(c)
	if err != nil {
		return err
	}

	app, err := newApplication(ctx)
	if err != nil {
		return err
	}
	defer app.Close()

	monitorTicker := time.NewTicker(30 * time.Second)
	defer monitorTicker.Stop()
	go func() {
		for range monitorTicker.C {
			app.Collect()
		}
	}()

	reportTicker := time.NewTicker(60 * time.Second)
	defer reportTicker.Stop()
	go func() {
		for range reportTicker.C {
			app.Report()
		}
	}()

	if c.Bool(FlagServer) {
		port := c.String(cmd.FlagPort)
		srv := newServer(app)
		h := http.Server{Addr: ":" + port, Handler: srv}
		defer func() {
			h.Shutdown(context.Background())
		}()
		go func() {
			ctx.Logger().Info("Starting on port " + port)
			if err := h.ListenAndServe(); err != nil {
				if err != http.ErrServerClosed {
					log.Fatal(err)
				}
			}
		}()
	}

	<-cmd.WaitForSignals()

	return nil
}

func newServer(app *kage.Application) http.Handler {
	return server.New(app)
}
