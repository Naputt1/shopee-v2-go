package goshopee

import (
	"context"
)

type AdsService interface {
	// CheckCreateGmsProductCampaignEligibility Check the seller's eligibility in creating a GMS campaign
	// Path: /api/v2/ads/check_create_gms_product_campaign_eligibility
	// https://open.shopee.com/documents/v2/v2.ads.check_create_gms_product_campaign_eligibility?module=117&type=1
	CheckCreateGmsProductCampaignEligibility(ctx context.Context, sid uint64, mid uint64, tok string) (*CheckCreateGmsProductCampaignEligibilityResponse, error)
	// CreateGmsProductCampaign Create a GMS campaign
	// Path: /api/v2/ads/create_gms_product_campaign
	// https://open.shopee.com/documents/v2/v2.ads.create_gms_product_campaign?module=117&type=1
	CreateGmsProductCampaign(ctx context.Context, sid uint64, mid uint64, tok string, req CreateGmsProductCampaignRequest) (*CreateGmsProductCampaignResponse, error)
	// CreateManualProductAds Use this API to create Manual Selection Product Ads
	// Path: /api/v2/ads/create_manual_product_ads
	// https://open.shopee.com/documents/v2/v2.ads.create_manual_product_ads?module=117&type=1
	CreateManualProductAds(ctx context.Context, sid uint64, mid uint64, tok string, req CreateManualProductAdsRequest) (*CreateManualProductAdsResponse, error)
	// EditGmsItemProductCampaign Add/remove items to/from the GMS Campaign
	// Path: /api/v2/ads/edit_gms_item_product_campaign
	// https://open.shopee.com/documents/v2/v2.ads.edit_gms_item_product_campaign?module=117&type=1
	EditGmsItemProductCampaign(ctx context.Context, sid uint64, mid uint64, tok string, req EditGmsItemProductCampaignRequest) (*EditGmsItemProductCampaignResponse, error)
	// EditGmsProductCampaign Edit a GMS campaign
	// Path: /api/v2/ads/edit_gms_product_campaign
	// https://open.shopee.com/documents/v2/v2.ads.edit_gms_product_campaign?module=117&type=1
	EditGmsProductCampaign(ctx context.Context, sid uint64, mid uint64, tok string, req EditGmsProductCampaignRequest) (*EditGmsProductCampaignResponse, error)
	// EditManualProductAdKeywords Use this API to edit Manual Selection Product Ad Keywords
	// Path: /api/v2/ads/edit_manual_product_ad_keywords
	// https://open.shopee.com/documents/v2/v2.ads.edit_manual_product_ad_keywords?module=117&type=1
	EditManualProductAdKeywords(ctx context.Context, sid uint64, mid uint64, tok string, req EditManualProductAdKeywordsRequest) (*EditManualProductAdKeywordsResponse, error)
	// EditManualProductAds Use this API to edit Manual Selection Product Ads
	// Path: /api/v2/ads/edit_manual_product_ads
	// https://open.shopee.com/documents/v2/v2.ads.edit_manual_product_ads?module=117&type=1
	EditManualProductAds(ctx context.Context, sid uint64, mid uint64, tok string, req EditManualProductAdsRequest) (*EditManualProductAdsResponse, error)
	// GetAdsFcilShopRate Get shop rate for Ads Facil Program
	// Path: /api/v2/ads/get_ads_facil_shop_rate
	// https://open.shopee.com/documents/v2/v2.ads.get_ads_fácil_shop_rate?module=117&type=1
	GetAdsFcilShopRate(ctx context.Context, sid uint64, mid uint64, tok string) (*GetAdsFcilShopRateResponse, error)
	// GetAllCpcAdsDailyPerformance Use this API to get Shop level CPC ads multiple-days daily performance.
	// Path: /api/v2/ads/get_all_cpc_ads_daily_performance
	// https://open.shopee.com/documents/v2/v2.ads.get_all_cpc_ads_daily_performance?module=117&type=1
	GetAllCpcAdsDailyPerformance(ctx context.Context, sid uint64, mid uint64, tok string, opt GetAllCpcAdsDailyPerformanceRequest) (*GetAllCpcAdsDailyPerformanceResponse, error)
	// GetAllCpcAdsHourlyPerformance Use this API to get Shop level CPC ads single-date hourly performance.
	// Path: /api/v2/ads/get_all_cpc_ads_hourly_performance
	// https://open.shopee.com/documents/v2/v2.ads.get_all_cpc_ads_hourly_performance?module=117&type=1
	GetAllCpcAdsHourlyPerformance(ctx context.Context, sid uint64, mid uint64, tok string, opt GetAllCpcAdsHourlyPerformanceRequest) (*GetAllCpcAdsHourlyPerformanceResponse, error)
	// GetCreateProductAdBudgetSuggestion Call this API to get budget suggestion for product ads creation
	// Path: /api/v2/ads/get_create_product_ad_budget_suggestion
	// https://open.shopee.com/documents/v2/v2.ads.get_create_product_ad_budget_suggestion?module=117&type=1
	GetCreateProductAdBudgetSuggestion(ctx context.Context, sid uint64, mid uint64, tok string, opt GetCreateProductAdBudgetSuggestionRequest) (*GetCreateProductAdBudgetSuggestionResponse, error)
	// GetGmsCampaignPerformance Get GMS Campaign performance
	// Path: /api/v2/ads/get_gms_campaign_performance
	// https://open.shopee.com/documents/v2/v2.ads.get_gms_campaign_performance?module=117&type=1
	GetGmsCampaignPerformance(ctx context.Context, sid uint64, mid uint64, tok string, req GetGmsCampaignPerformanceRequest) (*GetGmsCampaignPerformanceResponse, error)
	// GetGmsItemPerformance Get GMS Item performance
	// 1. The response returned is sorted by item_id
	// 2. Only items with performance will be returned
	// Path: /api/v2/ads/get_gms_item_performance
	// https://open.shopee.com/documents/v2/v2.ads.get_gms_item_performance?module=117&type=1
	GetGmsItemPerformance(ctx context.Context, sid uint64, mid uint64, tok string, req GetGmsItemPerformanceRequest) (*GetGmsItemPerformanceResponse, error)
	// GetProductCampaignDailyPerformance Use this API to get Product level ads multiple-days daily performance.
	// Path: /api/v2/ads/get_product_campaign_daily_performance
	// https://open.shopee.com/documents/v2/v2.ads.get_product_campaign_daily_performance?module=117&type=1
	GetProductCampaignDailyPerformance(ctx context.Context, sid uint64, mid uint64, tok string, opt GetProductCampaignDailyPerformanceRequest) (*GetProductCampaignDailyPerformanceResponse, error)
	// GetProductCampaignHourlyPerformance Use this API to get Product level ads single-day hourly performance.
	// Path: /api/v2/ads/get_product_campaign_hourly_performance
	// https://open.shopee.com/documents/v2/v2.ads.get_product_campaign_hourly_performance?module=117&type=1
	GetProductCampaignHourlyPerformance(ctx context.Context, sid uint64, mid uint64, tok string, opt GetProductCampaignHourlyPerformanceRequest) (*GetProductCampaignHourlyPerformanceResponse, error)
	// GetProductLevelCampaignIdList Call this API to fetch all the product campaign ids displayed on advertiser platform under a specific Shop
	// Path: /api/v2/ads/get_product_level_campaign_id_list
	// https://open.shopee.com/documents/v2/v2.ads.get_product_level_campaign_id_list?module=117&type=1
	GetProductLevelCampaignIdList(ctx context.Context, sid uint64, mid uint64, tok string, opt GetProductLevelCampaignIdListRequest) (*GetProductLevelCampaignIdListResponse, error)
	// GetProductLevelCampaignSettingInfo Call this API to fetch all the campaign setting info under this Shop.
	// Path: /api/v2/ads/get_product_level_campaign_setting_info
	// https://open.shopee.com/documents/v2/v2.ads.get_product_level_campaign_setting_info?module=117&type=1
	GetProductLevelCampaignSettingInfo(ctx context.Context, sid uint64, mid uint64, tok string, opt GetProductLevelCampaignSettingInfoRequest) (*GetProductLevelCampaignSettingInfoResponse, error)
	// GetProductRecommendedRoiTarget Get Product Recommended ROI Target
	// Path: /api/v2/ads/get_product_recommended_roi_target
	// https://open.shopee.com/documents/v2/v2.ads.get_product_recommended_roi_target?module=117&type=1
	GetProductRecommendedRoiTarget(ctx context.Context, sid uint64, mid uint64, tok string, opt GetProductRecommendedRoiTargetRequest) (*GetProductRecommendedRoiTargetResponse, error)
	// GetRecommendedItemList Use this API to get the list of recommended SKU (Shop level) with the corresponding tag, i.e top search/best selling/best ROI tag.
	// Path: /api/v2/ads/get_recommended_item_list
	// https://open.shopee.com/documents/v2/v2.ads.get_recommended_item_list?module=117&type=1
	GetRecommendedItemList(ctx context.Context, sid uint64, mid uint64, tok string) (*GetRecommendedItemListResponse, error)
	// GetRecommendedKeywordList Use this API to get the list of Recommended keywords by item and optionally a search keyword
	// Path: /api/v2/ads/get_recommended_keyword_list
	// https://open.shopee.com/documents/v2/v2.ads.get_recommended_keyword_list?module=117&type=1
	GetRecommendedKeywordList(ctx context.Context, sid uint64, mid uint64, tok string, opt GetRecommendedKeywordListRequest) (*GetRecommendedKeywordListResponse, error)
	// GetShopToggleInfo Use this API to get Shop level info - i.e. seller's toggle status is on/off
	// Path: /api/v2/ads/get_shop_toggle_info
	// https://open.shopee.com/documents/v2/v2.ads.get_shop_toggle_info?module=117&type=1
	GetShopToggleInfo(ctx context.Context, sid uint64, mid uint64, tok string) (*GetShopToggleInfoResponse, error)
	// GetTotalBalance Use this API to return the seller's Real-time total balance of their ads credit including the paid credits and free credits.
	// Path: /api/v2/ads/get_total_balance
	// https://open.shopee.com/documents/v2/v2.ads.get_total_balance?module=117&type=1
	GetTotalBalance(ctx context.Context, sid uint64, mid uint64, tok string) (*GetTotalBalanceResponse, error)
	// ListGmsUserDeletedItem List GMS items that have been removed from the Campaign by seller
	// Path: /api/v2/ads/list_gms_user_deleted_item
	// https://open.shopee.com/documents/v2/v2.ads.list_gms_user_deleted_item?module=117&type=1
	ListGmsUserDeletedItem(ctx context.Context, sid uint64, mid uint64, tok string, req ListGmsUserDeletedItemRequest) (*ListGmsUserDeletedItemResponse, error)
}

type AdsServiceOp[T any] struct {
	client *Client[T]
}

// CheckCreateGmsProductCampaignEligibility Check the seller's eligibility in creating a GMS campaign
// Path: /api/v2/ads/check_create_gms_product_campaign_eligibility
// https://open.shopee.com/documents/v2/v2.ads.check_create_gms_product_campaign_eligibility?module=117&type=1
func (s *AdsServiceOp[T]) CheckCreateGmsProductCampaignEligibility(ctx context.Context, sid uint64, mid uint64, tok string) (*CheckCreateGmsProductCampaignEligibilityResponse, error) {
	path := "/ads/check_create_gms_product_campaign_eligibility"
	resp := new(CheckCreateGmsProductCampaignEligibilityResponse)
	err := s.client.Get(ctx, path, resp, nil, sid, mid, tok)
	return resp, err
}

// CreateGmsProductCampaign Create a GMS campaign
// Path: /api/v2/ads/create_gms_product_campaign
// https://open.shopee.com/documents/v2/v2.ads.create_gms_product_campaign?module=117&type=1
func (s *AdsServiceOp[T]) CreateGmsProductCampaign(ctx context.Context, sid uint64, mid uint64, tok string, req CreateGmsProductCampaignRequest) (*CreateGmsProductCampaignResponse, error) {
	path := "/ads/create_gms_product_campaign"
	resp := new(CreateGmsProductCampaignResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// CreateManualProductAds Use this API to create Manual Selection Product Ads
// Path: /api/v2/ads/create_manual_product_ads
// https://open.shopee.com/documents/v2/v2.ads.create_manual_product_ads?module=117&type=1
func (s *AdsServiceOp[T]) CreateManualProductAds(ctx context.Context, sid uint64, mid uint64, tok string, req CreateManualProductAdsRequest) (*CreateManualProductAdsResponse, error) {
	path := "/ads/create_manual_product_ads"
	resp := new(CreateManualProductAdsResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// EditGmsItemProductCampaign Add/remove items to/from the GMS Campaign
// Path: /api/v2/ads/edit_gms_item_product_campaign
// https://open.shopee.com/documents/v2/v2.ads.edit_gms_item_product_campaign?module=117&type=1
func (s *AdsServiceOp[T]) EditGmsItemProductCampaign(ctx context.Context, sid uint64, mid uint64, tok string, req EditGmsItemProductCampaignRequest) (*EditGmsItemProductCampaignResponse, error) {
	path := "/ads/edit_gms_item_product_campaign"
	resp := new(EditGmsItemProductCampaignResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// EditGmsProductCampaign Edit a GMS campaign
// Path: /api/v2/ads/edit_gms_product_campaign
// https://open.shopee.com/documents/v2/v2.ads.edit_gms_product_campaign?module=117&type=1
func (s *AdsServiceOp[T]) EditGmsProductCampaign(ctx context.Context, sid uint64, mid uint64, tok string, req EditGmsProductCampaignRequest) (*EditGmsProductCampaignResponse, error) {
	path := "/ads/edit_gms_product_campaign"
	resp := new(EditGmsProductCampaignResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// EditManualProductAdKeywords Use this API to edit Manual Selection Product Ad Keywords
// Path: /api/v2/ads/edit_manual_product_ad_keywords
// https://open.shopee.com/documents/v2/v2.ads.edit_manual_product_ad_keywords?module=117&type=1
func (s *AdsServiceOp[T]) EditManualProductAdKeywords(ctx context.Context, sid uint64, mid uint64, tok string, req EditManualProductAdKeywordsRequest) (*EditManualProductAdKeywordsResponse, error) {
	path := "/ads/edit_manual_product_ad_keywords"
	resp := new(EditManualProductAdKeywordsResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// EditManualProductAds Use this API to edit Manual Selection Product Ads
// Path: /api/v2/ads/edit_manual_product_ads
// https://open.shopee.com/documents/v2/v2.ads.edit_manual_product_ads?module=117&type=1
func (s *AdsServiceOp[T]) EditManualProductAds(ctx context.Context, sid uint64, mid uint64, tok string, req EditManualProductAdsRequest) (*EditManualProductAdsResponse, error) {
	path := "/ads/edit_manual_product_ads"
	resp := new(EditManualProductAdsResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// GetAdsFcilShopRate Get shop rate for Ads Facil Program
// Path: /api/v2/ads/get_ads_facil_shop_rate
// https://open.shopee.com/documents/v2/v2.ads.get_ads_fácil_shop_rate?module=117&type=1
func (s *AdsServiceOp[T]) GetAdsFcilShopRate(ctx context.Context, sid uint64, mid uint64, tok string) (*GetAdsFcilShopRateResponse, error) {
	path := "/ads/get_ads_facil_shop_rate"
	resp := new(GetAdsFcilShopRateResponse)
	err := s.client.Post(ctx, path, nil, resp, sid, mid, tok)
	return resp, err
}

// GetAllCpcAdsDailyPerformance Use this API to get Shop level CPC ads multiple-days daily performance.
// Path: /api/v2/ads/get_all_cpc_ads_daily_performance
// https://open.shopee.com/documents/v2/v2.ads.get_all_cpc_ads_daily_performance?module=117&type=1
func (s *AdsServiceOp[T]) GetAllCpcAdsDailyPerformance(ctx context.Context, sid uint64, mid uint64, tok string, opt GetAllCpcAdsDailyPerformanceRequest) (*GetAllCpcAdsDailyPerformanceResponse, error) {
	path := "/ads/get_all_cpc_ads_daily_performance"
	resp := new(GetAllCpcAdsDailyPerformanceResponse)
	err := s.client.Get(ctx, path, resp, opt, sid, mid, tok)
	return resp, err
}

// GetAllCpcAdsHourlyPerformance Use this API to get Shop level CPC ads single-date hourly performance.
// Path: /api/v2/ads/get_all_cpc_ads_hourly_performance
// https://open.shopee.com/documents/v2/v2.ads.get_all_cpc_ads_hourly_performance?module=117&type=1
func (s *AdsServiceOp[T]) GetAllCpcAdsHourlyPerformance(ctx context.Context, sid uint64, mid uint64, tok string, opt GetAllCpcAdsHourlyPerformanceRequest) (*GetAllCpcAdsHourlyPerformanceResponse, error) {
	path := "/ads/get_all_cpc_ads_hourly_performance"
	resp := new(GetAllCpcAdsHourlyPerformanceResponse)
	err := s.client.Get(ctx, path, resp, opt, sid, mid, tok)
	return resp, err
}

// GetCreateProductAdBudgetSuggestion Call this API to get budget suggestion for product ads creation
// Path: /api/v2/ads/get_create_product_ad_budget_suggestion
// https://open.shopee.com/documents/v2/v2.ads.get_create_product_ad_budget_suggestion?module=117&type=1
func (s *AdsServiceOp[T]) GetCreateProductAdBudgetSuggestion(ctx context.Context, sid uint64, mid uint64, tok string, opt GetCreateProductAdBudgetSuggestionRequest) (*GetCreateProductAdBudgetSuggestionResponse, error) {
	path := "/ads/get_create_product_ad_budget_suggestion"
	resp := new(GetCreateProductAdBudgetSuggestionResponse)
	err := s.client.Get(ctx, path, resp, opt, sid, mid, tok)
	return resp, err
}

// GetGmsCampaignPerformance Get GMS Campaign performance
// Path: /api/v2/ads/get_gms_campaign_performance
// https://open.shopee.com/documents/v2/v2.ads.get_gms_campaign_performance?module=117&type=1
func (s *AdsServiceOp[T]) GetGmsCampaignPerformance(ctx context.Context, sid uint64, mid uint64, tok string, req GetGmsCampaignPerformanceRequest) (*GetGmsCampaignPerformanceResponse, error) {
	path := "/ads/get_gms_campaign_performance"
	resp := new(GetGmsCampaignPerformanceResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// GetGmsItemPerformance Get GMS Item performance
// 1. The response returned is sorted by item_id
// 2. Only items with performance will be returned
// Path: /api/v2/ads/get_gms_item_performance
// https://open.shopee.com/documents/v2/v2.ads.get_gms_item_performance?module=117&type=1
func (s *AdsServiceOp[T]) GetGmsItemPerformance(ctx context.Context, sid uint64, mid uint64, tok string, req GetGmsItemPerformanceRequest) (*GetGmsItemPerformanceResponse, error) {
	path := "/ads/get_gms_item_performance"
	resp := new(GetGmsItemPerformanceResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// GetProductCampaignDailyPerformance Use this API to get Product level ads multiple-days daily performance.
// Path: /api/v2/ads/get_product_campaign_daily_performance
// https://open.shopee.com/documents/v2/v2.ads.get_product_campaign_daily_performance?module=117&type=1
func (s *AdsServiceOp[T]) GetProductCampaignDailyPerformance(ctx context.Context, sid uint64, mid uint64, tok string, opt GetProductCampaignDailyPerformanceRequest) (*GetProductCampaignDailyPerformanceResponse, error) {
	path := "/ads/get_product_campaign_daily_performance"
	resp := new(GetProductCampaignDailyPerformanceResponse)
	err := s.client.Get(ctx, path, resp, opt, sid, mid, tok)
	return resp, err
}

// GetProductCampaignHourlyPerformance Use this API to get Product level ads single-day hourly performance.
// Path: /api/v2/ads/get_product_campaign_hourly_performance
// https://open.shopee.com/documents/v2/v2.ads.get_product_campaign_hourly_performance?module=117&type=1
func (s *AdsServiceOp[T]) GetProductCampaignHourlyPerformance(ctx context.Context, sid uint64, mid uint64, tok string, opt GetProductCampaignHourlyPerformanceRequest) (*GetProductCampaignHourlyPerformanceResponse, error) {
	path := "/ads/get_product_campaign_hourly_performance"
	resp := new(GetProductCampaignHourlyPerformanceResponse)
	err := s.client.Get(ctx, path, resp, opt, sid, mid, tok)
	return resp, err
}

// GetProductLevelCampaignIdList Call this API to fetch all the product campaign ids displayed on advertiser platform under a specific Shop
// Path: /api/v2/ads/get_product_level_campaign_id_list
// https://open.shopee.com/documents/v2/v2.ads.get_product_level_campaign_id_list?module=117&type=1
func (s *AdsServiceOp[T]) GetProductLevelCampaignIdList(ctx context.Context, sid uint64, mid uint64, tok string, opt GetProductLevelCampaignIdListRequest) (*GetProductLevelCampaignIdListResponse, error) {
	path := "/ads/get_product_level_campaign_id_list"
	resp := new(GetProductLevelCampaignIdListResponse)
	err := s.client.Get(ctx, path, resp, opt, sid, mid, tok)
	return resp, err
}

// GetProductLevelCampaignSettingInfo Call this API to fetch all the campaign setting info under this Shop.
// Path: /api/v2/ads/get_product_level_campaign_setting_info
// https://open.shopee.com/documents/v2/v2.ads.get_product_level_campaign_setting_info?module=117&type=1
func (s *AdsServiceOp[T]) GetProductLevelCampaignSettingInfo(ctx context.Context, sid uint64, mid uint64, tok string, opt GetProductLevelCampaignSettingInfoRequest) (*GetProductLevelCampaignSettingInfoResponse, error) {
	path := "/ads/get_product_level_campaign_setting_info"
	resp := new(GetProductLevelCampaignSettingInfoResponse)
	err := s.client.Get(ctx, path, resp, opt, sid, mid, tok)
	return resp, err
}

// GetProductRecommendedRoiTarget Get Product Recommended ROI Target
// Path: /api/v2/ads/get_product_recommended_roi_target
// https://open.shopee.com/documents/v2/v2.ads.get_product_recommended_roi_target?module=117&type=1
func (s *AdsServiceOp[T]) GetProductRecommendedRoiTarget(ctx context.Context, sid uint64, mid uint64, tok string, opt GetProductRecommendedRoiTargetRequest) (*GetProductRecommendedRoiTargetResponse, error) {
	path := "/ads/get_product_recommended_roi_target"
	resp := new(GetProductRecommendedRoiTargetResponse)
	err := s.client.Get(ctx, path, resp, opt, sid, mid, tok)
	return resp, err
}

// GetRecommendedItemList Use this API to get the list of recommended SKU (Shop level) with the corresponding tag, i.e top search/best selling/best ROI tag.
// Path: /api/v2/ads/get_recommended_item_list
// https://open.shopee.com/documents/v2/v2.ads.get_recommended_item_list?module=117&type=1
func (s *AdsServiceOp[T]) GetRecommendedItemList(ctx context.Context, sid uint64, mid uint64, tok string) (*GetRecommendedItemListResponse, error) {
	path := "/ads/get_recommended_item_list"
	resp := new(GetRecommendedItemListResponse)
	err := s.client.Get(ctx, path, resp, nil, sid, mid, tok)
	return resp, err
}

// GetRecommendedKeywordList Use this API to get the list of Recommended keywords by item and optionally a search keyword
// Path: /api/v2/ads/get_recommended_keyword_list
// https://open.shopee.com/documents/v2/v2.ads.get_recommended_keyword_list?module=117&type=1
func (s *AdsServiceOp[T]) GetRecommendedKeywordList(ctx context.Context, sid uint64, mid uint64, tok string, opt GetRecommendedKeywordListRequest) (*GetRecommendedKeywordListResponse, error) {
	path := "/ads/get_recommended_keyword_list"
	resp := new(GetRecommendedKeywordListResponse)
	err := s.client.Get(ctx, path, resp, opt, sid, mid, tok)
	return resp, err
}

// GetShopToggleInfo Use this API to get Shop level info - i.e. seller's toggle status is on/off
// Path: /api/v2/ads/get_shop_toggle_info
// https://open.shopee.com/documents/v2/v2.ads.get_shop_toggle_info?module=117&type=1
func (s *AdsServiceOp[T]) GetShopToggleInfo(ctx context.Context, sid uint64, mid uint64, tok string) (*GetShopToggleInfoResponse, error) {
	path := "/ads/get_shop_toggle_info"
	resp := new(GetShopToggleInfoResponse)
	err := s.client.Get(ctx, path, resp, nil, sid, mid, tok)
	return resp, err
}

// GetTotalBalance Use this API to return the seller's Real-time total balance of their ads credit including the paid credits and free credits.
// Path: /api/v2/ads/get_total_balance
// https://open.shopee.com/documents/v2/v2.ads.get_total_balance?module=117&type=1
func (s *AdsServiceOp[T]) GetTotalBalance(ctx context.Context, sid uint64, mid uint64, tok string) (*GetTotalBalanceResponse, error) {
	path := "/ads/get_total_balance"
	resp := new(GetTotalBalanceResponse)
	err := s.client.Get(ctx, path, resp, nil, sid, mid, tok)
	return resp, err
}

// ListGmsUserDeletedItem List GMS items that have been removed from the Campaign by seller
// Path: /api/v2/ads/list_gms_user_deleted_item
// https://open.shopee.com/documents/v2/v2.ads.list_gms_user_deleted_item?module=117&type=1
func (s *AdsServiceOp[T]) ListGmsUserDeletedItem(ctx context.Context, sid uint64, mid uint64, tok string, req ListGmsUserDeletedItemRequest) (*ListGmsUserDeletedItemResponse, error) {
	path := "/ads/list_gms_user_deleted_item"
	resp := new(ListGmsUserDeletedItemResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}
