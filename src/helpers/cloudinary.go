package helpers

import (
	"context"
	"mime/multipart"
	"os"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
)

func UploadImages(ext string, file multipart.File, handle *multipart.FileHeader) (*uploader.UploadResult, error) {
	name := os.Getenv("CLOUD_NAME")
	key := os.Getenv("CLOUD_KEY")
	secret := os.Getenv("CLOUD_SECRET")

	cld, err := cloudinary.NewFromParams(name, key, secret)
	if err != nil {
		return nil, err
	}

	rand := GenToken(8)

	var ctx = context.Background()
	result, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{
		PublicID: "lectronic/" + ext + "/" + rand + "-" + handle.Filename,
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}
