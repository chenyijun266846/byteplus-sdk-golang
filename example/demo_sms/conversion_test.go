package demo_sms

import (
	"testing"

	"github.com/chenyijun266846/byteplus-sdk-golang/base"
	"github.com/chenyijun266846/byteplus-sdk-golang/service/sms"
)

/*
return message_ids(the message_ids are the ones in Send/BatchSend response body) which have converted
*/
func TestSmsConversion(t *testing.T) {
	i18nInstance := sms.NewInstanceI18n(base.RegionApSingapore)
	i18nInstance.Client.SetAccessKey(testAk)
	i18nInstance.Client.SetSecretKey(testSk)
	req := &sms.ConversionRequest{
		MessageIDs: []string{"MessageID"},
	}
	result, statusCode, err := i18nInstance.Conversion(req)
	t.Logf("result = %+v\n", result)
	t.Logf("statusCode = %+v\n", statusCode)
	t.Logf("err = %+v\n", err)
}
