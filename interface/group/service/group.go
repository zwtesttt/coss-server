package service

import (
	"context"
	"errors"
	"github.com/cossim/coss-server/interface/group/api/model"
	"github.com/cossim/coss-server/pkg/code"
	"github.com/cossim/coss-server/pkg/constants"
	"github.com/cossim/coss-server/pkg/msg_queue"
	"github.com/cossim/coss-server/pkg/utils/time"
	groupgrpcv1 "github.com/cossim/coss-server/service/group/api/v1"
	relationgrpcv1 "github.com/cossim/coss-server/service/relation/api/v1"
	"github.com/dtm-labs/client/dtmcli"
	"github.com/dtm-labs/client/dtmgrpc"
	"github.com/dtm-labs/client/workflow"
	"github.com/lithammer/shortuuid/v3"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Service) CreateGroup(ctx context.Context, req *groupgrpcv1.Group) (*model.CreateGroupResponse, error) {
	var err error
	friends, err := s.relationUserClient.GetUserRelationByUserIds(ctx, &relationgrpcv1.GetUserRelationByUserIdsRequest{UserId: req.CreatorId, FriendIds: req.Member})
	if err != nil {
		s.logger.Error("获取好友关系失败", zap.Error(err))
		return nil, code.RelationErrCreateGroupFailed
	}

	if len(req.Member) != len(friends.Users) {
		return nil, code.RelationUserErrFriendRelationNotFound
	}
	for _, friend := range friends.Users {
		if friend.Status != relationgrpcv1.RelationStatus_RELATION_NORMAL {
			return nil, code.StatusNotAvailable
		}
	}

	r1 := &groupgrpcv1.CreateGroupRequest{Group: &groupgrpcv1.Group{
		Type:            req.Type,
		MaxMembersLimit: req.MaxMembersLimit,
		CreatorId:       req.CreatorId,
		Name:            req.Name,
		Avatar:          req.Avatar,
		Member:          req.Member,
	}}

	r22 := &relationgrpcv1.CreateGroupAndInviteUsersRequest{
		UserID: req.CreatorId,
		Member: req.Member,
	}

	resp1 := &groupgrpcv1.Group{}
	var groupID uint32
	var DialogID uint32
	// 创建 DTM 分布式事务工作流
	workflow.InitGrpc(s.dtmGrpcServer, s.relationGrpcServer, grpc.NewServer())
	gid := shortuuid.New()
	wfName := "create_group_workflow_" + gid
	if err = workflow.Register(wfName, func(wf *workflow.Workflow, data []byte) error {
		// 创建群聊
		resp1, err = s.groupClient.CreateGroup(wf.Context, r1)
		if err != nil {
			return err
		}
		wf.NewBranch().OnRollback(func(bb *dtmcli.BranchBarrier) error {
			_, err = s.groupClient.CreateGroupRevert(wf.Context, &groupgrpcv1.CreateGroupRequest{Group: &groupgrpcv1.Group{
				Id: resp1.Id,
			}})
			return err
		})
		groupID = resp1.Id
		r22.GroupId = groupID
		r22.Member = req.Member
		r22.UserID = req.CreatorId
		resp2, err := s.relationGroupClient.CreateGroupAndInviteUsers(wf.Context, r22)
		if err != nil {
			return err
		}
		DialogID = resp2.DialogId
		wf.NewBranch().OnRollback(func(bb *dtmcli.BranchBarrier) error {
			_, err = s.relationGroupClient.CreateGroupAndInviteUsersRevert(wf.Context, r22)
			return err
		})

		return err
	}); err != nil {
		s.logger.Error("WorkFlow CreateGroup", zap.Error(err))
		return nil, code.RelationErrCreateGroupFailed
	}
	if err = workflow.Execute(wfName, gid, nil); err != nil {
		s.logger.Error("WorkFlow CreateGroup", zap.Error(err))
		return nil, code.RelationErrCreateGroupFailed
	}

	// 给被邀请的用户推送
	for _, id := range req.Member {
		msg := constants.WsMsg{Uid: id, Event: constants.InviteJoinGroupEvent, Data: map[string]interface{}{"group_id": groupID, "inviter_id": req.CreatorId}, SendAt: time.Now()}
		//通知消息服务有消息需要发送
		err = s.rabbitMQClient.PublishServiceMessage(msg_queue.RelationService, msg_queue.MsgService, msg_queue.Service_Exchange, msg_queue.SendMessage, msg)
	}

	return &model.CreateGroupResponse{
		Id:              resp1.Id,
		Avatar:          resp1.Avatar,
		Name:            resp1.Name,
		Type:            uint32(resp1.Type),
		Status:          int32(resp1.Status),
		MaxMembersLimit: resp1.MaxMembersLimit,
		CreatorId:       resp1.CreatorId,
		DialogId:        DialogID,
	}, nil
}

func (s *Service) DeleteGroup(ctx context.Context, groupID uint32, userID string) (uint32, error) {
	_, err := s.groupClient.GetGroupInfoByGid(ctx, &groupgrpcv1.GetGroupInfoRequest{
		Gid: groupID,
	})
	if err != nil {
		return 0, err
	}
	sf, err := s.relationGroupClient.GetGroupRelation(ctx, &relationgrpcv1.GetGroupRelationRequest{
		UserId:  userID,
		GroupId: groupID,
	})
	if err != nil {
		return 0, err
	}
	if sf.Identity == relationgrpcv1.GroupIdentity_IDENTITY_USER {
		return 0, errors.New("不允许操作")
	}

	dialog, err := s.relationDialogClient.GetDialogByGroupId(ctx, &relationgrpcv1.GetDialogByGroupIdRequest{GroupId: groupID})
	if err != nil {
		return 0, err
	}

	r1 := &relationgrpcv1.DeleteDialogUsersByDialogIDRequest{
		DialogId: dialog.DialogId,
	}
	r2 := &relationgrpcv1.DeleteDialogByIdRequest{
		DialogId: dialog.DialogId,
	}
	r3 := &relationgrpcv1.GroupIDRequest{
		GroupId: groupID,
	}
	r4 := &groupgrpcv1.DeleteGroupRequest{
		Gid: groupID,
	}
	gid := shortuuid.New()
	if err = dtmgrpc.TccGlobalTransaction(s.dtmGrpcServer, gid, func(tcc *dtmgrpc.TccGrpc) error {
		r := &emptypb.Empty{}
		// 删除对话用户
		if err = tcc.CallBranch(r1, s.relationGrpcServer+relationgrpcv1.DialogService_DeleteDialogUsersByDialogID_FullMethodName, "", s.relationGrpcServer+relationgrpcv1.DialogService_DeleteDialogUsersByDialogIDRevert_FullMethodName, r); err != nil {
			return err
		}
		// 删除对话
		if err = tcc.CallBranch(r2, s.relationGrpcServer+relationgrpcv1.DialogService_DeleteDialogById_FullMethodName, "", s.relationGrpcServer+relationgrpcv1.DialogService_DeleteDialogByIdRevert_FullMethodName, r); err != nil {
			return err
		}
		// 删除群聊成员
		if err = tcc.CallBranch(r3, s.relationGrpcServer+relationgrpcv1.GroupRelationService_DeleteGroupRelationByGroupId_FullMethodName, "", s.relationGrpcServer+relationgrpcv1.GroupRelationService_DeleteGroupRelationByGroupIdRevert_FullMethodName, r); err != nil {
			return err
		}
		// 删除群聊
		if err = tcc.CallBranch(r4, s.groupGrpcServer+groupgrpcv1.GroupService_DeleteGroup_FullMethodName, "", s.groupGrpcServer+groupgrpcv1.GroupService_DeleteGroupRevert_FullMethodName, r); err != nil {
			return err
		}
		return err
	}); err != nil {
		return 0, err
	}

	return groupID, err
}

func (s *Service) UpdateGroup(ctx context.Context, req *model.UpdateGroupRequest, userID string) (interface{}, error) {
	group, err := s.groupClient.GetGroupInfoByGid(ctx, &groupgrpcv1.GetGroupInfoRequest{
		Gid: req.GroupId,
	})
	if err != nil {
		s.logger.Error("更新群聊信息失败", zap.Error(err))
		return nil, err
	}

	sf, err := s.relationGroupClient.GetGroupRelation(ctx, &relationgrpcv1.GetGroupRelationRequest{
		UserId:  userID,
		GroupId: req.GroupId,
	})
	if err != nil {
		s.logger.Error("获取用户群聊关系失败", zap.Error(err))
		return nil, err
	}

	if sf.Identity == relationgrpcv1.GroupIdentity_IDENTITY_USER {
		return nil, code.Unauthorized
	}

	group.Type = groupgrpcv1.GroupType(req.Type)
	group.Name = req.Name
	group.Avatar = req.Avatar
	group.Id = req.GroupId
	switch req.Type {
	case uint32(groupgrpcv1.GroupType_TypeEncrypted):
		group.MaxMembersLimit = model.EncryptedGroup
	default:
		group.MaxMembersLimit = model.DefaultGroup
	}

	resp, err := s.groupClient.UpdateGroup(ctx, &groupgrpcv1.UpdateGroupRequest{
		Group: group,
	})
	if err != nil {
		s.logger.Error("更新群聊信息失败", zap.Error(err))
		return nil, err
	}

	return resp, nil
}

func (s *Service) GetBatchGroupInfoByIDs(ctx context.Context, ids []uint32) (interface{}, error) {
	groups, err := s.groupClient.GetBatchGroupInfoByIDs(ctx, &groupgrpcv1.GetBatchGroupInfoRequest{
		GroupIds: ids,
	})
	if err != nil {
		s.logger.Error("批量获取群聊信息失败", zap.Error(err))
		return nil, err
	}

	return groups.Groups, nil
}

func (s *Service) GetGroupInfoByGid(ctx context.Context, gid uint32, userID string) (interface{}, error) {
	group, err := s.groupClient.GetGroupInfoByGid(ctx, &groupgrpcv1.GetGroupInfoRequest{
		Gid: gid,
	})
	if err != nil {
		return nil, err
	}

	id, err := s.relationDialogClient.GetDialogByGroupId(ctx, &relationgrpcv1.GetDialogByGroupIdRequest{GroupId: gid})
	if err != nil {
		return nil, err
	}

	relation, err := s.relationGroupClient.GetGroupRelation(ctx, &relationgrpcv1.GetGroupRelationRequest{
		GroupId: gid,
		UserId:  userID,
	})
	if err != nil {
		return nil, err
	}

	per := &model.Preferences{
		OpenBurnAfterReading: model.OpenBurnAfterReadingType(relation.OpenBurnAfterReading),
		SilentNotification:   model.SilentNotification(relation.IsSilent),
		Remark:               relation.Remark,
		EntryMethod:          model.EntryMethod(relation.JoinMethod),
		GroupNickname:        relation.GroupNickname,
		Inviter:              relation.Inviter,
		JoinedAt:             relation.JoinTime,
		MuteEndTime:          relation.MuteEndTime,
		Identity:             model.GroupIdentity(relation.Identity),
	}

	return &model.GroupInfo{
		Id:              group.Id,
		Avatar:          group.Avatar,
		Name:            group.Name,
		Type:            uint32(group.Type),
		Status:          int32(group.Status),
		MaxMembersLimit: group.MaxMembersLimit,
		CreatorId:       group.CreatorId,
		DialogId:        id.DialogId,
		Preferences:     per,
	}, nil
}
