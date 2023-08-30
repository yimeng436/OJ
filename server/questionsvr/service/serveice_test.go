package service

import (
	"encoding/json"
	"fmt"
	"testing"
)

type ListString struct {
	tags []string
}

type String struct {
	tags string
}

func TestJson(t *testing.T) {
	s := &String{
		tags: "[1,2,3]",
	}
	res, err := json.Marshal(s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}
