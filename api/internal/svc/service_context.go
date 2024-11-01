package svc

import (
	"fox-server-core/api/internal/config"
	"fox-server-core/api/internal/middleware"
	"fox-server-core/rpc/core_client"

	"github.com/mojocn/base64Captcha"
	"github.com/redis/go-redis/v9"
	"github.com/suyuan32/simple-admin-common/utils/captcha"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config    config.Config
	Authority rest.Middleware
	CoreRpc   core_client.Core
	Redis     redis.UniversalClient
	Captcha   *base64Captcha.Captcha
}

func NewServiceContext(c config.Config) *ServiceContext {
	rds := c.RedisConf.MustNewUniversalRedis()

	return &ServiceContext{
		Config:    c,
		Authority: middleware.NewAuthorityMiddleware().Handle,
		CoreRpc:   core_client.NewCore(zrpc.MustNewClient(c.CoreRpcConf)),
		Captcha:   captcha.MustNewOriginalRedisCaptcha(c.Captcha, rds),
		Redis:     c.RedisConf.MustNewUniversalRedis(),
	}
}
