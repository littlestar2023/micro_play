package service

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"micro_play/app/user/repository/db"
	"micro_play/idl/pb"
	"micro_play/pkg/code"
	"micro_play/pkg/model"
	"sync"
)

type UserSrv struct {
}

var UserSrvInstance *UserSrv
var UserSrvOnce sync.Once

func GetUserSrv() *UserSrv {
	UserSrvOnce.Do(func() {
		UserSrvInstance = &UserSrv{}
	})
	return UserSrvInstance
}

func (srv *UserSrv) UserLogin(ctx context.Context, req *pb.UserRequest, resp *pb.UserDetailResponse) (err error) {
	resp.Code = code.SUCCESS

	var user *model.User
	user, err = db.FindUserByUserName(req.UserName)
	if err != nil {
		resp.Code = code.ERROR
		return err
	}

	if !user.CheckPassword(req.Password) {
		resp.Code = code.InvalidParams
		return err
	}

	resp.UserDetail = BuildUser(user)
	return err
}

func (srv *UserSrv) UserRegister(ctx context.Context, req *pb.UserRequest, resp *pb.UserDetailResponse) (err error) {
	if req.Password != req.PasswordConfirm {
		err = errors.New("两次密码输入不一致")
		return
	}

	resp.Code = code.SUCCESS
	_, err = db.FindUserByUserName(req.UserName)
	if err != nil && err != gorm.ErrRecordNotFound {
		resp.Code = code.ERROR
		resp.Message = err.Error()
		return
	}

	user := &model.User{
		UserName: req.UserName,
	}
	if err = user.SetPassword(req.Password); err != nil {
		resp.Code = code.ERROR
		resp.Message = err.Error()
		return
	}
	if err = db.CreateUser(user); err != nil {
		resp.Code = code.ERROR
		resp.Message = err.Error()
		return
	}

	resp.UserDetail = BuildUser(user)
	return
}

func BuildUser(item *model.User) *pb.UserModel {
	userModel := pb.UserModel{
		Id:        uint32(item.ID),
		UserName:  item.UserName,
		CreatedAt: item.CreatedAt.Unix(),
		UpdatedAt: item.UpdatedAt.Unix(),
	}
	return &userModel
}
