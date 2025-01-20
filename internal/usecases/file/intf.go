package file

import (
	"context"
	"mime/multipart"
)

type UseCase interface {
	UploadFile(context.Context, multipart.File, string) (string, error)
}
