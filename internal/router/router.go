package router

import (
	"io/fs"
	"log"
	"net/http"

	controlerv1 "github.com/anasrar/lien/internal/controller/v1"
	"github.com/anasrar/lien/internal/webui"
	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humafiber"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

func Register(
	app *fiber.App,
	port int,
	cwd string,
	resumable bool,
	foreceDownload bool,
) {
	app.Route("/api/v1", func(r fiber.Router) {
		config := huma.DefaultConfig("API v1", "0.0.0")
		RegisterCustomOpenApiDocScalar(&config, port, r, "/api/v1")
		api := humafiber.NewWithGroup(app, r, config)

		api.UseMiddleware(func(ctx huma.Context, next func(huma.Context)) {
			ctx = huma.WithValue(ctx, "cwd", cwd)
			next(ctx)
		})

		huma.Register(api, controlerv1.GetDirectoryEntriesOperation, controlerv1.GetDirectoryEntriesHandler)
	})

	app.Static("/dl", cwd, fiber.Static{
		Compress:  false,
		ByteRange: resumable,
		Browse:    false,
		Download:  foreceDownload,
	})

	{
		root, err := fs.Sub(webui.BuildFiles, "build")
		if err != nil {
			log.Panic(err)
		}

		app.Use("/", filesystem.New(filesystem.Config{
			Root:   http.FS(root),
			Browse: false,
		}))
	}
}
