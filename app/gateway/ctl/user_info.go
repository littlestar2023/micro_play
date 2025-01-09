package ctl

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/littlestar2023/common_pkg/global"
	"github.com/redis/go-redis/v9"
	"time"
)

var userKey = "user_key"

type UserInfo struct {
	Id uint `json:"id"`
}

func InitUserInfo(c *gin.Context) error {
	cachedUser, err := GetUserInfo(c.Request.Context(), c.GetUint(userKey))
	if err != nil {
		return err
	}

	// 如果缓存中存在用户信息，直接返回
	if cachedUser != nil {
		return nil
	}

	// 如果缓存中不存在用户信息，将用户信息存储到缓存中
	if err = SetUserInCache(c.Request.Context(), &UserInfo{Id: c.GetUint(userKey)}, global.CMP_REDIS); err != nil {
		return err
	}

	return nil
}

func GetUserInfo(ctx context.Context, userId uint) (*UserInfo, error) {

	data, err := global.CMP_REDIS.Get(ctx, GetUserKey(userId)).Bytes()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	var user UserInfo
	if err = json.Unmarshal(data, &user); err != nil {
		// 解码数据时出错
		return nil, err
	}
	return &user, nil
}

func SetUserInCache(ctx context.Context, user *UserInfo, redisClient *redis.Client) error {
	data, err := json.Marshal(user)
	if err != nil {
		return err
	}
	if err = redisClient.Set(ctx, GetUserKey(user.Id), data, time.Second*30).Err(); err != nil {
		return err
	}
	return nil
}

func GetUserKey(userId uint) string {

	return fmt.Sprintf("%v_%v", userKey, userId)
}
