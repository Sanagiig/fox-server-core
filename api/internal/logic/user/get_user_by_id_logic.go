package user

import (
	"context"

	"fox-server-core/api/internal/svc"
	"fox-server-core/api/internal/types"
	"fox-server-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByIdLogic {
	return &GetUserByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserByIdLogic) GetUserById(req *types.UUIDReq) (*types.UserInfoResp, error) {
	res, err := l.svcCtx.CoreRpc.GetUserById(l.ctx, &core.UUIDReq{Id: req.Id})

	if err != nil {
		return nil, err
	}
	return &types.UserInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 200,
			Msg:  "获取用户成功",
		},
		Data: types.UserInfo{
			BaseUUIDInfo: types.BaseUUIDInfo{
				Id:        res.Id,
				CreatedAt: res.CreatedAt,
				UpdatedAt: res.UpdatedAt,
			},
			Username: res.Username,
			Nickname: res.Nickname,
			HomePath: res.HomePath,
			Email:    res.Email,
			Avatar:   res.Avatar,
			Status:   res.Status,
		},
	}, nil
}
