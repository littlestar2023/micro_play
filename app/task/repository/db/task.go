package db

import (
	"github.com/littlestar2023/common_pkg/global"
	"micro_play/idl/pb"
	"micro_play/pkg/model"
)

// ListTaskByUserId 获取user id
func ListTaskByUserId(userId uint64, start, limit int) (r []*model.Task, count int64, err error) {
	err = global.CMP_DB.Model(&model.Task{}).Offset(start).
		Limit(limit).Where("uid = ?", userId).
		Find(&r).Error

	err = global.CMP_DB.Model(&model.Task{}).Where("uid = ?", userId).
		Count(&count).Error

	return
}

// GetTaskByTaskIdAndUserId 通过 task id 和 user id 获取task
func GetTaskByTaskIdAndUserId(taskId, userId uint64) (r *model.Task, err error) {
	err = global.CMP_DB.Model(&model.Task{}).
		Where("id = ? AND uid = ?", taskId, userId).
		First(&r).Error

	return
}

// UpdateTask 更新task
func UpdateTask(req *pb.TaskRequest) (r *model.Task, err error) {
	r = new(model.Task)
	err = global.CMP_DB.Model(&model.Task{}).
		Where("id = ? AND uid = ?", req.Id, req.Uid).
		First(&r).Error
	if err != nil {
		return
	}
	r.Title = req.Title
	r.Status = int(req.Status)
	r.Content = req.Content

	err = global.CMP_DB.Save(&r).Error
	return
}

func DeleteTaskByIdAndUserId(id, uId uint64) error {
	return global.CMP_DB.Model(&model.Task{}).
		Where("id =? AND uid=?", id, uId).
		Delete(&model.Task{}).Error

}

func CreateTask(in *model.Task) error {
	return global.CMP_DB.Model(&model.Task{}).Create(&in).Error

}
