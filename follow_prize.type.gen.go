package goshopee

type AddFollowPrizeRequest struct {
	DiscountAmount *float64 `json:"discount_amount,omitempty"` // [Optional] <p>The discount amount set for this particular follow prize.Only fill in when you are creating a fix amount follow prize.<br /></p>
	EndTime int64 `json:"end_time"` // [Required] <p>The timing until when the follow prize is still valid,the end time must be greater than the start time by at least 1 day and end time cannot exceed 3 months after start time.If the start time and end time passed in by the seller overlap with other upcoming/ongoing activities, it will prompt "Another Follow Prize voucher already exists during this time period, please set another period."<br /></p>
	FollowPrizeName string `json:"follow_prize_name"` // [Required] <p>The name of the follow prize,The follow prize name length max limit is 20.<br /></p>
	MaxPrice *float64 `json:"max_price,omitempty"` // [Optional] <p>The max amount of discount/value a user can enjoy by using this particular follow prize. It is required to fill in when you are creating a discount percentage follow prize or coins cashback follow prize. max_price &gt;=1<br /></p>
	MinSpend float64 `json:"min_spend"` // [Required] <p>The minimum spend required for using this follow prize.<br /></p>
	Percentage *int64 `json:"percentage,omitempty"` // [Optional] <p>The discount percentage set for this particular follow prize. Only fill in when you are creating a discount percentage follow prize or coins cashback follow prize.Discount percentage (reward_type ==2) or Percentage of coins cash back (reward_type==3).<br /></p>
	RewardType int64 `json:"reward_type"` // [Required] <p>The reward type of the follow prize.The available values are:1:discount---fix amount,2:discount---by percentage,3:coin cash back.<br /></p>
	StartTime int64 `json:"start_time"` // [Required] <p>The timing from when the follow prize is valid,the start time later than the current time.If the start time and end time passed in by the seller overlap with other upcoming/ongoing activities, it will prompt "Another Follow Prize voucher already exists during this time period, please set another period."<br /></p>
	UsageQuantity int64 `json:"usage_quantity"` // [Required] <p>Please enter a value between 1 and 200000.</p>
}
type AddFollowPrizeResponse struct {
	BaseResponse // Common response fields
	Response AddFollowPrizeResponseData `json:"response"` // <p>Detailed informations you are querying.<br /></p>
}
type AddFollowPrizeResponseData struct {
	CampaginId int64 `json:"campagin_id"` // [Required] <p>The unique identifier for the created follow prize.<br /></p>
}
type DeleteFollowPrizeRequest struct {
	CampaignId int64 `json:"campaign_id"` // [Required] <p>The unique identifier for the created follow prize.<br /></p>
}
type DeleteFollowPrizeResponse struct {
	BaseResponse // Common response fields
	Response DeleteFollowPrizeResponseData `json:"response"` // <p>Detailed informations you are querying.<br /></p>
}
type DeleteFollowPrizeResponseData struct {
	CampaginId int64 `json:"campagin_id"` // [Required] <p>The unique identifier for the created follow prize.<br /></p>
}
type EndFollowPrizeRequest struct {
	CampaignId int64 `json:"campaign_id"` // [Required] <p>The unique identifier for the created follow prize.<br /></p>
}
type EndFollowPrizeResponse struct {
	BaseResponse // Common response fields
	Response EndFollowPrizeResponseData `json:"response"` // <p>Detailed informations you are querying.<br /></p>
}
type EndFollowPrizeResponseData struct {
	CampaignId int64 `json:"campaign_id"` // [Required] <p>The unique identifier for the created follow prize.<br /></p>
}
type FollowPrize struct {
	CampaignId int64 `json:"campaign_id"` // [Required] <p>The unique identifier for the created follow prize.<br /></p>
	CampaignStatus CampaignStatus `json:"campaign_status"` // [Required] <p>The status of follow prize,the campagin status have upcoming/ongoing/expired.<br /></p>
	FollowPrizeName string `json:"follow_prize_name"` // [Required] <p>The name of the follow prize,The follow prize name length max limit is 20.<br /></p>
	StartTime int64 `json:"start_time"` // [Required] <p>The timing from when the follow prize is valid,the start time later than the current time.If the start time and end time passed in by the seller overlap with other upcoming/ongoing activities, it will prompt "Another Follow Prize voucher already exists during this time period, please set another period."<br /></p>
	EndTime int64 `json:"end_time"` // [Required] <p>The timing until when the follow prize is still valid,the end time must be greater than the start time by at least 1 day and end time cannot exceed 3 months after start time.If the start time and end time passed in by the seller overlap with other upcoming/ongoing activities, it will prompt "Another Follow Prize voucher already exists during this time period, please set another period."<br /></p>
	UsageQuantity int64 `json:"usage_quantity"` // [Required] <p>Please enter a value between 1 and 200000.<br /></p>
	Claimed int64 `json:"claimed"` // [Required] <p>This is to indicate the quantity of voucher claimed.<br /></p>
}
type GetFollowPrizeDetailRequest struct {
	CampaignId *int64 `json:"campaign_id,omitempty"` // [Optional] <p>The unique identifier for the created follow prize.<br /></p>
}
type GetFollowPrizeDetailResponse struct {
	BaseResponse // Common response fields
	Response GetFollowPrizeDetailResponseData `json:"response"` // <p>Detailed informations you are querying.<br /></p>
}
type GetFollowPrizeDetailResponseData struct {
	CampaignStatus CampaignStatus `json:"campaign_status"` // [Required] <p>The status of follow prize,the campagin status have upcoming/ongoing/expired.<br /></p>
	CampaignId int64 `json:"campaign_id"` // [Required] <p>The unique identifier for the created follow prize.<br /></p>
	UsageQuantity int64 `json:"usage_quantity"` // [Required] <p>Please enter a value between 1 and 200000.<br /></p>
	StartTime int64 `json:"start_time"` // [Required] <p>The timing from when the follow prize is valid,the start time later than the current time.If the start time and end time passed in by the seller overlap with other upcoming/ongoing activities, it will prompt "Another Follow Prize voucher already exists during this time period, please set another period."<br /></p>
	EndTime int64 `json:"end_time"` // [Required] <p>The timing until when the follow prize is still valid,the end time must be greater than the start time by at least 1 day and end time cannot exceed 3 months after start time.If the start time and end time passed in by the seller overlap with other upcoming/ongoing activities, it will prompt "Another Follow Prize voucher already exists during this time period, please set another period."<br /></p>
	MinSpend int64 `json:"min_spend"` // [Required] <p>The minimum spend required for using this follow prize.<br /></p>
	RewardType int64 `json:"reward_type"` // [Required] <p>The reward type of the follow prize.The available values are:1:discount---fix amount,2:discount---by percentage,3:coin cash back.<br /></p>
	FollowPrizeName string `json:"follow_prize_name"` // [Required] <p>The name of the follow prize,The follow prize name length max limit is 20.<br /></p>
	DiscountAmount int64 `json:"discount_amount"` // [Required] <p>The discount amount set for this particular follow prize.Only fill in when you are creating a fix amount follow prize.<br /></p>
	Percentage int64 `json:"percentage"` // [Required] <p>The discount percentage set for this particular follow prize. Only fill in when you are creating a discount percentage follow prize or coins cashback follow prize.Discount percentage (reward_type ==2) or Percentage of coins cash back (reward_type==3).<br /></p>
	MaxPrice int64 `json:"max_price"` // [Required] <p>The max amount of discount/value a user can enjoy by using this particular follow prize. Only fill in when you are creating a discount percentage follow prize or coins cashback follow prize.<br /></p>
}
type GetFollowPrizeListRequest struct {
	PageNo *int64 `json:"page_no,omitempty"` // [Optional] <p>Specifies the page number of data to return in the current call. Default to be 1.<br /></p>
	PageSize *int64 `json:"page_size,omitempty"` // [Optional] <p>Use the 'page_size' filters to control the maximum number of entries to retrieve per page (i.e., per call). Default to be 20 and allowed input is from 1- 100.<br /></p>
	Status string `json:"status"` // [Required] <p>The status filter for retrieving follow prize list. Available value: upcoming/ongoing/expired/all.<br /></p>
}
type GetFollowPrizeListResponse struct {
	BaseResponse // Common response fields
	Response GetFollowPrizeListResponseData `json:"response"` // <p>Detailed informations you are querying.<br /></p>
}
type GetFollowPrizeListResponseData struct {
	More bool `json:"more"` // [Required] <p>This is to indicate whether the comment list is more than one page. If this value is true, you may want to continue to check next page to retrieve the rest of comments.<br /></p>
	FollowPrizeList *FollowPrize `json:"follow_prize_list"` // [Required] <p>The list of follow prize.<br /></p>
}
type UpdateFollowPrizeRequest struct {
	CampaignId int64 `json:"campaign_id"` // [Required] <p>The unique identifier for the created follow prize.<br /></p>
	EndTime *int64 `json:"end_time,omitempty"` // [Optional] <p>The timing until when the follow prize is still valid,the end time must be greater than the start time by at least 1 day and end time cannot exceed 3 months after start time.If the start time and end time passed in by the seller overlap with other upcoming/ongoing activities, it will prompt "Another Follow Prize voucher already exists during this time period, please set another period."<br /></p>
	FollowPrizeName *string `json:"follow_prize_name,omitempty"` // [Optional] <p>The name of the follow prize,The follow prize name length max limit is 20.<br /></p>
	MinSpend *float64 `json:"min_spend,omitempty"` // [Optional] <p>The minimum spend required for using this follow prize.<br /></p>
	StartTime *int64 `json:"start_time,omitempty"` // [Optional] <p>The timing from when the follow prize is valid,the start time later than the current time.If the start time and end time passed in by the seller overlap with other upcoming/ongoing activities, it will prompt "Another Follow Prize voucher already exists during this time period, please set another period."<br /></p>
	UsageQuantity *int64 `json:"usage_quantity,omitempty"` // [Optional] <p>Please enter a value between 1 and 200000.</p>
}
type UpdateFollowPrizeResponse struct {
	BaseResponse // Common response fields
	Response UpdateFollowPrizeResponseData `json:"response"` // <p>Detailed informations you are querying.<br /></p>
}
type UpdateFollowPrizeResponseData struct {
	CampaginId int64 `json:"campagin_id"` // [Required] <p>The unique identifier for the created follow prize.<br /></p>
}
