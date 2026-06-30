package goshopee

type AuthedMerchant struct {
	Region string `json:"region"` // [Required] Merchant's area
	MerchantId int64 `json:"merchant_id"` // [Required] Shopee's unique identifier for a merchant.
	AuthTime int64 `json:"auth_time"` // [Required] The timestamp when the merchant was authorized to the partner.
	ExpireTime int64 `json:"expire_time"` // [Required] Use this field to indicate the expiration date for merchant authorization.
}
type AuthedShop struct {
	Region string `json:"region"` // [Required] Shop's area
	ShopId int64 `json:"shop_id"` // [Required] Shop id
	AuthTime int64 `json:"auth_time"` // [Required] The timestamp when the shop was authorized to the partner.
	ExpireTime int64 `json:"expire_time"` // [Required] Use this field to indicate the expiration date for shop authorization.
	SipAffiShopList []SipAffiShop `json:"sip_affi_shop_list"` // [Required] SIP affiliate shops info list
}
type GetMerchantsByPartnerRequest struct {
	PageNo *int64 `json:"page_no,omitempty" url:"page_no,omitempty"` // [Optional] Specifies the page number of data to return in the current call. Starting from 1. if data is more than one page, the page_no can be some entry to start next call.
	PageSize *int64 `json:"page_size,omitempty" url:"page_size,omitempty"` // [Optional] Each result set is returned as a page of entries. Use the "page_size" filters to control the maximum number of entries to retrieve per page (i.e., per call), and the "page_no" to start next call. This integer value is used to specify the maximum number of entries to return in a single "page" of data.
}
type GetMerchantsByPartnerResponse struct {
	BaseResponse // Common response fields
	AuthedMerchantList []AuthedMerchant `json:"authed_merchant_list,omitempty"` // A list of merchants authorized to the partner.
	More bool `json:"more,omitempty"` // This is to indicate whether the list is more than one page. If this value is true, you may want to continue to check next page to retrieve the rest of datas.
}
type GetShopeeIpRangesResponse struct {
	BaseResponse // Common response fields
	IpList []string `json:"ip_list,omitempty"` // <p>IP address ranges of Shopee</p>
}
type GetShopsByPartnerRequest struct {
	PageNo *int64 `json:"page_no,omitempty" url:"page_no,omitempty"` // [Optional] Specifies the page number of data to return in the current call. Starting from 1. if data is more than one page, the page_no can be some entry to start next call.
	PageSize *int64 `json:"page_size,omitempty" url:"page_size,omitempty"` // [Optional] Each result set is returned as a page of entries. Use the "page_size" filters to control the maximum number of entries to retrieve per page (i.e., per call), and the "page_no" to start next call. This integer value is used to specify the maximum number of entries to return in a single "page" of data.
}
type GetShopsByPartnerResponse struct {
	BaseResponse // Common response fields
	AuthedShopList []AuthedShop `json:"authed_shop_list,omitempty"` // A list of shops authorized to the partner.
	More bool `json:"more,omitempty"` // This is to indicate whether the list is more than one page. If this value is true, you may want to continue to check next page to retrieve the rest of datas.
}
type GetTokenByResendCodeRequest struct {
	ResendCode string `json:"resend_code"` // [Required] the code in redirect url after you resend code in shop authorization management page. valid for one-time use, expires in 10minutes.
}
type GetTokenByResendCodeResponse struct {
	BaseResponse // Common response fields
	AccessToken string `json:"access_token,omitempty"` // The token for API access, using to identify your permission to the api. Valid for multiple use and expires in 4 hours.
	ExpireIn int64 `json:"expire_in,omitempty"` // Access_token expiration time, unit is second.
	MerchantIdList []int64 `json:"merchant_id_list,omitempty"` // Return when resend code in merchant module
	RefreshToken string `json:"refresh_token,omitempty"` // Use refresh_token to obtain new access_token. Valid for each shop_id and merchant_id respectively one-time use, expires in 30 days.
	ShopIdList []int64 `json:"shop_id_list,omitempty"` // Return when resend code in shop module
}
