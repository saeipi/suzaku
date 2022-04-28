package api_minio

import (
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"mime/multipart"
	"suzaku/internal/interface/dto/dto_api"
	"suzaku/pkg/common/config"
	"suzaku/pkg/common/log"
	"suzaku/pkg/common/minioc"
	"suzaku/pkg/constant"
	"suzaku/pkg/http"
	"suzaku/pkg/utils"
)

func UploadFile(c *gin.Context) {
	var (
		req         dto_api.MinioUploadFileReq
		resp        dto_api.MinioUploadFileResp
		fileHeader  *multipart.FileHeader
		file        multipart.File
		objectName  string
		contentType string
		err         error
	)

	if err = c.Bind(&req); err != nil {
		log.NewError("0", utils.GetSelfFuncName(), "BindJSON failed ", err.Error())
		http.Error(c, err, http.ErrorCodeHttpReqDeserializeFailed)
		return
	}
	switch req.FileType {
	// imageType upload snapShot
	case constant.ImageType:
		fileHeader, err = c.FormFile("image")
		if err != nil {
			log.NewError(req.OperationID, utils.GetSelfFuncName(), "missing file arg: ", err.Error())
			http.Error(c, err, http.ErrorCodeHttp400)
			return
		}
		file, err = fileHeader.Open()
		if err != nil {
			log.NewError(req.OperationID, utils.GetSelfFuncName(), "Open file error", err.Error())
			http.Error(c, err, http.ErrorCodeHttp400)
			return
		}
		objectName, contentType = utils.GetNewFileNameAndContentType(fileHeader.Filename, constant.ImageType)
		log.Debug(req.OperationID, utils.GetSelfFuncName(), objectName, contentType)
		_, err = minioc.PutObject(objectName, file, fileHeader.Size, minio.PutObjectOptions{ContentType: contentType})
		if err != nil {
			log.NewError(req.OperationID, utils.GetSelfFuncName(), "PutObject error", err.Error())
			http.Error(c, err, http.ErrorCodeHttp400)
			return
		}
		resp.Url = config.Config.Credential.Minio.Endpoint + "/" + config.Config.Credential.Minio.Bucket + "/" + objectName
		resp.NewName = objectName
		resp.Name = fileHeader.Filename
	}
	http.Success(c, resp)
	return
}
