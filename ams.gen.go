package goshopee

import (
	"context"
)

type AMSService interface {
	// AddAllProductsToOpenCampaign Add all eligible products into the Open Campaign. We will only return the general error that caused the whole task failure, without returning the specific error for each product in the v2.ams.get_open_campaign_batch_task_result API. If you want to get the result for each products, you can use v2.ams.batch_add_products_to_open_campaign by pagination manually, or check the product status by using the GET API after the task progress turn to 100%.
	// Path: /api/v2/ams/add_all_products_to_open_campaign
	// https://open.shopee.com/documents/v2/v2.ams.add_all_products_to_open_campaign?module=127&type=1
	AddAllProductsToOpenCampaign(ctx context.Context, sid uint64, mid uint64, tok string, req AddAllProductsToOpenCampaignRequest) (*AddAllProductsToOpenCampaignResponse, error)
	// BatchAddProductsToOpenCampaign Batch add products to the Open Campaign for a given list of product IDs
	// Path: /api/v2/ams/batch_add_products_to_open_campaign
	// https://open.shopee.com/documents/v2/v2.ams.batch_add_products_to_open_campaign?module=127&type=1
	BatchAddProductsToOpenCampaign(ctx context.Context, sid uint64, mid uint64, tok string, req BatchAddProductsToOpenCampaignRequest) (*BatchAddProductsToOpenCampaignResponse, error)
	// BatchEditProductsOpenCampaignSetting Batch update open campaign settings for a given list of product IDs
	// Path: /api/v2/ams/batch_edit_products_open_campaign_setting
	// https://open.shopee.com/documents/v2/v2.ams.batch_edit_products_open_campaign_setting?module=127&type=1
	BatchEditProductsOpenCampaignSetting(ctx context.Context, sid uint64, mid uint64, tok string, req BatchEditProductsOpenCampaignSettingRequest) (*BatchEditProductsOpenCampaignSettingResponse, error)
	// BatchGetProductsSuggestedRate Fetch suggested rates for a given list of product IDs
	// Path: /api/v2/ams/batch_get_products_suggested_rate
	// https://open.shopee.com/documents/v2/v2.ams.batch_get_products_suggested_rate?module=127&type=1
	BatchGetProductsSuggestedRate(ctx context.Context, sid uint64, mid uint64, tok string, opt BatchGetProductsSuggestedRateRequest) (*BatchGetProductsSuggestedRateResponse, error)
	// BatchRemoveProductsOpenCampaignSetting Batch update products from Open Campaign for a given list of product IDs
	// Path: /api/v2/ams/batch_remove_products_open_campaign_setting
	// https://open.shopee.com/documents/v2/v2.ams.batch_remove_products_open_campaign_setting?module=127&type=1
	BatchRemoveProductsOpenCampaignSetting(ctx context.Context, sid uint64, mid uint64, tok string, req BatchRemoveProductsOpenCampaignSettingRequest) (*BatchRemoveProductsOpenCampaignSettingResponse, error)
	// CreateNewTargetedCampaign Create a new campaign with custom product & affiliate selections, and basic info filling.
	// Path: /api/v2/ams/create_new_targeted_campaign
	// https://open.shopee.com/documents/v2/v2.ams.create_new_targeted_campaign?module=127&type=1
	CreateNewTargetedCampaign(ctx context.Context, sid uint64, mid uint64, tok string, req CreateNewTargetedCampaignRequest) (*CreateNewTargetedCampaignResponse, error)
	// EditAffiliateListOfTargetedCampaign Modify the selected affiliate list in an existing target campaign
	// Path: /api/v2/ams/edit_affiliate_list_of_targeted_campaign
	// https://open.shopee.com/documents/v2/v2.ams.edit_affiliate_list_of_targeted_campaign?module=127&type=1
	EditAffiliateListOfTargetedCampaign(ctx context.Context, sid uint64, mid uint64, tok string, req EditAffiliateListOfTargetedCampaignRequest) (*EditAffiliateListOfTargetedCampaignResponse, error)
	// EditAllProductsOpenCampaignSetting Update for all products in the Open Campaign. We will only return the general error that caused the whole task failure, without returning the specific error for each product in the v2.ams.get_open_campaign_batch_task_result API. If you want to get the result for each products, you can use v2.ams.batch_edit_products_open_campaign_setting by pagination manually, or check the product status by using the GET API after the task progress turn to 100%.
	// Path: /api/v2/ams/edit_all_products_open_campaign_setting
	// https://open.shopee.com/documents/v2/v2.ams.edit_all_products_open_campaign_setting?module=127&type=1
	EditAllProductsOpenCampaignSetting(ctx context.Context, sid uint64, mid uint64, tok string, req EditAllProductsOpenCampaignSettingRequest) (*EditAllProductsOpenCampaignSettingResponse, error)
	// EditProductListOfTargetedCampaign Modify the selected product list in an existing target campaign
	// Path: /api/v2/ams/edit_product_list_of_targeted_campaign
	// https://open.shopee.com/documents/v2/v2.ams.edit_product_list_of_targeted_campaign?module=127&type=1
	EditProductListOfTargetedCampaign(ctx context.Context, sid uint64, mid uint64, tok string, req EditProductListOfTargetedCampaignRequest) (*EditProductListOfTargetedCampaignResponse, error)
	// GetAffiliatePerformance Retrieve affiliate performance of the shop.
	// Path: /api/v2/ams/get_affiliate_performance
	// https://open.shopee.com/documents/v2/v2.ams.get_affiliate_performance?module=127&type=1
	GetAffiliatePerformance(ctx context.Context, sid uint64, mid uint64, tok string, opt GetAffiliatePerformanceRequest) (*GetAffiliatePerformanceResponse, error)
	// GetAutoAddNewProductToggleStatus Check if auto-add new product is currently enabled
	// Path: /api/v2/ams/get_auto_add_new_product_toggle_status
	// https://open.shopee.com/documents/v2/v2.ams.get_auto_add_new_product_toggle_status?module=127&type=1
	GetAutoAddNewProductToggleStatus(ctx context.Context, sid uint64, mid uint64, tok string) (*GetAutoAddNewProductToggleStatusResponse, error)
	// GetCampaignKeyMetricsPerformance Retrieve key metrics for Open and Targeted campaigns
	// Path: /api/v2/ams/get_campaign_key_metrics_performance
	// https://open.shopee.com/documents/v2/v2.ams.get_campaign_key_metrics_performance?module=127&type=1
	GetCampaignKeyMetricsPerformance(ctx context.Context, sid uint64, mid uint64, tok string, opt GetCampaignKeyMetricsPerformanceRequest) (*GetCampaignKeyMetricsPerformanceResponse, error)
	// GetContentPerformance Retrieve content performance of the shop
	// Path: /api/v2/ams/get_content_performance
	// https://open.shopee.com/documents/v2/v2.ams.get_content_performance?module=127&type=1
	GetContentPerformance(ctx context.Context, sid uint64, mid uint64, tok string, opt GetContentPerformanceRequest) (*GetContentPerformanceResponse, error)
	// GetConversionReport Retrieve the shop's conversion report with details about each order, item, affiliate, campaign.You can filter results using one or multiple time ranges, and the final result will be the intersection of these ranges. Due to data volume limitations, the maximum queryable time span is three months, etc.Maximum data can be viewed is 500 pages, please export data for more details.
	// Path: /api/v2/ams/get_conversion_report
	// https://open.shopee.com/documents/v2/v2.ams.get_conversion_report?module=127&type=1
	GetConversionReport(ctx context.Context, sid uint64, mid uint64, tok string, opt GetConversionReportRequest) (*GetConversionReportResponse, error)
	// GetManagedAffiliateList Returns affiliates that are saved to managed affiliate list
	// Path: /api/v2/ams/get_managed_affiliate_list
	// https://open.shopee.com/documents/v2/v2.ams.get_managed_affiliate_list?module=127&type=1
	GetManagedAffiliateList(ctx context.Context, sid uint64, mid uint64, tok string, opt GetManagedAffiliateListRequest) (*GetManagedAffiliateListResponse, error)
	// GetOpenCampaignAddedProduct Retrieve all products currently in the Open Campaign, including campaign status, commission rate, and promotion period
	// Path: /api/v2/ams/get_open_campaign_added_product
	// https://open.shopee.com/documents/v2/v2.ams.get_open_campaign_added_product?module=127&type=1
	GetOpenCampaignAddedProduct(ctx context.Context, sid uint64, mid uint64, tok string, opt GetOpenCampaignAddedProductRequest) (*GetOpenCampaignAddedProductResponse, error)
	// GetOpenCampaignBatchTaskResult Get open campaign batch task result
	// Path: /api/v2/ams/get_open_campaign_batch_task_result
	// https://open.shopee.com/documents/v2/v2.ams.get_open_campaign_batch_task_result?module=127&type=1
	GetOpenCampaignBatchTaskResult(ctx context.Context, sid uint64, mid uint64, tok string, opt GetOpenCampaignBatchTaskResultRequest) (*GetOpenCampaignBatchTaskResultResponse, error)
	// GetOpenCampaignNotAddedProduct Retrieve eligible products not yet added to the Open Campaign
	// Path: /api/v2/ams/get_open_campaign_not_added_product
	// https://open.shopee.com/documents/v2/v2.ams.get_open_campaign_not_added_product?module=127&type=1
	GetOpenCampaignNotAddedProduct(ctx context.Context, sid uint64, mid uint64, tok string, opt GetOpenCampaignNotAddedProductRequest) (*GetOpenCampaignNotAddedProductResponse, error)
	// GetOpenCampaignPerformance Retrieve all products in the Open Campaign along with performance data
	// Path: /api/v2/ams/get_open_campaign_performance
	// https://open.shopee.com/documents/v2/v2.ams.get_open_campaign_performance?module=127&type=1
	GetOpenCampaignPerformance(ctx context.Context, sid uint64, mid uint64, tok string, opt GetOpenCampaignPerformanceRequest) (*GetOpenCampaignPerformanceResponse, error)
	// GetOptimizationSuggestionProduct Retrieve products with suggestions to improve performance
	// Path: /api/v2/ams/get_optimization_suggestion_product
	// https://open.shopee.com/documents/v2/v2.ams.get_optimization_suggestion_product?module=127&type=1
	GetOptimizationSuggestionProduct(ctx context.Context, sid uint64, mid uint64, tok string, opt GetOptimizationSuggestionProductRequest) (*GetOptimizationSuggestionProductResponse, error)
	// GetPerformanceDataUpdateTime Retrieve the latest date of AMS dashboard data metrics update.
	// Path: /api/v2/ams/get_performance_data_update_time
	// https://open.shopee.com/documents/v2/v2.ams.get_performance_data_update_time?module=127&type=1
	GetPerformanceDataUpdateTime(ctx context.Context, sid uint64, mid uint64, tok string, opt GetPerformanceDataUpdateTimeRequest) (*GetPerformanceDataUpdateTimeResponse, error)
	// GetProductPerformance Retrieve product performance of the shop.
	// Path: /api/v2/ams/get_product_performance
	// https://open.shopee.com/documents/v2/v2.ams.get_product_performance?module=127&type=1
	GetProductPerformance(ctx context.Context, sid uint64, mid uint64, tok string, opt GetProductPerformanceRequest) (*GetProductPerformanceResponse, error)
	// GetRecommendedAffiliateList Returns top 200 recommended affiliates that can be added to a campaign
	// Path: /api/v2/ams/get_recommended_affiliate_list
	// https://open.shopee.com/documents/v2/v2.ams.get_recommended_affiliate_list?module=127&type=1
	GetRecommendedAffiliateList(ctx context.Context, sid uint64, mid uint64, tok string, opt GetRecommendedAffiliateListRequest) (*GetRecommendedAffiliateListResponse, error)
	// GetShopPerformance Retrieve overall key metrics for all channels or specific channels.
	// Path: /api/v2/ams/get_shop_performance
	// https://open.shopee.com/documents/v2/v2.ams.get_shop_performance?module=127&type=1
	GetShopPerformance(ctx context.Context, sid uint64, mid uint64, tok string, opt GetShopPerformanceRequest) (*AMSGetShopPerformanceResponse, error)
	// GetShopSuggestedRate Retrieve suggested rates for all eligible products
	// Path: /api/v2/ams/get_shop_suggested_rate
	// https://open.shopee.com/documents/v2/v2.ams.get_shop_suggested_rate?module=127&type=1
	GetShopSuggestedRate(ctx context.Context, sid uint64, mid uint64, tok string) (*GetShopSuggestedRateResponse, error)
	// GetTargetedCampaignAddableProductList Returns a list of products that can be added to a targeted campaign
	// Path: /api/v2/ams/get_targeted_campaign_addable_product_list
	// https://open.shopee.com/documents/v2/v2.ams.get_targeted_campaign_addable_product_list?module=127&type=1
	GetTargetedCampaignAddableProductList(ctx context.Context, sid uint64, mid uint64, tok string, opt GetTargetedCampaignAddableProductListRequest) (*GetTargetedCampaignAddableProductListResponse, error)
	// GetTargetedCampaignList {"content":"<p>Retrieve all current targeted campaigns created by the seller</p>","raw_content":[{"name":"paragraph","children":[{"data":"Retrieve all current targeted campaigns created by the seller"}]}]}
	// Path: /api/v2/ams/get_targeted_campaign_list
	// https://open.shopee.com/documents/v2/v2.ams.get_targeted_campaign_list?module=127&type=1
	GetTargetedCampaignList(ctx context.Context, sid uint64, mid uint64, tok string, opt GetTargetedCampaignListRequest) (*GetTargetedCampaignListResponse, error)
	// GetTargetedCampaignPerformance Retrieve a list of Targeted Campaigns and their performance data
	// Path: /api/v2/ams/get_targeted_campaign_performance
	// https://open.shopee.com/documents/v2/v2.ams.get_targeted_campaign_performance?module=127&type=1
	GetTargetedCampaignPerformance(ctx context.Context, sid uint64, mid uint64, tok string, opt GetTargetedCampaignPerformanceRequest) (*GetTargetedCampaignPerformanceResponse, error)
	// GetTargetedCampaignSettings For each campaign, return: campaign basic info (name, status, promotion period, message), selected product list (with product name & ID), selected affiliate list (with affiliate names)
	// Path: /api/v2/ams/get_targeted_campaign_settings
	// https://open.shopee.com/documents/v2/v2.ams.get_targeted_campaign_settings?module=127&type=1
	GetTargetedCampaignSettings(ctx context.Context, sid uint64, mid uint64, tok string, opt GetTargetedCampaignSettingsRequest) (*GetTargetedCampaignSettingsResponse, error)
	// GetValidationList Retrieve the seller's AMS validation bill
	// Path: /api/v2/ams/get_validation_list
	// https://open.shopee.com/documents/v2/v2.ams.get_validation_list?module=127&type=1
	GetValidationList(ctx context.Context, sid uint64, mid uint64, tok string) (*GetValidationListResponse, error)
	// GetValidationReport Retrieve detailed information for a specific validation bill
	// Path: /api/v2/ams/get_validation_report
	// https://open.shopee.com/documents/v2/v2.ams.get_validation_report?module=127&type=1
	GetValidationReport(ctx context.Context, sid uint64, mid uint64, tok string, opt GetValidationReportRequest) (*GetValidationReportResponse, error)
	// QueryAffiliateList Retrieve affiliate information by affiliate id.
	// Path: /api/v2/ams/query_affiliate_list
	// https://open.shopee.com/documents/v2/v2.ams.query_affiliate_list?module=127&type=1
	QueryAffiliateList(ctx context.Context, sid uint64, mid uint64, tok string, opt QueryAffiliateListRequest) (*QueryAffiliateListResponse, error)
	// RemoveAllProductsOpenCampaignSetting Remove the entire product list of Open Campaign. We will only return the general error that caused the whole task failure, without returning the specific error for each product in the v2.ams.get_open_campaign_batch_task_result API. If you want to get the result for each products, you can use v2.ams. batch_remove_products_open_campaign_setting by pagination manually, or check the product status by using the GET API after the task progress turn to 100%.
	// Path: /api/v2/ams/remove_all_products_open_campaign_setting
	// https://open.shopee.com/documents/v2/v2.ams.remove_all_products_open_campaign_setting?module=127&type=1
	RemoveAllProductsOpenCampaignSetting(ctx context.Context, sid uint64, mid uint64, tok string) (*RemoveAllProductsOpenCampaignSettingResponse, error)
	// TerminateTargetedCampaign Change target campaign status to "terminated" to stop all affiliate promotion activity
	// Path: /api/v2/ams/terminate_targeted_campaign
	// https://open.shopee.com/documents/v2/v2.ams.terminate_targeted_campaign?module=127&type=1
	TerminateTargetedCampaign(ctx context.Context, sid uint64, mid uint64, tok string, req TerminateTargetedCampaignRequest) (*TerminateTargetedCampaignResponse, error)
	// UpdateAutoAddNewProductSetting Change auto-add toggle and default commission rate setting
	// Path: /api/v2/ams/update_auto_add_new_product_setting
	// https://open.shopee.com/documents/v2/v2.ams.update_auto_add_new_product_setting?module=127&type=1
	UpdateAutoAddNewProductSetting(ctx context.Context, sid uint64, mid uint64, tok string, req UpdateAutoAddNewProductSettingRequest) (*UpdateAutoAddNewProductSettingResponse, error)
	// UpdateBasicInfoOfTargetedCampaign Edit campaign name, promotion period, message, and budget (if the shop is whitelisted) of target campaign
	// Path: /api/v2/ams/update_basic_info_of_targeted_campaign
	// https://open.shopee.com/documents/v2/v2.ams.update_basic_info_of_targeted_campaign?module=127&type=1
	UpdateBasicInfoOfTargetedCampaign(ctx context.Context, sid uint64, mid uint64, tok string, req UpdateBasicInfoOfTargetedCampaignRequest) (*UpdateBasicInfoOfTargetedCampaignResponse, error)
}

type AMSServiceOp[T any] struct {
	client *Client[T]
}

// AddAllProductsToOpenCampaign Add all eligible products into the Open Campaign. We will only return the general error that caused the whole task failure, without returning the specific error for each product in the v2.ams.get_open_campaign_batch_task_result API. If you want to get the result for each products, you can use v2.ams.batch_add_products_to_open_campaign by pagination manually, or check the product status by using the GET API after the task progress turn to 100%.
// Path: /api/v2/ams/add_all_products_to_open_campaign
// https://open.shopee.com/documents/v2/v2.ams.add_all_products_to_open_campaign?module=127&type=1
func (s *AMSServiceOp[T]) AddAllProductsToOpenCampaign(ctx context.Context, sid uint64, mid uint64, tok string, req AddAllProductsToOpenCampaignRequest) (*AddAllProductsToOpenCampaignResponse, error) {
	path := "/ams/add_all_products_to_open_campaign"
	resp := new(AddAllProductsToOpenCampaignResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// BatchAddProductsToOpenCampaign Batch add products to the Open Campaign for a given list of product IDs
// Path: /api/v2/ams/batch_add_products_to_open_campaign
// https://open.shopee.com/documents/v2/v2.ams.batch_add_products_to_open_campaign?module=127&type=1
func (s *AMSServiceOp[T]) BatchAddProductsToOpenCampaign(ctx context.Context, sid uint64, mid uint64, tok string, req BatchAddProductsToOpenCampaignRequest) (*BatchAddProductsToOpenCampaignResponse, error) {
	path := "/ams/batch_add_products_to_open_campaign"
	resp := new(BatchAddProductsToOpenCampaignResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// BatchEditProductsOpenCampaignSetting Batch update open campaign settings for a given list of product IDs
// Path: /api/v2/ams/batch_edit_products_open_campaign_setting
// https://open.shopee.com/documents/v2/v2.ams.batch_edit_products_open_campaign_setting?module=127&type=1
func (s *AMSServiceOp[T]) BatchEditProductsOpenCampaignSetting(ctx context.Context, sid uint64, mid uint64, tok string, req BatchEditProductsOpenCampaignSettingRequest) (*BatchEditProductsOpenCampaignSettingResponse, error) {
	path := "/ams/batch_edit_products_open_campaign_setting"
	resp := new(BatchEditProductsOpenCampaignSettingResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// BatchGetProductsSuggestedRate Fetch suggested rates for a given list of product IDs
// Path: /api/v2/ams/batch_get_products_suggested_rate
// https://open.shopee.com/documents/v2/v2.ams.batch_get_products_suggested_rate?module=127&type=1
func (s *AMSServiceOp[T]) BatchGetProductsSuggestedRate(ctx context.Context, sid uint64, mid uint64, tok string, opt BatchGetProductsSuggestedRateRequest) (*BatchGetProductsSuggestedRateResponse, error) {
	path := "/ams/batch_get_products_suggested_rate"
	resp := new(BatchGetProductsSuggestedRateResponse)
	err := s.client.Get(ctx, path, resp, opt, sid, mid, tok)
	return resp, err
}

// BatchRemoveProductsOpenCampaignSetting Batch update products from Open Campaign for a given list of product IDs
// Path: /api/v2/ams/batch_remove_products_open_campaign_setting
// https://open.shopee.com/documents/v2/v2.ams.batch_remove_products_open_campaign_setting?module=127&type=1
func (s *AMSServiceOp[T]) BatchRemoveProductsOpenCampaignSetting(ctx context.Context, sid uint64, mid uint64, tok string, req BatchRemoveProductsOpenCampaignSettingRequest) (*BatchRemoveProductsOpenCampaignSettingResponse, error) {
	path := "/ams/batch_remove_products_open_campaign_setting"
	resp := new(BatchRemoveProductsOpenCampaignSettingResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// CreateNewTargetedCampaign Create a new campaign with custom product & affiliate selections, and basic info filling.
// Path: /api/v2/ams/create_new_targeted_campaign
// https://open.shopee.com/documents/v2/v2.ams.create_new_targeted_campaign?module=127&type=1
func (s *AMSServiceOp[T]) CreateNewTargetedCampaign(ctx context.Context, sid uint64, mid uint64, tok string, req CreateNewTargetedCampaignRequest) (*CreateNewTargetedCampaignResponse, error) {
	path := "/ams/create_new_targeted_campaign"
	resp := new(CreateNewTargetedCampaignResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// EditAffiliateListOfTargetedCampaign Modify the selected affiliate list in an existing target campaign
// Path: /api/v2/ams/edit_affiliate_list_of_targeted_campaign
// https://open.shopee.com/documents/v2/v2.ams.edit_affiliate_list_of_targeted_campaign?module=127&type=1
func (s *AMSServiceOp[T]) EditAffiliateListOfTargetedCampaign(ctx context.Context, sid uint64, mid uint64, tok string, req EditAffiliateListOfTargetedCampaignRequest) (*EditAffiliateListOfTargetedCampaignResponse, error) {
	path := "/ams/edit_affiliate_list_of_targeted_campaign"
	resp := new(EditAffiliateListOfTargetedCampaignResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// EditAllProductsOpenCampaignSetting Update for all products in the Open Campaign. We will only return the general error that caused the whole task failure, without returning the specific error for each product in the v2.ams.get_open_campaign_batch_task_result API. If you want to get the result for each products, you can use v2.ams.batch_edit_products_open_campaign_setting by pagination manually, or check the product status by using the GET API after the task progress turn to 100%.
// Path: /api/v2/ams/edit_all_products_open_campaign_setting
// https://open.shopee.com/documents/v2/v2.ams.edit_all_products_open_campaign_setting?module=127&type=1
func (s *AMSServiceOp[T]) EditAllProductsOpenCampaignSetting(ctx context.Context, sid uint64, mid uint64, tok string, req EditAllProductsOpenCampaignSettingRequest) (*EditAllProductsOpenCampaignSettingResponse, error) {
	path := "/ams/edit_all_products_open_campaign_setting"
	resp := new(EditAllProductsOpenCampaignSettingResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// EditProductListOfTargetedCampaign Modify the selected product list in an existing target campaign
// Path: /api/v2/ams/edit_product_list_of_targeted_campaign
// https://open.shopee.com/documents/v2/v2.ams.edit_product_list_of_targeted_campaign?module=127&type=1
func (s *AMSServiceOp[T]) EditProductListOfTargetedCampaign(ctx context.Context, sid uint64, mid uint64, tok string, req EditProductListOfTargetedCampaignRequest) (*EditProductListOfTargetedCampaignResponse, error) {
	path := "/ams/edit_product_list_of_targeted_campaign"
	resp := new(EditProductListOfTargetedCampaignResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// GetAffiliatePerformance Retrieve affiliate performance of the shop.
// Path: /api/v2/ams/get_affiliate_performance
// https://open.shopee.com/documents/v2/v2.ams.get_affiliate_performance?module=127&type=1
func (s *AMSServiceOp[T]) GetAffiliatePerformance(ctx context.Context, sid uint64, mid uint64, tok string, opt GetAffiliatePerformanceRequest) (*GetAffiliatePerformanceResponse, error) {
	path := "/ams/get_affiliate_performance"
	resp := new(GetAffiliatePerformanceResponse)
	err := s.client.Get(ctx, path, resp, opt, sid, mid, tok)
	return resp, err
}

// GetAutoAddNewProductToggleStatus Check if auto-add new product is currently enabled
// Path: /api/v2/ams/get_auto_add_new_product_toggle_status
// https://open.shopee.com/documents/v2/v2.ams.get_auto_add_new_product_toggle_status?module=127&type=1
func (s *AMSServiceOp[T]) GetAutoAddNewProductToggleStatus(ctx context.Context, sid uint64, mid uint64, tok string) (*GetAutoAddNewProductToggleStatusResponse, error) {
	path := "/ams/get_auto_add_new_product_toggle_status"
	resp := new(GetAutoAddNewProductToggleStatusResponse)
	err := s.client.Get(ctx, path, resp, nil, sid, mid, tok)
	return resp, err
}

// GetCampaignKeyMetricsPerformance Retrieve key metrics for Open and Targeted campaigns
// Path: /api/v2/ams/get_campaign_key_metrics_performance
// https://open.shopee.com/documents/v2/v2.ams.get_campaign_key_metrics_performance?module=127&type=1
func (s *AMSServiceOp[T]) GetCampaignKeyMetricsPerformance(ctx context.Context, sid uint64, mid uint64, tok string, opt GetCampaignKeyMetricsPerformanceRequest) (*GetCampaignKeyMetricsPerformanceResponse, error) {
	path := "/ams/get_campaign_key_metrics_performance"
	resp := new(GetCampaignKeyMetricsPerformanceResponse)
	err := s.client.Get(ctx, path, resp, opt, sid, mid, tok)
	return resp, err
}

// GetContentPerformance Retrieve content performance of the shop
// Path: /api/v2/ams/get_content_performance
// https://open.shopee.com/documents/v2/v2.ams.get_content_performance?module=127&type=1
func (s *AMSServiceOp[T]) GetContentPerformance(ctx context.Context, sid uint64, mid uint64, tok string, opt GetContentPerformanceRequest) (*GetContentPerformanceResponse, error) {
	path := "/ams/get_content_performance"
	resp := new(GetContentPerformanceResponse)
	err := s.client.Get(ctx, path, resp, opt, sid, mid, tok)
	return resp, err
}

// GetConversionReport Retrieve the shop's conversion report with details about each order, item, affiliate, campaign.You can filter results using one or multiple time ranges, and the final result will be the intersection of these ranges. Due to data volume limitations, the maximum queryable time span is three months, etc.Maximum data can be viewed is 500 pages, please export data for more details.
// Path: /api/v2/ams/get_conversion_report
// https://open.shopee.com/documents/v2/v2.ams.get_conversion_report?module=127&type=1
func (s *AMSServiceOp[T]) GetConversionReport(ctx context.Context, sid uint64, mid uint64, tok string, opt GetConversionReportRequest) (*GetConversionReportResponse, error) {
	path := "/ams/get_conversion_report"
	resp := new(GetConversionReportResponse)
	err := s.client.Get(ctx, path, resp, opt, sid, mid, tok)
	return resp, err
}

// GetManagedAffiliateList Returns affiliates that are saved to managed affiliate list
// Path: /api/v2/ams/get_managed_affiliate_list
// https://open.shopee.com/documents/v2/v2.ams.get_managed_affiliate_list?module=127&type=1
func (s *AMSServiceOp[T]) GetManagedAffiliateList(ctx context.Context, sid uint64, mid uint64, tok string, opt GetManagedAffiliateListRequest) (*GetManagedAffiliateListResponse, error) {
	path := "/ams/get_managed_affiliate_list"
	resp := new(GetManagedAffiliateListResponse)
	err := s.client.Get(ctx, path, resp, opt, sid, mid, tok)
	return resp, err
}

// GetOpenCampaignAddedProduct Retrieve all products currently in the Open Campaign, including campaign status, commission rate, and promotion period
// Path: /api/v2/ams/get_open_campaign_added_product
// https://open.shopee.com/documents/v2/v2.ams.get_open_campaign_added_product?module=127&type=1
func (s *AMSServiceOp[T]) GetOpenCampaignAddedProduct(ctx context.Context, sid uint64, mid uint64, tok string, opt GetOpenCampaignAddedProductRequest) (*GetOpenCampaignAddedProductResponse, error) {
	path := "/ams/get_open_campaign_added_product"
	resp := new(GetOpenCampaignAddedProductResponse)
	err := s.client.Get(ctx, path, resp, opt, sid, mid, tok)
	return resp, err
}

// GetOpenCampaignBatchTaskResult Get open campaign batch task result
// Path: /api/v2/ams/get_open_campaign_batch_task_result
// https://open.shopee.com/documents/v2/v2.ams.get_open_campaign_batch_task_result?module=127&type=1
func (s *AMSServiceOp[T]) GetOpenCampaignBatchTaskResult(ctx context.Context, sid uint64, mid uint64, tok string, opt GetOpenCampaignBatchTaskResultRequest) (*GetOpenCampaignBatchTaskResultResponse, error) {
	path := "/ams/get_open_campaign_batch_task_result"
	resp := new(GetOpenCampaignBatchTaskResultResponse)
	err := s.client.Get(ctx, path, resp, opt, sid, mid, tok)
	return resp, err
}

// GetOpenCampaignNotAddedProduct Retrieve eligible products not yet added to the Open Campaign
// Path: /api/v2/ams/get_open_campaign_not_added_product
// https://open.shopee.com/documents/v2/v2.ams.get_open_campaign_not_added_product?module=127&type=1
func (s *AMSServiceOp[T]) GetOpenCampaignNotAddedProduct(ctx context.Context, sid uint64, mid uint64, tok string, opt GetOpenCampaignNotAddedProductRequest) (*GetOpenCampaignNotAddedProductResponse, error) {
	path := "/ams/get_open_campaign_not_added_product"
	resp := new(GetOpenCampaignNotAddedProductResponse)
	err := s.client.Get(ctx, path, resp, opt, sid, mid, tok)
	return resp, err
}

// GetOpenCampaignPerformance Retrieve all products in the Open Campaign along with performance data
// Path: /api/v2/ams/get_open_campaign_performance
// https://open.shopee.com/documents/v2/v2.ams.get_open_campaign_performance?module=127&type=1
func (s *AMSServiceOp[T]) GetOpenCampaignPerformance(ctx context.Context, sid uint64, mid uint64, tok string, opt GetOpenCampaignPerformanceRequest) (*GetOpenCampaignPerformanceResponse, error) {
	path := "/ams/get_open_campaign_performance"
	resp := new(GetOpenCampaignPerformanceResponse)
	err := s.client.Get(ctx, path, resp, opt, sid, mid, tok)
	return resp, err
}

// GetOptimizationSuggestionProduct Retrieve products with suggestions to improve performance
// Path: /api/v2/ams/get_optimization_suggestion_product
// https://open.shopee.com/documents/v2/v2.ams.get_optimization_suggestion_product?module=127&type=1
func (s *AMSServiceOp[T]) GetOptimizationSuggestionProduct(ctx context.Context, sid uint64, mid uint64, tok string, opt GetOptimizationSuggestionProductRequest) (*GetOptimizationSuggestionProductResponse, error) {
	path := "/ams/get_optimization_suggestion_product"
	resp := new(GetOptimizationSuggestionProductResponse)
	err := s.client.Get(ctx, path, resp, opt, sid, mid, tok)
	return resp, err
}

// GetPerformanceDataUpdateTime Retrieve the latest date of AMS dashboard data metrics update.
// Path: /api/v2/ams/get_performance_data_update_time
// https://open.shopee.com/documents/v2/v2.ams.get_performance_data_update_time?module=127&type=1
func (s *AMSServiceOp[T]) GetPerformanceDataUpdateTime(ctx context.Context, sid uint64, mid uint64, tok string, opt GetPerformanceDataUpdateTimeRequest) (*GetPerformanceDataUpdateTimeResponse, error) {
	path := "/ams/get_performance_data_update_time"
	resp := new(GetPerformanceDataUpdateTimeResponse)
	err := s.client.Get(ctx, path, resp, opt, sid, mid, tok)
	return resp, err
}

// GetProductPerformance Retrieve product performance of the shop.
// Path: /api/v2/ams/get_product_performance
// https://open.shopee.com/documents/v2/v2.ams.get_product_performance?module=127&type=1
func (s *AMSServiceOp[T]) GetProductPerformance(ctx context.Context, sid uint64, mid uint64, tok string, opt GetProductPerformanceRequest) (*GetProductPerformanceResponse, error) {
	path := "/ams/get_product_performance"
	resp := new(GetProductPerformanceResponse)
	err := s.client.Get(ctx, path, resp, opt, sid, mid, tok)
	return resp, err
}

// GetRecommendedAffiliateList Returns top 200 recommended affiliates that can be added to a campaign
// Path: /api/v2/ams/get_recommended_affiliate_list
// https://open.shopee.com/documents/v2/v2.ams.get_recommended_affiliate_list?module=127&type=1
func (s *AMSServiceOp[T]) GetRecommendedAffiliateList(ctx context.Context, sid uint64, mid uint64, tok string, opt GetRecommendedAffiliateListRequest) (*GetRecommendedAffiliateListResponse, error) {
	path := "/ams/get_recommended_affiliate_list"
	resp := new(GetRecommendedAffiliateListResponse)
	err := s.client.Get(ctx, path, resp, opt, sid, mid, tok)
	return resp, err
}

// GetShopPerformance Retrieve overall key metrics for all channels or specific channels.
// Path: /api/v2/ams/get_shop_performance
// https://open.shopee.com/documents/v2/v2.ams.get_shop_performance?module=127&type=1
func (s *AMSServiceOp[T]) GetShopPerformance(ctx context.Context, sid uint64, mid uint64, tok string, opt GetShopPerformanceRequest) (*AMSGetShopPerformanceResponse, error) {
	path := "/ams/get_shop_performance"
	resp := new(AMSGetShopPerformanceResponse)
	err := s.client.Get(ctx, path, resp, opt, sid, mid, tok)
	return resp, err
}

// GetShopSuggestedRate Retrieve suggested rates for all eligible products
// Path: /api/v2/ams/get_shop_suggested_rate
// https://open.shopee.com/documents/v2/v2.ams.get_shop_suggested_rate?module=127&type=1
func (s *AMSServiceOp[T]) GetShopSuggestedRate(ctx context.Context, sid uint64, mid uint64, tok string) (*GetShopSuggestedRateResponse, error) {
	path := "/ams/get_shop_suggested_rate"
	resp := new(GetShopSuggestedRateResponse)
	err := s.client.Get(ctx, path, resp, nil, sid, mid, tok)
	return resp, err
}

// GetTargetedCampaignAddableProductList Returns a list of products that can be added to a targeted campaign
// Path: /api/v2/ams/get_targeted_campaign_addable_product_list
// https://open.shopee.com/documents/v2/v2.ams.get_targeted_campaign_addable_product_list?module=127&type=1
func (s *AMSServiceOp[T]) GetTargetedCampaignAddableProductList(ctx context.Context, sid uint64, mid uint64, tok string, opt GetTargetedCampaignAddableProductListRequest) (*GetTargetedCampaignAddableProductListResponse, error) {
	path := "/ams/get_targeted_campaign_addable_product_list"
	resp := new(GetTargetedCampaignAddableProductListResponse)
	err := s.client.Get(ctx, path, resp, opt, sid, mid, tok)
	return resp, err
}

// GetTargetedCampaignList {"content":"<p>Retrieve all current targeted campaigns created by the seller</p>","raw_content":[{"name":"paragraph","children":[{"data":"Retrieve all current targeted campaigns created by the seller"}]}]}
// Path: /api/v2/ams/get_targeted_campaign_list
// https://open.shopee.com/documents/v2/v2.ams.get_targeted_campaign_list?module=127&type=1
func (s *AMSServiceOp[T]) GetTargetedCampaignList(ctx context.Context, sid uint64, mid uint64, tok string, opt GetTargetedCampaignListRequest) (*GetTargetedCampaignListResponse, error) {
	path := "/ams/get_targeted_campaign_list"
	resp := new(GetTargetedCampaignListResponse)
	err := s.client.Get(ctx, path, resp, opt, sid, mid, tok)
	return resp, err
}

// GetTargetedCampaignPerformance Retrieve a list of Targeted Campaigns and their performance data
// Path: /api/v2/ams/get_targeted_campaign_performance
// https://open.shopee.com/documents/v2/v2.ams.get_targeted_campaign_performance?module=127&type=1
func (s *AMSServiceOp[T]) GetTargetedCampaignPerformance(ctx context.Context, sid uint64, mid uint64, tok string, opt GetTargetedCampaignPerformanceRequest) (*GetTargetedCampaignPerformanceResponse, error) {
	path := "/ams/get_targeted_campaign_performance"
	resp := new(GetTargetedCampaignPerformanceResponse)
	err := s.client.Get(ctx, path, resp, opt, sid, mid, tok)
	return resp, err
}

// GetTargetedCampaignSettings For each campaign, return: campaign basic info (name, status, promotion period, message), selected product list (with product name & ID), selected affiliate list (with affiliate names)
// Path: /api/v2/ams/get_targeted_campaign_settings
// https://open.shopee.com/documents/v2/v2.ams.get_targeted_campaign_settings?module=127&type=1
func (s *AMSServiceOp[T]) GetTargetedCampaignSettings(ctx context.Context, sid uint64, mid uint64, tok string, opt GetTargetedCampaignSettingsRequest) (*GetTargetedCampaignSettingsResponse, error) {
	path := "/ams/get_targeted_campaign_settings"
	resp := new(GetTargetedCampaignSettingsResponse)
	err := s.client.Get(ctx, path, resp, opt, sid, mid, tok)
	return resp, err
}

// GetValidationList Retrieve the seller's AMS validation bill
// Path: /api/v2/ams/get_validation_list
// https://open.shopee.com/documents/v2/v2.ams.get_validation_list?module=127&type=1
func (s *AMSServiceOp[T]) GetValidationList(ctx context.Context, sid uint64, mid uint64, tok string) (*GetValidationListResponse, error) {
	path := "/ams/get_validation_list"
	resp := new(GetValidationListResponse)
	err := s.client.Get(ctx, path, resp, nil, sid, mid, tok)
	return resp, err
}

// GetValidationReport Retrieve detailed information for a specific validation bill
// Path: /api/v2/ams/get_validation_report
// https://open.shopee.com/documents/v2/v2.ams.get_validation_report?module=127&type=1
func (s *AMSServiceOp[T]) GetValidationReport(ctx context.Context, sid uint64, mid uint64, tok string, opt GetValidationReportRequest) (*GetValidationReportResponse, error) {
	path := "/ams/get_validation_report"
	resp := new(GetValidationReportResponse)
	err := s.client.Get(ctx, path, resp, opt, sid, mid, tok)
	return resp, err
}

// QueryAffiliateList Retrieve affiliate information by affiliate id.
// Path: /api/v2/ams/query_affiliate_list
// https://open.shopee.com/documents/v2/v2.ams.query_affiliate_list?module=127&type=1
func (s *AMSServiceOp[T]) QueryAffiliateList(ctx context.Context, sid uint64, mid uint64, tok string, opt QueryAffiliateListRequest) (*QueryAffiliateListResponse, error) {
	path := "/ams/query_affiliate_list"
	resp := new(QueryAffiliateListResponse)
	err := s.client.Get(ctx, path, resp, opt, sid, mid, tok)
	return resp, err
}

// RemoveAllProductsOpenCampaignSetting Remove the entire product list of Open Campaign. We will only return the general error that caused the whole task failure, without returning the specific error for each product in the v2.ams.get_open_campaign_batch_task_result API. If you want to get the result for each products, you can use v2.ams. batch_remove_products_open_campaign_setting by pagination manually, or check the product status by using the GET API after the task progress turn to 100%.
// Path: /api/v2/ams/remove_all_products_open_campaign_setting
// https://open.shopee.com/documents/v2/v2.ams.remove_all_products_open_campaign_setting?module=127&type=1
func (s *AMSServiceOp[T]) RemoveAllProductsOpenCampaignSetting(ctx context.Context, sid uint64, mid uint64, tok string) (*RemoveAllProductsOpenCampaignSettingResponse, error) {
	path := "/ams/remove_all_products_open_campaign_setting"
	resp := new(RemoveAllProductsOpenCampaignSettingResponse)
	err := s.client.Post(ctx, path, nil, resp, sid, mid, tok)
	return resp, err
}

// TerminateTargetedCampaign Change target campaign status to "terminated" to stop all affiliate promotion activity
// Path: /api/v2/ams/terminate_targeted_campaign
// https://open.shopee.com/documents/v2/v2.ams.terminate_targeted_campaign?module=127&type=1
func (s *AMSServiceOp[T]) TerminateTargetedCampaign(ctx context.Context, sid uint64, mid uint64, tok string, req TerminateTargetedCampaignRequest) (*TerminateTargetedCampaignResponse, error) {
	path := "/ams/terminate_targeted_campaign"
	resp := new(TerminateTargetedCampaignResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// UpdateAutoAddNewProductSetting Change auto-add toggle and default commission rate setting
// Path: /api/v2/ams/update_auto_add_new_product_setting
// https://open.shopee.com/documents/v2/v2.ams.update_auto_add_new_product_setting?module=127&type=1
func (s *AMSServiceOp[T]) UpdateAutoAddNewProductSetting(ctx context.Context, sid uint64, mid uint64, tok string, req UpdateAutoAddNewProductSettingRequest) (*UpdateAutoAddNewProductSettingResponse, error) {
	path := "/ams/update_auto_add_new_product_setting"
	resp := new(UpdateAutoAddNewProductSettingResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// UpdateBasicInfoOfTargetedCampaign Edit campaign name, promotion period, message, and budget (if the shop is whitelisted) of target campaign
// Path: /api/v2/ams/update_basic_info_of_targeted_campaign
// https://open.shopee.com/documents/v2/v2.ams.update_basic_info_of_targeted_campaign?module=127&type=1
func (s *AMSServiceOp[T]) UpdateBasicInfoOfTargetedCampaign(ctx context.Context, sid uint64, mid uint64, tok string, req UpdateBasicInfoOfTargetedCampaignRequest) (*UpdateBasicInfoOfTargetedCampaignResponse, error) {
	path := "/ams/update_basic_info_of_targeted_campaign"
	resp := new(UpdateBasicInfoOfTargetedCampaignResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}
