package service

import "github.com/yimeng436/OJ/pkg/pb"

type CodeSandFactory struct {
}

func (CodeSandFactory) CreateCodeSand(name string) pb.CodeSandServiceServer {
	switch name {
	case "normal":
		return new(CodeSandService)
	case "remote":
		return new(RemoteCodeSandService)
	case "thirdpart":
		return new(ThirdPartCodeSandService)
	}
	return nil
}
