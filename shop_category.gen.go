package goshopee

import (
	"context"
)

type ShopCategoryService interface {
	// AddItemList Use this call to add items list to certain shop_category
	// Path: /api/v2/shop_category/add_item_list
	// https://open.shopee.com/documents/v2/v2.shop_category.add_item_list?module=101&type=1
	AddItemList(ctx context.Context, sid uint64, mid uint64, tok string, req ShopCategoryAddItemListRequest) (*ShopCategoryAddItemListResponse, error)
	// AddShopCategory Use this call to add a new shop collecion
	// Path: /api/v2/shop_category/add_shop_category
	// https://open.shopee.com/documents/v2/v2.shop_category.add_shop_category?module=101&type=1
	AddShopCategory(ctx context.Context, sid uint64, mid uint64, tok string, req AddShopCategoryRequest) (*AddShopCategoryResponse, error)
	// DeleteItemList Use this api to delete items from shop category
	// Path: /api/v2/shop_category/delete_item_list
	// https://open.shopee.com/documents/v2/v2.shop_category.delete_item_list?module=101&type=1
	DeleteItemList(ctx context.Context, sid uint64, mid uint64, tok string, req ShopCategoryDeleteItemListRequest) (*ShopCategoryDeleteItemListResponse, error)
	// DeleteShopCategory Use this call to delete a existing shop collecion
	// Path: /api/v2/shop_category/delete_shop_category
	// https://open.shopee.com/documents/v2/v2.shop_category.delete_shop_category?module=101&type=1
	DeleteShopCategory(ctx context.Context, sid uint64, mid uint64, tok string, req DeleteShopCategoryRequest) (*DeleteShopCategoryResponse, error)
	// GetItemList Use this call to get items list of certain shop_category
	// Path: /api/v2/shop_category/get_item_list
	// https://open.shopee.com/documents/v2/v2.shop_category.get_item_list?module=101&type=1
	GetItemList(ctx context.Context, sid uint64, mid uint64, tok string, req ShopCategoryGetItemListRequest) (*ShopCategoryGetItemListResponse, error)
	// GetShopCategoryList Use this call to get list of shop categories
	// Path: /api/v2/shop_category/get_shop_category_list
	// https://open.shopee.com/documents/v2/v2.shop_category.get_shop_category_list?module=101&type=1
	GetShopCategoryList(ctx context.Context, sid uint64, mid uint64, tok string, req GetShopCategoryListRequest) (*GetShopCategoryListResponse, error)
	// UpdateShopCategory Use this call to update a existing collecion
	// Path: /api/v2/shop_category/update_shop_category
	// https://open.shopee.com/documents/v2/v2.shop_category.update_shop_category?module=101&type=1
	UpdateShopCategory(ctx context.Context, sid uint64, mid uint64, tok string, req UpdateShopCategoryRequest) (*UpdateShopCategoryResponse, error)
}

type ShopCategoryServiceOp[T any] struct {
	client *Client[T]
}

// AddItemList Use this call to add items list to certain shop_category
// Path: /api/v2/shop_category/add_item_list
// https://open.shopee.com/documents/v2/v2.shop_category.add_item_list?module=101&type=1
func (s *ShopCategoryServiceOp[T]) AddItemList(ctx context.Context, sid uint64, mid uint64, tok string, req ShopCategoryAddItemListRequest) (*ShopCategoryAddItemListResponse, error) {
	path := "/shop_category/add_item_list"
	resp := new(ShopCategoryAddItemListResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// AddShopCategory Use this call to add a new shop collecion
// Path: /api/v2/shop_category/add_shop_category
// https://open.shopee.com/documents/v2/v2.shop_category.add_shop_category?module=101&type=1
func (s *ShopCategoryServiceOp[T]) AddShopCategory(ctx context.Context, sid uint64, mid uint64, tok string, req AddShopCategoryRequest) (*AddShopCategoryResponse, error) {
	path := "/shop_category/add_shop_category"
	resp := new(AddShopCategoryResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// DeleteItemList Use this api to delete items from shop category
// Path: /api/v2/shop_category/delete_item_list
// https://open.shopee.com/documents/v2/v2.shop_category.delete_item_list?module=101&type=1
func (s *ShopCategoryServiceOp[T]) DeleteItemList(ctx context.Context, sid uint64, mid uint64, tok string, req ShopCategoryDeleteItemListRequest) (*ShopCategoryDeleteItemListResponse, error) {
	path := "/shop_category/delete_item_list"
	resp := new(ShopCategoryDeleteItemListResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// DeleteShopCategory Use this call to delete a existing shop collecion
// Path: /api/v2/shop_category/delete_shop_category
// https://open.shopee.com/documents/v2/v2.shop_category.delete_shop_category?module=101&type=1
func (s *ShopCategoryServiceOp[T]) DeleteShopCategory(ctx context.Context, sid uint64, mid uint64, tok string, req DeleteShopCategoryRequest) (*DeleteShopCategoryResponse, error) {
	path := "/shop_category/delete_shop_category"
	resp := new(DeleteShopCategoryResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// GetItemList Use this call to get items list of certain shop_category
// Path: /api/v2/shop_category/get_item_list
// https://open.shopee.com/documents/v2/v2.shop_category.get_item_list?module=101&type=1
func (s *ShopCategoryServiceOp[T]) GetItemList(ctx context.Context, sid uint64, mid uint64, tok string, req ShopCategoryGetItemListRequest) (*ShopCategoryGetItemListResponse, error) {
	path := "/shop_category/get_item_list"
	resp := new(ShopCategoryGetItemListResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// GetShopCategoryList Use this call to get list of shop categories
// Path: /api/v2/shop_category/get_shop_category_list
// https://open.shopee.com/documents/v2/v2.shop_category.get_shop_category_list?module=101&type=1
func (s *ShopCategoryServiceOp[T]) GetShopCategoryList(ctx context.Context, sid uint64, mid uint64, tok string, req GetShopCategoryListRequest) (*GetShopCategoryListResponse, error) {
	path := "/shop_category/get_shop_category_list"
	resp := new(GetShopCategoryListResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// UpdateShopCategory Use this call to update a existing collecion
// Path: /api/v2/shop_category/update_shop_category
// https://open.shopee.com/documents/v2/v2.shop_category.update_shop_category?module=101&type=1
func (s *ShopCategoryServiceOp[T]) UpdateShopCategory(ctx context.Context, sid uint64, mid uint64, tok string, req UpdateShopCategoryRequest) (*UpdateShopCategoryResponse, error) {
	path := "/shop_category/update_shop_category"
	resp := new(UpdateShopCategoryResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}
