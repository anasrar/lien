package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/anasrar/lien/internal/router"
	"github.com/danielgtaylor/huma/v2/humacli"
	"github.com/gofiber/fiber/v2"
)

type Options struct {
	Port int `help:"Port to listen on" short:"p" default:"3000"`
}

func main() {
	cwd, err := os.Getwd()
	// TODO: argument for set the CWD
	if err != nil {
		log.Fatalln(err)
	}

	cli := humacli.New(func(hooks humacli.Hooks, options *Options) {
		app := fiber.New()
		router.Register(app, options.Port, cwd)

		hooks.OnStart(func() {
			app.Listen(fmt.Sprintf(":%d", options.Port))
		})

		hooks.OnStop(func() {
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			if err := app.ShutdownWithContext(ctx); err != nil {
				log.Fatalln(err)
			}
		})
	})

	cli.Run()
}
