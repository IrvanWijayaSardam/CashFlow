package dto

import "mime/multipart"

type UploadFileProfile struct {
	Profile *multipart.FileHeader `form:"profile"`
}
