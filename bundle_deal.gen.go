package goshopee

import (
	"context"
)

type BundleDealService interface {
	// AddBundleDeal create bundle deal. Relevant restrictions refer to FAQ：https://open.shopee.com/faq/254
	// Path: /api/v2/bundle_deal/add_bundle_deal
	// https://open.shopee.com/documents/v2/v2.bundle_deal.add_bundle_deal?module=110&type=1
	AddBundleDeal(ctx context.Context, sid uint64, mid uint64, tok string, req AddBundleDealRequest) (*AddBundleDealResponse, error)
	// AddBundleDealItem add product to bundle deal
	// Path: /api/v2/bundle_deal/add_bundle_deal_item
	// https://open.shopee.com/documents/v2/v2.bundle_deal.add_bundle_deal_item?module=110&type=1
	AddBundleDealItem(ctx context.Context, sid uint64, mid uint64, tok string, req AddBundleDealItemRequest) (*AddBundleDealItemResponse, error)
	// DeleteBundleDeal delete bundle deal
	// Path: /api/v2/bundle_deal/delete_bundle_deal
	// https://open.shopee.com/documents/v2/v2.bundle_deal.delete_bundle_deal?module=110&type=1
	DeleteBundleDeal(ctx context.Context, sid uint64, mid uint64, tok string, req DeleteBundleDealRequest) (*DeleteBundleDealResponse, error)
	// DeleteBundleDealItem delete product in bundle deal
	// Path: /api/v2/bundle_deal/delete_bundle_deal_item
	// https://open.shopee.com/documents/v2/v2.bundle_deal.delete_bundle_deal_item?module=110&type=1
	DeleteBundleDealItem(ctx context.Context, sid uint64, mid uint64, tok string, req DeleteBundleDealItemRequest) (*DeleteBundleDealItemResponse, error)
	// EndBundleDeal end bundle deal
	// Path: /api/v2/bundle_deal/end_bundle_deal
	// https://open.shopee.com/documents/v2/v2.bundle_deal.end_bundle_deal?module=110&type=1
	EndBundleDeal(ctx context.Context, sid uint64, mid uint64, tok string, req EndBundleDealRequest) (*EndBundleDealResponse, error)
	// GetBundleDeal get bundle deal detail
	// Path: /api/v2/bundle_deal/get_bundle_deal
	// https://open.shopee.com/documents/v2/v2.bundle_deal.get_bundle_deal?module=110&type=1
	GetBundleDeal(ctx context.Context, sid uint64, mid uint64, tok string, opt GetBundleDealRequest) (*GetBundleDealResponse, error)
	// GetBundleDealItem get bundle deal item
	// Path: /api/v2/bundle_deal/get_bundle_deal_item
	// https://open.shopee.com/documents/v2/v2.bundle_deal.get_bundle_deal_item?module=110&type=1
	GetBundleDealItem(ctx context.Context, sid uint64, mid uint64, tok string, req GetBundleDealItemRequest) (*GetBundleDealItemResponse, error)
	// GetBundleDealList get bundle deal list
	// Path: /api/v2/bundle_deal/get_bundle_deal_list
	// https://open.shopee.com/documents/v2/v2.bundle_deal.get_bundle_deal_list?module=110&type=1
	GetBundleDealList(ctx context.Context, sid uint64, mid uint64, tok string, opt GetBundleDealListRequest) (*GetBundleDealListResponse, error)
	// UpdateBundleDeal update bundle deal. Relevant restrictions refer to FAQ：https://open.shopee.com/faq/254
	// Path: /api/v2/bundle_deal/update_bundle_deal
	// https://open.shopee.com/documents/v2/v2.bundle_deal.update_bundle_deal?module=110&type=1
	UpdateBundleDeal(ctx context.Context, sid uint64, mid uint64, tok string, req UpdateBundleDealRequest) (*UpdateBundleDealResponse, error)
	// UpdateBundleDealItem update product in bundle deal
	// Path: /api/v2/bundle_deal/update_bundle_deal_item
	// https://open.shopee.com/documents/v2/v2.bundle_deal.update_bundle_deal_item?module=110&type=1
	UpdateBundleDealItem(ctx context.Context, sid uint64, mid uint64, tok string, req UpdateBundleDealItemRequest) (*UpdateBundleDealItemResponse, error)
}

type BundleDealServiceOp[T any] struct {
	client *Client[T]
}

// AddBundleDeal create bundle deal. Relevant restrictions refer to FAQ：https://open.shopee.com/faq/254
// Path: /api/v2/bundle_deal/add_bundle_deal
// https://open.shopee.com/documents/v2/v2.bundle_deal.add_bundle_deal?module=110&type=1
func (s *BundleDealServiceOp[T]) AddBundleDeal(ctx context.Context, sid uint64, mid uint64, tok string, req AddBundleDealRequest) (*AddBundleDealResponse, error) {
	path := "/bundle_deal/add_bundle_deal"
	resp := new(AddBundleDealResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// AddBundleDealItem add product to bundle deal
// Path: /api/v2/bundle_deal/add_bundle_deal_item
// https://open.shopee.com/documents/v2/v2.bundle_deal.add_bundle_deal_item?module=110&type=1
func (s *BundleDealServiceOp[T]) AddBundleDealItem(ctx context.Context, sid uint64, mid uint64, tok string, req AddBundleDealItemRequest) (*AddBundleDealItemResponse, error) {
	path := "/bundle_deal/add_bundle_deal_item"
	resp := new(AddBundleDealItemResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// DeleteBundleDeal delete bundle deal
// Path: /api/v2/bundle_deal/delete_bundle_deal
// https://open.shopee.com/documents/v2/v2.bundle_deal.delete_bundle_deal?module=110&type=1
func (s *BundleDealServiceOp[T]) DeleteBundleDeal(ctx context.Context, sid uint64, mid uint64, tok string, req DeleteBundleDealRequest) (*DeleteBundleDealResponse, error) {
	path := "/bundle_deal/delete_bundle_deal"
	resp := new(DeleteBundleDealResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// DeleteBundleDealItem delete product in bundle deal
// Path: /api/v2/bundle_deal/delete_bundle_deal_item
// https://open.shopee.com/documents/v2/v2.bundle_deal.delete_bundle_deal_item?module=110&type=1
func (s *BundleDealServiceOp[T]) DeleteBundleDealItem(ctx context.Context, sid uint64, mid uint64, tok string, req DeleteBundleDealItemRequest) (*DeleteBundleDealItemResponse, error) {
	path := "/bundle_deal/delete_bundle_deal_item"
	resp := new(DeleteBundleDealItemResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// EndBundleDeal end bundle deal
// Path: /api/v2/bundle_deal/end_bundle_deal
// https://open.shopee.com/documents/v2/v2.bundle_deal.end_bundle_deal?module=110&type=1
func (s *BundleDealServiceOp[T]) EndBundleDeal(ctx context.Context, sid uint64, mid uint64, tok string, req EndBundleDealRequest) (*EndBundleDealResponse, error) {
	path := "/bundle_deal/end_bundle_deal"
	resp := new(EndBundleDealResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// GetBundleDeal get bundle deal detail
// Path: /api/v2/bundle_deal/get_bundle_deal
// https://open.shopee.com/documents/v2/v2.bundle_deal.get_bundle_deal?module=110&type=1
func (s *BundleDealServiceOp[T]) GetBundleDeal(ctx context.Context, sid uint64, mid uint64, tok string, opt GetBundleDealRequest) (*GetBundleDealResponse, error) {
	path := "/bundle_deal/get_bundle_deal"
	resp := new(GetBundleDealResponse)
	err := s.client.Get(ctx, path, resp, opt, sid, mid, tok)
	return resp, err
}

// GetBundleDealItem get bundle deal item
// Path: /api/v2/bundle_deal/get_bundle_deal_item
// https://open.shopee.com/documents/v2/v2.bundle_deal.get_bundle_deal_item?module=110&type=1
func (s *BundleDealServiceOp[T]) GetBundleDealItem(ctx context.Context, sid uint64, mid uint64, tok string, req GetBundleDealItemRequest) (*GetBundleDealItemResponse, error) {
	path := "/bundle_deal/get_bundle_deal_item"
	resp := new(GetBundleDealItemResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// GetBundleDealList get bundle deal list
// Path: /api/v2/bundle_deal/get_bundle_deal_list
// https://open.shopee.com/documents/v2/v2.bundle_deal.get_bundle_deal_list?module=110&type=1
func (s *BundleDealServiceOp[T]) GetBundleDealList(ctx context.Context, sid uint64, mid uint64, tok string, opt GetBundleDealListRequest) (*GetBundleDealListResponse, error) {
	path := "/bundle_deal/get_bundle_deal_list"
	resp := new(GetBundleDealListResponse)
	err := s.client.Get(ctx, path, resp, opt, sid, mid, tok)
	return resp, err
}

// UpdateBundleDeal update bundle deal. Relevant restrictions refer to FAQ：https://open.shopee.com/faq/254
// Path: /api/v2/bundle_deal/update_bundle_deal
// https://open.shopee.com/documents/v2/v2.bundle_deal.update_bundle_deal?module=110&type=1
func (s *BundleDealServiceOp[T]) UpdateBundleDeal(ctx context.Context, sid uint64, mid uint64, tok string, req UpdateBundleDealRequest) (*UpdateBundleDealResponse, error) {
	path := "/bundle_deal/update_bundle_deal"
	resp := new(UpdateBundleDealResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// UpdateBundleDealItem update product in bundle deal
// Path: /api/v2/bundle_deal/update_bundle_deal_item
// https://open.shopee.com/documents/v2/v2.bundle_deal.update_bundle_deal_item?module=110&type=1
func (s *BundleDealServiceOp[T]) UpdateBundleDealItem(ctx context.Context, sid uint64, mid uint64, tok string, req UpdateBundleDealItemRequest) (*UpdateBundleDealItemResponse, error) {
	path := "/bundle_deal/update_bundle_deal_item"
	resp := new(UpdateBundleDealItemResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}
