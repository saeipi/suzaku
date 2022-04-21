package dto_api

type MinioStorageCredentialReq struct {
	OperationID string `json:"operation_id"`
}

type MiniostorageCredentialResp struct {
	AccessKey      string `json:"access_key"`
	SecretKey      string `json:"secret_key"`
	SessionToken   string `json:"session_token"`
	BucketName     string `json:"bucket_name"`
	StsEndpointUrl string `json:"sts_endpoint_url"`
}

type MinioUploadFileReq struct {
	OperationID string `form:"operation_id" binding:"required"`
	FileType    int    `form:"file_type" binding:"required"`
}

type MinioUploadFileResp struct {
	Url     string `json:"url"`
	Name    string `json:"name"`
	NewName string `json:"new_name"`
}
