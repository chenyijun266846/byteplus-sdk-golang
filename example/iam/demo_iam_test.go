package main

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/chenyijun266846/byteplus-sdk-golang/service/iam"
)

const (
	region = "ap-singapore-1"
	testAk = "ak"
	testSk = "sk"
)

func TestIAM(t *testing.T) {
	//iam.DefaultInstance.Client.SetAccessKey(testAk)
	//iam.DefaultInstance.Client.SetSecretKey(testSk)
	iam.DefaultInstance.SetRegion(region)

	list, status, err := iam.DefaultInstance.ListUsers(nil)
	fmt.Println(status, err)
	b, _ := json.Marshal(list)
	fmt.Println(string(b))
}
