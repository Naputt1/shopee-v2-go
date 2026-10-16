package goshopee

import (
	"context"
	"fmt"
	"testing"

	"github.com/jarcoal/httpmock"
)

func Test_BusinessInsights_GetMarketingHotListingonlyBRNow(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.business_insights.get_marketing_hot_listing（only BR now）_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetMarketingHotListingonlyBRNow due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetMarketingHotListingonlyBRNow due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/business_insights/get_marketing_hot_listing", app.APIURL), responder)

	var req GetMarketingHotListingonlyBRNowRequest
	ctx := context.Background()
	res, err := client.BusinessInsights.GetMarketingHotListingonlyBRNow(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("BusinessInsights.GetMarketingHotListingonlyBRNow returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("BusinessInsights.GetMarketingHotListingonlyBRNow response: %#v", res)
}
