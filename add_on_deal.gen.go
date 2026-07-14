package goshopee

import (
	"context"
)

type AddOnDealService interface {
	// AddAddOnDeal Add Add-on Deal
	// Path: /api/v2/add_on_deal/add_add_on_deal
	// https://open.shopee.com/documents/v2/v2.add_on_deal.add_add_on_deal?module=111&type=1
	AddAddOnDeal(ctx context.Context, sid uint64, mid uint64, tok string, req AddAddOnDealRequest) (*AddAddOnDealResponse, error)
	// AddAddOnDealMainItem Add Add-on Deal Main Item
	// Path: /api/v2/add_on_deal/add_add_on_deal_main_item
	// https://open.shopee.com/documents/v2/v2.add_on_deal.add_add_on_deal_main_item?module=111&type=1
	AddAddOnDealMainItem(ctx context.Context, sid uint64, mid uint64, tok string, req AddAddOnDealMainItemRequest) (*AddAddOnDealMainItemResponse, error)
	// AddAddOnDealSubItem Add Add-on Deal Sub Item
	// Path: /api/v2/add_on_deal/add_add_on_deal_sub_item
	// https://open.shopee.com/documents/v2/v2.add_on_deal.add_add_on_deal_sub_item?module=111&type=1
	AddAddOnDealSubItem(ctx context.Context, sid uint64, mid uint64, tok string, req AddAddOnDealSubItemRequest) (*AddAddOnDealSubItemResponse, error)
	// DeleteAddOnDeal Delete Add-on Deal
	// Path: /api/v2/add_on_deal/delete_add_on_deal
	// https://open.shopee.com/documents/v2/v2.add_on_deal.delete_add_on_deal?module=111&type=1
	DeleteAddOnDeal(ctx context.Context, sid uint64, mid uint64, tok string, req DeleteAddOnDealRequest) (*DeleteAddOnDealResponse, error)
	// DeleteAddOnDealMainItem Delete Add-on Deal Main Item
	// Path: /api/v2/add_on_deal/delete_add_on_deal_main_item
	// https://open.shopee.com/documents/v2/v2.add_on_deal.delete_add_on_deal_main_item?module=111&type=1
	DeleteAddOnDealMainItem(ctx context.Context, sid uint64, mid uint64, tok string, req DeleteAddOnDealMainItemRequest) (*DeleteAddOnDealMainItemResponse, error)
	// DeleteAddOnDealSubItem Delete Add-on Deal Sub Item
	// Path: /api/v2/add_on_deal/delete_add_on_deal_sub_item
	// https://open.shopee.com/documents/v2/v2.add_on_deal.delete_add_on_deal_sub_item?module=111&type=1
	DeleteAddOnDealSubItem(ctx context.Context, sid uint64, mid uint64, tok string, req DeleteAddOnDealSubItemRequest) (*DeleteAddOnDealSubItemResponse, error)
	// EndAddOnDeal End Add-on Deal
	// Path: /api/v2/add_on_deal/end_add_on_deal
	// https://open.shopee.com/documents/v2/v2.add_on_deal.end_add_on_deal?module=111&type=1
	EndAddOnDeal(ctx context.Context, sid uint64, mid uint64, tok string, req EndAddOnDealRequest) (*EndAddOnDealResponse, error)
	// GetAddOnDeal Get Add-on Deal
	// Path: /api/v2/add_on_deal/get_add_on_deal
	// https://open.shopee.com/documents/v2/v2.add_on_deal.get_add_on_deal?module=111&type=1
	GetAddOnDeal(ctx context.Context, sid uint64, mid uint64, tok string, req GetAddOnDealRequest) (*GetAddOnDealResponse, error)
	// GetAddOnDealList Get Add-on Deal List
	// Path: /api/v2/add_on_deal/get_add_on_deal_list
	// https://open.shopee.com/documents/v2/v2.add_on_deal.get_add_on_deal_list?module=111&type=1
	GetAddOnDealList(ctx context.Context, sid uint64, mid uint64, tok string, req GetAddOnDealListRequest) (*GetAddOnDealListResponse, error)
	// GetAddOnDealMainItem Get Add-on Deal Main Item
	// Path: /api/v2/add_on_deal/get_add_on_deal_main_item
	// https://open.shopee.com/documents/v2/v2.add_on_deal.get_add_on_deal_main_item?module=111&type=1
	GetAddOnDealMainItem(ctx context.Context, sid uint64, mid uint64, tok string, req GetAddOnDealMainItemRequest) (*GetAddOnDealMainItemResponse, error)
	// GetAddOnDealSubItem Get Add-on Deal Sub Item
	// Path: /api/v2/add_on_deal/get_add_on_deal_sub_item
	// https://open.shopee.com/documents/v2/v2.add_on_deal.get_add_on_deal_sub_item?module=111&type=1
	GetAddOnDealSubItem(ctx context.Context, sid uint64, mid uint64, tok string, req GetAddOnDealSubItemRequest) (*GetAddOnDealSubItemResponse, error)
	// UpdateAddOnDeal Update Add-on Deal
	// Path: /api/v2/add_on_deal/update_add_on_deal
	// https://open.shopee.com/documents/v2/v2.add_on_deal.update_add_on_deal?module=111&type=1
	UpdateAddOnDeal(ctx context.Context, sid uint64, mid uint64, tok string, req UpdateAddOnDealRequest) (*UpdateAddOnDealResponse, error)
	// UpdateAddOnDealMainItem Update Add-on Deal Main Item
	// Path: /api/v2/add_on_deal/update_add_on_deal_main_item
	// https://open.shopee.com/documents/v2/v2.add_on_deal.update_add_on_deal_main_item?module=111&type=1
	UpdateAddOnDealMainItem(ctx context.Context, sid uint64, mid uint64, tok string, req UpdateAddOnDealMainItemRequest) (*UpdateAddOnDealMainItemResponse, error)
	// UpdateAddOnDealSubItem Update Add-on Deal Sub Item
	// Path: /api/v2/add_on_deal/update_add_on_deal_sub_item
	// https://open.shopee.com/documents/v2/v2.add_on_deal.update_add_on_deal_sub_item?module=111&type=1
	UpdateAddOnDealSubItem(ctx context.Context, sid uint64, mid uint64, tok string, req UpdateAddOnDealSubItemRequest) (*UpdateAddOnDealSubItemResponse, error)
}

type AddOnDealServiceOp[T any] struct {
	client *Client[T]
}

// AddAddOnDeal Add Add-on Deal
// Path: /api/v2/add_on_deal/add_add_on_deal
// https://open.shopee.com/documents/v2/v2.add_on_deal.add_add_on_deal?module=111&type=1
func (s *AddOnDealServiceOp[T]) AddAddOnDeal(ctx context.Context, sid uint64, mid uint64, tok string, req AddAddOnDealRequest) (*AddAddOnDealResponse, error) {
	path := "/add_on_deal/add_add_on_deal"
	resp := new(AddAddOnDealResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// AddAddOnDealMainItem Add Add-on Deal Main Item
// Path: /api/v2/add_on_deal/add_add_on_deal_main_item
// https://open.shopee.com/documents/v2/v2.add_on_deal.add_add_on_deal_main_item?module=111&type=1
func (s *AddOnDealServiceOp[T]) AddAddOnDealMainItem(ctx context.Context, sid uint64, mid uint64, tok string, req AddAddOnDealMainItemRequest) (*AddAddOnDealMainItemResponse, error) {
	path := "/add_on_deal/add_add_on_deal_main_item"
	resp := new(AddAddOnDealMainItemResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// AddAddOnDealSubItem Add Add-on Deal Sub Item
// Path: /api/v2/add_on_deal/add_add_on_deal_sub_item
// https://open.shopee.com/documents/v2/v2.add_on_deal.add_add_on_deal_sub_item?module=111&type=1
func (s *AddOnDealServiceOp[T]) AddAddOnDealSubItem(ctx context.Context, sid uint64, mid uint64, tok string, req AddAddOnDealSubItemRequest) (*AddAddOnDealSubItemResponse, error) {
	path := "/add_on_deal/add_add_on_deal_sub_item"
	resp := new(AddAddOnDealSubItemResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// DeleteAddOnDeal Delete Add-on Deal
// Path: /api/v2/add_on_deal/delete_add_on_deal
// https://open.shopee.com/documents/v2/v2.add_on_deal.delete_add_on_deal?module=111&type=1
func (s *AddOnDealServiceOp[T]) DeleteAddOnDeal(ctx context.Context, sid uint64, mid uint64, tok string, req DeleteAddOnDealRequest) (*DeleteAddOnDealResponse, error) {
	path := "/add_on_deal/delete_add_on_deal"
	resp := new(DeleteAddOnDealResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// DeleteAddOnDealMainItem Delete Add-on Deal Main Item
// Path: /api/v2/add_on_deal/delete_add_on_deal_main_item
// https://open.shopee.com/documents/v2/v2.add_on_deal.delete_add_on_deal_main_item?module=111&type=1
func (s *AddOnDealServiceOp[T]) DeleteAddOnDealMainItem(ctx context.Context, sid uint64, mid uint64, tok string, req DeleteAddOnDealMainItemRequest) (*DeleteAddOnDealMainItemResponse, error) {
	path := "/add_on_deal/delete_add_on_deal_main_item"
	resp := new(DeleteAddOnDealMainItemResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// DeleteAddOnDealSubItem Delete Add-on Deal Sub Item
// Path: /api/v2/add_on_deal/delete_add_on_deal_sub_item
// https://open.shopee.com/documents/v2/v2.add_on_deal.delete_add_on_deal_sub_item?module=111&type=1
func (s *AddOnDealServiceOp[T]) DeleteAddOnDealSubItem(ctx context.Context, sid uint64, mid uint64, tok string, req DeleteAddOnDealSubItemRequest) (*DeleteAddOnDealSubItemResponse, error) {
	path := "/add_on_deal/delete_add_on_deal_sub_item"
	resp := new(DeleteAddOnDealSubItemResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// EndAddOnDeal End Add-on Deal
// Path: /api/v2/add_on_deal/end_add_on_deal
// https://open.shopee.com/documents/v2/v2.add_on_deal.end_add_on_deal?module=111&type=1
func (s *AddOnDealServiceOp[T]) EndAddOnDeal(ctx context.Context, sid uint64, mid uint64, tok string, req EndAddOnDealRequest) (*EndAddOnDealResponse, error) {
	path := "/add_on_deal/end_add_on_deal"
	resp := new(EndAddOnDealResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// GetAddOnDeal Get Add-on Deal
// Path: /api/v2/add_on_deal/get_add_on_deal
// https://open.shopee.com/documents/v2/v2.add_on_deal.get_add_on_deal?module=111&type=1
func (s *AddOnDealServiceOp[T]) GetAddOnDeal(ctx context.Context, sid uint64, mid uint64, tok string, req GetAddOnDealRequest) (*GetAddOnDealResponse, error) {
	path := "/add_on_deal/get_add_on_deal"
	resp := new(GetAddOnDealResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// GetAddOnDealList Get Add-on Deal List
// Path: /api/v2/add_on_deal/get_add_on_deal_list
// https://open.shopee.com/documents/v2/v2.add_on_deal.get_add_on_deal_list?module=111&type=1
func (s *AddOnDealServiceOp[T]) GetAddOnDealList(ctx context.Context, sid uint64, mid uint64, tok string, req GetAddOnDealListRequest) (*GetAddOnDealListResponse, error) {
	path := "/add_on_deal/get_add_on_deal_list"
	resp := new(GetAddOnDealListResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// GetAddOnDealMainItem Get Add-on Deal Main Item
// Path: /api/v2/add_on_deal/get_add_on_deal_main_item
// https://open.shopee.com/documents/v2/v2.add_on_deal.get_add_on_deal_main_item?module=111&type=1
func (s *AddOnDealServiceOp[T]) GetAddOnDealMainItem(ctx context.Context, sid uint64, mid uint64, tok string, req GetAddOnDealMainItemRequest) (*GetAddOnDealMainItemResponse, error) {
	path := "/add_on_deal/get_add_on_deal_main_item"
	resp := new(GetAddOnDealMainItemResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// GetAddOnDealSubItem Get Add-on Deal Sub Item
// Path: /api/v2/add_on_deal/get_add_on_deal_sub_item
// https://open.shopee.com/documents/v2/v2.add_on_deal.get_add_on_deal_sub_item?module=111&type=1
func (s *AddOnDealServiceOp[T]) GetAddOnDealSubItem(ctx context.Context, sid uint64, mid uint64, tok string, req GetAddOnDealSubItemRequest) (*GetAddOnDealSubItemResponse, error) {
	path := "/add_on_deal/get_add_on_deal_sub_item"
	resp := new(GetAddOnDealSubItemResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// UpdateAddOnDeal Update Add-on Deal
// Path: /api/v2/add_on_deal/update_add_on_deal
// https://open.shopee.com/documents/v2/v2.add_on_deal.update_add_on_deal?module=111&type=1
func (s *AddOnDealServiceOp[T]) UpdateAddOnDeal(ctx context.Context, sid uint64, mid uint64, tok string, req UpdateAddOnDealRequest) (*UpdateAddOnDealResponse, error) {
	path := "/add_on_deal/update_add_on_deal"
	resp := new(UpdateAddOnDealResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// UpdateAddOnDealMainItem Update Add-on Deal Main Item
// Path: /api/v2/add_on_deal/update_add_on_deal_main_item
// https://open.shopee.com/documents/v2/v2.add_on_deal.update_add_on_deal_main_item?module=111&type=1
func (s *AddOnDealServiceOp[T]) UpdateAddOnDealMainItem(ctx context.Context, sid uint64, mid uint64, tok string, req UpdateAddOnDealMainItemRequest) (*UpdateAddOnDealMainItemResponse, error) {
	path := "/add_on_deal/update_add_on_deal_main_item"
	resp := new(UpdateAddOnDealMainItemResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// UpdateAddOnDealSubItem Update Add-on Deal Sub Item
// Path: /api/v2/add_on_deal/update_add_on_deal_sub_item
// https://open.shopee.com/documents/v2/v2.add_on_deal.update_add_on_deal_sub_item?module=111&type=1
func (s *AddOnDealServiceOp[T]) UpdateAddOnDealSubItem(ctx context.Context, sid uint64, mid uint64, tok string, req UpdateAddOnDealSubItemRequest) (*UpdateAddOnDealSubItemResponse, error) {
	path := "/add_on_deal/update_add_on_deal_sub_item"
	resp := new(UpdateAddOnDealSubItemResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}
