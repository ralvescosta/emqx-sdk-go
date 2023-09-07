package authentication

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"os"

	"go.uber.org/zap"
)

func (c *Client) PostGlobalImportUsersIntoAuthenticator(ctx context.Context, authenticatorID string, file *os.File) error {
	urlPath := fmt.Sprintf("api/v5/authentication/%s/import_users", authenticatorID)

	byt := []byte{}
	reqBody := bytes.NewBuffer(byt)

	multipartWriter := multipart.NewWriter(reqBody)
	defer multipartWriter.Close()

	formFile, err := multipartWriter.CreateFormFile("filename", file.Name())
	if err != nil {
		c.options.Logger.Error("error to create form file", zap.Error(err))
		return err
	}

	_, err = io.Copy(formFile, file)
	if err != nil {
		c.options.Logger.Error("error moving file", zap.Error(err))
		return err
	}

	if err := c.perform(
		ctx,
		NewRequestParams().Post(urlPath).ReqBody(reqBody).ReqContentType(multipartWriter.FormDataContentType()),
	); err != nil {
		return err
	}

	return nil
}
