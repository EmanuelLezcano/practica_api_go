package entities

import (
	"fmt"
	"strings"

	"github.com/Corrientes-Telecomunicaciones/api_go_recolector/pkg/dtos/tools"
	"gorm.io/gorm"
)

type Configuracione struct {
	gorm.Model
	MunicipiosId    uint
	BotonesId       uint
	Nombre          string
	Descripcion     string
	Valor           string
	DiasRecoleccion string
	Boton           Botone `gorm:"foreignKey:botones_id"`
}

func (c *Configuracione) IsValid() error {

	if tools.EsStringVacio(c.Nombre) {
		return fmt.Errorf("el campo nombre es obligatorio")
	}
	if tools.EsStringVacio(c.Valor) {
		return fmt.Errorf("el campo valor es obligatorio")
	}

	c.Nombre = strings.ToUpper(c.Nombre)
	c.Nombre = strings.TrimSpace(c.Nombre)

	return nil
}
