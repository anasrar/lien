package router

import (
	"fmt"

	"github.com/danielgtaylor/huma/v2"
	"github.com/gofiber/fiber/v2"
)

func RegisterCustomOpenApiDocScalar(c *huma.Config, port int, r fiber.Router, subdir string) {
	c.Servers = []*huma.Server{&huma.Server{
		URL:         `http://localhost:{port}` + subdir,
		Description: `Development`,
		Variables: map[string]*huma.ServerVariable{
			"port": &huma.ServerVariable{
				Default: fmt.Sprintf("%d", port),
			},
		},
	}}

	c.DocsPath = ""
	r.Get("/docs", func(c *fiber.Ctx) error {
		c.Response().Header.Set("Content-Type", "text/html")
		return c.SendString(`<!doctype html>
<html>
  <head>
    <title>Scalar API Reference</title>
    <meta charset="utf-8" />
    <meta
      name="viewport"
      content="width=device-width, initial-scale=1" />
  </head>

  <body>
    <div id="app"></div>

    <!-- Load the Script -->
    <script src="https://cdn.jsdelivr.net/npm/@scalar/api-reference"></script>

    <!-- Initialize the Scalar API Reference -->
    <script>
      Scalar.createApiReference('#app', {
        telemetry: false,
        hideClientButton: true,
        showDeveloperTools: 'localhost',
        sources: [{
          url: './openapi.yaml',
        }],
      })
    </script>
  </body>
</html>`)
	})
}

func RegisterCustomOpenApiDocStoplight(c *huma.Config, port int, r fiber.Router, subdir string) {
	c.Servers = []*huma.Server{&huma.Server{
		URL:         `http://localhost:{port}` + subdir,
		Description: `Development`,
		Variables: map[string]*huma.ServerVariable{
			"port": &huma.ServerVariable{
				Default: fmt.Sprintf("%d", port),
			},
		},
	}}
	c.DocsPath = ""

	title := c.Info.Title
	r.Get("/docs", func(c *fiber.Ctx) error {
		c.Response().Header.Set("Content-Type", "text/html")
		return c.SendString(`<!doctype html>
<html lang="en">
  <head>
    <meta charset="utf-8" />
    <meta name="referrer" content="same-origin" />
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no" />
    <title>` + title + ` Reference</title>
    <!-- Embed elements Elements via Web Component -->
    <link href="https://unpkg.com/@stoplight/elements@9.0.0/styles.min.css" rel="stylesheet" />
    <script src="https://unpkg.com/@stoplight/elements@9.0.0/web-components.min.js" integrity="sha256-Tqvw1qE2abI+G6dPQBc5zbeHqfVwGoamETU3/TSpUw4="
            crossorigin="anonymous"></script>
  </head>
  <body style="height: 100vh;">

    <elements-api
      apiDescriptionUrl="./openapi.yaml"
      router="hash"
      layout="sidebar"
      tryItCredentialsPolicy="same-origin"
    />

  </body>
</html>
`)
	})
}
