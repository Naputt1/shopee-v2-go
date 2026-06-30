package goshopee

type VoucherService interface {
	// AddVoucher Add voucher
	// Path: /api/v2/voucher/add_voucher
	// https://open.shopee.com/documents/v2/v2.voucher.add_voucher?module=112&type=1
	AddVoucher(sid uint64, req AddVoucherRequest, tok string) (*AddVoucherResponse, error)
	// DeleteVoucher Delete voucher
	// Path: /api/v2/voucher/delete_voucher
	// https://open.shopee.com/documents/v2/v2.voucher.delete_voucher?module=112&type=1
	DeleteVoucher(sid uint64, req DeleteVoucherRequest, tok string) (*DeleteVoucherResponse, error)
	// EndVoucher End Voucher
	// Path: /api/v2/voucher/end_voucher
	// https://open.shopee.com/documents/v2/v2.voucher.end_voucher?module=112&type=1
	EndVoucher(sid uint64, req EndVoucherRequest, tok string) (*EndVoucherResponse, error)
	// GetVoucher Get Voucher Detail
	// Path: /api/v2/voucher/get_voucher
	// https://open.shopee.com/documents/v2/v2.voucher.get_voucher?module=112&type=1
	GetVoucher(sid uint64, req GetVoucherRequest, tok string) (*GetVoucherResponse, error)
	// GetVoucherList Get Voucher List
	// Path: /api/v2/voucher/get_voucher_list
	// https://open.shopee.com/documents/v2/v2.voucher.get_voucher_list?module=112&type=1
	GetVoucherList(sid uint64, req GetVoucherListRequest, tok string) (*GetVoucherListResponse, error)
	// UpdateVoucher Update voucher
	// Path: /api/v2/voucher/update_voucher
	// https://open.shopee.com/documents/v2/v2.voucher.update_voucher?module=112&type=1
	UpdateVoucher(sid uint64, req UpdateVoucherRequest, tok string) (*UpdateVoucherResponse, error)
}

type VoucherServiceOp[T any] struct {
	client *Client[T]
}

// AddVoucher Add voucher
// Path: /api/v2/voucher/add_voucher
// https://open.shopee.com/documents/v2/v2.voucher.add_voucher?module=112&type=1
func (s *VoucherServiceOp[T]) AddVoucher(sid uint64, req AddVoucherRequest, tok string) (*AddVoucherResponse, error) {
	path := "/voucher/add_voucher"
	resp := new(AddVoucherResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}
// DeleteVoucher Delete voucher
// Path: /api/v2/voucher/delete_voucher
// https://open.shopee.com/documents/v2/v2.voucher.delete_voucher?module=112&type=1
func (s *VoucherServiceOp[T]) DeleteVoucher(sid uint64, req DeleteVoucherRequest, tok string) (*DeleteVoucherResponse, error) {
	path := "/voucher/delete_voucher"
	resp := new(DeleteVoucherResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}
// EndVoucher End Voucher
// Path: /api/v2/voucher/end_voucher
// https://open.shopee.com/documents/v2/v2.voucher.end_voucher?module=112&type=1
func (s *VoucherServiceOp[T]) EndVoucher(sid uint64, req EndVoucherRequest, tok string) (*EndVoucherResponse, error) {
	path := "/voucher/end_voucher"
	resp := new(EndVoucherResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}
// GetVoucher Get Voucher Detail
// Path: /api/v2/voucher/get_voucher
// https://open.shopee.com/documents/v2/v2.voucher.get_voucher?module=112&type=1
func (s *VoucherServiceOp[T]) GetVoucher(sid uint64, req GetVoucherRequest, tok string) (*GetVoucherResponse, error) {
	path := "/voucher/get_voucher"
	resp := new(GetVoucherResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}
// GetVoucherList Get Voucher List
// Path: /api/v2/voucher/get_voucher_list
// https://open.shopee.com/documents/v2/v2.voucher.get_voucher_list?module=112&type=1
func (s *VoucherServiceOp[T]) GetVoucherList(sid uint64, req GetVoucherListRequest, tok string) (*GetVoucherListResponse, error) {
	path := "/voucher/get_voucher_list"
	resp := new(GetVoucherListResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}
// UpdateVoucher Update voucher
// Path: /api/v2/voucher/update_voucher
// https://open.shopee.com/documents/v2/v2.voucher.update_voucher?module=112&type=1
func (s *VoucherServiceOp[T]) UpdateVoucher(sid uint64, req UpdateVoucherRequest, tok string) (*UpdateVoucherResponse, error) {
	path := "/voucher/update_voucher"
	resp := new(UpdateVoucherResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}
