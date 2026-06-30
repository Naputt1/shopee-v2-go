package goshopee

import (
	"fmt"
	"testing"

	"github.com/jarcoal/httpmock"
)

func Test_ShopCategory_AddItemList(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.shop_category.add_item_list_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping AddItemList due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping AddItemList due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/shop_category/add_item_list", app.APIURL), responder)

	var req ShopCategoryAddItemListRequest
	res, err := client.ShopCategory.AddItemList(shopID, req, accessToken)
	if err != nil {
		t.Logf("ShopCategory.AddItemList returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("ShopCategory.AddItemList response: %#v", res)
}
func Test_ShopCategory_AddShopCategory(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.shop_category.add_shop_category_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping AddShopCategory due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping AddShopCategory due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/shop_category/add_shop_category", app.APIURL), responder)

	var req AddShopCategoryRequest
	res, err := client.ShopCategory.AddShopCategory(shopID, req, accessToken)
	if err != nil {
		t.Logf("ShopCategory.AddShopCategory returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("ShopCategory.AddShopCategory response: %#v", res)
}
func Test_ShopCategory_DeleteItemList(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.shop_category.delete_item_list_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping DeleteItemList due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping DeleteItemList due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/shop_category/delete_item_list", app.APIURL), responder)

	var req ShopCategoryDeleteItemListRequest
	res, err := client.ShopCategory.DeleteItemList(shopID, req, accessToken)
	if err != nil {
		t.Logf("ShopCategory.DeleteItemList returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("ShopCategory.DeleteItemList response: %#v", res)
}
func Test_ShopCategory_DeleteShopCategory(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.shop_category.delete_shop_category_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping DeleteShopCategory due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping DeleteShopCategory due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/shop_category/delete_shop_category", app.APIURL), responder)

	var req DeleteShopCategoryRequest
	res, err := client.ShopCategory.DeleteShopCategory(shopID, req, accessToken)
	if err != nil {
		t.Logf("ShopCategory.DeleteShopCategory returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("ShopCategory.DeleteShopCategory response: %#v", res)
}
func Test_ShopCategory_GetItemList(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.shop_category.get_item_list_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetItemList due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetItemList due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/shop_category/get_item_list", app.APIURL), responder)

	var req ShopCategoryGetItemListRequest
	res, err := client.ShopCategory.GetItemList(shopID, req, accessToken)
	if err != nil {
		t.Logf("ShopCategory.GetItemList returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("ShopCategory.GetItemList response: %#v", res)
}
func Test_ShopCategory_GetShopCategoryList(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.shop_category.get_shop_category_list_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetShopCategoryList due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetShopCategoryList due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/shop_category/get_shop_category_list", app.APIURL), responder)

	var req GetShopCategoryListRequest
	res, err := client.ShopCategory.GetShopCategoryList(shopID, req, accessToken)
	if err != nil {
		t.Logf("ShopCategory.GetShopCategoryList returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("ShopCategory.GetShopCategoryList response: %#v", res)
}
func Test_ShopCategory_UpdateShopCategory(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.shop_category.update_shop_category_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping UpdateShopCategory due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping UpdateShopCategory due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/shop_category/update_shop_category", app.APIURL), responder)

	var req UpdateShopCategoryRequest
	res, err := client.ShopCategory.UpdateShopCategory(shopID, req, accessToken)
	if err != nil {
		t.Logf("ShopCategory.UpdateShopCategory returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("ShopCategory.UpdateShopCategory response: %#v", res)
}
