package http

import (
	"fmt"
	"github.com/littlestar2023/common_pkg/global"
	"micro_play/idl/pb"
	"micro_play/pkg/model"
	"micro_play/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"

	"micro_play/app/gateway/ctl"
	"micro_play/app/gateway/rpc"
	//log "github.com/CocaineCong/micro-todoList/pkg/logger"
	//"github.com/CocaineCong/micro-todoList/pkg/utils"
	//"github.com/CocaineCong/micro-todoList/types"
)

// UserRegisterHandler 用户注册
func UserRegisterHandler(ctx *gin.Context) {
	var req pb.UserRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ctl.RespError(ctx, err, "UserRegister Bind 绑定参数失败"))
		return
	}
	userResp, err := rpc.UserRegister(ctx, &req)
	if err != nil {
		global.CMP_LOG.Error(fmt.Sprintf("UserRegister:%v", err))
		ctx.JSON(http.StatusInternalServerError, ctl.RespError(ctx, err, "UserRegister RPC 调用失败"))
		return
	}
	ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, userResp))
}

// UserLoginHandler 用户登录
func UserLoginHandler(ctx *gin.Context) {
	var req pb.UserRequest
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ctl.RespError(ctx, err, "UserLogin Bind 绑定参数失败"))
		return
	}
	userResp, err := rpc.UserLogin(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ctl.RespError(ctx, err, "UserLogin RPC 调用失败"))
		return
	}
	token, err := utils.GenerateToken(uint(userResp.UserDetail.Id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ctl.RespError(ctx, err, "GenerateToken 失败"))
		return
	}
	res := &model.TokenData{
		User:  userResp,
		Token: token,
	}
	ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, res))
}
