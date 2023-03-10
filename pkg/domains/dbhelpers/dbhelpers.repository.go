package dbhelpers

import (
	"errors"

	"github.com/Corrientes-Telecomunicaciones/api_go_recolector/internal/database"
)

type DbHelpersRepository interface {
	CheckIfModelsExists(tableName string, id uint) (exists bool, erro error)
}

type dbHelpersRepository struct {
	SqlClient *database.MySQLClient
}

func NewDbHelpersRepository(conn *database.MySQLClient) DbHelpersRepository {
	return &dbHelpersRepository{
		SqlClient: conn,
	}
}

// retorna si existe o no un registro segun un id y un nombre de tabla
func (r *dbHelpersRepository) CheckIfModelsExists(tableName string, id uint) (exists bool, erro error) {

	erro = r.SqlClient.Table(tableName).Select("count(*) > 0").Where("id = ?", id).Find(&exists).Error

	if erro != nil {
		erro = errors.New(ERROR_DB_ACCESS)
	}
	return
}
