package service

import (
	"context"
	"github.com/littlestar2023/common_pkg/global"
	"go.uber.org/zap"
	"micro_play/app/task/repository/db"
	"micro_play/idl/pb"
	"micro_play/pkg/code"
	"micro_play/pkg/model"
	"sync"
)

var TaskSrvIns *TaskSrv
var TaskSrvOnce sync.Once

type TaskSrv struct {
}

func GetTaskSrv() *TaskSrv {
	TaskSrvOnce.Do(func() {
		TaskSrvIns = &TaskSrv{}
	})
	return TaskSrvIns
}

// CreateTask 创建备忘录，将备忘录信息生产，放到rabbitMQ消息队列中
func (t *TaskSrv) CreateTask(ctx context.Context, req *pb.TaskRequest, resp *pb.TaskDetailResponse) (err error) {

	resp.Code = code.SUCCESS

	task := &model.Task{
		Uid:       uint(req.Uid),
		Title:     req.Title,
		Status:    int(req.Status),
		Content:   req.Content,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
	}
	err = db.CreateTask(task)
	if err != nil {
		resp.Code = code.ERROR
		resp.Message = err.Error()
		return
	}
	return
}

// GetTasksList 实现备忘录服务接口 获取备忘录列表
func (t *TaskSrv) GetTasksList(ctx context.Context, req *pb.TaskRequest, resp *pb.TaskListResponse) (err error) {
	resp.Code = code.SUCCESS
	if req.Limit == 0 {
		req.Limit = 10
	}
	// 查找备忘录
	r, count, err := db.ListTaskByUserId(req.Uid, int(req.Start), int(req.Limit))
	if err != nil {
		resp.Code = code.ERROR
		resp.Message = err.Error()
		global.CMP_LOG.Error("ListTaskByUserId error", zap.Error(err))
		return
	}
	// 返回proto里面定义的类型
	var taskRes []*pb.TaskModel
	for _, item := range r {
		taskRes = append(taskRes, BuildTask(item))
	}
	resp.TaskList = taskRes
	resp.Count = uint32(count)
	return
}

// GetTask 获取详细的备忘录
func (t *TaskSrv) GetTask(ctx context.Context, req *pb.TaskRequest, resp *pb.TaskDetailResponse) (err error) {
	resp.Code = code.SUCCESS
	r, err := db.GetTaskByTaskIdAndUserId(req.Id, req.Uid)
	if err != nil {
		resp.Code = code.ERROR
		resp.Message = err.Error()
		global.CMP_LOG.Error("GetTask error", zap.Error(err))
		return
	}

	taskRes := BuildTask(r)
	resp.TaskDetail = taskRes
	return
}

// UpdateTask 修改备忘录
func (t *TaskSrv) UpdateTask(ctx context.Context, req *pb.TaskRequest, resp *pb.TaskDetailResponse) (err error) {
	// 查找该用户的这条信息
	resp.Code = code.SUCCESS
	taskData, err := db.UpdateTask(req)
	if err != nil {
		resp.Code = code.ERROR
		resp.Message = err.Error()
		global.CMP_LOG.Error("UpdateTask error", zap.Error(err))
		return
	}
	resp.TaskDetail = BuildTask(taskData)
	return
}

// DeleteTask 删除备忘录
func (t *TaskSrv) DeleteTask(ctx context.Context, req *pb.TaskRequest, resp *pb.TaskDetailResponse) (err error) {
	resp.Code = code.SUCCESS
	err = db.DeleteTaskByIdAndUserId(req.Id, req.Uid)
	if err != nil {
		resp.Code = code.ERROR
		resp.Message = err.Error()
		global.CMP_LOG.Error("DeleteTask error", zap.Error(err))
		return
	}
	return
}

func BuildTask(item *model.Task) *pb.TaskModel {
	taskModel := pb.TaskModel{
		Id:         uint64(item.ID),
		Uid:        uint64(item.Uid),
		Title:      item.Title,
		Content:    item.Content,
		StartTime:  item.StartTime,
		EndTime:    item.EndTime,
		Status:     int64(item.Status),
		CreateTime: item.CreatedAt.Unix(),
		UpdateTime: item.UpdatedAt.Unix(),
	}
	return &taskModel
}
