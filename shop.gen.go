package goshopee

type ShopService interface {
	// GetAuthorisedResellerBrand Get the authorised reseller brand list for the shop.
	// Path: /api/v2/shop/get_authorised_reseller_brand
	// https://open.shopee.com/documents/v2/v2.shop.get_authorised_reseller_brand?module=92&type=1
	GetAuthorisedResellerBrand(sid uint64, opt GetAuthorisedResellerBrandRequest, tok string) (*GetAuthorisedResellerBrandResponse, error)
	// GetBrShopOnboardingInfo [For BR Shop Only] Use this API to get shop KYC registration and qualification information.
	// Path: /api/v2/shop/get_br_shop_onboarding_info
	// https://open.shopee.com/documents/v2/v2.shop.get_br_shop_onboarding_info?module=92&type=1
	GetBrShopOnboardingInfo(sid uint64, tok string) (*GetBrShopOnboardingInfoResponse, error)
	// GetProfile This API support to get information of shop.
	// Path: /api/v2/shop/get_profile
	// https://open.shopee.com/documents/v2/v2.shop.get_profile?module=92&type=1
	GetProfile(sid uint64, tok string) (*GetProfileResponse, error)
	// GetShopHolidayMode Use this API to check whether a shop has enabled holiday mode and its ongoing and upcoming holiday mode period.
	// Path: /api/v2/shop/get_shop_holiday_mode
	// https://open.shopee.com/documents/v2/v2.shop.get_shop_holiday_mode?module=92&type=1
	GetShopHolidayMode(sid uint64, tok string) (*GetShopHolidayModeResponse, error)
	// GetShopInfo {"content":"<p>Use this call to get information of shop</p>","raw_content":[{"name":"paragraph","children":[{"data":"Use this call to get information of shop"}]}]}
	// Path: /api/v2/shop/get_shop_info
	// https://open.shopee.com/documents/v2/v2.shop.get_shop_info?module=92&type=1
	GetShopInfo(sid uint64, tok string) (*GetShopInfoResponse, error)
	// GetShopNotification get Seller Center notification, the permission is controlled by App type
	// Path: /api/v2/shop/get_shop_notification
	// https://open.shopee.com/documents/v2/v2.shop.get_shop_notification?module=92&type=1
	GetShopNotification(sid uint64, opt GetShopNotificationRequest, tok string) (*GetShopNotificationResponse, error)
	// GetWarehouseDetail For given shop id and region, return warehouse info including warehouse id, address id and location id, return all warehouse with once call.
	// Path: /api/v2/shop/get_warehouse_detail
	// https://open.shopee.com/documents/v2/v2.shop.get_warehouse_detail?module=92&type=1
	GetWarehouseDetail(sid uint64, opt GetWarehouseDetailRequest, tok string) (*GetWarehouseDetailResponse, error)
	// SetShopHolidayMode Use this API to set holiday periods in advance for automatic on/off of holiday mode and there are two holiday modes allowing sellers to choose whether to accept new orders during holiday.
	// Path: /api/v2/shop/set_shop_holiday_mode
	// https://open.shopee.com/documents/v2/v2.shop.set_shop_holiday_mode?module=92&type=1
	SetShopHolidayMode(sid uint64, req SetShopHolidayModeRequest, tok string) (*SetShopHolidayModeResponse, error)
	// UpdateProfile This API support to let sellers to update the shop name, shop logo, and shop description.
	// Path: /api/v2/shop/update_profile
	// https://open.shopee.com/documents/v2/v2.shop.update_profile?module=92&type=1
	UpdateProfile(sid uint64, req UpdateProfileRequest, tok string) (*UpdateProfileResponse, error)
}

type ShopServiceOp[T any] struct {
	client *Client[T]
}

// GetAuthorisedResellerBrand Get the authorised reseller brand list for the shop.
// Path: /api/v2/shop/get_authorised_reseller_brand
// https://open.shopee.com/documents/v2/v2.shop.get_authorised_reseller_brand?module=92&type=1
func (s *ShopServiceOp[T]) GetAuthorisedResellerBrand(sid uint64, opt GetAuthorisedResellerBrandRequest, tok string) (*GetAuthorisedResellerBrandResponse, error) {
	path := "/shop/get_authorised_reseller_brand"
	resp := new(GetAuthorisedResellerBrandResponse)
	err := s.client.WithShop(sid, tok).Get(path, resp, opt)
	return resp, err
}

// GetBrShopOnboardingInfo [For BR Shop Only] Use this API to get shop KYC registration and qualification information.
// Path: /api/v2/shop/get_br_shop_onboarding_info
// https://open.shopee.com/documents/v2/v2.shop.get_br_shop_onboarding_info?module=92&type=1
func (s *ShopServiceOp[T]) GetBrShopOnboardingInfo(sid uint64, tok string) (*GetBrShopOnboardingInfoResponse, error) {
	path := "/shop/get_br_shop_onboarding_info"
	resp := new(GetBrShopOnboardingInfoResponse)
	err := s.client.WithShop(sid, tok).Get(path, resp, nil)
	return resp, err
}

// GetProfile This API support to get information of shop.
// Path: /api/v2/shop/get_profile
// https://open.shopee.com/documents/v2/v2.shop.get_profile?module=92&type=1
func (s *ShopServiceOp[T]) GetProfile(sid uint64, tok string) (*GetProfileResponse, error) {
	path := "/shop/get_profile"
	resp := new(GetProfileResponse)
	err := s.client.WithShop(sid, tok).Get(path, resp, nil)
	return resp, err
}

// GetShopHolidayMode Use this API to check whether a shop has enabled holiday mode and its ongoing and upcoming holiday mode period.
// Path: /api/v2/shop/get_shop_holiday_mode
// https://open.shopee.com/documents/v2/v2.shop.get_shop_holiday_mode?module=92&type=1
func (s *ShopServiceOp[T]) GetShopHolidayMode(sid uint64, tok string) (*GetShopHolidayModeResponse, error) {
	path := "/shop/get_shop_holiday_mode"
	resp := new(GetShopHolidayModeResponse)
	err := s.client.WithShop(sid, tok).Get(path, resp, nil)
	return resp, err
}

// GetShopInfo {"content":"<p>Use this call to get information of shop</p>","raw_content":[{"name":"paragraph","children":[{"data":"Use this call to get information of shop"}]}]}
// Path: /api/v2/shop/get_shop_info
// https://open.shopee.com/documents/v2/v2.shop.get_shop_info?module=92&type=1
func (s *ShopServiceOp[T]) GetShopInfo(sid uint64, tok string) (*GetShopInfoResponse, error) {
	path := "/shop/get_shop_info"
	resp := new(GetShopInfoResponse)
	err := s.client.WithShop(sid, tok).Get(path, resp, nil)
	return resp, err
}

// GetShopNotification get Seller Center notification, the permission is controlled by App type
// Path: /api/v2/shop/get_shop_notification
// https://open.shopee.com/documents/v2/v2.shop.get_shop_notification?module=92&type=1
func (s *ShopServiceOp[T]) GetShopNotification(sid uint64, opt GetShopNotificationRequest, tok string) (*GetShopNotificationResponse, error) {
	path := "/shop/get_shop_notification"
	resp := new(GetShopNotificationResponse)
	err := s.client.WithShop(sid, tok).Get(path, resp, opt)
	return resp, err
}

// GetWarehouseDetail For given shop id and region, return warehouse info including warehouse id, address id and location id, return all warehouse with once call.
// Path: /api/v2/shop/get_warehouse_detail
// https://open.shopee.com/documents/v2/v2.shop.get_warehouse_detail?module=92&type=1
func (s *ShopServiceOp[T]) GetWarehouseDetail(sid uint64, opt GetWarehouseDetailRequest, tok string) (*GetWarehouseDetailResponse, error) {
	path := "/shop/get_warehouse_detail"
	resp := new(GetWarehouseDetailResponse)
	err := s.client.WithShop(sid, tok).Get(path, resp, opt)
	return resp, err
}

// SetShopHolidayMode Use this API to set holiday periods in advance for automatic on/off of holiday mode and there are two holiday modes allowing sellers to choose whether to accept new orders during holiday.
// Path: /api/v2/shop/set_shop_holiday_mode
// https://open.shopee.com/documents/v2/v2.shop.set_shop_holiday_mode?module=92&type=1
func (s *ShopServiceOp[T]) SetShopHolidayMode(sid uint64, req SetShopHolidayModeRequest, tok string) (*SetShopHolidayModeResponse, error) {
	path := "/shop/set_shop_holiday_mode"
	resp := new(SetShopHolidayModeResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

// UpdateProfile This API support to let sellers to update the shop name, shop logo, and shop description.
// Path: /api/v2/shop/update_profile
// https://open.shopee.com/documents/v2/v2.shop.update_profile?module=92&type=1
func (s *ShopServiceOp[T]) UpdateProfile(sid uint64, req UpdateProfileRequest, tok string) (*UpdateProfileResponse, error) {
	path := "/shop/update_profile"
	resp := new(UpdateProfileResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}
