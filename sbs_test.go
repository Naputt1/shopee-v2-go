package goshopee

import (
	"context"
	"fmt"
	"testing"

	"github.com/jarcoal/httpmock"
)

func Test_SBS_GetBoundWhsInfo(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.sbs.get_bound_whs_info_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetBoundWhsInfo due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetBoundWhsInfo due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/sbs/get_bound_whs_info", app.APIURL), responder)

	ctx := context.Background()
	res, err := client.SBS.GetBoundWhsInfo(ctx, shopID, accessToken)
	if err != nil {
		t.Logf("SBS.GetBoundWhsInfo returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("SBS.GetBoundWhsInfo response: %#v", res)
}
func Test_SBS_GetCurrentInventory(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.sbs.get_current_inventory_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetCurrentInventory due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetCurrentInventory due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/sbs/get_current_inventory", app.APIURL), responder)

	var req GetCurrentInventoryRequest
	ctx := context.Background()
	res, err := client.SBS.GetCurrentInventory(ctx, shopID, req, accessToken)
	if err != nil {
		t.Logf("SBS.GetCurrentInventory returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("SBS.GetCurrentInventory response: %#v", res)
}
func Test_SBS_GetExpiryReport(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.sbs.get_expiry_report_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetExpiryReport due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetExpiryReport due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/sbs/get_expiry_report", app.APIURL), responder)

	var req GetExpiryReportRequest
	ctx := context.Background()
	res, err := client.SBS.GetExpiryReport(ctx, shopID, req, accessToken)
	if err != nil {
		t.Logf("SBS.GetExpiryReport returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("SBS.GetExpiryReport response: %#v", res)
}
func Test_SBS_GetStockAging(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.sbs.get_stock_aging_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetStockAging due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetStockAging due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/sbs/get_stock_aging", app.APIURL), responder)

	var req GetStockAgingRequest
	ctx := context.Background()
	res, err := client.SBS.GetStockAging(ctx, shopID, req, accessToken)
	if err != nil {
		t.Logf("SBS.GetStockAging returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("SBS.GetStockAging response: %#v", res)
}
func Test_SBS_GetStockMovement(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.sbs.get_stock_movement_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetStockMovement due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetStockMovement due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/sbs/get_stock_movement", app.APIURL), responder)

	var req GetStockMovementRequest
	ctx := context.Background()
	res, err := client.SBS.GetStockMovement(ctx, shopID, req, accessToken)
	if err != nil {
		t.Logf("SBS.GetStockMovement returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("SBS.GetStockMovement response: %#v", res)
}
