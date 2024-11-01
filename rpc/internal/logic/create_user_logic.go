package logic

import (
	"context"

	"fox-server-core/rpc/internal/svc"
	"fox-server-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// User management
func (l *CreateUserLogic) CreateUser(in *core.UserInfo) (*core.BaseUUIDResp, error) {

	u, err := l.svcCtx.DB.User.Create().
		SetUsername(*in.Username).
		SetNickname(*in.Nickname).
		SetNillableHomePath(in.HomePath).
		SetPassword(*in.Password).
		SetNillableEmail(in.Email).
		SetNillableAvatar(in.Avatar).
		SetNillableDescription(in.Description).
		Save(l.ctx)

	if err != nil {
		return nil, err
	}

	return &core.BaseUUIDResp{
		Id: u.ID.String(),
	}, nil
}
