package goshopee

import (
	"context"
	"io"
)

type OrderService interface {
	// CancelOrder {"content":"<p>Use this api to cancel an order. This action can only be performed before the order has been shipped.</p>","raw_content":[{"name":"paragraph","children":[{"data":"Use this api to cancel an order. This action can only be performed before the order has been shipped."}]}]}
	// Path: /api/v2/order/cancel_order
	// https://open.shopee.com/documents/v2/v2.order.cancel_order?module=94&type=1
	CancelOrder(ctx context.Context, sid uint64, mid uint64, tok string, req CancelOrderRequest) (*CancelOrderResponse, error)
	// DownloadFbsInvoices This API allows you to download FBS invoices. To use this API, the client must first call v2.order.generate_fbs_invoices to create a new shipping document task, followed by calling v2.order.get_fbs_invoices_result to check the task status. The document can only be downloaded once the task status is "READY."
	// Path: /api/v2/order/download_fbs_invoices
	// https://open.shopee.com/documents/v2/v2.order.download_fbs_invoices?module=94&type=1
	DownloadFbsInvoices(ctx context.Context, sid uint64, mid uint64, tok string, req DownloadFbsInvoicesRequest) (*DownloadFbsInvoicesResponse, error)
	// DownloadInvoiceDoc This endpoint only for PH and BR local seller. Seller can download the invoice uploaded before through this endpoint.
	//
	// Path: /api/v2/order/download_invoice_doc
	// https://open.shopee.com/documents/v2/v2.order.download_invoice_doc?module=94&type=1
	DownloadInvoiceDoc(ctx context.Context, sid uint64, mid uint64, tok string, opt DownloadInvoiceDocRequest) (*DownloadInvoiceDocResponse, error)
	// GenerateFbsInvoices This API creates a task to download a specific tax document (e.g., sales invoice, remessa invoice) for the seller's account, available only after the document is issued by the system as part of the Fulfilled by Shopee (FBS) process.
	// The workflow is as follows: (1) v2.order.generate_fbs_invoices; (2) v2.order.get_fbs_invoices_result; (3) v2.order.download_fbs_invoices.
	// Please note: The download link for the document will expire 30 minutes after being generated.
	// Path: /api/v2/order/generate_fbs_invoices
	// https://open.shopee.com/documents/v2/v2.order.generate_fbs_invoices?module=94&type=1
	GenerateFbsInvoices(ctx context.Context, sid uint64, mid uint64, tok string, req GenerateFbsInvoicesRequest) (*GenerateFbsInvoicesResponse, error)
	// GetBookingDetail Use this api to get booking detail.
	// Path: /api/v2/order/get_booking_detail
	// https://open.shopee.com/documents/v2/v2.order.get_booking_detail?module=94&type=1
	GetBookingDetail(ctx context.Context, sid uint64, mid uint64, tok string, opt GetBookingDetailRequest) (*GetBookingDetailResponse, error)
	// GetBookingList Use this api to search bookings. You may also filter them by status, if needed.
	// Path: /api/v2/order/get_booking_list
	// https://open.shopee.com/documents/v2/v2.order.get_booking_list?module=94&type=1
	GetBookingList(ctx context.Context, sid uint64, mid uint64, tok string, opt GetBookingListRequest) (*GetBookingListResponse, error)
	// GetBuyerInvoiceInfo {"content":"<p>API to obtain buyer submitted invoice info for VN, TH and PH local sellers only.</p>","raw_content":[{"name":"paragraph","children":[{"data":"API to obtain buyer submitted invoice info for VN, TH and PH local sellers only."}]}]}
	// Path: /api/v2/order/get_buyer_invoice_info
	// https://open.shopee.com/documents/v2/v2.order.get_buyer_invoice_info?module=94&type=1
	GetBuyerInvoiceInfo(ctx context.Context, sid uint64, mid uint64, tok string, req GetBuyerInvoiceInfoRequest) (*GetBuyerInvoiceInfoResponse, error)
	// GetEstimateCancelValue {"content":"<p>Returns the estimated refund value for a partial order cancellation given the specified items to cancel.</p>","raw_content":[{"name":"paragraph","children":[{"data":"Returns the estimated refund value for a partial order cancellation given the specified items to cancel."}]}]}
	// Path: /api/v2/order/get_estimate_cancel_value
	// https://open.shopee.com/documents/v2/v2.order.get_estimate_cancel_value?module=94&type=1
	GetEstimateCancelValue(ctx context.Context, sid uint64, mid uint64, tok string, req GetEstimateCancelValueRequest) (*GetEstimateCancelValueResponse, error)
	// GetFbsInvoicesResult This API allows you to consult the status of a previously requested batch download for FBS tax documents.
	// Path: /api/v2/order/get_fbs_invoices_result
	// https://open.shopee.com/documents/v2/v2.order.get_fbs_invoices_result?module=94&type=1
	GetFbsInvoicesResult(ctx context.Context, sid uint64, mid uint64, tok string, req GetFbsInvoicesResultRequest) (*GetFbsInvoicesResultResponse, error)
	// GetOrderDetail {"content":"<p>Use this api to get order detail.</p>","raw_content":[{"name":"paragraph","children":[{"data":"Use this api to get order detail."}]}]}
	// Path: /api/v2/order/get_order_detail
	// https://open.shopee.com/documents/v2/v2.order.get_order_detail?module=94&type=1
	GetOrderDetail(ctx context.Context, sid uint64, mid uint64, tok string, opt GetOrderDetailRequest) (*GetOrderDetailResponse, error)
	// GetOrderList Use this api to search orders. You may also filter them by status, if needed.
	// Path: /api/v2/order/get_order_list
	// https://open.shopee.com/documents/v2/v2.order.get_order_list?module=94&type=1
	GetOrderList(ctx context.Context, sid uint64, mid uint64, tok string, opt GetOrderListRequest) (*GetOrderListResponse, error)
	// GetPackageDetail {"content":"<p>Use this api to get package detail.<br>&nbsp;</p>","raw_content":[{"name":"paragraph","children":[{"data":"Use this api to get package detail."},{"name":"softBreak"},{"data":" "}]}]}
	// Path: /api/v2/order/get_package_detail
	// https://open.shopee.com/documents/v2/v2.order.get_package_detail?module=94&type=1
	GetPackageDetail(ctx context.Context, sid uint64, mid uint64, tok string, opt GetPackageDetailRequest) (*GetPackageDetailResponse, error)
	// GetPendingBuyerInvoiceOrderList This endpoint only for PH and BR local sellers only. This API is used for seller to retrieve a list of order IDs that are pending invoice upload.
	// Path: /api/v2/order/get_pending_buyer_invoice_order_list
	// https://open.shopee.com/documents/v2/v2.order.get_pending_buyer_invoice_order_list?module=94&type=1
	GetPendingBuyerInvoiceOrderList(ctx context.Context, sid uint64, mid uint64, tok string, opt GetPendingBuyerInvoiceOrderListRequest) (*GetPendingBuyerInvoiceOrderListResponse, error)
	// GetShipmentList Use this api to get order list which order_status is READY_TO_SHIP or RETRY_SHIP to start process the whole shipping progress.
	// Path: /api/v2/order/get_shipment_list
	// https://open.shopee.com/documents/v2/v2.order.get_shipment_list?module=94&type=1
	GetShipmentList(ctx context.Context, sid uint64, mid uint64, tok string, opt GetShipmentListRequest) (*GetShipmentListResponse, error)
	// GetWarehouseFilterConfig For multi-warehouse shops, return all warehouses with packages that have not been SHIPPED including product_location_id and address_id. Compared to v2.shop.get_warehouse_detail, it covers some edge cases like warehouse that have been unlinked but still retain packages that have not been SHIPPED, and does not cover some cases like single warehouse with default product_location_id and FBS shop.
	// Path: /api/v2/order/get_warehouse_filter_config
	// https://open.shopee.com/documents/v2/v2.order.get_warehouse_filter_config?module=94&type=1
	GetWarehouseFilterConfig(ctx context.Context, sid uint64, mid uint64, tok string) (*GetWarehouseFilterConfigResponse, error)
	// HandleBuyerCancellation Use this api to handle buyer's cancellation application.
	// Path: /api/v2/order/handle_buyer_cancellation
	// https://open.shopee.com/documents/v2/v2.order.handle_buyer_cancellation?module=94&type=1
	HandleBuyerCancellation(ctx context.Context, sid uint64, mid uint64, tok string, req HandleBuyerCancellationRequest) (*HandleBuyerCancellationResponse, error)
	// HandlePrescriptionCheck Use this API to approve or reject a prescription
	// Path: /api/v2/order/handle_prescription_check
	// https://open.shopee.com/documents/v2/v2.order.handle_prescription_check?module=94&type=1
	HandlePrescriptionCheck(ctx context.Context, sid uint64, mid uint64, tok string, req HandlePrescriptionCheckRequest) (*HandlePrescriptionCheckResponse, error)
	// SearchPackageList Use this API to search the list of packages that have not been SHIPPED to proceed arranging shipment, and it supports various filters and sort fields.
	// Path: /api/v2/order/search_package_list
	// https://open.shopee.com/documents/v2/v2.order.search_package_list?module=94&type=1
	SearchPackageList(ctx context.Context, sid uint64, mid uint64, tok string, req SearchPackageListRequest) (*SearchPackageListResponse, error)
	// SetNote Use this api to set note for an order.
	// Path: /api/v2/order/set_note
	// https://open.shopee.com/documents/v2/v2.order.set_note?module=94&type=1
	SetNote(ctx context.Context, sid uint64, mid uint64, tok string, req SetNoteRequest) (*SetNoteResponse, error)
	// SplitOrder Use this api to split an order into multiple packages. Orders that include installation services cannot be split by quantity.
	// Path: /api/v2/order/split_order
	// https://open.shopee.com/documents/v2/v2.order.split_order?module=94&type=1
	SplitOrder(ctx context.Context, sid uint64, mid uint64, tok string, req SplitOrderRequest) (*SplitOrderResponse, error)
	// UnsplitOrder Use this ai to undo split of order. After undo split, the order will have only one package. It can only be used when order status still at READY_TO_SHIP.
	// Path: /api/v2/order/unsplit_order
	// https://open.shopee.com/documents/v2/v2.order.unsplit_order?module=94&type=1
	UnsplitOrder(ctx context.Context, sid uint64, mid uint64, tok string, req UnsplitOrderRequest) (*UnsplitOrderResponse, error)
	// UploadInvoiceDoc This endpoint is for PH and BR local seller. Upload the invoice document
	//
	// Path: /api/v2/order/upload_invoice_doc
	// https://open.shopee.com/documents/v2/v2.order.upload_invoice_doc?module=94&type=1
	UploadInvoiceDoc(ctx context.Context, sid uint64, mid uint64, tok string, filename string) (*UploadInvoiceDocResponse, error)
	UploadInvoiceDocFromReader(ctx context.Context, sid uint64, mid uint64, tok string, filename string, reader io.Reader) (*UploadInvoiceDocResponse, error)
}

type OrderServiceOp[T any] struct {
	client *Client[T]
}

// CancelOrder {"content":"<p>Use this api to cancel an order. This action can only be performed before the order has been shipped.</p>","raw_content":[{"name":"paragraph","children":[{"data":"Use this api to cancel an order. This action can only be performed before the order has been shipped."}]}]}
// Path: /api/v2/order/cancel_order
// https://open.shopee.com/documents/v2/v2.order.cancel_order?module=94&type=1
func (s *OrderServiceOp[T]) CancelOrder(ctx context.Context, sid uint64, mid uint64, tok string, req CancelOrderRequest) (*CancelOrderResponse, error) {
	path := "/order/cancel_order"
	resp := new(CancelOrderResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// DownloadFbsInvoices This API allows you to download FBS invoices. To use this API, the client must first call v2.order.generate_fbs_invoices to create a new shipping document task, followed by calling v2.order.get_fbs_invoices_result to check the task status. The document can only be downloaded once the task status is "READY."
// Path: /api/v2/order/download_fbs_invoices
// https://open.shopee.com/documents/v2/v2.order.download_fbs_invoices?module=94&type=1
func (s *OrderServiceOp[T]) DownloadFbsInvoices(ctx context.Context, sid uint64, mid uint64, tok string, req DownloadFbsInvoicesRequest) (*DownloadFbsInvoicesResponse, error) {
	path := "/order/download_fbs_invoices"
	resp := new(DownloadFbsInvoicesResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// DownloadInvoiceDoc This endpoint only for PH and BR local seller. Seller can download the invoice uploaded before through this endpoint.
//
// Path: /api/v2/order/download_invoice_doc
// https://open.shopee.com/documents/v2/v2.order.download_invoice_doc?module=94&type=1
func (s *OrderServiceOp[T]) DownloadInvoiceDoc(ctx context.Context, sid uint64, mid uint64, tok string, opt DownloadInvoiceDocRequest) (*DownloadInvoiceDocResponse, error) {
	path := "/order/download_invoice_doc"
	resp := new(DownloadInvoiceDocResponse)
	err := s.client.Get(ctx, path, resp, opt, sid, mid, tok)
	return resp, err
}

// GenerateFbsInvoices This API creates a task to download a specific tax document (e.g., sales invoice, remessa invoice) for the seller's account, available only after the document is issued by the system as part of the Fulfilled by Shopee (FBS) process.
// The workflow is as follows: (1) v2.order.generate_fbs_invoices; (2) v2.order.get_fbs_invoices_result; (3) v2.order.download_fbs_invoices.
// Please note: The download link for the document will expire 30 minutes after being generated.
// Path: /api/v2/order/generate_fbs_invoices
// https://open.shopee.com/documents/v2/v2.order.generate_fbs_invoices?module=94&type=1
func (s *OrderServiceOp[T]) GenerateFbsInvoices(ctx context.Context, sid uint64, mid uint64, tok string, req GenerateFbsInvoicesRequest) (*GenerateFbsInvoicesResponse, error) {
	path := "/order/generate_fbs_invoices"
	resp := new(GenerateFbsInvoicesResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// GetBookingDetail Use this api to get booking detail.
// Path: /api/v2/order/get_booking_detail
// https://open.shopee.com/documents/v2/v2.order.get_booking_detail?module=94&type=1
func (s *OrderServiceOp[T]) GetBookingDetail(ctx context.Context, sid uint64, mid uint64, tok string, opt GetBookingDetailRequest) (*GetBookingDetailResponse, error) {
	path := "/order/get_booking_detail"
	resp := new(GetBookingDetailResponse)
	err := s.client.Get(ctx, path, resp, opt, sid, mid, tok)
	return resp, err
}

// GetBookingList Use this api to search bookings. You may also filter them by status, if needed.
// Path: /api/v2/order/get_booking_list
// https://open.shopee.com/documents/v2/v2.order.get_booking_list?module=94&type=1
func (s *OrderServiceOp[T]) GetBookingList(ctx context.Context, sid uint64, mid uint64, tok string, opt GetBookingListRequest) (*GetBookingListResponse, error) {
	path := "/order/get_booking_list"
	resp := new(GetBookingListResponse)
	err := s.client.Get(ctx, path, resp, opt, sid, mid, tok)
	return resp, err
}

// GetBuyerInvoiceInfo {"content":"<p>API to obtain buyer submitted invoice info for VN, TH and PH local sellers only.</p>","raw_content":[{"name":"paragraph","children":[{"data":"API to obtain buyer submitted invoice info for VN, TH and PH local sellers only."}]}]}
// Path: /api/v2/order/get_buyer_invoice_info
// https://open.shopee.com/documents/v2/v2.order.get_buyer_invoice_info?module=94&type=1
func (s *OrderServiceOp[T]) GetBuyerInvoiceInfo(ctx context.Context, sid uint64, mid uint64, tok string, req GetBuyerInvoiceInfoRequest) (*GetBuyerInvoiceInfoResponse, error) {
	path := "/order/get_buyer_invoice_info"
	resp := new(GetBuyerInvoiceInfoResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// GetEstimateCancelValue {"content":"<p>Returns the estimated refund value for a partial order cancellation given the specified items to cancel.</p>","raw_content":[{"name":"paragraph","children":[{"data":"Returns the estimated refund value for a partial order cancellation given the specified items to cancel."}]}]}
// Path: /api/v2/order/get_estimate_cancel_value
// https://open.shopee.com/documents/v2/v2.order.get_estimate_cancel_value?module=94&type=1
func (s *OrderServiceOp[T]) GetEstimateCancelValue(ctx context.Context, sid uint64, mid uint64, tok string, req GetEstimateCancelValueRequest) (*GetEstimateCancelValueResponse, error) {
	path := "/order/get_estimate_cancel_value"
	resp := new(GetEstimateCancelValueResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// GetFbsInvoicesResult This API allows you to consult the status of a previously requested batch download for FBS tax documents.
// Path: /api/v2/order/get_fbs_invoices_result
// https://open.shopee.com/documents/v2/v2.order.get_fbs_invoices_result?module=94&type=1
func (s *OrderServiceOp[T]) GetFbsInvoicesResult(ctx context.Context, sid uint64, mid uint64, tok string, req GetFbsInvoicesResultRequest) (*GetFbsInvoicesResultResponse, error) {
	path := "/order/get_fbs_invoices_result"
	resp := new(GetFbsInvoicesResultResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// GetOrderDetail {"content":"<p>Use this api to get order detail.</p>","raw_content":[{"name":"paragraph","children":[{"data":"Use this api to get order detail."}]}]}
// Path: /api/v2/order/get_order_detail
// https://open.shopee.com/documents/v2/v2.order.get_order_detail?module=94&type=1
func (s *OrderServiceOp[T]) GetOrderDetail(ctx context.Context, sid uint64, mid uint64, tok string, opt GetOrderDetailRequest) (*GetOrderDetailResponse, error) {
	path := "/order/get_order_detail"
	resp := new(GetOrderDetailResponse)
	err := s.client.Get(ctx, path, resp, opt, sid, mid, tok)
	return resp, err
}

// GetOrderList Use this api to search orders. You may also filter them by status, if needed.
// Path: /api/v2/order/get_order_list
// https://open.shopee.com/documents/v2/v2.order.get_order_list?module=94&type=1
func (s *OrderServiceOp[T]) GetOrderList(ctx context.Context, sid uint64, mid uint64, tok string, opt GetOrderListRequest) (*GetOrderListResponse, error) {
	path := "/order/get_order_list"
	resp := new(GetOrderListResponse)
	err := s.client.Get(ctx, path, resp, opt, sid, mid, tok)
	return resp, err
}

// GetPackageDetail {"content":"<p>Use this api to get package detail.<br>&nbsp;</p>","raw_content":[{"name":"paragraph","children":[{"data":"Use this api to get package detail."},{"name":"softBreak"},{"data":" "}]}]}
// Path: /api/v2/order/get_package_detail
// https://open.shopee.com/documents/v2/v2.order.get_package_detail?module=94&type=1
func (s *OrderServiceOp[T]) GetPackageDetail(ctx context.Context, sid uint64, mid uint64, tok string, opt GetPackageDetailRequest) (*GetPackageDetailResponse, error) {
	path := "/order/get_package_detail"
	resp := new(GetPackageDetailResponse)
	err := s.client.Get(ctx, path, resp, opt, sid, mid, tok)
	return resp, err
}

// GetPendingBuyerInvoiceOrderList This endpoint only for PH and BR local sellers only. This API is used for seller to retrieve a list of order IDs that are pending invoice upload.
// Path: /api/v2/order/get_pending_buyer_invoice_order_list
// https://open.shopee.com/documents/v2/v2.order.get_pending_buyer_invoice_order_list?module=94&type=1
func (s *OrderServiceOp[T]) GetPendingBuyerInvoiceOrderList(ctx context.Context, sid uint64, mid uint64, tok string, opt GetPendingBuyerInvoiceOrderListRequest) (*GetPendingBuyerInvoiceOrderListResponse, error) {
	path := "/order/get_pending_buyer_invoice_order_list"
	resp := new(GetPendingBuyerInvoiceOrderListResponse)
	err := s.client.Get(ctx, path, resp, opt, sid, mid, tok)
	return resp, err
}

// GetShipmentList Use this api to get order list which order_status is READY_TO_SHIP or RETRY_SHIP to start process the whole shipping progress.
// Path: /api/v2/order/get_shipment_list
// https://open.shopee.com/documents/v2/v2.order.get_shipment_list?module=94&type=1
func (s *OrderServiceOp[T]) GetShipmentList(ctx context.Context, sid uint64, mid uint64, tok string, opt GetShipmentListRequest) (*GetShipmentListResponse, error) {
	path := "/order/get_shipment_list"
	resp := new(GetShipmentListResponse)
	err := s.client.Get(ctx, path, resp, opt, sid, mid, tok)
	return resp, err
}

// GetWarehouseFilterConfig For multi-warehouse shops, return all warehouses with packages that have not been SHIPPED including product_location_id and address_id. Compared to v2.shop.get_warehouse_detail, it covers some edge cases like warehouse that have been unlinked but still retain packages that have not been SHIPPED, and does not cover some cases like single warehouse with default product_location_id and FBS shop.
// Path: /api/v2/order/get_warehouse_filter_config
// https://open.shopee.com/documents/v2/v2.order.get_warehouse_filter_config?module=94&type=1
func (s *OrderServiceOp[T]) GetWarehouseFilterConfig(ctx context.Context, sid uint64, mid uint64, tok string) (*GetWarehouseFilterConfigResponse, error) {
	path := "/order/get_warehouse_filter_config"
	resp := new(GetWarehouseFilterConfigResponse)
	err := s.client.Get(ctx, path, resp, nil, sid, mid, tok)
	return resp, err
}

// HandleBuyerCancellation Use this api to handle buyer's cancellation application.
// Path: /api/v2/order/handle_buyer_cancellation
// https://open.shopee.com/documents/v2/v2.order.handle_buyer_cancellation?module=94&type=1
func (s *OrderServiceOp[T]) HandleBuyerCancellation(ctx context.Context, sid uint64, mid uint64, tok string, req HandleBuyerCancellationRequest) (*HandleBuyerCancellationResponse, error) {
	path := "/order/handle_buyer_cancellation"
	resp := new(HandleBuyerCancellationResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// HandlePrescriptionCheck Use this API to approve or reject a prescription
// Path: /api/v2/order/handle_prescription_check
// https://open.shopee.com/documents/v2/v2.order.handle_prescription_check?module=94&type=1
func (s *OrderServiceOp[T]) HandlePrescriptionCheck(ctx context.Context, sid uint64, mid uint64, tok string, req HandlePrescriptionCheckRequest) (*HandlePrescriptionCheckResponse, error) {
	path := "/order/handle_prescription_check"
	resp := new(HandlePrescriptionCheckResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// SearchPackageList Use this API to search the list of packages that have not been SHIPPED to proceed arranging shipment, and it supports various filters and sort fields.
// Path: /api/v2/order/search_package_list
// https://open.shopee.com/documents/v2/v2.order.search_package_list?module=94&type=1
func (s *OrderServiceOp[T]) SearchPackageList(ctx context.Context, sid uint64, mid uint64, tok string, req SearchPackageListRequest) (*SearchPackageListResponse, error) {
	path := "/order/search_package_list"
	resp := new(SearchPackageListResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// SetNote Use this api to set note for an order.
// Path: /api/v2/order/set_note
// https://open.shopee.com/documents/v2/v2.order.set_note?module=94&type=1
func (s *OrderServiceOp[T]) SetNote(ctx context.Context, sid uint64, mid uint64, tok string, req SetNoteRequest) (*SetNoteResponse, error) {
	path := "/order/set_note"
	resp := new(SetNoteResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// SplitOrder Use this api to split an order into multiple packages. Orders that include installation services cannot be split by quantity.
// Path: /api/v2/order/split_order
// https://open.shopee.com/documents/v2/v2.order.split_order?module=94&type=1
func (s *OrderServiceOp[T]) SplitOrder(ctx context.Context, sid uint64, mid uint64, tok string, req SplitOrderRequest) (*SplitOrderResponse, error) {
	path := "/order/split_order"
	resp := new(SplitOrderResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// UnsplitOrder Use this ai to undo split of order. After undo split, the order will have only one package. It can only be used when order status still at READY_TO_SHIP.
// Path: /api/v2/order/unsplit_order
// https://open.shopee.com/documents/v2/v2.order.unsplit_order?module=94&type=1
func (s *OrderServiceOp[T]) UnsplitOrder(ctx context.Context, sid uint64, mid uint64, tok string, req UnsplitOrderRequest) (*UnsplitOrderResponse, error) {
	path := "/order/unsplit_order"
	resp := new(UnsplitOrderResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// UploadInvoiceDoc This endpoint is for PH and BR local seller. Upload the invoice document
//
// Path: /api/v2/order/upload_invoice_doc
// https://open.shopee.com/documents/v2/v2.order.upload_invoice_doc?module=94&type=1
func (s *OrderServiceOp[T]) UploadInvoiceDoc(ctx context.Context, sid uint64, mid uint64, tok string, filename string) (*UploadInvoiceDocResponse, error) {
	path := "/order/upload_invoice_doc"
	resp := new(UploadInvoiceDocResponse)
	err := s.client.Upload(ctx, path, "image", filename, resp, sid, mid, tok)
	return resp, err
}

func (s *OrderServiceOp[T]) UploadInvoiceDocFromReader(ctx context.Context, sid uint64, mid uint64, tok string, filename string, reader io.Reader) (*UploadInvoiceDocResponse, error) {
	path := "/order/upload_invoice_doc"
	resp := new(UploadInvoiceDocResponse)
	err := s.client.UploadFromReader(ctx, path, "image", filename, reader, resp, sid, mid, tok)
	return resp, err
}
