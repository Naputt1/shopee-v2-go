package goshopee

type ShopFlashSaleService interface {
	// AddShopFlashSaleItems add shop flash sale item
	// Path: /api/v2/shop_flash_sale/add_shop_flash_sale_items
	// https://open.shopee.com/documents/v2/v2.shop_flash_sale.add_shop_flash_sale_items?module=123&type=1
	AddShopFlashSaleItems(sid uint64, req AddShopFlashSaleItemsRequest, tok string) (*AddShopFlashSaleItemsResponse, error)
	// CreateShopFlashSale creat shop flash sale
	// Path: /api/v2/shop_flash_sale/create_shop_flash_sale
	// https://open.shopee.com/documents/v2/v2.shop_flash_sale.create_shop_flash_sale?module=123&type=1
	CreateShopFlashSale(sid uint64, req CreateShopFlashSaleRequest, tok string) (*CreateShopFlashSaleResponse, error)
	// DeleteShopFlashSale delete shop flash sale
	// Path: /api/v2/shop_flash_sale/delete_shop_flash_sale
	// https://open.shopee.com/documents/v2/v2.shop_flash_sale.delete_shop_flash_sale?module=123&type=1
	DeleteShopFlashSale(sid uint64, req DeleteShopFlashSaleRequest, tok string) (*DeleteShopFlashSaleResponse, error)
	// DeleteShopFlashSaleItems delete shop flash sale items
	// Path: /api/v2/shop_flash_sale/delete_shop_flash_sale_items
	// https://open.shopee.com/documents/v2/v2.shop_flash_sale.delete_shop_flash_sale_items?module=123&type=1
	DeleteShopFlashSaleItems(sid uint64, req DeleteShopFlashSaleItemsRequest, tok string) (*DeleteShopFlashSaleItemsResponse, error)
	// GetItemCriteria get shop flash sale item criteria
	// Path: /api/v2/shop_flash_sale/get_item_criteria
	// https://open.shopee.com/documents/v2/v2.shop_flash_sale.get_item_criteria?module=123&type=1
	GetItemCriteria(sid uint64, tok string) (*GetItemCriteriaResponse, error)
	// GetShopFlashSale get shop flash sale detail
	// Path: /api/v2/shop_flash_sale/get_shop_flash_sale
	// https://open.shopee.com/documents/v2/v2.shop_flash_sale.get_shop_flash_sale?module=123&type=1
	GetShopFlashSale(sid uint64, req GetShopFlashSaleRequest, tok string) (*GetShopFlashSaleResponse, error)
	// GetShopFlashSaleItems get shop flash sale items and item detail
	// Path: /api/v2/shop_flash_sale/get_shop_flash_sale_items
	// https://open.shopee.com/documents/v2/v2.shop_flash_sale.get_shop_flash_sale_items?module=123&type=1
	GetShopFlashSaleItems(sid uint64, req GetShopFlashSaleItemsRequest, tok string) (*GetShopFlashSaleItemsResponse, error)
	// GetShopFlashSaleList get shop flash sale list
	// Path: /api/v2/shop_flash_sale/get_shop_flash_sale_list
	// https://open.shopee.com/documents/v2/v2.shop_flash_sale.get_shop_flash_sale_list?module=123&type=1
	GetShopFlashSaleList(sid uint64, req GetShopFlashSaleListRequest, tok string) (*GetShopFlashSaleListResponse, error)
	// GetTimeSlotId get time slot id
	// Path: /api/v2/shop_flash_sale/get_time_slot_id
	// https://open.shopee.com/documents/v2/v2.shop_flash_sale.get_time_slot_id?module=123&type=1
	GetTimeSlotId(sid uint64, req GetTimeSlotIdRequest, tok string) (*GetTimeSlotIdResponse, error)
	// UpdateShopFlashSale edit shop flash sale(enable, disable)
	// Path: /api/v2/shop_flash_sale/update_shop_flash_sale
	// https://open.shopee.com/documents/v2/v2.shop_flash_sale.update_shop_flash_sale?module=123&type=1
	UpdateShopFlashSale(sid uint64, req UpdateShopFlashSaleRequest, tok string) (*UpdateShopFlashSaleResponse, error)
	// UpdateShopFlashSaleItems edit shop flash sale item, you can only edit the models in disbaled or enabled status
	// Path: /api/v2/shop_flash_sale/update_shop_flash_sale_items
	// https://open.shopee.com/documents/v2/v2.shop_flash_sale.update_shop_flash_sale_items?module=123&type=1
	UpdateShopFlashSaleItems(sid uint64, req UpdateShopFlashSaleItemsRequest, tok string) (*UpdateShopFlashSaleItemsResponse, error)
}

type ShopFlashSaleServiceOp[T any] struct {
	client *Client[T]
}

// AddShopFlashSaleItems add shop flash sale item
// Path: /api/v2/shop_flash_sale/add_shop_flash_sale_items
// https://open.shopee.com/documents/v2/v2.shop_flash_sale.add_shop_flash_sale_items?module=123&type=1
func (s *ShopFlashSaleServiceOp[T]) AddShopFlashSaleItems(sid uint64, req AddShopFlashSaleItemsRequest, tok string) (*AddShopFlashSaleItemsResponse, error) {
	path := "/shop_flash_sale/add_shop_flash_sale_items"
	resp := new(AddShopFlashSaleItemsResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}
// CreateShopFlashSale creat shop flash sale
// Path: /api/v2/shop_flash_sale/create_shop_flash_sale
// https://open.shopee.com/documents/v2/v2.shop_flash_sale.create_shop_flash_sale?module=123&type=1
func (s *ShopFlashSaleServiceOp[T]) CreateShopFlashSale(sid uint64, req CreateShopFlashSaleRequest, tok string) (*CreateShopFlashSaleResponse, error) {
	path := "/shop_flash_sale/create_shop_flash_sale"
	resp := new(CreateShopFlashSaleResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}
// DeleteShopFlashSale delete shop flash sale
// Path: /api/v2/shop_flash_sale/delete_shop_flash_sale
// https://open.shopee.com/documents/v2/v2.shop_flash_sale.delete_shop_flash_sale?module=123&type=1
func (s *ShopFlashSaleServiceOp[T]) DeleteShopFlashSale(sid uint64, req DeleteShopFlashSaleRequest, tok string) (*DeleteShopFlashSaleResponse, error) {
	path := "/shop_flash_sale/delete_shop_flash_sale"
	resp := new(DeleteShopFlashSaleResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}
// DeleteShopFlashSaleItems delete shop flash sale items
// Path: /api/v2/shop_flash_sale/delete_shop_flash_sale_items
// https://open.shopee.com/documents/v2/v2.shop_flash_sale.delete_shop_flash_sale_items?module=123&type=1
func (s *ShopFlashSaleServiceOp[T]) DeleteShopFlashSaleItems(sid uint64, req DeleteShopFlashSaleItemsRequest, tok string) (*DeleteShopFlashSaleItemsResponse, error) {
	path := "/shop_flash_sale/delete_shop_flash_sale_items"
	resp := new(DeleteShopFlashSaleItemsResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}
// GetItemCriteria get shop flash sale item criteria
// Path: /api/v2/shop_flash_sale/get_item_criteria
// https://open.shopee.com/documents/v2/v2.shop_flash_sale.get_item_criteria?module=123&type=1
func (s *ShopFlashSaleServiceOp[T]) GetItemCriteria(sid uint64, tok string) (*GetItemCriteriaResponse, error) {
	path := "/shop_flash_sale/get_item_criteria"
	resp := new(GetItemCriteriaResponse)
	err := s.client.WithShop(sid, tok).Post(path, nil, resp)
	return resp, err
}
// GetShopFlashSale get shop flash sale detail
// Path: /api/v2/shop_flash_sale/get_shop_flash_sale
// https://open.shopee.com/documents/v2/v2.shop_flash_sale.get_shop_flash_sale?module=123&type=1
func (s *ShopFlashSaleServiceOp[T]) GetShopFlashSale(sid uint64, req GetShopFlashSaleRequest, tok string) (*GetShopFlashSaleResponse, error) {
	path := "/shop_flash_sale/get_shop_flash_sale"
	resp := new(GetShopFlashSaleResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}
// GetShopFlashSaleItems get shop flash sale items and item detail
// Path: /api/v2/shop_flash_sale/get_shop_flash_sale_items
// https://open.shopee.com/documents/v2/v2.shop_flash_sale.get_shop_flash_sale_items?module=123&type=1
func (s *ShopFlashSaleServiceOp[T]) GetShopFlashSaleItems(sid uint64, req GetShopFlashSaleItemsRequest, tok string) (*GetShopFlashSaleItemsResponse, error) {
	path := "/shop_flash_sale/get_shop_flash_sale_items"
	resp := new(GetShopFlashSaleItemsResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}
// GetShopFlashSaleList get shop flash sale list
// Path: /api/v2/shop_flash_sale/get_shop_flash_sale_list
// https://open.shopee.com/documents/v2/v2.shop_flash_sale.get_shop_flash_sale_list?module=123&type=1
func (s *ShopFlashSaleServiceOp[T]) GetShopFlashSaleList(sid uint64, req GetShopFlashSaleListRequest, tok string) (*GetShopFlashSaleListResponse, error) {
	path := "/shop_flash_sale/get_shop_flash_sale_list"
	resp := new(GetShopFlashSaleListResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}
// GetTimeSlotId get time slot id
// Path: /api/v2/shop_flash_sale/get_time_slot_id
// https://open.shopee.com/documents/v2/v2.shop_flash_sale.get_time_slot_id?module=123&type=1
func (s *ShopFlashSaleServiceOp[T]) GetTimeSlotId(sid uint64, req GetTimeSlotIdRequest, tok string) (*GetTimeSlotIdResponse, error) {
	path := "/shop_flash_sale/get_time_slot_id"
	resp := new(GetTimeSlotIdResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}
// UpdateShopFlashSale edit shop flash sale(enable, disable)
// Path: /api/v2/shop_flash_sale/update_shop_flash_sale
// https://open.shopee.com/documents/v2/v2.shop_flash_sale.update_shop_flash_sale?module=123&type=1
func (s *ShopFlashSaleServiceOp[T]) UpdateShopFlashSale(sid uint64, req UpdateShopFlashSaleRequest, tok string) (*UpdateShopFlashSaleResponse, error) {
	path := "/shop_flash_sale/update_shop_flash_sale"
	resp := new(UpdateShopFlashSaleResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}
// UpdateShopFlashSaleItems edit shop flash sale item, you can only edit the models in disbaled or enabled status
// Path: /api/v2/shop_flash_sale/update_shop_flash_sale_items
// https://open.shopee.com/documents/v2/v2.shop_flash_sale.update_shop_flash_sale_items?module=123&type=1
func (s *ShopFlashSaleServiceOp[T]) UpdateShopFlashSaleItems(sid uint64, req UpdateShopFlashSaleItemsRequest, tok string) (*UpdateShopFlashSaleItemsResponse, error) {
	path := "/shop_flash_sale/update_shop_flash_sale_items"
	resp := new(UpdateShopFlashSaleItemsResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}
