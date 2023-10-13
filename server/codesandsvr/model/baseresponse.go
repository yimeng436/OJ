package model

import "github.com/yimeng436/OJ/pkg/pb"

type BaseResponse struct {
	Code    int64                  `json:"code"`
	Data    pb.ExecuteCodeResponse `json:"data"`
	Message string                 `json:"message"`
}
