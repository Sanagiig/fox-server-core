package errs

type BaseErrResp struct {
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
	ErrMsg string `json:"errMsg"`
	Data   any    `json:"data"`
}
