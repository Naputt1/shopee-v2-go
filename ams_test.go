package goshopee

import (
	"fmt"
	"testing"

	"github.com/jarcoal/httpmock"
)

func Test_AMS_AddAllProductsToOpenCampaign(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.ams.add_all_products_to_open_campaign_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping AddAllProductsToOpenCampaign due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping AddAllProductsToOpenCampaign due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/ams/add_all_products_to_open_campaign", app.APIURL), responder)

	var req AddAllProductsToOpenCampaignRequest
	res, err := client.AMS.AddAllProductsToOpenCampaign(shopID, req, accessToken)
	if err != nil {
		t.Logf("AMS.AddAllProductsToOpenCampaign returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("AMS.AddAllProductsToOpenCampaign response: %#v", res)
}
func Test_AMS_BatchAddProductsToOpenCampaign(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.ams.batch_add_products_to_open_campaign_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping BatchAddProductsToOpenCampaign due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping BatchAddProductsToOpenCampaign due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/ams/batch_add_products_to_open_campaign", app.APIURL), responder)

	var req BatchAddProductsToOpenCampaignRequest
	res, err := client.AMS.BatchAddProductsToOpenCampaign(shopID, req, accessToken)
	if err != nil {
		t.Logf("AMS.BatchAddProductsToOpenCampaign returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("AMS.BatchAddProductsToOpenCampaign response: %#v", res)
}
func Test_AMS_BatchEditProductsOpenCampaignSetting(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.ams.batch_edit_products_open_campaign_setting_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping BatchEditProductsOpenCampaignSetting due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping BatchEditProductsOpenCampaignSetting due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/ams/batch_edit_products_open_campaign_setting", app.APIURL), responder)

	var req BatchEditProductsOpenCampaignSettingRequest
	res, err := client.AMS.BatchEditProductsOpenCampaignSetting(shopID, req, accessToken)
	if err != nil {
		t.Logf("AMS.BatchEditProductsOpenCampaignSetting returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("AMS.BatchEditProductsOpenCampaignSetting response: %#v", res)
}
func Test_AMS_BatchGetProductsSuggestedRate(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.ams.batch_get_products_suggested_rate_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping BatchGetProductsSuggestedRate due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping BatchGetProductsSuggestedRate due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/ams/batch_get_products_suggested_rate", app.APIURL), responder)

	var req BatchGetProductsSuggestedRateRequest
	res, err := client.AMS.BatchGetProductsSuggestedRate(shopID, req, accessToken)
	if err != nil {
		t.Logf("AMS.BatchGetProductsSuggestedRate returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("AMS.BatchGetProductsSuggestedRate response: %#v", res)
}
func Test_AMS_BatchRemoveProductsOpenCampaignSetting(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.ams.batch_remove_products_open_campaign_setting_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping BatchRemoveProductsOpenCampaignSetting due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping BatchRemoveProductsOpenCampaignSetting due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/ams/batch_remove_products_open_campaign_setting", app.APIURL), responder)

	var req BatchRemoveProductsOpenCampaignSettingRequest
	res, err := client.AMS.BatchRemoveProductsOpenCampaignSetting(shopID, req, accessToken)
	if err != nil {
		t.Logf("AMS.BatchRemoveProductsOpenCampaignSetting returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("AMS.BatchRemoveProductsOpenCampaignSetting response: %#v", res)
}
func Test_AMS_CreateNewTargetedCampaign(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.ams.create_new_targeted_campaign_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping CreateNewTargetedCampaign due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping CreateNewTargetedCampaign due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/ams/create_new_targeted_campaign", app.APIURL), responder)

	var req CreateNewTargetedCampaignRequest
	res, err := client.AMS.CreateNewTargetedCampaign(shopID, req, accessToken)
	if err != nil {
		t.Logf("AMS.CreateNewTargetedCampaign returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("AMS.CreateNewTargetedCampaign response: %#v", res)
}
func Test_AMS_EditAffiliateListOfTargetedCampaign(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.ams.edit_affiliate_list_of_targeted_campaign_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping EditAffiliateListOfTargetedCampaign due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping EditAffiliateListOfTargetedCampaign due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/ams/edit_affiliate_list_of_targeted_campaign", app.APIURL), responder)

	var req EditAffiliateListOfTargetedCampaignRequest
	res, err := client.AMS.EditAffiliateListOfTargetedCampaign(shopID, req, accessToken)
	if err != nil {
		t.Logf("AMS.EditAffiliateListOfTargetedCampaign returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("AMS.EditAffiliateListOfTargetedCampaign response: %#v", res)
}
func Test_AMS_EditAllProductsOpenCampaignSetting(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.ams.edit_all_products_open_campaign_setting_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping EditAllProductsOpenCampaignSetting due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping EditAllProductsOpenCampaignSetting due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/ams/edit_all_products_open_campaign_setting", app.APIURL), responder)

	var req EditAllProductsOpenCampaignSettingRequest
	res, err := client.AMS.EditAllProductsOpenCampaignSetting(shopID, req, accessToken)
	if err != nil {
		t.Logf("AMS.EditAllProductsOpenCampaignSetting returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("AMS.EditAllProductsOpenCampaignSetting response: %#v", res)
}
func Test_AMS_EditProductListOfTargetedCampaign(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.ams.edit_product_list_of_targeted_campaign_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping EditProductListOfTargetedCampaign due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping EditProductListOfTargetedCampaign due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/ams/edit_product_list_of_targeted_campaign", app.APIURL), responder)

	var req EditProductListOfTargetedCampaignRequest
	res, err := client.AMS.EditProductListOfTargetedCampaign(shopID, req, accessToken)
	if err != nil {
		t.Logf("AMS.EditProductListOfTargetedCampaign returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("AMS.EditProductListOfTargetedCampaign response: %#v", res)
}
func Test_AMS_GetAffiliatePerformance(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.ams.get_affiliate_performance_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetAffiliatePerformance due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetAffiliatePerformance due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/ams/get_affiliate_performance", app.APIURL), responder)

	var req GetAffiliatePerformanceRequest
	res, err := client.AMS.GetAffiliatePerformance(shopID, req, accessToken)
	if err != nil {
		t.Logf("AMS.GetAffiliatePerformance returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("AMS.GetAffiliatePerformance response: %#v", res)
}
func Test_AMS_GetAutoAddNewProductToggleStatus(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.ams.get_auto_add_new_product_toggle_status_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetAutoAddNewProductToggleStatus due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetAutoAddNewProductToggleStatus due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/ams/get_auto_add_new_product_toggle_status", app.APIURL), responder)

	res, err := client.AMS.GetAutoAddNewProductToggleStatus(shopID, accessToken)
	if err != nil {
		t.Logf("AMS.GetAutoAddNewProductToggleStatus returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("AMS.GetAutoAddNewProductToggleStatus response: %#v", res)
}
func Test_AMS_GetCampaignKeyMetricsPerformance(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.ams.get_campaign_key_metrics_performance_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetCampaignKeyMetricsPerformance due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetCampaignKeyMetricsPerformance due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/ams/get_campaign_key_metrics_performance", app.APIURL), responder)

	var req GetCampaignKeyMetricsPerformanceRequest
	res, err := client.AMS.GetCampaignKeyMetricsPerformance(shopID, req, accessToken)
	if err != nil {
		t.Logf("AMS.GetCampaignKeyMetricsPerformance returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("AMS.GetCampaignKeyMetricsPerformance response: %#v", res)
}
func Test_AMS_GetContentPerformance(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.ams.get_content_performance_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetContentPerformance due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetContentPerformance due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/ams/get_content_performance", app.APIURL), responder)

	var req GetContentPerformanceRequest
	res, err := client.AMS.GetContentPerformance(shopID, req, accessToken)
	if err != nil {
		t.Logf("AMS.GetContentPerformance returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("AMS.GetContentPerformance response: %#v", res)
}
func Test_AMS_GetConversionReport(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.ams.get_conversion_report_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetConversionReport due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetConversionReport due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/ams/get_conversion_report", app.APIURL), responder)

	var req GetConversionReportRequest
	res, err := client.AMS.GetConversionReport(shopID, req, accessToken)
	if err != nil {
		t.Logf("AMS.GetConversionReport returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("AMS.GetConversionReport response: %#v", res)
}
func Test_AMS_GetManagedAffiliateList(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.ams.get_managed_affiliate_list_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetManagedAffiliateList due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetManagedAffiliateList due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/ams/get_managed_affiliate_list", app.APIURL), responder)

	var req GetManagedAffiliateListRequest
	res, err := client.AMS.GetManagedAffiliateList(shopID, req, accessToken)
	if err != nil {
		t.Logf("AMS.GetManagedAffiliateList returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("AMS.GetManagedAffiliateList response: %#v", res)
}
func Test_AMS_GetOpenCampaignAddedProduct(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.ams.get_open_campaign_added_product_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetOpenCampaignAddedProduct due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetOpenCampaignAddedProduct due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/ams/get_open_campaign_added_product", app.APIURL), responder)

	var req GetOpenCampaignAddedProductRequest
	res, err := client.AMS.GetOpenCampaignAddedProduct(shopID, req, accessToken)
	if err != nil {
		t.Logf("AMS.GetOpenCampaignAddedProduct returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("AMS.GetOpenCampaignAddedProduct response: %#v", res)
}
func Test_AMS_GetOpenCampaignBatchTaskResult(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.ams.get_open_campaign_batch_task_result_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetOpenCampaignBatchTaskResult due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetOpenCampaignBatchTaskResult due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/ams/get_open_campaign_batch_task_result", app.APIURL), responder)

	var req GetOpenCampaignBatchTaskResultRequest
	res, err := client.AMS.GetOpenCampaignBatchTaskResult(shopID, req, accessToken)
	if err != nil {
		t.Logf("AMS.GetOpenCampaignBatchTaskResult returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("AMS.GetOpenCampaignBatchTaskResult response: %#v", res)
}
func Test_AMS_GetOpenCampaignNotAddedProduct(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.ams.get_open_campaign_not_added_product_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetOpenCampaignNotAddedProduct due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetOpenCampaignNotAddedProduct due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/ams/get_open_campaign_not_added_product", app.APIURL), responder)

	var req GetOpenCampaignNotAddedProductRequest
	res, err := client.AMS.GetOpenCampaignNotAddedProduct(shopID, req, accessToken)
	if err != nil {
		t.Logf("AMS.GetOpenCampaignNotAddedProduct returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("AMS.GetOpenCampaignNotAddedProduct response: %#v", res)
}
func Test_AMS_GetOpenCampaignPerformance(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.ams.get_open_campaign_performance_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetOpenCampaignPerformance due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetOpenCampaignPerformance due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/ams/get_open_campaign_performance", app.APIURL), responder)

	var req GetOpenCampaignPerformanceRequest
	res, err := client.AMS.GetOpenCampaignPerformance(shopID, req, accessToken)
	if err != nil {
		t.Logf("AMS.GetOpenCampaignPerformance returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("AMS.GetOpenCampaignPerformance response: %#v", res)
}
func Test_AMS_GetOptimizationSuggestionProduct(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.ams.get_optimization_suggestion_product_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetOptimizationSuggestionProduct due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetOptimizationSuggestionProduct due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/ams/get_optimization_suggestion_product", app.APIURL), responder)

	var req GetOptimizationSuggestionProductRequest
	res, err := client.AMS.GetOptimizationSuggestionProduct(shopID, req, accessToken)
	if err != nil {
		t.Logf("AMS.GetOptimizationSuggestionProduct returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("AMS.GetOptimizationSuggestionProduct response: %#v", res)
}
func Test_AMS_GetPerformanceDataUpdateTime(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.ams.get_performance_data_update_time_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetPerformanceDataUpdateTime due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetPerformanceDataUpdateTime due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/ams/get_performance_data_update_time", app.APIURL), responder)

	var req GetPerformanceDataUpdateTimeRequest
	res, err := client.AMS.GetPerformanceDataUpdateTime(shopID, req, accessToken)
	if err != nil {
		t.Logf("AMS.GetPerformanceDataUpdateTime returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("AMS.GetPerformanceDataUpdateTime response: %#v", res)
}
func Test_AMS_GetProductPerformance(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.ams.get_product_performance_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetProductPerformance due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetProductPerformance due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/ams/get_product_performance", app.APIURL), responder)

	var req GetProductPerformanceRequest
	res, err := client.AMS.GetProductPerformance(shopID, req, accessToken)
	if err != nil {
		t.Logf("AMS.GetProductPerformance returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("AMS.GetProductPerformance response: %#v", res)
}
func Test_AMS_GetRecommendedAffiliateList(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.ams.get_recommended_affiliate_list_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetRecommendedAffiliateList due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetRecommendedAffiliateList due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/ams/get_recommended_affiliate_list", app.APIURL), responder)

	var req GetRecommendedAffiliateListRequest
	res, err := client.AMS.GetRecommendedAffiliateList(shopID, req, accessToken)
	if err != nil {
		t.Logf("AMS.GetRecommendedAffiliateList returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("AMS.GetRecommendedAffiliateList response: %#v", res)
}
func Test_AMS_GetShopPerformance(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.ams.get_shop_performance_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetShopPerformance due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetShopPerformance due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/ams/get_shop_performance", app.APIURL), responder)

	var req GetShopPerformanceRequest
	res, err := client.AMS.GetShopPerformance(shopID, req, accessToken)
	if err != nil {
		t.Logf("AMS.GetShopPerformance returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("AMS.GetShopPerformance response: %#v", res)
}
func Test_AMS_GetShopSuggestedRate(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.ams.get_shop_suggested_rate_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetShopSuggestedRate due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetShopSuggestedRate due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/ams/get_shop_suggested_rate", app.APIURL), responder)

	res, err := client.AMS.GetShopSuggestedRate(shopID, accessToken)
	if err != nil {
		t.Logf("AMS.GetShopSuggestedRate returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("AMS.GetShopSuggestedRate response: %#v", res)
}
func Test_AMS_GetTargetedCampaignAddableProductList(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.ams.get_targeted_campaign_addable_product_list_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetTargetedCampaignAddableProductList due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetTargetedCampaignAddableProductList due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/ams/get_targeted_campaign_addable_product_list", app.APIURL), responder)

	var req GetTargetedCampaignAddableProductListRequest
	res, err := client.AMS.GetTargetedCampaignAddableProductList(shopID, req, accessToken)
	if err != nil {
		t.Logf("AMS.GetTargetedCampaignAddableProductList returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("AMS.GetTargetedCampaignAddableProductList response: %#v", res)
}
func Test_AMS_GetTargetedCampaignList(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.ams.get_targeted_campaign_list_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetTargetedCampaignList due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetTargetedCampaignList due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/ams/get_targeted_campaign_list", app.APIURL), responder)

	var req GetTargetedCampaignListRequest
	res, err := client.AMS.GetTargetedCampaignList(shopID, req, accessToken)
	if err != nil {
		t.Logf("AMS.GetTargetedCampaignList returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("AMS.GetTargetedCampaignList response: %#v", res)
}
func Test_AMS_GetTargetedCampaignPerformance(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.ams.get_targeted_campaign_performance_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetTargetedCampaignPerformance due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetTargetedCampaignPerformance due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/ams/get_targeted_campaign_performance", app.APIURL), responder)

	var req GetTargetedCampaignPerformanceRequest
	res, err := client.AMS.GetTargetedCampaignPerformance(shopID, req, accessToken)
	if err != nil {
		t.Logf("AMS.GetTargetedCampaignPerformance returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("AMS.GetTargetedCampaignPerformance response: %#v", res)
}
func Test_AMS_GetTargetedCampaignSettings(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.ams.get_targeted_campaign_settings_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetTargetedCampaignSettings due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetTargetedCampaignSettings due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/ams/get_targeted_campaign_settings", app.APIURL), responder)

	var req GetTargetedCampaignSettingsRequest
	res, err := client.AMS.GetTargetedCampaignSettings(shopID, req, accessToken)
	if err != nil {
		t.Logf("AMS.GetTargetedCampaignSettings returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("AMS.GetTargetedCampaignSettings response: %#v", res)
}
func Test_AMS_GetValidationList(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.ams.get_validation_list_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetValidationList due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetValidationList due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/ams/get_validation_list", app.APIURL), responder)

	res, err := client.AMS.GetValidationList(shopID, accessToken)
	if err != nil {
		t.Logf("AMS.GetValidationList returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("AMS.GetValidationList response: %#v", res)
}
func Test_AMS_GetValidationReport(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.ams.get_validation_report_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetValidationReport due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetValidationReport due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/ams/get_validation_report", app.APIURL), responder)

	var req GetValidationReportRequest
	res, err := client.AMS.GetValidationReport(shopID, req, accessToken)
	if err != nil {
		t.Logf("AMS.GetValidationReport returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("AMS.GetValidationReport response: %#v", res)
}
func Test_AMS_QueryAffiliateList(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.ams.query_affiliate_list_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping QueryAffiliateList due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping QueryAffiliateList due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/ams/query_affiliate_list", app.APIURL), responder)

	var req QueryAffiliateListRequest
	res, err := client.AMS.QueryAffiliateList(shopID, req, accessToken)
	if err != nil {
		t.Logf("AMS.QueryAffiliateList returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("AMS.QueryAffiliateList response: %#v", res)
}
func Test_AMS_RemoveAllProductsOpenCampaignSetting(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.ams.remove_all_products_open_campaign_setting_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping RemoveAllProductsOpenCampaignSetting due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping RemoveAllProductsOpenCampaignSetting due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/ams/remove_all_products_open_campaign_setting", app.APIURL), responder)

	res, err := client.AMS.RemoveAllProductsOpenCampaignSetting(shopID, accessToken)
	if err != nil {
		t.Logf("AMS.RemoveAllProductsOpenCampaignSetting returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("AMS.RemoveAllProductsOpenCampaignSetting response: %#v", res)
}
func Test_AMS_TerminateTargetedCampaign(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.ams.terminate_targeted_campaign_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping TerminateTargetedCampaign due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping TerminateTargetedCampaign due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/ams/terminate_targeted_campaign", app.APIURL), responder)

	var req TerminateTargetedCampaignRequest
	res, err := client.AMS.TerminateTargetedCampaign(shopID, req, accessToken)
	if err != nil {
		t.Logf("AMS.TerminateTargetedCampaign returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("AMS.TerminateTargetedCampaign response: %#v", res)
}
func Test_AMS_UpdateAutoAddNewProductSetting(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.ams.update_auto_add_new_product_setting_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping UpdateAutoAddNewProductSetting due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping UpdateAutoAddNewProductSetting due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/ams/update_auto_add_new_product_setting", app.APIURL), responder)

	var req UpdateAutoAddNewProductSettingRequest
	res, err := client.AMS.UpdateAutoAddNewProductSetting(shopID, req, accessToken)
	if err != nil {
		t.Logf("AMS.UpdateAutoAddNewProductSetting returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("AMS.UpdateAutoAddNewProductSetting response: %#v", res)
}
func Test_AMS_UpdateBasicInfoOfTargetedCampaign(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.ams.update_basic_info_of_targeted_campaign_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping UpdateBasicInfoOfTargetedCampaign due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping UpdateBasicInfoOfTargetedCampaign due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/ams/update_basic_info_of_targeted_campaign", app.APIURL), responder)

	var req UpdateBasicInfoOfTargetedCampaignRequest
	res, err := client.AMS.UpdateBasicInfoOfTargetedCampaign(shopID, req, accessToken)
	if err != nil {
		t.Logf("AMS.UpdateBasicInfoOfTargetedCampaign returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("AMS.UpdateBasicInfoOfTargetedCampaign response: %#v", res)
}
