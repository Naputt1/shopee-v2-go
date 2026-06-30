package goshopee

type MerchantService interface {
	// GetMerchantInfo Use this call to get information of merchant
	// Path: /api/v2/merchant/get_merchant_info
	// https://open.shopee.com/documents/v2/v2.merchant.get_merchant_info?module=93&type=1
	GetMerchantInfo(sid uint64, tok string) (*GetMerchantInfoResponse, error)
	// GetMerchantPrepaidAccountList Use this api to get seller’s courier prepaid account.
	// Path: /api/v2/merchant/get_merchant_prepaid_account_list
	// https://open.shopee.com/documents/v2/v2.merchant.get_merchant_prepaid_account_list?module=93&type=1
	GetMerchantPrepaidAccountList(sid uint64, opt GetMerchantPrepaidAccountListRequest, tok string) (*GetMerchantPrepaidAccountListResponse, error)
	// GetMerchantWarehouseList Get merchant warehouse with page
	// Path: /api/v2/merchant/get_merchant_warehouse_list
	// https://open.shopee.com/documents/v2/v2.merchant.get_merchant_warehouse_list?module=93&type=1
	GetMerchantWarehouseList(sid uint64, req GetMerchantWarehouseListRequest, tok string) (*GetMerchantWarehouseListResponse, error)
	// GetMerchantWarehouseLocationList get merchant warehouse location list
	// Path: /api/v2/merchant/get_merchant_warehouse_location_list
	// https://open.shopee.com/documents/v2/v2.merchant.get_merchant_warehouse_location_list?module=93&type=1
	GetMerchantWarehouseLocationList(sid uint64, tok string) (*GetMerchantWarehouseLocationListResponse, error)
	// GetShopListByMerchant Use this call to get shop_list bound to merchant_id.
	// 
	// 
	// Path: /api/v2/merchant/get_shop_list_by_merchant
	// https://open.shopee.com/documents/v2/v2.merchant.get_shop_list_by_merchant?module=93&type=1
	GetShopListByMerchant(sid uint64, opt GetShopListByMerchantRequest, tok string) (*GetShopListByMerchantResponse, error)
	// GetWarehouseEligibleShopList Get eligible shop list by warehouse id
	// Path: /api/v2/merchant/get_warehouse_eligible_shop_list
	// https://open.shopee.com/documents/v2/v2.merchant.get_warehouse_eligible_shop_list?module=93&type=1
	GetWarehouseEligibleShopList(sid uint64, req GetWarehouseEligibleShopListRequest, tok string) (*GetWarehouseEligibleShopListResponse, error)
}

type MerchantServiceOp[T any] struct {
	client *Client[T]
}

// GetMerchantInfo Use this call to get information of merchant
// Path: /api/v2/merchant/get_merchant_info
// https://open.shopee.com/documents/v2/v2.merchant.get_merchant_info?module=93&type=1
func (s *MerchantServiceOp[T]) GetMerchantInfo(sid uint64, tok string) (*GetMerchantInfoResponse, error) {
	path := "/merchant/get_merchant_info"
	resp := new(GetMerchantInfoResponse)
	err := s.client.WithMerchant(sid, tok).Get(path, resp, nil)
	return resp, err
}
// GetMerchantPrepaidAccountList Use this api to get seller’s courier prepaid account.
// Path: /api/v2/merchant/get_merchant_prepaid_account_list
// https://open.shopee.com/documents/v2/v2.merchant.get_merchant_prepaid_account_list?module=93&type=1
func (s *MerchantServiceOp[T]) GetMerchantPrepaidAccountList(sid uint64, opt GetMerchantPrepaidAccountListRequest, tok string) (*GetMerchantPrepaidAccountListResponse, error) {
	path := "/merchant/get_merchant_prepaid_account_list"
	resp := new(GetMerchantPrepaidAccountListResponse)
	err := s.client.WithMerchant(sid, tok).Get(path, resp, opt)
	return resp, err
}
// GetMerchantWarehouseList Get merchant warehouse with page
// Path: /api/v2/merchant/get_merchant_warehouse_list
// https://open.shopee.com/documents/v2/v2.merchant.get_merchant_warehouse_list?module=93&type=1
func (s *MerchantServiceOp[T]) GetMerchantWarehouseList(sid uint64, req GetMerchantWarehouseListRequest, tok string) (*GetMerchantWarehouseListResponse, error) {
	path := "/merchant/get_merchant_warehouse_list"
	resp := new(GetMerchantWarehouseListResponse)
	err := s.client.WithMerchant(sid, tok).Post(path, req, resp)
	return resp, err
}
// GetMerchantWarehouseLocationList get merchant warehouse location list
// Path: /api/v2/merchant/get_merchant_warehouse_location_list
// https://open.shopee.com/documents/v2/v2.merchant.get_merchant_warehouse_location_list?module=93&type=1
func (s *MerchantServiceOp[T]) GetMerchantWarehouseLocationList(sid uint64, tok string) (*GetMerchantWarehouseLocationListResponse, error) {
	path := "/merchant/get_merchant_warehouse_location_list"
	resp := new(GetMerchantWarehouseLocationListResponse)
	err := s.client.WithMerchant(sid, tok).Post(path, nil, resp)
	return resp, err
}
// GetShopListByMerchant Use this call to get shop_list bound to merchant_id.
// 
// 
// Path: /api/v2/merchant/get_shop_list_by_merchant
// https://open.shopee.com/documents/v2/v2.merchant.get_shop_list_by_merchant?module=93&type=1
func (s *MerchantServiceOp[T]) GetShopListByMerchant(sid uint64, opt GetShopListByMerchantRequest, tok string) (*GetShopListByMerchantResponse, error) {
	path := "/merchant/get_shop_list_by_merchant"
	resp := new(GetShopListByMerchantResponse)
	err := s.client.WithMerchant(sid, tok).Get(path, resp, opt)
	return resp, err
}
// GetWarehouseEligibleShopList Get eligible shop list by warehouse id
// Path: /api/v2/merchant/get_warehouse_eligible_shop_list
// https://open.shopee.com/documents/v2/v2.merchant.get_warehouse_eligible_shop_list?module=93&type=1
func (s *MerchantServiceOp[T]) GetWarehouseEligibleShopList(sid uint64, req GetWarehouseEligibleShopListRequest, tok string) (*GetWarehouseEligibleShopListResponse, error) {
	path := "/merchant/get_warehouse_eligible_shop_list"
	resp := new(GetWarehouseEligibleShopListResponse)
	err := s.client.WithMerchant(sid, tok).Post(path, req, resp)
	return resp, err
}
