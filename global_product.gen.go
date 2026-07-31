package goshopee

import (
	"context"
)

type GlobalProductService interface {
	// AddGlobalItem Add global item. Only for China mainland sellers using China Seller Centre(CNSC). More details in https://shopee.cn/cooperate/46/53/926.
	// Path: /api/v2/global_product/add_global_item
	// https://open.shopee.com/documents/v2/v2.global_product.add_global_item?module=90&type=1
	AddGlobalItem(ctx context.Context, sid uint64, mid uint64, tok string, req AddGlobalItemRequest) (*AddGlobalItemResponse, error)
	// AddGlobalModel Add global model. Only for China mainland sellers and Korean sellers.
	//
	// Path: /api/v2/global_product/add_global_model
	// https://open.shopee.com/documents/v2/v2.global_product.add_global_model?module=90&type=1
	AddGlobalModel(ctx context.Context, sid uint64, mid uint64, tok string, req AddGlobalModelRequest) (*AddGlobalModelResponse, error)
	// CategoryRecommend Recommend category by item name. Only for China mainland sellers and Korean sellers.
	// Path: /api/v2/global_product/category_recommend
	// https://open.shopee.com/documents/v2/v2.global_product.category_recommend?module=90&type=1
	CategoryRecommend(ctx context.Context, sid uint64, mid uint64, tok string, opt CategoryRecommendRequest) (*CategoryRecommendResponse, error)
	// CreatePublishTask Create publish task for global item. Only for China mainland sellers and Korean sellers.
	// Path: /api/v2/global_product/create_publish_task
	// https://open.shopee.com/documents/v2/v2.global_product.create_publish_task?module=90&type=1
	CreatePublishTask(ctx context.Context, sid uint64, mid uint64, tok string, req CreatePublishTaskRequest) (*CreatePublishTaskResponse, error)
	// DeleteGlobalItem Delete global item. Only for China mainland sellers and Korean sellers.
	// Path: /api/v2/global_product/delete_global_item
	// https://open.shopee.com/documents/v2/v2.global_product.delete_global_item?module=90&type=1
	DeleteGlobalItem(ctx context.Context, sid uint64, mid uint64, tok string, req DeleteGlobalItemRequest) (*DeleteGlobalItemResponse, error)
	// DeleteGlobalModel Delete global model. Only for China mainland sellers and Korean sellers.
	// Path: /api/v2/global_product/delete_global_model
	// https://open.shopee.com/documents/v2/v2.global_product.delete_global_model?module=90&type=1
	DeleteGlobalModel(ctx context.Context, sid uint64, mid uint64, tok string, req DeleteGlobalModelRequest) (*DeleteGlobalModelResponse, error)
	// GetAttributeTree Get the mtsku attribute trees for categories
	// Path: /api/v2/global_product/get_attribute_tree
	// https://open.shopee.com/documents/v2/v2.global_product.get_attribute_tree?module=90&type=1
	GetAttributeTree(ctx context.Context, sid uint64, mid uint64, tok string, opt GetAttributeTreeRequest) (*GetAttributeTreeResponse, error)
	// GetBrandList Use this call to get a list of brand. Only for China mainland sellers and Korean sellers.
	// Path: /api/v2/global_product/get_brand_list
	// https://open.shopee.com/documents/v2/v2.global_product.get_brand_list?module=90&type=1
	GetBrandList(ctx context.Context, sid uint64, mid uint64, tok string, opt GetBrandListRequest) (*GetBrandListResponse, error)
	// GetCategory Get global category. Only for China mainland sellers and Korean sellers.
	// Path: /api/v2/global_product/get_category
	// https://open.shopee.com/documents/v2/v2.global_product.get_category?module=90&type=1
	GetCategory(ctx context.Context, sid uint64, mid uint64, tok string, opt GetCategoryRequest) (*GetCategoryResponse, error)
	// GetGlobalItemId Get get_global_item_id by item_id. Only for China mainland sellers and Korean sellers.
	// Path: /api/v2/global_product/get_global_item_id
	// https://open.shopee.com/documents/v2/v2.global_product.get_global_item_id?module=90&type=1
	GetGlobalItemId(ctx context.Context, sid uint64, mid uint64, tok string, opt GetGlobalItemIdRequest) (*GetGlobalItemIdResponse, error)
	// GetGlobalItemInfo Get global item info.Only for China mainland sellers and Korean sellers.
	// Path: /api/v2/global_product/get_global_item_info
	// https://open.shopee.com/documents/v2/v2.global_product.get_global_item_info?module=90&type=1
	GetGlobalItemInfo(ctx context.Context, sid uint64, mid uint64, tok string, opt GetGlobalItemInfoRequest) (*GetGlobalItemInfoResponse, error)
	// GetGlobalItemLimit Get global item upload control.
	// Path: /api/v2/global_product/get_global_item_limit
	// https://open.shopee.com/documents/v2/v2.global_product.get_global_item_limit?module=90&type=1
	GetGlobalItemLimit(ctx context.Context, sid uint64, mid uint64, tok string, opt GetGlobalItemLimitRequest) (*GetGlobalItemLimitResponse, error)
	// GetGlobalItemList Get global item id list. Only for China mainland sellers and Korean sellers.
	// Path: /api/v2/global_product/get_global_item_list
	// https://open.shopee.com/documents/v2/v2.global_product.get_global_item_list?module=90&type=1
	GetGlobalItemList(ctx context.Context, sid uint64, mid uint64, tok string, opt GetGlobalItemListRequest) (*GetGlobalItemListResponse, error)
	// GetGlobalModelList Get global model list. Only for China mainland sellers and Korean sellers.
	// Path: /api/v2/global_product/get_global_model_list
	// https://open.shopee.com/documents/v2/v2.global_product.get_global_model_list?module=90&type=1
	GetGlobalModelList(ctx context.Context, sid uint64, mid uint64, tok string, opt GetGlobalModelListRequest) (*GetGlobalModelListResponse, error)
	// GetLocalAdjustmentRate Retrieves the adjustment rate that converts CB stock price into local-warehouse price for a specific shop.
	// Path: /api/v2/global_product/get_local_adjustment_rate
	// https://open.shopee.com/documents/v2/v2.global_product.get_local_adjustment_rate?module=90&type=1
	GetLocalAdjustmentRate(ctx context.Context, sid uint64, mid uint64, tok string) (*GetLocalAdjustmentRateResponse, error)
	// GetPublishableShop Get publishable shop list for global item. Only for China mainland sellers and Korean sellers.
	//
	// Path: /api/v2/global_product/get_publishable_shop
	// https://open.shopee.com/documents/v2/v2.global_product.get_publishable_shop?module=90&type=1
	GetPublishableShop(ctx context.Context, sid uint64, mid uint64, tok string, opt GetPublishableShopRequest) (*GetPublishableShopResponse, error)
	// GetPublishedList Get published item list of global item. Only for China mainland sellers and Korean sellers.
	//
	// Path: /api/v2/global_product/get_published_list
	// https://open.shopee.com/documents/v2/v2.global_product.get_published_list?module=90&type=1
	GetPublishedList(ctx context.Context, sid uint64, mid uint64, tok string, opt GetPublishedListRequest) (*GetPublishedListResponse, error)
	// GetPublishTaskResult Get publish task result for global item. Only for China mainland sellers and Korean sellers.
	// Path: /api/v2/global_product/get_publish_task_result
	// https://open.shopee.com/documents/v2/v2.global_product.get_publish_task_result?module=90&type=1
	GetPublishTaskResult(ctx context.Context, sid uint64, mid uint64, tok string, opt GetPublishTaskResultRequest) (*GetPublishTaskResultResponse, error)
	// GetRecommendAttribute Get recommend attributes. Only for China mainland sellers and Korean sellers.
	// Path: /api/v2/global_product/get_recommend_attribute
	// https://open.shopee.com/documents/v2/v2.global_product.get_recommend_attribute?module=90&type=1
	GetRecommendAttribute(ctx context.Context, sid uint64, mid uint64, tok string, opt GetRecommendAttributeRequest) (*GetRecommendAttributeResponse, error)
	// GetShopPublishableStatus Get publishable shop list for global item in pages.
	// Path: /api/v2/global_product/get_shop_publishable_status
	// https://open.shopee.com/documents/v2/v2.global_product.get_shop_publishable_status?module=90&type=1
	GetShopPublishableStatus(ctx context.Context, sid uint64, mid uint64, tok string, opt GetShopPublishableStatusRequest) (*GetShopPublishableStatusResponse, error)
	// GetSizeChartDetail Get new size chart detail
	// Path: /api/v2/global_product/get_size_chart_detail
	// https://open.shopee.com/documents/v2/v2.global_product.get_size_chart_detail?module=90&type=1
	GetSizeChartDetail(ctx context.Context, sid uint64, mid uint64, tok string, req GetSizeChartDetailRequest) (*GetSizeChartDetailResponse, error)
	// GetSizeChartList {"content":"<p>Get size chart list</p>","raw_content":[{"name":"paragraph","children":[{"data":"Get size chart list"}]}]}
	// Path: /api/v2/global_product/get_size_chart_list
	// https://open.shopee.com/documents/v2/v2.global_product.get_size_chart_list?module=90&type=1
	GetSizeChartList(ctx context.Context, sid uint64, mid uint64, tok string, req GetSizeChartListRequest) (*GetSizeChartListResponse, error)
	// GetVariations Get the standardized tier variation defined by Shopee, which is currently a three-layer tree structure. The top layer is variations, the second layer is groups, groups are used to divide options, and the third layer is options.
	// Path: /api/v2/global_product/get_variations
	// https://open.shopee.com/documents/v2/v2.global_product.get_variations?module=90&type=1
	GetVariations(ctx context.Context, sid uint64, mid uint64, tok string, opt GetVariationsRequest) (*GetVariationsResponse, error)
	// InitTierVariation Only for China mainland sellers and Korean sellers. If you only define color, it is one tier, if you define color and size, it is two tier. Support two tier structures at most. This API can change no tier to one tier, no tier to two tier, one tier to two tier, two tier to one tier, one tier to no tier, two tier to no tier. Please create variants after an interval of 5 seconds after creating an item, as there may be a delay.
	// Path: /api/v2/global_product/init_tier_variation
	// https://open.shopee.com/documents/v2/v2.global_product.init_tier_variation?module=90&type=1
	InitTierVariation(ctx context.Context, sid uint64, mid uint64, tok string, req InitTierVariationRequest) (*InitTierVariationResponse, error)
	// SearchGlobalAttributeValueList this api is for searching attribute value list for attribute with support_search_value flag
	// Path: /api/v2/global_product/search_global_attribute_value_list
	// https://open.shopee.com/documents/v2/v2.global_product.search_global_attribute_value_list?module=90&type=1
	SearchGlobalAttributeValueList(ctx context.Context, sid uint64, mid uint64, tok string, req SearchGlobalAttributeValueListRequest) (*SearchGlobalAttributeValueListResponse, error)
	// SetSyncField Set auto sync field. Only for China mainland sellers and Korean sellers.
	// Path: /api/v2/global_product/set_sync_field
	// https://open.shopee.com/documents/v2/v2.global_product.set_sync_field?module=90&type=1
	SetSyncField(ctx context.Context, sid uint64, mid uint64, tok string, req SetSyncFieldRequest) (*SetSyncFieldResponse, error)
	// SupportSizeChart Get category support size chart. Only for China mainland sellers and Korean sellers.
	//
	// Path: /api/v2/global_product/support_size_chart
	// https://open.shopee.com/documents/v2/v2.global_product.support_size_chart?module=90&type=1
	SupportSizeChart(ctx context.Context, sid uint64, mid uint64, tok string, opt SupportSizeChartRequest) (*SupportSizeChartResponse, error)
	// UpdateGlobalItem Update global item. Only for China mainland sellers and Korean sellers.
	// Path: /api/v2/global_product/update_global_item
	// https://open.shopee.com/documents/v2/v2.global_product.update_global_item?module=90&type=1
	UpdateGlobalItem(ctx context.Context, sid uint64, mid uint64, tok string, req UpdateGlobalItemRequest) (*UpdateGlobalItemResponse, error)
	// UpdateGlobalModel Update global model. Only for China mainland sellers and Korean sellers.
	// Path: /api/v2/global_product/update_global_model
	// https://open.shopee.com/documents/v2/v2.global_product.update_global_model?module=90&type=1
	UpdateGlobalModel(ctx context.Context, sid uint64, mid uint64, tok string, req UpdateGlobalModelRequest) (*UpdateGlobalModelResponse, error)
	// UpdateLocalAdjustmentRate A multiplier that automatically converts your CB stock price into the local-warehouse price. It ensures your local inventory prices reflect regional costs, currency factors, and margin targets.
	// Path: /api/v2/global_product/update_local_adjustment_rate
	// https://open.shopee.com/documents/v2/v2.global_product.update_local_adjustment_rate?module=90&type=1
	UpdateLocalAdjustmentRate(ctx context.Context, sid uint64, mid uint64, tok string, req UpdateLocalAdjustmentRateRequest) (*UpdateLocalAdjustmentRateResponse, error)
	// UpdatePrice Update global price. Only for China mainland sellers and Korean sellers.
	// Path: /api/v2/global_product/update_price
	// https://open.shopee.com/documents/v2/v2.global_product.update_price?module=90&type=1
	UpdatePrice(ctx context.Context, sid uint64, mid uint64, tok string, req UpdatePriceRequest) (*UpdatePriceResponse, error)
	// UpdateSizeChart Update size chart for global item. Only for China mainland sellers and Korean sellers.
	//
	// Path: /api/v2/global_product/update_size_chart
	// https://open.shopee.com/documents/v2/v2.global_product.update_size_chart?module=90&type=1
	UpdateSizeChart(ctx context.Context, sid uint64, mid uint64, tok string, req UpdateSizeChartRequest) (*UpdateSizeChartResponse, error)
	// UpdateStock Update global stock. Only for China mainland sellers and Korean sellers.
	// Path: /api/v2/global_product/update_stock
	// https://open.shopee.com/documents/v2/v2.global_product.update_stock?module=90&type=1
	UpdateStock(ctx context.Context, sid uint64, mid uint64, tok string, req UpdateStockRequest) (*UpdateStockResponse, error)
	// UpdateTierVariation Update global product tier variation. Only for China mainland sellers and Korean sellers.This api can only be used without changing the tier structure, you can add options, delete options, and update the option image by this api.
	//
	// Path: /api/v2/global_product/update_tier_variation
	// https://open.shopee.com/documents/v2/v2.global_product.update_tier_variation?module=90&type=1
	UpdateTierVariation(ctx context.Context, sid uint64, mid uint64, tok string, req UpdateTierVariationRequest) (*UpdateTierVariationResponse, error)
}

type GlobalProductServiceOp[T any] struct {
	client *Client[T]
}

// AddGlobalItem Add global item. Only for China mainland sellers using China Seller Centre(CNSC). More details in https://shopee.cn/cooperate/46/53/926.
// Path: /api/v2/global_product/add_global_item
// https://open.shopee.com/documents/v2/v2.global_product.add_global_item?module=90&type=1
func (s *GlobalProductServiceOp[T]) AddGlobalItem(ctx context.Context, sid uint64, mid uint64, tok string, req AddGlobalItemRequest) (*AddGlobalItemResponse, error) {
	path := "/global_product/add_global_item"
	resp := new(AddGlobalItemResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// AddGlobalModel Add global model. Only for China mainland sellers and Korean sellers.
//
// Path: /api/v2/global_product/add_global_model
// https://open.shopee.com/documents/v2/v2.global_product.add_global_model?module=90&type=1
func (s *GlobalProductServiceOp[T]) AddGlobalModel(ctx context.Context, sid uint64, mid uint64, tok string, req AddGlobalModelRequest) (*AddGlobalModelResponse, error) {
	path := "/global_product/add_global_model"
	resp := new(AddGlobalModelResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// CategoryRecommend Recommend category by item name. Only for China mainland sellers and Korean sellers.
// Path: /api/v2/global_product/category_recommend
// https://open.shopee.com/documents/v2/v2.global_product.category_recommend?module=90&type=1
func (s *GlobalProductServiceOp[T]) CategoryRecommend(ctx context.Context, sid uint64, mid uint64, tok string, opt CategoryRecommendRequest) (*CategoryRecommendResponse, error) {
	path := "/global_product/category_recommend"
	resp := new(CategoryRecommendResponse)
	err := s.client.Get(ctx, path, resp, opt, sid, mid, tok)
	return resp, err
}

// CreatePublishTask Create publish task for global item. Only for China mainland sellers and Korean sellers.
// Path: /api/v2/global_product/create_publish_task
// https://open.shopee.com/documents/v2/v2.global_product.create_publish_task?module=90&type=1
func (s *GlobalProductServiceOp[T]) CreatePublishTask(ctx context.Context, sid uint64, mid uint64, tok string, req CreatePublishTaskRequest) (*CreatePublishTaskResponse, error) {
	path := "/global_product/create_publish_task"
	resp := new(CreatePublishTaskResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// DeleteGlobalItem Delete global item. Only for China mainland sellers and Korean sellers.
// Path: /api/v2/global_product/delete_global_item
// https://open.shopee.com/documents/v2/v2.global_product.delete_global_item?module=90&type=1
func (s *GlobalProductServiceOp[T]) DeleteGlobalItem(ctx context.Context, sid uint64, mid uint64, tok string, req DeleteGlobalItemRequest) (*DeleteGlobalItemResponse, error) {
	path := "/global_product/delete_global_item"
	resp := new(DeleteGlobalItemResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// DeleteGlobalModel Delete global model. Only for China mainland sellers and Korean sellers.
// Path: /api/v2/global_product/delete_global_model
// https://open.shopee.com/documents/v2/v2.global_product.delete_global_model?module=90&type=1
func (s *GlobalProductServiceOp[T]) DeleteGlobalModel(ctx context.Context, sid uint64, mid uint64, tok string, req DeleteGlobalModelRequest) (*DeleteGlobalModelResponse, error) {
	path := "/global_product/delete_global_model"
	resp := new(DeleteGlobalModelResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// GetAttributeTree Get the mtsku attribute trees for categories
// Path: /api/v2/global_product/get_attribute_tree
// https://open.shopee.com/documents/v2/v2.global_product.get_attribute_tree?module=90&type=1
func (s *GlobalProductServiceOp[T]) GetAttributeTree(ctx context.Context, sid uint64, mid uint64, tok string, opt GetAttributeTreeRequest) (*GetAttributeTreeResponse, error) {
	path := "/global_product/get_attribute_tree"
	resp := new(GetAttributeTreeResponse)
	err := s.client.Get(ctx, path, resp, opt, sid, mid, tok)
	return resp, err
}

// GetBrandList Use this call to get a list of brand. Only for China mainland sellers and Korean sellers.
// Path: /api/v2/global_product/get_brand_list
// https://open.shopee.com/documents/v2/v2.global_product.get_brand_list?module=90&type=1
func (s *GlobalProductServiceOp[T]) GetBrandList(ctx context.Context, sid uint64, mid uint64, tok string, opt GetBrandListRequest) (*GetBrandListResponse, error) {
	path := "/global_product/get_brand_list"
	resp := new(GetBrandListResponse)
	err := s.client.Get(ctx, path, resp, opt, sid, mid, tok)
	return resp, err
}

// GetCategory Get global category. Only for China mainland sellers and Korean sellers.
// Path: /api/v2/global_product/get_category
// https://open.shopee.com/documents/v2/v2.global_product.get_category?module=90&type=1
func (s *GlobalProductServiceOp[T]) GetCategory(ctx context.Context, sid uint64, mid uint64, tok string, opt GetCategoryRequest) (*GetCategoryResponse, error) {
	path := "/global_product/get_category"
	resp := new(GetCategoryResponse)
	err := s.client.Get(ctx, path, resp, opt, sid, mid, tok)
	return resp, err
}

// GetGlobalItemId Get get_global_item_id by item_id. Only for China mainland sellers and Korean sellers.
// Path: /api/v2/global_product/get_global_item_id
// https://open.shopee.com/documents/v2/v2.global_product.get_global_item_id?module=90&type=1
func (s *GlobalProductServiceOp[T]) GetGlobalItemId(ctx context.Context, sid uint64, mid uint64, tok string, opt GetGlobalItemIdRequest) (*GetGlobalItemIdResponse, error) {
	path := "/global_product/get_global_item_id"
	resp := new(GetGlobalItemIdResponse)
	err := s.client.Get(ctx, path, resp, opt, sid, mid, tok)
	return resp, err
}

// GetGlobalItemInfo Get global item info.Only for China mainland sellers and Korean sellers.
// Path: /api/v2/global_product/get_global_item_info
// https://open.shopee.com/documents/v2/v2.global_product.get_global_item_info?module=90&type=1
func (s *GlobalProductServiceOp[T]) GetGlobalItemInfo(ctx context.Context, sid uint64, mid uint64, tok string, opt GetGlobalItemInfoRequest) (*GetGlobalItemInfoResponse, error) {
	path := "/global_product/get_global_item_info"
	resp := new(GetGlobalItemInfoResponse)
	err := s.client.Get(ctx, path, resp, opt, sid, mid, tok)
	return resp, err
}

// GetGlobalItemLimit Get global item upload control.
// Path: /api/v2/global_product/get_global_item_limit
// https://open.shopee.com/documents/v2/v2.global_product.get_global_item_limit?module=90&type=1
func (s *GlobalProductServiceOp[T]) GetGlobalItemLimit(ctx context.Context, sid uint64, mid uint64, tok string, opt GetGlobalItemLimitRequest) (*GetGlobalItemLimitResponse, error) {
	path := "/global_product/get_global_item_limit"
	resp := new(GetGlobalItemLimitResponse)
	err := s.client.Get(ctx, path, resp, opt, sid, mid, tok)
	return resp, err
}

// GetGlobalItemList Get global item id list. Only for China mainland sellers and Korean sellers.
// Path: /api/v2/global_product/get_global_item_list
// https://open.shopee.com/documents/v2/v2.global_product.get_global_item_list?module=90&type=1
func (s *GlobalProductServiceOp[T]) GetGlobalItemList(ctx context.Context, sid uint64, mid uint64, tok string, opt GetGlobalItemListRequest) (*GetGlobalItemListResponse, error) {
	path := "/global_product/get_global_item_list"
	resp := new(GetGlobalItemListResponse)
	err := s.client.Get(ctx, path, resp, opt, sid, mid, tok)
	return resp, err
}

// GetGlobalModelList Get global model list. Only for China mainland sellers and Korean sellers.
// Path: /api/v2/global_product/get_global_model_list
// https://open.shopee.com/documents/v2/v2.global_product.get_global_model_list?module=90&type=1
func (s *GlobalProductServiceOp[T]) GetGlobalModelList(ctx context.Context, sid uint64, mid uint64, tok string, opt GetGlobalModelListRequest) (*GetGlobalModelListResponse, error) {
	path := "/global_product/get_global_model_list"
	resp := new(GetGlobalModelListResponse)
	err := s.client.Get(ctx, path, resp, opt, sid, mid, tok)
	return resp, err
}

// GetLocalAdjustmentRate Retrieves the adjustment rate that converts CB stock price into local-warehouse price for a specific shop.
// Path: /api/v2/global_product/get_local_adjustment_rate
// https://open.shopee.com/documents/v2/v2.global_product.get_local_adjustment_rate?module=90&type=1
func (s *GlobalProductServiceOp[T]) GetLocalAdjustmentRate(ctx context.Context, sid uint64, mid uint64, tok string) (*GetLocalAdjustmentRateResponse, error) {
	path := "/global_product/get_local_adjustment_rate"
	resp := new(GetLocalAdjustmentRateResponse)
	err := s.client.Post(ctx, path, nil, resp, sid, mid, tok)
	return resp, err
}

// GetPublishableShop Get publishable shop list for global item. Only for China mainland sellers and Korean sellers.
//
// Path: /api/v2/global_product/get_publishable_shop
// https://open.shopee.com/documents/v2/v2.global_product.get_publishable_shop?module=90&type=1
func (s *GlobalProductServiceOp[T]) GetPublishableShop(ctx context.Context, sid uint64, mid uint64, tok string, opt GetPublishableShopRequest) (*GetPublishableShopResponse, error) {
	path := "/global_product/get_publishable_shop"
	resp := new(GetPublishableShopResponse)
	err := s.client.Get(ctx, path, resp, opt, sid, mid, tok)
	return resp, err
}

// GetPublishedList Get published item list of global item. Only for China mainland sellers and Korean sellers.
//
// Path: /api/v2/global_product/get_published_list
// https://open.shopee.com/documents/v2/v2.global_product.get_published_list?module=90&type=1
func (s *GlobalProductServiceOp[T]) GetPublishedList(ctx context.Context, sid uint64, mid uint64, tok string, opt GetPublishedListRequest) (*GetPublishedListResponse, error) {
	path := "/global_product/get_published_list"
	resp := new(GetPublishedListResponse)
	err := s.client.Get(ctx, path, resp, opt, sid, mid, tok)
	return resp, err
}

// GetPublishTaskResult Get publish task result for global item. Only for China mainland sellers and Korean sellers.
// Path: /api/v2/global_product/get_publish_task_result
// https://open.shopee.com/documents/v2/v2.global_product.get_publish_task_result?module=90&type=1
func (s *GlobalProductServiceOp[T]) GetPublishTaskResult(ctx context.Context, sid uint64, mid uint64, tok string, opt GetPublishTaskResultRequest) (*GetPublishTaskResultResponse, error) {
	path := "/global_product/get_publish_task_result"
	resp := new(GetPublishTaskResultResponse)
	err := s.client.Get(ctx, path, resp, opt, sid, mid, tok)
	return resp, err
}

// GetRecommendAttribute Get recommend attributes. Only for China mainland sellers and Korean sellers.
// Path: /api/v2/global_product/get_recommend_attribute
// https://open.shopee.com/documents/v2/v2.global_product.get_recommend_attribute?module=90&type=1
func (s *GlobalProductServiceOp[T]) GetRecommendAttribute(ctx context.Context, sid uint64, mid uint64, tok string, opt GetRecommendAttributeRequest) (*GetRecommendAttributeResponse, error) {
	path := "/global_product/get_recommend_attribute"
	resp := new(GetRecommendAttributeResponse)
	err := s.client.Get(ctx, path, resp, opt, sid, mid, tok)
	return resp, err
}

// GetShopPublishableStatus Get publishable shop list for global item in pages.
// Path: /api/v2/global_product/get_shop_publishable_status
// https://open.shopee.com/documents/v2/v2.global_product.get_shop_publishable_status?module=90&type=1
func (s *GlobalProductServiceOp[T]) GetShopPublishableStatus(ctx context.Context, sid uint64, mid uint64, tok string, opt GetShopPublishableStatusRequest) (*GetShopPublishableStatusResponse, error) {
	path := "/global_product/get_shop_publishable_status"
	resp := new(GetShopPublishableStatusResponse)
	err := s.client.Get(ctx, path, resp, opt, sid, mid, tok)
	return resp, err
}

// GetSizeChartDetail Get new size chart detail
// Path: /api/v2/global_product/get_size_chart_detail
// https://open.shopee.com/documents/v2/v2.global_product.get_size_chart_detail?module=90&type=1
func (s *GlobalProductServiceOp[T]) GetSizeChartDetail(ctx context.Context, sid uint64, mid uint64, tok string, req GetSizeChartDetailRequest) (*GetSizeChartDetailResponse, error) {
	path := "/global_product/get_size_chart_detail"
	resp := new(GetSizeChartDetailResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// GetSizeChartList {"content":"<p>Get size chart list</p>","raw_content":[{"name":"paragraph","children":[{"data":"Get size chart list"}]}]}
// Path: /api/v2/global_product/get_size_chart_list
// https://open.shopee.com/documents/v2/v2.global_product.get_size_chart_list?module=90&type=1
func (s *GlobalProductServiceOp[T]) GetSizeChartList(ctx context.Context, sid uint64, mid uint64, tok string, req GetSizeChartListRequest) (*GetSizeChartListResponse, error) {
	path := "/global_product/get_size_chart_list"
	resp := new(GetSizeChartListResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// GetVariations Get the standardized tier variation defined by Shopee, which is currently a three-layer tree structure. The top layer is variations, the second layer is groups, groups are used to divide options, and the third layer is options.
// Path: /api/v2/global_product/get_variations
// https://open.shopee.com/documents/v2/v2.global_product.get_variations?module=90&type=1
func (s *GlobalProductServiceOp[T]) GetVariations(ctx context.Context, sid uint64, mid uint64, tok string, opt GetVariationsRequest) (*GetVariationsResponse, error) {
	path := "/global_product/get_variations"
	resp := new(GetVariationsResponse)
	err := s.client.Get(ctx, path, resp, opt, sid, mid, tok)
	return resp, err
}

// InitTierVariation Only for China mainland sellers and Korean sellers. If you only define color, it is one tier, if you define color and size, it is two tier. Support two tier structures at most. This API can change no tier to one tier, no tier to two tier, one tier to two tier, two tier to one tier, one tier to no tier, two tier to no tier. Please create variants after an interval of 5 seconds after creating an item, as there may be a delay.
// Path: /api/v2/global_product/init_tier_variation
// https://open.shopee.com/documents/v2/v2.global_product.init_tier_variation?module=90&type=1
func (s *GlobalProductServiceOp[T]) InitTierVariation(ctx context.Context, sid uint64, mid uint64, tok string, req InitTierVariationRequest) (*InitTierVariationResponse, error) {
	path := "/global_product/init_tier_variation"
	resp := new(InitTierVariationResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// SearchGlobalAttributeValueList this api is for searching attribute value list for attribute with support_search_value flag
// Path: /api/v2/global_product/search_global_attribute_value_list
// https://open.shopee.com/documents/v2/v2.global_product.search_global_attribute_value_list?module=90&type=1
func (s *GlobalProductServiceOp[T]) SearchGlobalAttributeValueList(ctx context.Context, sid uint64, mid uint64, tok string, req SearchGlobalAttributeValueListRequest) (*SearchGlobalAttributeValueListResponse, error) {
	path := "/global_product/search_global_attribute_value_list"
	resp := new(SearchGlobalAttributeValueListResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// SetSyncField Set auto sync field. Only for China mainland sellers and Korean sellers.
// Path: /api/v2/global_product/set_sync_field
// https://open.shopee.com/documents/v2/v2.global_product.set_sync_field?module=90&type=1
func (s *GlobalProductServiceOp[T]) SetSyncField(ctx context.Context, sid uint64, mid uint64, tok string, req SetSyncFieldRequest) (*SetSyncFieldResponse, error) {
	path := "/global_product/set_sync_field"
	resp := new(SetSyncFieldResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// SupportSizeChart Get category support size chart. Only for China mainland sellers and Korean sellers.
//
// Path: /api/v2/global_product/support_size_chart
// https://open.shopee.com/documents/v2/v2.global_product.support_size_chart?module=90&type=1
func (s *GlobalProductServiceOp[T]) SupportSizeChart(ctx context.Context, sid uint64, mid uint64, tok string, opt SupportSizeChartRequest) (*SupportSizeChartResponse, error) {
	path := "/global_product/support_size_chart"
	resp := new(SupportSizeChartResponse)
	err := s.client.Get(ctx, path, resp, opt, sid, mid, tok)
	return resp, err
}

// UpdateGlobalItem Update global item. Only for China mainland sellers and Korean sellers.
// Path: /api/v2/global_product/update_global_item
// https://open.shopee.com/documents/v2/v2.global_product.update_global_item?module=90&type=1
func (s *GlobalProductServiceOp[T]) UpdateGlobalItem(ctx context.Context, sid uint64, mid uint64, tok string, req UpdateGlobalItemRequest) (*UpdateGlobalItemResponse, error) {
	path := "/global_product/update_global_item"
	resp := new(UpdateGlobalItemResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// UpdateGlobalModel Update global model. Only for China mainland sellers and Korean sellers.
// Path: /api/v2/global_product/update_global_model
// https://open.shopee.com/documents/v2/v2.global_product.update_global_model?module=90&type=1
func (s *GlobalProductServiceOp[T]) UpdateGlobalModel(ctx context.Context, sid uint64, mid uint64, tok string, req UpdateGlobalModelRequest) (*UpdateGlobalModelResponse, error) {
	path := "/global_product/update_global_model"
	resp := new(UpdateGlobalModelResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// UpdateLocalAdjustmentRate A multiplier that automatically converts your CB stock price into the local-warehouse price. It ensures your local inventory prices reflect regional costs, currency factors, and margin targets.
// Path: /api/v2/global_product/update_local_adjustment_rate
// https://open.shopee.com/documents/v2/v2.global_product.update_local_adjustment_rate?module=90&type=1
func (s *GlobalProductServiceOp[T]) UpdateLocalAdjustmentRate(ctx context.Context, sid uint64, mid uint64, tok string, req UpdateLocalAdjustmentRateRequest) (*UpdateLocalAdjustmentRateResponse, error) {
	path := "/global_product/update_local_adjustment_rate"
	resp := new(UpdateLocalAdjustmentRateResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// UpdatePrice Update global price. Only for China mainland sellers and Korean sellers.
// Path: /api/v2/global_product/update_price
// https://open.shopee.com/documents/v2/v2.global_product.update_price?module=90&type=1
func (s *GlobalProductServiceOp[T]) UpdatePrice(ctx context.Context, sid uint64, mid uint64, tok string, req UpdatePriceRequest) (*UpdatePriceResponse, error) {
	path := "/global_product/update_price"
	resp := new(UpdatePriceResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// UpdateSizeChart Update size chart for global item. Only for China mainland sellers and Korean sellers.
//
// Path: /api/v2/global_product/update_size_chart
// https://open.shopee.com/documents/v2/v2.global_product.update_size_chart?module=90&type=1
func (s *GlobalProductServiceOp[T]) UpdateSizeChart(ctx context.Context, sid uint64, mid uint64, tok string, req UpdateSizeChartRequest) (*UpdateSizeChartResponse, error) {
	path := "/global_product/update_size_chart"
	resp := new(UpdateSizeChartResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// UpdateStock Update global stock. Only for China mainland sellers and Korean sellers.
// Path: /api/v2/global_product/update_stock
// https://open.shopee.com/documents/v2/v2.global_product.update_stock?module=90&type=1
func (s *GlobalProductServiceOp[T]) UpdateStock(ctx context.Context, sid uint64, mid uint64, tok string, req UpdateStockRequest) (*UpdateStockResponse, error) {
	path := "/global_product/update_stock"
	resp := new(UpdateStockResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// UpdateTierVariation Update global product tier variation. Only for China mainland sellers and Korean sellers.This api can only be used without changing the tier structure, you can add options, delete options, and update the option image by this api.
//
// Path: /api/v2/global_product/update_tier_variation
// https://open.shopee.com/documents/v2/v2.global_product.update_tier_variation?module=90&type=1
func (s *GlobalProductServiceOp[T]) UpdateTierVariation(ctx context.Context, sid uint64, mid uint64, tok string, req UpdateTierVariationRequest) (*UpdateTierVariationResponse, error) {
	path := "/global_product/update_tier_variation"
	resp := new(UpdateTierVariationResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}
