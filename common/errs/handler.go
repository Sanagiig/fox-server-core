package errs

import (
	"errors"
	"fmt"
)

func DefaultErrHandler(err error) (int, any) {
	fmt.Printf("err: %v\n", err)
	switch {
	case errors.Is(err, DBNotFoundERR):
		return 200, BaseErrResp{
			Code:   200,
			Msg:    "DB not found",
			ErrMsg: err.Error(),
		}
	}

	return 200, BaseErrResp{
		Code:   500,
		Msg:    "Internal Server Error",
		ErrMsg: err.Error(),
	}
}
