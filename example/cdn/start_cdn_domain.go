package cdn

import (
	"testing"

	"github.com/chenyijun266846/byteplus-sdk-golang/service/cdn"
	"github.com/stretchr/testify/assert"
)

func StartCdnDomain(t *testing.T) {
	resp, err := DefaultInstance.StartCdnDomain(&cdn.StartCdnDomainRequest{Domain: operateDomain})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}
