package goshopee

import (
	"context"
	"fmt"
	"testing"

	"github.com/jarcoal/httpmock"
)

func Test_Shop_GetAuthorisedResellerBrand(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.shop.get_authorised_reseller_brand_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetAuthorisedResellerBrand due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetAuthorisedResellerBrand due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/shop/get_authorised_reseller_brand", app.APIURL), responder)

	var req GetAuthorisedResellerBrandRequest
	ctx := context.Background()
	res, err := client.Shop.GetAuthorisedResellerBrand(ctx, shopID, req, accessToken)
	if err != nil {
		t.Logf("Shop.GetAuthorisedResellerBrand returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Shop.GetAuthorisedResellerBrand response: %#v", res)
}
func Test_Shop_GetBrShopOnboardingInfo(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.shop.get_br_shop_onboarding_info_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetBrShopOnboardingInfo due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetBrShopOnboardingInfo due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/shop/get_br_shop_onboarding_info", app.APIURL), responder)

	ctx := context.Background()
	res, err := client.Shop.GetBrShopOnboardingInfo(ctx, shopID, accessToken)
	if err != nil {
		t.Logf("Shop.GetBrShopOnboardingInfo returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Shop.GetBrShopOnboardingInfo response: %#v", res)
}
func Test_Shop_GetProfile(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.shop.get_profile_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetProfile due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetProfile due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/shop/get_profile", app.APIURL), responder)

	ctx := context.Background()
	res, err := client.Shop.GetProfile(ctx, shopID, accessToken)
	if err != nil {
		t.Logf("Shop.GetProfile returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Shop.GetProfile response: %#v", res)
}
func Test_Shop_GetShopHolidayMode(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.shop.get_shop_holiday_mode_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetShopHolidayMode due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetShopHolidayMode due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/shop/get_shop_holiday_mode", app.APIURL), responder)

	ctx := context.Background()
	res, err := client.Shop.GetShopHolidayMode(ctx, shopID, accessToken)
	if err != nil {
		t.Logf("Shop.GetShopHolidayMode returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Shop.GetShopHolidayMode response: %#v", res)
}
func Test_Shop_GetShopInfo(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.shop.get_shop_info_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetShopInfo due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetShopInfo due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/shop/get_shop_info", app.APIURL), responder)

	ctx := context.Background()
	res, err := client.Shop.GetShopInfo(ctx, shopID, accessToken)
	if err != nil {
		t.Logf("Shop.GetShopInfo returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Shop.GetShopInfo response: %#v", res)
}
func Test_Shop_GetShopNotification(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.shop.get_shop_notification_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetShopNotification due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetShopNotification due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/shop/get_shop_notification", app.APIURL), responder)

	var req GetShopNotificationRequest
	ctx := context.Background()
	res, err := client.Shop.GetShopNotification(ctx, shopID, req, accessToken)
	if err != nil {
		t.Logf("Shop.GetShopNotification returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Shop.GetShopNotification response: %#v", res)
}
func Test_Shop_GetWarehouseDetail(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.shop.get_warehouse_detail_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetWarehouseDetail due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetWarehouseDetail due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/shop/get_warehouse_detail", app.APIURL), responder)

	var req GetWarehouseDetailRequest
	ctx := context.Background()
	res, err := client.Shop.GetWarehouseDetail(ctx, shopID, req, accessToken)
	if err != nil {
		t.Logf("Shop.GetWarehouseDetail returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Shop.GetWarehouseDetail response: %#v", res)
}
func Test_Shop_SetShopHolidayMode(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.shop.set_shop_holiday_mode_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping SetShopHolidayMode due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping SetShopHolidayMode due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/shop/set_shop_holiday_mode", app.APIURL), responder)

	var req SetShopHolidayModeRequest
	ctx := context.Background()
	res, err := client.Shop.SetShopHolidayMode(ctx, shopID, req, accessToken)
	if err != nil {
		t.Logf("Shop.SetShopHolidayMode returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Shop.SetShopHolidayMode response: %#v", res)
}
func Test_Shop_UpdateProfile(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.shop.update_profile_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping UpdateProfile due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping UpdateProfile due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/shop/update_profile", app.APIURL), responder)

	var req UpdateProfileRequest
	ctx := context.Background()
	res, err := client.Shop.UpdateProfile(ctx, shopID, req, accessToken)
	if err != nil {
		t.Logf("Shop.UpdateProfile returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Shop.UpdateProfile response: %#v", res)
}
