package goshopee

import (
	"context"
)

type BundleDealService interface {
	// AddBundleDeal create bundle deal. Relevant restrictions refer to FAQ：https://open.shopee.com/faq/254
	// Path: /api/v2/bundle_deal/add_bundle_deal
	// https://open.shopee.com/documents/v2/v2.bundle_deal.add_bundle_deal?module=110&type=1
	AddBundleDeal(ctx context.Context, sid uint64, req AddBundleDealRequest, tok string) (*AddBundleDealResponse, error)
	// AddBundleDealItem add product to bundle deal
	// Path: /api/v2/bundle_deal/add_bundle_deal_item
	// https://open.shopee.com/documents/v2/v2.bundle_deal.add_bundle_deal_item?module=110&type=1
	AddBundleDealItem(ctx context.Context, sid uint64, req AddBundleDealItemRequest, tok string) (*AddBundleDealItemResponse, error)
	// DeleteBundleDeal delete bundle deal
	// Path: /api/v2/bundle_deal/delete_bundle_deal
	// https://open.shopee.com/documents/v2/v2.bundle_deal.delete_bundle_deal?module=110&type=1
	DeleteBundleDeal(ctx context.Context, sid uint64, req DeleteBundleDealRequest, tok string) (*DeleteBundleDealResponse, error)
	// DeleteBundleDealItem delete product in bundle deal
	// Path: /api/v2/bundle_deal/delete_bundle_deal_item
	// https://open.shopee.com/documents/v2/v2.bundle_deal.delete_bundle_deal_item?module=110&type=1
	DeleteBundleDealItem(ctx context.Context, sid uint64, req DeleteBundleDealItemRequest, tok string) (*DeleteBundleDealItemResponse, error)
	// EndBundleDeal end bundle deal
	// Path: /api/v2/bundle_deal/end_bundle_deal
	// https://open.shopee.com/documents/v2/v2.bundle_deal.end_bundle_deal?module=110&type=1
	EndBundleDeal(ctx context.Context, sid uint64, req EndBundleDealRequest, tok string) (*EndBundleDealResponse, error)
	// GetBundleDeal get bundle deal detail
	// Path: /api/v2/bundle_deal/get_bundle_deal
	// https://open.shopee.com/documents/v2/v2.bundle_deal.get_bundle_deal?module=110&type=1
	GetBundleDeal(ctx context.Context, sid uint64, opt GetBundleDealRequest, tok string) (*GetBundleDealResponse, error)
	// GetBundleDealItem get bundle deal item
	// Path: /api/v2/bundle_deal/get_bundle_deal_item
	// https://open.shopee.com/documents/v2/v2.bundle_deal.get_bundle_deal_item?module=110&type=1
	GetBundleDealItem(ctx context.Context, sid uint64, req GetBundleDealItemRequest, tok string) (*GetBundleDealItemResponse, error)
	// GetBundleDealList get bundle deal list
	// Path: /api/v2/bundle_deal/get_bundle_deal_list
	// https://open.shopee.com/documents/v2/v2.bundle_deal.get_bundle_deal_list?module=110&type=1
	GetBundleDealList(ctx context.Context, sid uint64, opt GetBundleDealListRequest, tok string) (*GetBundleDealListResponse, error)
	// UpdateBundleDeal update bundle deal. Relevant restrictions refer to FAQ：https://open.shopee.com/faq/254
	// Path: /api/v2/bundle_deal/update_bundle_deal
	// https://open.shopee.com/documents/v2/v2.bundle_deal.update_bundle_deal?module=110&type=1
	UpdateBundleDeal(ctx context.Context, sid uint64, req UpdateBundleDealRequest, tok string) (*UpdateBundleDealResponse, error)
	// UpdateBundleDealItem update product in bundle deal
	// Path: /api/v2/bundle_deal/update_bundle_deal_item
	// https://open.shopee.com/documents/v2/v2.bundle_deal.update_bundle_deal_item?module=110&type=1
	UpdateBundleDealItem(ctx context.Context, sid uint64, req UpdateBundleDealItemRequest, tok string) (*UpdateBundleDealItemResponse, error)
}

type BundleDealServiceOp[T any] struct {
	client *Client[T]
}

// AddBundleDeal create bundle deal. Relevant restrictions refer to FAQ：https://open.shopee.com/faq/254
// Path: /api/v2/bundle_deal/add_bundle_deal
// https://open.shopee.com/documents/v2/v2.bundle_deal.add_bundle_deal?module=110&type=1
func (s *BundleDealServiceOp[T]) AddBundleDeal(ctx context.Context, sid uint64, req AddBundleDealRequest, tok string) (*AddBundleDealResponse, error) {
	path := "/bundle_deal/add_bundle_deal"
	resp := new(AddBundleDealResponse)
	err := s.client.WithShop(sid, tok).Post(ctx, path, req, resp)
	return resp, err
}

// AddBundleDealItem add product to bundle deal
// Path: /api/v2/bundle_deal/add_bundle_deal_item
// https://open.shopee.com/documents/v2/v2.bundle_deal.add_bundle_deal_item?module=110&type=1
func (s *BundleDealServiceOp[T]) AddBundleDealItem(ctx context.Context, sid uint64, req AddBundleDealItemRequest, tok string) (*AddBundleDealItemResponse, error) {
	path := "/bundle_deal/add_bundle_deal_item"
	resp := new(AddBundleDealItemResponse)
	err := s.client.WithShop(sid, tok).Post(ctx, path, req, resp)
	return resp, err
}

// DeleteBundleDeal delete bundle deal
// Path: /api/v2/bundle_deal/delete_bundle_deal
// https://open.shopee.com/documents/v2/v2.bundle_deal.delete_bundle_deal?module=110&type=1
func (s *BundleDealServiceOp[T]) DeleteBundleDeal(ctx context.Context, sid uint64, req DeleteBundleDealRequest, tok string) (*DeleteBundleDealResponse, error) {
	path := "/bundle_deal/delete_bundle_deal"
	resp := new(DeleteBundleDealResponse)
	err := s.client.WithShop(sid, tok).Post(ctx, path, req, resp)
	return resp, err
}

// DeleteBundleDealItem delete product in bundle deal
// Path: /api/v2/bundle_deal/delete_bundle_deal_item
// https://open.shopee.com/documents/v2/v2.bundle_deal.delete_bundle_deal_item?module=110&type=1
func (s *BundleDealServiceOp[T]) DeleteBundleDealItem(ctx context.Context, sid uint64, req DeleteBundleDealItemRequest, tok string) (*DeleteBundleDealItemResponse, error) {
	path := "/bundle_deal/delete_bundle_deal_item"
	resp := new(DeleteBundleDealItemResponse)
	err := s.client.WithShop(sid, tok).Post(ctx, path, req, resp)
	return resp, err
}

// EndBundleDeal end bundle deal
// Path: /api/v2/bundle_deal/end_bundle_deal
// https://open.shopee.com/documents/v2/v2.bundle_deal.end_bundle_deal?module=110&type=1
func (s *BundleDealServiceOp[T]) EndBundleDeal(ctx context.Context, sid uint64, req EndBundleDealRequest, tok string) (*EndBundleDealResponse, error) {
	path := "/bundle_deal/end_bundle_deal"
	resp := new(EndBundleDealResponse)
	err := s.client.WithShop(sid, tok).Post(ctx, path, req, resp)
	return resp, err
}

// GetBundleDeal get bundle deal detail
// Path: /api/v2/bundle_deal/get_bundle_deal
// https://open.shopee.com/documents/v2/v2.bundle_deal.get_bundle_deal?module=110&type=1
func (s *BundleDealServiceOp[T]) GetBundleDeal(ctx context.Context, sid uint64, opt GetBundleDealRequest, tok string) (*GetBundleDealResponse, error) {
	path := "/bundle_deal/get_bundle_deal"
	resp := new(GetBundleDealResponse)
	err := s.client.WithShop(sid, tok).Get(ctx, path, resp, opt)
	return resp, err
}

// GetBundleDealItem get bundle deal item
// Path: /api/v2/bundle_deal/get_bundle_deal_item
// https://open.shopee.com/documents/v2/v2.bundle_deal.get_bundle_deal_item?module=110&type=1
func (s *BundleDealServiceOp[T]) GetBundleDealItem(ctx context.Context, sid uint64, req GetBundleDealItemRequest, tok string) (*GetBundleDealItemResponse, error) {
	path := "/bundle_deal/get_bundle_deal_item"
	resp := new(GetBundleDealItemResponse)
	err := s.client.WithShop(sid, tok).Post(ctx, path, req, resp)
	return resp, err
}

// GetBundleDealList get bundle deal list
// Path: /api/v2/bundle_deal/get_bundle_deal_list
// https://open.shopee.com/documents/v2/v2.bundle_deal.get_bundle_deal_list?module=110&type=1
func (s *BundleDealServiceOp[T]) GetBundleDealList(ctx context.Context, sid uint64, opt GetBundleDealListRequest, tok string) (*GetBundleDealListResponse, error) {
	path := "/bundle_deal/get_bundle_deal_list"
	resp := new(GetBundleDealListResponse)
	err := s.client.WithShop(sid, tok).Get(ctx, path, resp, opt)
	return resp, err
}

// UpdateBundleDeal update bundle deal. Relevant restrictions refer to FAQ：https://open.shopee.com/faq/254
// Path: /api/v2/bundle_deal/update_bundle_deal
// https://open.shopee.com/documents/v2/v2.bundle_deal.update_bundle_deal?module=110&type=1
func (s *BundleDealServiceOp[T]) UpdateBundleDeal(ctx context.Context, sid uint64, req UpdateBundleDealRequest, tok string) (*UpdateBundleDealResponse, error) {
	path := "/bundle_deal/update_bundle_deal"
	resp := new(UpdateBundleDealResponse)
	err := s.client.WithShop(sid, tok).Post(ctx, path, req, resp)
	return resp, err
}

// UpdateBundleDealItem update product in bundle deal
// Path: /api/v2/bundle_deal/update_bundle_deal_item
// https://open.shopee.com/documents/v2/v2.bundle_deal.update_bundle_deal_item?module=110&type=1
func (s *BundleDealServiceOp[T]) UpdateBundleDealItem(ctx context.Context, sid uint64, req UpdateBundleDealItemRequest, tok string) (*UpdateBundleDealItemResponse, error) {
	path := "/bundle_deal/update_bundle_deal_item"
	resp := new(UpdateBundleDealItemResponse)
	err := s.client.WithShop(sid, tok).Post(ctx, path, req, resp)
	return resp, err
}
