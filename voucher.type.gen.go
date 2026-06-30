package goshopee

type AddVoucherRequest struct {
	DiscountAmount *float64 `json:"discount_amount,omitempty"` // [Optional] The discount amount set for this particular voucher. Only fill in when you are creating a fix amount voucher.
	DisplayChannelList []int64 `json:"display_channel_list,omitempty"` // [Optional] <p>The FE channel where the voucher will be displayed. The available values are: 1: display_all&nbsp; 3: feed, 4: live streaming,   [] (empty - which is hidden).</p>
	DisplayStartTime *int64 `json:"display_start_time,omitempty"` // [Optional] <p>The timing of when voucher is displayed on shop pages; so buyer is allowed to claim.display_start_time must be left empty if the&nbsp;display_channel_list is empty (when voucher is hidden), otherwise it will show error.</p>
	EndTime int64 `json:"end_time"` // [Required] The timing until when the voucher is still valid. Any time after this end_time, buyer is not allowed to claim or to use. 
	ItemIdList []int64 `json:"item_id_list,omitempty"` // [Optional] The list of items which is applicable for the voucher. Only fill in when you are creating a product type of voucher.
	MaxPrice *float64 `json:"max_price,omitempty"` // [Optional] <p>The max amount of discount/value a user can enjoy by using this particular voucher. Only fill in when you are creating a discount percentage voucher or coins cashback voucher.If no cap limit, can set to be 0.</p>
	MinBasketPrice float64 `json:"min_basket_price"` // [Required] The minimum spend required for using this voucher. 
	Percentage *int64 `json:"percentage,omitempty"` // [Optional] The discount percentage set for this particular voucher. Only fill in when you are creating a discount percentage voucher or coins cashback voucher.
	RewardType int64 `json:"reward_type"` // [Required] The reward type of the voucher. The available values are: 1: fix_amount voucher, 2: discount_percentage voucher, 3: coin_cashback voucher.
	StartTime int64 `json:"start_time"` // [Required] <p>The timing from when the voucher is valid; so buyer is allowed to claim and to use. Field can only be updated if voucher has not started.</p>
	UsageQuantity int64 `json:"usage_quantity"` // [Required] The number of times for this particular voucher could be used.
	VoucherCode string `json:"voucher_code"` // [Required] The code of the voucher.
	VoucherName string `json:"voucher_name"` // [Required] The name of the voucher.
	VoucherType int64 `json:"voucher_type"` // [Required] The type of the voucher. The available values are: 1: shop voucher, 2: product voucher.
}
type AddVoucherResponse struct {
	BaseResponse // Common response fields
	Response AddVoucherResponseData `json:"response"` // Detailed informations you are querying.
}
type AddVoucherResponseData struct {
	VoucherId int64 `json:"voucher_id"` // [Required] The unique identifier for the created voucher.
}
type DeleteVoucherRequest struct {
	VoucherId int64 `json:"voucher_id"` // [Required] The unique identifier for the voucher you want to delete.
}
type DeleteVoucherResponse struct {
	BaseResponse // Common response fields
	Response DeleteVoucherResponseData `json:"response"` // Detail informations you are querying.
}
type DeleteVoucherResponseData struct {
	VoucherId int64 `json:"voucher_id"` // [Required] The unique identifier for the voucher it is being deleted.
}
type EndVoucherRequest struct {
	VoucherId int64 `json:"voucher_id"` // [Required] The unique identifier for the voucher you want to end now.
}
type EndVoucherResponse struct {
	BaseResponse // Common response fields
	Response EndVoucherResponseData `json:"response"` // Detailed informations you are querying.
}
type EndVoucherResponseData struct {
	VoucherId int64 `json:"voucher_id"` // [Required] The unique identifier for the voucher it is being ended.
}
type GetVoucherListRequest struct {
	PageNo *int64 `json:"page_no,omitempty"` // [Optional] Specifies the page number of data to return in the current call. Default to be 1 and allowed input is from 1 - 5000.
	PageSize *int64 `json:"page_size,omitempty"` // [Optional] Use the 'page_size' filters to control the maximum number of entries to retrieve per page (i.e., per call). Default to be 20 and allowed input is from 1- 100.
	Status string `json:"status"` // [Required] The status filter for retrieving voucher list. Available value: upcoming/ongoing/expired/all.
}
type GetVoucherListResponse struct {
	BaseResponse // Common response fields
	Response GetVoucherListResponseData `json:"response"` // Detailed informations you are querying.
}
type GetVoucherListResponseData struct {
	More bool `json:"more"` // [Required] This is to indicate whether the comment list is more than one page. If this value is true, you may want to continue to check next page to retrieve the rest of comments.
	VoucherList []Voucher `json:"voucher_list"` // [Required] The list of voucher.
}
type GetVoucherRequest struct {
	VoucherId int64 `json:"voucher_id"` // [Required] The unique identifier of a voucher used to query the voucher details.
}
type GetVoucherResponse struct {
	BaseResponse // Common response fields
	Response GetVoucherResponseData `json:"response"` // Detailed informations you are querying.
}
type GetVoucherResponseData struct {
	VoucherId int64 `json:"voucher_id"` // [Required] The unique identifier of the voucher whose details are returned.
	VoucherCode string `json:"voucher_code"` // [Required] Voucher Code
	VoucherName string `json:"voucher_name"` // [Required] Voucher Name
	VoucherType int64 `json:"voucher_type"` // [Required] The type of the voucher. The available values are: 1: shop voucher, 2: product voucher.
	RewardType int64 `json:"reward_type"` // [Required] The reward type of the voucher. The available values are: 1: fix_amount voucher, 2: discount_percentage voucher, 3: coin_cashback voucher.
	UsageQuantity int64 `json:"usage_quantity"` // [Required] The number of times for this particular voucher could be used.
	CurrentUsage int64 `json:"current_usage"` // [Required] Up till now, how many times has this particular voucher already been used.
	StartTime int64 `json:"start_time"` // [Required] The timing from when the voucher is valid; so buyer is allowed to claim and to use. 
	EndTime int64 `json:"end_time"` // [Required] The timing until when the voucher is still valid. Any time after this end_time, buyer is not allowed to claim or to use. 
	IsAdmin bool `json:"is_admin"` // [Required] If the voucher is created by Shopee or not.
	VoucherPurpose int64 `json:"voucher_purpose"` // [Required] <p>The use case for the voucher. The available values are: 0: normal; 1: welcome, 2: referral; 3: shop_follow; 4:shop_game, 5: free_gift, 6: membership，7: Ads</p>
	DisplayChannelList []int64 `json:"display_channel_list"` // [Required] The FE channel where the voucher will be displayed. The available values are: 1: display_all, 2: order page, 3: feed, 4: live streaming,   [] (empty - which is hidden).
	MinBasketPrice float64 `json:"min_basket_price"` // [Required] The minimum spend required for using this voucher. 
	Percentage int64 `json:"percentage"` // [Required] <p>The discount percentage is set for this particular voucher. Only when it is a discount percentage voucher or coins cashback voucher, api will return a value.</p>
	MaxPrice float64 `json:"max_price"` // [Required] The max amount of discount/value a user can enjoy by using this particular voucher. Only when it is a discount percentage voucher or coins cashback voucher, api will return a value.
	DiscountAmount float64 `json:"discount_amount"` // [Required] The discount amount set for this particular voucher. Only when it is a fix amount voucher, api will return a value.
	CmtVoucherStatus int64 `json:"cmt_voucher_status"` // [Required] The voucher status in CMT. The available values are: 1:review, 2: approved, 3:reject. Only when this voucher is attending CMT campaign and not being rejected, api will return a value.
	ItemIdList []int64 `json:"item_id_list"` // [Required] The list of items which is applicable for the voucher. Only return a value when it is a product type of voucher.
	DisplayStartTime int64 `json:"display_start_time"` // [Required] The timing of when voucher is displayed on shop pages; so buyer is allowed to claim.
	TargetVoucher int64 `json:"target_voucher"` // [Required] <p>The field is used to mark new user voucher/ repeat buyer voucher&nbsp;</p><p>1: new user voucher&nbsp;</p><p>2: repeat buyer voucher: with 1 orders</p><p>3. repeat buyer voucher: with 2 orders<br /></p>
	Usecase int64 `json:"usecase"` // [Required] <p>usecase indicates a specific business scenario that the voucher is created and used for, i.e., new buyer voucher, live voucher, follow shop voucher, etc.</p><p>shop voucher:1</p><p>product voucher:2</p><p>new buyer voucher:3</p><p>repeat buyer voucher:4</p><p>private voucher:5</p><p>live voucher:6</p><p>video voucher:7</p><p>campaign voucher:8</p><p>follow prize voucher:9</p><p>membership voucher:10</p><p>game prize voucher:11</p><p>sample voucher:12<br /></p>
}
type UpdateVoucherRequest struct {
	DiscountAmount *float64 `json:"discount_amount,omitempty"` // [Optional] The discount amount set for this particular voucher. Only fill in when you are updating a fix amount voucher.
	DisplayChannelList []int64 `json:"display_channel_list,omitempty"` // [Optional] The FE channel where the voucher will be displayed. The available values are: 1: display_all, 2: order page, 3: feed, 4: live streaming,   [] (empty - which is hidden).
	DisplayStartTime *int64 `json:"display_start_time,omitempty"` // [Optional] The timing of when voucher is displayed on shop pages; so buyer is allowed to claim.
	EndTime *int64 `json:"end_time,omitempty"` // [Optional] The timing until when the voucher is still valid. Any time after this end_time, buyer is not allowed to claim or to use. 
	ItemIdList []int64 `json:"item_id_list,omitempty"` // [Optional] The list of items which is applicable for the voucher. Only fill in when you are updating a product type of voucher.
	MaxPrice *float64 `json:"max_price,omitempty"` // [Optional] The max amount of discount/value a user can enjoy by using this particular voucher. Only fill in when you are updating a discount percentage voucher or coins cashback voucher.
	MinBasketPrice *float64 `json:"min_basket_price,omitempty"` // [Optional] The minimum spend required for using this voucher. 
	Percentage *int64 `json:"percentage,omitempty"` // [Optional] The discount percentage set for this particular voucher. Only fill in when you are updating a discount percentage voucher or coins cashback voucher.
	StartTime *int64 `json:"start_time,omitempty"` // [Optional] <p>The timing from when the voucher is valid; so buyer is allowed to claim and to use. Field can only be updated if voucher has not started.</p>
	UsageQuantity *int64 `json:"usage_quantity,omitempty"` // [Optional] The number of times for this particular voucher could be used.
	VoucherId int64 `json:"voucher_id"` // [Required] The unique identifier of the voucher which is going to be updated.
	VoucherName *string `json:"voucher_name,omitempty"` // [Optional] The name of the voucher
}
type UpdateVoucherResponse struct {
	BaseResponse // Common response fields
	Response UpdateVoucherResponseData `json:"response"` // Detailed informations you are querying.
}
type UpdateVoucherResponseData struct {
	VoucherId int64 `json:"voucher_id"` // [Required] The unique identifier of the voucher which is being updated.
}
type Voucher struct {
	VoucherId int64 `json:"voucher_id"` // [Required] The unique identifier for a voucher.
	VoucherCode string `json:"voucher_code"` // [Required] Voucher Code
	VoucherName string `json:"voucher_name"` // [Required] Voucher Name
	VoucherType int64 `json:"voucher_type"` // [Required] The type of the voucher. The available values are: 1: shop voucher, 2: product voucher.
	RewardType int64 `json:"reward_type"` // [Required] The reward type of the voucher. The available values are: 1: fix_amount voucher, 2: discount_percentage voucher, 3: coin_cashback voucher.
	UsageQuantity int64 `json:"usage_quantity"` // [Required] The number of times for this particular voucher could be used.
	CurrentUsage int64 `json:"current_usage"` // [Required] Up till now, how many times has this particular voucher already been used.
	StartTime int64 `json:"start_time"` // [Required] The timing from when the voucher is valid; so buyer is allowed to claim and to use. 
	EndTime int64 `json:"end_time"` // [Required] The timing until when the voucher is still valid. Any time after this end_time, buyer is not allowed to claim or to use. 
	IsAdmin bool `json:"is_admin"` // [Required] If the voucher is created by Shopee or not.
	VoucherPurpose int64 `json:"voucher_purpose"` // [Required] The use case for the voucher. The available values are: 0: normal; 1: welcome, 2: referral; 3: shop_follow; 4:shop_game, 5: free_gift, 6: membership
	DiscountAmount float64 `json:"discount_amount"` // [Required] The discount amount set for this particular voucher. Only when it is a fix amount voucher, api will return a value.
	Percentage int64 `json:"percentage"` // [Required] The discount percentage set for this particular voucher. Only when it is a discount percentage voucher or coins cashback voucher, api will return a value.
	CmtVoucherStatus int64 `json:"cmt_voucher_status"` // [Required] The voucher status in CMT. The available values are: 1:review, 2: approved, 3:reject. Only when this voucher is attending CMT campaign and not being rejected, api will return a value.
	DisplayStartTime int64 `json:"display_start_time"` // [Required] The timing of when voucher is displayed on shop pages; so buyer is allowed to claim.
}
