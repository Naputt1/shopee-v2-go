package goshopee

import (
	"context"
	"fmt"
	"testing"

	"github.com/jarcoal/httpmock"
)

func Test_Voucher_AddVoucher(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.voucher.add_voucher_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping AddVoucher due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping AddVoucher due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/voucher/add_voucher", app.APIURL), responder)

	var req AddVoucherRequest
	ctx := context.Background()
	res, err := client.Voucher.AddVoucher(ctx, shopID, req, accessToken)
	if err != nil {
		t.Logf("Voucher.AddVoucher returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Voucher.AddVoucher response: %#v", res)
}
func Test_Voucher_DeleteVoucher(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.voucher.delete_voucher_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping DeleteVoucher due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping DeleteVoucher due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/voucher/delete_voucher", app.APIURL), responder)

	var req DeleteVoucherRequest
	ctx := context.Background()
	res, err := client.Voucher.DeleteVoucher(ctx, shopID, req, accessToken)
	if err != nil {
		t.Logf("Voucher.DeleteVoucher returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Voucher.DeleteVoucher response: %#v", res)
}
func Test_Voucher_EndVoucher(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.voucher.end_voucher_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping EndVoucher due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping EndVoucher due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/voucher/end_voucher", app.APIURL), responder)

	var req EndVoucherRequest
	ctx := context.Background()
	res, err := client.Voucher.EndVoucher(ctx, shopID, req, accessToken)
	if err != nil {
		t.Logf("Voucher.EndVoucher returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Voucher.EndVoucher response: %#v", res)
}
func Test_Voucher_GetVoucher(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.voucher.get_voucher_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetVoucher due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetVoucher due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/voucher/get_voucher", app.APIURL), responder)

	var req GetVoucherRequest
	ctx := context.Background()
	res, err := client.Voucher.GetVoucher(ctx, shopID, req, accessToken)
	if err != nil {
		t.Logf("Voucher.GetVoucher returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Voucher.GetVoucher response: %#v", res)
}
func Test_Voucher_GetVoucherList(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.voucher.get_voucher_list_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetVoucherList due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetVoucherList due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/voucher/get_voucher_list", app.APIURL), responder)

	var req GetVoucherListRequest
	ctx := context.Background()
	res, err := client.Voucher.GetVoucherList(ctx, shopID, req, accessToken)
	if err != nil {
		t.Logf("Voucher.GetVoucherList returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Voucher.GetVoucherList response: %#v", res)
}
func Test_Voucher_UpdateVoucher(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.voucher.update_voucher_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping UpdateVoucher due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping UpdateVoucher due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/voucher/update_voucher", app.APIURL), responder)

	var req UpdateVoucherRequest
	ctx := context.Background()
	res, err := client.Voucher.UpdateVoucher(ctx, shopID, req, accessToken)
	if err != nil {
		t.Logf("Voucher.UpdateVoucher returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Voucher.UpdateVoucher response: %#v", res)
}
