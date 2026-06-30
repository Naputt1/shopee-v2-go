package goshopee

import (
	"context"
)

type DiscountService interface {
	// AddDiscount Use this api to add shop discount activity
	// Path: /api/v2/discount/add_discount
	// https://open.shopee.com/documents/v2/v2.discount.add_discount?module=99&type=1
	AddDiscount(ctx context.Context, sid uint64, req AddDiscountRequest, tok string) (*AddDiscountResponse, error)
	// AddDiscountItem Use this api to add shop discount item.
	// Path: /api/v2/discount/add_discount_item
	// https://open.shopee.com/documents/v2/v2.discount.add_discount_item?module=99&type=1
	AddDiscountItem(ctx context.Context, sid uint64, req AddDiscountItemRequest, tok string) (*AddDiscountItemResponse, error)
	// DeleteDiscount Use this api to delete one discount activity
	// Path: /api/v2/discount/delete_discount
	// https://open.shopee.com/documents/v2/v2.discount.delete_discount?module=99&type=1
	DeleteDiscount(ctx context.Context, sid uint64, req DeleteDiscountRequest, tok string) (*DeleteDiscountResponse, error)
	// DeleteDiscountItem Use this api to delete items of the discount activity
	// Path: /api/v2/discount/delete_discount_item
	// https://open.shopee.com/documents/v2/v2.discount.delete_discount_item?module=99&type=1
	DeleteDiscountItem(ctx context.Context, sid uint64, req DeleteDiscountItemRequest, tok string) (*DeleteDiscountItemResponse, error)
	// DeleteSipDiscount Delete SIP Overseas Discounts for SIP affiliate region. Please use Primary shop's Shop ID to request, and provide the region of the Affiliate shop to be deleted, the API will delete the discount from that region's Affiliate shop.
	// Path: /api/v2/discount/delete_sip_discount
	// https://open.shopee.com/documents/v2/v2.discount.delete_sip_discount?module=99&type=1
	DeleteSipDiscount(ctx context.Context, sid uint64, req DeleteSipDiscountRequest, tok string) (*DeleteSipDiscountResponse, error)
	// EndDiscount Use this api to end shop discount activity
	// Path: /api/v2/discount/end_discount
	// https://open.shopee.com/documents/v2/v2.discount.end_discount?module=99&type=1
	EndDiscount(ctx context.Context, sid uint64, req EndDiscountRequest, tok string) (*EndDiscountResponse, error)
	// GetDiscount Use this api to get one shop discount activity detail
	// Path: /api/v2/discount/get_discount
	// https://open.shopee.com/documents/v2/v2.discount.get_discount?module=99&type=1
	GetDiscount(ctx context.Context, sid uint64, opt GetDiscountRequest, tok string) (*GetDiscountResponse, error)
	// GetDiscountList Use this api to get shop discount activity list
	// Path: /api/v2/discount/get_discount_list
	// https://open.shopee.com/documents/v2/v2.discount.get_discount_list?module=99&type=1
	GetDiscountList(ctx context.Context, sid uint64, req GetDiscountListRequest, tok string) (*GetDiscountListResponse, error)
	// GetSipDiscounts Get SIP Overseas Discounts. Only regions that have upcoming/ongoing discounts will be returned. Please use Primary shop's Shop ID to request, the API will return the list of Affiliate shops under this Primary shop that have set discounts, along with the discount details.
	// Path: /api/v2/discount/get_sip_discounts
	// https://open.shopee.com/documents/v2/v2.discount.get_sip_discounts?module=99&type=1
	GetSipDiscounts(ctx context.Context, sid uint64, opt GetSipDiscountsRequest, tok string) (*GetSipDiscountsResponse, error)
	// SetSipDiscount Set SIP Overseas Discount for SIP affiliate region. Please use Primary shop's Shop ID to request, and provide the region and discount rate of the Affiliate shop to be set or update, the API will set or update the discount rate for that region's Affiliate shop.
	// Path: /api/v2/discount/set_sip_discount
	// https://open.shopee.com/documents/v2/v2.discount.set_sip_discount?module=99&type=1
	SetSipDiscount(ctx context.Context, sid uint64, req SetSipDiscountRequest, tok string) (*SetSipDiscountResponse, error)
	// UpdateDiscount Use this api to update one discount information
	// Path: /api/v2/discount/update_discount
	// https://open.shopee.com/documents/v2/v2.discount.update_discount?module=99&type=1
	UpdateDiscount(ctx context.Context, sid uint64, req UpdateDiscountRequest, tok string) (*UpdateDiscountResponse, error)
	// UpdateDiscountItem Use this api to update items of the discount promotion.
	// Path: /api/v2/discount/update_discount_item
	// https://open.shopee.com/documents/v2/v2.discount.update_discount_item?module=99&type=1
	UpdateDiscountItem(ctx context.Context, sid uint64, req UpdateDiscountItemRequest, tok string) (*UpdateDiscountItemResponse, error)
}

type DiscountServiceOp[T any] struct {
	client *Client[T]
}

// AddDiscount Use this api to add shop discount activity
// Path: /api/v2/discount/add_discount
// https://open.shopee.com/documents/v2/v2.discount.add_discount?module=99&type=1
func (s *DiscountServiceOp[T]) AddDiscount(ctx context.Context, sid uint64, req AddDiscountRequest, tok string) (*AddDiscountResponse, error) {
	path := "/discount/add_discount"
	resp := new(AddDiscountResponse)
	err := s.client.WithShop(sid, tok).Post(ctx, path, req, resp)
	return resp, err
}

// AddDiscountItem Use this api to add shop discount item.
// Path: /api/v2/discount/add_discount_item
// https://open.shopee.com/documents/v2/v2.discount.add_discount_item?module=99&type=1
func (s *DiscountServiceOp[T]) AddDiscountItem(ctx context.Context, sid uint64, req AddDiscountItemRequest, tok string) (*AddDiscountItemResponse, error) {
	path := "/discount/add_discount_item"
	resp := new(AddDiscountItemResponse)
	err := s.client.WithShop(sid, tok).Post(ctx, path, req, resp)
	return resp, err
}

// DeleteDiscount Use this api to delete one discount activity
// Path: /api/v2/discount/delete_discount
// https://open.shopee.com/documents/v2/v2.discount.delete_discount?module=99&type=1
func (s *DiscountServiceOp[T]) DeleteDiscount(ctx context.Context, sid uint64, req DeleteDiscountRequest, tok string) (*DeleteDiscountResponse, error) {
	path := "/discount/delete_discount"
	resp := new(DeleteDiscountResponse)
	err := s.client.WithShop(sid, tok).Post(ctx, path, req, resp)
	return resp, err
}

// DeleteDiscountItem Use this api to delete items of the discount activity
// Path: /api/v2/discount/delete_discount_item
// https://open.shopee.com/documents/v2/v2.discount.delete_discount_item?module=99&type=1
func (s *DiscountServiceOp[T]) DeleteDiscountItem(ctx context.Context, sid uint64, req DeleteDiscountItemRequest, tok string) (*DeleteDiscountItemResponse, error) {
	path := "/discount/delete_discount_item"
	resp := new(DeleteDiscountItemResponse)
	err := s.client.WithShop(sid, tok).Post(ctx, path, req, resp)
	return resp, err
}

// DeleteSipDiscount Delete SIP Overseas Discounts for SIP affiliate region. Please use Primary shop's Shop ID to request, and provide the region of the Affiliate shop to be deleted, the API will delete the discount from that region's Affiliate shop.
// Path: /api/v2/discount/delete_sip_discount
// https://open.shopee.com/documents/v2/v2.discount.delete_sip_discount?module=99&type=1
func (s *DiscountServiceOp[T]) DeleteSipDiscount(ctx context.Context, sid uint64, req DeleteSipDiscountRequest, tok string) (*DeleteSipDiscountResponse, error) {
	path := "/discount/delete_sip_discount"
	resp := new(DeleteSipDiscountResponse)
	err := s.client.WithShop(sid, tok).Post(ctx, path, req, resp)
	return resp, err
}

// EndDiscount Use this api to end shop discount activity
// Path: /api/v2/discount/end_discount
// https://open.shopee.com/documents/v2/v2.discount.end_discount?module=99&type=1
func (s *DiscountServiceOp[T]) EndDiscount(ctx context.Context, sid uint64, req EndDiscountRequest, tok string) (*EndDiscountResponse, error) {
	path := "/discount/end_discount"
	resp := new(EndDiscountResponse)
	err := s.client.WithShop(sid, tok).Post(ctx, path, req, resp)
	return resp, err
}

// GetDiscount Use this api to get one shop discount activity detail
// Path: /api/v2/discount/get_discount
// https://open.shopee.com/documents/v2/v2.discount.get_discount?module=99&type=1
func (s *DiscountServiceOp[T]) GetDiscount(ctx context.Context, sid uint64, opt GetDiscountRequest, tok string) (*GetDiscountResponse, error) {
	path := "/discount/get_discount"
	resp := new(GetDiscountResponse)
	err := s.client.WithShop(sid, tok).Get(ctx, path, resp, opt)
	return resp, err
}

// GetDiscountList Use this api to get shop discount activity list
// Path: /api/v2/discount/get_discount_list
// https://open.shopee.com/documents/v2/v2.discount.get_discount_list?module=99&type=1
func (s *DiscountServiceOp[T]) GetDiscountList(ctx context.Context, sid uint64, req GetDiscountListRequest, tok string) (*GetDiscountListResponse, error) {
	path := "/discount/get_discount_list"
	resp := new(GetDiscountListResponse)
	err := s.client.WithShop(sid, tok).Post(ctx, path, req, resp)
	return resp, err
}

// GetSipDiscounts Get SIP Overseas Discounts. Only regions that have upcoming/ongoing discounts will be returned. Please use Primary shop's Shop ID to request, the API will return the list of Affiliate shops under this Primary shop that have set discounts, along with the discount details.
// Path: /api/v2/discount/get_sip_discounts
// https://open.shopee.com/documents/v2/v2.discount.get_sip_discounts?module=99&type=1
func (s *DiscountServiceOp[T]) GetSipDiscounts(ctx context.Context, sid uint64, opt GetSipDiscountsRequest, tok string) (*GetSipDiscountsResponse, error) {
	path := "/discount/get_sip_discounts"
	resp := new(GetSipDiscountsResponse)
	err := s.client.WithShop(sid, tok).Get(ctx, path, resp, opt)
	return resp, err
}

// SetSipDiscount Set SIP Overseas Discount for SIP affiliate region. Please use Primary shop's Shop ID to request, and provide the region and discount rate of the Affiliate shop to be set or update, the API will set or update the discount rate for that region's Affiliate shop.
// Path: /api/v2/discount/set_sip_discount
// https://open.shopee.com/documents/v2/v2.discount.set_sip_discount?module=99&type=1
func (s *DiscountServiceOp[T]) SetSipDiscount(ctx context.Context, sid uint64, req SetSipDiscountRequest, tok string) (*SetSipDiscountResponse, error) {
	path := "/discount/set_sip_discount"
	resp := new(SetSipDiscountResponse)
	err := s.client.WithShop(sid, tok).Post(ctx, path, req, resp)
	return resp, err
}

// UpdateDiscount Use this api to update one discount information
// Path: /api/v2/discount/update_discount
// https://open.shopee.com/documents/v2/v2.discount.update_discount?module=99&type=1
func (s *DiscountServiceOp[T]) UpdateDiscount(ctx context.Context, sid uint64, req UpdateDiscountRequest, tok string) (*UpdateDiscountResponse, error) {
	path := "/discount/update_discount"
	resp := new(UpdateDiscountResponse)
	err := s.client.WithShop(sid, tok).Post(ctx, path, req, resp)
	return resp, err
}

// UpdateDiscountItem Use this api to update items of the discount promotion.
// Path: /api/v2/discount/update_discount_item
// https://open.shopee.com/documents/v2/v2.discount.update_discount_item?module=99&type=1
func (s *DiscountServiceOp[T]) UpdateDiscountItem(ctx context.Context, sid uint64, req UpdateDiscountItemRequest, tok string) (*UpdateDiscountItemResponse, error) {
	path := "/discount/update_discount_item"
	resp := new(UpdateDiscountItemResponse)
	err := s.client.WithShop(sid, tok).Post(ctx, path, req, resp)
	return resp, err
}
