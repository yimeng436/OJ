package service

import (
	"bytes"
	"codesandsvr/config"
	"codesandsvr/model"
	"context"
	"encoding/json"
	"errors"
	"github.com/yimeng436/OJ/pkg/pb"
	"io"
	"net/http"
)

type RemoteCodeSandService struct {
	pb.UnimplementedCodeSandServiceServer
}

func (RemoteCodeSandService) ExecuteCode(ctx context.Context, request *pb.ExecuteCodeRequest) (*pb.ExecuteCodeResponse, error) {

	requestBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, "http://localhost:8080/executeCode", bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set(config.GetGlobalConfig().CodeSand.AuthHead, config.GetGlobalConfig().CodeSand.AuthKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New("请求/localhost:8080/executeCode：异常")
	}

	// 把resp的响应体从reader读为 []byte
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	baseResp := new(model.BaseResponse)
	err = json.Unmarshal(respBody, baseResp)
	if err != nil {
		return nil, err
	}
	executeResp := new(pb.ExecuteCodeResponse)
	if baseResp.Code != 0 {
		executeResp.Status = baseResp.Message
		executeResp.Message = baseResp.Message
		executeResp.JudgeInfo = &pb.JudgeInfo{
			Message: baseResp.Message,
			Time:    0,
			Memory:  0,
		}
	} else {
		executeResp = &baseResp.Data
	}
	return executeResp, nil
}
