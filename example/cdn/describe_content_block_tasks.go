package cdn

import (
	"testing"

	"github.com/chenyijun266846/byteplus-sdk-golang/service/cdn"
	"github.com/stretchr/testify/assert"
)

func DescribeContentBlockTasks(t *testing.T) {
	resp, err := DefaultInstance.DescribeContentBlockTasks(&cdn.DescribeContentBlockTasksRequest{
		TaskType:  "block_url",
		StartTime: &testStartTime,
		EndTime:   &testEndTime,
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, resp.Result.Data)
}
