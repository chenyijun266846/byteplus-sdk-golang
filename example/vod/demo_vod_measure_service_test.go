// Code generated by protoc-gen-byteplus-sdk
// source: VodMeasureService
// DO NOT EDIT!

package vod

import (
	"fmt"
	"testing"

	"github.com/chenyijun266846/byteplus-sdk-golang/base"
	"github.com/chenyijun266846/byteplus-sdk-golang/service/vod"
	"github.com/chenyijun266846/byteplus-sdk-golang/service/vod/models/request"
)

func Test_DescribeVodSpaceTranscodeData(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.DescribeVodSpaceTranscodeDataRequest{
		SpaceList:       "your SpaceList",
		StartTime:       "your StartTime",
		EndTime:         "your EndTime",
		TranscodeType:   "your TranscodeType",
		Specification:   "your Specification",
		TaskStageList:   "your TaskStageList",
		Aggregation:     0,
		DetailFieldList: "your DetailFieldList",
	}

	resp, status, err := instance.DescribeVodSpaceTranscodeData(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_DescribeVodSnapshotData(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.DescribeVodSnapshotDataRequest{
		SpaceList:       "your SpaceList",
		StartTime:       "your StartTime",
		EndTime:         "your EndTime",
		SnapshotType:    "your SnapshotType",
		TaskStageList:   "your TaskStageList",
		Aggregation:     0,
		DetailFieldList: "your DetailFieldList",
	}

	resp, status, err := instance.DescribeVodSnapshotData(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}
