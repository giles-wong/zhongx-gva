package business

import (
	"encoding/json"
	"errors"
	"github.com/giles-wong/zhongx-gva/server/global"
	"github.com/giles-wong/zhongx-gva/server/model/business"
	businessReq "github.com/giles-wong/zhongx-gva/server/model/business/request"
	businessRes "github.com/giles-wong/zhongx-gva/server/model/business/response"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type ReadingService struct {
}

func (readingService *ReadingService) GetReadingList(info businessReq.GetReadingList) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&business.CertReading{})
	var readingList []business.CertReading
	var readingResList []businessRes.ReadingResponse

	//db = db.Where("is_deleted = ?", 0)
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.Mobile != "" {
		db = db.Where("mobile = ?", info.Mobile)
	}
	if info.Number != "" {
		db = db.Where("number = ?", info.Number)
	}
	if info.IdCard != "" {
		db = db.Where("id_card = ", info.IdCard)
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&readingList).Error

	for _, item := range readingList {
		//把json 转为切片
		var fileList []businessReq.CertFile
		_ = json.Unmarshal([]byte(item.File), &fileList)
		readingResList = append(readingResList, businessRes.ReadingResponse{
			ID:       item.ID,
			Name:     item.Name,
			Mobile:   item.Mobile,
			IdCard:   item.IdCard,
			Number:   item.Number,
			Group:    item.Group,
			File:     fileList,
			Referrer: item.Referrer,
			Profile:  item.Profile,
			Remark:   item.Remark,
			Status:   item.Status,
		})

	}

	return readingResList, total, err
}

func (readingService *ReadingService) AddReading(cr business.CertReading) (reading business.CertReading, err error) {
	var crd business.CertReading
	if !errors.Is(global.GVA_DB.Where("number = ?", cr.Number).First(&crd).Error, gorm.ErrRecordNotFound) { // 判断证书是否存在
		return reading, errors.New("证书已经存在")
	}

	err = global.GVA_DB.Create(&cr).Error
	return cr, err
}

func (readingService *ReadingService) EditReading(req business.CertReading) error {
	return global.GVA_DB.Model(&business.CertReading{}).
		Select("updated_at", "name", "number", "mobile", "id_card", "file", "group", "referrer", "profile", "status", "remark").
		Where("id=?", req.ID).
		Updates(map[string]interface{}{
			"updated_at": time.Now(),
			"mobile":     req.Mobile,
			"id_card":    req.IdCard,
			"group":      req.Group,
			"referrer":   req.Referrer,
			"profile":    req.Profile,
			"status":     req.Status,
			"remark":     req.Remark,
		}).Error
}

func (readingService *ReadingService) DeleteReading(id uint) (err error) {
	global.GVA_LOG.Info("删除证书信息", zap.String("id", strconv.Itoa(int(id))))

	return global.GVA_DB.Model(&business.CertReading{}).
		Where("id=?", id).
		Updates(map[string]interface{}{
			"updated_at": time.Now(),
			"deleted_at": time.Now(),
		}).Error
}
