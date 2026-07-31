package goshopee

import (
	"context"
	"fmt"
	"testing"

	"github.com/jarcoal/httpmock"
)

func Test_BrandPortal_GetClipVideoPerformance(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.principal.get_clip_video_performance_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetClipVideoPerformance due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetClipVideoPerformance due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/principal/get_clip_video_performance", app.APIURL), responder)

	var req GetClipVideoPerformanceRequest
	ctx := context.Background()
	res, err := client.BrandPortal.GetClipVideoPerformance(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("BrandPortal.GetClipVideoPerformance returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("BrandPortal.GetClipVideoPerformance response: %#v", res)
}
func Test_BrandPortal_GetContentAffiliatePerformance(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.principal.get_content_affiliate_performance_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetContentAffiliatePerformance due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetContentAffiliatePerformance due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/principal/get_content_affiliate_performance", app.APIURL), responder)

	var req GetContentAffiliatePerformanceRequest
	ctx := context.Background()
	res, err := client.BrandPortal.GetContentAffiliatePerformance(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("BrandPortal.GetContentAffiliatePerformance returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("BrandPortal.GetContentAffiliatePerformance response: %#v", res)
}
func Test_BrandPortal_GetPrincipalAffiliatePerformance(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.principal.get_principal_affiliate_performance_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetPrincipalAffiliatePerformance due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetPrincipalAffiliatePerformance due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/principal/get_principal_affiliate_performance", app.APIURL), responder)

	var req GetPrincipalAffiliatePerformanceRequest
	ctx := context.Background()
	res, err := client.BrandPortal.GetPrincipalAffiliatePerformance(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("BrandPortal.GetPrincipalAffiliatePerformance returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("BrandPortal.GetPrincipalAffiliatePerformance response: %#v", res)
}
func Test_BrandPortal_GetPrincipalLivestreamPerformance(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.principal.get_principal_livestream_performance_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetPrincipalLivestreamPerformance due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetPrincipalLivestreamPerformance due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/principal/get_principal_livestream_performance", app.APIURL), responder)

	var req GetPrincipalLivestreamPerformanceRequest
	ctx := context.Background()
	res, err := client.BrandPortal.GetPrincipalLivestreamPerformance(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("BrandPortal.GetPrincipalLivestreamPerformance returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("BrandPortal.GetPrincipalLivestreamPerformance response: %#v", res)
}
func Test_BrandPortal_GetPrincipalSalesPerformanceDetail(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.principal.get_principal_sales_performance_detail_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetPrincipalSalesPerformanceDetail due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetPrincipalSalesPerformanceDetail due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/principal/get_principal_sales_performance_detail", app.APIURL), responder)

	var req GetPrincipalSalesPerformanceDetailRequest
	ctx := context.Background()
	res, err := client.BrandPortal.GetPrincipalSalesPerformanceDetail(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("BrandPortal.GetPrincipalSalesPerformanceDetail returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("BrandPortal.GetPrincipalSalesPerformanceDetail response: %#v", res)
}
func Test_BrandPortal_GetPrincipalVideoPerformance(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.principal.get_principal_video_performance_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetPrincipalVideoPerformance due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetPrincipalVideoPerformance due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/principal/get_principal_video_performance", app.APIURL), responder)

	var req GetPrincipalVideoPerformanceRequest
	ctx := context.Background()
	res, err := client.BrandPortal.GetPrincipalVideoPerformance(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("BrandPortal.GetPrincipalVideoPerformance returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("BrandPortal.GetPrincipalVideoPerformance response: %#v", res)
}
func Test_BrandPortal_GetSessionLivestreamPerformance(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.principal.get_session_livestream_performance_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetSessionLivestreamPerformance due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetSessionLivestreamPerformance due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/principal/get_session_livestream_performance", app.APIURL), responder)

	var req GetSessionLivestreamPerformanceRequest
	ctx := context.Background()
	res, err := client.BrandPortal.GetSessionLivestreamPerformance(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("BrandPortal.GetSessionLivestreamPerformance returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("BrandPortal.GetSessionLivestreamPerformance response: %#v", res)
}
func Test_BrandPortal_GetShopAffiliatePerformance(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.principal.get_shop_affiliate_performance_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetShopAffiliatePerformance due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetShopAffiliatePerformance due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/principal/get_shop_affiliate_performance", app.APIURL), responder)

	var req GetShopAffiliatePerformanceRequest
	ctx := context.Background()
	res, err := client.BrandPortal.GetShopAffiliatePerformance(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("BrandPortal.GetShopAffiliatePerformance returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("BrandPortal.GetShopAffiliatePerformance response: %#v", res)
}
func Test_BrandPortal_GetShopLivestreamPerformance(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.principal.get_shop_livestream_performance_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetShopLivestreamPerformance due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetShopLivestreamPerformance due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/principal/get_shop_livestream_performance", app.APIURL), responder)

	var req GetShopLivestreamPerformanceRequest
	ctx := context.Background()
	res, err := client.BrandPortal.GetShopLivestreamPerformance(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("BrandPortal.GetShopLivestreamPerformance returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("BrandPortal.GetShopLivestreamPerformance response: %#v", res)
}
func Test_BrandPortal_GetShopSalesPerformanceDetail(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.principal.get_shop_sales_performance_detail_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetShopSalesPerformanceDetail due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetShopSalesPerformanceDetail due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/principal/get_shop_sales_performance_detail", app.APIURL), responder)

	var req GetShopSalesPerformanceDetailRequest
	ctx := context.Background()
	res, err := client.BrandPortal.GetShopSalesPerformanceDetail(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("BrandPortal.GetShopSalesPerformanceDetail returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("BrandPortal.GetShopSalesPerformanceDetail response: %#v", res)
}
func Test_BrandPortal_GetShopVideoPerformance(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.principal.get_shop_video_performance_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetShopVideoPerformance due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetShopVideoPerformance due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/principal/get_shop_video_performance", app.APIURL), responder)

	var req GetShopVideoPerformanceRequest
	ctx := context.Background()
	res, err := client.BrandPortal.GetShopVideoPerformance(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("BrandPortal.GetShopVideoPerformance returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("BrandPortal.GetShopVideoPerformance response: %#v", res)
}
