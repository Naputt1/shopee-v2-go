package goshopee

import (
	"context"
	"fmt"
	"testing"

	"github.com/jarcoal/httpmock"
)

func Test_FBS_QueryBrShopBlockStatus(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.fbs.query_br_shop_block_status_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping QueryBrShopBlockStatus due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping QueryBrShopBlockStatus due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/fbs/query_br_shop_block_status", app.APIURL), responder)

	ctx := context.Background()
	res, err := client.FBS.QueryBrShopBlockStatus(ctx, shopID, accessToken)
	if err != nil {
		t.Logf("FBS.QueryBrShopBlockStatus returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("FBS.QueryBrShopBlockStatus response: %#v", res)
}
func Test_FBS_QueryBrShopEnrollmentStatus(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.fbs.query_br_shop_enrollment_status_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping QueryBrShopEnrollmentStatus due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping QueryBrShopEnrollmentStatus due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/fbs/query_br_shop_enrollment_status", app.APIURL), responder)

	ctx := context.Background()
	res, err := client.FBS.QueryBrShopEnrollmentStatus(ctx, shopID, accessToken)
	if err != nil {
		t.Logf("FBS.QueryBrShopEnrollmentStatus returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("FBS.QueryBrShopEnrollmentStatus response: %#v", res)
}
func Test_FBS_QueryBrShopInvoiceError(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.fbs.query_br_shop_invoice_error_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping QueryBrShopInvoiceError due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping QueryBrShopInvoiceError due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/fbs/query_br_shop_invoice_error", app.APIURL), responder)

	var req QueryBrShopInvoiceErrorRequest
	ctx := context.Background()
	res, err := client.FBS.QueryBrShopInvoiceError(ctx, shopID, req, accessToken)
	if err != nil {
		t.Logf("FBS.QueryBrShopInvoiceError returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("FBS.QueryBrShopInvoiceError response: %#v", res)
}
func Test_FBS_QueryBrSkuBlockStatus(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.fbs.query_br_sku_block_status_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping QueryBrSkuBlockStatus due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping QueryBrSkuBlockStatus due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/fbs/query_br_sku_block_status", app.APIURL), responder)

	var req QueryBrSkuBlockStatusRequest
	ctx := context.Background()
	res, err := client.FBS.QueryBrSkuBlockStatus(ctx, shopID, req, accessToken)
	if err != nil {
		t.Logf("FBS.QueryBrSkuBlockStatus returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("FBS.QueryBrSkuBlockStatus response: %#v", res)
}
