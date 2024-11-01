package base

import (
	"context"

	"fox-server-core/api/internal/svc"
	"fox-server-core/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InitJobDatabaseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInitJobDatabaseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InitJobDatabaseLogic {
	return &InitJobDatabaseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InitJobDatabaseLogic) InitJobDatabase() (resp *types.BaseMsgResp, err error) {
	// todo: add your logic here and delete this line

	return
}
