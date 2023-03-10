package middlewares

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/Corrientes-Telecomunicaciones/api_go_recolector/internal/config"

	"github.com/gofiber/fiber/v2"
)

type MiddlewareManager struct {
	HTTPClient *http.Client
}

func (m *MiddlewareManager) ValidarPermiso(scope string) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		bearer := c.Get("Authorization")
		if len(bearer) <= 0 {
			return errors.New("acceso no autorizado, debe enviar un token de autenticación")
		}
		var result struct {
			Acceso string `json:"acceso"`
			ID     int64  `json:"user_id"`
		}

		base, err := url.Parse(config.AUTH)
		if err != nil {
			return fmt.Errorf("error al crear base url" + err.Error())
		}

		//para login de alerta permisos y para loguin telco permiso
		base.Path += "/users/permisos"

		var values struct {
			SistemaID string `json:"sistema_id"`
			Scope     string `json:"permiso_slug"`
		}
		//FIXME definir como vamos a setear el sistemaId
		// idSitema, err := strconv.ParseInt(config.ID_SISTEMAS, 10, 64)
		// if err != nil {
		// 	return fmt.Errorf("error al convertir id sistema")
		// }
		values.SistemaID = config.ID_SISTEMAS
		values.Scope = scope
		json_data, _ := json.Marshal(values)

		req, _ := http.NewRequest("POST", base.String(), bytes.NewBuffer(json_data))

		req.Header.Add("Authorization", bearer)
		req.Header.Add("Content-Type", "application/json")

		resp, err := m.HTTPClient.Do(req)

		if err != nil {
			return fmt.Errorf("error al enviar solicitud a api externa")
		}

		if resp.StatusCode != 200 {
			info, _ := ioutil.ReadAll(req.Body)
			erro := fmt.Errorf("acceso denegado o permisos insuficientes: %s", info)
			return fiber.NewError(403, erro.Error())
		}

		json.NewDecoder(resp.Body).Decode(&result)
		c.Set("user_id", fmt.Sprint(result.ID))

		return c.Next()
	}
}
