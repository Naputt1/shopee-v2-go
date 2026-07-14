package goshopee

import (
	"context"
	"fmt"
	"testing"

	"github.com/jarcoal/httpmock"
)

func Test_Product_AddItem(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.product.add_item_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping AddItem due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping AddItem due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/product/add_item", app.APIURL), responder)

	var req AddItemRequest
	ctx := context.Background()
	res, err := client.Product.AddItem(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Product.AddItem returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Product.AddItem response: %#v", res)
}
func Test_Product_AddKitItem(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.product.add_kit_item_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping AddKitItem due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping AddKitItem due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/product/add_kit_item", app.APIURL), responder)

	var req AddKitItemRequest
	ctx := context.Background()
	res, err := client.Product.AddKitItem(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Product.AddKitItem returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Product.AddKitItem response: %#v", res)
}
func Test_Product_AddModel(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.product.add_model_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping AddModel due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping AddModel due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/product/add_model", app.APIURL), responder)

	var req AddModelRequest
	ctx := context.Background()
	res, err := client.Product.AddModel(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Product.AddModel returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Product.AddModel response: %#v", res)
}
func Test_Product_BatchAddItem(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.product.batch_add_item_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping BatchAddItem due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping BatchAddItem due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/product/batch_add_item", app.APIURL), responder)

	var req BatchAddItemRequest
	ctx := context.Background()
	res, err := client.Product.BatchAddItem(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Product.BatchAddItem returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Product.BatchAddItem response: %#v", res)
}
func Test_Product_BatchPublishItemToOutletShop(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.product.batch_publish_item_to_outlet_shop_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping BatchPublishItemToOutletShop due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping BatchPublishItemToOutletShop due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/product/batch_publish_item_to_outlet_shop", app.APIURL), responder)

	var req BatchPublishItemToOutletShopRequest
	ctx := context.Background()
	res, err := client.Product.BatchPublishItemToOutletShop(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Product.BatchPublishItemToOutletShop returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Product.BatchPublishItemToOutletShop response: %#v", res)
}
func Test_Product_BatchUpdateOutletPrice(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.product.batch_update_outlet_price_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping BatchUpdateOutletPrice due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping BatchUpdateOutletPrice due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/product/batch_update_outlet_price", app.APIURL), responder)

	var req BatchUpdateOutletPriceRequest
	ctx := context.Background()
	res, err := client.Product.BatchUpdateOutletPrice(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Product.BatchUpdateOutletPrice returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Product.BatchUpdateOutletPrice response: %#v", res)
}
func Test_Product_BatchUpdateOutletStock(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.product.batch_update_outlet_stock_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping BatchUpdateOutletStock due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping BatchUpdateOutletStock due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/product/batch_update_outlet_stock", app.APIURL), responder)

	var req BatchUpdateOutletStockRequest
	ctx := context.Background()
	res, err := client.Product.BatchUpdateOutletStock(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Product.BatchUpdateOutletStock returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Product.BatchUpdateOutletStock response: %#v", res)
}
func Test_Product_BoostItem(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.product.boost_item_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping BoostItem due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping BoostItem due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/product/boost_item", app.APIURL), responder)

	var req BoostItemRequest
	ctx := context.Background()
	res, err := client.Product.BoostItem(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Product.BoostItem returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Product.BoostItem response: %#v", res)
}
func Test_Product_CategoryRecommend(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.product.category_recommend_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping CategoryRecommend due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping CategoryRecommend due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/product/category_recommend", app.APIURL), responder)

	var req ProductCategoryRecommendRequest
	ctx := context.Background()
	res, err := client.Product.CategoryRecommend(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Product.CategoryRecommend returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Product.CategoryRecommend response: %#v", res)
}
func Test_Product_DeleteItem(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.product.delete_item_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping DeleteItem due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping DeleteItem due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/product/delete_item", app.APIURL), responder)

	var req DeleteItemRequest
	ctx := context.Background()
	res, err := client.Product.DeleteItem(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Product.DeleteItem returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Product.DeleteItem response: %#v", res)
}
func Test_Product_DeleteModel(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.product.delete_model_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping DeleteModel due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping DeleteModel due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/product/delete_model", app.APIURL), responder)

	var req DeleteModelRequest
	ctx := context.Background()
	res, err := client.Product.DeleteModel(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Product.DeleteModel returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Product.DeleteModel response: %#v", res)
}
func Test_Product_GenerateKitImage(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.product.generate_kit_image_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GenerateKitImage due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GenerateKitImage due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/product/generate_kit_image", app.APIURL), responder)

	var req GenerateKitImageRequest
	ctx := context.Background()
	res, err := client.Product.GenerateKitImage(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Product.GenerateKitImage returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Product.GenerateKitImage response: %#v", res)
}
func Test_Product_GetAitemByPitemId(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.product.get_aitem_by_pitem_id_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetAitemByPitemId due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetAitemByPitemId due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/product/get_aitem_by_pitem_id", app.APIURL), responder)

	var req GetAitemByPitemIdRequest
	ctx := context.Background()
	res, err := client.Product.GetAitemByPitemId(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Product.GetAitemByPitemId returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Product.GetAitemByPitemId response: %#v", res)
}
func Test_Product_GetAllVehicleList(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.product.get_all_vehicle_list_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetAllVehicleList due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetAllVehicleList due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/product/get_all_vehicle_list", app.APIURL), responder)

	var req GetAllVehicleListRequest
	ctx := context.Background()
	res, err := client.Product.GetAllVehicleList(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Product.GetAllVehicleList returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Product.GetAllVehicleList response: %#v", res)
}
func Test_Product_GetAttributeTree(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.product.get_attribute_tree_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetAttributeTree due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetAttributeTree due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/product/get_attribute_tree", app.APIURL), responder)

	var req ProductGetAttributeTreeRequest
	ctx := context.Background()
	res, err := client.Product.GetAttributeTree(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Product.GetAttributeTree returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Product.GetAttributeTree response: %#v", res)
}
func Test_Product_GetBatchTaskResult(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.product.get_batch_task_result_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetBatchTaskResult due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetBatchTaskResult due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/product/get_batch_task_result", app.APIURL), responder)

	var req GetBatchTaskResultRequest
	ctx := context.Background()
	res, err := client.Product.GetBatchTaskResult(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Product.GetBatchTaskResult returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Product.GetBatchTaskResult response: %#v", res)
}
func Test_Product_GetBoostedList(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.product.get_boosted_list_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetBoostedList due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetBoostedList due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/product/get_boosted_list", app.APIURL), responder)

	ctx := context.Background()
	res, err := client.Product.GetBoostedList(ctx, sid, mid, tok)
	if err != nil {
		t.Logf("Product.GetBoostedList returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Product.GetBoostedList response: %#v", res)
}
func Test_Product_GetBrandList(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.product.get_brand_list_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetBrandList due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetBrandList due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/product/get_brand_list", app.APIURL), responder)

	var req ProductGetBrandListRequest
	ctx := context.Background()
	res, err := client.Product.GetBrandList(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Product.GetBrandList returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Product.GetBrandList response: %#v", res)
}
func Test_Product_GetCategory(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.product.get_category_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetCategory due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetCategory due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/product/get_category", app.APIURL), responder)

	var req ProductGetCategoryRequest
	ctx := context.Background()
	res, err := client.Product.GetCategory(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Product.GetCategory returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Product.GetCategory response: %#v", res)
}
func Test_Product_GetComment(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.product.get_comment_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetComment due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetComment due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/product/get_comment", app.APIURL), responder)

	var req GetCommentRequest
	ctx := context.Background()
	res, err := client.Product.GetComment(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Product.GetComment returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Product.GetComment response: %#v", res)
}
func Test_Product_GetDirectItemList(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.product.get_direct_item_list_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetDirectItemList due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetDirectItemList due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/product/get_direct_item_list", app.APIURL), responder)

	var req GetDirectItemListRequest
	ctx := context.Background()
	res, err := client.Product.GetDirectItemList(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Product.GetDirectItemList returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Product.GetDirectItemList response: %#v", res)
}
func Test_Product_GetDirectShopRecommendedPrice(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.product.get_direct_shop_recommended_price_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetDirectShopRecommendedPrice due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetDirectShopRecommendedPrice due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/product/get_direct_shop_recommended_price", app.APIURL), responder)

	var req GetDirectShopRecommendedPriceRequest
	ctx := context.Background()
	res, err := client.Product.GetDirectShopRecommendedPrice(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Product.GetDirectShopRecommendedPrice returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Product.GetDirectShopRecommendedPrice response: %#v", res)
}
func Test_Product_GetItemBaseInfo(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.product.get_item_base_info_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetItemBaseInfo due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetItemBaseInfo due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/product/get_item_base_info", app.APIURL), responder)

	var req GetItemBaseInfoRequest
	ctx := context.Background()
	res, err := client.Product.GetItemBaseInfo(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Product.GetItemBaseInfo returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Product.GetItemBaseInfo response: %#v", res)
}
func Test_Product_GetItemContentDiagnosisResult(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.product.get_item_content_diagnosis_result_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetItemContentDiagnosisResult due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetItemContentDiagnosisResult due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/product/get_item_content_diagnosis_result", app.APIURL), responder)

	var req GetItemContentDiagnosisResultRequest
	ctx := context.Background()
	res, err := client.Product.GetItemContentDiagnosisResult(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Product.GetItemContentDiagnosisResult returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Product.GetItemContentDiagnosisResult response: %#v", res)
}
func Test_Product_GetItemExtraInfo(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.product.get_item_extra_info_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetItemExtraInfo due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetItemExtraInfo due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/product/get_item_extra_info", app.APIURL), responder)

	var req GetItemExtraInfoRequest
	ctx := context.Background()
	res, err := client.Product.GetItemExtraInfo(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Product.GetItemExtraInfo returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Product.GetItemExtraInfo response: %#v", res)
}
func Test_Product_GetItemLimit(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.product.get_item_limit_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetItemLimit due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetItemLimit due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/product/get_item_limit", app.APIURL), responder)

	var req GetItemLimitRequest
	ctx := context.Background()
	res, err := client.Product.GetItemLimit(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Product.GetItemLimit returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Product.GetItemLimit response: %#v", res)
}
func Test_Product_GetItemList(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.product.get_item_list_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetItemList due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetItemList due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/product/get_item_list", app.APIURL), responder)

	var req ProductGetItemListRequest
	ctx := context.Background()
	res, err := client.Product.GetItemList(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Product.GetItemList returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Product.GetItemList response: %#v", res)
}
func Test_Product_GetItemListByContentDiagnosis(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.product.get_item_list_by_content_diagnosis_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetItemListByContentDiagnosis due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetItemListByContentDiagnosis due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/product/get_item_list_by_content_diagnosis", app.APIURL), responder)

	var req GetItemListByContentDiagnosisRequest
	ctx := context.Background()
	res, err := client.Product.GetItemListByContentDiagnosis(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Product.GetItemListByContentDiagnosis returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Product.GetItemListByContentDiagnosis response: %#v", res)
}
func Test_Product_GetItemPromotion(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.product.get_item_promotion_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetItemPromotion due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetItemPromotion due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/product/get_item_promotion", app.APIURL), responder)

	var req GetItemPromotionRequest
	ctx := context.Background()
	res, err := client.Product.GetItemPromotion(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Product.GetItemPromotion returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Product.GetItemPromotion response: %#v", res)
}
func Test_Product_GetItemViolationInfo(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.product.get_item_violation_info_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetItemViolationInfo due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetItemViolationInfo due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/product/get_item_violation_info", app.APIURL), responder)

	var req GetItemViolationInfoRequest
	ctx := context.Background()
	res, err := client.Product.GetItemViolationInfo(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Product.GetItemViolationInfo returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Product.GetItemViolationInfo response: %#v", res)
}
func Test_Product_GetKitItemInfo(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.product.get_kit_item_info_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetKitItemInfo due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetKitItemInfo due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/product/get_kit_item_info", app.APIURL), responder)

	var req GetKitItemInfoRequest
	ctx := context.Background()
	res, err := client.Product.GetKitItemInfo(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Product.GetKitItemInfo returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Product.GetKitItemInfo response: %#v", res)
}
func Test_Product_GetKitItemLimit(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.product.get_kit_item_limit_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetKitItemLimit due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetKitItemLimit due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/product/get_kit_item_limit", app.APIURL), responder)

	var req GetKitItemLimitRequest
	ctx := context.Background()
	res, err := client.Product.GetKitItemLimit(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Product.GetKitItemLimit returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Product.GetKitItemLimit response: %#v", res)
}
func Test_Product_GetMainItemList(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.product.get_main_item_list_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetMainItemList due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetMainItemList due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/product/get_main_item_list", app.APIURL), responder)

	var req GetMainItemListRequest
	ctx := context.Background()
	res, err := client.Product.GetMainItemList(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Product.GetMainItemList returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Product.GetMainItemList response: %#v", res)
}
func Test_Product_GetMartItemByOutletItemId(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.product.get_mart_item_by_outlet_item_id_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetMartItemByOutletItemId due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetMartItemByOutletItemId due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/product/get_mart_item_by_outlet_item_id", app.APIURL), responder)

	var req GetMartItemByOutletItemIdRequest
	ctx := context.Background()
	res, err := client.Product.GetMartItemByOutletItemId(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Product.GetMartItemByOutletItemId returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Product.GetMartItemByOutletItemId response: %#v", res)
}
func Test_Product_GetMartItemMappingById(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.product.get_mart_item_mapping_by_id_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetMartItemMappingById due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetMartItemMappingById due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/product/get_mart_item_mapping_by_id", app.APIURL), responder)

	var req GetMartItemMappingByIdRequest
	ctx := context.Background()
	res, err := client.Product.GetMartItemMappingById(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Product.GetMartItemMappingById returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Product.GetMartItemMappingById response: %#v", res)
}
func Test_Product_GetModelList(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.product.get_model_list_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetModelList due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetModelList due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/product/get_model_list", app.APIURL), responder)

	var req GetModelListRequest
	ctx := context.Background()
	res, err := client.Product.GetModelList(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Product.GetModelList returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Product.GetModelList response: %#v", res)
}
func Test_Product_GetProductCertificationRule(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.product.get_product_certification_rule_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetProductCertificationRule due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetProductCertificationRule due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/product/get_product_certification_rule", app.APIURL), responder)

	var req GetProductCertificationRuleRequest
	ctx := context.Background()
	res, err := client.Product.GetProductCertificationRule(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Product.GetProductCertificationRule returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Product.GetProductCertificationRule response: %#v", res)
}
func Test_Product_GetRecommendAttribute(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.product.get_recommend_attribute_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetRecommendAttribute due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetRecommendAttribute due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/product/get_recommend_attribute", app.APIURL), responder)

	var req ProductGetRecommendAttributeRequest
	ctx := context.Background()
	res, err := client.Product.GetRecommendAttribute(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Product.GetRecommendAttribute returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Product.GetRecommendAttribute response: %#v", res)
}
func Test_Product_GetSizeChartDetail(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.product.get_size_chart_detail_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetSizeChartDetail due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetSizeChartDetail due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/product/get_size_chart_detail", app.APIURL), responder)

	var req ProductGetSizeChartDetailRequest
	ctx := context.Background()
	res, err := client.Product.GetSizeChartDetail(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Product.GetSizeChartDetail returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Product.GetSizeChartDetail response: %#v", res)
}
func Test_Product_GetSizeChartList(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.product.get_size_chart_list_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetSizeChartList due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetSizeChartList due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/product/get_size_chart_list", app.APIURL), responder)

	var req ProductGetSizeChartListRequest
	ctx := context.Background()
	res, err := client.Product.GetSizeChartList(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Product.GetSizeChartList returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Product.GetSizeChartList response: %#v", res)
}
func Test_Product_GetVariations(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.product.get_variations_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetVariations due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetVariations due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/product/get_variation_tree", app.APIURL), responder)

	var req ProductGetVariationsRequest
	ctx := context.Background()
	res, err := client.Product.GetVariations(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Product.GetVariations returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Product.GetVariations response: %#v", res)
}
func Test_Product_GetVehicleListByCompatibilityDetail(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.product.get_vehicle_list_by_compatibility_detail_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetVehicleListByCompatibilityDetail due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetVehicleListByCompatibilityDetail due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/product/get_vehicle_list_by_compatibility_detail", app.APIURL), responder)

	var req GetVehicleListByCompatibilityDetailRequest
	ctx := context.Background()
	res, err := client.Product.GetVehicleListByCompatibilityDetail(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Product.GetVehicleListByCompatibilityDetail returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Product.GetVehicleListByCompatibilityDetail response: %#v", res)
}
func Test_Product_GetWeightRecommendation(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.product.get_weight_recommendation_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetWeightRecommendation due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetWeightRecommendation due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/product/get_weight_recommendation", app.APIURL), responder)

	var req GetWeightRecommendationRequest
	ctx := context.Background()
	res, err := client.Product.GetWeightRecommendation(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Product.GetWeightRecommendation returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Product.GetWeightRecommendation response: %#v", res)
}
func Test_Product_InitTierVariation(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.product.init_tier_variation_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping InitTierVariation due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping InitTierVariation due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/product/init_tier_variation", app.APIURL), responder)

	var req ProductInitTierVariationRequest
	ctx := context.Background()
	res, err := client.Product.InitTierVariation(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Product.InitTierVariation returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Product.InitTierVariation response: %#v", res)
}
func Test_Product_PublishItemToOutletShop(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.product.publish_item_to_outlet_shop_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping PublishItemToOutletShop due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping PublishItemToOutletShop due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/", app.APIURL), responder)

	ctx := context.Background()
	res, err := client.Product.PublishItemToOutletShop(ctx, sid, mid, tok)
	if err != nil {
		t.Logf("Product.PublishItemToOutletShop returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Product.PublishItemToOutletShop response: %#v", res)
}
func Test_Product_RegisterBrand(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.product.register_brand_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping RegisterBrand due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping RegisterBrand due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/product/register_brand", app.APIURL), responder)

	var req RegisterBrandRequest
	ctx := context.Background()
	res, err := client.Product.RegisterBrand(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Product.RegisterBrand returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Product.RegisterBrand response: %#v", res)
}
func Test_Product_ReplyComment(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.product.reply_comment_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping ReplyComment due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping ReplyComment due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/product/reply_comment", app.APIURL), responder)

	var req ReplyCommentRequest
	ctx := context.Background()
	res, err := client.Product.ReplyComment(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Product.ReplyComment returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Product.ReplyComment response: %#v", res)
}
func Test_Product_SearchAttributeValueList(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.product.search_attribute_value_list_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping SearchAttributeValueList due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping SearchAttributeValueList due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/product/search_attribute_value_list", app.APIURL), responder)

	var req SearchAttributeValueListRequest
	ctx := context.Background()
	res, err := client.Product.SearchAttributeValueList(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Product.SearchAttributeValueList returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Product.SearchAttributeValueList response: %#v", res)
}
func Test_Product_SearchItem(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.product.search_item_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping SearchItem due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping SearchItem due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/product/search_item", app.APIURL), responder)

	var req SearchItemRequest
	ctx := context.Background()
	res, err := client.Product.SearchItem(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Product.SearchItem returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Product.SearchItem response: %#v", res)
}
func Test_Product_SearchUnpackagedModelList(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.product.search_unpackaged_model_list_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping SearchUnpackagedModelList due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping SearchUnpackagedModelList due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/product/search_unpackaged_model_list", app.APIURL), responder)

	var req SearchUnpackagedModelListRequest
	ctx := context.Background()
	res, err := client.Product.SearchUnpackagedModelList(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Product.SearchUnpackagedModelList returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Product.SearchUnpackagedModelList response: %#v", res)
}
func Test_Product_UnlistItem(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.product.unlist_item_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping UnlistItem due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping UnlistItem due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/product/unlist_item", app.APIURL), responder)

	var req UnlistItemRequest
	ctx := context.Background()
	res, err := client.Product.UnlistItem(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Product.UnlistItem returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Product.UnlistItem response: %#v", res)
}
func Test_Product_UpdateItem(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.product.update_item_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping UpdateItem due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping UpdateItem due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/product/update_item", app.APIURL), responder)

	var req UpdateItemRequest
	ctx := context.Background()
	res, err := client.Product.UpdateItem(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Product.UpdateItem returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Product.UpdateItem response: %#v", res)
}
func Test_Product_UpdateKitItem(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.product.update_kit_item_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping UpdateKitItem due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping UpdateKitItem due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/product/update_kit_item", app.APIURL), responder)

	var req UpdateKitItemRequest
	ctx := context.Background()
	res, err := client.Product.UpdateKitItem(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Product.UpdateKitItem returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Product.UpdateKitItem response: %#v", res)
}
func Test_Product_UpdateModel(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.product.update_model_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping UpdateModel due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping UpdateModel due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/product/update_model", app.APIURL), responder)

	var req UpdateModelRequest
	ctx := context.Background()
	res, err := client.Product.UpdateModel(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Product.UpdateModel returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Product.UpdateModel response: %#v", res)
}
func Test_Product_UpdatePrice(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.product.update_price_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping UpdatePrice due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping UpdatePrice due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/product/update_price", app.APIURL), responder)

	var req ProductUpdatePriceRequest
	ctx := context.Background()
	res, err := client.Product.UpdatePrice(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Product.UpdatePrice returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Product.UpdatePrice response: %#v", res)
}
func Test_Product_UpdateSipItemPrice(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.product.update_sip_item_price_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping UpdateSipItemPrice due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping UpdateSipItemPrice due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/product/update_sip_item_price", app.APIURL), responder)

	var req UpdateSipItemPriceRequest
	ctx := context.Background()
	res, err := client.Product.UpdateSipItemPrice(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Product.UpdateSipItemPrice returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Product.UpdateSipItemPrice response: %#v", res)
}
func Test_Product_UpdateStock(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.product.update_stock_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping UpdateStock due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping UpdateStock due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/product/update_stock", app.APIURL), responder)

	var req ProductUpdateStockRequest
	ctx := context.Background()
	res, err := client.Product.UpdateStock(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Product.UpdateStock returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Product.UpdateStock response: %#v", res)
}
func Test_Product_UpdateTierVariation(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.product.update_tier_variation_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping UpdateTierVariation due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping UpdateTierVariation due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/product/update_tier_variation", app.APIURL), responder)

	var req ProductUpdateTierVariationRequest
	ctx := context.Background()
	res, err := client.Product.UpdateTierVariation(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Product.UpdateTierVariation returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Product.UpdateTierVariation response: %#v", res)
}
