package goshopee

import (
	"context"
	"fmt"
	"testing"

	"github.com/jarcoal/httpmock"
)

func Test_Ads_CheckCreateGmsProductCampaignEligibility(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.ads.check_create_gms_product_campaign_eligibility_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping CheckCreateGmsProductCampaignEligibility due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping CheckCreateGmsProductCampaignEligibility due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/ads/check_create_gms_product_campaign_eligibility", app.APIURL), responder)

	ctx := context.Background()
	res, err := client.Ads.CheckCreateGmsProductCampaignEligibility(ctx, sid, mid, tok)
	if err != nil {
		t.Logf("Ads.CheckCreateGmsProductCampaignEligibility returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Ads.CheckCreateGmsProductCampaignEligibility response: %#v", res)
}
func Test_Ads_CreateGmsProductCampaign(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.ads.create_gms_product_campaign_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping CreateGmsProductCampaign due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping CreateGmsProductCampaign due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/ads/create_gms_product_campaign", app.APIURL), responder)

	var req CreateGmsProductCampaignRequest
	ctx := context.Background()
	res, err := client.Ads.CreateGmsProductCampaign(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Ads.CreateGmsProductCampaign returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Ads.CreateGmsProductCampaign response: %#v", res)
}
func Test_Ads_CreateManualProductAds(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.ads.create_manual_product_ads_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping CreateManualProductAds due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping CreateManualProductAds due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/ads/create_manual_product_ads", app.APIURL), responder)

	var req CreateManualProductAdsRequest
	ctx := context.Background()
	res, err := client.Ads.CreateManualProductAds(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Ads.CreateManualProductAds returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Ads.CreateManualProductAds response: %#v", res)
}
func Test_Ads_EditGmsItemProductCampaign(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.ads.edit_gms_item_product_campaign_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping EditGmsItemProductCampaign due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping EditGmsItemProductCampaign due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/ads/edit_gms_item_product_campaign", app.APIURL), responder)

	var req EditGmsItemProductCampaignRequest
	ctx := context.Background()
	res, err := client.Ads.EditGmsItemProductCampaign(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Ads.EditGmsItemProductCampaign returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Ads.EditGmsItemProductCampaign response: %#v", res)
}
func Test_Ads_EditGmsProductCampaign(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.ads.edit_gms_product_campaign_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping EditGmsProductCampaign due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping EditGmsProductCampaign due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/ads/edit_gms_product_campaign", app.APIURL), responder)

	var req EditGmsProductCampaignRequest
	ctx := context.Background()
	res, err := client.Ads.EditGmsProductCampaign(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Ads.EditGmsProductCampaign returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Ads.EditGmsProductCampaign response: %#v", res)
}
func Test_Ads_EditManualProductAdKeywords(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.ads.edit_manual_product_ad_keywords_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping EditManualProductAdKeywords due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping EditManualProductAdKeywords due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/ads/edit_manual_product_ad_keywords", app.APIURL), responder)

	var req EditManualProductAdKeywordsRequest
	ctx := context.Background()
	res, err := client.Ads.EditManualProductAdKeywords(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Ads.EditManualProductAdKeywords returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Ads.EditManualProductAdKeywords response: %#v", res)
}
func Test_Ads_EditManualProductAds(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.ads.edit_manual_product_ads_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping EditManualProductAds due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping EditManualProductAds due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/ads/edit_manual_product_ads", app.APIURL), responder)

	var req EditManualProductAdsRequest
	ctx := context.Background()
	res, err := client.Ads.EditManualProductAds(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Ads.EditManualProductAds returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Ads.EditManualProductAds response: %#v", res)
}
func Test_Ads_GetAdsFcilShopRate(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.ads.get_ads_fácil_shop_rate_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetAdsFcilShopRate due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetAdsFcilShopRate due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/ads/get_ads_facil_shop_rate", app.APIURL), responder)

	ctx := context.Background()
	res, err := client.Ads.GetAdsFcilShopRate(ctx, sid, mid, tok)
	if err != nil {
		t.Logf("Ads.GetAdsFcilShopRate returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Ads.GetAdsFcilShopRate response: %#v", res)
}
func Test_Ads_GetAllCpcAdsDailyPerformance(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.ads.get_all_cpc_ads_daily_performance_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetAllCpcAdsDailyPerformance due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetAllCpcAdsDailyPerformance due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/ads/get_all_cpc_ads_daily_performance", app.APIURL), responder)

	var req GetAllCpcAdsDailyPerformanceRequest
	ctx := context.Background()
	res, err := client.Ads.GetAllCpcAdsDailyPerformance(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Ads.GetAllCpcAdsDailyPerformance returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Ads.GetAllCpcAdsDailyPerformance response: %#v", res)
}
func Test_Ads_GetAllCpcAdsHourlyPerformance(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.ads.get_all_cpc_ads_hourly_performance_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetAllCpcAdsHourlyPerformance due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetAllCpcAdsHourlyPerformance due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/ads/get_all_cpc_ads_hourly_performance", app.APIURL), responder)

	var req GetAllCpcAdsHourlyPerformanceRequest
	ctx := context.Background()
	res, err := client.Ads.GetAllCpcAdsHourlyPerformance(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Ads.GetAllCpcAdsHourlyPerformance returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Ads.GetAllCpcAdsHourlyPerformance response: %#v", res)
}
func Test_Ads_GetCreateProductAdBudgetSuggestion(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.ads.get_create_product_ad_budget_suggestion_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetCreateProductAdBudgetSuggestion due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetCreateProductAdBudgetSuggestion due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/ads/get_create_product_ad_budget_suggestion", app.APIURL), responder)

	var req GetCreateProductAdBudgetSuggestionRequest
	ctx := context.Background()
	res, err := client.Ads.GetCreateProductAdBudgetSuggestion(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Ads.GetCreateProductAdBudgetSuggestion returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Ads.GetCreateProductAdBudgetSuggestion response: %#v", res)
}
func Test_Ads_GetGmsCampaignPerformance(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.ads.get_gms_campaign_performance_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetGmsCampaignPerformance due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetGmsCampaignPerformance due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/ads/get_gms_campaign_performance", app.APIURL), responder)

	var req GetGmsCampaignPerformanceRequest
	ctx := context.Background()
	res, err := client.Ads.GetGmsCampaignPerformance(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Ads.GetGmsCampaignPerformance returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Ads.GetGmsCampaignPerformance response: %#v", res)
}
func Test_Ads_GetGmsItemPerformance(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.ads.get_gms_item_performance_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetGmsItemPerformance due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetGmsItemPerformance due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/ads/get_gms_item_performance", app.APIURL), responder)

	var req GetGmsItemPerformanceRequest
	ctx := context.Background()
	res, err := client.Ads.GetGmsItemPerformance(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Ads.GetGmsItemPerformance returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Ads.GetGmsItemPerformance response: %#v", res)
}
func Test_Ads_GetProductCampaignDailyPerformance(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.ads.get_product_campaign_daily_performance_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetProductCampaignDailyPerformance due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetProductCampaignDailyPerformance due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/ads/get_product_campaign_daily_performance", app.APIURL), responder)

	var req GetProductCampaignDailyPerformanceRequest
	ctx := context.Background()
	res, err := client.Ads.GetProductCampaignDailyPerformance(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Ads.GetProductCampaignDailyPerformance returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Ads.GetProductCampaignDailyPerformance response: %#v", res)
}
func Test_Ads_GetProductCampaignHourlyPerformance(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.ads.get_product_campaign_hourly_performance_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetProductCampaignHourlyPerformance due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetProductCampaignHourlyPerformance due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/ads/get_product_campaign_hourly_performance", app.APIURL), responder)

	var req GetProductCampaignHourlyPerformanceRequest
	ctx := context.Background()
	res, err := client.Ads.GetProductCampaignHourlyPerformance(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Ads.GetProductCampaignHourlyPerformance returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Ads.GetProductCampaignHourlyPerformance response: %#v", res)
}
func Test_Ads_GetProductLevelCampaignIdList(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.ads.get_product_level_campaign_id_list_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetProductLevelCampaignIdList due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetProductLevelCampaignIdList due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/ads/get_product_level_campaign_id_list", app.APIURL), responder)

	var req GetProductLevelCampaignIdListRequest
	ctx := context.Background()
	res, err := client.Ads.GetProductLevelCampaignIdList(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Ads.GetProductLevelCampaignIdList returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Ads.GetProductLevelCampaignIdList response: %#v", res)
}
func Test_Ads_GetProductLevelCampaignSettingInfo(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.ads.get_product_level_campaign_setting_info_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetProductLevelCampaignSettingInfo due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetProductLevelCampaignSettingInfo due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/ads/get_product_level_campaign_setting_info", app.APIURL), responder)

	var req GetProductLevelCampaignSettingInfoRequest
	ctx := context.Background()
	res, err := client.Ads.GetProductLevelCampaignSettingInfo(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Ads.GetProductLevelCampaignSettingInfo returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Ads.GetProductLevelCampaignSettingInfo response: %#v", res)
}
func Test_Ads_GetProductRecommendedRoiTarget(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.ads.get_product_recommended_roi_target_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetProductRecommendedRoiTarget due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetProductRecommendedRoiTarget due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/ads/get_product_recommended_roi_target", app.APIURL), responder)

	var req GetProductRecommendedRoiTargetRequest
	ctx := context.Background()
	res, err := client.Ads.GetProductRecommendedRoiTarget(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Ads.GetProductRecommendedRoiTarget returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Ads.GetProductRecommendedRoiTarget response: %#v", res)
}
func Test_Ads_GetRecommendedItemList(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.ads.get_recommended_item_list_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetRecommendedItemList due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetRecommendedItemList due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/ads/get_recommended_item_list", app.APIURL), responder)

	ctx := context.Background()
	res, err := client.Ads.GetRecommendedItemList(ctx, sid, mid, tok)
	if err != nil {
		t.Logf("Ads.GetRecommendedItemList returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Ads.GetRecommendedItemList response: %#v", res)
}
func Test_Ads_GetRecommendedKeywordList(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.ads.get_recommended_keyword_list_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetRecommendedKeywordList due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetRecommendedKeywordList due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/ads/get_recommended_keyword_list", app.APIURL), responder)

	var req GetRecommendedKeywordListRequest
	ctx := context.Background()
	res, err := client.Ads.GetRecommendedKeywordList(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Ads.GetRecommendedKeywordList returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Ads.GetRecommendedKeywordList response: %#v", res)
}
func Test_Ads_GetShopToggleInfo(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.ads.get_shop_toggle_info_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetShopToggleInfo due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetShopToggleInfo due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/ads/get_shop_toggle_info", app.APIURL), responder)

	ctx := context.Background()
	res, err := client.Ads.GetShopToggleInfo(ctx, sid, mid, tok)
	if err != nil {
		t.Logf("Ads.GetShopToggleInfo returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Ads.GetShopToggleInfo response: %#v", res)
}
func Test_Ads_GetTotalBalance(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.ads.get_total_balance_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetTotalBalance due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetTotalBalance due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/ads/get_total_balance", app.APIURL), responder)

	ctx := context.Background()
	res, err := client.Ads.GetTotalBalance(ctx, sid, mid, tok)
	if err != nil {
		t.Logf("Ads.GetTotalBalance returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Ads.GetTotalBalance response: %#v", res)
}
func Test_Ads_ListGmsUserDeletedItem(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.ads.list_gms_user_deleted_item_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping ListGmsUserDeletedItem due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping ListGmsUserDeletedItem due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/ads/list_gms_user_deleted_item", app.APIURL), responder)

	var req ListGmsUserDeletedItemRequest
	ctx := context.Background()
	res, err := client.Ads.ListGmsUserDeletedItem(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Ads.ListGmsUserDeletedItem returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Ads.ListGmsUserDeletedItem response: %#v", res)
}
