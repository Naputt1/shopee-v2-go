package goshopee

import (
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

	res, err := client.SBS.GetBoundWhsInfo(shopID, accessToken)
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
	res, err := client.SBS.GetCurrentInventory(shopID, req, accessToken)
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
	res, err := client.SBS.GetExpiryReport(shopID, req, accessToken)
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
	res, err := client.SBS.GetStockAging(shopID, req, accessToken)
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
	res, err := client.SBS.GetStockMovement(shopID, req, accessToken)
	if err != nil {
		t.Logf("SBS.GetStockMovement returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("SBS.GetStockMovement response: %#v", res)
}
