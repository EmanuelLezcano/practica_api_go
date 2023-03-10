package storage

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/Corrientes-Telecomunicaciones/api_go_recolector/internal/config"
	"github.com/Corrientes-Telecomunicaciones/api_go_recolector/internal/logs"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type Minio struct {
	S3 *s3.S3
}

func NewMinioSession() Minio {
	minioSession := session.Must(session.NewSession(&aws.Config{
		Credentials:      credentials.NewStaticCredentials(config.AWS_KEY, config.AWS_SECRET, ""),
		Region:           aws.String(config.AWS_REGION),
		Endpoint:         aws.String(config.MINIO_ENDPOINT),
		DisableSSL:       aws.Bool(true),
		S3ForcePathStyle: aws.Bool(true),
	}))
	return Minio{
		S3: s3.New(minioSession),
	}
}

func (m Minio) GetObject(ctx context.Context, key string) (string, error) {
	inputS3 := &s3.ListObjectsInput{
		Bucket:  aws.String(config.AWS_BUCKET),
		MaxKeys: aws.Int64(1000),
		Prefix:  aws.String(key),
	}
	result1, err := m.S3.ListObjects(inputS3)
	if err != nil {
		return "", fmt.Errorf("s3.GetObjectWithContext: %w", err)
	}

	directorio := time.Now().Local().Format("02-01-2006") + config.DIR_TEMP_NAME
	direcName, err := ioutil.TempDir(".."+config.DIR_IMAGES_RECOLECCION, directorio)
	if err != nil {
		return "", fmt.Errorf("error creando directorio: %w", err)
	}
	for _, v := range result1.Contents {
		stringArray := strings.Split(*aws.String(*v.Key), "/")
		if stringArray[2] != "" {
			obj, err := m.S3.GetObject(&s3.GetObjectInput{
				Bucket: aws.String(config.AWS_BUCKET),
				Key:    aws.String(*v.Key),
			})
			if err != nil {
				return "", fmt.Errorf("error al obtener los archivos: %w", err)
			}
			stringArray := strings.Split(*aws.String(*v.Key), "/")
			file, err := os.Create(direcName + "/" + stringArray[2])
			if err != nil {
				return "", fmt.Errorf("error al crear archivo temporal: %w", err)
			}
			io.Copy(file, obj.Body)
		}
	}
	return direcName, nil
}

func (m Minio) PutObject(ctx context.Context, data []byte, fileName, fileType string) (bool, error) {
	var res bool
	result := &s3.PutObjectInput{
		Body:   bytes.NewReader(data),
		Bucket: aws.String(config.AWS_BUCKET),
		Key:    aws.String(fileName),
	}
	logs.Info(result)
	_, err := m.S3.PutObjectWithContext(ctx, result)
	if err != nil {
		return res, fmt.Errorf("s3.PutObjectWithContext: %w, bucket: %s", err, config.AWS_BUCKET)
	}
	res = true
	return res, nil
}

func (m Minio) DeleteObject(ctx context.Context, key string) error {
	_, err := m.S3.DeleteObjectWithContext(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(config.AWS_BUCKET),
		Key:    aws.String(key),
	})
	if err != nil {
		return fmt.Errorf("error al intentar borrar los archivos del minio: %w", err)
	}

	return nil
}
