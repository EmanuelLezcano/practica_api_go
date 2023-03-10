package storage

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/Corrientes-Telecomunicaciones/api_go_recolector/internal/logs"
)

// recibe la carpeta remota, la carpeta de origen temporal donde se aloja el archivo, y el nombre del archivo
func LeerDatosArchivo(remoteFolder string, sourcePath string, fileNameAndType string) (data []byte, archivonombre string, archivotipo string, erro error) {
	// leer el archivo desde el source y guardar su contenido en byte
	data, erro = ioutil.ReadFile(fmt.Sprintf("%s/%s", sourcePath, fileNameAndType))

	if erro != nil {
		msj := "error a leer datos del archivo:" + fileNameAndType
		logs.Error(msj)
		erro = errors.New(msj)
		return
	}

	fileNameAndExtension := strings.Split(fileNameAndType, ".")
	archivonombre = fmt.Sprintf("%s/%s.%s", remoteFolder, fileNameAndExtension[0], fileNameAndExtension[1])
	archivotipo = fileNameAndExtension[len(fileNameAndExtension)-1]
	return
}
