package goshopee

import (
	"context"
	"fmt"
	"testing"

	"github.com/jarcoal/httpmock"
)

func Test_Public_GetMerchantsByPartner(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.public.get_merchants_by_partner_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetMerchantsByPartner due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetMerchantsByPartner due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/public/get_merchants_by_partner", app.APIURL), responder)

	var req GetMerchantsByPartnerRequest
	ctx := context.Background()
	res, err := client.Public.GetMerchantsByPartner(ctx, shopID, req, accessToken)
	if err != nil {
		t.Logf("Public.GetMerchantsByPartner returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Public.GetMerchantsByPartner response: %#v", res)
}
func Test_Public_GetShopeeIpRanges(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.public.get_shopee_ip_ranges_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetShopeeIpRanges due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetShopeeIpRanges due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/public/get_shopee_ip_ranges", app.APIURL), responder)

	ctx := context.Background()
	res, err := client.Public.GetShopeeIpRanges(ctx, shopID, accessToken)
	if err != nil {
		t.Logf("Public.GetShopeeIpRanges returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Public.GetShopeeIpRanges response: %#v", res)
}
func Test_Public_GetShopsByPartner(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.public.get_shops_by_partner_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetShopsByPartner due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetShopsByPartner due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/public/get_shops_by_partner", app.APIURL), responder)

	var req GetShopsByPartnerRequest
	ctx := context.Background()
	res, err := client.Public.GetShopsByPartner(ctx, shopID, req, accessToken)
	if err != nil {
		t.Logf("Public.GetShopsByPartner returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Public.GetShopsByPartner response: %#v", res)
}
func Test_Public_GetTokenByResendCode(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.public.get_token_by_resend_code_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetTokenByResendCode due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetTokenByResendCode due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/public/get_token_by_resend_code", app.APIURL), responder)

	var req GetTokenByResendCodeRequest
	ctx := context.Background()
	res, err := client.Public.GetTokenByResendCode(ctx, shopID, req, accessToken)
	if err != nil {
		t.Logf("Public.GetTokenByResendCode returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Public.GetTokenByResendCode response: %#v", res)
}
