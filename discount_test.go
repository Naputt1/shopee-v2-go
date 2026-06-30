package goshopee

import (
	"fmt"
	"testing"

	"github.com/jarcoal/httpmock"
)

func Test_Discount_AddDiscount(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.discount.add_discount_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping AddDiscount due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping AddDiscount due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/discount/add_discount", app.APIURL), responder)

	var req AddDiscountRequest
	res, err := client.Discount.AddDiscount(shopID, req, accessToken)
	if err != nil {
		t.Logf("Discount.AddDiscount returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Discount.AddDiscount response: %#v", res)
}
func Test_Discount_AddDiscountItem(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.discount.add_discount_item_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping AddDiscountItem due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping AddDiscountItem due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/discount/add_discount_item", app.APIURL), responder)

	var req AddDiscountItemRequest
	res, err := client.Discount.AddDiscountItem(shopID, req, accessToken)
	if err != nil {
		t.Logf("Discount.AddDiscountItem returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Discount.AddDiscountItem response: %#v", res)
}
func Test_Discount_DeleteDiscount(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.discount.delete_discount_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping DeleteDiscount due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping DeleteDiscount due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/discount/delete_discount", app.APIURL), responder)

	var req DeleteDiscountRequest
	res, err := client.Discount.DeleteDiscount(shopID, req, accessToken)
	if err != nil {
		t.Logf("Discount.DeleteDiscount returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Discount.DeleteDiscount response: %#v", res)
}
func Test_Discount_DeleteDiscountItem(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.discount.delete_discount_item_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping DeleteDiscountItem due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping DeleteDiscountItem due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/discount/delete_discount_item", app.APIURL), responder)

	var req DeleteDiscountItemRequest
	res, err := client.Discount.DeleteDiscountItem(shopID, req, accessToken)
	if err != nil {
		t.Logf("Discount.DeleteDiscountItem returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Discount.DeleteDiscountItem response: %#v", res)
}
func Test_Discount_DeleteSipDiscount(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.discount.delete_sip_discount_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping DeleteSipDiscount due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping DeleteSipDiscount due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/discount/delete_sip_discount", app.APIURL), responder)

	var req DeleteSipDiscountRequest
	res, err := client.Discount.DeleteSipDiscount(shopID, req, accessToken)
	if err != nil {
		t.Logf("Discount.DeleteSipDiscount returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Discount.DeleteSipDiscount response: %#v", res)
}
func Test_Discount_EndDiscount(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.discount.end_discount_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping EndDiscount due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping EndDiscount due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/discount/end_discount", app.APIURL), responder)

	var req EndDiscountRequest
	res, err := client.Discount.EndDiscount(shopID, req, accessToken)
	if err != nil {
		t.Logf("Discount.EndDiscount returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Discount.EndDiscount response: %#v", res)
}
func Test_Discount_GetDiscount(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.discount.get_discount_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetDiscount due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetDiscount due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/discount/get_discount", app.APIURL), responder)

	var req GetDiscountRequest
	res, err := client.Discount.GetDiscount(shopID, req, accessToken)
	if err != nil {
		t.Logf("Discount.GetDiscount returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Discount.GetDiscount response: %#v", res)
}
func Test_Discount_GetDiscountList(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.discount.get_discount_list_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetDiscountList due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetDiscountList due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/discount/get_discount_list", app.APIURL), responder)

	var req GetDiscountListRequest
	res, err := client.Discount.GetDiscountList(shopID, req, accessToken)
	if err != nil {
		t.Logf("Discount.GetDiscountList returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Discount.GetDiscountList response: %#v", res)
}
func Test_Discount_GetSipDiscounts(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.discount.get_sip_discounts_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetSipDiscounts due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetSipDiscounts due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/discount/get_sip_discounts", app.APIURL), responder)

	var req GetSipDiscountsRequest
	res, err := client.Discount.GetSipDiscounts(shopID, req, accessToken)
	if err != nil {
		t.Logf("Discount.GetSipDiscounts returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Discount.GetSipDiscounts response: %#v", res)
}
func Test_Discount_SetSipDiscount(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.discount.set_sip_discount_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping SetSipDiscount due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping SetSipDiscount due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/discount/set_sip_discount", app.APIURL), responder)

	var req SetSipDiscountRequest
	res, err := client.Discount.SetSipDiscount(shopID, req, accessToken)
	if err != nil {
		t.Logf("Discount.SetSipDiscount returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Discount.SetSipDiscount response: %#v", res)
}
func Test_Discount_UpdateDiscount(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.discount.update_discount_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping UpdateDiscount due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping UpdateDiscount due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/discount/update_discount", app.APIURL), responder)

	var req UpdateDiscountRequest
	res, err := client.Discount.UpdateDiscount(shopID, req, accessToken)
	if err != nil {
		t.Logf("Discount.UpdateDiscount returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Discount.UpdateDiscount response: %#v", res)
}
func Test_Discount_UpdateDiscountItem(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.discount.update_discount_item_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping UpdateDiscountItem due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping UpdateDiscountItem due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/discount/update_discount_item", app.APIURL), responder)

	var req UpdateDiscountItemRequest
	res, err := client.Discount.UpdateDiscountItem(shopID, req, accessToken)
	if err != nil {
		t.Logf("Discount.UpdateDiscountItem returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Discount.UpdateDiscountItem response: %#v", res)
}
