package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rayhan889/go-template/api"
)

type DocsController struct{}

func NewDocsController() *DocsController {
	return &DocsController{}
}

func (c *DocsController) Spec(ctx *fiber.Ctx) error {
	ctx.Set(fiber.HeaderContentType, "application/yaml")
	return ctx.Send(api.OpenAPISpec)
}

func (c *DocsController) Reference(ctx *fiber.Ctx) error {
	ctx.Set(fiber.HeaderContentType, fiber.MIMETextHTMLCharsetUTF8)
	return ctx.SendString(scalarReferenceHtml)
}

const scalarReferenceHtml = `<!doctype html>
<html>
  <head>
    <title>go-template API Reference</title>
    <meta charset="utf-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1" />
  </head>
  <body>
    <script id="api-reference" data-url="/api/openapi.yaml"></script>
    <script src="https://cdn.jsdelivr.net/npm/@scalar/api-reference"></script>
  </body>
</html>
`
