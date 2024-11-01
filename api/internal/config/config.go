package config

import (
	"github.com/suyuan32/simple-admin-common/config"
	"github.com/suyuan32/simple-admin-common/utils/captcha"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	CoreRpcConf zrpc.RpcClientConf
	Auth        AuthConf
	Captcha     captcha.Conf
	RedisConf   config.RedisConf
}

// AuthConf is a JWT config
type AuthConf struct {
	AccessSecret string `json:",optional,env=AUTH_SECRET"`
	AccessExpire int64  `json:",optional,env=AUTH_EXPIRE"`
}
