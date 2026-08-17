package goshopee

import (
	"context"
)

type SBSService interface {
	// GetBoundWhsInfo get bound warehouse by shop id
	// Path: /api/v2/sbs/get_bound_whs_info
	// https://open.shopee.com/documents/v2/v2.sbs.get_bound_whs_info?module=124&type=1
	GetBoundWhsInfo(ctx context.Context, sid uint64, mid uint64, tok string) (*GetBoundWhsInfoResponse, error)
	// GetCurrentInventory Get Seller Center Current Inventory Page Data
	// Path: /api/v2/sbs/get_current_inventory
	// https://open.shopee.com/documents/v2/v2.sbs.get_current_inventory?module=124&type=1
	GetCurrentInventory(ctx context.Context, sid uint64, mid uint64, tok string, req GetCurrentInventoryRequest) (*GetCurrentInventoryResponse, error)
	// GetExpiryReport Seller Center Expiry Report page data
	// Path: /api/v2/sbs/get_expiry_report
	// https://open.shopee.com/documents/v2/v2.sbs.get_expiry_report?module=124&type=1
	GetExpiryReport(ctx context.Context, sid uint64, mid uint64, tok string, req GetExpiryReportRequest) (*GetExpiryReportResponse, error)
	// GetFulfillmentMappingInventoryList {"content":"<p><span style=\"color:rgb(23,43,77);\">This API is designed for sellers using Fulfillment Mapping and their ERP systems.</span><br><span style=\"color:rgb(23,43,77);\">It allows callers to query the corresponding mapping and inventory information using the MTSKU ID of either a Bundle SKU or a Parent SKU, supporting automated inventory reconciliation and planning, improving Parent SKU inventory visibility, and reducing manual operations and cross-channel overselling risks.</span></p>","raw_content":[{"name":"paragraph","children":[{"attributes":{"fontColor":"rgb(23,43,77)"},"data":"This API is designed for sellers using Fulfillment Mapping and their ERP systems."},{"name":"softBreak"},{"attributes":{"fontColor":"rgb(23,43,77)"},"data":"It allows callers to query the corresponding mapping and inventory information using the MTSKU ID of either a Bundle SKU or a Parent SKU, supporting automated inventory reconciliation and planning, improving Parent SKU inventory visibility, and reducing manual operations and cross-channel overselling risks."}]}]}
	// Path: /api/v2/sbs/get_fulfillment_mapping_inventory_list
	// https://open.shopee.com/documents/v2/v2.sbs.get_fulfillment_mapping_inventory_list?module=124&type=1
	GetFulfillmentMappingInventoryList(ctx context.Context, sid uint64, mid uint64, tok string, opt GetFulfillmentMappingInventoryListRequest) (*GetFulfillmentMappingInventoryListResponse, error)
	// GetStockAging Get Seller Center Stock Aging page data
	// Path: /api/v2/sbs/get_stock_aging
	// https://open.shopee.com/documents/v2/v2.sbs.get_stock_aging?module=124&type=1
	GetStockAging(ctx context.Context, sid uint64, mid uint64, tok string, req GetStockAgingRequest) (*GetStockAgingResponse, error)
	// GetStockMovement Get Seller Center，Stock Movement page data
	// Path: /api/v2/sbs/get_stock_movement
	// https://open.shopee.com/documents/v2/v2.sbs.get_stock_movement?module=124&type=1
	GetStockMovement(ctx context.Context, sid uint64, mid uint64, tok string, req GetStockMovementRequest) (*GetStockMovementResponse, error)
}

type SBSServiceOp[T any] struct {
	client *Client[T]
}

// GetBoundWhsInfo get bound warehouse by shop id
// Path: /api/v2/sbs/get_bound_whs_info
// https://open.shopee.com/documents/v2/v2.sbs.get_bound_whs_info?module=124&type=1
func (s *SBSServiceOp[T]) GetBoundWhsInfo(ctx context.Context, sid uint64, mid uint64, tok string) (*GetBoundWhsInfoResponse, error) {
	path := "/sbs/get_bound_whs_info"
	resp := new(GetBoundWhsInfoResponse)
	err := s.client.Get(ctx, path, resp, nil, sid, mid, tok)
	return resp, err
}

// GetCurrentInventory Get Seller Center Current Inventory Page Data
// Path: /api/v2/sbs/get_current_inventory
// https://open.shopee.com/documents/v2/v2.sbs.get_current_inventory?module=124&type=1
func (s *SBSServiceOp[T]) GetCurrentInventory(ctx context.Context, sid uint64, mid uint64, tok string, req GetCurrentInventoryRequest) (*GetCurrentInventoryResponse, error) {
	path := "/sbs/get_current_inventory"
	resp := new(GetCurrentInventoryResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// GetExpiryReport Seller Center Expiry Report page data
// Path: /api/v2/sbs/get_expiry_report
// https://open.shopee.com/documents/v2/v2.sbs.get_expiry_report?module=124&type=1
func (s *SBSServiceOp[T]) GetExpiryReport(ctx context.Context, sid uint64, mid uint64, tok string, req GetExpiryReportRequest) (*GetExpiryReportResponse, error) {
	path := "/sbs/get_expiry_report"
	resp := new(GetExpiryReportResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// GetFulfillmentMappingInventoryList {"content":"<p><span style=\"color:rgb(23,43,77);\">This API is designed for sellers using Fulfillment Mapping and their ERP systems.</span><br><span style=\"color:rgb(23,43,77);\">It allows callers to query the corresponding mapping and inventory information using the MTSKU ID of either a Bundle SKU or a Parent SKU, supporting automated inventory reconciliation and planning, improving Parent SKU inventory visibility, and reducing manual operations and cross-channel overselling risks.</span></p>","raw_content":[{"name":"paragraph","children":[{"attributes":{"fontColor":"rgb(23,43,77)"},"data":"This API is designed for sellers using Fulfillment Mapping and their ERP systems."},{"name":"softBreak"},{"attributes":{"fontColor":"rgb(23,43,77)"},"data":"It allows callers to query the corresponding mapping and inventory information using the MTSKU ID of either a Bundle SKU or a Parent SKU, supporting automated inventory reconciliation and planning, improving Parent SKU inventory visibility, and reducing manual operations and cross-channel overselling risks."}]}]}
// Path: /api/v2/sbs/get_fulfillment_mapping_inventory_list
// https://open.shopee.com/documents/v2/v2.sbs.get_fulfillment_mapping_inventory_list?module=124&type=1
func (s *SBSServiceOp[T]) GetFulfillmentMappingInventoryList(ctx context.Context, sid uint64, mid uint64, tok string, opt GetFulfillmentMappingInventoryListRequest) (*GetFulfillmentMappingInventoryListResponse, error) {
	path := "/sbs/get_fulfillment_mapping_inventory_list"
	resp := new(GetFulfillmentMappingInventoryListResponse)
	err := s.client.Get(ctx, path, resp, opt, sid, mid, tok)
	return resp, err
}

// GetStockAging Get Seller Center Stock Aging page data
// Path: /api/v2/sbs/get_stock_aging
// https://open.shopee.com/documents/v2/v2.sbs.get_stock_aging?module=124&type=1
func (s *SBSServiceOp[T]) GetStockAging(ctx context.Context, sid uint64, mid uint64, tok string, req GetStockAgingRequest) (*GetStockAgingResponse, error) {
	path := "/sbs/get_stock_aging"
	resp := new(GetStockAgingResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// GetStockMovement Get Seller Center，Stock Movement page data
// Path: /api/v2/sbs/get_stock_movement
// https://open.shopee.com/documents/v2/v2.sbs.get_stock_movement?module=124&type=1
func (s *SBSServiceOp[T]) GetStockMovement(ctx context.Context, sid uint64, mid uint64, tok string, req GetStockMovementRequest) (*GetStockMovementResponse, error) {
	path := "/sbs/get_stock_movement"
	resp := new(GetStockMovementResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}
