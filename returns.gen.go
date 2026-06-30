package goshopee

import (
	"context"
	"io"
)

type ReturnsService interface {
	// AcceptOffer v2.returns.accept_offer
	// Path: /api/v2/returns/accept_offer
	// https://open.shopee.com/documents/v2/v2.returns.accept_offer?module=102&type=1
	AcceptOffer(ctx context.Context, sid uint64, req AcceptOfferRequest, tok string) (*AcceptOfferResponse, error)
	// CancelDispute Sellers can only cancel compensation disputes, not normal disputes. This means that sellers can only cancel disputes when the return_status is ACCEPTED and the compensation_status is COMPENSATION_REQUESTED.
	// Path: /api/v2/returns/cancel_dispute
	// https://open.shopee.com/documents/v2/v2.returns.cancel_dispute?module=102&type=1
	CancelDispute(ctx context.Context, sid uint64, req CancelDisputeRequest, tok string) (*CancelDisputeResponse, error)
	// Confirm Confirm refund
	// Path: /api/v2/returns/confirm
	// https://open.shopee.com/documents/v2/v2.returns.confirm?module=102&type=1
	Confirm(ctx context.Context, sid uint64, req ConfirmRequest, tok string) (*ConfirmResponse, error)
	// ConvertImage Convert a specific format and pictures within 10M into url.
	// Path: /api/v2/returns/convert_image
	// https://open.shopee.com/documents/v2/v2.returns.convert_image?module=102&type=1
	ConvertImage(ctx context.Context, sid uint64, req ConvertImageRequest, tok string) (*ConvertImageResponse, error)
	// Dispute Dispute return.
	//
	// Support to raise dispute when return_status in REQUESTED / PROCESSING/ACCEPTED
	// Path: /api/v2/returns/dispute
	// https://open.shopee.com/documents/v2/v2.returns.dispute?module=102&type=1
	Dispute(ctx context.Context, sid uint64, req DisputeRequest, tok string) (*DisputeResponse, error)
	// GetAvailableSolutions Get the available solutions offered to buyers.
	// Path: /api/v2/returns/get_available_solutions
	// https://open.shopee.com/documents/v2/v2.returns.get_available_solutions?module=102&type=1
	GetAvailableSolutions(ctx context.Context, sid uint64, opt GetAvailableSolutionsRequest, tok string) (*GetAvailableSolutionsResponse, error)
	// GetReturnDetail Use this api to get detail information of a return by return sn.
	// Path: /api/v2/returns/get_return_detail
	// https://open.shopee.com/documents/v2/v2.returns.get_return_detail?module=102&type=1
	GetReturnDetail(ctx context.Context, sid uint64, opt GetReturnDetailRequest, tok string) (*GetReturnDetailResponse, error)
	// GetReturnDisputeReason To get the dispute return reason.
	// Path: /api/v2/returns/get_return_dispute_reason
	// https://open.shopee.com/documents/v2/v2.returns.get_return_dispute_reason?module=102&type=1
	GetReturnDisputeReason(ctx context.Context, sid uint64, opt GetReturnDisputeReasonRequest, tok string) (*GetReturnDisputeReasonResponse, error)
	// GetReturnList Use this api to get detail information of many return by shop id.
	// Path: /api/v2/returns/get_return_list
	// https://open.shopee.com/documents/v2/v2.returns.get_return_list?module=102&type=1
	GetReturnList(ctx context.Context, sid uint64, opt GetReturnListRequest, tok string) (*GetReturnListResponse, error)
	// GetReverseTrackingInfo {"content":"<p>Get reverse and post-return logistics information of return request. For Normal RR, return complete reverse logistics information, for In-transit RR and Return-on-the-Spot, only return latest reverse logistics status, without providing complete reverse logistics information. For seller_validation, only one segment of reverse (buyer to seller), use tracking_info, for warehouse_validation, two segment of reverse (buyer to warehouse and warehouse to seller), use post_return_logistics_tracking_info.</p>","raw_content":[{"name":"paragraph","children":[{"data":"Get reverse and post-return logistics information of return request. For Normal RR, return complete reverse logistics information, for In-transit RR and Return-on-the-Spot, only return latest reverse logistics status, without providing complete reverse logistics information. For seller_validation, only one segment of reverse (buyer to seller), use tracking_info, for warehouse_validation, two segment of reverse (buyer to warehouse and warehouse to seller), use post_return_logistics_tracking_info."}]}]}
	// Path: /api/v2/returns/get_reverse_tracking_info
	// https://open.shopee.com/documents/v2/v2.returns.get_reverse_tracking_info?module=102&type=1
	GetReverseTrackingInfo(ctx context.Context, sid uint64, opt GetReverseTrackingInfoRequest, tok string) (*GetReverseTrackingInfoResponse, error)
	// GetShippingCarrier Use this API to get the list of shipping carriers and request parameters needed before calling v2.returns.upload_shipping_proof. Only for TW and BR returns with is_seller_arrange = true.
	// Path: /api/v2/returns/get_shipping_carrier
	// https://open.shopee.com/documents/v2/v2.returns.get_shipping_carrier?module=102&type=1
	GetShippingCarrier(ctx context.Context, sid uint64, opt GetShippingCarrierRequest, tok string) (*GetShippingCarrierResponse, error)
	// Offer v2.returns.offer
	// Path: /api/v2/returns/offer
	// https://open.shopee.com/documents/v2/v2.returns.offer?module=102&type=1
	Offer(ctx context.Context, sid uint64, req OfferRequest, tok string) (*OfferResponse, error)
	// QueryProof Support sellers to query the evidence uploaded through the upload evidence API.
	// Path: /api/v2/returns/query_proof
	// https://open.shopee.com/documents/v2/v2.returns.query_proof?module=102&type=1
	QueryProof(ctx context.Context, sid uint64, opt QueryProofRequest, tok string) (*QueryProofResponse, error)
	// UploadProof Support sellers to upload evidence, including text and pictures and videos converted into URLs.
	// Path: /api/v2/returns/upload_proof
	// https://open.shopee.com/documents/v2/v2.returns.upload_proof?module=102&type=1
	UploadProof(ctx context.Context, sid uint64, filename string, tok string) (*UploadProofResponse, error)
	UploadProofFromReader(ctx context.Context, sid uint64, filename string, reader io.Reader, tok string) (*UploadProofResponse, error)
	// UploadShippingProof Use this API to upload shipping proof (Only for TW and BR returns with is_seller_arrange = true). This is not to upload evidence for disputes.
	// Path: /api/v2/returns/upload_shipping_proof
	// https://open.shopee.com/documents/v2/v2.returns.upload_shipping_proof?module=102&type=1
	UploadShippingProof(ctx context.Context, sid uint64, filename string, tok string) (*UploadShippingProofResponse, error)
	UploadShippingProofFromReader(ctx context.Context, sid uint64, filename string, reader io.Reader, tok string) (*UploadShippingProofResponse, error)
}

type ReturnsServiceOp[T any] struct {
	client *Client[T]
}

// AcceptOffer v2.returns.accept_offer
// Path: /api/v2/returns/accept_offer
// https://open.shopee.com/documents/v2/v2.returns.accept_offer?module=102&type=1
func (s *ReturnsServiceOp[T]) AcceptOffer(ctx context.Context, sid uint64, req AcceptOfferRequest, tok string) (*AcceptOfferResponse, error) {
	path := "/returns/accept_offer"
	resp := new(AcceptOfferResponse)
	err := s.client.WithShop(sid, tok).Post(ctx, path, req, resp)
	return resp, err
}

// CancelDispute Sellers can only cancel compensation disputes, not normal disputes. This means that sellers can only cancel disputes when the return_status is ACCEPTED and the compensation_status is COMPENSATION_REQUESTED.
// Path: /api/v2/returns/cancel_dispute
// https://open.shopee.com/documents/v2/v2.returns.cancel_dispute?module=102&type=1
func (s *ReturnsServiceOp[T]) CancelDispute(ctx context.Context, sid uint64, req CancelDisputeRequest, tok string) (*CancelDisputeResponse, error) {
	path := "/returns/cancel_dispute"
	resp := new(CancelDisputeResponse)
	err := s.client.WithShop(sid, tok).Post(ctx, path, req, resp)
	return resp, err
}

// Confirm Confirm refund
// Path: /api/v2/returns/confirm
// https://open.shopee.com/documents/v2/v2.returns.confirm?module=102&type=1
func (s *ReturnsServiceOp[T]) Confirm(ctx context.Context, sid uint64, req ConfirmRequest, tok string) (*ConfirmResponse, error) {
	path := "/returns/confirm"
	resp := new(ConfirmResponse)
	err := s.client.WithShop(sid, tok).Post(ctx, path, req, resp)
	return resp, err
}

// ConvertImage Convert a specific format and pictures within 10M into url.
// Path: /api/v2/returns/convert_image
// https://open.shopee.com/documents/v2/v2.returns.convert_image?module=102&type=1
func (s *ReturnsServiceOp[T]) ConvertImage(ctx context.Context, sid uint64, req ConvertImageRequest, tok string) (*ConvertImageResponse, error) {
	path := "/returns/convert_image"
	resp := new(ConvertImageResponse)
	err := s.client.WithShop(sid, tok).Post(ctx, path, req, resp)
	return resp, err
}

// Dispute Dispute return.
//
// Support to raise dispute when return_status in REQUESTED / PROCESSING/ACCEPTED
// Path: /api/v2/returns/dispute
// https://open.shopee.com/documents/v2/v2.returns.dispute?module=102&type=1
func (s *ReturnsServiceOp[T]) Dispute(ctx context.Context, sid uint64, req DisputeRequest, tok string) (*DisputeResponse, error) {
	path := "/returns/dispute"
	resp := new(DisputeResponse)
	err := s.client.WithShop(sid, tok).Post(ctx, path, req, resp)
	return resp, err
}

// GetAvailableSolutions Get the available solutions offered to buyers.
// Path: /api/v2/returns/get_available_solutions
// https://open.shopee.com/documents/v2/v2.returns.get_available_solutions?module=102&type=1
func (s *ReturnsServiceOp[T]) GetAvailableSolutions(ctx context.Context, sid uint64, opt GetAvailableSolutionsRequest, tok string) (*GetAvailableSolutionsResponse, error) {
	path := "/returns/get_available_solutions"
	resp := new(GetAvailableSolutionsResponse)
	err := s.client.WithShop(sid, tok).Get(ctx, path, resp, opt)
	return resp, err
}

// GetReturnDetail Use this api to get detail information of a return by return sn.
// Path: /api/v2/returns/get_return_detail
// https://open.shopee.com/documents/v2/v2.returns.get_return_detail?module=102&type=1
func (s *ReturnsServiceOp[T]) GetReturnDetail(ctx context.Context, sid uint64, opt GetReturnDetailRequest, tok string) (*GetReturnDetailResponse, error) {
	path := "/returns/get_return_detail"
	resp := new(GetReturnDetailResponse)
	err := s.client.WithShop(sid, tok).Get(ctx, path, resp, opt)
	return resp, err
}

// GetReturnDisputeReason To get the dispute return reason.
// Path: /api/v2/returns/get_return_dispute_reason
// https://open.shopee.com/documents/v2/v2.returns.get_return_dispute_reason?module=102&type=1
func (s *ReturnsServiceOp[T]) GetReturnDisputeReason(ctx context.Context, sid uint64, opt GetReturnDisputeReasonRequest, tok string) (*GetReturnDisputeReasonResponse, error) {
	path := "/returns/get_return_dispute_reason"
	resp := new(GetReturnDisputeReasonResponse)
	err := s.client.WithShop(sid, tok).Get(ctx, path, resp, opt)
	return resp, err
}

// GetReturnList Use this api to get detail information of many return by shop id.
// Path: /api/v2/returns/get_return_list
// https://open.shopee.com/documents/v2/v2.returns.get_return_list?module=102&type=1
func (s *ReturnsServiceOp[T]) GetReturnList(ctx context.Context, sid uint64, opt GetReturnListRequest, tok string) (*GetReturnListResponse, error) {
	path := "/returns/get_return_list"
	resp := new(GetReturnListResponse)
	err := s.client.WithShop(sid, tok).Get(ctx, path, resp, opt)
	return resp, err
}

// GetReverseTrackingInfo {"content":"<p>Get reverse and post-return logistics information of return request. For Normal RR, return complete reverse logistics information, for In-transit RR and Return-on-the-Spot, only return latest reverse logistics status, without providing complete reverse logistics information. For seller_validation, only one segment of reverse (buyer to seller), use tracking_info, for warehouse_validation, two segment of reverse (buyer to warehouse and warehouse to seller), use post_return_logistics_tracking_info.</p>","raw_content":[{"name":"paragraph","children":[{"data":"Get reverse and post-return logistics information of return request. For Normal RR, return complete reverse logistics information, for In-transit RR and Return-on-the-Spot, only return latest reverse logistics status, without providing complete reverse logistics information. For seller_validation, only one segment of reverse (buyer to seller), use tracking_info, for warehouse_validation, two segment of reverse (buyer to warehouse and warehouse to seller), use post_return_logistics_tracking_info."}]}]}
// Path: /api/v2/returns/get_reverse_tracking_info
// https://open.shopee.com/documents/v2/v2.returns.get_reverse_tracking_info?module=102&type=1
func (s *ReturnsServiceOp[T]) GetReverseTrackingInfo(ctx context.Context, sid uint64, opt GetReverseTrackingInfoRequest, tok string) (*GetReverseTrackingInfoResponse, error) {
	path := "/returns/get_reverse_tracking_info"
	resp := new(GetReverseTrackingInfoResponse)
	err := s.client.WithShop(sid, tok).Get(ctx, path, resp, opt)
	return resp, err
}

// GetShippingCarrier Use this API to get the list of shipping carriers and request parameters needed before calling v2.returns.upload_shipping_proof. Only for TW and BR returns with is_seller_arrange = true.
// Path: /api/v2/returns/get_shipping_carrier
// https://open.shopee.com/documents/v2/v2.returns.get_shipping_carrier?module=102&type=1
func (s *ReturnsServiceOp[T]) GetShippingCarrier(ctx context.Context, sid uint64, opt GetShippingCarrierRequest, tok string) (*GetShippingCarrierResponse, error) {
	path := "/returns/get_shipping_carrier"
	resp := new(GetShippingCarrierResponse)
	err := s.client.WithShop(sid, tok).Get(ctx, path, resp, opt)
	return resp, err
}

// Offer v2.returns.offer
// Path: /api/v2/returns/offer
// https://open.shopee.com/documents/v2/v2.returns.offer?module=102&type=1
func (s *ReturnsServiceOp[T]) Offer(ctx context.Context, sid uint64, req OfferRequest, tok string) (*OfferResponse, error) {
	path := "/returns/offer"
	resp := new(OfferResponse)
	err := s.client.WithShop(sid, tok).Post(ctx, path, req, resp)
	return resp, err
}

// QueryProof Support sellers to query the evidence uploaded through the upload evidence API.
// Path: /api/v2/returns/query_proof
// https://open.shopee.com/documents/v2/v2.returns.query_proof?module=102&type=1
func (s *ReturnsServiceOp[T]) QueryProof(ctx context.Context, sid uint64, opt QueryProofRequest, tok string) (*QueryProofResponse, error) {
	path := "/returns/query_proof"
	resp := new(QueryProofResponse)
	err := s.client.WithShop(sid, tok).Get(ctx, path, resp, opt)
	return resp, err
}

// UploadProof Support sellers to upload evidence, including text and pictures and videos converted into URLs.
// Path: /api/v2/returns/upload_proof
// https://open.shopee.com/documents/v2/v2.returns.upload_proof?module=102&type=1
func (s *ReturnsServiceOp[T]) UploadProof(ctx context.Context, sid uint64, filename string, tok string) (*UploadProofResponse, error) {
	path := "/returns/upload_proof"
	resp := new(UploadProofResponse)
	err := s.client.WithShop(sid, tok).Upload(ctx, path, "image", filename, resp)
	return resp, err
}

func (s *ReturnsServiceOp[T]) UploadProofFromReader(ctx context.Context, sid uint64, filename string, reader io.Reader, tok string) (*UploadProofResponse, error) {
	path := "/returns/upload_proof"
	resp := new(UploadProofResponse)
	err := s.client.WithShop(sid, tok).UploadFromReader(ctx, path, "image", filename, reader, resp)
	return resp, err
}

// UploadShippingProof Use this API to upload shipping proof (Only for TW and BR returns with is_seller_arrange = true). This is not to upload evidence for disputes.
// Path: /api/v2/returns/upload_shipping_proof
// https://open.shopee.com/documents/v2/v2.returns.upload_shipping_proof?module=102&type=1
func (s *ReturnsServiceOp[T]) UploadShippingProof(ctx context.Context, sid uint64, filename string, tok string) (*UploadShippingProofResponse, error) {
	path := "/returns/upload_shipping_proof"
	resp := new(UploadShippingProofResponse)
	err := s.client.WithShop(sid, tok).Upload(ctx, path, "image", filename, resp)
	return resp, err
}

func (s *ReturnsServiceOp[T]) UploadShippingProofFromReader(ctx context.Context, sid uint64, filename string, reader io.Reader, tok string) (*UploadShippingProofResponse, error) {
	path := "/returns/upload_shipping_proof"
	resp := new(UploadShippingProofResponse)
	err := s.client.WithShop(sid, tok).UploadFromReader(ctx, path, "image", filename, reader, resp)
	return resp, err
}
