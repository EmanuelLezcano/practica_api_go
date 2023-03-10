package routes

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/Corrientes-Telecomunicaciones/api_go_recolector/api/middlewares"
	"github.com/Corrientes-Telecomunicaciones/api_go_recolector/internal/config"
	"github.com/Corrientes-Telecomunicaciones/api_go_recolector/pkg/domains/recolecciones"
	"github.com/Corrientes-Telecomunicaciones/api_go_recolector/pkg/dtos/recoleccionesdtos"
	filtros "github.com/Corrientes-Telecomunicaciones/api_go_recolector/pkg/filtros"
	filtros_recolecciones "github.com/Corrientes-Telecomunicaciones/api_go_recolector/pkg/filtros/recolecciones"

	"github.com/gofiber/fiber/v2"
)

func RecoleccionesRoutes(app fiber.Router, middlewares middlewares.MiddlewareManager, recoleccionesService recolecciones.RecoleccionesService) {
	app.Get("/configuraciones", getConfiguraciones(recoleccionesService))
	app.Get("/obtener-datos-ciudadano", getDatosUserApp(recoleccionesService))
	app.Get("/provincia", getProvincia(recoleccionesService))
	app.Get("/departamento", getDepartamento(recoleccionesService))
	app.Get("/localidad", getLocalidad(recoleccionesService))
	app.Get("/municipio", getMunicipio(recoleccionesService))
	app.Get("/barrio", getBarrio(recoleccionesService))
	app.Get("/planillas", getPlanillas(recoleccionesService))

	app.Put("/planilla-estado", putPlanillaEstado(recoleccionesService))
	app.Put("/recorrido-estado", putRecorridoEstado(recoleccionesService))

	app.Post("/registrar-barrio", middlewares.ValidarPermiso("recoleccion.create"), postBarrio(recoleccionesService))
	// app.Post("/recoleccion", middlewares.ValidarPermiso("recoleccion.create"), postRecoleccion(recoleccionesService))
	app.Post("/recoleccion", postRecoleccion(recoleccionesService))
	app.Get("/recoleccion", getRecoleccion(recoleccionesService))
	app.Put("/recoleccion", putRecoleccion(recoleccionesService))
	// asociar imagenes a una recoleccion
	app.Post("/images-recoleccion", postImagesRecoleccion(recoleccionesService))
}

func getConfiguraciones(recoleccionesService recolecciones.RecoleccionesService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var request string
		var status bool
		var msj string
		request = ctx.Query("MunicipioId", "")

		if len(request) == 0 {
			return fiber.NewError(400, "error al recuperar parametro enviado.")
		}
		municipio_id, err := strconv.Atoi(request)
		if err != nil {
			return fiber.NewError(400, "error al convertir parametro enviado: "+err.Error())
		}
		response, err := recoleccionesService.GetConfiguracionesByMunicipioIdService(uint(municipio_id))
		if err != nil {
			return fiber.NewError(400, "error: "+err.Error())
		}
		status = true
		msj = "datos obtenidos con exito"
		if len(response.Configuraciones) == 0 {
			status = false
			msj = "no existe configuracion registrada"
			return ctx.Status(200).JSON(&fiber.Map{
				"status":  status,
				"message": msj,
				"data":    response.Configuraciones,
			})
		}
		return ctx.Status(200).JSON(&fiber.Map{
			"status":  status,
			"message": msj,
			"data":    response.Configuraciones,
		})
	}
}

func getDatosUserApp(recoleccionesService recolecciones.RecoleccionesService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var request recoleccionesdtos.RequestDatosCiudadanos
		var status bool
		var msj string
		err := ctx.QueryParser(&request)
		if err != nil {
			return fiber.NewError(400, "error al obtener parametro: "+err.Error())
		}
		responseAppUser, rowAffected, err := recoleccionesService.GetAppUserByUserIdService(request.CiudadanoId)
		if err != nil {
			return fiber.NewError(400, "error : "+err.Error())
		}
		status = true
		msj = "datos obtenidos con exito"
		if rowAffected {
			status = false
			msj = "no existe registros"
			return ctx.Status(200).JSON(&fiber.Map{
				"status":  status,
				"message": msj,
				"data":    responseAppUser,
			})
		}
		return ctx.Status(200).JSON(&fiber.Map{
			"status":  status,
			"message": msj,
			"data":    responseAppUser,
		})
	}
}

func getProvincia(recoleccionesService recolecciones.RecoleccionesService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var (
			status bool
			msj    string
		)

		// filtros de provincia
		var request filtros.ProvinciaFiltro
		err := ctx.QueryParser(&request)
		if err != nil {
			return fiber.NewError(400, "Error en los parámetros enviados: "+err.Error())
		}

		response, err := recoleccionesService.GetProvinciaService(request)

		if err != nil {
			return fiber.NewError(400, "error: "+err.Error())
		}
		status = true
		msj = "datos obtenidos con exito"
		if len(response.Provincias) == 0 {
			status = false
			msj = recolecciones.ERROR_GET_PROVINCIA
			return ctx.Status(200).JSON(&fiber.Map{
				"status":  status,
				"message": msj,
				"data":    "",
			})
		}

		return ctx.Status(200).JSON(&fiber.Map{
			"status":  status,
			"message": msj,
			"data":    response,
		})
	}
}

func getDepartamento(recoleccionesService recolecciones.RecoleccionesService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var (
			status bool
			msj    string
		)

		// filtros de departamento
		var request filtros.DepartamentoFiltro
		err := ctx.QueryParser(&request)
		if err != nil {
			return fiber.NewError(400, "Error en los parámetros enviados: "+err.Error())
		}

		response, err := recoleccionesService.GetDepartamentoService(request)
		if err != nil {
			return fiber.NewError(400, "error: "+err.Error())
		}
		status = true
		msj = "datos obtenidos con exito"
		if len(response.Departamentos) == 0 {
			status = false
			msj = recolecciones.ERROR_GET_DEPARTAMENTO
			return ctx.Status(200).JSON(&fiber.Map{
				"status":  status,
				"message": msj,
				"data":    response.Departamentos,
			})
		}
		return ctx.Status(200).JSON(&fiber.Map{
			"status":  status,
			"message": msj,
			"data":    response,
		})
	}
}

func getLocalidad(recoleccionesService recolecciones.RecoleccionesService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var (
			status bool
			msj    string
		)

		// filtros de localidades
		var request filtros.LocalidadFiltro
		err := ctx.QueryParser(&request)
		if err != nil {
			return fiber.NewError(400, "Error en los parámetros enviados: "+err.Error())
		}

		response, err := recoleccionesService.GetLocalidadService(request)
		if err != nil {
			return fiber.NewError(400, "error: "+err.Error())
		}
		status = true
		msj = "datos obtenidos con exito"
		if len(response.Localidades) == 0 {
			status = false
			msj = recolecciones.ERROR_GET_LOCALIDAD
			return ctx.Status(200).JSON(&fiber.Map{
				"status":  status,
				"message": msj,
				"data":    response.Localidades,
			})
		}
		return ctx.Status(200).JSON(&fiber.Map{
			"status":  status,
			"message": msj,
			"data":    response,
		})
	}
}

func getMunicipio(recoleccionesService recolecciones.RecoleccionesService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var (
			status bool
			msj    string
		)

		var request filtros.MunicipioFiltro
		err := ctx.QueryParser(&request)
		if err != nil {
			return fiber.NewError(400, "Error en los parámetros enviados: "+err.Error())
		}
		response, err := recoleccionesService.GetMunicipioService(request)
		if err != nil {
			return fiber.NewError(400, "error: "+err.Error())
		}
		status = true
		msj = "datos obtenidos con exito"
		if len(response.Municipios) == 0 {
			status = false
			msj = recolecciones.ERROR_GET_MUNICIPIO
			return ctx.Status(200).JSON(&fiber.Map{
				"status":  status,
				"message": msj,
				"data":    response.Municipios,
			})
		}
		return ctx.Status(200).JSON(&fiber.Map{
			"status":  status,
			"message": msj,
			"data":    response,
		})
	}
}

func getBarrio(recoleccionesService recolecciones.RecoleccionesService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var (
			status bool
			msj    string
		)

		var request filtros.BarrioFiltro
		err := ctx.QueryParser(&request)
		if err != nil {
			return fiber.NewError(400, "Error en los parámetros enviados: "+err.Error())
		}
		response, err := recoleccionesService.GetBarrioService(request)
		if err != nil {
			return fiber.NewError(400, "error: "+err.Error())
		}
		status = true
		msj = "datos obtenidos con exito"
		if len(response.Barrios) == 0 {
			status = false
			msj = recolecciones.ERROR_GET_BARRIO
			return ctx.Status(200).JSON(&fiber.Map{
				"status":  status,
				"message": msj,
				"data":    response.Barrios,
			})
		}
		return ctx.Status(200).JSON(&fiber.Map{
			"status":  status,
			"message": msj,
			"data":    response,
		})
	}
}

func getPlanillas(recoleccionesService recolecciones.RecoleccionesService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var (
			status bool
			msj    string
		)

		// filtros de planillas
		var request filtros.PlanillaFiltro
		err := ctx.QueryParser(&request)
		if err != nil {
			return fiber.NewError(400, "Error en los parámetros enviados: "+err.Error())
		}

		response, err := recoleccionesService.GetPlanillasService(request)

		if err != nil {
			return fiber.NewError(400, "error: "+err.Error())
		}
		status = true
		msj = "datos obtenidos con exito"

		return ctx.Status(200).JSON(&fiber.Map{
			"status":  status,
			"message": msj,
			"data":    response,
		})
	}
}

func putRecorridoEstado(recoleccionesService recolecciones.RecoleccionesService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var (
			status bool
			msj    string
		)

		// filtros de planillas
		var request recoleccionesdtos.RequestEstadoRecorrido
		err := ctx.BodyParser(&request)
		if err != nil {
			return fiber.NewError(400, "Error en los parámetros enviados: "+err.Error())
		}

		err = request.Validar()
		if err != nil {
			return fiber.NewError(400, err.Error())
		}

		response, err := recoleccionesService.PutEstadoRecorridoService(request)

		if err != nil {
			return fiber.NewError(400, "error: "+err.Error())
		}
		status = true
		msj = "datos actualizados con exito"

		return ctx.Status(200).JSON(&fiber.Map{
			"status":  status,
			"message": msj,
			"data":    response,
		})
	}
}

func putPlanillaEstado(recoleccionesService recolecciones.RecoleccionesService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var (
			status bool
			msj    string
		)

		// filtros de planillas
		var request recoleccionesdtos.RequestEstadoPlanilla
		err := ctx.BodyParser(&request)
		if err != nil {
			return fiber.NewError(400, "Error en los parámetros enviados: "+err.Error())
		}

		err = request.Validar()
		if err != nil {
			return fiber.NewError(400, err.Error())
		}

		response, err := recoleccionesService.PutEstadoPlanillaService(request)

		if err != nil {
			return fiber.NewError(400, "error: "+err.Error())
		}
		status = true
		msj = "datos actualizados con exito"

		return ctx.Status(200).JSON(&fiber.Map{
			"status":  status,
			"message": msj,
			"data":    response,
		})
	}
}

func postBarrio(recoleccionesService recolecciones.RecoleccionesService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		ctx.Accepts("application/json")
		var (
			request recoleccionesdtos.RequestRegistroAppUser
			status  bool
			msj     string
		)
		err := ctx.BodyParser(&request)
		if err != nil {
			return fiber.NewError(400, "error al obtener parametro: "+err.Error())
		}
		request.Validar()
		if err != nil {
			return fiber.NewError(400, "error - "+err.Error())
		}
		status, err = recoleccionesService.PostRegistrarAppUsuarioService(request)
		if err != nil {
			return fiber.NewError(400, "error: "+err.Error())
		}
		msj = "datos registrado con exito"
		return ctx.Status(200).JSON(&fiber.Map{
			"status":  status,
			"message": msj,
		})
	}
}

func postRecoleccion(recoleccionesService recolecciones.RecoleccionesService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		ctx.Accepts("application/json")
		var (
			request recoleccionesdtos.RequestPostRecoleccion
			status  bool
			msj     string
			id      uint
		)

		err := ctx.BodyParser(&request)

		if err != nil {
			return fiber.NewError(400, "error al obtener parametro: "+err.Error())
		}

		id, status, err = recoleccionesService.PostRecoleccionService(request)
		if err != nil {
			return fiber.NewError(400, "error: "+err.Error())
		}
		msj = "dato registrado con exito"
		return ctx.Status(200).JSON(&fiber.Map{
			"status":  status,
			"message": msj,
			"id":      id,
		})
	}
}

func getRecoleccion(recoleccionesService recolecciones.RecoleccionesService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var (
			status bool
			msj    string
		)

		var filtro filtros_recolecciones.GetRecoleccioneFiltro

		err := ctx.QueryParser(&filtro)

		if err != nil {
			return fiber.NewError(400, "Error en los parámetros enviados: "+err.Error())
		}

		response, rowAffected, err := recoleccionesService.GetRecoleccionService(filtro)

		if err != nil {
			return fiber.NewError(400, "error: "+err.Error())
		}

		if rowAffected {
			status = false
			msj = "no existe registro"
			return ctx.Status(200).JSON(&fiber.Map{
				"status":  status,
				"message": msj,
				"data":    "",
			})
		}

		return ctx.Status(200).JSON(&fiber.Map{
			"status":  true,
			"message": "datos obtenidos con exito",
			"data":    response,
		})
	}
}

func putRecoleccion(recoleccionesService recolecciones.RecoleccionesService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		ctx.Accepts("application/json")
		var (
			request recoleccionesdtos.RequestPostRecoleccion
			status  bool
			msj     string
		)

		err := ctx.BodyParser(&request)

		if err != nil {
			return fiber.NewError(400, "error al obtener parametro: "+err.Error())
		}

		status, err = recoleccionesService.PutRecoleccionService(request)
		if err != nil {
			return fiber.NewError(400, "error: "+err.Error())
		}
		msj = "dato modificado con exito"
		return ctx.Status(200).JSON(&fiber.Map{
			"status":  status,
			"message": msj,
		})
	}
}

// recibe imagen con recoleccion id y name y guarda en tabla y storage
func postImagesRecoleccion(recoleccionesService recolecciones.RecoleccionesService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		var (
			id  uint
			res bool
		)
		// estructuta vacia del dto
		request := recoleccionesdtos.RequestImagesRecoleccionesDTO{}

		err := ctx.BodyParser(&request)

		if err != nil {
			return fiber.NewError(400, "error al obtener parametro: "+err.Error())
		}

		// consultar si el registro de recoleccion existe
		res, err = recoleccionesService.CheckIfModelExists("recolecciones", request.RecoleccionesId)

		if err != nil {
			return fiber.NewError(400, "error al consultar registro de recoleccion: "+err.Error())
		}
		if !res {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":   true,
				"message": "error el registro de recoleccion no existe",
				"status":  false,
			})
			// return fiber.NewError(404, "error el registro de recoleccion no existe")
		}

		// obtener el archvio de la request
		file, err := ctx.FormFile("Images")

		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": true,
				"msg":   err.Error(),
			})
		}

		fecha := fmt.Sprintf("%v_%v-%v-%v_", time.Now().Format("02-Jan-2006"), time.Now().Hour(), time.Now().Minute(), time.Now().Second())

		if len(request.FileName) == 0 {
			request.FileName = fecha + file.Filename
		} else {
			request.FileName = fecha + request.FileName
		}

		// Se crea la carpeta temporal en caso de que no exista
		carpetaTempImages := config.DIR_BASE + config.DIR_IMAGES_RECOLECCION
		if _, err := os.Stat(carpetaTempImages); os.IsNotExist(err) {
			err = os.MkdirAll(carpetaTempImages, 0755)
			if err != nil {
				return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error":   true,
					"message": err.Error(),
					"status":  false,
				})
			}
		}

		// guardar el archivo en carpeta temporal
		err = ctx.SaveFile(file, fmt.Sprintf("%s%s/%s", config.DIR_BASE, config.DIR_IMAGES_RECOLECCION, request.FileName))
		if err != nil {

			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":   true,
				"message": err.Error(),
				"status":  false,
			})
		}
		request.Path = fmt.Sprintf("%s%s/%s", config.DIR_BASE, config.DIR_IMAGES_RECOLECCION, request.FileName)
		// Comunicarse al servicio. pasar la request
		id, err = recoleccionesService.PostImagesRecoleccionService(request)

		// Capturar el error del servicio
		if err != nil {
			return fiber.NewError(400, "error: "+err.Error())
		}

		// Responder
		// msj = "dato registrado con exito"
		return ctx.Status(200).JSON(&fiber.Map{
			"status":  true,
			"message": "Se guardaron las imagenes de la recoleccion",
			"id":      id,
		})

	}
}
