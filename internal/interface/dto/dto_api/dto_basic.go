package dto_api

type PageReq struct {
	Page     int `form:"page" json:"page" example:"1"  binding:"required,min=1"`
	PageSize int `form:"page_size" json:"page_size" example:"10" binding:"required,min=1,max=100"`
}
