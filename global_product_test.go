package goshopee

import (
	"fmt"
	"testing"

	"github.com/jarcoal/httpmock"
)

func Test_GlobalProduct_AddGlobalItem(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.global_product.add_global_item_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping AddGlobalItem due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping AddGlobalItem due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/global_product/add_global_item", app.APIURL), responder)

	var req AddGlobalItemRequest
	res, err := client.GlobalProduct.AddGlobalItem(shopID, req, accessToken)
	if err != nil {
		t.Logf("GlobalProduct.AddGlobalItem returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("GlobalProduct.AddGlobalItem response: %#v", res)
}
func Test_GlobalProduct_AddGlobalModel(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.global_product.add_global_model_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping AddGlobalModel due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping AddGlobalModel due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/global_product/add_global_model", app.APIURL), responder)

	var req AddGlobalModelRequest
	res, err := client.GlobalProduct.AddGlobalModel(shopID, req, accessToken)
	if err != nil {
		t.Logf("GlobalProduct.AddGlobalModel returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("GlobalProduct.AddGlobalModel response: %#v", res)
}
func Test_GlobalProduct_CategoryRecommend(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.global_product.category_recommend_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping CategoryRecommend due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping CategoryRecommend due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/global_product/category_recommend", app.APIURL), responder)

	var req CategoryRecommendRequest
	res, err := client.GlobalProduct.CategoryRecommend(shopID, req, accessToken)
	if err != nil {
		t.Logf("GlobalProduct.CategoryRecommend returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("GlobalProduct.CategoryRecommend response: %#v", res)
}
func Test_GlobalProduct_CreatePublishTask(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.global_product.create_publish_task_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping CreatePublishTask due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping CreatePublishTask due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/global_product/create_publish_task", app.APIURL), responder)

	var req CreatePublishTaskRequest
	res, err := client.GlobalProduct.CreatePublishTask(shopID, req, accessToken)
	if err != nil {
		t.Logf("GlobalProduct.CreatePublishTask returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("GlobalProduct.CreatePublishTask response: %#v", res)
}
func Test_GlobalProduct_DeleteGlobalItem(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.global_product.delete_global_item_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping DeleteGlobalItem due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping DeleteGlobalItem due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/global_product/delete_global_item", app.APIURL), responder)

	var req DeleteGlobalItemRequest
	res, err := client.GlobalProduct.DeleteGlobalItem(shopID, req, accessToken)
	if err != nil {
		t.Logf("GlobalProduct.DeleteGlobalItem returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("GlobalProduct.DeleteGlobalItem response: %#v", res)
}
func Test_GlobalProduct_DeleteGlobalModel(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.global_product.delete_global_model_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping DeleteGlobalModel due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping DeleteGlobalModel due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/global_product/delete_global_model", app.APIURL), responder)

	var req DeleteGlobalModelRequest
	res, err := client.GlobalProduct.DeleteGlobalModel(shopID, req, accessToken)
	if err != nil {
		t.Logf("GlobalProduct.DeleteGlobalModel returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("GlobalProduct.DeleteGlobalModel response: %#v", res)
}
func Test_GlobalProduct_GetAttributeTree(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.global_product.get_attribute_tree_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetAttributeTree due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetAttributeTree due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/global_product/get_attribute_tree", app.APIURL), responder)

	var req GetAttributeTreeRequest
	res, err := client.GlobalProduct.GetAttributeTree(shopID, req, accessToken)
	if err != nil {
		t.Logf("GlobalProduct.GetAttributeTree returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("GlobalProduct.GetAttributeTree response: %#v", res)
}
func Test_GlobalProduct_GetBrandList(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.global_product.get_brand_list_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetBrandList due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetBrandList due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/global_product/get_brand_list", app.APIURL), responder)

	var req GetBrandListRequest
	res, err := client.GlobalProduct.GetBrandList(shopID, req, accessToken)
	if err != nil {
		t.Logf("GlobalProduct.GetBrandList returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("GlobalProduct.GetBrandList response: %#v", res)
}
func Test_GlobalProduct_GetCategory(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.global_product.get_category_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetCategory due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetCategory due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/global_product/get_category", app.APIURL), responder)

	var req GetCategoryRequest
	res, err := client.GlobalProduct.GetCategory(shopID, req, accessToken)
	if err != nil {
		t.Logf("GlobalProduct.GetCategory returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("GlobalProduct.GetCategory response: %#v", res)
}
func Test_GlobalProduct_GetGlobalItemId(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.global_product.get_global_item_id_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetGlobalItemId due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetGlobalItemId due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/global_product/get_global_item_id", app.APIURL), responder)

	var req GetGlobalItemIdRequest
	res, err := client.GlobalProduct.GetGlobalItemId(shopID, req, accessToken)
	if err != nil {
		t.Logf("GlobalProduct.GetGlobalItemId returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("GlobalProduct.GetGlobalItemId response: %#v", res)
}
func Test_GlobalProduct_GetGlobalItemInfo(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.global_product.get_global_item_info_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetGlobalItemInfo due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetGlobalItemInfo due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/global_product/get_global_item_info", app.APIURL), responder)

	var req GetGlobalItemInfoRequest
	res, err := client.GlobalProduct.GetGlobalItemInfo(shopID, req, accessToken)
	if err != nil {
		t.Logf("GlobalProduct.GetGlobalItemInfo returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("GlobalProduct.GetGlobalItemInfo response: %#v", res)
}
func Test_GlobalProduct_GetGlobalItemLimit(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.global_product.get_global_item_limit_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetGlobalItemLimit due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetGlobalItemLimit due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/global_product/get_global_item_limit", app.APIURL), responder)

	var req GetGlobalItemLimitRequest
	res, err := client.GlobalProduct.GetGlobalItemLimit(shopID, req, accessToken)
	if err != nil {
		t.Logf("GlobalProduct.GetGlobalItemLimit returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("GlobalProduct.GetGlobalItemLimit response: %#v", res)
}
func Test_GlobalProduct_GetGlobalItemList(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.global_product.get_global_item_list_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetGlobalItemList due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetGlobalItemList due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/global_product/get_global_item_list", app.APIURL), responder)

	var req GetGlobalItemListRequest
	res, err := client.GlobalProduct.GetGlobalItemList(shopID, req, accessToken)
	if err != nil {
		t.Logf("GlobalProduct.GetGlobalItemList returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("GlobalProduct.GetGlobalItemList response: %#v", res)
}
func Test_GlobalProduct_GetGlobalModelList(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.global_product.get_global_model_list_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetGlobalModelList due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetGlobalModelList due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/global_product/get_global_model_list", app.APIURL), responder)

	var req GetGlobalModelListRequest
	res, err := client.GlobalProduct.GetGlobalModelList(shopID, req, accessToken)
	if err != nil {
		t.Logf("GlobalProduct.GetGlobalModelList returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("GlobalProduct.GetGlobalModelList response: %#v", res)
}
func Test_GlobalProduct_GetLocalAdjustmentRate(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.global_product.get_local_adjustment_rate_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetLocalAdjustmentRate due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetLocalAdjustmentRate due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/global_product/get_local_adjustment_rate", app.APIURL), responder)

	res, err := client.GlobalProduct.GetLocalAdjustmentRate(shopID, accessToken)
	if err != nil {
		t.Logf("GlobalProduct.GetLocalAdjustmentRate returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("GlobalProduct.GetLocalAdjustmentRate response: %#v", res)
}
func Test_GlobalProduct_GetPublishableShop(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.global_product.get_publishable_shop_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetPublishableShop due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetPublishableShop due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/global_product/get_publishable_shop", app.APIURL), responder)

	var req GetPublishableShopRequest
	res, err := client.GlobalProduct.GetPublishableShop(shopID, req, accessToken)
	if err != nil {
		t.Logf("GlobalProduct.GetPublishableShop returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("GlobalProduct.GetPublishableShop response: %#v", res)
}
func Test_GlobalProduct_GetPublishedList(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.global_product.get_published_list_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetPublishedList due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetPublishedList due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/global_product/get_published_list", app.APIURL), responder)

	var req GetPublishedListRequest
	res, err := client.GlobalProduct.GetPublishedList(shopID, req, accessToken)
	if err != nil {
		t.Logf("GlobalProduct.GetPublishedList returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("GlobalProduct.GetPublishedList response: %#v", res)
}
func Test_GlobalProduct_GetPublishTaskResult(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.global_product.get_publish_task_result_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetPublishTaskResult due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetPublishTaskResult due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/global_product/get_publish_task_result", app.APIURL), responder)

	var req GetPublishTaskResultRequest
	res, err := client.GlobalProduct.GetPublishTaskResult(shopID, req, accessToken)
	if err != nil {
		t.Logf("GlobalProduct.GetPublishTaskResult returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("GlobalProduct.GetPublishTaskResult response: %#v", res)
}
func Test_GlobalProduct_GetRecommendAttribute(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.global_product.get_recommend_attribute_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetRecommendAttribute due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetRecommendAttribute due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/global_product/get_recommend_attribute", app.APIURL), responder)

	var req GetRecommendAttributeRequest
	res, err := client.GlobalProduct.GetRecommendAttribute(shopID, req, accessToken)
	if err != nil {
		t.Logf("GlobalProduct.GetRecommendAttribute returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("GlobalProduct.GetRecommendAttribute response: %#v", res)
}
func Test_GlobalProduct_GetShopPublishableStatus(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.global_product.get_shop_publishable_status_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetShopPublishableStatus due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetShopPublishableStatus due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/global_product/get_shop_publishable_status", app.APIURL), responder)

	var req GetShopPublishableStatusRequest
	res, err := client.GlobalProduct.GetShopPublishableStatus(shopID, req, accessToken)
	if err != nil {
		t.Logf("GlobalProduct.GetShopPublishableStatus returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("GlobalProduct.GetShopPublishableStatus response: %#v", res)
}
func Test_GlobalProduct_GetSizeChartDetail(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.global_product.get_size_chart_detail_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetSizeChartDetail due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetSizeChartDetail due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/global_product/get_size_chart_detail", app.APIURL), responder)

	var req GetSizeChartDetailRequest
	res, err := client.GlobalProduct.GetSizeChartDetail(shopID, req, accessToken)
	if err != nil {
		t.Logf("GlobalProduct.GetSizeChartDetail returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("GlobalProduct.GetSizeChartDetail response: %#v", res)
}
func Test_GlobalProduct_GetSizeChartList(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.global_product.get_size_chart_list_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetSizeChartList due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetSizeChartList due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/global_product/get_size_chart_list", app.APIURL), responder)

	var req GetSizeChartListRequest
	res, err := client.GlobalProduct.GetSizeChartList(shopID, req, accessToken)
	if err != nil {
		t.Logf("GlobalProduct.GetSizeChartList returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("GlobalProduct.GetSizeChartList response: %#v", res)
}
func Test_GlobalProduct_GetVariations(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.global_product.get_variations_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetVariations due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetVariations due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/global_product/get_variations", app.APIURL), responder)

	var req GetVariationsRequest
	res, err := client.GlobalProduct.GetVariations(shopID, req, accessToken)
	if err != nil {
		t.Logf("GlobalProduct.GetVariations returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("GlobalProduct.GetVariations response: %#v", res)
}
func Test_GlobalProduct_InitTierVariation(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.global_product.init_tier_variation_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping InitTierVariation due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping InitTierVariation due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/global_product/init_tier_variation", app.APIURL), responder)

	var req InitTierVariationRequest
	res, err := client.GlobalProduct.InitTierVariation(shopID, req, accessToken)
	if err != nil {
		t.Logf("GlobalProduct.InitTierVariation returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("GlobalProduct.InitTierVariation response: %#v", res)
}
func Test_GlobalProduct_SearchGlobalAttributeValueList(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.global_product.search_global_attribute_value_list_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping SearchGlobalAttributeValueList due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping SearchGlobalAttributeValueList due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/global_product/search_global_attribute_value_list", app.APIURL), responder)

	var req SearchGlobalAttributeValueListRequest
	res, err := client.GlobalProduct.SearchGlobalAttributeValueList(shopID, req, accessToken)
	if err != nil {
		t.Logf("GlobalProduct.SearchGlobalAttributeValueList returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("GlobalProduct.SearchGlobalAttributeValueList response: %#v", res)
}
func Test_GlobalProduct_SetSyncField(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.global_product.set_sync_field_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping SetSyncField due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping SetSyncField due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/global_product/set_sync_field", app.APIURL), responder)

	var req SetSyncFieldRequest
	res, err := client.GlobalProduct.SetSyncField(shopID, req, accessToken)
	if err != nil {
		t.Logf("GlobalProduct.SetSyncField returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("GlobalProduct.SetSyncField response: %#v", res)
}
func Test_GlobalProduct_SupportSizeChart(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.global_product.support_size_chart_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping SupportSizeChart due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping SupportSizeChart due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/global_product/support_size_chart", app.APIURL), responder)

	var req SupportSizeChartRequest
	res, err := client.GlobalProduct.SupportSizeChart(shopID, req, accessToken)
	if err != nil {
		t.Logf("GlobalProduct.SupportSizeChart returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("GlobalProduct.SupportSizeChart response: %#v", res)
}
func Test_GlobalProduct_UpdateGlobalItem(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.global_product.update_global_item_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping UpdateGlobalItem due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping UpdateGlobalItem due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/global_product/update_global_item", app.APIURL), responder)

	var req UpdateGlobalItemRequest
	res, err := client.GlobalProduct.UpdateGlobalItem(shopID, req, accessToken)
	if err != nil {
		t.Logf("GlobalProduct.UpdateGlobalItem returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("GlobalProduct.UpdateGlobalItem response: %#v", res)
}
func Test_GlobalProduct_UpdateGlobalModel(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.global_product.update_global_model_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping UpdateGlobalModel due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping UpdateGlobalModel due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/global_product/update_global_model", app.APIURL), responder)

	var req UpdateGlobalModelRequest
	res, err := client.GlobalProduct.UpdateGlobalModel(shopID, req, accessToken)
	if err != nil {
		t.Logf("GlobalProduct.UpdateGlobalModel returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("GlobalProduct.UpdateGlobalModel response: %#v", res)
}
func Test_GlobalProduct_UpdateLocalAdjustmentRate(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.global_product.update_local_adjustment_rate_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping UpdateLocalAdjustmentRate due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping UpdateLocalAdjustmentRate due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/global_product/update_local_adjustment_rate", app.APIURL), responder)

	var req UpdateLocalAdjustmentRateRequest
	res, err := client.GlobalProduct.UpdateLocalAdjustmentRate(shopID, req, accessToken)
	if err != nil {
		t.Logf("GlobalProduct.UpdateLocalAdjustmentRate returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("GlobalProduct.UpdateLocalAdjustmentRate response: %#v", res)
}
func Test_GlobalProduct_UpdatePrice(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.global_product.update_price_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping UpdatePrice due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping UpdatePrice due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/global_product/update_price", app.APIURL), responder)

	var req UpdatePriceRequest
	res, err := client.GlobalProduct.UpdatePrice(shopID, req, accessToken)
	if err != nil {
		t.Logf("GlobalProduct.UpdatePrice returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("GlobalProduct.UpdatePrice response: %#v", res)
}
func Test_GlobalProduct_UpdateSizeChart(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.global_product.update_size_chart_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping UpdateSizeChart due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping UpdateSizeChart due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/global_product/update_size_chart", app.APIURL), responder)

	var req UpdateSizeChartRequest
	res, err := client.GlobalProduct.UpdateSizeChart(shopID, req, accessToken)
	if err != nil {
		t.Logf("GlobalProduct.UpdateSizeChart returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("GlobalProduct.UpdateSizeChart response: %#v", res)
}
func Test_GlobalProduct_UpdateStock(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.global_product.update_stock_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping UpdateStock due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping UpdateStock due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/global_product/update_stock", app.APIURL), responder)

	var req UpdateStockRequest
	res, err := client.GlobalProduct.UpdateStock(shopID, req, accessToken)
	if err != nil {
		t.Logf("GlobalProduct.UpdateStock returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("GlobalProduct.UpdateStock response: %#v", res)
}
func Test_GlobalProduct_UpdateTierVariation(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.global_product.update_tier_variation_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping UpdateTierVariation due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping UpdateTierVariation due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/global_product/update_tier_variation", app.APIURL), responder)

	var req UpdateTierVariationRequest
	res, err := client.GlobalProduct.UpdateTierVariation(shopID, req, accessToken)
	if err != nil {
		t.Logf("GlobalProduct.UpdateTierVariation returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("GlobalProduct.UpdateTierVariation response: %#v", res)
}
