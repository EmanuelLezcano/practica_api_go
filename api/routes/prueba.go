package routes

import (
	"github.com/Corrientes-Telecomunicaciones/api_go_recolector/api/middlewares"
	"github.com/Corrientes-Telecomunicaciones/api_go_recolector/pkg/domains/utils"

	"github.com/gofiber/fiber/v2"
)

func PruebaRoutes(app fiber.Router, middlewares middlewares.MiddlewareManager, utilService utils.UtilService) {
	app.Get("/prueba", getPruebas(utilService))

}

func getPruebas(utilService utils.UtilService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return ctx.JSON(&fiber.Map{
			"message": "Hola mundo",
		})
	}
}
