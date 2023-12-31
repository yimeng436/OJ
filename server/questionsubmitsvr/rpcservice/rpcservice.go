package rpcservice

import (
	"context"
	"fmt"
	"github.com/yimeng436/OJ/pkg/pb"
	"google.golang.org/grpc"
	"questionsubmitsvr/config"
	"questionsubmitsvr/log"
	"time"
)

var (
	UserSvrClient     pb.UserServiceClient
	QuestionSvrClient pb.QuestionServiceClient
	JudgeSvrClient    pb.JudgeServiceClient
)

func InitSvrConn() {
	UserSvrClient = NewUserSvrClient(config.GetGlobalConfig().SvrConfig.UserSvrName)
	QuestionSvrClient = NewQuestionSvrClient(config.GetGlobalConfig().SvrConfig.QuestionSvrName)
	JudgeSvrClient = NewJudgeSvrClient(config.GetGlobalConfig().SvrConfig.JudgeSvrName)
}

func NewJudgeSvrClient(svrName string) pb.JudgeServiceClient {
	conn, err := NewSvrConn(svrName)
	if err != nil {
		return nil
	}
	return pb.NewJudgeServiceClient(conn)
}

func NewUserSvrClient(svrName string) pb.UserServiceClient {
	conn, err := NewSvrConn(svrName)
	if err != nil {
		return nil
	}
	return pb.NewUserServiceClient(conn)
}

func NewQuestionSvrClient(svrName string) pb.QuestionServiceClient {
	conn, err := NewSvrConn(svrName)
	if err != nil {
		return nil
	}
	return pb.NewQuestionServiceClient(conn)
}

func NewSvrConn(svrName string) (*grpc.ClientConn, error) {
	consulInfo := config.GetGlobalConfig().ConsulConfig
	fmt.Println(consulInfo)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	conn, err := grpc.DialContext(
		ctx,
		//这里必须要引入上面github.com/mbobakov/grpc-consul-resolver 这个包
		//因为底层是通过这个解析器将这个请求转为http请求找到对应的服务
		//只有把这个包引进来了，才会执行这个包里面的consul.go的init函数，将解析器注入
		fmt.Sprintf("consul://%s:%d/%s?wait=14s", consulInfo.Host, consulInfo.Port, svrName),
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		log.Errorf("NewSvrConn with svrname %s err:%v", svrName, err)
		return nil, err
	}
	log.Info("NewSvrConn success")
	return conn, nil
}
func GetUserServiceClient() pb.UserServiceClient {
	return UserSvrClient
}

func GetQuestionSvrClient() pb.QuestionServiceClient {
	return QuestionSvrClient
}

func GetJudgeServiceClient() pb.JudgeServiceClient {
	return JudgeSvrClient
}
