package rpc

import (
	"go-micro.dev/v4"
	"micro_play/app/gateway/wrappers"
	"micro_play/idl/pb"
)

var (
	UserService pb.UserService
	TaskService pb.TaskService
)

func InitRPC() {
	// 用户服务调用实例
	userMicroService := micro.NewService(
		micro.Name("userService.client"),
		micro.WrapClient(wrappers.NewUserWrapper),
	)
	userService := pb.NewUserService("micro-user-service", userMicroService.Client())
	UserService = userService

	// task
	taskMicroService := micro.NewService(
		micro.Name("taskService.client"),
		micro.WrapClient(wrappers.NewTaskWrapper),
	)
	taskService := pb.NewTaskService("micro-task-service", taskMicroService.Client())
	TaskService = taskService
}
