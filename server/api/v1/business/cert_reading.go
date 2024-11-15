package business

import (
	"encoding/json"
	"github.com/giles-wong/zhongx-gva/server/global"
	"github.com/giles-wong/zhongx-gva/server/model/business"
	businessReq "github.com/giles-wong/zhongx-gva/server/model/business/request"
	"github.com/giles-wong/zhongx-gva/server/model/common/response"
	"github.com/giles-wong/zhongx-gva/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"log"
)

type CertReadingApi struct {
}

// GetReadingList
// @Tags      Gert
// @Summary   分页获取证书列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      systemReq.GetUserList                                        true  "页码, 每页大小"
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "分页获取用户列表,返回包括列表,总数,页码,每页数量"
// @Router    /cert/getList [post]
func (cr *CertReadingApi) GetReadingList(ctx *gin.Context) {
	var pageInfo businessReq.GetReadingList
	err := ctx.ShouldBindJSON(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	err = utils.Verify(pageInfo, utils.PageInfoVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	list, total, err := readingService.GetReadingList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", ctx)
		return
	}

	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", ctx)
}

func (cr *CertReadingApi) AddReading(c *gin.Context) {
	var r businessReq.AddReading
	err := c.ShouldBindJSON(&r)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 将文件数组转换为 JSON
	fileJSON, err := json.Marshal(r.File)
	if err != nil {
		log.Fatalf("Error marshaling file data: %v", err)
	}

	reading := &business.CertReading{
		Number:   r.Number,
		Name:     r.Name,
		IdCard:   r.IdCard,
		Mobile:   r.Mobile,
		Profile:  r.Profile,
		Referrer: r.Referrer,
		Remark:   r.Remark,
		Status:   r.Status,
		File:     string(fileJSON),
		Group:    r.Group,
	}
	_, err = readingService.AddReading(*reading)
	if err != nil {
		global.GVA_LOG.Error("添加失败!", zap.Error(err))
		response.FailWithDetailed(r, "添加失败", c)
		return
	}
	response.OkWithDetailed(r, "添加成功", c)
}

func (cr *CertReadingApi) EditReading(c *gin.Context) {
	var r businessReq.EditReading
	err := c.ShouldBindJSON(&r)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 将文件数组转换为 JSON
	fileJSON, err := json.Marshal(r.File)
	if err != nil {
		log.Fatalf("Error marshaling file data: %v", err)
	}
	reading := &business.CertReading{
		ID:       r.ID,
		Number:   r.Number,
		Name:     r.Name,
		IdCard:   r.IdCard,
		Mobile:   r.Mobile,
		Profile:  r.Profile,
		Referrer: r.Referrer,
		Remark:   r.Remark,
		Status:   r.Status,
		File:     string(fileJSON),
		Group:    r.Group,
	}

	err = readingService.EditReading(*reading)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

func (cr *CertReadingApi) DeleteReading(c *gin.Context) {
	var r businessReq.ReadingId
	err := c.ShouldBindJSON(&r)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = readingService.DeleteReading(r.ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}
