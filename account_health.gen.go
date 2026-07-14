package goshopee

import (
	"context"
)

type AccountHealthService interface {
	// GetLateOrders Get the Late Orders to take action to avoid order cancellation and penalty points.
	// Path: /api/v2/account_health/get_late_orders
	// https://open.shopee.com/documents/v2/v2.account_health.get_late_orders?module=103&type=1
	GetLateOrders(ctx context.Context, sid uint64, mid uint64, tok string, opt GetLateOrdersRequest) (*GetLateOrdersResponse, error)
	// GetListingsWithIssues Get the Problematic Listings to improve the listings to avoid incurring penalty points.
	// Path: /api/v2/account_health/get_listings_with_issues
	// https://open.shopee.com/documents/v2/v2.account_health.get_listings_with_issues?module=103&type=1
	GetListingsWithIssues(ctx context.Context, sid uint64, mid uint64, tok string, opt GetListingsWithIssuesRequest) (*GetListingsWithIssuesResponse, error)
	// GetMetricSourceDetail {"content":"<p>Get the Affected Orders / Relevant Listings / Relevant Violations details of metrics.</p>","raw_content":[{"name":"paragraph","children":[{"data":"Get the Affected Orders / Relevant Listings / Relevant Violations details of metrics."}]}]}
	// Path: /api/v2/account_health/get_metric_source_detail
	// https://open.shopee.com/documents/v2/v2.account_health.get_metric_source_detail?module=103&type=1
	GetMetricSourceDetail(ctx context.Context, sid uint64, mid uint64, tok string, opt GetMetricSourceDetailRequest) (*GetMetricSourceDetailResponse, error)
	// GetPenaltyPointHistory Get the penalty point records generated in the current quarter.
	// Path: /api/v2/account_health/get_penalty_point_history
	// https://open.shopee.com/documents/v2/v2.account_health.get_penalty_point_history?module=103&type=1
	GetPenaltyPointHistory(ctx context.Context, sid uint64, mid uint64, tok string, opt GetPenaltyPointHistoryRequest) (*GetPenaltyPointHistoryResponse, error)
	// GetPunishmentHistory Get the punishment records generated in the current quarter.
	// Path: /api/v2/account_health/get_punishment_history
	// https://open.shopee.com/documents/v2/v2.account_health.get_punishment_history?module=103&type=1
	GetPunishmentHistory(ctx context.Context, sid uint64, mid uint64, tok string, opt GetPunishmentHistoryRequest) (*GetPunishmentHistoryResponse, error)
	// GetShopPerformance {"content":"<p>The data metrics of shop performance.<br>&nbsp;</p>","raw_content":[{"name":"paragraph","children":[{"data":"The data metrics of shop performance."},{"name":"softBreak"},{"data":" "}]}]}
	// Path: /api/v2/account_health/get_shop_performance
	// https://open.shopee.com/documents/v2/v2.account_health.get_shop_performance?module=103&type=1
	GetShopPerformance(ctx context.Context, sid uint64, mid uint64, tok string) (*GetShopPerformanceResponse, error)
}

type AccountHealthServiceOp[T any] struct {
	client *Client[T]
}

// GetLateOrders Get the Late Orders to take action to avoid order cancellation and penalty points.
// Path: /api/v2/account_health/get_late_orders
// https://open.shopee.com/documents/v2/v2.account_health.get_late_orders?module=103&type=1
func (s *AccountHealthServiceOp[T]) GetLateOrders(ctx context.Context, sid uint64, mid uint64, tok string, opt GetLateOrdersRequest) (*GetLateOrdersResponse, error) {
	path := "/account_health/get_late_orders"
	resp := new(GetLateOrdersResponse)
	err := s.client.Get(ctx, path, resp, opt, sid, mid, tok)
	return resp, err
}

// GetListingsWithIssues Get the Problematic Listings to improve the listings to avoid incurring penalty points.
// Path: /api/v2/account_health/get_listings_with_issues
// https://open.shopee.com/documents/v2/v2.account_health.get_listings_with_issues?module=103&type=1
func (s *AccountHealthServiceOp[T]) GetListingsWithIssues(ctx context.Context, sid uint64, mid uint64, tok string, opt GetListingsWithIssuesRequest) (*GetListingsWithIssuesResponse, error) {
	path := "/account_health/get_listings_with_issues"
	resp := new(GetListingsWithIssuesResponse)
	err := s.client.Get(ctx, path, resp, opt, sid, mid, tok)
	return resp, err
}

// GetMetricSourceDetail {"content":"<p>Get the Affected Orders / Relevant Listings / Relevant Violations details of metrics.</p>","raw_content":[{"name":"paragraph","children":[{"data":"Get the Affected Orders / Relevant Listings / Relevant Violations details of metrics."}]}]}
// Path: /api/v2/account_health/get_metric_source_detail
// https://open.shopee.com/documents/v2/v2.account_health.get_metric_source_detail?module=103&type=1
func (s *AccountHealthServiceOp[T]) GetMetricSourceDetail(ctx context.Context, sid uint64, mid uint64, tok string, opt GetMetricSourceDetailRequest) (*GetMetricSourceDetailResponse, error) {
	path := "/account_health/get_metric_source_detail"
	resp := new(GetMetricSourceDetailResponse)
	err := s.client.Get(ctx, path, resp, opt, sid, mid, tok)
	return resp, err
}

// GetPenaltyPointHistory Get the penalty point records generated in the current quarter.
// Path: /api/v2/account_health/get_penalty_point_history
// https://open.shopee.com/documents/v2/v2.account_health.get_penalty_point_history?module=103&type=1
func (s *AccountHealthServiceOp[T]) GetPenaltyPointHistory(ctx context.Context, sid uint64, mid uint64, tok string, opt GetPenaltyPointHistoryRequest) (*GetPenaltyPointHistoryResponse, error) {
	path := "/account_health/get_penalty_point_history"
	resp := new(GetPenaltyPointHistoryResponse)
	err := s.client.Get(ctx, path, resp, opt, sid, mid, tok)
	return resp, err
}

// GetPunishmentHistory Get the punishment records generated in the current quarter.
// Path: /api/v2/account_health/get_punishment_history
// https://open.shopee.com/documents/v2/v2.account_health.get_punishment_history?module=103&type=1
func (s *AccountHealthServiceOp[T]) GetPunishmentHistory(ctx context.Context, sid uint64, mid uint64, tok string, opt GetPunishmentHistoryRequest) (*GetPunishmentHistoryResponse, error) {
	path := "/account_health/get_punishment_history"
	resp := new(GetPunishmentHistoryResponse)
	err := s.client.Get(ctx, path, resp, opt, sid, mid, tok)
	return resp, err
}

// GetShopPerformance {"content":"<p>The data metrics of shop performance.<br>&nbsp;</p>","raw_content":[{"name":"paragraph","children":[{"data":"The data metrics of shop performance."},{"name":"softBreak"},{"data":" "}]}]}
// Path: /api/v2/account_health/get_shop_performance
// https://open.shopee.com/documents/v2/v2.account_health.get_shop_performance?module=103&type=1
func (s *AccountHealthServiceOp[T]) GetShopPerformance(ctx context.Context, sid uint64, mid uint64, tok string) (*GetShopPerformanceResponse, error) {
	path := "/account_health/get_shop_performance"
	resp := new(GetShopPerformanceResponse)
	err := s.client.Get(ctx, path, resp, nil, sid, mid, tok)
	return resp, err
}
