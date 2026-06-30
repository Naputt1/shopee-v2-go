package goshopee

import (
	"context"
)

type FBSService interface {
	// QueryBrShopBlockStatus This API checks whether an FBS shop is blocked due to invoice-related issues. When blocked, the shop cannot create new Inbound Requests, and its warehouse inventory is restricted from being sold.
	// Path: /api/v2/fbs/query_br_shop_block_status
	// https://open.shopee.com/documents/v2/v2.fbs.query_br_shop_block_status?module=126&type=1
	QueryBrShopBlockStatus(ctx context.Context, sid uint64, tok string) (*QueryBrShopBlockStatusResponse, error)
	// QueryBrShopEnrollmentStatus This API checks whether a given shop_id is eligible to enroll in the Brazil Fulfilled-by-Shopee (FBS) service.
	// Path: /api/v2/fbs/query_br_shop_enrollment_status
	// https://open.shopee.com/documents/v2/v2.fbs.query_br_shop_enrollment_status?module=126&type=1
	QueryBrShopEnrollmentStatus(ctx context.Context, sid uint64, tok string) (*QueryBrShopEnrollmentStatusResponse, error)
	// QueryBrShopInvoiceError This API handles failed invoice issuance for FBS-related processes, covering Inbound Requests, RTS Requests, Sales Orders, and Move Transfer Orders.
	// Path: /api/v2/fbs/query_br_shop_invoice_error
	// https://open.shopee.com/documents/v2/v2.fbs.query_br_shop_invoice_error?module=126&type=1
	QueryBrShopInvoiceError(ctx context.Context, sid uint64, req QueryBrShopInvoiceErrorRequest, tok string) (*QueryBrShopInvoiceErrorResponse, error)
	// QueryBrSkuBlockStatus This API checks whether an FBS product is blocked due to invoice-related issues. When blocked, the product cannot be included in new Inbound Requests, and its warehouse inventory is restricted from being sold.
	// Path: /api/v2/fbs/query_br_sku_block_status
	// https://open.shopee.com/documents/v2/v2.fbs.query_br_sku_block_status?module=126&type=1
	QueryBrSkuBlockStatus(ctx context.Context, sid uint64, req QueryBrSkuBlockStatusRequest, tok string) (*QueryBrSkuBlockStatusResponse, error)
}

type FBSServiceOp[T any] struct {
	client *Client[T]
}

// QueryBrShopBlockStatus This API checks whether an FBS shop is blocked due to invoice-related issues. When blocked, the shop cannot create new Inbound Requests, and its warehouse inventory is restricted from being sold.
// Path: /api/v2/fbs/query_br_shop_block_status
// https://open.shopee.com/documents/v2/v2.fbs.query_br_shop_block_status?module=126&type=1
func (s *FBSServiceOp[T]) QueryBrShopBlockStatus(ctx context.Context, sid uint64, tok string) (*QueryBrShopBlockStatusResponse, error) {
	path := "/fbs/query_br_shop_block_status"
	resp := new(QueryBrShopBlockStatusResponse)
	err := s.client.Post(ctx, path, nil, resp, sid, tok)
	return resp, err
}

// QueryBrShopEnrollmentStatus This API checks whether a given shop_id is eligible to enroll in the Brazil Fulfilled-by-Shopee (FBS) service.
// Path: /api/v2/fbs/query_br_shop_enrollment_status
// https://open.shopee.com/documents/v2/v2.fbs.query_br_shop_enrollment_status?module=126&type=1
func (s *FBSServiceOp[T]) QueryBrShopEnrollmentStatus(ctx context.Context, sid uint64, tok string) (*QueryBrShopEnrollmentStatusResponse, error) {
	path := "/fbs/query_br_shop_enrollment_status"
	resp := new(QueryBrShopEnrollmentStatusResponse)
	err := s.client.Post(ctx, path, nil, resp, sid, tok)
	return resp, err
}

// QueryBrShopInvoiceError This API handles failed invoice issuance for FBS-related processes, covering Inbound Requests, RTS Requests, Sales Orders, and Move Transfer Orders.
// Path: /api/v2/fbs/query_br_shop_invoice_error
// https://open.shopee.com/documents/v2/v2.fbs.query_br_shop_invoice_error?module=126&type=1
func (s *FBSServiceOp[T]) QueryBrShopInvoiceError(ctx context.Context, sid uint64, req QueryBrShopInvoiceErrorRequest, tok string) (*QueryBrShopInvoiceErrorResponse, error) {
	path := "/fbs/query_br_shop_invoice_error"
	resp := new(QueryBrShopInvoiceErrorResponse)
	err := s.client.Post(ctx, path, req, resp, sid, tok)
	return resp, err
}

// QueryBrSkuBlockStatus This API checks whether an FBS product is blocked due to invoice-related issues. When blocked, the product cannot be included in new Inbound Requests, and its warehouse inventory is restricted from being sold.
// Path: /api/v2/fbs/query_br_sku_block_status
// https://open.shopee.com/documents/v2/v2.fbs.query_br_sku_block_status?module=126&type=1
func (s *FBSServiceOp[T]) QueryBrSkuBlockStatus(ctx context.Context, sid uint64, req QueryBrSkuBlockStatusRequest, tok string) (*QueryBrSkuBlockStatusResponse, error) {
	path := "/fbs/query_br_sku_block_status"
	resp := new(QueryBrSkuBlockStatusResponse)
	err := s.client.Post(ctx, path, req, resp, sid, tok)
	return resp, err
}
