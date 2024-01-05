package http

import (
	"context"
	"github.com/cossim/coss-server/pkg/http"
	pkghttp "github.com/cossim/coss-server/pkg/http"
	"github.com/cossim/coss-server/pkg/http/response"
	"github.com/cossim/coss-server/pkg/utils/usersorter"
	relationApi "github.com/cossim/coss-server/services/relation/api/v1"
	userApi "github.com/cossim/coss-server/services/user/api/v1"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type blackListRequest struct {
	UserID string `json:"user_id" binding:"required"`
}

// @Summary 黑名单
// @Description 黑名单
// @Accept  json
// @Produce  json
// @param request body blackListRequest true "request"
// @Success		200 {object} utils.Response{}
// @Router /relation/blacklist [get]
func blackList(c *gin.Context) {
	userID := c.Query("user_id")
	if userID == "" {
		logger.Error("用户id为空")
		response.Fail(c, "用户id为空", nil)
		return
	}

	uid, err := http.ParseTokenReUid(c)
	if err != nil {
		return
	}

	// 检查用户是否在查询自己的黑名单列表
	if userID != uid {
		logger.Error("用户权限不足：不允许查询其他用户的黑名单列表")
		response.Fail(c, "用户权限不足：不允许查询其他用户的黑名单列表", nil)
		return
	}

	// 检查用户是否存在
	user, err := userClient.UserInfo(context.Background(), &userApi.UserInfoRequest{UserId: userID})
	if err != nil {
		c.Error(err)
		return
	}

	if user == nil {
		response.Fail(c, "用户不存在", nil)
		return
	}

	// 获取黑名单列表
	blacklistResp, err := relationClient.GetBlacklist(context.Background(), &relationApi.GetBlacklistRequest{UserId: userID})
	if err != nil {
		c.Error(err)
		return
	}

	var users []string
	for _, user := range blacklistResp.Blacklist {
		users = append(users, user.UserId)
	}

	blacklist, err := userClient.GetBatchUserInfo(context.Background(), &userApi.GetBatchUserInfoRequest{UserIds: users})
	if err != nil {
		return
	}

	response.Success(c, "获取黑名单列表成功", gin.H{"blacklist": blacklist})
}

type friendListRequest struct {
	UserID string `json:"user_id" binding:"required"`
}

// @Summary 好友列表
// @Description 好友列表
// @Accept  json
// @Produce  json
// @param request body friendListRequest true "request"
// @Success		200 {object} utils.Response{}
// @Router /relation/friend_list [get]
func friendList(c *gin.Context) {
	userID := c.Query("user_id")
	if userID == "" {
		logger.Error("用户id为空")
		response.Fail(c, "用户id为空", nil)
		return
	}

	uid, err := http.ParseTokenReUid(c)
	if err != nil {
		return
	}

	// 检查用户是否在查询自己的黑名单列表
	if userID != uid {
		logger.Error("用户权限不足：不允许查询其他用户的黑名单列表")
		response.Fail(c, "用户权限不足：不允许查询其他用户的黑名单列表", nil)
		return
	}

	// 检查用户是否存在
	user, err := userClient.UserInfo(context.Background(), &userApi.UserInfoRequest{UserId: userID})
	if err != nil {
		logger.Error("user service UserInfo", zap.Error(err))
		c.Error(err)
		return
	}

	if user == nil {
		response.Fail(c, "用户不存在", nil)
		return
	}

	// 获取好友列表
	friendListResp, err := relationClient.GetFriendList(context.Background(), &relationApi.GetFriendListRequest{UserId: userID})
	if err != nil {
		logger.Error("user service GetFriendList", zap.Error(err))
		c.Error(err)
		return
	}

	var users []string
	for _, user := range friendListResp.FriendList {
		users = append(users, user.UserId)
	}

	userInfos, err := userClient.GetBatchUserInfo(context.Background(), &userApi.GetBatchUserInfoRequest{UserIds: users})
	if err != nil {
		logger.Error("user service GetBatchUserInfo", zap.Error(err))
		c.Error(err)
		return
	}

	var data []usersorter.User
	for _, v := range userInfos.Users {
		data = append(data, usersorter.CustomUserData{
			UserID:   v.UserId,
			NickName: v.NickName,
			Email:    v.Email,
			Tel:      v.Tel,
			Avatar:   v.Avatar,
			Status:   uint(v.Status),
		})
	}

	// Sort and group by specified field
	groupedUsers := usersorter.SortAndGroupUsers(data, "NickName")

	response.Success(c, "获取好友列表成功", usersorter.ConvertToGinH(groupedUsers))
}

type deleteBlacklistRequest struct {
	UserID   string `json:"user_id" binding:"required"`
	FriendID string `json:"friend_id" binding:"required"`
}

// @Summary 删除黑名单
// @Description 删除黑名单
// @Accept  json
// @Produce  json
// @param request body deleteBlacklistRequest true "request"
// @Success		200 {object} utils.Response{}
// @Router /relation/delete_blacklist [post]
func deleteBlacklist(c *gin.Context) {
	req := new(deleteBlacklistRequest)
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("参数验证失败", zap.Error(err))
		response.Fail(c, "参数验证失败", nil)
		return
	}

	// 检查用户是否存在
	user, err := userClient.UserInfo(context.Background(), &userApi.UserInfoRequest{UserId: req.UserID})
	if err != nil {
		c.Error(err)
		return
	}

	if user == nil {
		response.Fail(c, "用户不存在", nil)
		return
	}

	// 进行删除黑名单操作
	if _, err = relationClient.DeleteBlacklist(context.Background(), &relationApi.DeleteBlacklistRequest{UserId: req.UserID, FriendId: req.FriendID}); err != nil {
		c.Error(err)
		return
	}

	response.Success(c, "删除黑名单成功", gin.H{})
}

type addBlacklistRequest struct {
	UserID   string `json:"user_id" binding:"required"`
	FriendID string `json:"friend_id" binding:"required"`
}

// @Summary 添加黑名单
// @Description 添加黑名单
// @Accept  json
// @Produce  json
// @param request body addBlacklistRequest true "request"
// @Success		200 {object} utils.Response{}
// @Router /relation/add_blacklist [post]
func addBlacklist(c *gin.Context) {
	req := new(addBlacklistRequest)
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("参数验证失败", zap.Error(err))
		response.Fail(c, "参数验证失败", nil)
		return
	}

	// 检查用户是否存在
	user, err := userClient.UserInfo(context.Background(), &userApi.UserInfoRequest{UserId: req.UserID})
	if err != nil {
		c.Error(err)
		return
	}

	if user == nil {
		response.Fail(c, "用户不存在", nil)
		return
	}

	// 进行添加黑名单操作
	if _, err = relationClient.AddBlacklist(context.Background(), &relationApi.AddBlacklistRequest{UserId: req.UserID, FriendId: req.FriendID}); err != nil {
		c.Error(err)
		return
	}

	response.Success(c, "添加到黑名单成功", gin.H{})
}

type deleteFriendRequest struct {
	UserID   string `json:"user_id" binding:"required"`
	FriendID string `json:"friend_id" binding:"required"`
}

// @Summary 删除好友
// @Description 删除好友
// @Accept  json
// @Produce  json
// @param request body deleteFriendRequest true "request"
// @Success		200 {object} utils.Response{}
// @Router /relation/delete_friend [post]
func deleteFriend(c *gin.Context) {
	req := new(deleteFriendRequest)
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("参数验证失败", zap.Error(err))
		response.Fail(c, "参数验证失败", nil)
		return
	}

	// 检查用户是否存在
	user, err := userClient.UserInfo(context.Background(), &userApi.UserInfoRequest{UserId: req.UserID})
	if err != nil {
		c.Error(err)
		return
	}

	if user == nil {
		response.Fail(c, "用户不存在", nil)
		return
	}

	// 进行删除好友操作
	if _, err = relationClient.DeleteFriend(context.Background(), &relationApi.DeleteFriendRequest{UserId: req.UserID, FriendId: req.FriendID}); err != nil {
		c.Error(err)
		return
	}

	response.Success(c, "删除好友成功", gin.H{})
}

type confirmFriendRequest struct {
	UserID   string `json:"user_id" binding:"required"`
	FriendID string `json:"friend_id" binding:"required"`
}

// @Summary 确认添加好友
// @Description 确认添加好友
// @Accept  json
// @Produce  json
// @param request body confirmFriendRequest true "request"
// @Success		200 {object} utils.Response{}
// @Router /relation/confirm_friend [post]
func confirmFriend(c *gin.Context) {
	req := new(confirmFriendRequest)
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("参数验证失败", zap.Error(err))
		response.Fail(c, "参数验证失败", nil)
		return
	}

	// 检查用户是否存在
	user, err := userClient.UserInfo(context.Background(), &userApi.UserInfoRequest{UserId: req.UserID})
	if err != nil {
		c.Error(err)
		return
	}

	if user == nil {
		response.Fail(c, "用户不存在", nil)
		return
	}

	// 进行确认好友操作
	if _, err = relationClient.ConfirmFriend(context.Background(), &relationApi.ConfirmFriendRequest{UserId: req.UserID, FriendId: req.FriendID}); err != nil {
		c.Error(err)
		return
	}

	response.Success(c, "确认好友成功", gin.H{})
}

type addFriendRequest struct {
	FriendId string `json:"friend_id" binding:"required"`
	Msg      string `json:"msg"`
}

// @Summary 添加好友
// @Description 添加好友
// @Accept  json
// @Produce  json
// @param request body addFriendRequest true "request"
// @Success		200 {object} utils.Response{}
// @Router /relation/add_friend [post]
func addFriend(c *gin.Context) {
	req := new(addFriendRequest)
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("参数验证失败", zap.Error(err))
		response.Fail(c, "参数验证失败", nil)
		return
	}

	thisId, err := pkghttp.ParseTokenReUid(c)
	if err != nil {
		response.Fail(c, err.Error(), nil)
		return
	}

	user, err := userClient.UserInfo(context.Background(), &userApi.UserInfoRequest{UserId: req.FriendId})
	if err != nil {
		c.Error(err)
		return
	}
	if user == nil {
		response.Fail(c, "用户不存在", nil)
		return
	}

	if _, err := relationClient.AddFriend(context.Background(), &relationApi.AddFriendRequest{UserId: thisId, FriendId: user.UserId}); err != nil {
		c.Error(err)
		return
	}

	response.Success(c, "发送好友请求成功", gin.H{})
}