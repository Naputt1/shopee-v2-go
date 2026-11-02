package goshopee

import (
	"context"
)

type FirstMileService interface {
	// BindCourierDeliveryFirstMileTrackingNumber Use this api to bind first mile tracking number for courier delivery method.
	// Path: /api/v2/first_mile/bind_courier_delivery_first_mile_tracking_number
	// https://open.shopee.com/documents/v2/v2.first_mile.bind_courier_delivery_first_mile_tracking_number?module=96&type=1
	BindCourierDeliveryFirstMileTrackingNumber(ctx context.Context, sid uint64, mid uint64, tok string, req BindCourierDeliveryFirstMileTrackingNumberRequest) (*BindCourierDeliveryFirstMileTrackingNumberResponse, error)
	// BindFirstMileTrackingNumber Use this api to bind first mile tracking number.
	// Path: /api/v2/first_mile/bind_first_mile_tracking_number
	// https://open.shopee.com/documents/v2/v2.first_mile.bind_first_mile_tracking_number?module=96&type=1
	BindFirstMileTrackingNumber(ctx context.Context, sid uint64, mid uint64, tok string, req BindFirstMileTrackingNumberRequest) (*BindFirstMileTrackingNumberResponse, error)
	// CancelCourierDeliveryFirstMileTrackingNumber {"content":"<p>Use this api to cancel first mile tracking number for courier delivery method.</p>","raw_content":[{"name":"paragraph","children":[{"data":"Use this api to cancel first mile tracking number for courier delivery method."}]}]}
	// Path: /api/v2/first_mile/cancel_courier_delivery_first_mile_tracking_number
	// https://open.shopee.com/documents/v2/v2.first_mile.cancel_courier_delivery_first_mile_tracking_number?module=96&type=1
	CancelCourierDeliveryFirstMileTrackingNumber(ctx context.Context, sid uint64, mid uint64, tok string, req CancelCourierDeliveryFirstMileTrackingNumberRequest) (*CancelCourierDeliveryFirstMileTrackingNumberResponse, error)
	// GenerateAndBindFirstMileTrackingNumber Use this api to generate first mile tracking number for courier delivery method.
	// Path: /api/v2/first_mile/generate_and_bind_first_mile_tracking_number
	// https://open.shopee.com/documents/v2/v2.first_mile.generate_and_bind_first_mile_tracking_number?module=96&type=1
	GenerateAndBindFirstMileTrackingNumber(ctx context.Context, sid uint64, mid uint64, tok string, req GenerateAndBindFirstMileTrackingNumberRequest) (*GenerateAndBindFirstMileTrackingNumberResponse, error)
	// GenerateFirstMileTrackingNumber Use this api to generate first mile tracking number.
	// Path: /api/v2/first_mile/generate_first_mile_tracking_number
	// https://open.shopee.com/documents/v2/v2.first_mile.generate_first_mile_tracking_number?module=96&type=1
	GenerateFirstMileTrackingNumber(ctx context.Context, sid uint64, mid uint64, tok string, req GenerateFirstMileTrackingNumberRequest) (*GenerateFirstMileTrackingNumberResponse, error)
	// GetChannelList Use this api to get first mile channel list.
	// Path: /api/v2/first_mile/get_channel_list
	// https://open.shopee.com/documents/v2/v2.first_mile.get_channel_list?module=96&type=1
	GetChannelList(ctx context.Context, sid uint64, mid uint64, tok string, req GetChannelListRequest) (*GetChannelListResponse, error)
	// GetCourierDeliveryChannelList {"content":"<p>Use this api to get courier information for courier delivery method.<br>&nbsp;</p>","raw_content":[{"name":"paragraph","children":[{"data":"Use this api to get courier information for courier delivery method."},{"name":"softBreak"},{"data":" "}]}]}
	// Path: /api/v2/first_mile/get_courier_delivery_channel_list
	// https://open.shopee.com/documents/v2/v2.first_mile.get_courier_delivery_channel_list?module=96&type=1
	GetCourierDeliveryChannelList(ctx context.Context, sid uint64, mid uint64, tok string, opt GetCourierDeliveryChannelListRequest) (*GetCourierDeliveryChannelListResponse, error)
	// GetCourierDeliveryDetail Use this api to get first mile detail for courier delivery method.
	//
	// Path: /api/v2/first_mile/get_courier_delivery_detail
	// https://open.shopee.com/documents/v2/v2.first_mile.get_courier_delivery_detail?module=96&type=1
	GetCourierDeliveryDetail(ctx context.Context, sid uint64, mid uint64, tok string, opt GetCourierDeliveryDetailRequest) (*GetCourierDeliveryDetailResponse, error)
	// GetCourierDeliveryTrackingNumberList {"content":"<p>Use this api to get tracking number for courier delivery method.</p>","raw_content":[{"name":"paragraph","children":[{"data":"Use this api to get tracking number for courier delivery method."}]}]}
	// Path: /api/v2/first_mile/get_courier_delivery_tracking_number_list
	// https://open.shopee.com/documents/v2/v2.first_mile.get_courier_delivery_tracking_number_list?module=96&type=1
	GetCourierDeliveryTrackingNumberList(ctx context.Context, sid uint64, mid uint64, tok string, req GetCourierDeliveryTrackingNumberListRequest) (*GetCourierDeliveryTrackingNumberListResponse, error)
	// GetCourierDeliveryWaybill Use this api to get first mile waybill file for courier delivery method.
	// Path: /api/v2/first_mile/get_courier_delivery_waybill
	// https://open.shopee.com/documents/v2/v2.first_mile.get_courier_delivery_waybill?module=96&type=1
	GetCourierDeliveryWaybill(ctx context.Context, sid uint64, mid uint64, tok string, req GetCourierDeliveryWaybillRequest) (*GetCourierDeliveryWaybillResponse, error)
	// GetDetail Use this api to get first mile detail.
	// Path: /api/v2/first_mile/get_detail
	// https://open.shopee.com/documents/v2/v2.first_mile.get_detail?module=96&type=1
	GetDetail(ctx context.Context, sid uint64, mid uint64, tok string, opt GetDetailRequest) (*GetDetailResponse, error)
	// GetTrackingNumberList Use this api to get first mile tracking number list.
	// Path: /api/v2/first_mile/get_tracking_number_list
	// https://open.shopee.com/documents/v2/v2.first_mile.get_tracking_number_list?module=96&type=1
	GetTrackingNumberList(ctx context.Context, sid uint64, mid uint64, tok string, opt GetTrackingNumberListRequest) (*GetTrackingNumberListResponse, error)
	// GetTransitWarehouseList Use this api to get transit warehouse list which is used for first mile tracking number generation for courier delivery method.
	// Path: /api/v2/first_mile/get_transit_warehouse_list
	// https://open.shopee.com/documents/v2/v2.first_mile.get_transit_warehouse_list?module=96&type=1
	GetTransitWarehouseList(ctx context.Context, sid uint64, mid uint64, tok string, opt GetTransitWarehouseListRequest) (*GetTransitWarehouseListResponse, error)
	// GetUnbindOrderList Use this api to get unbind order list. It will only return orders unbound to first-mile that were created within the past 6 months.
	// Path: /api/v2/first_mile/get_unbind_order_list
	// https://open.shopee.com/documents/v2/v2.first_mile.get_unbind_order_list?module=96&type=1
	GetUnbindOrderList(ctx context.Context, sid uint64, mid uint64, tok string, opt GetUnbindOrderListRequest) (*GetUnbindOrderListResponse, error)
	// GetWaybill Use this api to get first mile waybill file.
	// Path: /api/v2/first_mile/get_waybill
	// https://open.shopee.com/documents/v2/v2.first_mile.get_waybill?module=96&type=1
	GetWaybill(ctx context.Context, sid uint64, mid uint64, tok string, req GetWaybillRequest) (*GetWaybillResponse, error)
	// GetWaybillFilelink {"content":"<p><span style=\"color:rgb(51,51,51);\">Use this api to get first mile waybill file link.</span></p>","raw_content":[{"name":"paragraph","children":[{"attributes":{"fontColor":"rgb(51,51,51)"},"data":"Use this api to get first mile waybill file link."}]}]}
	// Path: /api/v2/first_mile/get_waybill_filelink
	// https://open.shopee.com/documents/v2/v2.first_mile.get_waybill_filelink?module=96&type=1
	GetWaybillFilelink(ctx context.Context, sid uint64, mid uint64, tok string, req GetWaybillFilelinkRequest) (*GetWaybillFilelinkResponse, error)
	// UnbindFirstMileTrackingNumber Use this api to unbind first mile.
	// Path: /api/v2/first_mile/unbind_first_mile_tracking_number
	// https://open.shopee.com/documents/v2/v2.first_mile.unbind_first_mile_tracking_number?module=96&type=1
	UnbindFirstMileTrackingNumber(ctx context.Context, sid uint64, mid uint64, tok string, req UnbindFirstMileTrackingNumberRequest) (*UnbindFirstMileTrackingNumberResponse, error)
	// UnbindFirstMileTrackingNumberAll Use this api to unbind orders from first mile tracking number or binding ID.
	//
	// Path: /api/v2/first_mile/unbind_first_mile_tracking_number_all
	// https://open.shopee.com/documents/v2/v2.first_mile.unbind_first_mile_tracking_number_all?module=96&type=1
	UnbindFirstMileTrackingNumberAll(ctx context.Context, sid uint64, mid uint64, tok string, req UnbindFirstMileTrackingNumberAllRequest) (*UnbindFirstMileTrackingNumberAllResponse, error)
}

type FirstMileServiceOp[T any] struct {
	client *Client[T]
}

// BindCourierDeliveryFirstMileTrackingNumber Use this api to bind first mile tracking number for courier delivery method.
// Path: /api/v2/first_mile/bind_courier_delivery_first_mile_tracking_number
// https://open.shopee.com/documents/v2/v2.first_mile.bind_courier_delivery_first_mile_tracking_number?module=96&type=1
func (s *FirstMileServiceOp[T]) BindCourierDeliveryFirstMileTrackingNumber(ctx context.Context, sid uint64, mid uint64, tok string, req BindCourierDeliveryFirstMileTrackingNumberRequest) (*BindCourierDeliveryFirstMileTrackingNumberResponse, error) {
	path := "/first_mile/bind_courier_delivery_first_mile_tracking_number"
	resp := new(BindCourierDeliveryFirstMileTrackingNumberResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// BindFirstMileTrackingNumber Use this api to bind first mile tracking number.
// Path: /api/v2/first_mile/bind_first_mile_tracking_number
// https://open.shopee.com/documents/v2/v2.first_mile.bind_first_mile_tracking_number?module=96&type=1
func (s *FirstMileServiceOp[T]) BindFirstMileTrackingNumber(ctx context.Context, sid uint64, mid uint64, tok string, req BindFirstMileTrackingNumberRequest) (*BindFirstMileTrackingNumberResponse, error) {
	path := "/first_mile/bind_first_mile_tracking_number"
	resp := new(BindFirstMileTrackingNumberResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// CancelCourierDeliveryFirstMileTrackingNumber {"content":"<p>Use this api to cancel first mile tracking number for courier delivery method.</p>","raw_content":[{"name":"paragraph","children":[{"data":"Use this api to cancel first mile tracking number for courier delivery method."}]}]}
// Path: /api/v2/first_mile/cancel_courier_delivery_first_mile_tracking_number
// https://open.shopee.com/documents/v2/v2.first_mile.cancel_courier_delivery_first_mile_tracking_number?module=96&type=1
func (s *FirstMileServiceOp[T]) CancelCourierDeliveryFirstMileTrackingNumber(ctx context.Context, sid uint64, mid uint64, tok string, req CancelCourierDeliveryFirstMileTrackingNumberRequest) (*CancelCourierDeliveryFirstMileTrackingNumberResponse, error) {
	path := "/first_mile/cancel_courier_delivery_first_mile_tracking_number"
	resp := new(CancelCourierDeliveryFirstMileTrackingNumberResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// GenerateAndBindFirstMileTrackingNumber Use this api to generate first mile tracking number for courier delivery method.
// Path: /api/v2/first_mile/generate_and_bind_first_mile_tracking_number
// https://open.shopee.com/documents/v2/v2.first_mile.generate_and_bind_first_mile_tracking_number?module=96&type=1
func (s *FirstMileServiceOp[T]) GenerateAndBindFirstMileTrackingNumber(ctx context.Context, sid uint64, mid uint64, tok string, req GenerateAndBindFirstMileTrackingNumberRequest) (*GenerateAndBindFirstMileTrackingNumberResponse, error) {
	path := "/first_mile/generate_and_bind_first_mile_tracking_number"
	resp := new(GenerateAndBindFirstMileTrackingNumberResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// GenerateFirstMileTrackingNumber Use this api to generate first mile tracking number.
// Path: /api/v2/first_mile/generate_first_mile_tracking_number
// https://open.shopee.com/documents/v2/v2.first_mile.generate_first_mile_tracking_number?module=96&type=1
func (s *FirstMileServiceOp[T]) GenerateFirstMileTrackingNumber(ctx context.Context, sid uint64, mid uint64, tok string, req GenerateFirstMileTrackingNumberRequest) (*GenerateFirstMileTrackingNumberResponse, error) {
	path := "/first_mile/generate_first_mile_tracking_number"
	resp := new(GenerateFirstMileTrackingNumberResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// GetChannelList Use this api to get first mile channel list.
// Path: /api/v2/first_mile/get_channel_list
// https://open.shopee.com/documents/v2/v2.first_mile.get_channel_list?module=96&type=1
func (s *FirstMileServiceOp[T]) GetChannelList(ctx context.Context, sid uint64, mid uint64, tok string, req GetChannelListRequest) (*GetChannelListResponse, error) {
	path := "/first_mile/get_channel_list"
	resp := new(GetChannelListResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// GetCourierDeliveryChannelList {"content":"<p>Use this api to get courier information for courier delivery method.<br>&nbsp;</p>","raw_content":[{"name":"paragraph","children":[{"data":"Use this api to get courier information for courier delivery method."},{"name":"softBreak"},{"data":" "}]}]}
// Path: /api/v2/first_mile/get_courier_delivery_channel_list
// https://open.shopee.com/documents/v2/v2.first_mile.get_courier_delivery_channel_list?module=96&type=1
func (s *FirstMileServiceOp[T]) GetCourierDeliveryChannelList(ctx context.Context, sid uint64, mid uint64, tok string, opt GetCourierDeliveryChannelListRequest) (*GetCourierDeliveryChannelListResponse, error) {
	path := "/first_mile/get_courier_delivery_channel_list"
	resp := new(GetCourierDeliveryChannelListResponse)
	err := s.client.Get(ctx, path, resp, opt, sid, mid, tok)
	return resp, err
}

// GetCourierDeliveryDetail Use this api to get first mile detail for courier delivery method.
//
// Path: /api/v2/first_mile/get_courier_delivery_detail
// https://open.shopee.com/documents/v2/v2.first_mile.get_courier_delivery_detail?module=96&type=1
func (s *FirstMileServiceOp[T]) GetCourierDeliveryDetail(ctx context.Context, sid uint64, mid uint64, tok string, opt GetCourierDeliveryDetailRequest) (*GetCourierDeliveryDetailResponse, error) {
	path := "/first_mile/get_courier_delivery_detail"
	resp := new(GetCourierDeliveryDetailResponse)
	err := s.client.Get(ctx, path, resp, opt, sid, mid, tok)
	return resp, err
}

// GetCourierDeliveryTrackingNumberList {"content":"<p>Use this api to get tracking number for courier delivery method.</p>","raw_content":[{"name":"paragraph","children":[{"data":"Use this api to get tracking number for courier delivery method."}]}]}
// Path: /api/v2/first_mile/get_courier_delivery_tracking_number_list
// https://open.shopee.com/documents/v2/v2.first_mile.get_courier_delivery_tracking_number_list?module=96&type=1
func (s *FirstMileServiceOp[T]) GetCourierDeliveryTrackingNumberList(ctx context.Context, sid uint64, mid uint64, tok string, req GetCourierDeliveryTrackingNumberListRequest) (*GetCourierDeliveryTrackingNumberListResponse, error) {
	path := "/first_mile/get_courier_delivery_tracking_number_list"
	resp := new(GetCourierDeliveryTrackingNumberListResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// GetCourierDeliveryWaybill Use this api to get first mile waybill file for courier delivery method.
// Path: /api/v2/first_mile/get_courier_delivery_waybill
// https://open.shopee.com/documents/v2/v2.first_mile.get_courier_delivery_waybill?module=96&type=1
func (s *FirstMileServiceOp[T]) GetCourierDeliveryWaybill(ctx context.Context, sid uint64, mid uint64, tok string, req GetCourierDeliveryWaybillRequest) (*GetCourierDeliveryWaybillResponse, error) {
	path := "/first_mile/get_courier_delivery_waybill"
	resp := new(GetCourierDeliveryWaybillResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// GetDetail Use this api to get first mile detail.
// Path: /api/v2/first_mile/get_detail
// https://open.shopee.com/documents/v2/v2.first_mile.get_detail?module=96&type=1
func (s *FirstMileServiceOp[T]) GetDetail(ctx context.Context, sid uint64, mid uint64, tok string, opt GetDetailRequest) (*GetDetailResponse, error) {
	path := "/first_mile/get_detail"
	resp := new(GetDetailResponse)
	err := s.client.Get(ctx, path, resp, opt, sid, mid, tok)
	return resp, err
}

// GetTrackingNumberList Use this api to get first mile tracking number list.
// Path: /api/v2/first_mile/get_tracking_number_list
// https://open.shopee.com/documents/v2/v2.first_mile.get_tracking_number_list?module=96&type=1
func (s *FirstMileServiceOp[T]) GetTrackingNumberList(ctx context.Context, sid uint64, mid uint64, tok string, opt GetTrackingNumberListRequest) (*GetTrackingNumberListResponse, error) {
	path := "/first_mile/get_tracking_number_list"
	resp := new(GetTrackingNumberListResponse)
	err := s.client.Get(ctx, path, resp, opt, sid, mid, tok)
	return resp, err
}

// GetTransitWarehouseList Use this api to get transit warehouse list which is used for first mile tracking number generation for courier delivery method.
// Path: /api/v2/first_mile/get_transit_warehouse_list
// https://open.shopee.com/documents/v2/v2.first_mile.get_transit_warehouse_list?module=96&type=1
func (s *FirstMileServiceOp[T]) GetTransitWarehouseList(ctx context.Context, sid uint64, mid uint64, tok string, opt GetTransitWarehouseListRequest) (*GetTransitWarehouseListResponse, error) {
	path := "/first_mile/get_transit_warehouse_list"
	resp := new(GetTransitWarehouseListResponse)
	err := s.client.Get(ctx, path, resp, opt, sid, mid, tok)
	return resp, err
}

// GetUnbindOrderList Use this api to get unbind order list. It will only return orders unbound to first-mile that were created within the past 6 months.
// Path: /api/v2/first_mile/get_unbind_order_list
// https://open.shopee.com/documents/v2/v2.first_mile.get_unbind_order_list?module=96&type=1
func (s *FirstMileServiceOp[T]) GetUnbindOrderList(ctx context.Context, sid uint64, mid uint64, tok string, opt GetUnbindOrderListRequest) (*GetUnbindOrderListResponse, error) {
	path := "/first_mile/get_unbind_order_list"
	resp := new(GetUnbindOrderListResponse)
	err := s.client.Get(ctx, path, resp, opt, sid, mid, tok)
	return resp, err
}

// GetWaybill Use this api to get first mile waybill file.
// Path: /api/v2/first_mile/get_waybill
// https://open.shopee.com/documents/v2/v2.first_mile.get_waybill?module=96&type=1
func (s *FirstMileServiceOp[T]) GetWaybill(ctx context.Context, sid uint64, mid uint64, tok string, req GetWaybillRequest) (*GetWaybillResponse, error) {
	path := "/first_mile/get_waybill"
	resp := new(GetWaybillResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// GetWaybillFilelink {"content":"<p><span style=\"color:rgb(51,51,51);\">Use this api to get first mile waybill file link.</span></p>","raw_content":[{"name":"paragraph","children":[{"attributes":{"fontColor":"rgb(51,51,51)"},"data":"Use this api to get first mile waybill file link."}]}]}
// Path: /api/v2/first_mile/get_waybill_filelink
// https://open.shopee.com/documents/v2/v2.first_mile.get_waybill_filelink?module=96&type=1
func (s *FirstMileServiceOp[T]) GetWaybillFilelink(ctx context.Context, sid uint64, mid uint64, tok string, req GetWaybillFilelinkRequest) (*GetWaybillFilelinkResponse, error) {
	path := "/first_mile/get_waybill_filelink"
	resp := new(GetWaybillFilelinkResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// UnbindFirstMileTrackingNumber Use this api to unbind first mile.
// Path: /api/v2/first_mile/unbind_first_mile_tracking_number
// https://open.shopee.com/documents/v2/v2.first_mile.unbind_first_mile_tracking_number?module=96&type=1
func (s *FirstMileServiceOp[T]) UnbindFirstMileTrackingNumber(ctx context.Context, sid uint64, mid uint64, tok string, req UnbindFirstMileTrackingNumberRequest) (*UnbindFirstMileTrackingNumberResponse, error) {
	path := "/first_mile/unbind_first_mile_tracking_number"
	resp := new(UnbindFirstMileTrackingNumberResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// UnbindFirstMileTrackingNumberAll Use this api to unbind orders from first mile tracking number or binding ID.
//
// Path: /api/v2/first_mile/unbind_first_mile_tracking_number_all
// https://open.shopee.com/documents/v2/v2.first_mile.unbind_first_mile_tracking_number_all?module=96&type=1
func (s *FirstMileServiceOp[T]) UnbindFirstMileTrackingNumberAll(ctx context.Context, sid uint64, mid uint64, tok string, req UnbindFirstMileTrackingNumberAllRequest) (*UnbindFirstMileTrackingNumberAllResponse, error) {
	path := "/first_mile/unbind_first_mile_tracking_number_all"
	resp := new(UnbindFirstMileTrackingNumberAllResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}
