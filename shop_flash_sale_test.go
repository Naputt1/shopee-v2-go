package goshopee

import (
	"fmt"
	"testing"

	"github.com/jarcoal/httpmock"
)

func Test_ShopFlashSale_AddShopFlashSaleItems(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.shop_flash_sale.add_shop_flash_sale_items_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping AddShopFlashSaleItems due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping AddShopFlashSaleItems due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/shop_flash_sale/add_shop_flash_sale_items", app.APIURL), responder)

	var req AddShopFlashSaleItemsRequest
	res, err := client.ShopFlashSale.AddShopFlashSaleItems(shopID, req, accessToken)
	if err != nil {
		t.Logf("ShopFlashSale.AddShopFlashSaleItems returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("ShopFlashSale.AddShopFlashSaleItems response: %#v", res)
}
func Test_ShopFlashSale_CreateShopFlashSale(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.shop_flash_sale.create_shop_flash_sale_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping CreateShopFlashSale due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping CreateShopFlashSale due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/shop_flash_sale/create_shop_flash_sale", app.APIURL), responder)

	var req CreateShopFlashSaleRequest
	res, err := client.ShopFlashSale.CreateShopFlashSale(shopID, req, accessToken)
	if err != nil {
		t.Logf("ShopFlashSale.CreateShopFlashSale returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("ShopFlashSale.CreateShopFlashSale response: %#v", res)
}
func Test_ShopFlashSale_DeleteShopFlashSale(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.shop_flash_sale.delete_shop_flash_sale_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping DeleteShopFlashSale due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping DeleteShopFlashSale due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/shop_flash_sale/delete_shop_flash_sale", app.APIURL), responder)

	var req DeleteShopFlashSaleRequest
	res, err := client.ShopFlashSale.DeleteShopFlashSale(shopID, req, accessToken)
	if err != nil {
		t.Logf("ShopFlashSale.DeleteShopFlashSale returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("ShopFlashSale.DeleteShopFlashSale response: %#v", res)
}
func Test_ShopFlashSale_DeleteShopFlashSaleItems(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.shop_flash_sale.delete_shop_flash_sale_items_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping DeleteShopFlashSaleItems due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping DeleteShopFlashSaleItems due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/shop_flash_sale/delete_shop_flash_sale_items", app.APIURL), responder)

	var req DeleteShopFlashSaleItemsRequest
	res, err := client.ShopFlashSale.DeleteShopFlashSaleItems(shopID, req, accessToken)
	if err != nil {
		t.Logf("ShopFlashSale.DeleteShopFlashSaleItems returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("ShopFlashSale.DeleteShopFlashSaleItems response: %#v", res)
}
func Test_ShopFlashSale_GetItemCriteria(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.shop_flash_sale.get_item_criteria_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetItemCriteria due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetItemCriteria due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/shop_flash_sale/get_item_criteria", app.APIURL), responder)

	res, err := client.ShopFlashSale.GetItemCriteria(shopID, accessToken)
	if err != nil {
		t.Logf("ShopFlashSale.GetItemCriteria returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("ShopFlashSale.GetItemCriteria response: %#v", res)
}
func Test_ShopFlashSale_GetShopFlashSale(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.shop_flash_sale.get_shop_flash_sale_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetShopFlashSale due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetShopFlashSale due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/shop_flash_sale/get_shop_flash_sale", app.APIURL), responder)

	var req GetShopFlashSaleRequest
	res, err := client.ShopFlashSale.GetShopFlashSale(shopID, req, accessToken)
	if err != nil {
		t.Logf("ShopFlashSale.GetShopFlashSale returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("ShopFlashSale.GetShopFlashSale response: %#v", res)
}
func Test_ShopFlashSale_GetShopFlashSaleItems(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.shop_flash_sale.get_shop_flash_sale_items_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetShopFlashSaleItems due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetShopFlashSaleItems due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/shop_flash_sale/get_shop_flash_sale_items", app.APIURL), responder)

	var req GetShopFlashSaleItemsRequest
	res, err := client.ShopFlashSale.GetShopFlashSaleItems(shopID, req, accessToken)
	if err != nil {
		t.Logf("ShopFlashSale.GetShopFlashSaleItems returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("ShopFlashSale.GetShopFlashSaleItems response: %#v", res)
}
func Test_ShopFlashSale_GetShopFlashSaleList(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.shop_flash_sale.get_shop_flash_sale_list_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetShopFlashSaleList due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetShopFlashSaleList due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/shop_flash_sale/get_shop_flash_sale_list", app.APIURL), responder)

	var req GetShopFlashSaleListRequest
	res, err := client.ShopFlashSale.GetShopFlashSaleList(shopID, req, accessToken)
	if err != nil {
		t.Logf("ShopFlashSale.GetShopFlashSaleList returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("ShopFlashSale.GetShopFlashSaleList response: %#v", res)
}
func Test_ShopFlashSale_GetTimeSlotId(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.shop_flash_sale.get_time_slot_id_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetTimeSlotId due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetTimeSlotId due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/shop_flash_sale/get_time_slot_id", app.APIURL), responder)

	var req GetTimeSlotIdRequest
	res, err := client.ShopFlashSale.GetTimeSlotId(shopID, req, accessToken)
	if err != nil {
		t.Logf("ShopFlashSale.GetTimeSlotId returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("ShopFlashSale.GetTimeSlotId response: %#v", res)
}
func Test_ShopFlashSale_UpdateShopFlashSale(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.shop_flash_sale.update_shop_flash_sale_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping UpdateShopFlashSale due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping UpdateShopFlashSale due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/shop_flash_sale/update_shop_flash_sale", app.APIURL), responder)

	var req UpdateShopFlashSaleRequest
	res, err := client.ShopFlashSale.UpdateShopFlashSale(shopID, req, accessToken)
	if err != nil {
		t.Logf("ShopFlashSale.UpdateShopFlashSale returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("ShopFlashSale.UpdateShopFlashSale response: %#v", res)
}
func Test_ShopFlashSale_UpdateShopFlashSaleItems(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.shop_flash_sale.update_shop_flash_sale_items_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping UpdateShopFlashSaleItems due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping UpdateShopFlashSaleItems due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/shop_flash_sale/update_shop_flash_sale_items", app.APIURL), responder)

	var req UpdateShopFlashSaleItemsRequest
	res, err := client.ShopFlashSale.UpdateShopFlashSaleItems(shopID, req, accessToken)
	if err != nil {
		t.Logf("ShopFlashSale.UpdateShopFlashSaleItems returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("ShopFlashSale.UpdateShopFlashSaleItems response: %#v", res)
}
