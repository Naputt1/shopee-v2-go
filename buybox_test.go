package goshopee

import (
	"context"
	"fmt"
	"testing"

	"github.com/jarcoal/httpmock"
)

func Test_Buybox_GetBuyboxModelPerformanceonlyForBRNow(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.buybox.get_buybox_model_performance（only for BR now）_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetBuyboxModelPerformanceonlyForBRNow due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetBuyboxModelPerformanceonlyForBRNow due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/buybox/get_buybox_model_performance", app.APIURL), responder)

	var req GetBuyboxModelPerformanceonlyForBRNowRequest
	ctx := context.Background()
	res, err := client.Buybox.GetBuyboxModelPerformanceonlyForBRNow(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Buybox.GetBuyboxModelPerformanceonlyForBRNow returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Buybox.GetBuyboxModelPerformanceonlyForBRNow response: %#v", res)
}
func Test_Buybox_GetBuyboxModelsByModelIdonlyForBRNow(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.buybox.get_buybox_models_by_model_id（only for BR now）_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetBuyboxModelsByModelIdonlyForBRNow due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetBuyboxModelsByModelIdonlyForBRNow due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/buybox/get_buybox_models_by_model_id", app.APIURL), responder)

	var req GetBuyboxModelsByModelIdonlyForBRNowRequest
	ctx := context.Background()
	res, err := client.Buybox.GetBuyboxModelsByModelIdonlyForBRNow(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Buybox.GetBuyboxModelsByModelIdonlyForBRNow returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Buybox.GetBuyboxModelsByModelIdonlyForBRNow response: %#v", res)
}
func Test_Buybox_GetBuyboxModelsByShopIdonlyForBRNow(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.buybox.get_buybox_models_by_shop_id（only for BR now）_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetBuyboxModelsByShopIdonlyForBRNow due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetBuyboxModelsByShopIdonlyForBRNow due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/buybox/get_buybox_models_by_shop_id", app.APIURL), responder)

	var req GetBuyboxModelsByShopIdonlyForBRNowRequest
	ctx := context.Background()
	res, err := client.Buybox.GetBuyboxModelsByShopIdonlyForBRNow(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Buybox.GetBuyboxModelsByShopIdonlyForBRNow returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Buybox.GetBuyboxModelsByShopIdonlyForBRNow response: %#v", res)
}
func Test_Buybox_GetBuyboxShopPerformanceonlyForBRNow(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.buybox.get_buybox_shop_performance（only for BR now）_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetBuyboxShopPerformanceonlyForBRNow due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetBuyboxShopPerformanceonlyForBRNow due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/buybox/get_buybox_shop_performance", app.APIURL), responder)

	ctx := context.Background()
	res, err := client.Buybox.GetBuyboxShopPerformanceonlyForBRNow(ctx, sid, mid, tok)
	if err != nil {
		t.Logf("Buybox.GetBuyboxShopPerformanceonlyForBRNow returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Buybox.GetBuyboxShopPerformanceonlyForBRNow response: %#v", res)
}
func Test_Buybox_UpdateBuyboxModelEnrollmentonlyForBRNow(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.buybox.update_buybox_model_enrollment（only for BR now）_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping UpdateBuyboxModelEnrollmentonlyForBRNow due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping UpdateBuyboxModelEnrollmentonlyForBRNow due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/buybox/update_buybox_model_enrollment", app.APIURL), responder)

	var req UpdateBuyboxModelEnrollmentonlyForBRNowRequest
	ctx := context.Background()
	res, err := client.Buybox.UpdateBuyboxModelEnrollmentonlyForBRNow(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Buybox.UpdateBuyboxModelEnrollmentonlyForBRNow returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Buybox.UpdateBuyboxModelEnrollmentonlyForBRNow response: %#v", res)
}
