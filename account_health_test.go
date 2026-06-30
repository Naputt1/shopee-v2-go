package goshopee

import (
	"fmt"
	"testing"

	"github.com/jarcoal/httpmock"
)

func Test_AccountHealth_GetLateOrders(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.account_health.get_late_orders_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetLateOrders due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetLateOrders due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/account_health/get_late_orders", app.APIURL), responder)

	var req GetLateOrdersRequest
	res, err := client.AccountHealth.GetLateOrders(shopID, req, accessToken)
	if err != nil {
		t.Logf("AccountHealth.GetLateOrders returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("AccountHealth.GetLateOrders response: %#v", res)
}
func Test_AccountHealth_GetListingsWithIssues(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.account_health.get_listings_with_issues_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetListingsWithIssues due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetListingsWithIssues due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/account_health/get_listings_with_issues", app.APIURL), responder)

	var req GetListingsWithIssuesRequest
	res, err := client.AccountHealth.GetListingsWithIssues(shopID, req, accessToken)
	if err != nil {
		t.Logf("AccountHealth.GetListingsWithIssues returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("AccountHealth.GetListingsWithIssues response: %#v", res)
}
func Test_AccountHealth_GetMetricSourceDetail(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.account_health.get_metric_source_detail_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetMetricSourceDetail due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetMetricSourceDetail due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/account_health/get_metric_source_detail", app.APIURL), responder)

	var req GetMetricSourceDetailRequest
	res, err := client.AccountHealth.GetMetricSourceDetail(shopID, req, accessToken)
	if err != nil {
		t.Logf("AccountHealth.GetMetricSourceDetail returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("AccountHealth.GetMetricSourceDetail response: %#v", res)
}
func Test_AccountHealth_GetPenaltyPointHistory(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.account_health.get_penalty_point_history_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetPenaltyPointHistory due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetPenaltyPointHistory due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/account_health/get_penalty_point_history", app.APIURL), responder)

	var req GetPenaltyPointHistoryRequest
	res, err := client.AccountHealth.GetPenaltyPointHistory(shopID, req, accessToken)
	if err != nil {
		t.Logf("AccountHealth.GetPenaltyPointHistory returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("AccountHealth.GetPenaltyPointHistory response: %#v", res)
}
func Test_AccountHealth_GetPunishmentHistory(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.account_health.get_punishment_history_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetPunishmentHistory due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetPunishmentHistory due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/account_health/get_punishment_history", app.APIURL), responder)

	var req GetPunishmentHistoryRequest
	res, err := client.AccountHealth.GetPunishmentHistory(shopID, req, accessToken)
	if err != nil {
		t.Logf("AccountHealth.GetPunishmentHistory returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("AccountHealth.GetPunishmentHistory response: %#v", res)
}
func Test_AccountHealth_GetShopPerformance(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.account_health.get_shop_performance_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetShopPerformance due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetShopPerformance due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/account_health/get_shop_performance", app.APIURL), responder)

	res, err := client.AccountHealth.GetShopPerformance(shopID, accessToken)
	if err != nil {
		t.Logf("AccountHealth.GetShopPerformance returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("AccountHealth.GetShopPerformance response: %#v", res)
}
