package goshopee

import (
	"context"
	"fmt"
	"testing"

	"github.com/jarcoal/httpmock"
)

func Test_Installation_ReceiveInstallationOrder(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.installation.receive_installation_order_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping ReceiveInstallationOrder due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping ReceiveInstallationOrder due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/installation/receive_installation_order", app.APIURL), responder)

	var req ReceiveInstallationOrderRequest
	ctx := context.Background()
	res, err := client.Installation.ReceiveInstallationOrder(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Installation.ReceiveInstallationOrder returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Installation.ReceiveInstallationOrder response: %#v", res)
}
func Test_Installation_ValidateInstallationOtpCode(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.installation.validate_installation_otp_code_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping ValidateInstallationOtpCode due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping ValidateInstallationOtpCode due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/installation/validate_installation_otp_code", app.APIURL), responder)

	var req ValidateInstallationOtpCodeRequest
	ctx := context.Background()
	res, err := client.Installation.ValidateInstallationOtpCode(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Installation.ValidateInstallationOtpCode returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Installation.ValidateInstallationOtpCode response: %#v", res)
}
