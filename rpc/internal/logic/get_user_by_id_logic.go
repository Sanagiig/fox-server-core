package logic

import (
	"context"

	"fox-server-core/common/errs"
	"fox-server-core/rpc/ent/user"
	"fox-server-core/rpc/internal/svc"
	"fox-server-core/rpc/types/core"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/suyuan32/simple-admin-common/utils/uuidx"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByIdLogic {
	return &GetUserByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: user
func (l *GetUserByIdLogic) GetUserById(in *core.UUIDReq) (*core.UserInfo, error) {
	u, err := l.svcCtx.DB.User.Query().Where(user.IDEQ(uuidx.ParseUUIDString(in.Id))).Only(l.ctx)
	if err != nil {
		return nil, errs.DBNotFoundERR
		// return nil, err
	}
	return &core.UserInfo{
		Id:          pointy.GetPointer(u.ID.String()),
		Username:    &u.Username,
		Nickname:    &u.Nickname,
		HomePath:    &u.HomePath,
		Password:    &u.Password,
		Description: &u.Description,
		Avatar:      &u.Avatar,
		Mobile:      &u.Mobile,
		Email:       &u.Email,
	}, nil
}
