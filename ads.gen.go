package goshopee

type AdsService interface {
	// CheckCreateGmsProductCampaignEligibility Check the seller's eligibility in creating a GMS campaign
	// Path: /api/v2/ads/check_create_gms_product_campaign_eligibility
	// https://open.shopee.com/documents/v2/v2.ads.check_create_gms_product_campaign_eligibility?module=117&type=1
	CheckCreateGmsProductCampaignEligibility(sid uint64, tok string) (*CheckCreateGmsProductCampaignEligibilityResponse, error)
	// CreateGmsProductCampaign Create a GMS campaign
	// Path: /api/v2/ads/create_gms_product_campaign
	// https://open.shopee.com/documents/v2/v2.ads.create_gms_product_campaign?module=117&type=1
	CreateGmsProductCampaign(sid uint64, req CreateGmsProductCampaignRequest, tok string) (*CreateGmsProductCampaignResponse, error)
	// CreateManualProductAds Use this API to create Manual Selection Product Ads
	// Path: /api/v2/ads/create_manual_product_ads
	// https://open.shopee.com/documents/v2/v2.ads.create_manual_product_ads?module=117&type=1
	CreateManualProductAds(sid uint64, req CreateManualProductAdsRequest, tok string) (*CreateManualProductAdsResponse, error)
	// EditGmsItemProductCampaign Add/remove items to/from the GMS Campaign
	// Path: /api/v2/ads/edit_gms_item_product_campaign
	// https://open.shopee.com/documents/v2/v2.ads.edit_gms_item_product_campaign?module=117&type=1
	EditGmsItemProductCampaign(sid uint64, req EditGmsItemProductCampaignRequest, tok string) (*EditGmsItemProductCampaignResponse, error)
	// EditGmsProductCampaign Edit a GMS campaign
	// Path: /api/v2/ads/edit_gms_product_campaign
	// https://open.shopee.com/documents/v2/v2.ads.edit_gms_product_campaign?module=117&type=1
	EditGmsProductCampaign(sid uint64, req EditGmsProductCampaignRequest, tok string) (*EditGmsProductCampaignResponse, error)
	// EditManualProductAdKeywords Use this API to edit Manual Selection Product Ad Keywords
	// Path: /api/v2/ads/edit_manual_product_ad_keywords
	// https://open.shopee.com/documents/v2/v2.ads.edit_manual_product_ad_keywords?module=117&type=1
	EditManualProductAdKeywords(sid uint64, req EditManualProductAdKeywordsRequest, tok string) (*EditManualProductAdKeywordsResponse, error)
	// EditManualProductAds Use this API to edit Manual Selection Product Ads
	// Path: /api/v2/ads/edit_manual_product_ads
	// https://open.shopee.com/documents/v2/v2.ads.edit_manual_product_ads?module=117&type=1
	EditManualProductAds(sid uint64, req EditManualProductAdsRequest, tok string) (*EditManualProductAdsResponse, error)
	// GetAdsFcilShopRate Get shop rate for Ads Facil Program
	// Path: /api/v2/ads/get_ads_facil_shop_rate
	// https://open.shopee.com/documents/v2/v2.ads.get_ads_fácil_shop_rate?module=117&type=1
	GetAdsFcilShopRate(sid uint64, tok string) (*GetAdsFcilShopRateResponse, error)
	// GetAllCpcAdsDailyPerformance Use this API to get Shop level CPC ads multiple-days daily performance.
	// Path: /api/v2/ads/get_all_cpc_ads_daily_performance
	// https://open.shopee.com/documents/v2/v2.ads.get_all_cpc_ads_daily_performance?module=117&type=1
	GetAllCpcAdsDailyPerformance(sid uint64, opt GetAllCpcAdsDailyPerformanceRequest, tok string) (*GetAllCpcAdsDailyPerformanceResponse, error)
	// GetAllCpcAdsHourlyPerformance Use this API to get Shop level CPC ads single-date hourly performance.
	// Path: /api/v2/ads/get_all_cpc_ads_hourly_performance
	// https://open.shopee.com/documents/v2/v2.ads.get_all_cpc_ads_hourly_performance?module=117&type=1
	GetAllCpcAdsHourlyPerformance(sid uint64, opt GetAllCpcAdsHourlyPerformanceRequest, tok string) (*GetAllCpcAdsHourlyPerformanceResponse, error)
	// GetCreateProductAdBudgetSuggestion Call this API to get budget suggestion for product ads creation
	// Path: /api/v2/ads/get_create_product_ad_budget_suggestion
	// https://open.shopee.com/documents/v2/v2.ads.get_create_product_ad_budget_suggestion?module=117&type=1
	GetCreateProductAdBudgetSuggestion(sid uint64, opt GetCreateProductAdBudgetSuggestionRequest, tok string) (*GetCreateProductAdBudgetSuggestionResponse, error)
	// GetGmsCampaignPerformance Get GMS Campaign performance
	// Path: /api/v2/ads/get_gms_campaign_performance
	// https://open.shopee.com/documents/v2/v2.ads.get_gms_campaign_performance?module=117&type=1
	GetGmsCampaignPerformance(sid uint64, req GetGmsCampaignPerformanceRequest, tok string) (*GetGmsCampaignPerformanceResponse, error)
	// GetGmsItemPerformance Get GMS Item performance
	// 1. The response returned is sorted by item_id
	// 2. Only items with performance will be returned
	// Path: /api/v2/ads/get_gms_item_performance
	// https://open.shopee.com/documents/v2/v2.ads.get_gms_item_performance?module=117&type=1
	GetGmsItemPerformance(sid uint64, req GetGmsItemPerformanceRequest, tok string) (*GetGmsItemPerformanceResponse, error)
	// GetProductCampaignDailyPerformance Use this API to get Product level ads multiple-days daily performance.
	// Path: /api/v2/ads/get_product_campaign_daily_performance
	// https://open.shopee.com/documents/v2/v2.ads.get_product_campaign_daily_performance?module=117&type=1
	GetProductCampaignDailyPerformance(sid uint64, opt GetProductCampaignDailyPerformanceRequest, tok string) (*GetProductCampaignDailyPerformanceResponse, error)
	// GetProductCampaignHourlyPerformance Use this API to get Product level ads single-day hourly performance.
	// Path: /api/v2/ads/get_product_campaign_hourly_performance
	// https://open.shopee.com/documents/v2/v2.ads.get_product_campaign_hourly_performance?module=117&type=1
	GetProductCampaignHourlyPerformance(sid uint64, opt GetProductCampaignHourlyPerformanceRequest, tok string) (*GetProductCampaignHourlyPerformanceResponse, error)
	// GetProductLevelCampaignIdList Call this API to fetch all the product campaign ids displayed on advertiser platform under a specific Shop
	// Path: /api/v2/ads/get_product_level_campaign_id_list
	// https://open.shopee.com/documents/v2/v2.ads.get_product_level_campaign_id_list?module=117&type=1
	GetProductLevelCampaignIdList(sid uint64, opt GetProductLevelCampaignIdListRequest, tok string) (*GetProductLevelCampaignIdListResponse, error)
	// GetProductLevelCampaignSettingInfo Call this API to fetch all the campaign setting info under this Shop.
	// Path: /api/v2/ads/get_product_level_campaign_setting_info
	// https://open.shopee.com/documents/v2/v2.ads.get_product_level_campaign_setting_info?module=117&type=1
	GetProductLevelCampaignSettingInfo(sid uint64, opt GetProductLevelCampaignSettingInfoRequest, tok string) (*GetProductLevelCampaignSettingInfoResponse, error)
	// GetProductRecommendedRoiTarget Get Product Recommended ROI Target
	// Path: /api/v2/ads/get_product_recommended_roi_target
	// https://open.shopee.com/documents/v2/v2.ads.get_product_recommended_roi_target?module=117&type=1
	GetProductRecommendedRoiTarget(sid uint64, opt GetProductRecommendedRoiTargetRequest, tok string) (*GetProductRecommendedRoiTargetResponse, error)
	// GetRecommendedItemList Use this API to get the list of recommended SKU (Shop level) with the corresponding tag, i.e top search/best selling/best ROI tag.
	// Path: /api/v2/ads/get_recommended_item_list
	// https://open.shopee.com/documents/v2/v2.ads.get_recommended_item_list?module=117&type=1
	GetRecommendedItemList(sid uint64, tok string) (*GetRecommendedItemListResponse, error)
	// GetRecommendedKeywordList Use this API to get the list of Recommended keywords by item and optionally a search keyword
	// Path: /api/v2/ads/get_recommended_keyword_list
	// https://open.shopee.com/documents/v2/v2.ads.get_recommended_keyword_list?module=117&type=1
	GetRecommendedKeywordList(sid uint64, opt GetRecommendedKeywordListRequest, tok string) (*GetRecommendedKeywordListResponse, error)
	// GetShopToggleInfo Use this API to get Shop level info - i.e. seller's toggle status is on/off
	// Path: /api/v2/ads/get_shop_toggle_info
	// https://open.shopee.com/documents/v2/v2.ads.get_shop_toggle_info?module=117&type=1
	GetShopToggleInfo(sid uint64, tok string) (*GetShopToggleInfoResponse, error)
	// GetTotalBalance Use this API to return the seller's Real-time total balance of their ads credit including the paid credits and free credits.
	// Path: /api/v2/ads/get_total_balance
	// https://open.shopee.com/documents/v2/v2.ads.get_total_balance?module=117&type=1
	GetTotalBalance(sid uint64, tok string) (*GetTotalBalanceResponse, error)
	// ListGmsUserDeletedItem List GMS items that have been removed from the Campaign by seller
	// Path: /api/v2/ads/list_gms_user_deleted_item
	// https://open.shopee.com/documents/v2/v2.ads.list_gms_user_deleted_item?module=117&type=1
	ListGmsUserDeletedItem(sid uint64, req ListGmsUserDeletedItemRequest, tok string) (*ListGmsUserDeletedItemResponse, error)
}

type AdsServiceOp[T any] struct {
	client *Client[T]
}

// CheckCreateGmsProductCampaignEligibility Check the seller's eligibility in creating a GMS campaign
// Path: /api/v2/ads/check_create_gms_product_campaign_eligibility
// https://open.shopee.com/documents/v2/v2.ads.check_create_gms_product_campaign_eligibility?module=117&type=1
func (s *AdsServiceOp[T]) CheckCreateGmsProductCampaignEligibility(sid uint64, tok string) (*CheckCreateGmsProductCampaignEligibilityResponse, error) {
	path := "/ads/check_create_gms_product_campaign_eligibility"
	resp := new(CheckCreateGmsProductCampaignEligibilityResponse)
	err := s.client.WithShop(sid, tok).Get(path, resp, nil)
	return resp, err
}

// CreateGmsProductCampaign Create a GMS campaign
// Path: /api/v2/ads/create_gms_product_campaign
// https://open.shopee.com/documents/v2/v2.ads.create_gms_product_campaign?module=117&type=1
func (s *AdsServiceOp[T]) CreateGmsProductCampaign(sid uint64, req CreateGmsProductCampaignRequest, tok string) (*CreateGmsProductCampaignResponse, error) {
	path := "/ads/create_gms_product_campaign"
	resp := new(CreateGmsProductCampaignResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

// CreateManualProductAds Use this API to create Manual Selection Product Ads
// Path: /api/v2/ads/create_manual_product_ads
// https://open.shopee.com/documents/v2/v2.ads.create_manual_product_ads?module=117&type=1
func (s *AdsServiceOp[T]) CreateManualProductAds(sid uint64, req CreateManualProductAdsRequest, tok string) (*CreateManualProductAdsResponse, error) {
	path := "/ads/create_manual_product_ads"
	resp := new(CreateManualProductAdsResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

// EditGmsItemProductCampaign Add/remove items to/from the GMS Campaign
// Path: /api/v2/ads/edit_gms_item_product_campaign
// https://open.shopee.com/documents/v2/v2.ads.edit_gms_item_product_campaign?module=117&type=1
func (s *AdsServiceOp[T]) EditGmsItemProductCampaign(sid uint64, req EditGmsItemProductCampaignRequest, tok string) (*EditGmsItemProductCampaignResponse, error) {
	path := "/ads/edit_gms_item_product_campaign"
	resp := new(EditGmsItemProductCampaignResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

// EditGmsProductCampaign Edit a GMS campaign
// Path: /api/v2/ads/edit_gms_product_campaign
// https://open.shopee.com/documents/v2/v2.ads.edit_gms_product_campaign?module=117&type=1
func (s *AdsServiceOp[T]) EditGmsProductCampaign(sid uint64, req EditGmsProductCampaignRequest, tok string) (*EditGmsProductCampaignResponse, error) {
	path := "/ads/edit_gms_product_campaign"
	resp := new(EditGmsProductCampaignResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

// EditManualProductAdKeywords Use this API to edit Manual Selection Product Ad Keywords
// Path: /api/v2/ads/edit_manual_product_ad_keywords
// https://open.shopee.com/documents/v2/v2.ads.edit_manual_product_ad_keywords?module=117&type=1
func (s *AdsServiceOp[T]) EditManualProductAdKeywords(sid uint64, req EditManualProductAdKeywordsRequest, tok string) (*EditManualProductAdKeywordsResponse, error) {
	path := "/ads/edit_manual_product_ad_keywords"
	resp := new(EditManualProductAdKeywordsResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

// EditManualProductAds Use this API to edit Manual Selection Product Ads
// Path: /api/v2/ads/edit_manual_product_ads
// https://open.shopee.com/documents/v2/v2.ads.edit_manual_product_ads?module=117&type=1
func (s *AdsServiceOp[T]) EditManualProductAds(sid uint64, req EditManualProductAdsRequest, tok string) (*EditManualProductAdsResponse, error) {
	path := "/ads/edit_manual_product_ads"
	resp := new(EditManualProductAdsResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

// GetAdsFcilShopRate Get shop rate for Ads Facil Program
// Path: /api/v2/ads/get_ads_facil_shop_rate
// https://open.shopee.com/documents/v2/v2.ads.get_ads_fácil_shop_rate?module=117&type=1
func (s *AdsServiceOp[T]) GetAdsFcilShopRate(sid uint64, tok string) (*GetAdsFcilShopRateResponse, error) {
	path := "/ads/get_ads_facil_shop_rate"
	resp := new(GetAdsFcilShopRateResponse)
	err := s.client.WithShop(sid, tok).Post(path, nil, resp)
	return resp, err
}

// GetAllCpcAdsDailyPerformance Use this API to get Shop level CPC ads multiple-days daily performance.
// Path: /api/v2/ads/get_all_cpc_ads_daily_performance
// https://open.shopee.com/documents/v2/v2.ads.get_all_cpc_ads_daily_performance?module=117&type=1
func (s *AdsServiceOp[T]) GetAllCpcAdsDailyPerformance(sid uint64, opt GetAllCpcAdsDailyPerformanceRequest, tok string) (*GetAllCpcAdsDailyPerformanceResponse, error) {
	path := "/ads/get_all_cpc_ads_daily_performance"
	resp := new(GetAllCpcAdsDailyPerformanceResponse)
	err := s.client.WithShop(sid, tok).Get(path, resp, opt)
	return resp, err
}

// GetAllCpcAdsHourlyPerformance Use this API to get Shop level CPC ads single-date hourly performance.
// Path: /api/v2/ads/get_all_cpc_ads_hourly_performance
// https://open.shopee.com/documents/v2/v2.ads.get_all_cpc_ads_hourly_performance?module=117&type=1
func (s *AdsServiceOp[T]) GetAllCpcAdsHourlyPerformance(sid uint64, opt GetAllCpcAdsHourlyPerformanceRequest, tok string) (*GetAllCpcAdsHourlyPerformanceResponse, error) {
	path := "/ads/get_all_cpc_ads_hourly_performance"
	resp := new(GetAllCpcAdsHourlyPerformanceResponse)
	err := s.client.WithShop(sid, tok).Get(path, resp, opt)
	return resp, err
}

// GetCreateProductAdBudgetSuggestion Call this API to get budget suggestion for product ads creation
// Path: /api/v2/ads/get_create_product_ad_budget_suggestion
// https://open.shopee.com/documents/v2/v2.ads.get_create_product_ad_budget_suggestion?module=117&type=1
func (s *AdsServiceOp[T]) GetCreateProductAdBudgetSuggestion(sid uint64, opt GetCreateProductAdBudgetSuggestionRequest, tok string) (*GetCreateProductAdBudgetSuggestionResponse, error) {
	path := "/ads/get_create_product_ad_budget_suggestion"
	resp := new(GetCreateProductAdBudgetSuggestionResponse)
	err := s.client.WithShop(sid, tok).Get(path, resp, opt)
	return resp, err
}

// GetGmsCampaignPerformance Get GMS Campaign performance
// Path: /api/v2/ads/get_gms_campaign_performance
// https://open.shopee.com/documents/v2/v2.ads.get_gms_campaign_performance?module=117&type=1
func (s *AdsServiceOp[T]) GetGmsCampaignPerformance(sid uint64, req GetGmsCampaignPerformanceRequest, tok string) (*GetGmsCampaignPerformanceResponse, error) {
	path := "/ads/get_gms_campaign_performance"
	resp := new(GetGmsCampaignPerformanceResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

// GetGmsItemPerformance Get GMS Item performance
// 1. The response returned is sorted by item_id
// 2. Only items with performance will be returned
// Path: /api/v2/ads/get_gms_item_performance
// https://open.shopee.com/documents/v2/v2.ads.get_gms_item_performance?module=117&type=1
func (s *AdsServiceOp[T]) GetGmsItemPerformance(sid uint64, req GetGmsItemPerformanceRequest, tok string) (*GetGmsItemPerformanceResponse, error) {
	path := "/ads/get_gms_item_performance"
	resp := new(GetGmsItemPerformanceResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

// GetProductCampaignDailyPerformance Use this API to get Product level ads multiple-days daily performance.
// Path: /api/v2/ads/get_product_campaign_daily_performance
// https://open.shopee.com/documents/v2/v2.ads.get_product_campaign_daily_performance?module=117&type=1
func (s *AdsServiceOp[T]) GetProductCampaignDailyPerformance(sid uint64, opt GetProductCampaignDailyPerformanceRequest, tok string) (*GetProductCampaignDailyPerformanceResponse, error) {
	path := "/ads/get_product_campaign_daily_performance"
	resp := new(GetProductCampaignDailyPerformanceResponse)
	err := s.client.WithShop(sid, tok).Get(path, resp, opt)
	return resp, err
}

// GetProductCampaignHourlyPerformance Use this API to get Product level ads single-day hourly performance.
// Path: /api/v2/ads/get_product_campaign_hourly_performance
// https://open.shopee.com/documents/v2/v2.ads.get_product_campaign_hourly_performance?module=117&type=1
func (s *AdsServiceOp[T]) GetProductCampaignHourlyPerformance(sid uint64, opt GetProductCampaignHourlyPerformanceRequest, tok string) (*GetProductCampaignHourlyPerformanceResponse, error) {
	path := "/ads/get_product_campaign_hourly_performance"
	resp := new(GetProductCampaignHourlyPerformanceResponse)
	err := s.client.WithShop(sid, tok).Get(path, resp, opt)
	return resp, err
}

// GetProductLevelCampaignIdList Call this API to fetch all the product campaign ids displayed on advertiser platform under a specific Shop
// Path: /api/v2/ads/get_product_level_campaign_id_list
// https://open.shopee.com/documents/v2/v2.ads.get_product_level_campaign_id_list?module=117&type=1
func (s *AdsServiceOp[T]) GetProductLevelCampaignIdList(sid uint64, opt GetProductLevelCampaignIdListRequest, tok string) (*GetProductLevelCampaignIdListResponse, error) {
	path := "/ads/get_product_level_campaign_id_list"
	resp := new(GetProductLevelCampaignIdListResponse)
	err := s.client.WithShop(sid, tok).Get(path, resp, opt)
	return resp, err
}

// GetProductLevelCampaignSettingInfo Call this API to fetch all the campaign setting info under this Shop.
// Path: /api/v2/ads/get_product_level_campaign_setting_info
// https://open.shopee.com/documents/v2/v2.ads.get_product_level_campaign_setting_info?module=117&type=1
func (s *AdsServiceOp[T]) GetProductLevelCampaignSettingInfo(sid uint64, opt GetProductLevelCampaignSettingInfoRequest, tok string) (*GetProductLevelCampaignSettingInfoResponse, error) {
	path := "/ads/get_product_level_campaign_setting_info"
	resp := new(GetProductLevelCampaignSettingInfoResponse)
	err := s.client.WithShop(sid, tok).Get(path, resp, opt)
	return resp, err
}

// GetProductRecommendedRoiTarget Get Product Recommended ROI Target
// Path: /api/v2/ads/get_product_recommended_roi_target
// https://open.shopee.com/documents/v2/v2.ads.get_product_recommended_roi_target?module=117&type=1
func (s *AdsServiceOp[T]) GetProductRecommendedRoiTarget(sid uint64, opt GetProductRecommendedRoiTargetRequest, tok string) (*GetProductRecommendedRoiTargetResponse, error) {
	path := "/ads/get_product_recommended_roi_target"
	resp := new(GetProductRecommendedRoiTargetResponse)
	err := s.client.WithShop(sid, tok).Get(path, resp, opt)
	return resp, err
}

// GetRecommendedItemList Use this API to get the list of recommended SKU (Shop level) with the corresponding tag, i.e top search/best selling/best ROI tag.
// Path: /api/v2/ads/get_recommended_item_list
// https://open.shopee.com/documents/v2/v2.ads.get_recommended_item_list?module=117&type=1
func (s *AdsServiceOp[T]) GetRecommendedItemList(sid uint64, tok string) (*GetRecommendedItemListResponse, error) {
	path := "/ads/get_recommended_item_list"
	resp := new(GetRecommendedItemListResponse)
	err := s.client.WithShop(sid, tok).Get(path, resp, nil)
	return resp, err
}

// GetRecommendedKeywordList Use this API to get the list of Recommended keywords by item and optionally a search keyword
// Path: /api/v2/ads/get_recommended_keyword_list
// https://open.shopee.com/documents/v2/v2.ads.get_recommended_keyword_list?module=117&type=1
func (s *AdsServiceOp[T]) GetRecommendedKeywordList(sid uint64, opt GetRecommendedKeywordListRequest, tok string) (*GetRecommendedKeywordListResponse, error) {
	path := "/ads/get_recommended_keyword_list"
	resp := new(GetRecommendedKeywordListResponse)
	err := s.client.WithShop(sid, tok).Get(path, resp, opt)
	return resp, err
}

// GetShopToggleInfo Use this API to get Shop level info - i.e. seller's toggle status is on/off
// Path: /api/v2/ads/get_shop_toggle_info
// https://open.shopee.com/documents/v2/v2.ads.get_shop_toggle_info?module=117&type=1
func (s *AdsServiceOp[T]) GetShopToggleInfo(sid uint64, tok string) (*GetShopToggleInfoResponse, error) {
	path := "/ads/get_shop_toggle_info"
	resp := new(GetShopToggleInfoResponse)
	err := s.client.WithShop(sid, tok).Get(path, resp, nil)
	return resp, err
}

// GetTotalBalance Use this API to return the seller's Real-time total balance of their ads credit including the paid credits and free credits.
// Path: /api/v2/ads/get_total_balance
// https://open.shopee.com/documents/v2/v2.ads.get_total_balance?module=117&type=1
func (s *AdsServiceOp[T]) GetTotalBalance(sid uint64, tok string) (*GetTotalBalanceResponse, error) {
	path := "/ads/get_total_balance"
	resp := new(GetTotalBalanceResponse)
	err := s.client.WithShop(sid, tok).Get(path, resp, nil)
	return resp, err
}

// ListGmsUserDeletedItem List GMS items that have been removed from the Campaign by seller
// Path: /api/v2/ads/list_gms_user_deleted_item
// https://open.shopee.com/documents/v2/v2.ads.list_gms_user_deleted_item?module=117&type=1
func (s *AdsServiceOp[T]) ListGmsUserDeletedItem(sid uint64, req ListGmsUserDeletedItemRequest, tok string) (*ListGmsUserDeletedItemResponse, error) {
	path := "/ads/list_gms_user_deleted_item"
	resp := new(ListGmsUserDeletedItemResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}
