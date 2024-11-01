package public_user

import (
	"context"

	"fox-server-core/api/internal/svc"
	"fox-server-core/api/internal/types"
	"fox-server-core/rpc/core_client"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (*types.BaseMsgResp, error) {
	// match := l.svcCtx.Captcha.Verify(req.CaptchaId, req.Captcha, true)
	// if !match {
	// 	return &types.BaseMsgResp{Msg: "验证码错误"}, nil
	// }

	_, err := l.svcCtx.CoreRpc.CreateUser(l.ctx, &core_client.UserInfo{
		Username: &req.Username,
		Password: &req.Password,
		Email:    &req.Email,
	})

	if err != nil {
		return &types.BaseMsgResp{Msg: "注册失败"}, err
	}

	return &types.BaseMsgResp{
		Code: 200,
		Msg:  "注册成功",
	}, nil
}
