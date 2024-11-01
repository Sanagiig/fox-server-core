package base

import (
	"context"

	"fox-server-core/api/internal/svc"
	"fox-server-core/api/internal/types"
	"fox-server-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type InitDatabaseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInitDatabaseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InitDatabaseLogic {
	return &InitDatabaseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InitDatabaseLogic) InitDatabase() (*types.BaseMsgResp, error) {
	_, err := l.svcCtx.CoreRpc.InitDatabase(l.ctx, &core.Empty{})
	if err != nil {
		return nil, err
	}

	return &types.BaseMsgResp{Msg: "初始化数据库成功"}, nil
}
