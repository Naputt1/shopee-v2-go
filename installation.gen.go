package goshopee

import (
	"context"
)

type InstallationService interface {
	// ReceiveInstallationOrder {"content":"<p>Receive installation order. Use this API to move installation status to PARCEL_PICKED_UP.</p>","raw_content":[{"name":"paragraph","children":[{"data":"Receive installation order. Use this API to move installation status to PARCEL_PICKED_UP."}]}]}
	// Path: /api/v2/installation/receive_installation_order
	// https://open.shopee.com/documents/v2/v2.installation.receive_installation_order?module=138&type=1
	ReceiveInstallationOrder(ctx context.Context, sid uint64, mid uint64, tok string, req ReceiveInstallationOrderRequest) (*ReceiveInstallationOrderResponse, error)
	// ValidateInstallationOtpCode {"content":"<p>Validate installation order OTP code. Use this API to validate OTP code when parcel status is PARCEL_PICKED_UP, PARCEL_REQUEST_CANCELLED, or PARCEL_CANCELLED.</p>","raw_content":[{"name":"paragraph","children":[{"data":"Validate installation order OTP code. Use this API to validate OTP code when parcel status is PARCEL_PICKED_UP, PARCEL_REQUEST_CANCELLED, or PARCEL_CANCELLED."}]}]}
	// Path: /api/v2/installation/validate_installation_otp_code
	// https://open.shopee.com/documents/v2/v2.installation.validate_installation_otp_code?module=138&type=1
	ValidateInstallationOtpCode(ctx context.Context, sid uint64, mid uint64, tok string, req ValidateInstallationOtpCodeRequest) (*ValidateInstallationOtpCodeResponse, error)
}

type InstallationServiceOp[T any] struct {
	client *Client[T]
}

// ReceiveInstallationOrder {"content":"<p>Receive installation order. Use this API to move installation status to PARCEL_PICKED_UP.</p>","raw_content":[{"name":"paragraph","children":[{"data":"Receive installation order. Use this API to move installation status to PARCEL_PICKED_UP."}]}]}
// Path: /api/v2/installation/receive_installation_order
// https://open.shopee.com/documents/v2/v2.installation.receive_installation_order?module=138&type=1
func (s *InstallationServiceOp[T]) ReceiveInstallationOrder(ctx context.Context, sid uint64, mid uint64, tok string, req ReceiveInstallationOrderRequest) (*ReceiveInstallationOrderResponse, error) {
	path := "/installation/receive_installation_order"
	resp := new(ReceiveInstallationOrderResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// ValidateInstallationOtpCode {"content":"<p>Validate installation order OTP code. Use this API to validate OTP code when parcel status is PARCEL_PICKED_UP, PARCEL_REQUEST_CANCELLED, or PARCEL_CANCELLED.</p>","raw_content":[{"name":"paragraph","children":[{"data":"Validate installation order OTP code. Use this API to validate OTP code when parcel status is PARCEL_PICKED_UP, PARCEL_REQUEST_CANCELLED, or PARCEL_CANCELLED."}]}]}
// Path: /api/v2/installation/validate_installation_otp_code
// https://open.shopee.com/documents/v2/v2.installation.validate_installation_otp_code?module=138&type=1
func (s *InstallationServiceOp[T]) ValidateInstallationOtpCode(ctx context.Context, sid uint64, mid uint64, tok string, req ValidateInstallationOtpCodeRequest) (*ValidateInstallationOtpCodeResponse, error) {
	path := "/installation/validate_installation_otp_code"
	resp := new(ValidateInstallationOtpCodeResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}
