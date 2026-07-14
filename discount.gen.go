package goshopee

import (
	"context"
)

type DiscountService interface {
	// AddDiscount Use this api to add shop discount activity
	// Path: /api/v2/discount/add_discount
	// https://open.shopee.com/documents/v2/v2.discount.add_discount?module=99&type=1
	AddDiscount(ctx context.Context, sid uint64, mid uint64, tok string, req AddDiscountRequest) (*AddDiscountResponse, error)
	// AddDiscountItem Use this api to add shop discount item.
	// Path: /api/v2/discount/add_discount_item
	// https://open.shopee.com/documents/v2/v2.discount.add_discount_item?module=99&type=1
	AddDiscountItem(ctx context.Context, sid uint64, mid uint64, tok string, req AddDiscountItemRequest) (*AddDiscountItemResponse, error)
	// DeleteDiscount Use this api to delete one discount activity
	// Path: /api/v2/discount/delete_discount
	// https://open.shopee.com/documents/v2/v2.discount.delete_discount?module=99&type=1
	DeleteDiscount(ctx context.Context, sid uint64, mid uint64, tok string, req DeleteDiscountRequest) (*DeleteDiscountResponse, error)
	// DeleteDiscountItem Use this api to delete items of the discount activity
	// Path: /api/v2/discount/delete_discount_item
	// https://open.shopee.com/documents/v2/v2.discount.delete_discount_item?module=99&type=1
	DeleteDiscountItem(ctx context.Context, sid uint64, mid uint64, tok string, req DeleteDiscountItemRequest) (*DeleteDiscountItemResponse, error)
	// DeleteSipDiscount Delete SIP Overseas Discounts for SIP affiliate region. Please use Primary shop's Shop ID to request, and provide the region of the Affiliate shop to be deleted, the API will delete the discount from that region's Affiliate shop.
	// Path: /api/v2/discount/delete_sip_discount
	// https://open.shopee.com/documents/v2/v2.discount.delete_sip_discount?module=99&type=1
	DeleteSipDiscount(ctx context.Context, sid uint64, mid uint64, tok string, req DeleteSipDiscountRequest) (*DeleteSipDiscountResponse, error)
	// EndDiscount Use this api to end shop discount activity
	// Path: /api/v2/discount/end_discount
	// https://open.shopee.com/documents/v2/v2.discount.end_discount?module=99&type=1
	EndDiscount(ctx context.Context, sid uint64, mid uint64, tok string, req EndDiscountRequest) (*EndDiscountResponse, error)
	// GetDiscount Use this api to get one shop discount activity detail
	// Path: /api/v2/discount/get_discount
	// https://open.shopee.com/documents/v2/v2.discount.get_discount?module=99&type=1
	GetDiscount(ctx context.Context, sid uint64, mid uint64, tok string, opt GetDiscountRequest) (*GetDiscountResponse, error)
	// GetDiscountList Use this api to get shop discount activity list
	// Path: /api/v2/discount/get_discount_list
	// https://open.shopee.com/documents/v2/v2.discount.get_discount_list?module=99&type=1
	GetDiscountList(ctx context.Context, sid uint64, mid uint64, tok string, req GetDiscountListRequest) (*GetDiscountListResponse, error)
	// GetSipDiscounts Get SIP Overseas Discounts. Only regions that have upcoming/ongoing discounts will be returned. Please use Primary shop's Shop ID to request, the API will return the list of Affiliate shops under this Primary shop that have set discounts, along with the discount details.
	// Path: /api/v2/discount/get_sip_discounts
	// https://open.shopee.com/documents/v2/v2.discount.get_sip_discounts?module=99&type=1
	GetSipDiscounts(ctx context.Context, sid uint64, mid uint64, tok string, opt GetSipDiscountsRequest) (*GetSipDiscountsResponse, error)
	// SetSipDiscount Set SIP Overseas Discount for SIP affiliate region. Please use Primary shop's Shop ID to request, and provide the region and discount rate of the Affiliate shop to be set or update, the API will set or update the discount rate for that region's Affiliate shop.
	// Path: /api/v2/discount/set_sip_discount
	// https://open.shopee.com/documents/v2/v2.discount.set_sip_discount?module=99&type=1
	SetSipDiscount(ctx context.Context, sid uint64, mid uint64, tok string, req SetSipDiscountRequest) (*SetSipDiscountResponse, error)
	// UpdateDiscount Use this api to update one discount information
	// Path: /api/v2/discount/update_discount
	// https://open.shopee.com/documents/v2/v2.discount.update_discount?module=99&type=1
	UpdateDiscount(ctx context.Context, sid uint64, mid uint64, tok string, req UpdateDiscountRequest) (*UpdateDiscountResponse, error)
	// UpdateDiscountItem Use this api to update items of the discount promotion.
	// Path: /api/v2/discount/update_discount_item
	// https://open.shopee.com/documents/v2/v2.discount.update_discount_item?module=99&type=1
	UpdateDiscountItem(ctx context.Context, sid uint64, mid uint64, tok string, req UpdateDiscountItemRequest) (*UpdateDiscountItemResponse, error)
}

type DiscountServiceOp[T any] struct {
	client *Client[T]
}

// AddDiscount Use this api to add shop discount activity
// Path: /api/v2/discount/add_discount
// https://open.shopee.com/documents/v2/v2.discount.add_discount?module=99&type=1
func (s *DiscountServiceOp[T]) AddDiscount(ctx context.Context, sid uint64, mid uint64, tok string, req AddDiscountRequest) (*AddDiscountResponse, error) {
	path := "/discount/add_discount"
	resp := new(AddDiscountResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// AddDiscountItem Use this api to add shop discount item.
// Path: /api/v2/discount/add_discount_item
// https://open.shopee.com/documents/v2/v2.discount.add_discount_item?module=99&type=1
func (s *DiscountServiceOp[T]) AddDiscountItem(ctx context.Context, sid uint64, mid uint64, tok string, req AddDiscountItemRequest) (*AddDiscountItemResponse, error) {
	path := "/discount/add_discount_item"
	resp := new(AddDiscountItemResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// DeleteDiscount Use this api to delete one discount activity
// Path: /api/v2/discount/delete_discount
// https://open.shopee.com/documents/v2/v2.discount.delete_discount?module=99&type=1
func (s *DiscountServiceOp[T]) DeleteDiscount(ctx context.Context, sid uint64, mid uint64, tok string, req DeleteDiscountRequest) (*DeleteDiscountResponse, error) {
	path := "/discount/delete_discount"
	resp := new(DeleteDiscountResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// DeleteDiscountItem Use this api to delete items of the discount activity
// Path: /api/v2/discount/delete_discount_item
// https://open.shopee.com/documents/v2/v2.discount.delete_discount_item?module=99&type=1
func (s *DiscountServiceOp[T]) DeleteDiscountItem(ctx context.Context, sid uint64, mid uint64, tok string, req DeleteDiscountItemRequest) (*DeleteDiscountItemResponse, error) {
	path := "/discount/delete_discount_item"
	resp := new(DeleteDiscountItemResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// DeleteSipDiscount Delete SIP Overseas Discounts for SIP affiliate region. Please use Primary shop's Shop ID to request, and provide the region of the Affiliate shop to be deleted, the API will delete the discount from that region's Affiliate shop.
// Path: /api/v2/discount/delete_sip_discount
// https://open.shopee.com/documents/v2/v2.discount.delete_sip_discount?module=99&type=1
func (s *DiscountServiceOp[T]) DeleteSipDiscount(ctx context.Context, sid uint64, mid uint64, tok string, req DeleteSipDiscountRequest) (*DeleteSipDiscountResponse, error) {
	path := "/discount/delete_sip_discount"
	resp := new(DeleteSipDiscountResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// EndDiscount Use this api to end shop discount activity
// Path: /api/v2/discount/end_discount
// https://open.shopee.com/documents/v2/v2.discount.end_discount?module=99&type=1
func (s *DiscountServiceOp[T]) EndDiscount(ctx context.Context, sid uint64, mid uint64, tok string, req EndDiscountRequest) (*EndDiscountResponse, error) {
	path := "/discount/end_discount"
	resp := new(EndDiscountResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// GetDiscount Use this api to get one shop discount activity detail
// Path: /api/v2/discount/get_discount
// https://open.shopee.com/documents/v2/v2.discount.get_discount?module=99&type=1
func (s *DiscountServiceOp[T]) GetDiscount(ctx context.Context, sid uint64, mid uint64, tok string, opt GetDiscountRequest) (*GetDiscountResponse, error) {
	path := "/discount/get_discount"
	resp := new(GetDiscountResponse)
	err := s.client.Get(ctx, path, resp, opt, sid, mid, tok)
	return resp, err
}

// GetDiscountList Use this api to get shop discount activity list
// Path: /api/v2/discount/get_discount_list
// https://open.shopee.com/documents/v2/v2.discount.get_discount_list?module=99&type=1
func (s *DiscountServiceOp[T]) GetDiscountList(ctx context.Context, sid uint64, mid uint64, tok string, req GetDiscountListRequest) (*GetDiscountListResponse, error) {
	path := "/discount/get_discount_list"
	resp := new(GetDiscountListResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// GetSipDiscounts Get SIP Overseas Discounts. Only regions that have upcoming/ongoing discounts will be returned. Please use Primary shop's Shop ID to request, the API will return the list of Affiliate shops under this Primary shop that have set discounts, along with the discount details.
// Path: /api/v2/discount/get_sip_discounts
// https://open.shopee.com/documents/v2/v2.discount.get_sip_discounts?module=99&type=1
func (s *DiscountServiceOp[T]) GetSipDiscounts(ctx context.Context, sid uint64, mid uint64, tok string, opt GetSipDiscountsRequest) (*GetSipDiscountsResponse, error) {
	path := "/discount/get_sip_discounts"
	resp := new(GetSipDiscountsResponse)
	err := s.client.Get(ctx, path, resp, opt, sid, mid, tok)
	return resp, err
}

// SetSipDiscount Set SIP Overseas Discount for SIP affiliate region. Please use Primary shop's Shop ID to request, and provide the region and discount rate of the Affiliate shop to be set or update, the API will set or update the discount rate for that region's Affiliate shop.
// Path: /api/v2/discount/set_sip_discount
// https://open.shopee.com/documents/v2/v2.discount.set_sip_discount?module=99&type=1
func (s *DiscountServiceOp[T]) SetSipDiscount(ctx context.Context, sid uint64, mid uint64, tok string, req SetSipDiscountRequest) (*SetSipDiscountResponse, error) {
	path := "/discount/set_sip_discount"
	resp := new(SetSipDiscountResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// UpdateDiscount Use this api to update one discount information
// Path: /api/v2/discount/update_discount
// https://open.shopee.com/documents/v2/v2.discount.update_discount?module=99&type=1
func (s *DiscountServiceOp[T]) UpdateDiscount(ctx context.Context, sid uint64, mid uint64, tok string, req UpdateDiscountRequest) (*UpdateDiscountResponse, error) {
	path := "/discount/update_discount"
	resp := new(UpdateDiscountResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// UpdateDiscountItem Use this api to update items of the discount promotion.
// Path: /api/v2/discount/update_discount_item
// https://open.shopee.com/documents/v2/v2.discount.update_discount_item?module=99&type=1
func (s *DiscountServiceOp[T]) UpdateDiscountItem(ctx context.Context, sid uint64, mid uint64, tok string, req UpdateDiscountItemRequest) (*UpdateDiscountItemResponse, error) {
	path := "/discount/update_discount_item"
	resp := new(UpdateDiscountItemResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}
