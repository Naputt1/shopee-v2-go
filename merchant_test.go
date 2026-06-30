package goshopee

import (
	"context"
	"fmt"
	"testing"

	"github.com/jarcoal/httpmock"
)

func Test_Merchant_GetMerchantInfo(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.merchant.get_merchant_info_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetMerchantInfo due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetMerchantInfo due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/merchant/get_merchant_info", app.APIURL), responder)

	ctx := context.Background()
	res, err := client.Merchant.GetMerchantInfo(ctx, shopID, accessToken)
	if err != nil {
		t.Logf("Merchant.GetMerchantInfo returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Merchant.GetMerchantInfo response: %#v", res)
}
func Test_Merchant_GetMerchantPrepaidAccountList(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.merchant.get_merchant_prepaid_account_list_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetMerchantPrepaidAccountList due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetMerchantPrepaidAccountList due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/merchant/get_merchant_prepaid_account_list", app.APIURL), responder)

	var req GetMerchantPrepaidAccountListRequest
	ctx := context.Background()
	res, err := client.Merchant.GetMerchantPrepaidAccountList(ctx, shopID, req, accessToken)
	if err != nil {
		t.Logf("Merchant.GetMerchantPrepaidAccountList returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Merchant.GetMerchantPrepaidAccountList response: %#v", res)
}
func Test_Merchant_GetMerchantWarehouseList(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.merchant.get_merchant_warehouse_list_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetMerchantWarehouseList due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetMerchantWarehouseList due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/merchant/get_merchant_warehouse_list", app.APIURL), responder)

	var req GetMerchantWarehouseListRequest
	ctx := context.Background()
	res, err := client.Merchant.GetMerchantWarehouseList(ctx, shopID, req, accessToken)
	if err != nil {
		t.Logf("Merchant.GetMerchantWarehouseList returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Merchant.GetMerchantWarehouseList response: %#v", res)
}
func Test_Merchant_GetMerchantWarehouseLocationList(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.merchant.get_merchant_warehouse_location_list_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetMerchantWarehouseLocationList due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetMerchantWarehouseLocationList due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/merchant/get_merchant_warehouse_location_list", app.APIURL), responder)

	ctx := context.Background()
	res, err := client.Merchant.GetMerchantWarehouseLocationList(ctx, shopID, accessToken)
	if err != nil {
		t.Logf("Merchant.GetMerchantWarehouseLocationList returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Merchant.GetMerchantWarehouseLocationList response: %#v", res)
}
func Test_Merchant_GetShopListByMerchant(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.merchant.get_shop_list_by_merchant_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetShopListByMerchant due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetShopListByMerchant due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/merchant/get_shop_list_by_merchant", app.APIURL), responder)

	var req GetShopListByMerchantRequest
	ctx := context.Background()
	res, err := client.Merchant.GetShopListByMerchant(ctx, shopID, req, accessToken)
	if err != nil {
		t.Logf("Merchant.GetShopListByMerchant returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Merchant.GetShopListByMerchant response: %#v", res)
}
func Test_Merchant_GetWarehouseEligibleShopList(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.merchant.get_warehouse_eligible_shop_list_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetWarehouseEligibleShopList due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetWarehouseEligibleShopList due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/merchant/get_warehouse_eligible_shop_list", app.APIURL), responder)

	var req GetWarehouseEligibleShopListRequest
	ctx := context.Background()
	res, err := client.Merchant.GetWarehouseEligibleShopList(ctx, shopID, req, accessToken)
	if err != nil {
		t.Logf("Merchant.GetWarehouseEligibleShopList returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Merchant.GetWarehouseEligibleShopList response: %#v", res)
}
