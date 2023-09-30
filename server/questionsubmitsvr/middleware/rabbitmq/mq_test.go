package rabbitmq

import (
	"fmt"
	"questionsvr/config"
	"questionsvr/log"
	"testing"
)

func TestMethod(t *testing.T) {
	err := config.Init()
	if err != nil {
		fmt.Println(err)
	}
	log.InitLog()
	mq := GetMQ("question_test", "")
	fmt.Println(mq)

}
