package main

import (
	"net"
	"net/http"
	"os"
	"time"

	"github.com/Corrientes-Telecomunicaciones/api_go_recolector/api/middlewares"
	"github.com/Corrientes-Telecomunicaciones/api_go_recolector/api/routes"
	"github.com/Corrientes-Telecomunicaciones/api_go_recolector/internal/config"
	"github.com/Corrientes-Telecomunicaciones/api_go_recolector/internal/database"
	"github.com/Corrientes-Telecomunicaciones/api_go_recolector/internal/storage"
	"github.com/Corrientes-Telecomunicaciones/api_go_recolector/pkg/commons"
	"github.com/Corrientes-Telecomunicaciones/api_go_recolector/pkg/domains/dbhelpers"
	"github.com/Corrientes-Telecomunicaciones/api_go_recolector/pkg/domains/recolecciones"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	//"github.com/gofiber/template/html"
)

func InicializarApp(clienteHttp *http.Client, clienteSql *database.MySQLClient, clienteFile *os.File) *fiber.App {
	middlewares := middlewares.MiddlewareManager{HTTPClient: clienteHttp}

	//***** Servicios Comunes *****//
	fileRepository := commons.NewFileRepository(clienteFile)
	commonsService := commons.NewCommons(fileRepository)

	//***** Almacenamiento *****//
	minioStorage := storage.NewMinioSession() // Objeto storage minio que implementa la interfaz Storage

	//***** Repositories *****//
	configuracionesRepository := recolecciones.NewRecoleccionesRepository(clienteSql)
	databaseHelper := dbhelpers.NewDbHelpersRepository(clienteSql) // funciones auxiliares de consultas a DB

	//***** Services *****//
	configuracionesService := recolecciones.NewRecoleccionesService(configuracionesRepository, databaseHelper, minioStorage, commonsService)

	app := fiber.New(fiber.Config{
		//Views: engine,
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			var msg string
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
				msg = e.Message
			}

			if msg == "" {
				msg = "No se pudo procesar el llamado a la api: " + err.Error()
			}

			_ = ctx.Status(code).JSON(internalError{
				Message: msg,
			})

			return nil
		},
	})
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: config.ALLOW_ORIGIN + ", " + config.AUTH,
		AllowHeaders: "",
		AllowMethods: "GET,POST,PUT,DELETE",
	}))
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Send([]byte("Corrientes Telecomunicaciones Api Servicio " + config.API_NOMBRE))
	})

	api := app.Group("/api/v1")
	//routes.PruebaRoutes(api, middlewares, utilService)
	routes.RecoleccionesRoutes(api, middlewares, configuracionesService)

	return app
}

func main() {
	var HTTPTransport http.RoundTripper = &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		ForceAttemptHTTP2:     false, // <- this is my adjustment
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	var HTTPClient = &http.Client{
		Transport: HTTPTransport,
	}

	//HTTPClient.Timeout = time.Second * 120 //Todo validar si este tiempo está bien
	clienteSQL := database.NewMySQLClient()
	osFile := os.File{}

	app := InicializarApp(HTTPClient, clienteSQL, &osFile)
	// el puerto puede que se necesite cambiar
	_ = app.Listen(":3300")
}

type internalError struct {
	Message string `json:"message"`
}
