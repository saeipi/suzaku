package ctrl

import (
	"github.com/gin-gonic/gin"
	"suzaku/micro/etcd/common"
	"suzaku/micro/etcd/etcd_master/job"
	"suzaku/pkg/network"
)

type jobCtrl struct {
}

func NewJobCtrl() *jobCtrl {
	return &jobCtrl{}
}

func (ctrl *jobCtrl) Save(ctx *gin.Context) {
	var req *common.Job
	if err := ctx.ShouldBind(&req); err != nil {
		network.Error(ctx, err, network.ErrorCodeDeserializeErr)
		return
	}
	var resp network.Resp
	resp.Data, resp.Err = job.SG_JOBMGR.SaveJob(req)
	if resp.Err != nil {
		network.Error(ctx, resp.Err, resp.Code)
		return
	}
	network.Success(ctx, resp.Data)
}

func (ctrl *jobCtrl) Delete(ctx *gin.Context) {
	name := ctx.PostForm("name")
	var resp network.Resp
	resp.Data, resp.Err = job.SG_JOBMGR.DeleteJob(name)
	if resp.Err != nil {
		network.Error(ctx, resp.Err, resp.Code)
		return
	}
	network.Success(ctx, resp.Data)
}
