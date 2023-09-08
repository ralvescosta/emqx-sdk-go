package authentication

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"os"

	"github.com/ralvescosta/emqx-sdk-go/client"
	"go.uber.org/zap"
)

func (c *authenticationClient) PostGlobalImportUsersIntoAuthenticator(ctx context.Context, authenticatorID string, file *os.File) error {
	urlPath := fmt.Sprintf("api/v5/authentication/%s/import_users", authenticatorID)

	byt := []byte{}
	reqBody := bytes.NewBuffer(byt)

	multipartWriter := multipart.NewWriter(reqBody)
	defer multipartWriter.Close()

	formFile, err := multipartWriter.CreateFormFile("filename", file.Name())
	if err != nil {
		c.logger.Error("error to create form file", zap.Error(err))

		return err
	}

	_, err = io.Copy(formFile, file)
	if err != nil {
		c.logger.Error("error moving file", zap.Error(err))

		return err
	}

	_, err = c.apiClient.Perform(
		ctx,
		client.NewParams().Post(urlPath).RequestBody(reqBody).RequestContentType(multipartWriter.FormDataContentType()),
	)
	if err != nil {
		return err
	}

	return nil
}
