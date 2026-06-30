package goshopee

import (
	"context"
)

type ProductService interface {
	// AddItem {"content":"<p>Add a new item.</p>","raw_content":[{"name":"paragraph","children":[{"data":"Add a new item."}]}]}
	// Path: /api/v2/product/add_item
	// https://open.shopee.com/documents/v2/v2.product.add_item?module=89&type=1
	AddItem(ctx context.Context, sid uint64, req AddItemRequest, tok string) (*AddItemResponse, error)
	// AddKitItem Create the kit item by selecting multiple items and setting main component and quantity per kit.
	// Path: /api/v2/product/add_kit_item
	// https://open.shopee.com/documents/v2/v2.product.add_kit_item?module=89&type=1
	AddKitItem(ctx context.Context, sid uint64, req AddKitItemRequest, tok string) (*AddKitItemResponse, error)
	// AddModel Add model. More detail please check: https://open.shopee.com/developer-guide/219
	// Path: /api/v2/product/add_model
	// https://open.shopee.com/documents/v2/v2.product.add_model?module=89&type=1
	AddModel(ctx context.Context, sid uint64, req AddModelRequest, tok string) (*AddModelResponse, error)
	// BatchAddItem {"content":"<p>Create asynchronous task to batch add item</p>","raw_content":[{"name":"paragraph","children":[{"data":"Create asynchronous task to batch add item"}]}]}
	// Path: /api/v2/product/batch_add_item
	// https://open.shopee.com/documents/v2/v2.product.batch_add_item?module=89&type=1
	BatchAddItem(ctx context.Context, sid uint64, req BatchAddItemRequest, tok string) (*BatchAddItemResponse, error)
	// BatchPublishItemToOutletShop {"content":"<p>Create asynchronous task to batch publish outlet item</p>","raw_content":[{"name":"paragraph","children":[{"data":"Create asynchronous task to batch publish outlet item"}]}]}
	// Path: /api/v2/product/batch_publish_item_to_outlet_shop
	// https://open.shopee.com/documents/v2/v2.product.batch_publish_item_to_outlet_shop?module=89&type=1
	BatchPublishItemToOutletShop(ctx context.Context, sid uint64, req BatchPublishItemToOutletShopRequest, tok string) (*BatchPublishItemToOutletShopResponse, error)
	// BatchUpdateOutletPrice {"content":"<p>Create asynchronous task to batch update outlet item's price</p>","raw_content":[{"name":"paragraph","children":[{"data":"Create asynchronous task to batch update outlet item's price"}]}]}
	// Path: /api/v2/product/batch_update_outlet_price
	// https://open.shopee.com/documents/v2/v2.product.batch_update_outlet_price?module=89&type=1
	BatchUpdateOutletPrice(ctx context.Context, sid uint64, req BatchUpdateOutletPriceRequest, tok string) (*BatchUpdateOutletPriceResponse, error)
	// BatchUpdateOutletStock {"content":"<p>Create asynchronous task to batch update outlet stock</p>","raw_content":[{"name":"paragraph","children":[{"data":"Create asynchronous task to batch update outlet stock"}]}]}
	// Path: /api/v2/product/batch_update_outlet_stock
	// https://open.shopee.com/documents/v2/v2.product.batch_update_outlet_stock?module=89&type=1
	BatchUpdateOutletStock(ctx context.Context, sid uint64, req BatchUpdateOutletStockRequest, tok string) (*BatchUpdateOutletStockResponse, error)
	// BoostItem Boost item.
	// Path: /api/v2/product/boost_item
	// https://open.shopee.com/documents/v2/v2.product.boost_item?module=89&type=1
	BoostItem(ctx context.Context, sid uint64, req BoostItemRequest, tok string) (*BoostItemResponse, error)
	// CategoryRecommend Recommend category by item name.
	// Path: /api/v2/product/category_recommend
	// https://open.shopee.com/documents/v2/v2.product.category_recommend?module=89&type=1
	CategoryRecommend(ctx context.Context, sid uint64, opt ProductCategoryRecommendRequest, tok string) (*ProductCategoryRecommendResponse, error)
	// DeleteItem Use this call to delete a product item.
	// Path: /api/v2/product/delete_item
	// https://open.shopee.com/documents/v2/v2.product.delete_item?module=89&type=1
	DeleteItem(ctx context.Context, sid uint64, req DeleteItemRequest, tok string) (*DeleteItemResponse, error)
	// DeleteModel Delete item model.
	// Path: /api/v2/product/delete_model
	// https://open.shopee.com/documents/v2/v2.product.delete_model?module=89&type=1
	DeleteModel(ctx context.Context, sid uint64, req DeleteModelRequest, tok string) (*DeleteModelResponse, error)
	// GenerateKitImage This API generates a single consolidated image by combining the cover images of all selected items. It is typically used to create a unified product display image for kits or bundles.
	// Path: /api/v2/product/generate_kit_image
	// https://open.shopee.com/documents/v2/v2.product.generate_kit_image?module=89&type=1
	GenerateKitImage(ctx context.Context, sid uint64, req GenerateKitImageRequest, tok string) (*GenerateKitImageResponse, error)
	// GetAitemByPitemId Get the list of A Items under SIP Affiliate Shop corresponding to P Items under SIP Primary Shop.
	// Path: /api/v2/product/get_aitem_by_pitem_id
	// https://open.shopee.com/documents/v2/v2.product.get_aitem_by_pitem_id?module=89&type=1
	GetAitemByPitemId(ctx context.Context, sid uint64, opt GetAitemByPitemIdRequest, tok string) (*GetAitemByPitemIdResponse, error)
	// GetAllVehicleList Use this Open API to get all vehicle list.
	// Path: /api/v2/product/get_all_vehicle_list
	// https://open.shopee.com/documents/v2/v2.product.get_all_vehicle_list?module=89&type=1
	GetAllVehicleList(ctx context.Context, sid uint64, opt GetAllVehicleListRequest, tok string) (*GetAllVehicleListResponse, error)
	// GetAttributeTree Get the attribute tree for categories
	// Path: /api/v2/product/get_attribute_tree
	// https://open.shopee.com/documents/v2/v2.product.get_attribute_tree?module=89&type=1
	GetAttributeTree(ctx context.Context, sid uint64, opt ProductGetAttributeTreeRequest, tok string) (*ProductGetAttributeTreeResponse, error)
	// GetBatchTaskResult {"content":"<p>Query batch task result</p>","raw_content":[{"name":"paragraph","children":[{"data":"Query batch task result"}]}]}
	// Path: /api/v2/product/get_batch_task_result
	// https://open.shopee.com/documents/v2/v2.product.get_batch_task_result?module=89&type=1
	GetBatchTaskResult(ctx context.Context, sid uint64, req GetBatchTaskResultRequest, tok string) (*GetBatchTaskResultResponse, error)
	// GetBoostedList Get boosted item list.
	// Path: /api/v2/product/get_boosted_list
	// https://open.shopee.com/documents/v2/v2.product.get_boosted_list?module=89&type=1
	GetBoostedList(ctx context.Context, sid uint64, tok string) (*GetBoostedListResponse, error)
	// GetBrandList Get the brand data of a leaf category. More detail please check: https://open.shopee.com/developer-guide/209
	// Path: /api/v2/product/get_brand_list
	// https://open.shopee.com/documents/v2/v2.product.get_brand_list?module=89&type=1
	GetBrandList(ctx context.Context, sid uint64, opt ProductGetBrandListRequest, tok string) (*ProductGetBrandListResponse, error)
	// GetCategory Get category tree data. More detail please check https://open.shopee.com/developer-guide/209
	// Path: /api/v2/product/get_category
	// https://open.shopee.com/documents/v2/v2.product.get_category?module=89&type=1
	GetCategory(ctx context.Context, sid uint64, opt ProductGetCategoryRequest, tok string) (*ProductGetCategoryResponse, error)
	// GetComment Use this api to get comment by shop_id, item_id, or comment_id, get up to 1000 comments.
	// Path: /api/v2/product/get_comment
	// https://open.shopee.com/documents/v2/v2.product.get_comment?module=89&type=1
	GetComment(ctx context.Context, sid uint64, opt GetCommentRequest, tok string) (*GetCommentResponse, error)
	// GetDirectItemList get direct item by main item.
	// Path: /api/v2/product/get_direct_item_list
	// https://open.shopee.com/documents/v2/v2.product.get_direct_item_list?module=89&type=1
	GetDirectItemList(ctx context.Context, sid uint64, opt GetDirectItemListRequest, tok string) (*GetDirectItemListResponse, error)
	// GetDirectShopRecommendedPrice get recommend price for direct shop.
	// Path: /api/v2/product/get_direct_shop_recommended_price
	// https://open.shopee.com/documents/v2/v2.product.get_direct_shop_recommended_price?module=89&type=1
	GetDirectShopRecommendedPrice(ctx context.Context, sid uint64, opt GetDirectShopRecommendedPriceRequest, tok string) (*GetDirectShopRecommendedPriceResponse, error)
	// GetItemBaseInfo Use this api to get basic info of item by item_id list.
	// Path: /api/v2/product/get_item_base_info
	// https://open.shopee.com/documents/v2/v2.product.get_item_base_info?module=89&type=1
	GetItemBaseInfo(ctx context.Context, sid uint64, opt GetItemBaseInfoRequest, tok string) (*GetItemBaseInfoResponse, error)
	// GetItemContentDiagnosisResult Get the content quality details (including content quality level, content issues, and system suggestions) for specific product list.
	// Path: /api/v2/product/get_item_content_diagnosis_result
	// https://open.shopee.com/documents/v2/v2.product.get_item_content_diagnosis_result?module=89&type=1
	GetItemContentDiagnosisResult(ctx context.Context, sid uint64, req GetItemContentDiagnosisResultRequest, tok string) (*GetItemContentDiagnosisResultResponse, error)
	// GetItemExtraInfo Use this api to get extra info of item by item_id list.
	// Path: /api/v2/product/get_item_extra_info
	// https://open.shopee.com/documents/v2/v2.product.get_item_extra_info?module=89&type=1
	GetItemExtraInfo(ctx context.Context, sid uint64, opt GetItemExtraInfoRequest, tok string) (*GetItemExtraInfoResponse, error)
	// GetItemLimit Get item upload control.
	// Path: /api/v2/product/get_item_limit
	// https://open.shopee.com/documents/v2/v2.product.get_item_limit?module=89&type=1
	GetItemLimit(ctx context.Context, sid uint64, opt GetItemLimitRequest, tok string) (*GetItemLimitResponse, error)
	// GetItemList Use this call to get a list of items.
	// Path: /api/v2/product/get_item_list
	// https://open.shopee.com/documents/v2/v2.product.get_item_list?module=89&type=1
	GetItemList(ctx context.Context, sid uint64, opt ProductGetItemListRequest, tok string) (*ProductGetItemListResponse, error)
	// GetItemListByContentDiagnosis Query the list of products and their content quality details by content quality level or content issues.
	// Path: /api/v2/product/get_item_list_by_content_diagnosis
	// https://open.shopee.com/documents/v2/v2.product.get_item_list_by_content_diagnosis?module=89&type=1
	GetItemListByContentDiagnosis(ctx context.Context, sid uint64, req GetItemListByContentDiagnosisRequest, tok string) (*GetItemListByContentDiagnosisResponse, error)
	// GetItemPromotion Get item promotion info.
	// Path: /api/v2/product/get_item_promotion
	// https://open.shopee.com/documents/v2/v2.product.get_item_promotion?module=89&type=1
	GetItemPromotion(ctx context.Context, sid uint64, opt GetItemPromotionRequest, tok string) (*GetItemPromotionResponse, error)
	// GetItemViolationInfo get item violation info
	// Path: /api/v2/product/get_item_violation_info
	// https://open.shopee.com/documents/v2/v2.product.get_item_violation_info?module=89&type=1
	GetItemViolationInfo(ctx context.Context, sid uint64, req GetItemViolationInfoRequest, tok string) (*GetItemViolationInfoResponse, error)
	// GetKitItemInfo Get the kit basic information and kit components.
	// Path: /api/v2/product/get_kit_item_info
	// https://open.shopee.com/documents/v2/v2.product.get_kit_item_info?module=89&type=1
	GetKitItemInfo(ctx context.Context, sid uint64, opt GetKitItemInfoRequest, tok string) (*GetKitItemInfoResponse, error)
	// GetKitItemLimit Get the limit of Kit item.
	// Path: /api/v2/product/get_kit_item_limit
	// https://open.shopee.com/documents/v2/v2.product.get_kit_item_limit?module=89&type=1
	GetKitItemLimit(ctx context.Context, sid uint64, opt GetKitItemLimitRequest, tok string) (*GetKitItemLimitResponse, error)
	// GetMainItemList get main item by direct item.
	// Path: /api/v2/product/get_main_item_list
	// https://open.shopee.com/documents/v2/v2.product.get_main_item_list?module=89&type=1
	GetMainItemList(ctx context.Context, sid uint64, opt GetMainItemListRequest, tok string) (*GetMainItemListResponse, error)
	// GetMartItemByOutletItemId {"content":"<p>Get the mapping information between a Mart item and its corresponding outlet item by outlet item ID.</p><p><br>&nbsp;</p>","raw_content":[{"name":"paragraph","children":[{"data":"Get the mapping information between a Mart item and its corresponding outlet item by outlet item ID."}]},{"name":"paragraph","children":[{"name":"softBreak"},{"data":" "}]}]}
	// Path: /api/v2/product/get_mart_item_by_outlet_item_id
	// https://open.shopee.com/documents/v2/v2.product.get_mart_item_by_outlet_item_id?module=89&type=1
	GetMartItemByOutletItemId(ctx context.Context, sid uint64, req GetMartItemByOutletItemIdRequest, tok string) (*GetMartItemByOutletItemIdResponse, error)
	// GetMartItemMappingById Get the mapping information between a Mart item and its corresponding outlet item by item ID.
	// Path: /api/v2/product/get_mart_item_mapping_by_id
	// https://open.shopee.com/documents/v2/v2.product.get_mart_item_mapping_by_id?module=89&type=1
	GetMartItemMappingById(ctx context.Context, sid uint64, req GetMartItemMappingByIdRequest, tok string) (*GetMartItemMappingByIdResponse, error)
	// GetModelList Get model list of an item.
	// Path: /api/v2/product/get_model_list
	// https://open.shopee.com/documents/v2/v2.product.get_model_list?module=89&type=1
	GetModelList(ctx context.Context, sid uint64, opt GetModelListRequest, tok string) (*GetModelListResponse, error)
	// GetProductCertificationRule Get product certification rule
	// Path: /api/v2/product/get_product_certification_rule
	// https://open.shopee.com/documents/v2/v2.product.get_product_certification_rule?module=89&type=1
	GetProductCertificationRule(ctx context.Context, sid uint64, req GetProductCertificationRuleRequest, tok string) (*GetProductCertificationRuleResponse, error)
	// GetRecommendAttribute Get recommend attributes.
	// Path: /api/v2/product/get_recommend_attribute
	// https://open.shopee.com/documents/v2/v2.product.get_recommend_attribute?module=89&type=1
	GetRecommendAttribute(ctx context.Context, sid uint64, opt ProductGetRecommendAttributeRequest, tok string) (*ProductGetRecommendAttributeResponse, error)
	// GetSizeChartDetail Get new size chart detail. Now only local shop support to use this api to get new size chart detail.
	// Path: /api/v2/product/get_size_chart_detail
	// https://open.shopee.com/documents/v2/v2.product.get_size_chart_detail?module=89&type=1
	GetSizeChartDetail(ctx context.Context, sid uint64, opt ProductGetSizeChartDetailRequest, tok string) (*ProductGetSizeChartDetailResponse, error)
	// GetSizeChartList Get new size chat list. Now only support local shop to use new size chart.
	// Path: /api/v2/product/get_size_chart_list
	// https://open.shopee.com/documents/v2/v2.product.get_size_chart_list?module=89&type=1
	GetSizeChartList(ctx context.Context, sid uint64, opt ProductGetSizeChartListRequest, tok string) (*ProductGetSizeChartListResponse, error)
	// GetVariations Get the standardized tier variation defined by Shopee, which is currently a three-layer tree structure.
	// The top layer is variations, the second layer is groups, groups are used to divide options, and the third layer is options.
	// Path: /api/v2/product/get_variation_tree
	// https://open.shopee.com/documents/v2/v2.product.get_variations?module=89&type=1
	GetVariations(ctx context.Context, sid uint64, opt ProductGetVariationsRequest, tok string) (*ProductGetVariationsResponse, error)
	// GetVehicleListByCompatibilityDetail Use this Open API to get vehicle list by brand, model, year, and version.
	// Path: /api/v2/product/get_vehicle_list_by_compatibility_detail
	// https://open.shopee.com/documents/v2/v2.product.get_vehicle_list_by_compatibility_detail?module=89&type=1
	GetVehicleListByCompatibilityDetail(ctx context.Context, sid uint64, opt GetVehicleListByCompatibilityDetailRequest, tok string) (*GetVehicleListByCompatibilityDetailResponse, error)
	// GetWeightRecommendation Get recommended weight. Now only BR shop support to use this api to get recommended weight.
	// Path: /api/v2/product/get_weight_recommendation
	// https://open.shopee.com/documents/v2/v2.product.get_weight_recommendation?module=89&type=1
	GetWeightRecommendation(ctx context.Context, sid uint64, req GetWeightRecommendationRequest, tok string) (*GetWeightRecommendationResponse, error)
	// InitTierVariation This API allows you to update the tier structure of a product. Defining only color creates one tier, while color + size creates two tiers (maximum supported). Supported changes include: no tier ↔ one/two tiers, one tier ↔ two/no tier, and two tiers ↔ one/no tier. For details, see Developer Guide.  Please wait at least 5 seconds after creating an item before creating variants, as processing may be delayed.
	// Path: /api/v2/product/init_tier_variation
	// https://open.shopee.com/documents/v2/v2.product.init_tier_variation?module=89&type=1
	InitTierVariation(ctx context.Context, sid uint64, req ProductInitTierVariationRequest, tok string) (*ProductInitTierVariationResponse, error)
	// PublishItemToOutletShop
	// Path: /api/v2/
	// https://open.shopee.com/documents/v2/v2.product.publish_item_to_outlet_shop?module=89&type=1
	PublishItemToOutletShop(ctx context.Context, sid uint64, tok string) (*PublishItemToOutletShopResponse, error)
	// RegisterBrand Use this call to register a brand.
	// Path: /api/v2/product/register_brand
	// https://open.shopee.com/documents/v2/v2.product.register_brand?module=89&type=1
	RegisterBrand(ctx context.Context, sid uint64, req RegisterBrandRequest, tok string) (*RegisterBrandResponse, error)
	// ReplyComment Use this api to reply comments from buyers in batch.
	// Path: /api/v2/product/reply_comment
	// https://open.shopee.com/documents/v2/v2.product.reply_comment?module=89&type=1
	ReplyComment(ctx context.Context, sid uint64, req ReplyCommentRequest, tok string) (*ReplyCommentResponse, error)
	// SearchAttributeValueList this api is for searching attribute value list for attribute with support_search_value flag
	// Path: /api/v2/product/search_attribute_value_list
	// https://open.shopee.com/documents/v2/v2.product.search_attribute_value_list?module=89&type=1
	SearchAttributeValueList(ctx context.Context, sid uint64, req SearchAttributeValueListRequest, tok string) (*SearchAttributeValueListResponse, error)
	// SearchItem Use this call to search item.
	// Path: /api/v2/product/search_item
	// https://open.shopee.com/documents/v2/v2.product.search_item?module=89&type=1
	SearchItem(ctx context.Context, sid uint64, opt SearchItemRequest, tok string) (*SearchItemResponse, error)
	// SearchUnpackagedModelList Use this API to retrieve Unpackaged SKU ID information for items that toggle on logistics channel 30029.
	// Path: /api/v2/product/search_unpackaged_model_list
	// https://open.shopee.com/documents/v2/v2.product.search_unpackaged_model_list?module=89&type=1
	SearchUnpackagedModelList(ctx context.Context, sid uint64, req SearchUnpackagedModelListRequest, tok string) (*SearchUnpackagedModelListResponse, error)
	// UnlistItem Unlist item.
	// Path: /api/v2/product/unlist_item
	// https://open.shopee.com/documents/v2/v2.product.unlist_item?module=89&type=1
	UnlistItem(ctx context.Context, sid uint64, req UnlistItemRequest, tok string) (*UnlistItemResponse, error)
	// UpdateItem {"content":"<p>Update item.</p>","raw_content":[{"name":"paragraph","children":[{"data":"Update item."}]}]}
	// Path: /api/v2/product/update_item
	// https://open.shopee.com/documents/v2/v2.product.update_item?module=89&type=1
	UpdateItem(ctx context.Context, sid uint64, req UpdateItemRequest, tok string) (*UpdateItemResponse, error)
	// UpdateKitItem Update the kit basic information and kit components, only support adding kit variations and updating existing kit variation’s image, price, and model_sku, don’t support deleting existing kit variations and updating the items, main component and quantity per kit of existing kit variations.
	// Path: /api/v2/product/update_kit_item
	// https://open.shopee.com/documents/v2/v2.product.update_kit_item?module=89&type=1
	UpdateKitItem(ctx context.Context, sid uint64, req UpdateKitItemRequest, tok string) (*UpdateKitItemResponse, error)
	// UpdateModel Update seller sku/ pre order/ model status for model.
	// Path: /api/v2/product/update_model
	// https://open.shopee.com/documents/v2/v2.product.update_model?module=89&type=1
	UpdateModel(ctx context.Context, sid uint64, req UpdateModelRequest, tok string) (*UpdateModelResponse, error)
	// UpdatePrice Update price.
	// Path: /api/v2/product/update_price
	// https://open.shopee.com/documents/v2/v2.product.update_price?module=89&type=1
	UpdatePrice(ctx context.Context, sid uint64, req ProductUpdatePriceRequest, tok string) (*ProductUpdatePriceResponse, error)
	// UpdateSipItemPrice Update sip item price.
	// Path: /api/v2/product/update_sip_item_price
	// https://open.shopee.com/documents/v2/v2.product.update_sip_item_price?module=89&type=1
	UpdateSipItemPrice(ctx context.Context, sid uint64, req UpdateSipItemPriceRequest, tok string) (*UpdateSipItemPriceResponse, error)
	// UpdateStock Use this API to update one item_id for each call, but still can support updating multiple model_ids stock of the same item_id (If you need batch modification, please call multiple times)This API will update only "seller_stock".Whenever there is a promotion ongoing or upcoming, the total stock must be larger than or equal to real-time “reserved_stock” promotion stock (Please check v2.get_item_promotion API for more details). Items that are deleted will not be allowed to modify stock.
	// Path: /api/v2/product/update_stock
	// https://open.shopee.com/documents/v2/v2.product.update_stock?module=89&type=1
	UpdateStock(ctx context.Context, sid uint64, req ProductUpdateStockRequest, tok string) (*ProductUpdateStockResponse, error)
	// UpdateTierVariation This api can only be used without changing the tier structure, you can add options, delete options, and update the option image by this api. More detail please check: https://open.shopee.com/developer-guide/219
	// Path: /api/v2/product/update_tier_variation
	// https://open.shopee.com/documents/v2/v2.product.update_tier_variation?module=89&type=1
	UpdateTierVariation(ctx context.Context, sid uint64, req ProductUpdateTierVariationRequest, tok string) (*ProductUpdateTierVariationResponse, error)
}

type ProductServiceOp[T any] struct {
	client *Client[T]
}

// AddItem {"content":"<p>Add a new item.</p>","raw_content":[{"name":"paragraph","children":[{"data":"Add a new item."}]}]}
// Path: /api/v2/product/add_item
// https://open.shopee.com/documents/v2/v2.product.add_item?module=89&type=1
func (s *ProductServiceOp[T]) AddItem(ctx context.Context, sid uint64, req AddItemRequest, tok string) (*AddItemResponse, error) {
	path := "/product/add_item"
	resp := new(AddItemResponse)
	err := s.client.WithShop(sid, tok).Post(ctx, path, req, resp)
	return resp, err
}

// AddKitItem Create the kit item by selecting multiple items and setting main component and quantity per kit.
// Path: /api/v2/product/add_kit_item
// https://open.shopee.com/documents/v2/v2.product.add_kit_item?module=89&type=1
func (s *ProductServiceOp[T]) AddKitItem(ctx context.Context, sid uint64, req AddKitItemRequest, tok string) (*AddKitItemResponse, error) {
	path := "/product/add_kit_item"
	resp := new(AddKitItemResponse)
	err := s.client.WithShop(sid, tok).Post(ctx, path, req, resp)
	return resp, err
}

// AddModel Add model. More detail please check: https://open.shopee.com/developer-guide/219
// Path: /api/v2/product/add_model
// https://open.shopee.com/documents/v2/v2.product.add_model?module=89&type=1
func (s *ProductServiceOp[T]) AddModel(ctx context.Context, sid uint64, req AddModelRequest, tok string) (*AddModelResponse, error) {
	path := "/product/add_model"
	resp := new(AddModelResponse)
	err := s.client.WithShop(sid, tok).Post(ctx, path, req, resp)
	return resp, err
}

// BatchAddItem {"content":"<p>Create asynchronous task to batch add item</p>","raw_content":[{"name":"paragraph","children":[{"data":"Create asynchronous task to batch add item"}]}]}
// Path: /api/v2/product/batch_add_item
// https://open.shopee.com/documents/v2/v2.product.batch_add_item?module=89&type=1
func (s *ProductServiceOp[T]) BatchAddItem(ctx context.Context, sid uint64, req BatchAddItemRequest, tok string) (*BatchAddItemResponse, error) {
	path := "/product/batch_add_item"
	resp := new(BatchAddItemResponse)
	err := s.client.WithShop(sid, tok).Post(ctx, path, req, resp)
	return resp, err
}

// BatchPublishItemToOutletShop {"content":"<p>Create asynchronous task to batch publish outlet item</p>","raw_content":[{"name":"paragraph","children":[{"data":"Create asynchronous task to batch publish outlet item"}]}]}
// Path: /api/v2/product/batch_publish_item_to_outlet_shop
// https://open.shopee.com/documents/v2/v2.product.batch_publish_item_to_outlet_shop?module=89&type=1
func (s *ProductServiceOp[T]) BatchPublishItemToOutletShop(ctx context.Context, sid uint64, req BatchPublishItemToOutletShopRequest, tok string) (*BatchPublishItemToOutletShopResponse, error) {
	path := "/product/batch_publish_item_to_outlet_shop"
	resp := new(BatchPublishItemToOutletShopResponse)
	err := s.client.WithShop(sid, tok).Post(ctx, path, req, resp)
	return resp, err
}

// BatchUpdateOutletPrice {"content":"<p>Create asynchronous task to batch update outlet item's price</p>","raw_content":[{"name":"paragraph","children":[{"data":"Create asynchronous task to batch update outlet item's price"}]}]}
// Path: /api/v2/product/batch_update_outlet_price
// https://open.shopee.com/documents/v2/v2.product.batch_update_outlet_price?module=89&type=1
func (s *ProductServiceOp[T]) BatchUpdateOutletPrice(ctx context.Context, sid uint64, req BatchUpdateOutletPriceRequest, tok string) (*BatchUpdateOutletPriceResponse, error) {
	path := "/product/batch_update_outlet_price"
	resp := new(BatchUpdateOutletPriceResponse)
	err := s.client.WithShop(sid, tok).Post(ctx, path, req, resp)
	return resp, err
}

// BatchUpdateOutletStock {"content":"<p>Create asynchronous task to batch update outlet stock</p>","raw_content":[{"name":"paragraph","children":[{"data":"Create asynchronous task to batch update outlet stock"}]}]}
// Path: /api/v2/product/batch_update_outlet_stock
// https://open.shopee.com/documents/v2/v2.product.batch_update_outlet_stock?module=89&type=1
func (s *ProductServiceOp[T]) BatchUpdateOutletStock(ctx context.Context, sid uint64, req BatchUpdateOutletStockRequest, tok string) (*BatchUpdateOutletStockResponse, error) {
	path := "/product/batch_update_outlet_stock"
	resp := new(BatchUpdateOutletStockResponse)
	err := s.client.WithShop(sid, tok).Post(ctx, path, req, resp)
	return resp, err
}

// BoostItem Boost item.
// Path: /api/v2/product/boost_item
// https://open.shopee.com/documents/v2/v2.product.boost_item?module=89&type=1
func (s *ProductServiceOp[T]) BoostItem(ctx context.Context, sid uint64, req BoostItemRequest, tok string) (*BoostItemResponse, error) {
	path := "/product/boost_item"
	resp := new(BoostItemResponse)
	err := s.client.WithShop(sid, tok).Post(ctx, path, req, resp)
	return resp, err
}

// CategoryRecommend Recommend category by item name.
// Path: /api/v2/product/category_recommend
// https://open.shopee.com/documents/v2/v2.product.category_recommend?module=89&type=1
func (s *ProductServiceOp[T]) CategoryRecommend(ctx context.Context, sid uint64, opt ProductCategoryRecommendRequest, tok string) (*ProductCategoryRecommendResponse, error) {
	path := "/product/category_recommend"
	resp := new(ProductCategoryRecommendResponse)
	err := s.client.WithShop(sid, tok).Get(ctx, path, resp, opt)
	return resp, err
}

// DeleteItem Use this call to delete a product item.
// Path: /api/v2/product/delete_item
// https://open.shopee.com/documents/v2/v2.product.delete_item?module=89&type=1
func (s *ProductServiceOp[T]) DeleteItem(ctx context.Context, sid uint64, req DeleteItemRequest, tok string) (*DeleteItemResponse, error) {
	path := "/product/delete_item"
	resp := new(DeleteItemResponse)
	err := s.client.WithShop(sid, tok).Post(ctx, path, req, resp)
	return resp, err
}

// DeleteModel Delete item model.
// Path: /api/v2/product/delete_model
// https://open.shopee.com/documents/v2/v2.product.delete_model?module=89&type=1
func (s *ProductServiceOp[T]) DeleteModel(ctx context.Context, sid uint64, req DeleteModelRequest, tok string) (*DeleteModelResponse, error) {
	path := "/product/delete_model"
	resp := new(DeleteModelResponse)
	err := s.client.WithShop(sid, tok).Post(ctx, path, req, resp)
	return resp, err
}

// GenerateKitImage This API generates a single consolidated image by combining the cover images of all selected items. It is typically used to create a unified product display image for kits or bundles.
// Path: /api/v2/product/generate_kit_image
// https://open.shopee.com/documents/v2/v2.product.generate_kit_image?module=89&type=1
func (s *ProductServiceOp[T]) GenerateKitImage(ctx context.Context, sid uint64, req GenerateKitImageRequest, tok string) (*GenerateKitImageResponse, error) {
	path := "/product/generate_kit_image"
	resp := new(GenerateKitImageResponse)
	err := s.client.WithShop(sid, tok).Post(ctx, path, req, resp)
	return resp, err
}

// GetAitemByPitemId Get the list of A Items under SIP Affiliate Shop corresponding to P Items under SIP Primary Shop.
// Path: /api/v2/product/get_aitem_by_pitem_id
// https://open.shopee.com/documents/v2/v2.product.get_aitem_by_pitem_id?module=89&type=1
func (s *ProductServiceOp[T]) GetAitemByPitemId(ctx context.Context, sid uint64, opt GetAitemByPitemIdRequest, tok string) (*GetAitemByPitemIdResponse, error) {
	path := "/product/get_aitem_by_pitem_id"
	resp := new(GetAitemByPitemIdResponse)
	err := s.client.WithShop(sid, tok).Get(ctx, path, resp, opt)
	return resp, err
}

// GetAllVehicleList Use this Open API to get all vehicle list.
// Path: /api/v2/product/get_all_vehicle_list
// https://open.shopee.com/documents/v2/v2.product.get_all_vehicle_list?module=89&type=1
func (s *ProductServiceOp[T]) GetAllVehicleList(ctx context.Context, sid uint64, opt GetAllVehicleListRequest, tok string) (*GetAllVehicleListResponse, error) {
	path := "/product/get_all_vehicle_list"
	resp := new(GetAllVehicleListResponse)
	err := s.client.WithShop(sid, tok).Get(ctx, path, resp, opt)
	return resp, err
}

// GetAttributeTree Get the attribute tree for categories
// Path: /api/v2/product/get_attribute_tree
// https://open.shopee.com/documents/v2/v2.product.get_attribute_tree?module=89&type=1
func (s *ProductServiceOp[T]) GetAttributeTree(ctx context.Context, sid uint64, opt ProductGetAttributeTreeRequest, tok string) (*ProductGetAttributeTreeResponse, error) {
	path := "/product/get_attribute_tree"
	resp := new(ProductGetAttributeTreeResponse)
	err := s.client.WithShop(sid, tok).Get(ctx, path, resp, opt)
	return resp, err
}

// GetBatchTaskResult {"content":"<p>Query batch task result</p>","raw_content":[{"name":"paragraph","children":[{"data":"Query batch task result"}]}]}
// Path: /api/v2/product/get_batch_task_result
// https://open.shopee.com/documents/v2/v2.product.get_batch_task_result?module=89&type=1
func (s *ProductServiceOp[T]) GetBatchTaskResult(ctx context.Context, sid uint64, req GetBatchTaskResultRequest, tok string) (*GetBatchTaskResultResponse, error) {
	path := "/product/get_batch_task_result"
	resp := new(GetBatchTaskResultResponse)
	err := s.client.WithShop(sid, tok).Post(ctx, path, req, resp)
	return resp, err
}

// GetBoostedList Get boosted item list.
// Path: /api/v2/product/get_boosted_list
// https://open.shopee.com/documents/v2/v2.product.get_boosted_list?module=89&type=1
func (s *ProductServiceOp[T]) GetBoostedList(ctx context.Context, sid uint64, tok string) (*GetBoostedListResponse, error) {
	path := "/product/get_boosted_list"
	resp := new(GetBoostedListResponse)
	err := s.client.WithShop(sid, tok).Get(ctx, path, resp, nil)
	return resp, err
}

// GetBrandList Get the brand data of a leaf category. More detail please check: https://open.shopee.com/developer-guide/209
// Path: /api/v2/product/get_brand_list
// https://open.shopee.com/documents/v2/v2.product.get_brand_list?module=89&type=1
func (s *ProductServiceOp[T]) GetBrandList(ctx context.Context, sid uint64, opt ProductGetBrandListRequest, tok string) (*ProductGetBrandListResponse, error) {
	path := "/product/get_brand_list"
	resp := new(ProductGetBrandListResponse)
	err := s.client.WithShop(sid, tok).Get(ctx, path, resp, opt)
	return resp, err
}

// GetCategory Get category tree data. More detail please check https://open.shopee.com/developer-guide/209
// Path: /api/v2/product/get_category
// https://open.shopee.com/documents/v2/v2.product.get_category?module=89&type=1
func (s *ProductServiceOp[T]) GetCategory(ctx context.Context, sid uint64, opt ProductGetCategoryRequest, tok string) (*ProductGetCategoryResponse, error) {
	path := "/product/get_category"
	resp := new(ProductGetCategoryResponse)
	err := s.client.WithShop(sid, tok).Get(ctx, path, resp, opt)
	return resp, err
}

// GetComment Use this api to get comment by shop_id, item_id, or comment_id, get up to 1000 comments.
// Path: /api/v2/product/get_comment
// https://open.shopee.com/documents/v2/v2.product.get_comment?module=89&type=1
func (s *ProductServiceOp[T]) GetComment(ctx context.Context, sid uint64, opt GetCommentRequest, tok string) (*GetCommentResponse, error) {
	path := "/product/get_comment"
	resp := new(GetCommentResponse)
	err := s.client.WithShop(sid, tok).Get(ctx, path, resp, opt)
	return resp, err
}

// GetDirectItemList get direct item by main item.
// Path: /api/v2/product/get_direct_item_list
// https://open.shopee.com/documents/v2/v2.product.get_direct_item_list?module=89&type=1
func (s *ProductServiceOp[T]) GetDirectItemList(ctx context.Context, sid uint64, opt GetDirectItemListRequest, tok string) (*GetDirectItemListResponse, error) {
	path := "/product/get_direct_item_list"
	resp := new(GetDirectItemListResponse)
	err := s.client.WithShop(sid, tok).Get(ctx, path, resp, opt)
	return resp, err
}

// GetDirectShopRecommendedPrice get recommend price for direct shop.
// Path: /api/v2/product/get_direct_shop_recommended_price
// https://open.shopee.com/documents/v2/v2.product.get_direct_shop_recommended_price?module=89&type=1
func (s *ProductServiceOp[T]) GetDirectShopRecommendedPrice(ctx context.Context, sid uint64, opt GetDirectShopRecommendedPriceRequest, tok string) (*GetDirectShopRecommendedPriceResponse, error) {
	path := "/product/get_direct_shop_recommended_price"
	resp := new(GetDirectShopRecommendedPriceResponse)
	err := s.client.WithShop(sid, tok).Get(ctx, path, resp, opt)
	return resp, err
}

// GetItemBaseInfo Use this api to get basic info of item by item_id list.
// Path: /api/v2/product/get_item_base_info
// https://open.shopee.com/documents/v2/v2.product.get_item_base_info?module=89&type=1
func (s *ProductServiceOp[T]) GetItemBaseInfo(ctx context.Context, sid uint64, opt GetItemBaseInfoRequest, tok string) (*GetItemBaseInfoResponse, error) {
	path := "/product/get_item_base_info"
	resp := new(GetItemBaseInfoResponse)
	err := s.client.WithShop(sid, tok).Get(ctx, path, resp, opt)
	return resp, err
}

// GetItemContentDiagnosisResult Get the content quality details (including content quality level, content issues, and system suggestions) for specific product list.
// Path: /api/v2/product/get_item_content_diagnosis_result
// https://open.shopee.com/documents/v2/v2.product.get_item_content_diagnosis_result?module=89&type=1
func (s *ProductServiceOp[T]) GetItemContentDiagnosisResult(ctx context.Context, sid uint64, req GetItemContentDiagnosisResultRequest, tok string) (*GetItemContentDiagnosisResultResponse, error) {
	path := "/product/get_item_content_diagnosis_result"
	resp := new(GetItemContentDiagnosisResultResponse)
	err := s.client.WithShop(sid, tok).Post(ctx, path, req, resp)
	return resp, err
}

// GetItemExtraInfo Use this api to get extra info of item by item_id list.
// Path: /api/v2/product/get_item_extra_info
// https://open.shopee.com/documents/v2/v2.product.get_item_extra_info?module=89&type=1
func (s *ProductServiceOp[T]) GetItemExtraInfo(ctx context.Context, sid uint64, opt GetItemExtraInfoRequest, tok string) (*GetItemExtraInfoResponse, error) {
	path := "/product/get_item_extra_info"
	resp := new(GetItemExtraInfoResponse)
	err := s.client.WithShop(sid, tok).Get(ctx, path, resp, opt)
	return resp, err
}

// GetItemLimit Get item upload control.
// Path: /api/v2/product/get_item_limit
// https://open.shopee.com/documents/v2/v2.product.get_item_limit?module=89&type=1
func (s *ProductServiceOp[T]) GetItemLimit(ctx context.Context, sid uint64, opt GetItemLimitRequest, tok string) (*GetItemLimitResponse, error) {
	path := "/product/get_item_limit"
	resp := new(GetItemLimitResponse)
	err := s.client.WithShop(sid, tok).Get(ctx, path, resp, opt)
	return resp, err
}

// GetItemList Use this call to get a list of items.
// Path: /api/v2/product/get_item_list
// https://open.shopee.com/documents/v2/v2.product.get_item_list?module=89&type=1
func (s *ProductServiceOp[T]) GetItemList(ctx context.Context, sid uint64, opt ProductGetItemListRequest, tok string) (*ProductGetItemListResponse, error) {
	path := "/product/get_item_list"
	resp := new(ProductGetItemListResponse)
	err := s.client.WithShop(sid, tok).Get(ctx, path, resp, opt)
	return resp, err
}

// GetItemListByContentDiagnosis Query the list of products and their content quality details by content quality level or content issues.
// Path: /api/v2/product/get_item_list_by_content_diagnosis
// https://open.shopee.com/documents/v2/v2.product.get_item_list_by_content_diagnosis?module=89&type=1
func (s *ProductServiceOp[T]) GetItemListByContentDiagnosis(ctx context.Context, sid uint64, req GetItemListByContentDiagnosisRequest, tok string) (*GetItemListByContentDiagnosisResponse, error) {
	path := "/product/get_item_list_by_content_diagnosis"
	resp := new(GetItemListByContentDiagnosisResponse)
	err := s.client.WithShop(sid, tok).Post(ctx, path, req, resp)
	return resp, err
}

// GetItemPromotion Get item promotion info.
// Path: /api/v2/product/get_item_promotion
// https://open.shopee.com/documents/v2/v2.product.get_item_promotion?module=89&type=1
func (s *ProductServiceOp[T]) GetItemPromotion(ctx context.Context, sid uint64, opt GetItemPromotionRequest, tok string) (*GetItemPromotionResponse, error) {
	path := "/product/get_item_promotion"
	resp := new(GetItemPromotionResponse)
	err := s.client.WithShop(sid, tok).Get(ctx, path, resp, opt)
	return resp, err
}

// GetItemViolationInfo get item violation info
// Path: /api/v2/product/get_item_violation_info
// https://open.shopee.com/documents/v2/v2.product.get_item_violation_info?module=89&type=1
func (s *ProductServiceOp[T]) GetItemViolationInfo(ctx context.Context, sid uint64, req GetItemViolationInfoRequest, tok string) (*GetItemViolationInfoResponse, error) {
	path := "/product/get_item_violation_info"
	resp := new(GetItemViolationInfoResponse)
	err := s.client.WithShop(sid, tok).Post(ctx, path, req, resp)
	return resp, err
}

// GetKitItemInfo Get the kit basic information and kit components.
// Path: /api/v2/product/get_kit_item_info
// https://open.shopee.com/documents/v2/v2.product.get_kit_item_info?module=89&type=1
func (s *ProductServiceOp[T]) GetKitItemInfo(ctx context.Context, sid uint64, opt GetKitItemInfoRequest, tok string) (*GetKitItemInfoResponse, error) {
	path := "/product/get_kit_item_info"
	resp := new(GetKitItemInfoResponse)
	err := s.client.WithShop(sid, tok).Get(ctx, path, resp, opt)
	return resp, err
}

// GetKitItemLimit Get the limit of Kit item.
// Path: /api/v2/product/get_kit_item_limit
// https://open.shopee.com/documents/v2/v2.product.get_kit_item_limit?module=89&type=1
func (s *ProductServiceOp[T]) GetKitItemLimit(ctx context.Context, sid uint64, opt GetKitItemLimitRequest, tok string) (*GetKitItemLimitResponse, error) {
	path := "/product/get_kit_item_limit"
	resp := new(GetKitItemLimitResponse)
	err := s.client.WithShop(sid, tok).Get(ctx, path, resp, opt)
	return resp, err
}

// GetMainItemList get main item by direct item.
// Path: /api/v2/product/get_main_item_list
// https://open.shopee.com/documents/v2/v2.product.get_main_item_list?module=89&type=1
func (s *ProductServiceOp[T]) GetMainItemList(ctx context.Context, sid uint64, opt GetMainItemListRequest, tok string) (*GetMainItemListResponse, error) {
	path := "/product/get_main_item_list"
	resp := new(GetMainItemListResponse)
	err := s.client.WithShop(sid, tok).Get(ctx, path, resp, opt)
	return resp, err
}

// GetMartItemByOutletItemId {"content":"<p>Get the mapping information between a Mart item and its corresponding outlet item by outlet item ID.</p><p><br>&nbsp;</p>","raw_content":[{"name":"paragraph","children":[{"data":"Get the mapping information between a Mart item and its corresponding outlet item by outlet item ID."}]},{"name":"paragraph","children":[{"name":"softBreak"},{"data":" "}]}]}
// Path: /api/v2/product/get_mart_item_by_outlet_item_id
// https://open.shopee.com/documents/v2/v2.product.get_mart_item_by_outlet_item_id?module=89&type=1
func (s *ProductServiceOp[T]) GetMartItemByOutletItemId(ctx context.Context, sid uint64, req GetMartItemByOutletItemIdRequest, tok string) (*GetMartItemByOutletItemIdResponse, error) {
	path := "/product/get_mart_item_by_outlet_item_id"
	resp := new(GetMartItemByOutletItemIdResponse)
	err := s.client.WithShop(sid, tok).Post(ctx, path, req, resp)
	return resp, err
}

// GetMartItemMappingById Get the mapping information between a Mart item and its corresponding outlet item by item ID.
// Path: /api/v2/product/get_mart_item_mapping_by_id
// https://open.shopee.com/documents/v2/v2.product.get_mart_item_mapping_by_id?module=89&type=1
func (s *ProductServiceOp[T]) GetMartItemMappingById(ctx context.Context, sid uint64, req GetMartItemMappingByIdRequest, tok string) (*GetMartItemMappingByIdResponse, error) {
	path := "/product/get_mart_item_mapping_by_id"
	resp := new(GetMartItemMappingByIdResponse)
	err := s.client.WithShop(sid, tok).Post(ctx, path, req, resp)
	return resp, err
}

// GetModelList Get model list of an item.
// Path: /api/v2/product/get_model_list
// https://open.shopee.com/documents/v2/v2.product.get_model_list?module=89&type=1
func (s *ProductServiceOp[T]) GetModelList(ctx context.Context, sid uint64, opt GetModelListRequest, tok string) (*GetModelListResponse, error) {
	path := "/product/get_model_list"
	resp := new(GetModelListResponse)
	err := s.client.WithShop(sid, tok).Get(ctx, path, resp, opt)
	return resp, err
}

// GetProductCertificationRule Get product certification rule
// Path: /api/v2/product/get_product_certification_rule
// https://open.shopee.com/documents/v2/v2.product.get_product_certification_rule?module=89&type=1
func (s *ProductServiceOp[T]) GetProductCertificationRule(ctx context.Context, sid uint64, req GetProductCertificationRuleRequest, tok string) (*GetProductCertificationRuleResponse, error) {
	path := "/product/get_product_certification_rule"
	resp := new(GetProductCertificationRuleResponse)
	err := s.client.WithShop(sid, tok).Post(ctx, path, req, resp)
	return resp, err
}

// GetRecommendAttribute Get recommend attributes.
// Path: /api/v2/product/get_recommend_attribute
// https://open.shopee.com/documents/v2/v2.product.get_recommend_attribute?module=89&type=1
func (s *ProductServiceOp[T]) GetRecommendAttribute(ctx context.Context, sid uint64, opt ProductGetRecommendAttributeRequest, tok string) (*ProductGetRecommendAttributeResponse, error) {
	path := "/product/get_recommend_attribute"
	resp := new(ProductGetRecommendAttributeResponse)
	err := s.client.WithShop(sid, tok).Get(ctx, path, resp, opt)
	return resp, err
}

// GetSizeChartDetail Get new size chart detail. Now only local shop support to use this api to get new size chart detail.
// Path: /api/v2/product/get_size_chart_detail
// https://open.shopee.com/documents/v2/v2.product.get_size_chart_detail?module=89&type=1
func (s *ProductServiceOp[T]) GetSizeChartDetail(ctx context.Context, sid uint64, opt ProductGetSizeChartDetailRequest, tok string) (*ProductGetSizeChartDetailResponse, error) {
	path := "/product/get_size_chart_detail"
	resp := new(ProductGetSizeChartDetailResponse)
	err := s.client.WithShop(sid, tok).Get(ctx, path, resp, opt)
	return resp, err
}

// GetSizeChartList Get new size chat list. Now only support local shop to use new size chart.
// Path: /api/v2/product/get_size_chart_list
// https://open.shopee.com/documents/v2/v2.product.get_size_chart_list?module=89&type=1
func (s *ProductServiceOp[T]) GetSizeChartList(ctx context.Context, sid uint64, opt ProductGetSizeChartListRequest, tok string) (*ProductGetSizeChartListResponse, error) {
	path := "/product/get_size_chart_list"
	resp := new(ProductGetSizeChartListResponse)
	err := s.client.WithShop(sid, tok).Get(ctx, path, resp, opt)
	return resp, err
}

// GetVariations Get the standardized tier variation defined by Shopee, which is currently a three-layer tree structure.
// The top layer is variations, the second layer is groups, groups are used to divide options, and the third layer is options.
// Path: /api/v2/product/get_variation_tree
// https://open.shopee.com/documents/v2/v2.product.get_variations?module=89&type=1
func (s *ProductServiceOp[T]) GetVariations(ctx context.Context, sid uint64, opt ProductGetVariationsRequest, tok string) (*ProductGetVariationsResponse, error) {
	path := "/product/get_variation_tree"
	resp := new(ProductGetVariationsResponse)
	err := s.client.WithShop(sid, tok).Get(ctx, path, resp, opt)
	return resp, err
}

// GetVehicleListByCompatibilityDetail Use this Open API to get vehicle list by brand, model, year, and version.
// Path: /api/v2/product/get_vehicle_list_by_compatibility_detail
// https://open.shopee.com/documents/v2/v2.product.get_vehicle_list_by_compatibility_detail?module=89&type=1
func (s *ProductServiceOp[T]) GetVehicleListByCompatibilityDetail(ctx context.Context, sid uint64, opt GetVehicleListByCompatibilityDetailRequest, tok string) (*GetVehicleListByCompatibilityDetailResponse, error) {
	path := "/product/get_vehicle_list_by_compatibility_detail"
	resp := new(GetVehicleListByCompatibilityDetailResponse)
	err := s.client.WithShop(sid, tok).Get(ctx, path, resp, opt)
	return resp, err
}

// GetWeightRecommendation Get recommended weight. Now only BR shop support to use this api to get recommended weight.
// Path: /api/v2/product/get_weight_recommendation
// https://open.shopee.com/documents/v2/v2.product.get_weight_recommendation?module=89&type=1
func (s *ProductServiceOp[T]) GetWeightRecommendation(ctx context.Context, sid uint64, req GetWeightRecommendationRequest, tok string) (*GetWeightRecommendationResponse, error) {
	path := "/product/get_weight_recommendation"
	resp := new(GetWeightRecommendationResponse)
	err := s.client.WithShop(sid, tok).Post(ctx, path, req, resp)
	return resp, err
}

// InitTierVariation This API allows you to update the tier structure of a product. Defining only color creates one tier, while color + size creates two tiers (maximum supported). Supported changes include: no tier ↔ one/two tiers, one tier ↔ two/no tier, and two tiers ↔ one/no tier. For details, see Developer Guide.  Please wait at least 5 seconds after creating an item before creating variants, as processing may be delayed.
// Path: /api/v2/product/init_tier_variation
// https://open.shopee.com/documents/v2/v2.product.init_tier_variation?module=89&type=1
func (s *ProductServiceOp[T]) InitTierVariation(ctx context.Context, sid uint64, req ProductInitTierVariationRequest, tok string) (*ProductInitTierVariationResponse, error) {
	path := "/product/init_tier_variation"
	resp := new(ProductInitTierVariationResponse)
	err := s.client.WithShop(sid, tok).Post(ctx, path, req, resp)
	return resp, err
}

// PublishItemToOutletShop
// Path: /api/v2/
// https://open.shopee.com/documents/v2/v2.product.publish_item_to_outlet_shop?module=89&type=1
func (s *ProductServiceOp[T]) PublishItemToOutletShop(ctx context.Context, sid uint64, tok string) (*PublishItemToOutletShopResponse, error) {
	path := "/"
	resp := new(PublishItemToOutletShopResponse)
	err := s.client.WithShop(sid, tok).Post(ctx, path, nil, resp)
	return resp, err
}

// RegisterBrand Use this call to register a brand.
// Path: /api/v2/product/register_brand
// https://open.shopee.com/documents/v2/v2.product.register_brand?module=89&type=1
func (s *ProductServiceOp[T]) RegisterBrand(ctx context.Context, sid uint64, req RegisterBrandRequest, tok string) (*RegisterBrandResponse, error) {
	path := "/product/register_brand"
	resp := new(RegisterBrandResponse)
	err := s.client.WithShop(sid, tok).Post(ctx, path, req, resp)
	return resp, err
}

// ReplyComment Use this api to reply comments from buyers in batch.
// Path: /api/v2/product/reply_comment
// https://open.shopee.com/documents/v2/v2.product.reply_comment?module=89&type=1
func (s *ProductServiceOp[T]) ReplyComment(ctx context.Context, sid uint64, req ReplyCommentRequest, tok string) (*ReplyCommentResponse, error) {
	path := "/product/reply_comment"
	resp := new(ReplyCommentResponse)
	err := s.client.WithShop(sid, tok).Post(ctx, path, req, resp)
	return resp, err
}

// SearchAttributeValueList this api is for searching attribute value list for attribute with support_search_value flag
// Path: /api/v2/product/search_attribute_value_list
// https://open.shopee.com/documents/v2/v2.product.search_attribute_value_list?module=89&type=1
func (s *ProductServiceOp[T]) SearchAttributeValueList(ctx context.Context, sid uint64, req SearchAttributeValueListRequest, tok string) (*SearchAttributeValueListResponse, error) {
	path := "/product/search_attribute_value_list"
	resp := new(SearchAttributeValueListResponse)
	err := s.client.WithShop(sid, tok).Post(ctx, path, req, resp)
	return resp, err
}

// SearchItem Use this call to search item.
// Path: /api/v2/product/search_item
// https://open.shopee.com/documents/v2/v2.product.search_item?module=89&type=1
func (s *ProductServiceOp[T]) SearchItem(ctx context.Context, sid uint64, opt SearchItemRequest, tok string) (*SearchItemResponse, error) {
	path := "/product/search_item"
	resp := new(SearchItemResponse)
	err := s.client.WithShop(sid, tok).Get(ctx, path, resp, opt)
	return resp, err
}

// SearchUnpackagedModelList Use this API to retrieve Unpackaged SKU ID information for items that toggle on logistics channel 30029.
// Path: /api/v2/product/search_unpackaged_model_list
// https://open.shopee.com/documents/v2/v2.product.search_unpackaged_model_list?module=89&type=1
func (s *ProductServiceOp[T]) SearchUnpackagedModelList(ctx context.Context, sid uint64, req SearchUnpackagedModelListRequest, tok string) (*SearchUnpackagedModelListResponse, error) {
	path := "/product/search_unpackaged_model_list"
	resp := new(SearchUnpackagedModelListResponse)
	err := s.client.WithShop(sid, tok).Post(ctx, path, req, resp)
	return resp, err
}

// UnlistItem Unlist item.
// Path: /api/v2/product/unlist_item
// https://open.shopee.com/documents/v2/v2.product.unlist_item?module=89&type=1
func (s *ProductServiceOp[T]) UnlistItem(ctx context.Context, sid uint64, req UnlistItemRequest, tok string) (*UnlistItemResponse, error) {
	path := "/product/unlist_item"
	resp := new(UnlistItemResponse)
	err := s.client.WithShop(sid, tok).Post(ctx, path, req, resp)
	return resp, err
}

// UpdateItem {"content":"<p>Update item.</p>","raw_content":[{"name":"paragraph","children":[{"data":"Update item."}]}]}
// Path: /api/v2/product/update_item
// https://open.shopee.com/documents/v2/v2.product.update_item?module=89&type=1
func (s *ProductServiceOp[T]) UpdateItem(ctx context.Context, sid uint64, req UpdateItemRequest, tok string) (*UpdateItemResponse, error) {
	path := "/product/update_item"
	resp := new(UpdateItemResponse)
	err := s.client.WithShop(sid, tok).Post(ctx, path, req, resp)
	return resp, err
}

// UpdateKitItem Update the kit basic information and kit components, only support adding kit variations and updating existing kit variation’s image, price, and model_sku, don’t support deleting existing kit variations and updating the items, main component and quantity per kit of existing kit variations.
// Path: /api/v2/product/update_kit_item
// https://open.shopee.com/documents/v2/v2.product.update_kit_item?module=89&type=1
func (s *ProductServiceOp[T]) UpdateKitItem(ctx context.Context, sid uint64, req UpdateKitItemRequest, tok string) (*UpdateKitItemResponse, error) {
	path := "/product/update_kit_item"
	resp := new(UpdateKitItemResponse)
	err := s.client.WithShop(sid, tok).Post(ctx, path, req, resp)
	return resp, err
}

// UpdateModel Update seller sku/ pre order/ model status for model.
// Path: /api/v2/product/update_model
// https://open.shopee.com/documents/v2/v2.product.update_model?module=89&type=1
func (s *ProductServiceOp[T]) UpdateModel(ctx context.Context, sid uint64, req UpdateModelRequest, tok string) (*UpdateModelResponse, error) {
	path := "/product/update_model"
	resp := new(UpdateModelResponse)
	err := s.client.WithShop(sid, tok).Post(ctx, path, req, resp)
	return resp, err
}

// UpdatePrice Update price.
// Path: /api/v2/product/update_price
// https://open.shopee.com/documents/v2/v2.product.update_price?module=89&type=1
func (s *ProductServiceOp[T]) UpdatePrice(ctx context.Context, sid uint64, req ProductUpdatePriceRequest, tok string) (*ProductUpdatePriceResponse, error) {
	path := "/product/update_price"
	resp := new(ProductUpdatePriceResponse)
	err := s.client.WithShop(sid, tok).Post(ctx, path, req, resp)
	return resp, err
}

// UpdateSipItemPrice Update sip item price.
// Path: /api/v2/product/update_sip_item_price
// https://open.shopee.com/documents/v2/v2.product.update_sip_item_price?module=89&type=1
func (s *ProductServiceOp[T]) UpdateSipItemPrice(ctx context.Context, sid uint64, req UpdateSipItemPriceRequest, tok string) (*UpdateSipItemPriceResponse, error) {
	path := "/product/update_sip_item_price"
	resp := new(UpdateSipItemPriceResponse)
	err := s.client.WithShop(sid, tok).Post(ctx, path, req, resp)
	return resp, err
}

// UpdateStock Use this API to update one item_id for each call, but still can support updating multiple model_ids stock of the same item_id (If you need batch modification, please call multiple times)This API will update only "seller_stock".Whenever there is a promotion ongoing or upcoming, the total stock must be larger than or equal to real-time “reserved_stock” promotion stock (Please check v2.get_item_promotion API for more details). Items that are deleted will not be allowed to modify stock.
// Path: /api/v2/product/update_stock
// https://open.shopee.com/documents/v2/v2.product.update_stock?module=89&type=1
func (s *ProductServiceOp[T]) UpdateStock(ctx context.Context, sid uint64, req ProductUpdateStockRequest, tok string) (*ProductUpdateStockResponse, error) {
	path := "/product/update_stock"
	resp := new(ProductUpdateStockResponse)
	err := s.client.WithShop(sid, tok).Post(ctx, path, req, resp)
	return resp, err
}

// UpdateTierVariation This api can only be used without changing the tier structure, you can add options, delete options, and update the option image by this api. More detail please check: https://open.shopee.com/developer-guide/219
// Path: /api/v2/product/update_tier_variation
// https://open.shopee.com/documents/v2/v2.product.update_tier_variation?module=89&type=1
func (s *ProductServiceOp[T]) UpdateTierVariation(ctx context.Context, sid uint64, req ProductUpdateTierVariationRequest, tok string) (*ProductUpdateTierVariationResponse, error) {
	path := "/product/update_tier_variation"
	resp := new(ProductUpdateTierVariationResponse)
	err := s.client.WithShop(sid, tok).Post(ctx, path, req, resp)
	return resp, err
}
