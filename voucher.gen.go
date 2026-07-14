package goshopee

import (
	"context"
)

type VoucherService interface {
	// AddVoucher Add voucher
	// Path: /api/v2/voucher/add_voucher
	// https://open.shopee.com/documents/v2/v2.voucher.add_voucher?module=112&type=1
	AddVoucher(ctx context.Context, sid uint64, mid uint64, tok string, req AddVoucherRequest) (*AddVoucherResponse, error)
	// DeleteVoucher Delete voucher
	// Path: /api/v2/voucher/delete_voucher
	// https://open.shopee.com/documents/v2/v2.voucher.delete_voucher?module=112&type=1
	DeleteVoucher(ctx context.Context, sid uint64, mid uint64, tok string, req DeleteVoucherRequest) (*DeleteVoucherResponse, error)
	// EndVoucher End Voucher
	// Path: /api/v2/voucher/end_voucher
	// https://open.shopee.com/documents/v2/v2.voucher.end_voucher?module=112&type=1
	EndVoucher(ctx context.Context, sid uint64, mid uint64, tok string, req EndVoucherRequest) (*EndVoucherResponse, error)
	// GetVoucher Get Voucher Detail
	// Path: /api/v2/voucher/get_voucher
	// https://open.shopee.com/documents/v2/v2.voucher.get_voucher?module=112&type=1
	GetVoucher(ctx context.Context, sid uint64, mid uint64, tok string, req GetVoucherRequest) (*GetVoucherResponse, error)
	// GetVoucherList Get Voucher List
	// Path: /api/v2/voucher/get_voucher_list
	// https://open.shopee.com/documents/v2/v2.voucher.get_voucher_list?module=112&type=1
	GetVoucherList(ctx context.Context, sid uint64, mid uint64, tok string, req GetVoucherListRequest) (*GetVoucherListResponse, error)
	// UpdateVoucher Update voucher
	// Path: /api/v2/voucher/update_voucher
	// https://open.shopee.com/documents/v2/v2.voucher.update_voucher?module=112&type=1
	UpdateVoucher(ctx context.Context, sid uint64, mid uint64, tok string, req UpdateVoucherRequest) (*UpdateVoucherResponse, error)
}

type VoucherServiceOp[T any] struct {
	client *Client[T]
}

// AddVoucher Add voucher
// Path: /api/v2/voucher/add_voucher
// https://open.shopee.com/documents/v2/v2.voucher.add_voucher?module=112&type=1
func (s *VoucherServiceOp[T]) AddVoucher(ctx context.Context, sid uint64, mid uint64, tok string, req AddVoucherRequest) (*AddVoucherResponse, error) {
	path := "/voucher/add_voucher"
	resp := new(AddVoucherResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// DeleteVoucher Delete voucher
// Path: /api/v2/voucher/delete_voucher
// https://open.shopee.com/documents/v2/v2.voucher.delete_voucher?module=112&type=1
func (s *VoucherServiceOp[T]) DeleteVoucher(ctx context.Context, sid uint64, mid uint64, tok string, req DeleteVoucherRequest) (*DeleteVoucherResponse, error) {
	path := "/voucher/delete_voucher"
	resp := new(DeleteVoucherResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// EndVoucher End Voucher
// Path: /api/v2/voucher/end_voucher
// https://open.shopee.com/documents/v2/v2.voucher.end_voucher?module=112&type=1
func (s *VoucherServiceOp[T]) EndVoucher(ctx context.Context, sid uint64, mid uint64, tok string, req EndVoucherRequest) (*EndVoucherResponse, error) {
	path := "/voucher/end_voucher"
	resp := new(EndVoucherResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// GetVoucher Get Voucher Detail
// Path: /api/v2/voucher/get_voucher
// https://open.shopee.com/documents/v2/v2.voucher.get_voucher?module=112&type=1
func (s *VoucherServiceOp[T]) GetVoucher(ctx context.Context, sid uint64, mid uint64, tok string, req GetVoucherRequest) (*GetVoucherResponse, error) {
	path := "/voucher/get_voucher"
	resp := new(GetVoucherResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// GetVoucherList Get Voucher List
// Path: /api/v2/voucher/get_voucher_list
// https://open.shopee.com/documents/v2/v2.voucher.get_voucher_list?module=112&type=1
func (s *VoucherServiceOp[T]) GetVoucherList(ctx context.Context, sid uint64, mid uint64, tok string, req GetVoucherListRequest) (*GetVoucherListResponse, error) {
	path := "/voucher/get_voucher_list"
	resp := new(GetVoucherListResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// UpdateVoucher Update voucher
// Path: /api/v2/voucher/update_voucher
// https://open.shopee.com/documents/v2/v2.voucher.update_voucher?module=112&type=1
func (s *VoucherServiceOp[T]) UpdateVoucher(ctx context.Context, sid uint64, mid uint64, tok string, req UpdateVoucherRequest) (*UpdateVoucherResponse, error) {
	path := "/voucher/update_voucher"
	resp := new(UpdateVoucherResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}
