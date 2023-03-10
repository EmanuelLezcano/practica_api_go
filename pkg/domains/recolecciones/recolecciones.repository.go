package recolecciones

import (
	"errors"
	"strings"

	"github.com/Corrientes-Telecomunicaciones/api_go_recolector/internal/database"

	"github.com/Corrientes-Telecomunicaciones/api_go_recolector/pkg/entities"
	filtros "github.com/Corrientes-Telecomunicaciones/api_go_recolector/pkg/filtros"
	filtros_recolecciones "github.com/Corrientes-Telecomunicaciones/api_go_recolector/pkg/filtros/recolecciones"
	"gorm.io/gorm/clause"
)

// defino una interface
type RecoleccionesRepository interface {
	/*
		permite obtener la configuracion de botones por municipio desde la base de datos
		recibe como parametro municipioId y responde una lista de entidad configuraciones consus relaciones (botones y tipo)
	*/
	GetConfiguracionesByMunicipioIdRepository(municipioId uint) (ListEntityConfiguraciones []entities.Configuracione, erro error)
	GetProvinciaRepository(filtro filtros.ProvinciaFiltro) (provincias []entities.Provincias, erro error)
	GetDepartamentoRepository(filtro filtros.DepartamentoFiltro) (departamentos []entities.Departamentos, erro error)
	GetLocalidadRepository(filtro filtros.LocalidadFiltro) (localidads []entities.Localidads, erro error)
	GetMunicipioRepository(filtro filtros.MunicipioFiltro) (municipios []entities.Municipios, erro error)
	GetBarrioRepository(filtro filtros.BarrioFiltro) (barrios []entities.Barrio, erro error)
	GetPlanillasTrabajosRepository(filtro filtros.PlanillaFiltro) (planillasTrabajos []entities.Planillatrabajo, erro error)
	GetPlanillasRecorridosRepository(filtro filtros.PlanillaRecorridoFiltro) (planillasRecorridos []entities.Planillatrabajosrecorrido, erro error)
	GetRecolectoresBarriosRepository(filtro filtros.PlanillaFiltro) (recolectoresBarrios []entities.Recolectores_barrio, erro error)

	/*
		permite obtener el barrio y el municipio al que pertenece el ciudadano que utiliza la app
		recibe el users_app_id como parametro para obtener la informacion de la base de datos y retorna la entidad appusers
	*/
	GetAppUserByUserIdRepository(userAppId uint) (entityAppUser entities.Appuser, rowAffected bool, erro error)

	// Cambia el estado de un recurrido
	PutEstadoRecorridoRepository(entityRecorrido entities.Planillatrabajosrecorrido) (erro error)

	// Cambia el estado de una planilla
	PutEstadoPlanillaRepository(entityRecorrido entities.Planillatrabajo) (erro error)

	/*
		permite guardar en la BD tabla Appusers en idAppuser e id del barrio
	*/
	PostRegistrarAppUsuarioRepository(entityUsuarioApp entities.Appuser) (erro error)
	PostRecoleccionRepository(e entities.Recoleccione) (id uint, erro error)
	GetRecoleccioneRepository(filtro filtros_recolecciones.GetRecoleccioneFiltro) (entity entities.Recoleccione, rowAffected bool, erro error)

	PutRecoleccionRepository(e entities.Recoleccione) (erro error)
	// guardar en la tabla de imagenes relacion con recolecciones de una imagen subida
	PostImagesRecoleccionRepository(e entities.Recoleccionesimagen) (id uint, erro error)
}

// se define estructrura
type recoleccionRepository struct {
	SqlClient *database.MySQLClient
}

// defino constructor

func NewRecoleccionesRepository(conn *database.MySQLClient) RecoleccionesRepository {
	return &recoleccionRepository{
		SqlClient: conn,
	}
}

func (r *recoleccionRepository) GetConfiguracionesByMunicipioIdRepository(municipioId uint) (ListEntityConfiguraciones []entities.Configuracione, erro error) {

	resp := r.SqlClient.Table("configuraciones")
	resp.Preload("Boton")
	resp.Preload("Boton.Tipo")
	resp.Where("municipios_id = ?", municipioId)
	resp.Find(&ListEntityConfiguraciones)
	if resp.Error != nil {
		erro = errors.New(ERROR_OBTENER_CONFIGURACIONES)
	}
	return
}

func (r *recoleccionRepository) GetProvinciaRepository(filtro filtros.ProvinciaFiltro) (provincias []entities.Provincias, erro error) {
	resp := r.SqlClient.Table("provincias")
	// si GetAll es true, no se tiene en cuenta el filtro de id
	if !filtro.GetAll {
		resp.Where("id", filtro.ProvinciaId)
	}
	resp.Preload("Departamentos.Localidads.Municipios.Barrios")
	resp.Find(&provincias)
	if resp.Error != nil {
		erro = errors.New(ERROR_GET_PROVINCIA)
	}
	return
}

func (r *recoleccionRepository) GetDepartamentoRepository(filtro filtros.DepartamentoFiltro) (departamentos []entities.Departamentos, erro error) {
	resp := r.SqlClient.Table("departamentos")

	if !filtro.GetAll {
		if filtro.DepartamentoId != 0 {
			resp.Where("id", filtro.DepartamentoId)
		}
		if filtro.ProvinciaId != 0 {
			resp.Where("provincias_id", filtro.ProvinciaId)
		}
	}

	resp.Preload("Localidads.Municipios")

	resp.Find(&departamentos)
	if resp.Error != nil {
		erro = errors.New(ERROR_GET_DEPARTAMENTO)
	}
	return
}

func (r *recoleccionRepository) GetLocalidadRepository(filtro filtros.LocalidadFiltro) (localidads []entities.Localidads, erro error) {

	resp := r.SqlClient.Table("localidads")
	if !filtro.GetAll {
		if filtro.LocalidadId != 0 {
			resp.Where("id", filtro.LocalidadId)
		}
	}
	resp.Preload("Municipios")
	resp.Find(&localidads)

	if resp.Error != nil {
		erro = errors.New(ERROR_GET_LOCALIDAD)
	}
	return
}

func (r *recoleccionRepository) GetMunicipioRepository(filtro filtros.MunicipioFiltro) (municipios []entities.Municipios, erro error) {
	resp := r.SqlClient.Table("municipios")
	if !filtro.GetAll {
		if filtro.MunicipioId != 0 {
			resp.Where("id", filtro.MunicipioId)
		}
	}
	resp.Preload("Barrios")
	resp.Find(&municipios)
	if resp.Error != nil {
		erro = errors.New(ERROR_GET_MUNICIPIO)
	}
	return
}

func (r *recoleccionRepository) GetAppUserByUserIdRepository(userAppId uint) (entityAppUser entities.Appuser, rowAffected bool, erro error) {
	resp := r.SqlClient.Table("appusers")
	resp.Where("users_app_id = ?", userAppId)
	resp.Preload("Barrio")
	resp.Preload("Barrio.Municipio")
	resp.Find(&entityAppUser)
	if resp.Error != nil {
		erro = errors.New(ERROR_OBTENER_DATO_USERAPP)
	}
	if resp.RowsAffected == 0 {
		rowAffected = true
		return
	}
	return
}

func (r *recoleccionRepository) PostRegistrarAppUsuarioRepository(entityUsuarioApp entities.Appuser) (erro error) {
	resp := r.SqlClient.Omit(clause.Associations).Create(&entityUsuarioApp)
	if resp.Error != nil {
		erro = errors.New(ERROR_CREAR_APPUSERS)
		return
	}
	return
}

func (r *recoleccionRepository) GetBarrioRepository(filtro filtros.BarrioFiltro) (barrios []entities.Barrio, erro error) {
	resp := r.SqlClient.Table("barrios")

	if filtro.MunicipioId != 0 {
		resp.Where("municipios_id", filtro.MunicipioId)
	}
	resp.Preload("Municipio")
	resp.Find(&barrios)
	if resp.Error != nil {
		erro = errors.New(ERROR_GET_MUNICIPIO)
	}
	return
}

func (r *recoleccionRepository) GetPlanillasTrabajosRepository(filtro filtros.PlanillaFiltro) (planillasTrabajos []entities.Planillatrabajo, erro error) {
	resp := r.SqlClient.Table("planillatrabajos")
	if filtro.Recolectores_barrios_id != 0 {
		resp.Where("recolectores_barrios_id", filtro.Recolectores_barrios_id)
	}
	resp.Find(&planillasTrabajos)
	if resp.Error != nil {
		erro = errors.New(ERROR_GET_PLANILLAS)
	}
	return
}

func (r *recoleccionRepository) GetRecolectoresBarriosRepository(filtro filtros.PlanillaFiltro) (recolectoresBarrios []entities.Recolectores_barrio, erro error) {
	resp := r.SqlClient.Table("recolectores_barrios")
	if filtro.Recolectores_barrios_id != 0 {
		resp.Where("id", filtro.Recolectores_barrios_id)
	}
	if filtro.Recolectores_id != 0 {
		resp.Where("recolectores_id", filtro.Recolectores_id)
	}
	if filtro.Barrios_id != 0 {
		resp.Where("barrios_id", filtro.Barrios_id)
	}
	resp.Find(&recolectoresBarrios)
	if resp.Error != nil {
		erro = errors.New(ERROR_GET_RECOLECTORES_BARRIOS)
	}
	return
}

func (r *recoleccionRepository) GetPlanillasRecorridosRepository(filtro filtros.PlanillaRecorridoFiltro) (planillasRecorridos []entities.Planillatrabajosrecorrido, erro error) {
	resp := r.SqlClient.Table("planillatrabajosrecorridos")
	resp.Preload("Recoleccion")
	if filtro.PlanillaId != 0 {
		resp.Where("planillatrabajos_id", filtro.PlanillaId)
	}
	if filtro.Id != 0 {
		resp.Where("id", filtro.Id)
	}
	resp.Find(&planillasRecorridos)
	if resp.Error != nil {
		erro = errors.New(ERROR_GET_PLANILLAS_RECORRIDOS)
	}
	return
}

func (r *recoleccionRepository) PutEstadoRecorridoRepository(entityRecorrido entities.Planillatrabajosrecorrido) (erro error) {

	if r.SqlClient.Model(&entityRecorrido).Updates(entityRecorrido).RowsAffected == 0 {
		erro = errors.New(ERROR_UPDATE_ESTADO_RECORRIDO)
		return
	}
	return
}

func (r *recoleccionRepository) PutEstadoPlanillaRepository(entityPlanilla entities.Planillatrabajo) (erro error) {

	if r.SqlClient.Model(&entityPlanilla).Updates(entityPlanilla).RowsAffected == 0 {
		erro = errors.New(ERROR_UPDATE_ESTADO_PLANILLA)
		return
	}
	return
}

func (r *recoleccionRepository) PostRecoleccionRepository(e entities.Recoleccione) (id uint, erro error) {
	err := r.SqlClient.Create(&e).Error

	if err != nil {
		if strings.Contains(err.Error(), "1452") {
			erro = errors.New(ERROR_CREAR_RECOLECCION_FK)
			return
		}
		erro = errors.New(ERROR_CREAR_RECOLECCION)
		return
	}
	id = e.ID
	return
}

func (r *recoleccionRepository) GetRecoleccioneRepository(filtro filtros_recolecciones.GetRecoleccioneFiltro) (entity entities.Recoleccione, rowAffected bool, erro error) {

	resp := r.SqlClient.Table("recolecciones")

	resp.Where("id = ?", filtro.Id)

	resp.Find(&entity)

	if resp.Error != nil {
		erro = errors.New(ERROR_GET_RECOLECCIONE)
	}
	if resp.RowsAffected == 0 {
		rowAffected = true
		return
	}
	return
}

func (r *recoleccionRepository) PutRecoleccionRepository(e entities.Recoleccione) (erro error) {

	if r.SqlClient.Model(&e).Updates(e).RowsAffected == 0 {
		erro = errors.New(ERROR_UPDATE_ESTADO_PLANILLA)
		return
	}
	return

}

func (r *recoleccionRepository) PostImagesRecoleccionRepository(e entities.Recoleccionesimagen) (id uint, erro error) {
	// guardar en base de datos
	erro = r.SqlClient.Create(&e).Error

	if erro != nil {
		if strings.Contains(erro.Error(), "1452") {
			erro = errors.New(ERROR_IMAGEN_RECOLECCION_FK)
			return
		}
		erro = errors.New(ERROR_CREAR_IMAGEN)
		return
	}
	id = e.ID
	return

}
