package goshopee

import (
	"context"
)

type PublicService interface {
	// GetMerchantsByPartner Use this API to get basic info of merchants which have authorized to the partner.
	// Path: /api/v2/public/get_merchants_by_partner
	// https://open.shopee.com/documents/v2/v2.public.get_merchants_by_partner?module=104&type=1
	GetMerchantsByPartner(ctx context.Context, sid uint64, opt GetMerchantsByPartnerRequest, tok string) (*GetMerchantsByPartnerResponse, error)
	// GetShopeeIpRanges You can get shopee ip address ranges through this open api.
	// Path: /api/v2/public/get_shopee_ip_ranges
	// https://open.shopee.com/documents/v2/v2.public.get_shopee_ip_ranges?module=104&type=1
	GetShopeeIpRanges(ctx context.Context, sid uint64, tok string) (*GetShopeeIpRangesResponse, error)
	// GetShopsByPartner get basic info of shops which have authorized to the partner.
	// Path: /api/v2/public/get_shops_by_partner
	// https://open.shopee.com/documents/v2/v2.public.get_shops_by_partner?module=104&type=1
	GetShopsByPartner(ctx context.Context, sid uint64, opt GetShopsByPartnerRequest, tok string) (*GetShopsByPartnerResponse, error)
	// GetTokenByResendCode Use the resend code to get access token and refresh token. When you lost your access token or refresh token, you can go to authorization management page to resend code by yourselves. You can only use this endpoint in live environment, we don't support in test-stable environment.
	// Path: /api/v2/public/get_token_by_resend_code
	// https://open.shopee.com/documents/v2/v2.public.get_token_by_resend_code?module=104&type=1
	GetTokenByResendCode(ctx context.Context, sid uint64, req GetTokenByResendCodeRequest, tok string) (*GetTokenByResendCodeResponse, error)
}

type PublicServiceOp[T any] struct {
	client *Client[T]
}

// GetMerchantsByPartner Use this API to get basic info of merchants which have authorized to the partner.
// Path: /api/v2/public/get_merchants_by_partner
// https://open.shopee.com/documents/v2/v2.public.get_merchants_by_partner?module=104&type=1
func (s *PublicServiceOp[T]) GetMerchantsByPartner(ctx context.Context, sid uint64, opt GetMerchantsByPartnerRequest, tok string) (*GetMerchantsByPartnerResponse, error) {
	path := "/public/get_merchants_by_partner"
	resp := new(GetMerchantsByPartnerResponse)
	err := s.client.Get(ctx, path, resp, opt)
	return resp, err
}

// GetShopeeIpRanges You can get shopee ip address ranges through this open api.
// Path: /api/v2/public/get_shopee_ip_ranges
// https://open.shopee.com/documents/v2/v2.public.get_shopee_ip_ranges?module=104&type=1
func (s *PublicServiceOp[T]) GetShopeeIpRanges(ctx context.Context, sid uint64, tok string) (*GetShopeeIpRangesResponse, error) {
	path := "/public/get_shopee_ip_ranges"
	resp := new(GetShopeeIpRangesResponse)
	err := s.client.Get(ctx, path, resp, nil)
	return resp, err
}

// GetShopsByPartner get basic info of shops which have authorized to the partner.
// Path: /api/v2/public/get_shops_by_partner
// https://open.shopee.com/documents/v2/v2.public.get_shops_by_partner?module=104&type=1
func (s *PublicServiceOp[T]) GetShopsByPartner(ctx context.Context, sid uint64, opt GetShopsByPartnerRequest, tok string) (*GetShopsByPartnerResponse, error) {
	path := "/public/get_shops_by_partner"
	resp := new(GetShopsByPartnerResponse)
	err := s.client.Get(ctx, path, resp, opt)
	return resp, err
}

// GetTokenByResendCode Use the resend code to get access token and refresh token. When you lost your access token or refresh token, you can go to authorization management page to resend code by yourselves. You can only use this endpoint in live environment, we don't support in test-stable environment.
// Path: /api/v2/public/get_token_by_resend_code
// https://open.shopee.com/documents/v2/v2.public.get_token_by_resend_code?module=104&type=1
func (s *PublicServiceOp[T]) GetTokenByResendCode(ctx context.Context, sid uint64, req GetTokenByResendCodeRequest, tok string) (*GetTokenByResendCodeResponse, error) {
	path := "/public/get_token_by_resend_code"
	resp := new(GetTokenByResendCodeResponse)
	err := s.client.Post(ctx, path, req, resp)
	return resp, err
}
