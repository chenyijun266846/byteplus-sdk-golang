package cdn

import (
	"testing"

	"github.com/chenyijun266846/byteplus-sdk-golang/service/cdn"
	"github.com/stretchr/testify/assert"
)

func DescribeCdnDataDetail(t *testing.T) {
	resp, err := DefaultInstance.DescribeCdnDataDetail(&cdn.DescribeCdnDataDetailRequest{
		StartTime: testStartTime,
		EndTime:   testEndTime,
		Metric:    "flux",
		Domain:    exampleDomain,
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Result.Name)
	assert.NotEmpty(t, resp.Result.DataDetails)
}
