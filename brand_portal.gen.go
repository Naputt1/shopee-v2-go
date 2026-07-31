package goshopee

import (
	"context"
)

type BrandPortalService interface {
	// GetClipVideoPerformance {"content":"<p>Queries video clip performance data for the specified videos within the selected time range. Supports request granularity by day, week, month, quarter, year, or customize, and returns both overall summary metrics and video-level detailed metrics.</p>","raw_content":[{"name":"paragraph","children":[{"data":"Queries video clip performance data for the specified videos within the selected time range. Supports request granularity by day, week, month, quarter, year, or customize, and returns both overall summary metrics and video-level detailed metrics."}]}]}
	// Path: /api/v2/principal/get_clip_video_performance
	// https://open.shopee.com/documents/v2/v2.principal.get_clip_video_performance?module=139&type=1
	GetClipVideoPerformance(ctx context.Context, sid uint64, mid uint64, tok string, req GetClipVideoPerformanceRequest) (*GetClipVideoPerformanceResponse, error)
	// GetContentAffiliatePerformance {"content":"<p>Queries affiliate performance data for the specified content items within the selected time range. Supports request granularity by day, week, month, quarter, year, or customize, and returns both overall summary metrics and content-level detailed metrics with placed-order and confirmed-order views.</p>","raw_content":[{"name":"paragraph","children":[{"data":"Queries affiliate performance data for the specified content items within the selected time range. Supports request granularity by day, week, month, quarter, year, or customize, and returns both overall summary metrics and content-level detailed metrics with placed-order and confirmed-order views."}]}]}
	// Path: /api/v2/principal/get_content_affiliate_performance
	// https://open.shopee.com/documents/v2/v2.principal.get_content_affiliate_performance?module=139&type=1
	GetContentAffiliatePerformance(ctx context.Context, sid uint64, mid uint64, tok string, req GetContentAffiliatePerformanceRequest) (*GetContentAffiliatePerformanceResponse, error)
	// GetPrincipalAffiliatePerformance {"content":"<p>Queries affiliate performance data for the specified principal within the selected time range. Supports request granularity by day, week, month, quarter, year, or customize, and returns both overall summary metrics and region-level detailed metrics with placed-order and confirmed-order views.</p>","raw_content":[{"name":"paragraph","children":[{"data":"Queries affiliate performance data for the specified principal within the selected time range. Supports request granularity by day, week, month, quarter, year, or customize, and returns both overall summary metrics and region-level detailed metrics with placed-order and confirmed-order views."}]}]}
	// Path: /api/v2/principal/get_principal_affiliate_performance
	// https://open.shopee.com/documents/v2/v2.principal.get_principal_affiliate_performance?module=139&type=1
	GetPrincipalAffiliatePerformance(ctx context.Context, sid uint64, mid uint64, tok string, req GetPrincipalAffiliatePerformanceRequest) (*GetPrincipalAffiliatePerformanceResponse, error)
	// GetPrincipalLivestreamPerformance {"content":"<p>Queries livestream performance data for the specified principal within the selected time range. Supports request granularity by day, week, month, quarter, year, or customize, and returns both overall summary metrics and region-level detailed metrics.</p>","raw_content":[{"name":"paragraph","children":[{"data":"Queries livestream performance data for the specified principal within the selected time range. Supports request granularity by day, week, month, quarter, year, or customize, and returns both overall summary metrics and region-level detailed metrics."}]}]}
	// Path: /api/v2/principal/get_principal_livestream_performance
	// https://open.shopee.com/documents/v2/v2.principal.get_principal_livestream_performance?module=139&type=1
	GetPrincipalLivestreamPerformance(ctx context.Context, sid uint64, mid uint64, tok string, req GetPrincipalLivestreamPerformanceRequest) (*GetPrincipalLivestreamPerformanceResponse, error)
	// GetPrincipalSalesPerformanceDetail {"content":"<p>Queries the business performance data aggregated at principal level for the specified regions within the selected time range. Supports request granularity by day, week, month, quarter, year, or customize, and returns both overall summary metrics and region-level detailed metrics.</p>","raw_content":[{"name":"paragraph","children":[{"data":"Queries the business performance data aggregated at principal level for the specified regions within the selected time range. Supports request granularity by day, week, month, quarter, year, or customize, and returns both overall summary metrics and region-level detailed metrics."}]}]}
	// Path: /api/v2/principal/get_principal_sales_performance_detail
	// https://open.shopee.com/documents/v2/v2.principal.get_principal_sales_performance_detail?module=139&type=1
	GetPrincipalSalesPerformanceDetail(ctx context.Context, sid uint64, mid uint64, tok string, req GetPrincipalSalesPerformanceDetailRequest) (*GetPrincipalSalesPerformanceDetailResponse, error)
	// GetPrincipalVideoPerformance {"content":"<p>Queries video performance data for the specified principal within the selected time range. Supports request granularity by day, week, month, quarter, year, or customize, and returns both overall summary metrics and region-level detailed metrics.</p>","raw_content":[{"name":"paragraph","children":[{"data":"Queries video performance data for the specified principal within the selected time range. Supports request granularity by day, week, month, quarter, year, or customize, and returns both overall summary metrics and region-level detailed metrics."}]}]}
	// Path: /api/v2/principal/get_principal_video_performance
	// https://open.shopee.com/documents/v2/v2.principal.get_principal_video_performance?module=139&type=1
	GetPrincipalVideoPerformance(ctx context.Context, sid uint64, mid uint64, tok string, req GetPrincipalVideoPerformanceRequest) (*GetPrincipalVideoPerformanceResponse, error)
	// GetSessionLivestreamPerformance {"content":"<p>Queries livestream session performance data for the specified sessions within the selected time range. Supports request granularity by day, week, month, quarter, year, or customize, and returns both overall summary metrics and session-level detailed metrics.</p>","raw_content":[{"name":"paragraph","children":[{"data":"Queries livestream session performance data for the specified sessions within the selected time range. Supports request granularity by day, week, month, quarter, year, or customize, and returns both overall summary metrics and session-level detailed metrics."}]}]}
	// Path: /api/v2/principal/get_session_livestream_performance
	// https://open.shopee.com/documents/v2/v2.principal.get_session_livestream_performance?module=139&type=1
	GetSessionLivestreamPerformance(ctx context.Context, sid uint64, mid uint64, tok string, req GetSessionLivestreamPerformanceRequest) (*GetSessionLivestreamPerformanceResponse, error)
	// GetShopAffiliatePerformance {"content":"<p>Queries affiliate performance data for the specified shops within the selected time range. Supports request granularity by day, week, month, quarter, year, or customize, and returns both overall summary metrics and shop-level detailed metrics with placed-order and confirmed-order views.</p>","raw_content":[{"name":"paragraph","children":[{"data":"Queries affiliate performance data for the specified shops within the selected time range. Supports request granularity by day, week, month, quarter, year, or customize, and returns both overall summary metrics and shop-level detailed metrics with placed-order and confirmed-order views."}]}]}
	// Path: /api/v2/principal/get_shop_affiliate_performance
	// https://open.shopee.com/documents/v2/v2.principal.get_shop_affiliate_performance?module=139&type=1
	GetShopAffiliatePerformance(ctx context.Context, sid uint64, mid uint64, tok string, req GetShopAffiliatePerformanceRequest) (*GetShopAffiliatePerformanceResponse, error)
	// GetShopLivestreamPerformance {"content":"<p>Queries livestream performance data for the specified shops within the selected time range. Supports request granularity by day, week, month, quarter, year, or customize, and returns both overall summary metrics and shop-level detailed metrics.</p>","raw_content":[{"name":"paragraph","children":[{"data":"Queries livestream performance data for the specified shops within the selected time range. Supports request granularity by day, week, month, quarter, year, or customize, and returns both overall summary metrics and shop-level detailed metrics."}]}]}
	// Path: /api/v2/principal/get_shop_livestream_performance
	// https://open.shopee.com/documents/v2/v2.principal.get_shop_livestream_performance?module=139&type=1
	GetShopLivestreamPerformance(ctx context.Context, sid uint64, mid uint64, tok string, req GetShopLivestreamPerformanceRequest) (*GetShopLivestreamPerformanceResponse, error)
	// GetShopSalesPerformanceDetail {"content":"<p>Queries the business performance data of stores under the specified entity within the selected time range. Supports request granularity by day, week, month, quarter, year, or customize, and returns both overall summary metrics and store-level detailed metrics.</p>","raw_content":[{"name":"paragraph","children":[{"data":"Queries the business performance data of stores under the specified entity within the selected time range. Supports request granularity by day, week, month, quarter, year, or customize, and returns both overall summary metrics and store-level detailed metrics."}]}]}
	// Path: /api/v2/principal/get_shop_sales_performance_detail
	// https://open.shopee.com/documents/v2/v2.principal.get_shop_sales_performance_detail?module=139&type=1
	GetShopSalesPerformanceDetail(ctx context.Context, sid uint64, mid uint64, tok string, req GetShopSalesPerformanceDetailRequest) (*GetShopSalesPerformanceDetailResponse, error)
	// GetShopVideoPerformance {"content":"<p>Queries video performance data for the specified shops within the selected time range. Supports request granularity by day, week, month, quarter, year, or customize, and returns both overall summary metrics and shop-level detailed metrics.</p>","raw_content":[{"name":"paragraph","children":[{"data":"Queries video performance data for the specified shops within the selected time range. Supports request granularity by day, week, month, quarter, year, or customize, and returns both overall summary metrics and shop-level detailed metrics."}]}]}
	// Path: /api/v2/principal/get_shop_video_performance
	// https://open.shopee.com/documents/v2/v2.principal.get_shop_video_performance?module=139&type=1
	GetShopVideoPerformance(ctx context.Context, sid uint64, mid uint64, tok string, req GetShopVideoPerformanceRequest) (*GetShopVideoPerformanceResponse, error)
}

type BrandPortalServiceOp[T any] struct {
	client *Client[T]
}

// GetClipVideoPerformance {"content":"<p>Queries video clip performance data for the specified videos within the selected time range. Supports request granularity by day, week, month, quarter, year, or customize, and returns both overall summary metrics and video-level detailed metrics.</p>","raw_content":[{"name":"paragraph","children":[{"data":"Queries video clip performance data for the specified videos within the selected time range. Supports request granularity by day, week, month, quarter, year, or customize, and returns both overall summary metrics and video-level detailed metrics."}]}]}
// Path: /api/v2/principal/get_clip_video_performance
// https://open.shopee.com/documents/v2/v2.principal.get_clip_video_performance?module=139&type=1
func (s *BrandPortalServiceOp[T]) GetClipVideoPerformance(ctx context.Context, sid uint64, mid uint64, tok string, req GetClipVideoPerformanceRequest) (*GetClipVideoPerformanceResponse, error) {
	path := "/principal/get_clip_video_performance"
	resp := new(GetClipVideoPerformanceResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// GetContentAffiliatePerformance {"content":"<p>Queries affiliate performance data for the specified content items within the selected time range. Supports request granularity by day, week, month, quarter, year, or customize, and returns both overall summary metrics and content-level detailed metrics with placed-order and confirmed-order views.</p>","raw_content":[{"name":"paragraph","children":[{"data":"Queries affiliate performance data for the specified content items within the selected time range. Supports request granularity by day, week, month, quarter, year, or customize, and returns both overall summary metrics and content-level detailed metrics with placed-order and confirmed-order views."}]}]}
// Path: /api/v2/principal/get_content_affiliate_performance
// https://open.shopee.com/documents/v2/v2.principal.get_content_affiliate_performance?module=139&type=1
func (s *BrandPortalServiceOp[T]) GetContentAffiliatePerformance(ctx context.Context, sid uint64, mid uint64, tok string, req GetContentAffiliatePerformanceRequest) (*GetContentAffiliatePerformanceResponse, error) {
	path := "/principal/get_content_affiliate_performance"
	resp := new(GetContentAffiliatePerformanceResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// GetPrincipalAffiliatePerformance {"content":"<p>Queries affiliate performance data for the specified principal within the selected time range. Supports request granularity by day, week, month, quarter, year, or customize, and returns both overall summary metrics and region-level detailed metrics with placed-order and confirmed-order views.</p>","raw_content":[{"name":"paragraph","children":[{"data":"Queries affiliate performance data for the specified principal within the selected time range. Supports request granularity by day, week, month, quarter, year, or customize, and returns both overall summary metrics and region-level detailed metrics with placed-order and confirmed-order views."}]}]}
// Path: /api/v2/principal/get_principal_affiliate_performance
// https://open.shopee.com/documents/v2/v2.principal.get_principal_affiliate_performance?module=139&type=1
func (s *BrandPortalServiceOp[T]) GetPrincipalAffiliatePerformance(ctx context.Context, sid uint64, mid uint64, tok string, req GetPrincipalAffiliatePerformanceRequest) (*GetPrincipalAffiliatePerformanceResponse, error) {
	path := "/principal/get_principal_affiliate_performance"
	resp := new(GetPrincipalAffiliatePerformanceResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// GetPrincipalLivestreamPerformance {"content":"<p>Queries livestream performance data for the specified principal within the selected time range. Supports request granularity by day, week, month, quarter, year, or customize, and returns both overall summary metrics and region-level detailed metrics.</p>","raw_content":[{"name":"paragraph","children":[{"data":"Queries livestream performance data for the specified principal within the selected time range. Supports request granularity by day, week, month, quarter, year, or customize, and returns both overall summary metrics and region-level detailed metrics."}]}]}
// Path: /api/v2/principal/get_principal_livestream_performance
// https://open.shopee.com/documents/v2/v2.principal.get_principal_livestream_performance?module=139&type=1
func (s *BrandPortalServiceOp[T]) GetPrincipalLivestreamPerformance(ctx context.Context, sid uint64, mid uint64, tok string, req GetPrincipalLivestreamPerformanceRequest) (*GetPrincipalLivestreamPerformanceResponse, error) {
	path := "/principal/get_principal_livestream_performance"
	resp := new(GetPrincipalLivestreamPerformanceResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// GetPrincipalSalesPerformanceDetail {"content":"<p>Queries the business performance data aggregated at principal level for the specified regions within the selected time range. Supports request granularity by day, week, month, quarter, year, or customize, and returns both overall summary metrics and region-level detailed metrics.</p>","raw_content":[{"name":"paragraph","children":[{"data":"Queries the business performance data aggregated at principal level for the specified regions within the selected time range. Supports request granularity by day, week, month, quarter, year, or customize, and returns both overall summary metrics and region-level detailed metrics."}]}]}
// Path: /api/v2/principal/get_principal_sales_performance_detail
// https://open.shopee.com/documents/v2/v2.principal.get_principal_sales_performance_detail?module=139&type=1
func (s *BrandPortalServiceOp[T]) GetPrincipalSalesPerformanceDetail(ctx context.Context, sid uint64, mid uint64, tok string, req GetPrincipalSalesPerformanceDetailRequest) (*GetPrincipalSalesPerformanceDetailResponse, error) {
	path := "/principal/get_principal_sales_performance_detail"
	resp := new(GetPrincipalSalesPerformanceDetailResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// GetPrincipalVideoPerformance {"content":"<p>Queries video performance data for the specified principal within the selected time range. Supports request granularity by day, week, month, quarter, year, or customize, and returns both overall summary metrics and region-level detailed metrics.</p>","raw_content":[{"name":"paragraph","children":[{"data":"Queries video performance data for the specified principal within the selected time range. Supports request granularity by day, week, month, quarter, year, or customize, and returns both overall summary metrics and region-level detailed metrics."}]}]}
// Path: /api/v2/principal/get_principal_video_performance
// https://open.shopee.com/documents/v2/v2.principal.get_principal_video_performance?module=139&type=1
func (s *BrandPortalServiceOp[T]) GetPrincipalVideoPerformance(ctx context.Context, sid uint64, mid uint64, tok string, req GetPrincipalVideoPerformanceRequest) (*GetPrincipalVideoPerformanceResponse, error) {
	path := "/principal/get_principal_video_performance"
	resp := new(GetPrincipalVideoPerformanceResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// GetSessionLivestreamPerformance {"content":"<p>Queries livestream session performance data for the specified sessions within the selected time range. Supports request granularity by day, week, month, quarter, year, or customize, and returns both overall summary metrics and session-level detailed metrics.</p>","raw_content":[{"name":"paragraph","children":[{"data":"Queries livestream session performance data for the specified sessions within the selected time range. Supports request granularity by day, week, month, quarter, year, or customize, and returns both overall summary metrics and session-level detailed metrics."}]}]}
// Path: /api/v2/principal/get_session_livestream_performance
// https://open.shopee.com/documents/v2/v2.principal.get_session_livestream_performance?module=139&type=1
func (s *BrandPortalServiceOp[T]) GetSessionLivestreamPerformance(ctx context.Context, sid uint64, mid uint64, tok string, req GetSessionLivestreamPerformanceRequest) (*GetSessionLivestreamPerformanceResponse, error) {
	path := "/principal/get_session_livestream_performance"
	resp := new(GetSessionLivestreamPerformanceResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// GetShopAffiliatePerformance {"content":"<p>Queries affiliate performance data for the specified shops within the selected time range. Supports request granularity by day, week, month, quarter, year, or customize, and returns both overall summary metrics and shop-level detailed metrics with placed-order and confirmed-order views.</p>","raw_content":[{"name":"paragraph","children":[{"data":"Queries affiliate performance data for the specified shops within the selected time range. Supports request granularity by day, week, month, quarter, year, or customize, and returns both overall summary metrics and shop-level detailed metrics with placed-order and confirmed-order views."}]}]}
// Path: /api/v2/principal/get_shop_affiliate_performance
// https://open.shopee.com/documents/v2/v2.principal.get_shop_affiliate_performance?module=139&type=1
func (s *BrandPortalServiceOp[T]) GetShopAffiliatePerformance(ctx context.Context, sid uint64, mid uint64, tok string, req GetShopAffiliatePerformanceRequest) (*GetShopAffiliatePerformanceResponse, error) {
	path := "/principal/get_shop_affiliate_performance"
	resp := new(GetShopAffiliatePerformanceResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// GetShopLivestreamPerformance {"content":"<p>Queries livestream performance data for the specified shops within the selected time range. Supports request granularity by day, week, month, quarter, year, or customize, and returns both overall summary metrics and shop-level detailed metrics.</p>","raw_content":[{"name":"paragraph","children":[{"data":"Queries livestream performance data for the specified shops within the selected time range. Supports request granularity by day, week, month, quarter, year, or customize, and returns both overall summary metrics and shop-level detailed metrics."}]}]}
// Path: /api/v2/principal/get_shop_livestream_performance
// https://open.shopee.com/documents/v2/v2.principal.get_shop_livestream_performance?module=139&type=1
func (s *BrandPortalServiceOp[T]) GetShopLivestreamPerformance(ctx context.Context, sid uint64, mid uint64, tok string, req GetShopLivestreamPerformanceRequest) (*GetShopLivestreamPerformanceResponse, error) {
	path := "/principal/get_shop_livestream_performance"
	resp := new(GetShopLivestreamPerformanceResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// GetShopSalesPerformanceDetail {"content":"<p>Queries the business performance data of stores under the specified entity within the selected time range. Supports request granularity by day, week, month, quarter, year, or customize, and returns both overall summary metrics and store-level detailed metrics.</p>","raw_content":[{"name":"paragraph","children":[{"data":"Queries the business performance data of stores under the specified entity within the selected time range. Supports request granularity by day, week, month, quarter, year, or customize, and returns both overall summary metrics and store-level detailed metrics."}]}]}
// Path: /api/v2/principal/get_shop_sales_performance_detail
// https://open.shopee.com/documents/v2/v2.principal.get_shop_sales_performance_detail?module=139&type=1
func (s *BrandPortalServiceOp[T]) GetShopSalesPerformanceDetail(ctx context.Context, sid uint64, mid uint64, tok string, req GetShopSalesPerformanceDetailRequest) (*GetShopSalesPerformanceDetailResponse, error) {
	path := "/principal/get_shop_sales_performance_detail"
	resp := new(GetShopSalesPerformanceDetailResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// GetShopVideoPerformance {"content":"<p>Queries video performance data for the specified shops within the selected time range. Supports request granularity by day, week, month, quarter, year, or customize, and returns both overall summary metrics and shop-level detailed metrics.</p>","raw_content":[{"name":"paragraph","children":[{"data":"Queries video performance data for the specified shops within the selected time range. Supports request granularity by day, week, month, quarter, year, or customize, and returns both overall summary metrics and shop-level detailed metrics."}]}]}
// Path: /api/v2/principal/get_shop_video_performance
// https://open.shopee.com/documents/v2/v2.principal.get_shop_video_performance?module=139&type=1
func (s *BrandPortalServiceOp[T]) GetShopVideoPerformance(ctx context.Context, sid uint64, mid uint64, tok string, req GetShopVideoPerformanceRequest) (*GetShopVideoPerformanceResponse, error) {
	path := "/principal/get_shop_video_performance"
	resp := new(GetShopVideoPerformanceResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}
