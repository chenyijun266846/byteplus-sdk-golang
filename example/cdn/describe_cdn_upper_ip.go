package cdn

import (
	"testing"

	"github.com/chenyijun266846/byteplus-sdk-golang/service/cdn"
	"github.com/stretchr/testify/assert"
)

func DescribeCdnUpperIp(t *testing.T) {
	resp, err := DefaultInstance.DescribeCdnUpperIp(&cdn.DescribeCdnUpperIpRequest{Domain: exampleDomain})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}
