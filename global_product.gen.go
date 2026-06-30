package goshopee

type GlobalProductService interface {
	// AddGlobalItem Add global item. Only for China mainland sellers using China Seller Centre(CNSC). More details in https://shopee.cn/cooperate/46/53/926.
	// Path: /api/v2/global_product/add_global_item
	// https://open.shopee.com/documents/v2/v2.global_product.add_global_item?module=90&type=1
	AddGlobalItem(sid uint64, req AddGlobalItemRequest, tok string) (*AddGlobalItemResponse, error)
	// AddGlobalModel Add global model. Only for China mainland sellers and Korean sellers.
	//
	// Path: /api/v2/global_product/add_global_model
	// https://open.shopee.com/documents/v2/v2.global_product.add_global_model?module=90&type=1
	AddGlobalModel(sid uint64, req AddGlobalModelRequest, tok string) (*AddGlobalModelResponse, error)
	// CategoryRecommend Recommend category by item name. Only for China mainland sellers and Korean sellers.
	// Path: /api/v2/global_product/category_recommend
	// https://open.shopee.com/documents/v2/v2.global_product.category_recommend?module=90&type=1
	CategoryRecommend(sid uint64, opt CategoryRecommendRequest, tok string) (*CategoryRecommendResponse, error)
	// CreatePublishTask Create publish task for global item. Only for China mainland sellers and Korean sellers.
	// Path: /api/v2/global_product/create_publish_task
	// https://open.shopee.com/documents/v2/v2.global_product.create_publish_task?module=90&type=1
	CreatePublishTask(sid uint64, req CreatePublishTaskRequest, tok string) (*CreatePublishTaskResponse, error)
	// DeleteGlobalItem Delete global item. Only for China mainland sellers and Korean sellers.
	// Path: /api/v2/global_product/delete_global_item
	// https://open.shopee.com/documents/v2/v2.global_product.delete_global_item?module=90&type=1
	DeleteGlobalItem(sid uint64, req DeleteGlobalItemRequest, tok string) (*DeleteGlobalItemResponse, error)
	// DeleteGlobalModel Delete global model. Only for China mainland sellers and Korean sellers.
	// Path: /api/v2/global_product/delete_global_model
	// https://open.shopee.com/documents/v2/v2.global_product.delete_global_model?module=90&type=1
	DeleteGlobalModel(sid uint64, req DeleteGlobalModelRequest, tok string) (*DeleteGlobalModelResponse, error)
	// GetAttributeTree Get the mtsku attribute trees for categories
	// Path: /api/v2/global_product/get_attribute_tree
	// https://open.shopee.com/documents/v2/v2.global_product.get_attribute_tree?module=90&type=1
	GetAttributeTree(sid uint64, opt GetAttributeTreeRequest, tok string) (*GetAttributeTreeResponse, error)
	// GetBrandList Use this call to get a list of brand. Only for China mainland sellers and Korean sellers.
	// Path: /api/v2/global_product/get_brand_list
	// https://open.shopee.com/documents/v2/v2.global_product.get_brand_list?module=90&type=1
	GetBrandList(sid uint64, opt GetBrandListRequest, tok string) (*GetBrandListResponse, error)
	// GetCategory Get global category. Only for China mainland sellers and Korean sellers.
	// Path: /api/v2/global_product/get_category
	// https://open.shopee.com/documents/v2/v2.global_product.get_category?module=90&type=1
	GetCategory(sid uint64, opt GetCategoryRequest, tok string) (*GetCategoryResponse, error)
	// GetGlobalItemId Get get_global_item_id by item_id. Only for China mainland sellers and Korean sellers.
	// Path: /api/v2/global_product/get_global_item_id
	// https://open.shopee.com/documents/v2/v2.global_product.get_global_item_id?module=90&type=1
	GetGlobalItemId(sid uint64, opt GetGlobalItemIdRequest, tok string) (*GetGlobalItemIdResponse, error)
	// GetGlobalItemInfo Get global item info.Only for China mainland sellers and Korean sellers.
	// Path: /api/v2/global_product/get_global_item_info
	// https://open.shopee.com/documents/v2/v2.global_product.get_global_item_info?module=90&type=1
	GetGlobalItemInfo(sid uint64, opt GetGlobalItemInfoRequest, tok string) (*GetGlobalItemInfoResponse, error)
	// GetGlobalItemLimit Get global item upload control.
	// Path: /api/v2/global_product/get_global_item_limit
	// https://open.shopee.com/documents/v2/v2.global_product.get_global_item_limit?module=90&type=1
	GetGlobalItemLimit(sid uint64, opt GetGlobalItemLimitRequest, tok string) (*GetGlobalItemLimitResponse, error)
	// GetGlobalItemList Get global item id list. Only for China mainland sellers and Korean sellers.
	// Path: /api/v2/global_product/get_global_item_list
	// https://open.shopee.com/documents/v2/v2.global_product.get_global_item_list?module=90&type=1
	GetGlobalItemList(sid uint64, opt GetGlobalItemListRequest, tok string) (*GetGlobalItemListResponse, error)
	// GetGlobalModelList Get global model list. Only for China mainland sellers and Korean sellers.
	// Path: /api/v2/global_product/get_global_model_list
	// https://open.shopee.com/documents/v2/v2.global_product.get_global_model_list?module=90&type=1
	GetGlobalModelList(sid uint64, opt GetGlobalModelListRequest, tok string) (*GetGlobalModelListResponse, error)
	// GetLocalAdjustmentRate Retrieves the adjustment rate that converts CB stock price into local-warehouse price for a specific shop.
	// Path: /api/v2/global_product/get_local_adjustment_rate
	// https://open.shopee.com/documents/v2/v2.global_product.get_local_adjustment_rate?module=90&type=1
	GetLocalAdjustmentRate(sid uint64, tok string) (*GetLocalAdjustmentRateResponse, error)
	// GetPublishableShop Get publishable shop list for global item. Only for China mainland sellers and Korean sellers.
	//
	// Path: /api/v2/global_product/get_publishable_shop
	// https://open.shopee.com/documents/v2/v2.global_product.get_publishable_shop?module=90&type=1
	GetPublishableShop(sid uint64, opt GetPublishableShopRequest, tok string) (*GetPublishableShopResponse, error)
	// GetPublishedList Get published item list of global item. Only for China mainland sellers and Korean sellers.
	//
	// Path: /api/v2/global_product/get_published_list
	// https://open.shopee.com/documents/v2/v2.global_product.get_published_list?module=90&type=1
	GetPublishedList(sid uint64, opt GetPublishedListRequest, tok string) (*GetPublishedListResponse, error)
	// GetPublishTaskResult Get publish task result for global item. Only for China mainland sellers and Korean sellers.
	// Path: /api/v2/global_product/get_publish_task_result
	// https://open.shopee.com/documents/v2/v2.global_product.get_publish_task_result?module=90&type=1
	GetPublishTaskResult(sid uint64, opt GetPublishTaskResultRequest, tok string) (*GetPublishTaskResultResponse, error)
	// GetRecommendAttribute Get recommend attributes. Only for China mainland sellers and Korean sellers.
	// Path: /api/v2/global_product/get_recommend_attribute
	// https://open.shopee.com/documents/v2/v2.global_product.get_recommend_attribute?module=90&type=1
	GetRecommendAttribute(sid uint64, opt GetRecommendAttributeRequest, tok string) (*GetRecommendAttributeResponse, error)
	// GetShopPublishableStatus Get publishable shop list for global item in pages.
	// Path: /api/v2/global_product/get_shop_publishable_status
	// https://open.shopee.com/documents/v2/v2.global_product.get_shop_publishable_status?module=90&type=1
	GetShopPublishableStatus(sid uint64, opt GetShopPublishableStatusRequest, tok string) (*GetShopPublishableStatusResponse, error)
	// GetSizeChartDetail Get new size chart detail
	// Path: /api/v2/global_product/get_size_chart_detail
	// https://open.shopee.com/documents/v2/v2.global_product.get_size_chart_detail?module=90&type=1
	GetSizeChartDetail(sid uint64, req GetSizeChartDetailRequest, tok string) (*GetSizeChartDetailResponse, error)
	// GetSizeChartList Get size chat list
	// Path: /api/v2/global_product/get_size_chart_list
	// https://open.shopee.com/documents/v2/v2.global_product.get_size_chart_list?module=90&type=1
	GetSizeChartList(sid uint64, req GetSizeChartListRequest, tok string) (*GetSizeChartListResponse, error)
	// GetVariations Get the standardized tier variation defined by Shopee, which is currently a three-layer tree structure. The top layer is variations, the second layer is groups, groups are used to divide options, and the third layer is options.
	// Path: /api/v2/global_product/get_variations
	// https://open.shopee.com/documents/v2/v2.global_product.get_variations?module=90&type=1
	GetVariations(sid uint64, opt GetVariationsRequest, tok string) (*GetVariationsResponse, error)
	// InitTierVariation Only for China mainland sellers and Korean sellers. If you only define color, it is one tier, if you define color and size, it is two tier. Support two tier structures at most. This API can change no tier to one tier, no tier to two tier, one tier to two tier, two tier to one tier, one tier to no tier, two tier to no tier. Please create variants after an interval of 5 seconds after creating an item, as there may be a delay.
	// Path: /api/v2/global_product/init_tier_variation
	// https://open.shopee.com/documents/v2/v2.global_product.init_tier_variation?module=90&type=1
	InitTierVariation(sid uint64, req InitTierVariationRequest, tok string) (*InitTierVariationResponse, error)
	// SearchGlobalAttributeValueList this api is for searching attribute value list for attribute with support_search_value flag
	// Path: /api/v2/global_product/search_global_attribute_value_list
	// https://open.shopee.com/documents/v2/v2.global_product.search_global_attribute_value_list?module=90&type=1
	SearchGlobalAttributeValueList(sid uint64, req SearchGlobalAttributeValueListRequest, tok string) (*SearchGlobalAttributeValueListResponse, error)
	// SetSyncField Set auto sync field. Only for China mainland sellers and Korean sellers.
	// Path: /api/v2/global_product/set_sync_field
	// https://open.shopee.com/documents/v2/v2.global_product.set_sync_field?module=90&type=1
	SetSyncField(sid uint64, req SetSyncFieldRequest, tok string) (*SetSyncFieldResponse, error)
	// SupportSizeChart Get category support size chart. Only for China mainland sellers and Korean sellers.
	//
	// Path: /api/v2/global_product/support_size_chart
	// https://open.shopee.com/documents/v2/v2.global_product.support_size_chart?module=90&type=1
	SupportSizeChart(sid uint64, opt SupportSizeChartRequest, tok string) (*SupportSizeChartResponse, error)
	// UpdateGlobalItem Update global item. Only for China mainland sellers and Korean sellers.
	// Path: /api/v2/global_product/update_global_item
	// https://open.shopee.com/documents/v2/v2.global_product.update_global_item?module=90&type=1
	UpdateGlobalItem(sid uint64, req UpdateGlobalItemRequest, tok string) (*UpdateGlobalItemResponse, error)
	// UpdateGlobalModel Update global model. Only for China mainland sellers and Korean sellers.
	// Path: /api/v2/global_product/update_global_model
	// https://open.shopee.com/documents/v2/v2.global_product.update_global_model?module=90&type=1
	UpdateGlobalModel(sid uint64, req UpdateGlobalModelRequest, tok string) (*UpdateGlobalModelResponse, error)
	// UpdateLocalAdjustmentRate A multiplier that automatically converts your CB stock price into the local-warehouse price. It ensures your local inventory prices reflect regional costs, currency factors, and margin targets.
	// Path: /api/v2/global_product/update_local_adjustment_rate
	// https://open.shopee.com/documents/v2/v2.global_product.update_local_adjustment_rate?module=90&type=1
	UpdateLocalAdjustmentRate(sid uint64, req UpdateLocalAdjustmentRateRequest, tok string) (*UpdateLocalAdjustmentRateResponse, error)
	// UpdatePrice Update global price. Only for China mainland sellers and Korean sellers.
	// Path: /api/v2/global_product/update_price
	// https://open.shopee.com/documents/v2/v2.global_product.update_price?module=90&type=1
	UpdatePrice(sid uint64, req UpdatePriceRequest, tok string) (*UpdatePriceResponse, error)
	// UpdateSizeChart Update size chart for global item. Only for China mainland sellers and Korean sellers.
	//
	// Path: /api/v2/global_product/update_size_chart
	// https://open.shopee.com/documents/v2/v2.global_product.update_size_chart?module=90&type=1
	UpdateSizeChart(sid uint64, req UpdateSizeChartRequest, tok string) (*UpdateSizeChartResponse, error)
	// UpdateStock Update global stock. Only for China mainland sellers and Korean sellers.
	// Path: /api/v2/global_product/update_stock
	// https://open.shopee.com/documents/v2/v2.global_product.update_stock?module=90&type=1
	UpdateStock(sid uint64, req UpdateStockRequest, tok string) (*UpdateStockResponse, error)
	// UpdateTierVariation Update global product tier variation. Only for China mainland sellers and Korean sellers.This api can only be used without changing the tier structure, you can add options, delete options, and update the option image by this api.
	//
	// Path: /api/v2/global_product/update_tier_variation
	// https://open.shopee.com/documents/v2/v2.global_product.update_tier_variation?module=90&type=1
	UpdateTierVariation(sid uint64, req UpdateTierVariationRequest, tok string) (*UpdateTierVariationResponse, error)
}

type GlobalProductServiceOp[T any] struct {
	client *Client[T]
}

// AddGlobalItem Add global item. Only for China mainland sellers using China Seller Centre(CNSC). More details in https://shopee.cn/cooperate/46/53/926.
// Path: /api/v2/global_product/add_global_item
// https://open.shopee.com/documents/v2/v2.global_product.add_global_item?module=90&type=1
func (s *GlobalProductServiceOp[T]) AddGlobalItem(sid uint64, req AddGlobalItemRequest, tok string) (*AddGlobalItemResponse, error) {
	path := "/global_product/add_global_item"
	resp := new(AddGlobalItemResponse)
	err := s.client.WithMerchant(sid, tok).Post(path, req, resp)
	return resp, err
}

// AddGlobalModel Add global model. Only for China mainland sellers and Korean sellers.
//
// Path: /api/v2/global_product/add_global_model
// https://open.shopee.com/documents/v2/v2.global_product.add_global_model?module=90&type=1
func (s *GlobalProductServiceOp[T]) AddGlobalModel(sid uint64, req AddGlobalModelRequest, tok string) (*AddGlobalModelResponse, error) {
	path := "/global_product/add_global_model"
	resp := new(AddGlobalModelResponse)
	err := s.client.WithMerchant(sid, tok).Post(path, req, resp)
	return resp, err
}

// CategoryRecommend Recommend category by item name. Only for China mainland sellers and Korean sellers.
// Path: /api/v2/global_product/category_recommend
// https://open.shopee.com/documents/v2/v2.global_product.category_recommend?module=90&type=1
func (s *GlobalProductServiceOp[T]) CategoryRecommend(sid uint64, opt CategoryRecommendRequest, tok string) (*CategoryRecommendResponse, error) {
	path := "/global_product/category_recommend"
	resp := new(CategoryRecommendResponse)
	err := s.client.WithMerchant(sid, tok).Get(path, resp, opt)
	return resp, err
}

// CreatePublishTask Create publish task for global item. Only for China mainland sellers and Korean sellers.
// Path: /api/v2/global_product/create_publish_task
// https://open.shopee.com/documents/v2/v2.global_product.create_publish_task?module=90&type=1
func (s *GlobalProductServiceOp[T]) CreatePublishTask(sid uint64, req CreatePublishTaskRequest, tok string) (*CreatePublishTaskResponse, error) {
	path := "/global_product/create_publish_task"
	resp := new(CreatePublishTaskResponse)
	err := s.client.WithMerchant(sid, tok).Post(path, req, resp)
	return resp, err
}

// DeleteGlobalItem Delete global item. Only for China mainland sellers and Korean sellers.
// Path: /api/v2/global_product/delete_global_item
// https://open.shopee.com/documents/v2/v2.global_product.delete_global_item?module=90&type=1
func (s *GlobalProductServiceOp[T]) DeleteGlobalItem(sid uint64, req DeleteGlobalItemRequest, tok string) (*DeleteGlobalItemResponse, error) {
	path := "/global_product/delete_global_item"
	resp := new(DeleteGlobalItemResponse)
	err := s.client.WithMerchant(sid, tok).Post(path, req, resp)
	return resp, err
}

// DeleteGlobalModel Delete global model. Only for China mainland sellers and Korean sellers.
// Path: /api/v2/global_product/delete_global_model
// https://open.shopee.com/documents/v2/v2.global_product.delete_global_model?module=90&type=1
func (s *GlobalProductServiceOp[T]) DeleteGlobalModel(sid uint64, req DeleteGlobalModelRequest, tok string) (*DeleteGlobalModelResponse, error) {
	path := "/global_product/delete_global_model"
	resp := new(DeleteGlobalModelResponse)
	err := s.client.WithMerchant(sid, tok).Post(path, req, resp)
	return resp, err
}

// GetAttributeTree Get the mtsku attribute trees for categories
// Path: /api/v2/global_product/get_attribute_tree
// https://open.shopee.com/documents/v2/v2.global_product.get_attribute_tree?module=90&type=1
func (s *GlobalProductServiceOp[T]) GetAttributeTree(sid uint64, opt GetAttributeTreeRequest, tok string) (*GetAttributeTreeResponse, error) {
	path := "/global_product/get_attribute_tree"
	resp := new(GetAttributeTreeResponse)
	err := s.client.WithMerchant(sid, tok).Get(path, resp, opt)
	return resp, err
}

// GetBrandList Use this call to get a list of brand. Only for China mainland sellers and Korean sellers.
// Path: /api/v2/global_product/get_brand_list
// https://open.shopee.com/documents/v2/v2.global_product.get_brand_list?module=90&type=1
func (s *GlobalProductServiceOp[T]) GetBrandList(sid uint64, opt GetBrandListRequest, tok string) (*GetBrandListResponse, error) {
	path := "/global_product/get_brand_list"
	resp := new(GetBrandListResponse)
	err := s.client.WithMerchant(sid, tok).Get(path, resp, opt)
	return resp, err
}

// GetCategory Get global category. Only for China mainland sellers and Korean sellers.
// Path: /api/v2/global_product/get_category
// https://open.shopee.com/documents/v2/v2.global_product.get_category?module=90&type=1
func (s *GlobalProductServiceOp[T]) GetCategory(sid uint64, opt GetCategoryRequest, tok string) (*GetCategoryResponse, error) {
	path := "/global_product/get_category"
	resp := new(GetCategoryResponse)
	err := s.client.WithMerchant(sid, tok).Get(path, resp, opt)
	return resp, err
}

// GetGlobalItemId Get get_global_item_id by item_id. Only for China mainland sellers and Korean sellers.
// Path: /api/v2/global_product/get_global_item_id
// https://open.shopee.com/documents/v2/v2.global_product.get_global_item_id?module=90&type=1
func (s *GlobalProductServiceOp[T]) GetGlobalItemId(sid uint64, opt GetGlobalItemIdRequest, tok string) (*GetGlobalItemIdResponse, error) {
	path := "/global_product/get_global_item_id"
	resp := new(GetGlobalItemIdResponse)
	err := s.client.WithMerchant(sid, tok).Get(path, resp, opt)
	return resp, err
}

// GetGlobalItemInfo Get global item info.Only for China mainland sellers and Korean sellers.
// Path: /api/v2/global_product/get_global_item_info
// https://open.shopee.com/documents/v2/v2.global_product.get_global_item_info?module=90&type=1
func (s *GlobalProductServiceOp[T]) GetGlobalItemInfo(sid uint64, opt GetGlobalItemInfoRequest, tok string) (*GetGlobalItemInfoResponse, error) {
	path := "/global_product/get_global_item_info"
	resp := new(GetGlobalItemInfoResponse)
	err := s.client.WithMerchant(sid, tok).Get(path, resp, opt)
	return resp, err
}

// GetGlobalItemLimit Get global item upload control.
// Path: /api/v2/global_product/get_global_item_limit
// https://open.shopee.com/documents/v2/v2.global_product.get_global_item_limit?module=90&type=1
func (s *GlobalProductServiceOp[T]) GetGlobalItemLimit(sid uint64, opt GetGlobalItemLimitRequest, tok string) (*GetGlobalItemLimitResponse, error) {
	path := "/global_product/get_global_item_limit"
	resp := new(GetGlobalItemLimitResponse)
	err := s.client.WithMerchant(sid, tok).Get(path, resp, opt)
	return resp, err
}

// GetGlobalItemList Get global item id list. Only for China mainland sellers and Korean sellers.
// Path: /api/v2/global_product/get_global_item_list
// https://open.shopee.com/documents/v2/v2.global_product.get_global_item_list?module=90&type=1
func (s *GlobalProductServiceOp[T]) GetGlobalItemList(sid uint64, opt GetGlobalItemListRequest, tok string) (*GetGlobalItemListResponse, error) {
	path := "/global_product/get_global_item_list"
	resp := new(GetGlobalItemListResponse)
	err := s.client.WithMerchant(sid, tok).Get(path, resp, opt)
	return resp, err
}

// GetGlobalModelList Get global model list. Only for China mainland sellers and Korean sellers.
// Path: /api/v2/global_product/get_global_model_list
// https://open.shopee.com/documents/v2/v2.global_product.get_global_model_list?module=90&type=1
func (s *GlobalProductServiceOp[T]) GetGlobalModelList(sid uint64, opt GetGlobalModelListRequest, tok string) (*GetGlobalModelListResponse, error) {
	path := "/global_product/get_global_model_list"
	resp := new(GetGlobalModelListResponse)
	err := s.client.WithMerchant(sid, tok).Get(path, resp, opt)
	return resp, err
}

// GetLocalAdjustmentRate Retrieves the adjustment rate that converts CB stock price into local-warehouse price for a specific shop.
// Path: /api/v2/global_product/get_local_adjustment_rate
// https://open.shopee.com/documents/v2/v2.global_product.get_local_adjustment_rate?module=90&type=1
func (s *GlobalProductServiceOp[T]) GetLocalAdjustmentRate(sid uint64, tok string) (*GetLocalAdjustmentRateResponse, error) {
	path := "/global_product/get_local_adjustment_rate"
	resp := new(GetLocalAdjustmentRateResponse)
	err := s.client.WithMerchant(sid, tok).Post(path, nil, resp)
	return resp, err
}

// GetPublishableShop Get publishable shop list for global item. Only for China mainland sellers and Korean sellers.
//
// Path: /api/v2/global_product/get_publishable_shop
// https://open.shopee.com/documents/v2/v2.global_product.get_publishable_shop?module=90&type=1
func (s *GlobalProductServiceOp[T]) GetPublishableShop(sid uint64, opt GetPublishableShopRequest, tok string) (*GetPublishableShopResponse, error) {
	path := "/global_product/get_publishable_shop"
	resp := new(GetPublishableShopResponse)
	err := s.client.WithMerchant(sid, tok).Get(path, resp, opt)
	return resp, err
}

// GetPublishedList Get published item list of global item. Only for China mainland sellers and Korean sellers.
//
// Path: /api/v2/global_product/get_published_list
// https://open.shopee.com/documents/v2/v2.global_product.get_published_list?module=90&type=1
func (s *GlobalProductServiceOp[T]) GetPublishedList(sid uint64, opt GetPublishedListRequest, tok string) (*GetPublishedListResponse, error) {
	path := "/global_product/get_published_list"
	resp := new(GetPublishedListResponse)
	err := s.client.WithMerchant(sid, tok).Get(path, resp, opt)
	return resp, err
}

// GetPublishTaskResult Get publish task result for global item. Only for China mainland sellers and Korean sellers.
// Path: /api/v2/global_product/get_publish_task_result
// https://open.shopee.com/documents/v2/v2.global_product.get_publish_task_result?module=90&type=1
func (s *GlobalProductServiceOp[T]) GetPublishTaskResult(sid uint64, opt GetPublishTaskResultRequest, tok string) (*GetPublishTaskResultResponse, error) {
	path := "/global_product/get_publish_task_result"
	resp := new(GetPublishTaskResultResponse)
	err := s.client.WithMerchant(sid, tok).Get(path, resp, opt)
	return resp, err
}

// GetRecommendAttribute Get recommend attributes. Only for China mainland sellers and Korean sellers.
// Path: /api/v2/global_product/get_recommend_attribute
// https://open.shopee.com/documents/v2/v2.global_product.get_recommend_attribute?module=90&type=1
func (s *GlobalProductServiceOp[T]) GetRecommendAttribute(sid uint64, opt GetRecommendAttributeRequest, tok string) (*GetRecommendAttributeResponse, error) {
	path := "/global_product/get_recommend_attribute"
	resp := new(GetRecommendAttributeResponse)
	err := s.client.WithMerchant(sid, tok).Get(path, resp, opt)
	return resp, err
}

// GetShopPublishableStatus Get publishable shop list for global item in pages.
// Path: /api/v2/global_product/get_shop_publishable_status
// https://open.shopee.com/documents/v2/v2.global_product.get_shop_publishable_status?module=90&type=1
func (s *GlobalProductServiceOp[T]) GetShopPublishableStatus(sid uint64, opt GetShopPublishableStatusRequest, tok string) (*GetShopPublishableStatusResponse, error) {
	path := "/global_product/get_shop_publishable_status"
	resp := new(GetShopPublishableStatusResponse)
	err := s.client.WithMerchant(sid, tok).Get(path, resp, opt)
	return resp, err
}

// GetSizeChartDetail Get new size chart detail
// Path: /api/v2/global_product/get_size_chart_detail
// https://open.shopee.com/documents/v2/v2.global_product.get_size_chart_detail?module=90&type=1
func (s *GlobalProductServiceOp[T]) GetSizeChartDetail(sid uint64, req GetSizeChartDetailRequest, tok string) (*GetSizeChartDetailResponse, error) {
	path := "/global_product/get_size_chart_detail"
	resp := new(GetSizeChartDetailResponse)
	err := s.client.WithMerchant(sid, tok).Post(path, req, resp)
	return resp, err
}

// GetSizeChartList Get size chat list
// Path: /api/v2/global_product/get_size_chart_list
// https://open.shopee.com/documents/v2/v2.global_product.get_size_chart_list?module=90&type=1
func (s *GlobalProductServiceOp[T]) GetSizeChartList(sid uint64, req GetSizeChartListRequest, tok string) (*GetSizeChartListResponse, error) {
	path := "/global_product/get_size_chart_list"
	resp := new(GetSizeChartListResponse)
	err := s.client.WithMerchant(sid, tok).Post(path, req, resp)
	return resp, err
}

// GetVariations Get the standardized tier variation defined by Shopee, which is currently a three-layer tree structure. The top layer is variations, the second layer is groups, groups are used to divide options, and the third layer is options.
// Path: /api/v2/global_product/get_variations
// https://open.shopee.com/documents/v2/v2.global_product.get_variations?module=90&type=1
func (s *GlobalProductServiceOp[T]) GetVariations(sid uint64, opt GetVariationsRequest, tok string) (*GetVariationsResponse, error) {
	path := "/global_product/get_variations"
	resp := new(GetVariationsResponse)
	err := s.client.WithMerchant(sid, tok).Get(path, resp, opt)
	return resp, err
}

// InitTierVariation Only for China mainland sellers and Korean sellers. If you only define color, it is one tier, if you define color and size, it is two tier. Support two tier structures at most. This API can change no tier to one tier, no tier to two tier, one tier to two tier, two tier to one tier, one tier to no tier, two tier to no tier. Please create variants after an interval of 5 seconds after creating an item, as there may be a delay.
// Path: /api/v2/global_product/init_tier_variation
// https://open.shopee.com/documents/v2/v2.global_product.init_tier_variation?module=90&type=1
func (s *GlobalProductServiceOp[T]) InitTierVariation(sid uint64, req InitTierVariationRequest, tok string) (*InitTierVariationResponse, error) {
	path := "/global_product/init_tier_variation"
	resp := new(InitTierVariationResponse)
	err := s.client.WithMerchant(sid, tok).Post(path, req, resp)
	return resp, err
}

// SearchGlobalAttributeValueList this api is for searching attribute value list for attribute with support_search_value flag
// Path: /api/v2/global_product/search_global_attribute_value_list
// https://open.shopee.com/documents/v2/v2.global_product.search_global_attribute_value_list?module=90&type=1
func (s *GlobalProductServiceOp[T]) SearchGlobalAttributeValueList(sid uint64, req SearchGlobalAttributeValueListRequest, tok string) (*SearchGlobalAttributeValueListResponse, error) {
	path := "/global_product/search_global_attribute_value_list"
	resp := new(SearchGlobalAttributeValueListResponse)
	err := s.client.WithMerchant(sid, tok).Post(path, req, resp)
	return resp, err
}

// SetSyncField Set auto sync field. Only for China mainland sellers and Korean sellers.
// Path: /api/v2/global_product/set_sync_field
// https://open.shopee.com/documents/v2/v2.global_product.set_sync_field?module=90&type=1
func (s *GlobalProductServiceOp[T]) SetSyncField(sid uint64, req SetSyncFieldRequest, tok string) (*SetSyncFieldResponse, error) {
	path := "/global_product/set_sync_field"
	resp := new(SetSyncFieldResponse)
	err := s.client.WithMerchant(sid, tok).Post(path, req, resp)
	return resp, err
}

// SupportSizeChart Get category support size chart. Only for China mainland sellers and Korean sellers.
//
// Path: /api/v2/global_product/support_size_chart
// https://open.shopee.com/documents/v2/v2.global_product.support_size_chart?module=90&type=1
func (s *GlobalProductServiceOp[T]) SupportSizeChart(sid uint64, opt SupportSizeChartRequest, tok string) (*SupportSizeChartResponse, error) {
	path := "/global_product/support_size_chart"
	resp := new(SupportSizeChartResponse)
	err := s.client.WithMerchant(sid, tok).Get(path, resp, opt)
	return resp, err
}

// UpdateGlobalItem Update global item. Only for China mainland sellers and Korean sellers.
// Path: /api/v2/global_product/update_global_item
// https://open.shopee.com/documents/v2/v2.global_product.update_global_item?module=90&type=1
func (s *GlobalProductServiceOp[T]) UpdateGlobalItem(sid uint64, req UpdateGlobalItemRequest, tok string) (*UpdateGlobalItemResponse, error) {
	path := "/global_product/update_global_item"
	resp := new(UpdateGlobalItemResponse)
	err := s.client.WithMerchant(sid, tok).Post(path, req, resp)
	return resp, err
}

// UpdateGlobalModel Update global model. Only for China mainland sellers and Korean sellers.
// Path: /api/v2/global_product/update_global_model
// https://open.shopee.com/documents/v2/v2.global_product.update_global_model?module=90&type=1
func (s *GlobalProductServiceOp[T]) UpdateGlobalModel(sid uint64, req UpdateGlobalModelRequest, tok string) (*UpdateGlobalModelResponse, error) {
	path := "/global_product/update_global_model"
	resp := new(UpdateGlobalModelResponse)
	err := s.client.WithMerchant(sid, tok).Post(path, req, resp)
	return resp, err
}

// UpdateLocalAdjustmentRate A multiplier that automatically converts your CB stock price into the local-warehouse price. It ensures your local inventory prices reflect regional costs, currency factors, and margin targets.
// Path: /api/v2/global_product/update_local_adjustment_rate
// https://open.shopee.com/documents/v2/v2.global_product.update_local_adjustment_rate?module=90&type=1
func (s *GlobalProductServiceOp[T]) UpdateLocalAdjustmentRate(sid uint64, req UpdateLocalAdjustmentRateRequest, tok string) (*UpdateLocalAdjustmentRateResponse, error) {
	path := "/global_product/update_local_adjustment_rate"
	resp := new(UpdateLocalAdjustmentRateResponse)
	err := s.client.WithMerchant(sid, tok).Post(path, req, resp)
	return resp, err
}

// UpdatePrice Update global price. Only for China mainland sellers and Korean sellers.
// Path: /api/v2/global_product/update_price
// https://open.shopee.com/documents/v2/v2.global_product.update_price?module=90&type=1
func (s *GlobalProductServiceOp[T]) UpdatePrice(sid uint64, req UpdatePriceRequest, tok string) (*UpdatePriceResponse, error) {
	path := "/global_product/update_price"
	resp := new(UpdatePriceResponse)
	err := s.client.WithMerchant(sid, tok).Post(path, req, resp)
	return resp, err
}

// UpdateSizeChart Update size chart for global item. Only for China mainland sellers and Korean sellers.
//
// Path: /api/v2/global_product/update_size_chart
// https://open.shopee.com/documents/v2/v2.global_product.update_size_chart?module=90&type=1
func (s *GlobalProductServiceOp[T]) UpdateSizeChart(sid uint64, req UpdateSizeChartRequest, tok string) (*UpdateSizeChartResponse, error) {
	path := "/global_product/update_size_chart"
	resp := new(UpdateSizeChartResponse)
	err := s.client.WithMerchant(sid, tok).Post(path, req, resp)
	return resp, err
}

// UpdateStock Update global stock. Only for China mainland sellers and Korean sellers.
// Path: /api/v2/global_product/update_stock
// https://open.shopee.com/documents/v2/v2.global_product.update_stock?module=90&type=1
func (s *GlobalProductServiceOp[T]) UpdateStock(sid uint64, req UpdateStockRequest, tok string) (*UpdateStockResponse, error) {
	path := "/global_product/update_stock"
	resp := new(UpdateStockResponse)
	err := s.client.WithMerchant(sid, tok).Post(path, req, resp)
	return resp, err
}

// UpdateTierVariation Update global product tier variation. Only for China mainland sellers and Korean sellers.This api can only be used without changing the tier structure, you can add options, delete options, and update the option image by this api.
//
// Path: /api/v2/global_product/update_tier_variation
// https://open.shopee.com/documents/v2/v2.global_product.update_tier_variation?module=90&type=1
func (s *GlobalProductServiceOp[T]) UpdateTierVariation(sid uint64, req UpdateTierVariationRequest, tok string) (*UpdateTierVariationResponse, error) {
	path := "/global_product/update_tier_variation"
	resp := new(UpdateTierVariationResponse)
	err := s.client.WithMerchant(sid, tok).Post(path, req, resp)
	return resp, err
}
