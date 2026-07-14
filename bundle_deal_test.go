package goshopee

import (
	"context"
	"fmt"
	"testing"

	"github.com/jarcoal/httpmock"
)

func Test_BundleDeal_AddBundleDeal(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.bundle_deal.add_bundle_deal_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping AddBundleDeal due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping AddBundleDeal due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/bundle_deal/add_bundle_deal", app.APIURL), responder)

	var req AddBundleDealRequest
	ctx := context.Background()
	res, err := client.BundleDeal.AddBundleDeal(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("BundleDeal.AddBundleDeal returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("BundleDeal.AddBundleDeal response: %#v", res)
}
func Test_BundleDeal_AddBundleDealItem(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.bundle_deal.add_bundle_deal_item_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping AddBundleDealItem due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping AddBundleDealItem due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/bundle_deal/add_bundle_deal_item", app.APIURL), responder)

	var req AddBundleDealItemRequest
	ctx := context.Background()
	res, err := client.BundleDeal.AddBundleDealItem(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("BundleDeal.AddBundleDealItem returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("BundleDeal.AddBundleDealItem response: %#v", res)
}
func Test_BundleDeal_DeleteBundleDeal(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.bundle_deal.delete_bundle_deal_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping DeleteBundleDeal due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping DeleteBundleDeal due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/bundle_deal/delete_bundle_deal", app.APIURL), responder)

	var req DeleteBundleDealRequest
	ctx := context.Background()
	res, err := client.BundleDeal.DeleteBundleDeal(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("BundleDeal.DeleteBundleDeal returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("BundleDeal.DeleteBundleDeal response: %#v", res)
}
func Test_BundleDeal_DeleteBundleDealItem(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.bundle_deal.delete_bundle_deal_item_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping DeleteBundleDealItem due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping DeleteBundleDealItem due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/bundle_deal/delete_bundle_deal_item", app.APIURL), responder)

	var req DeleteBundleDealItemRequest
	ctx := context.Background()
	res, err := client.BundleDeal.DeleteBundleDealItem(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("BundleDeal.DeleteBundleDealItem returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("BundleDeal.DeleteBundleDealItem response: %#v", res)
}
func Test_BundleDeal_EndBundleDeal(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.bundle_deal.end_bundle_deal_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping EndBundleDeal due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping EndBundleDeal due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/bundle_deal/end_bundle_deal", app.APIURL), responder)

	var req EndBundleDealRequest
	ctx := context.Background()
	res, err := client.BundleDeal.EndBundleDeal(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("BundleDeal.EndBundleDeal returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("BundleDeal.EndBundleDeal response: %#v", res)
}
func Test_BundleDeal_GetBundleDeal(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.bundle_deal.get_bundle_deal_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetBundleDeal due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetBundleDeal due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/bundle_deal/get_bundle_deal", app.APIURL), responder)

	var req GetBundleDealRequest
	ctx := context.Background()
	res, err := client.BundleDeal.GetBundleDeal(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("BundleDeal.GetBundleDeal returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("BundleDeal.GetBundleDeal response: %#v", res)
}
func Test_BundleDeal_GetBundleDealItem(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.bundle_deal.get_bundle_deal_item_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetBundleDealItem due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetBundleDealItem due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/bundle_deal/get_bundle_deal_item", app.APIURL), responder)

	var req GetBundleDealItemRequest
	ctx := context.Background()
	res, err := client.BundleDeal.GetBundleDealItem(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("BundleDeal.GetBundleDealItem returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("BundleDeal.GetBundleDealItem response: %#v", res)
}
func Test_BundleDeal_GetBundleDealList(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.bundle_deal.get_bundle_deal_list_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetBundleDealList due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetBundleDealList due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/bundle_deal/get_bundle_deal_list", app.APIURL), responder)

	var req GetBundleDealListRequest
	ctx := context.Background()
	res, err := client.BundleDeal.GetBundleDealList(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("BundleDeal.GetBundleDealList returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("BundleDeal.GetBundleDealList response: %#v", res)
}
func Test_BundleDeal_UpdateBundleDeal(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.bundle_deal.update_bundle_deal_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping UpdateBundleDeal due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping UpdateBundleDeal due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/bundle_deal/update_bundle_deal", app.APIURL), responder)

	var req UpdateBundleDealRequest
	ctx := context.Background()
	res, err := client.BundleDeal.UpdateBundleDeal(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("BundleDeal.UpdateBundleDeal returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("BundleDeal.UpdateBundleDeal response: %#v", res)
}
func Test_BundleDeal_UpdateBundleDealItem(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.bundle_deal.update_bundle_deal_item_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping UpdateBundleDealItem due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping UpdateBundleDealItem due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/bundle_deal/update_bundle_deal_item", app.APIURL), responder)

	var req UpdateBundleDealItemRequest
	ctx := context.Background()
	res, err := client.BundleDeal.UpdateBundleDealItem(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("BundleDeal.UpdateBundleDealItem returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("BundleDeal.UpdateBundleDealItem response: %#v", res)
}
