package logic

import (
	"context"

	"fox-server-core/rpc/ent"
	"fox-server-core/rpc/internal/svc"
	"fox-server-core/rpc/types/core"

	"github.com/suyuan32/simple-admin-common/utils/encrypt"
	"github.com/zeromicro/go-zero/core/logx"
)

type InitDatabaseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInitDatabaseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InitDatabaseLogic {
	return &InitDatabaseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: base
func (l *InitDatabaseLogic) InitDatabase(in *core.Empty) (*core.BaseResp, error) {
	err := l.svcCtx.DB.Schema.Create(l.ctx)
	if err != nil {
		return nil, err
	}

	if err := l.insertUserData(); err != nil {
		return nil, err
	}

	return &core.BaseResp{}, nil
}

func (l *InitDatabaseLogic) insertUserData() error {
	var users []*ent.UserCreate
	users = append(users, l.svcCtx.DB.User.Create().
		SetUsername("admin").
		SetNickname("admin").
		SetPassword(encrypt.BcryptEncrypt("simple-admin")).
		SetEmail("simple_admin@gmail.com").
		SetDepartmentID(1))

	err := l.svcCtx.DB.User.CreateBulk(users...).Exec(l.ctx)
	if err != nil {
		logx.Errorw(err.Error())
		return err
	} else {
		return nil
	}
}
