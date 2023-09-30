package service

import (
	"fmt"
	"github.com/yimeng436/OJ/pkg/pb"
	"testing"
)

func TestMethod(t *testing.T) {
	remoteservice := new(RemoteCodeSandService)
	request := new(pb.ExecuteCodeRequest)
	request.Code = "class Main{\n    public static void main(String[] args) {\n        System.out.println(args[0]+args[1]);\n    }\n}"
	request.Language = "java"
	request.Inputs = []string{"1 2", "3 4"}
	resp, err := remoteservice.ExecuteCode(nil, request)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)
}
