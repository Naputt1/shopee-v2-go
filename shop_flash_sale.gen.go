package goshopee

import (
	"context"
)

type ShopFlashSaleService interface {
	// AddShopFlashSaleItems add shop flash sale item
	// Path: /api/v2/shop_flash_sale/add_shop_flash_sale_items
	// https://open.shopee.com/documents/v2/v2.shop_flash_sale.add_shop_flash_sale_items?module=123&type=1
	AddShopFlashSaleItems(ctx context.Context, sid uint64, mid uint64, tok string, req AddShopFlashSaleItemsRequest) (*AddShopFlashSaleItemsResponse, error)
	// CreateShopFlashSale creat shop flash sale
	// Path: /api/v2/shop_flash_sale/create_shop_flash_sale
	// https://open.shopee.com/documents/v2/v2.shop_flash_sale.create_shop_flash_sale?module=123&type=1
	CreateShopFlashSale(ctx context.Context, sid uint64, mid uint64, tok string, req CreateShopFlashSaleRequest) (*CreateShopFlashSaleResponse, error)
	// DeleteShopFlashSale delete shop flash sale
	// Path: /api/v2/shop_flash_sale/delete_shop_flash_sale
	// https://open.shopee.com/documents/v2/v2.shop_flash_sale.delete_shop_flash_sale?module=123&type=1
	DeleteShopFlashSale(ctx context.Context, sid uint64, mid uint64, tok string, req DeleteShopFlashSaleRequest) (*DeleteShopFlashSaleResponse, error)
	// DeleteShopFlashSaleItems delete shop flash sale items
	// Path: /api/v2/shop_flash_sale/delete_shop_flash_sale_items
	// https://open.shopee.com/documents/v2/v2.shop_flash_sale.delete_shop_flash_sale_items?module=123&type=1
	DeleteShopFlashSaleItems(ctx context.Context, sid uint64, mid uint64, tok string, req DeleteShopFlashSaleItemsRequest) (*DeleteShopFlashSaleItemsResponse, error)
	// GetItemCriteria get shop flash sale item criteria
	// Path: /api/v2/shop_flash_sale/get_item_criteria
	// https://open.shopee.com/documents/v2/v2.shop_flash_sale.get_item_criteria?module=123&type=1
	GetItemCriteria(ctx context.Context, sid uint64, mid uint64, tok string) (*GetItemCriteriaResponse, error)
	// GetShopFlashSale get shop flash sale detail
	// Path: /api/v2/shop_flash_sale/get_shop_flash_sale
	// https://open.shopee.com/documents/v2/v2.shop_flash_sale.get_shop_flash_sale?module=123&type=1
	GetShopFlashSale(ctx context.Context, sid uint64, mid uint64, tok string, req GetShopFlashSaleRequest) (*GetShopFlashSaleResponse, error)
	// GetShopFlashSaleItems get shop flash sale items and item detail
	// Path: /api/v2/shop_flash_sale/get_shop_flash_sale_items
	// https://open.shopee.com/documents/v2/v2.shop_flash_sale.get_shop_flash_sale_items?module=123&type=1
	GetShopFlashSaleItems(ctx context.Context, sid uint64, mid uint64, tok string, req GetShopFlashSaleItemsRequest) (*GetShopFlashSaleItemsResponse, error)
	// GetShopFlashSaleList get shop flash sale list
	// Path: /api/v2/shop_flash_sale/get_shop_flash_sale_list
	// https://open.shopee.com/documents/v2/v2.shop_flash_sale.get_shop_flash_sale_list?module=123&type=1
	GetShopFlashSaleList(ctx context.Context, sid uint64, mid uint64, tok string, req GetShopFlashSaleListRequest) (*GetShopFlashSaleListResponse, error)
	// GetTimeSlotId get time slot id
	// Path: /api/v2/shop_flash_sale/get_time_slot_id
	// https://open.shopee.com/documents/v2/v2.shop_flash_sale.get_time_slot_id?module=123&type=1
	GetTimeSlotId(ctx context.Context, sid uint64, mid uint64, tok string, req GetTimeSlotIdRequest) (*GetTimeSlotIdResponse, error)
	// UpdateShopFlashSale edit shop flash sale(enable, disable)
	// Path: /api/v2/shop_flash_sale/update_shop_flash_sale
	// https://open.shopee.com/documents/v2/v2.shop_flash_sale.update_shop_flash_sale?module=123&type=1
	UpdateShopFlashSale(ctx context.Context, sid uint64, mid uint64, tok string, req UpdateShopFlashSaleRequest) (*UpdateShopFlashSaleResponse, error)
	// UpdateShopFlashSaleItems edit shop flash sale item, you can only edit the models in disbaled or enabled status
	// Path: /api/v2/shop_flash_sale/update_shop_flash_sale_items
	// https://open.shopee.com/documents/v2/v2.shop_flash_sale.update_shop_flash_sale_items?module=123&type=1
	UpdateShopFlashSaleItems(ctx context.Context, sid uint64, mid uint64, tok string, req UpdateShopFlashSaleItemsRequest) (*UpdateShopFlashSaleItemsResponse, error)
}

type ShopFlashSaleServiceOp[T any] struct {
	client *Client[T]
}

// AddShopFlashSaleItems add shop flash sale item
// Path: /api/v2/shop_flash_sale/add_shop_flash_sale_items
// https://open.shopee.com/documents/v2/v2.shop_flash_sale.add_shop_flash_sale_items?module=123&type=1
func (s *ShopFlashSaleServiceOp[T]) AddShopFlashSaleItems(ctx context.Context, sid uint64, mid uint64, tok string, req AddShopFlashSaleItemsRequest) (*AddShopFlashSaleItemsResponse, error) {
	path := "/shop_flash_sale/add_shop_flash_sale_items"
	resp := new(AddShopFlashSaleItemsResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// CreateShopFlashSale creat shop flash sale
// Path: /api/v2/shop_flash_sale/create_shop_flash_sale
// https://open.shopee.com/documents/v2/v2.shop_flash_sale.create_shop_flash_sale?module=123&type=1
func (s *ShopFlashSaleServiceOp[T]) CreateShopFlashSale(ctx context.Context, sid uint64, mid uint64, tok string, req CreateShopFlashSaleRequest) (*CreateShopFlashSaleResponse, error) {
	path := "/shop_flash_sale/create_shop_flash_sale"
	resp := new(CreateShopFlashSaleResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// DeleteShopFlashSale delete shop flash sale
// Path: /api/v2/shop_flash_sale/delete_shop_flash_sale
// https://open.shopee.com/documents/v2/v2.shop_flash_sale.delete_shop_flash_sale?module=123&type=1
func (s *ShopFlashSaleServiceOp[T]) DeleteShopFlashSale(ctx context.Context, sid uint64, mid uint64, tok string, req DeleteShopFlashSaleRequest) (*DeleteShopFlashSaleResponse, error) {
	path := "/shop_flash_sale/delete_shop_flash_sale"
	resp := new(DeleteShopFlashSaleResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// DeleteShopFlashSaleItems delete shop flash sale items
// Path: /api/v2/shop_flash_sale/delete_shop_flash_sale_items
// https://open.shopee.com/documents/v2/v2.shop_flash_sale.delete_shop_flash_sale_items?module=123&type=1
func (s *ShopFlashSaleServiceOp[T]) DeleteShopFlashSaleItems(ctx context.Context, sid uint64, mid uint64, tok string, req DeleteShopFlashSaleItemsRequest) (*DeleteShopFlashSaleItemsResponse, error) {
	path := "/shop_flash_sale/delete_shop_flash_sale_items"
	resp := new(DeleteShopFlashSaleItemsResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// GetItemCriteria get shop flash sale item criteria
// Path: /api/v2/shop_flash_sale/get_item_criteria
// https://open.shopee.com/documents/v2/v2.shop_flash_sale.get_item_criteria?module=123&type=1
func (s *ShopFlashSaleServiceOp[T]) GetItemCriteria(ctx context.Context, sid uint64, mid uint64, tok string) (*GetItemCriteriaResponse, error) {
	path := "/shop_flash_sale/get_item_criteria"
	resp := new(GetItemCriteriaResponse)
	err := s.client.Post(ctx, path, nil, resp, sid, mid, tok)
	return resp, err
}

// GetShopFlashSale get shop flash sale detail
// Path: /api/v2/shop_flash_sale/get_shop_flash_sale
// https://open.shopee.com/documents/v2/v2.shop_flash_sale.get_shop_flash_sale?module=123&type=1
func (s *ShopFlashSaleServiceOp[T]) GetShopFlashSale(ctx context.Context, sid uint64, mid uint64, tok string, req GetShopFlashSaleRequest) (*GetShopFlashSaleResponse, error) {
	path := "/shop_flash_sale/get_shop_flash_sale"
	resp := new(GetShopFlashSaleResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// GetShopFlashSaleItems get shop flash sale items and item detail
// Path: /api/v2/shop_flash_sale/get_shop_flash_sale_items
// https://open.shopee.com/documents/v2/v2.shop_flash_sale.get_shop_flash_sale_items?module=123&type=1
func (s *ShopFlashSaleServiceOp[T]) GetShopFlashSaleItems(ctx context.Context, sid uint64, mid uint64, tok string, req GetShopFlashSaleItemsRequest) (*GetShopFlashSaleItemsResponse, error) {
	path := "/shop_flash_sale/get_shop_flash_sale_items"
	resp := new(GetShopFlashSaleItemsResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// GetShopFlashSaleList get shop flash sale list
// Path: /api/v2/shop_flash_sale/get_shop_flash_sale_list
// https://open.shopee.com/documents/v2/v2.shop_flash_sale.get_shop_flash_sale_list?module=123&type=1
func (s *ShopFlashSaleServiceOp[T]) GetShopFlashSaleList(ctx context.Context, sid uint64, mid uint64, tok string, req GetShopFlashSaleListRequest) (*GetShopFlashSaleListResponse, error) {
	path := "/shop_flash_sale/get_shop_flash_sale_list"
	resp := new(GetShopFlashSaleListResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// GetTimeSlotId get time slot id
// Path: /api/v2/shop_flash_sale/get_time_slot_id
// https://open.shopee.com/documents/v2/v2.shop_flash_sale.get_time_slot_id?module=123&type=1
func (s *ShopFlashSaleServiceOp[T]) GetTimeSlotId(ctx context.Context, sid uint64, mid uint64, tok string, req GetTimeSlotIdRequest) (*GetTimeSlotIdResponse, error) {
	path := "/shop_flash_sale/get_time_slot_id"
	resp := new(GetTimeSlotIdResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// UpdateShopFlashSale edit shop flash sale(enable, disable)
// Path: /api/v2/shop_flash_sale/update_shop_flash_sale
// https://open.shopee.com/documents/v2/v2.shop_flash_sale.update_shop_flash_sale?module=123&type=1
func (s *ShopFlashSaleServiceOp[T]) UpdateShopFlashSale(ctx context.Context, sid uint64, mid uint64, tok string, req UpdateShopFlashSaleRequest) (*UpdateShopFlashSaleResponse, error) {
	path := "/shop_flash_sale/update_shop_flash_sale"
	resp := new(UpdateShopFlashSaleResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// UpdateShopFlashSaleItems edit shop flash sale item, you can only edit the models in disbaled or enabled status
// Path: /api/v2/shop_flash_sale/update_shop_flash_sale_items
// https://open.shopee.com/documents/v2/v2.shop_flash_sale.update_shop_flash_sale_items?module=123&type=1
func (s *ShopFlashSaleServiceOp[T]) UpdateShopFlashSaleItems(ctx context.Context, sid uint64, mid uint64, tok string, req UpdateShopFlashSaleItemsRequest) (*UpdateShopFlashSaleItemsResponse, error) {
	path := "/shop_flash_sale/update_shop_flash_sale_items"
	resp := new(UpdateShopFlashSaleItemsResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}
