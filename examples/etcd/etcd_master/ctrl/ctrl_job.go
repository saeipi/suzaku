package ctrl

import (
	"github.com/gin-gonic/gin"
	"suzaku/examples/etcd/common"
	"suzaku/examples/etcd/etcd_master/job"
	"suzaku/pkg/http"
)

type jobCtrl struct {
}

func NewJobCtrl() *jobCtrl {
	return &jobCtrl{}
}

func (ctrl *jobCtrl) Save(ctx *gin.Context) {
	var req *common.Job
	if err := ctx.ShouldBind(&req); err != nil {
		http.Error(ctx, err, http.ErrorCodeHttpReqDeserializeFailed)
		return
	}
	var resp http.Resp
	resp.Data, resp.Err = job.SG_JOBMGR.SaveJob(req)
	if resp.Err != nil {
		http.Error(ctx, resp.Err, resp.Code)
		return
	}
	http.Success(ctx, resp.Data)
}

func (ctrl *jobCtrl) Delete(ctx *gin.Context) {
	name := ctx.PostForm("name")
	var resp http.Resp
	resp.Data, resp.Err = job.SG_JOBMGR.DeleteJob(name)
	if resp.Err != nil {
		http.Error(ctx, resp.Err, resp.Code)
		return
	}
	http.Success(ctx, resp.Data)
}
