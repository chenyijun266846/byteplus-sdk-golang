package cdn

import (
	"testing"

	"github.com/chenyijun266846/byteplus-sdk-golang/service/cdn"
	"github.com/stretchr/testify/assert"
)

func DescribeCdnRegionAndIsp(t *testing.T) {
	area := "China"
	resp, err := DefaultInstance.DescribeCdnRegionAndIsp(&cdn.DescribeCdnRegionAndIspRequest{Area: &area})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Result.Isps)
	assert.NotEmpty(t, resp.Result.Regions)
}
