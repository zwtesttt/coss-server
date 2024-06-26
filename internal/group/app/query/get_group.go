package query

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/cossim/coss-server/internal/group/app/command"
	"github.com/cossim/coss-server/internal/group/domain/repository"
	"github.com/cossim/coss-server/pkg/code"
	"github.com/cossim/coss-server/pkg/decorator"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type GetGroup struct {
	ID     uint32 `json:"id"`
	UserID string `json:"user_id"`
}

type GetGroupHandler decorator.CommandHandler[GetGroup, *GroupInfo]

var _ decorator.CommandHandler[GetGroup, *GroupInfo] = &getGroupHandler{}

type getGroupHandler struct {
	groupRepo             repository.Repository
	relationGroupService  command.RelationGroupService
	relationDialogService command.RelationDialogService
	logger                *zap.Logger

	dtmGrpcServer string
}

func NewGetGroupHandler(
	repo repository.Repository,
	logger *zap.Logger,
	dtmGrpcServer string,
	relationGroupService command.RelationGroupService,
	relationDialogService command.RelationDialogService,
) GetGroupHandler {
	if repo == nil {
		panic("nil repo")
	}

	h := &getGroupHandler{
		groupRepo: repo,
		logger:    logger,
		//dtmGrpcServer: dtmGrpcServer,
		relationGroupService:  relationGroupService,
		relationDialogService: relationDialogService,
	}
	return h
}

func (h *getGroupHandler) Handle(ctx context.Context, cmd GetGroup) (*GroupInfo, error) {
	fmt.Println("get group => ", cmd)
	r, err := h.groupRepo.Get(ctx, cmd.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, code.GroupErrGroupNotFound
		}
		h.logger.Error("get group failed", zap.Error(err))
		return nil, err
	}

	dialogID, err := h.relationDialogService.GetGroupDialogID(ctx, r.ID)
	if err != nil {
		return nil, err
	}

	relation, err := h.relationGroupService.GetRelation(ctx, r.ID, cmd.UserID)
	if err != nil {
		return nil, err
	}

	marshal, err := json.Marshal(relation)
	if err != nil {
		return nil, err
	}

	fmt.Println("marshal => ", string(marshal))

	per := &Preferences{}
	if relation != nil && relation.ID != 0 {
		per = &Preferences{
			SilentNotification: relation.IsSilent,
			Remark:             relation.Remark,
			EntryMethod:        relation.JoinMethod,
			Inviter:            relation.Inviter,
			JoinedAt:           relation.JoinTime,
			MuteEndTime:        relation.MuteEndTime,
			Identity:           relation.Identity,
		}
	}

	return &GroupInfo{
		Id:              r.ID,
		Avatar:          r.Avatar,
		Name:            r.Name,
		Type:            uint(r.Type),
		Status:          int(r.Status),
		MaxMembersLimit: int32(r.MaxMembersLimit),
		CreatorId:       r.CreatorID,
		DialogId:        dialogID,
		SilenceTime:     r.SilenceTime,
		Encrypt:         r.Encrypt,
		JoinApprove:     r.JoinApprove,
		Preferences:     per,
	}, nil
}
