package storage

import "context"

type Storage interface {
	// PutObject almacena un archivo en el storage
	PutObject(ctx context.Context, data []byte, filename, fileType string) (bool, error)

	// DeleteObject mediante el nombre de acceso de un archivo se lo elimina del storage.
	DeleteObject(ctx context.Context, key string) error

	// GetObject mediante el nombre de acceso devuelve del storage un archivo.s
	GetObject(ctx context.Context, key string) (string, error)
}
