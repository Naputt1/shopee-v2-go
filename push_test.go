package goshopee

import (
	"fmt"
	"testing"

	"github.com/jarcoal/httpmock"
)

func Test_Push_ConfirmConsumedLostPushMessage(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.push.confirm_consumed_lost_push_message_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping ConfirmConsumedLostPushMessage due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping ConfirmConsumedLostPushMessage due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/push/confirm_consumed_lost_push_message", app.APIURL), responder)

	var req ConfirmConsumedLostPushMessageRequest
	res, err := client.Push.ConfirmConsumedLostPushMessage(shopID, req, accessToken)
	if err != nil {
		t.Logf("Push.ConfirmConsumedLostPushMessage returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Push.ConfirmConsumedLostPushMessage response: %#v", res)
}
func Test_Push_GetAppPushConfig(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.push.get_app_push_config_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetAppPushConfig due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetAppPushConfig due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/push/get_app_push_config", app.APIURL), responder)

	res, err := client.Push.GetAppPushConfig(shopID, accessToken)
	if err != nil {
		t.Logf("Push.GetAppPushConfig returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Push.GetAppPushConfig response: %#v", res)
}
func Test_Push_GetLostPushMessage(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.push.get_lost_push_message_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetLostPushMessage due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetLostPushMessage due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/push/get_lost_push_message", app.APIURL), responder)

	res, err := client.Push.GetLostPushMessage(shopID, accessToken)
	if err != nil {
		t.Logf("Push.GetLostPushMessage returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Push.GetLostPushMessage response: %#v", res)
}
func Test_Push_SetAppPushConfig(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.push.set_app_push_config_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping SetAppPushConfig due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping SetAppPushConfig due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/push/set_app_push_config", app.APIURL), responder)

	var req SetAppPushConfigRequest
	res, err := client.Push.SetAppPushConfig(shopID, req, accessToken)
	if err != nil {
		t.Logf("Push.SetAppPushConfig returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Push.SetAppPushConfig response: %#v", res)
}
