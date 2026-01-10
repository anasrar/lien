package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/anasrar/lien/internal/router"
	"github.com/danielgtaylor/huma/v2/humacli"
	"github.com/gofiber/fiber/v2"
)

type Options struct {
	Port int    `help:"Port to listen on" short:"p" default:"3000"`
	Root string `help:"Root path to serve" short:"r" default:"./"`
}

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	cli := humacli.New(func(hooks humacli.Hooks, options *Options) {
		if filepath.IsAbs(options.Root) {
			cwd = filepath.Clean(options.Root)
		} else {
			cwd = filepath.Join(cwd, options.Root)
		}

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
