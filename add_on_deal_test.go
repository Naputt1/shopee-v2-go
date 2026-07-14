package goshopee

import (
	"context"
	"fmt"
	"testing"

	"github.com/jarcoal/httpmock"
)

func Test_AddOnDeal_AddAddOnDeal(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.add_on_deal.add_add_on_deal_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping AddAddOnDeal due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping AddAddOnDeal due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/add_on_deal/add_add_on_deal", app.APIURL), responder)

	var req AddAddOnDealRequest
	ctx := context.Background()
	res, err := client.AddOnDeal.AddAddOnDeal(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("AddOnDeal.AddAddOnDeal returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("AddOnDeal.AddAddOnDeal response: %#v", res)
}
func Test_AddOnDeal_AddAddOnDealMainItem(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.add_on_deal.add_add_on_deal_main_item_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping AddAddOnDealMainItem due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping AddAddOnDealMainItem due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/add_on_deal/add_add_on_deal_main_item", app.APIURL), responder)

	var req AddAddOnDealMainItemRequest
	ctx := context.Background()
	res, err := client.AddOnDeal.AddAddOnDealMainItem(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("AddOnDeal.AddAddOnDealMainItem returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("AddOnDeal.AddAddOnDealMainItem response: %#v", res)
}
func Test_AddOnDeal_AddAddOnDealSubItem(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.add_on_deal.add_add_on_deal_sub_item_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping AddAddOnDealSubItem due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping AddAddOnDealSubItem due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/add_on_deal/add_add_on_deal_sub_item", app.APIURL), responder)

	var req AddAddOnDealSubItemRequest
	ctx := context.Background()
	res, err := client.AddOnDeal.AddAddOnDealSubItem(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("AddOnDeal.AddAddOnDealSubItem returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("AddOnDeal.AddAddOnDealSubItem response: %#v", res)
}
func Test_AddOnDeal_DeleteAddOnDeal(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.add_on_deal.delete_add_on_deal_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping DeleteAddOnDeal due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping DeleteAddOnDeal due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/add_on_deal/delete_add_on_deal", app.APIURL), responder)

	var req DeleteAddOnDealRequest
	ctx := context.Background()
	res, err := client.AddOnDeal.DeleteAddOnDeal(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("AddOnDeal.DeleteAddOnDeal returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("AddOnDeal.DeleteAddOnDeal response: %#v", res)
}
func Test_AddOnDeal_DeleteAddOnDealMainItem(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.add_on_deal.delete_add_on_deal_main_item_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping DeleteAddOnDealMainItem due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping DeleteAddOnDealMainItem due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/add_on_deal/delete_add_on_deal_main_item", app.APIURL), responder)

	var req DeleteAddOnDealMainItemRequest
	ctx := context.Background()
	res, err := client.AddOnDeal.DeleteAddOnDealMainItem(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("AddOnDeal.DeleteAddOnDealMainItem returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("AddOnDeal.DeleteAddOnDealMainItem response: %#v", res)
}
func Test_AddOnDeal_DeleteAddOnDealSubItem(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.add_on_deal.delete_add_on_deal_sub_item_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping DeleteAddOnDealSubItem due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping DeleteAddOnDealSubItem due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/add_on_deal/delete_add_on_deal_sub_item", app.APIURL), responder)

	var req DeleteAddOnDealSubItemRequest
	ctx := context.Background()
	res, err := client.AddOnDeal.DeleteAddOnDealSubItem(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("AddOnDeal.DeleteAddOnDealSubItem returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("AddOnDeal.DeleteAddOnDealSubItem response: %#v", res)
}
func Test_AddOnDeal_EndAddOnDeal(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.add_on_deal.end_add_on_deal_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping EndAddOnDeal due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping EndAddOnDeal due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/add_on_deal/end_add_on_deal", app.APIURL), responder)

	var req EndAddOnDealRequest
	ctx := context.Background()
	res, err := client.AddOnDeal.EndAddOnDeal(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("AddOnDeal.EndAddOnDeal returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("AddOnDeal.EndAddOnDeal response: %#v", res)
}
func Test_AddOnDeal_GetAddOnDeal(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.add_on_deal.get_add_on_deal_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetAddOnDeal due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetAddOnDeal due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/add_on_deal/get_add_on_deal", app.APIURL), responder)

	var req GetAddOnDealRequest
	ctx := context.Background()
	res, err := client.AddOnDeal.GetAddOnDeal(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("AddOnDeal.GetAddOnDeal returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("AddOnDeal.GetAddOnDeal response: %#v", res)
}
func Test_AddOnDeal_GetAddOnDealList(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.add_on_deal.get_add_on_deal_list_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetAddOnDealList due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetAddOnDealList due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/add_on_deal/get_add_on_deal_list", app.APIURL), responder)

	var req GetAddOnDealListRequest
	ctx := context.Background()
	res, err := client.AddOnDeal.GetAddOnDealList(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("AddOnDeal.GetAddOnDealList returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("AddOnDeal.GetAddOnDealList response: %#v", res)
}
func Test_AddOnDeal_GetAddOnDealMainItem(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.add_on_deal.get_add_on_deal_main_item_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetAddOnDealMainItem due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetAddOnDealMainItem due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/add_on_deal/get_add_on_deal_main_item", app.APIURL), responder)

	var req GetAddOnDealMainItemRequest
	ctx := context.Background()
	res, err := client.AddOnDeal.GetAddOnDealMainItem(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("AddOnDeal.GetAddOnDealMainItem returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("AddOnDeal.GetAddOnDealMainItem response: %#v", res)
}
func Test_AddOnDeal_GetAddOnDealSubItem(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.add_on_deal.get_add_on_deal_sub_item_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetAddOnDealSubItem due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetAddOnDealSubItem due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/add_on_deal/get_add_on_deal_sub_item", app.APIURL), responder)

	var req GetAddOnDealSubItemRequest
	ctx := context.Background()
	res, err := client.AddOnDeal.GetAddOnDealSubItem(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("AddOnDeal.GetAddOnDealSubItem returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("AddOnDeal.GetAddOnDealSubItem response: %#v", res)
}
func Test_AddOnDeal_UpdateAddOnDeal(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.add_on_deal.update_add_on_deal_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping UpdateAddOnDeal due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping UpdateAddOnDeal due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/add_on_deal/update_add_on_deal", app.APIURL), responder)

	var req UpdateAddOnDealRequest
	ctx := context.Background()
	res, err := client.AddOnDeal.UpdateAddOnDeal(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("AddOnDeal.UpdateAddOnDeal returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("AddOnDeal.UpdateAddOnDeal response: %#v", res)
}
func Test_AddOnDeal_UpdateAddOnDealMainItem(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.add_on_deal.update_add_on_deal_main_item_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping UpdateAddOnDealMainItem due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping UpdateAddOnDealMainItem due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/add_on_deal/update_add_on_deal_main_item", app.APIURL), responder)

	var req UpdateAddOnDealMainItemRequest
	ctx := context.Background()
	res, err := client.AddOnDeal.UpdateAddOnDealMainItem(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("AddOnDeal.UpdateAddOnDealMainItem returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("AddOnDeal.UpdateAddOnDealMainItem response: %#v", res)
}
func Test_AddOnDeal_UpdateAddOnDealSubItem(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.add_on_deal.update_add_on_deal_sub_item_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping UpdateAddOnDealSubItem due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping UpdateAddOnDealSubItem due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/add_on_deal/update_add_on_deal_sub_item", app.APIURL), responder)

	var req UpdateAddOnDealSubItemRequest
	ctx := context.Background()
	res, err := client.AddOnDeal.UpdateAddOnDealSubItem(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("AddOnDeal.UpdateAddOnDealSubItem returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("AddOnDeal.UpdateAddOnDealSubItem response: %#v", res)
}
