package cdn

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/chenyijun266846/byteplus-sdk-golang/service/cdn"
	"github.com/stretchr/testify/assert"
)

func ListCdnDomains(t *testing.T) {
	resp, err := DefaultInstance.ListCdnDomains(&cdn.ListCdnDomainsRequest{})
	test, _ := json.Marshal(resp)
	fmt.Println(string(test))
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Result.Data)
}
