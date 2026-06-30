package goshopee

type AddAllProductsToOpenCampaignRequest struct {
	CommissionRate float64 `json:"commission_rate"` // [Required] <p>Commission Rate, 1.1 means 1.1%, support two decimal places</p>
	PeriodEndTime *int64 `json:"period_end_time,omitempty"` // [Optional] <p>Period end time, in seconds, if missing, will set&nbsp;<span style="font-size:14px;">32503651199 (</span><span style="font-size:14px;">2999-12-31 23:59:59</span><span style="font-size:14px;">) represent of no limit</span></p>
	PeriodStartTime *int64 `json:"period_start_time,omitempty"` // [Optional] <p>Period start time, in seconds, if missing, will set&nbsp;10 minutes later</p>
}
type AddAllProductsToOpenCampaignResponse struct {
	BaseResponse // Common response fields
	Response AddAllProductsToOpenCampaignResponseData `json:"response"` // 
}
type AddAllProductsToOpenCampaignResponseData struct {
	TaskType string `json:"task_type"` // [Required] <p>Task type. Applicable values:&nbsp;</p><p>batch_add_open_campaigns</p><p>batch_remove_open_campaigns</p><p>batch_update_open_campaigns</p><p><br /></p><p>For this API, task type will be&nbsp;batch_add_open_campaigns</p>
	TaskId string `json:"task_id"` // [Required] <p>Task id, used to query task progress when calling v2.ams.get_open_campaign_batch_task_result API</p>
}
type Affiliate struct {
	AffiliateId int64 `json:"affiliate_id"` // [Required] <p>The unique key for affiliate, can call v2.ams.query_affiliate_list to get affiliate details.</p>
}
type AMSGetShopPerformanceResponse struct {
	BaseResponse // Common response fields
	Response AMSGetShopPerformanceResponseData `json:"response"` // 
}
type AMSGetShopPerformanceResponseData struct {
	Sales string `json:"sales"` // [Required] <p>Total value of orders generated through affiliate marketing during the selected period.</p>
	GrossItemSold int64 `json:"gross_item_sold"` // [Required] <p>Total number of items sold through affiliate marketing during the selected period.</p>
	Orders int64 `json:"orders"` // [Required] <p>Total number of orders generated through affiliate marketing during the selected period.</p>
	Clicks int64 `json:"clicks"` // [Required] <p>Total clicks on your product links through affiliate marketing during the selected period.</p>
	EstCommission string `json:"est_commission"` // [Required] <p>Estimated total payout from your affiliate marketing orders.</p>
	Roi string `json:"roi"` // [Required] <p>Return on Investment, equal to Sales divided by Est. Commission. It can be used to evaluate the efficiency of your investment in affiliate marketing. If it does not exist, the return value is --.</p>
	TotalBuyers int64 `json:"total_buyers"` // [Required] <p>Total number of buyers who order from your shop through affiliate marketing.</p>
	NewBuyers int64 `json:"new_buyers"` // [Required] <p>Total number of new buyers who order from your shop through affiliate marketing.</p>
	FetchedDateRange string `json:"fetched_date_range"` // [Required] <p>Effective query date range. Invalid input ranges will be automatically shifted.</p>
}
type BatchAddProductsToOpenCampaignRequest struct {
	CommissionRate float64 `json:"commission_rate"` // [Required] <p>Commission Rate, 1.1 means 1.1%, support two decimal places</p>
	ItemIdList []int64 `json:"item_id_list"` // [Required] <p>The list of item_id,&nbsp;max limit: 50</p>
	PeriodEndTime *int64 `json:"period_end_time,omitempty"` // [Optional] <p>Period end time, in seconds, if missing, will set&nbsp;32503651199 (2999-12-31 23:59:59) represent of no limit</p>
	PeriodStartTime *int64 `json:"period_start_time,omitempty"` // [Optional] <p>Period start time, in seconds, if missing, will set 10 minutes later</p>
}
type BatchAddProductsToOpenCampaignResponse struct {
	BaseResponse // Common response fields
	Response BatchAddProductsToOpenCampaignResponseData `json:"response"` // 
}
type BatchAddProductsToOpenCampaignResponseData struct {
	FailedList []Failed `json:"failed_list"` // [Required] 
	SuccessList []int64 `json:"success_list"` // [Required] <p>Success Item ID List</p>
}
type BatchEditProductsOpenCampaignSettingRequest struct {
	CampaignIds []int64 `json:"campaign_ids"` // [Required] <p>The list of campaign_id,&nbsp;max limit: 50</p>
	CommissionRate *float64 `json:"commission_rate,omitempty"` // [Optional] <p>Commission Rate, 1.1 means 1.1%, support two decimal places</p>
	PeriodEndTime *int64 `json:"period_end_time,omitempty"` // [Optional] <p>Period end time, in seconds, if missing, will skip and do not update</p><p>Can set 32503651199 to make period no limit</p>
	PeriodStartTime *int64 `json:"period_start_time,omitempty"` // [Optional] <p>Period start time, in seconds, if missing, will skip and do not update</p><p>Only allow to update on UPCOMING status, when in other status, will skip too</p>
}
type BatchEditProductsOpenCampaignSettingResponse struct {
	BaseResponse // Common response fields
	Response BatchEditProductsOpenCampaignSettingResponseData `json:"response"` // 
}
type BatchEditProductsOpenCampaignSettingResponseData struct {
	FailedList []ResponseDataFailed `json:"failed_list"` // [Required] 
	SuccessList []int64 `json:"success_list"` // [Required] <p>Success Campaign ID List</p>
}
type BatchGetProductsSuggestedRateRequest struct {
	ItemIdList string `json:"item_id_list" url:"item_id_list"` // [Required] <p>The list of item_id, different item id should be split by comma and at most 20 items</p>
}
type BatchGetProductsSuggestedRateResponse struct {
	BaseResponse // Common response fields
	Response BatchGetProductsSuggestedRateResponseData `json:"response"` // <p>Detail informations you are querying.</p>
}
type BatchGetProductsSuggestedRateResponseData struct {
	Rates []Rates `json:"rates"` // [Required] 
}
type BatchRemoveProductsOpenCampaignSettingRequest struct {
	CampaignIds []int64 `json:"campaign_ids"` // [Required] <p>The list of campaign_id,&nbsp;max limit: 50</p>
}
type BatchRemoveProductsOpenCampaignSettingResponse struct {
	BaseResponse // Common response fields
	Response BatchRemoveProductsOpenCampaignSettingResponseData `json:"response"` // 
}
type BatchRemoveProductsOpenCampaignSettingResponseData struct {
	FailedList *ResponseDataFailed `json:"failed_list"` // [Required] 
	SuccessList []int64 `json:"success_list"` // [Required] <p>Success Campaign ID List</p>
}
type CommissionProtection struct {
	CommissionRate float64 `json:"commission_rate"` // [Required] <p>Commission Rate,&nbsp;1.1 means 1.1%, support two decimal places.</p>
	ProtectionPeriodEndTime int64 `json:"protection_period_end_time"` // [Required] <p>Protection Period End Time.</p>
}
type CreateNewTargetedCampaignRequest struct {
	AffiliateList []Affiliate `json:"affiliate_list"` // [Required] <p>The&nbsp;list of affiliates associated with the current campaign.</p>
	Budget *float64 `json:"budget,omitempty"` // [Optional] <p>Budget value set for the current campaign.</p><p><br /></p><p>Note: TH not supported</p>
	CampaignName string `json:"campaign_name"` // [Required] <p>The name of the current campaign.</p>
	IsSetBudget *bool `json:"is_set_budget,omitempty"` // [Optional] <p>Budget allocation toggle for the current campaign.</p><p><br /></p><p>Note: TH not supported</p>
	ItemList []Item `json:"item_list"` // [Required] <p>The&nbsp;list of items associated with the current campaign.</p>
	PeriodEndTime int64 `json:"period_end_time"` // [Required] <p>The period end time of campaign, in seconds.</p><p>Can set&nbsp;32503651199 (2999-12-31 23:59:59) represent of no limit.</p>
	PeriodStartTime int64 `json:"period_start_time"` // [Required] <p>The period start time of campaign, in seconds.</p>
	SellerMessage string `json:"seller_message"` // [Required] <p>The message displayed to affiliates.</p>
}
type CreateNewTargetedCampaignResponse struct {
	BaseResponse // Common response fields
	Response CreateNewTargetedCampaignResponseData `json:"response"` // <p>Successful elements will not be returned.&nbsp;</p><p>If error and message are not empty, it means that all settings have failed.&nbsp;</p><p>If error and message are empty,&nbsp;but fail_xx_list has content, it means that the element in the fail_xx_list setting failed and the rest were successful.</p><p>If error, message, fail_xx_list are all empty, it means that all settings are successful.</p>
}
type CreateNewTargetedCampaignResponseData struct {
	CampaignId int64 `json:"campaign_id"` // [Required] <p>The unique key for campaign.</p>
	FailItemList []Failed `json:"fail_item_list"` // [Required] <p>Failed Item List.</p>
	FailAffiliateList []FailAffiliate `json:"fail_affiliate_list"` // [Required] <p>Fail Affiliate List.</p>
}
type EditAffiliateListOfTargetedCampaignRequest struct {
	AffiliateList []Affiliate `json:"affiliate_list"` // [Required] <p>The list of affiliates to be modified.</p>
	CampaignId int64 `json:"campaign_id"` // [Required] <p>Campaign id for update.</p>
	EditType string `json:"edit_type"` // [Required] <p>Edit type. Applicable values:</p><p>add<br />delete<br /></p>
}
type EditAffiliateListOfTargetedCampaignResponse struct {
	BaseResponse // Common response fields
	Response EditAffiliateListOfTargetedCampaignResponseData `json:"response"` // <p>Successful elements will not be returned.&nbsp;</p><p>If error and message are not empty, it means that all settings have failed.&nbsp;</p><p>If error and message are empty,&nbsp;but fail_xx_list has content, it means that the element in the fail_xx_list setting failed and the rest were successful.</p><p>If error, message, fail_xx_list are all empty, it means that all settings are successful.</p>
}
type EditAffiliateListOfTargetedCampaignResponseData struct {
	FailAffiliateList []FailAffiliate `json:"fail_affiliate_list"` // [Required] <p>Failed Affiliate List.</p>
}
type EditAllProductsOpenCampaignSettingRequest struct {
	CommissionRate *float64 `json:"commission_rate,omitempty"` // [Optional] <p>Commission Rate, 1.1 means 1.1%, support two decimal places, if miss, will skip and do not update</p>
	PeriodEndTime *int64 `json:"period_end_time,omitempty"` // [Optional] <p>Period end time, in seconds, if missing, will skip and do not update</p><p>Can set 32503651199 to make period no limit</p>
	PeriodStartTime *int64 `json:"period_start_time,omitempty"` // [Optional] <p>Period start time, in seconds, if missing, will skip and do not update</p><p>Only allow to update on UPCOMING status, when in other status, will skip too</p>
}
type EditAllProductsOpenCampaignSettingResponse struct {
	BaseResponse // Common response fields
	Response EditAllProductsOpenCampaignSettingResponseData `json:"response"` // 
}
type EditAllProductsOpenCampaignSettingResponseData struct {
	TaskType string `json:"task_type"` // [Required] <p>Task type. Applicable values:&nbsp;</p><p>batch_add_open_campaigns</p><p>batch_remove_open_campaigns</p><p>batch_update_open_campaigns</p><p><br /></p><p>For this API, task type will be&nbsp;batch_update_open_campaigns</p>
	TaskId string `json:"task_id"` // [Required] <p>Task id, used to query task progress when calling v2.ams.get_open_campaign_batch_task_result API</p>
}
type EditProductListOfTargetedCampaignRequest struct {
	CampaignId int64 `json:"campaign_id"` // [Required] <p>Campaign id for update.</p>
	EditType string `json:"edit_type"` // [Required] <p>Edit type. Applicable values:&nbsp;</p><p>add<br />delete<br />update<br /></p>
	ItemList []Item `json:"item_list"` // [Required] <p>The list of items to be modified.</p>
}
type EditProductListOfTargetedCampaignResponse struct {
	BaseResponse // Common response fields
	Response EditProductListOfTargetedCampaignResponseData `json:"response"` // <p>Successful elements will not be returned.&nbsp;</p><p>If error and message are not empty, it means that all settings have failed.&nbsp;</p><p>If error and message are empty,&nbsp;but fail_xx_list has content, it means that the element in the fail_xx_list setting failed and the rest were successful.</p><p>If error, message, fail_xx_list are all empty, it means that all settings are successful.</p>
}
type EditProductListOfTargetedCampaignResponseData struct {
	FailItemList []Failed `json:"fail_item_list"` // [Required] <p>Failed Item List.</p>
}
type FailAffiliate struct {
	AffiliateId int64 `json:"affiliate_id"` // [Required] <p>The unique key for affiliate.</p>
	FailError string `json:"fail_error"` // [Required] <p>Indicate error type if hit error. Empty if no error happened.</p>
	FailMessage string `json:"fail_message"` // [Required] <p>Indicate error details if hit error. Empty if no error happened.</p>
}
type GetAffiliatePerformanceRequest struct {
	AffiliateId *int64 `json:"affiliate_id,omitempty" url:"affiliate_id,omitempty"` // [Optional] <p>Affiliate ID for query.</p>
	Channel string `json:"channel" url:"channel"` // [Required] <p>Channel. Applicable values:&nbsp;<br />- AllChannel<br />- SocialMedia<br />- ShopeeVideo<br />- LiveStreaming</p>
	EndDate string `json:"end_date" url:"end_date"` // [Required] <p>The end_date must be:&nbsp;</p><p>- Equal to start_date for "Day" period type</p><p>- Saturday for "Week" period type</p><p>- The last day of a Month for "Month" period type. If the selected month is the current month, the end_date should be the latest data date</p><p>- The latest data date for "Last7d" period type</p><p>- The latest data date for "Last30d" period type</p><p><br /></p><p>Note:&nbsp;</p><p>- The end_date must be later than the start_date and earlier than the latest data date</p><p>- The latest data date can be obtained by using "AmsMarker" in the v2.ams.get_performance_data_update_time API.</p>
	OrderType string `json:"order_type" url:"order_type"` // [Required] <p>Order Type. Applicable values:&nbsp;</p><p>PlacedOrder<br />ConfirmedOrder:&nbsp;<br /><br />Note:&nbsp;</p><p>- Placed orders are orders (COD and non-COD) that buyers have successfully placed, including paid and unpaid orders.</p><p>- Confirmed orders are either non-COD orders that have been paid for or COD orders that have been confirmed for shipping (usually 30 mins after placing the order).</p>
	PageNo int64 `json:"page_no" url:"page_no"` // [Required] <p>Specifies the page number of data to return in the current call. Starting from 1. if data is more than one page, the page_no can be some entry to start next call.</p>
	PageSize int64 `json:"page_size" url:"page_size"` // [Required] <p>Each result set is returned as a page of entries. Use the "page_size" filters to control the maximum number of entries to retrieve per page (i.e., per call). This integer value is used to specify the maximum number of entries to return in a single "page" of data. The limit of page_size if between 1 and 20.</p>
	PeriodType string `json:"period_type" url:"period_type"` // [Required] <p>Period Type. Applicable values:<br />Day<br />Week<br />Month<br />Last7d<br />Last30d<br /><br /></p><p>Note: The start date and end date must align with the Period Type.</p>
	StartDate string `json:"start_date" url:"start_date"` // [Required] <p>The start_date must be:</p><p>- Any day in the past three calendar months for "Day" period type</p><p>- Sunday for "Week" period type</p><p>- The 1st day of a Month for "Month" period type</p><p>- The date that is 6 days prior to the latest data date for "Last7d" period type</p><p>- The date that is 29 days prior to the latest data date for "Last30d" period type</p><p><br /></p><p>Note: The latest data date can be obtained by using "AmsMarker" in the v2.ams.get_performance_data_update_time API.</p>
}
type GetAffiliatePerformanceResponse struct {
	BaseResponse // Common response fields
	Response GetAffiliatePerformanceResponseData `json:"response"` // 
}
type GetAffiliatePerformanceResponseData struct {
	List []List `json:"list"` // [Required] 
	TotalCount int64 `json:"total_count"` // [Required] <p>Total number of affiliates that match the condition.</p>
	HasMore bool `json:"has_more"` // [Required] <p>This is to indicate whether the list is more than one page. If this value is true, you may want to continue to check next page to retrieve the rest of datas.</p>
	FetchedDateRange string `json:"fetched_date_range"` // [Required] <p>Effective query date range. Invalid input ranges will be automatically shifted.</p>
}
type GetAutoAddNewProductToggleStatusResponse struct {
	BaseResponse // Common response fields
	Response GetAutoAddNewProductToggleStatusResponseData `json:"response"` // <p>Detail informations you are querying.</p>
}
type GetAutoAddNewProductToggleStatusResponseData struct {
	IsOpen bool `json:"is_open"` // [Required] <p>If auto-add new product is currently enabled</p>
	CommissionRate float64 `json:"commission_rate"` // [Required] <p>Commission Rate, 1.11 means 1.11%, support two decimal places</p>
}
type GetCampaignKeyMetricsPerformanceRequest struct {
	EndDate string `json:"end_date" url:"end_date"` // [Required] <p>The end_date must be:&nbsp;</p><p>- Equal to start_date for "Day" period type</p><p>- Saturday for "Week" period type</p><p>- The last day of a Month for "Month" period type. If the selected month is the current month, the end_date should be the latest data date</p><p>- The latest data date for "Last7d" period type</p><p>- The latest data date for "Last30d" period type</p><p><br /></p><p>Note:&nbsp;</p><p>- The end_date must be later than the start_date and earlier than the latest data date</p><p>- The latest data date can be obtained by using "AmsMarker" in the v2.ams.get_performance_data_update_time API.</p>
	PeriodType string `json:"period_type" url:"period_type"` // [Required] <p>Period Type. Applicable values:<br />Day<br />Week<br />Month<br />Last7d<br />Last30d<br /><br /></p><p>Note: The start date and end date must align with the Period Type.</p>
	StartDate string `json:"start_date" url:"start_date"` // [Required] <p>The start_date must be:</p><p>- Any day in the past three calendar months for "Day" period type</p><p>- Sunday for "Week" period type</p><p>- The 1st day of a Month for "Month" period type</p><p>- The date that is 6 days prior to the latest data date for "Last7d" period type</p><p>- The date that is 29 days prior to the latest data date for "Last30d" period type</p><p><br /></p><p>Note: The latest data date can be obtained by using "AmsMarker" in the v2.ams.get_performance_data_update_time API.</p>
}
type GetCampaignKeyMetricsPerformanceResponse struct {
	BaseResponse // Common response fields
	Response GetCampaignKeyMetricsPerformanceResponseData `json:"response"` // 
}
type GetCampaignKeyMetricsPerformanceResponseData struct {
	OpenCampaignKeyMetircs *OpenCampaignKeyMetircs `json:"open_campaign_key_metircs"` // [Required] <p>Performance data of Open Campaign.</p>
	TargetedCampaignKeyMetircs *OpenCampaignKeyMetircs `json:"targeted_campaign_key_metircs"` // [Required] <p>Performance data of Target Campaign.</p>
	FetchedDateRange string `json:"fetched_date_range"` // [Required] <p>Effective query date range. Invalid input ranges will be automatically shifted.</p>
}
type GetContentPerformanceRequest struct {
	AffiliateId *int64 `json:"affiliate_id,omitempty" url:"affiliate_id,omitempty"` // [Optional] <p>Search for the contents published by affiliates with the affiliate id entered.</p>
	Channel string `json:"channel" url:"channel"` // [Required] <p>Channel. Applicable values:&nbsp;<br />- ShopeeVideo<br />- LiveStreaming</p>
	EndDate string `json:"end_date" url:"end_date"` // [Required] <p>The end_date must be:&nbsp;</p><p>- Equal to start_date for "Day" period type</p><p>- Saturday for "Week" period type</p><p>- The last day of a Month for "Month" period type. If the selected month is the current month, the end_date should be the latest data date</p><p>- The latest data date for "Last7d" period type</p><p>- The latest data date for "Last30d" period type</p><p><br /></p><p>Note:&nbsp;</p><p>- The end_date must be later than the start_date and earlier than the latest data date</p><p>- The latest data date can be obtained by using "AmsMarker" in the v2.ams.get_performance_data_update_time API.</p>
	ItemId *int64 `json:"item_id,omitempty" url:"item_id,omitempty"` // [Optional] <p>Search for the contents with the searched product included (precise search).</p>
	OrderType string `json:"order_type" url:"order_type"` // [Required] <p>Order Type. Applicable values:&nbsp;</p><p>PlacedOrder<br />ConfirmedOrder:&nbsp;<br /><br />Note:&nbsp;</p><p>- Placed orders are orders (COD and non-COD) that buyers have successfully placed, including paid and unpaid orders.</p><p>- Confirmed orders are either non-COD orders that have been paid for or COD orders that have been confirmed for shipping (usually 30 mins after placing the order).</p>
	PageNo int64 `json:"page_no" url:"page_no"` // [Required] <p>Specifies the page number of data to return in the current call. Starting from 1. if data is more than one page, the page_no can be some entry to start next call.</p>
	PageSize int64 `json:"page_size" url:"page_size"` // [Required] <p>Each result set is returned as a page of entries. Use the "page_size" filters to control the maximum number of entries to retrieve per page (i.e., per call). This integer value is used to specify the maximum number of entries to return in a single "page" of data. The limit of page_size if between 1 and 20.</p>
	PeriodType string `json:"period_type" url:"period_type"` // [Required] <p>Period Type. Applicable values:<br />Day<br />Week<br />Month<br />Last7d<br />Last30d<br /><br /></p><p>Note: The start date and end date must align with the Period Type.</p>
	StartDate string `json:"start_date" url:"start_date"` // [Required] <p>The start_date must be:</p><p>- Any day in the past three calendar months for "Day" period type</p><p>- Sunday for "Week" period type</p><p>- The 1st day of a Month for "Month" period type</p><p>- The date that is 6 days prior to the latest data date for "Last7d" period type</p><p>- The date that is 29 days prior to the latest data date for "Last30d" period type</p><p><br /></p><p>Note: The latest data date can be obtained by using "AmsMarker" in the v2.ams.get_performance_data_update_time API.</p>
}
type GetContentPerformanceResponse struct {
	BaseResponse // Common response fields
	Response GetContentPerformanceResponseData `json:"response"` // 
}
type GetContentPerformanceResponseData struct {
	List []ResponseDataList `json:"list"` // [Required] 
	TotalCount int64 `json:"total_count"` // [Required] <p>This is to indicate the whole number of items.</p>
	HasMore bool `json:"has_more"` // [Required] <p>This is to indicate whether the list is more than one page. If this value is true, you may want to continue to check next page to retrieve the rest of datas.</p>
	FetchedDateRange string `json:"fetched_date_range"` // [Required] <p>Effective query date range. Invalid input ranges will be automatically shifted.</p>
}
type GetConversionReportRequest struct {
	AffiliateId *int64 `json:"affiliate_id,omitempty" url:"affiliate_id,omitempty"` // [Optional] <p>ID of the affiliate who promoted the item.</p>
	AmsDeductionTimeEnd *int64 `json:"ams_deduction_time_end,omitempty" url:"ams_deduction_time_end,omitempty"` // [Optional] <p>End time (inclusive) of fee deduction, in timestamp format.</p>
	AmsDeductionTimeStart *int64 `json:"ams_deduction_time_start,omitempty" url:"ams_deduction_time_start,omitempty"` // [Optional] <p>Start time (inclusive) of fee deduction, in timestamp format.</p>
	AttrCampaignId *int64 `json:"attr_campaign_id,omitempty" url:"attr_campaign_id,omitempty"` // [Optional] <p>ID referencing the campaign rule applied.</p>
	BuyerStatus *string `json:"buyer_status,omitempty" url:"buyer_status,omitempty"` // [Optional] <p>Buyer Status. Applicable values:</p><p>New</p><p>Existing</p>
	CampaignPartner *string `json:"campaign_partner,omitempty" url:"campaign_partner,omitempty"` // [Optional] <p>Name/ID of campaign partner.</p>
	ConversionCompletedTimeEnd *int64 `json:"conversion_completed_time_end,omitempty" url:"conversion_completed_time_end,omitempty"` // [Optional] <p>End time (inclusive) of final completion, in timestamp format.</p>
	ConversionCompletedTimeStart *int64 `json:"conversion_completed_time_start,omitempty" url:"conversion_completed_time_start,omitempty"` // [Optional] <p>Start time (inclusive) of final completion, in timestamp format.</p>
	DeductionMethod *string `json:"deduction_method,omitempty" url:"deduction_method,omitempty"` // [Optional] <p>Deduction Method. Applicable values:</p><p>OrderEscrow<br />SellerWallet<br />AutoAdjustment<br />SVSPaymentLink<br />OfflineSettlement<br />AMSCredit</p>
	DeductionStatus *string `json:"deduction_status,omitempty" url:"deduction_status,omitempty"` // [Optional] <p>Deduction Status. Applicable values:</p><p>PendingDeduction</p><p>Deducted</p>
	ItemId *int64 `json:"item_id,omitempty" url:"item_id,omitempty"` // [Optional] <p>ID of the product purchased.</p>
	ItemName *string `json:"item_name,omitempty" url:"item_name,omitempty"` // [Optional] <p>Product's name.</p>
	L1CategoryId *int64 `json:"l1_category_id,omitempty" url:"l1_category_id,omitempty"` // [Optional] <p>Hierarchical product category classification. (L1 Category)</p>
	L2CategoryId *int64 `json:"l2_category_id,omitempty" url:"l2_category_id,omitempty"` // [Optional] <p>Hierarchical product category classification. (L2 category)</p>
	L3CategoryId *int64 `json:"l3_category_id,omitempty" url:"l3_category_id,omitempty"` // [Optional] <p>Hierarchical product category classification. (L3 Category)</p>
	OrderCompletedTimeEnd *int64 `json:"order_completed_time_end,omitempty" url:"order_completed_time_end,omitempty"` // [Optional] <p>End time (inclusive) of order completion, in timestamp format.</p>
	OrderCompletedTimeStart *int64 `json:"order_completed_time_start,omitempty" url:"order_completed_time_start,omitempty"` // [Optional] <p>Start time (inclusive) of order completion, in timestamp format.</p>
	OrderSn *string `json:"order_sn,omitempty" url:"order_sn,omitempty"` // [Optional] <p>Unique identifier of the order.</p>
	OrderStatus *OrderStatus `json:"order_status,omitempty" url:"order_status,omitempty"` // [Optional] <p>Order Status. Applicable values:</p><p>Unpaid</p><p>Pending</p><p>Completed</p><p>Cancelled</p>
	PageNo int64 `json:"page_no" url:"page_no"` // [Required] <p>Specifies the page number of data to return in the current call. Starting from 1. If data is more than one page, the page_no can be some entry to start next call.</p>
	PageSize int64 `json:"page_size" url:"page_size"` // [Required] <p>Number of records returned per page, the maximum limit is 500, and page_no * page_size must be &lt;= 10000.</p>
	PlaceOrderTimeEnd *int64 `json:"place_order_time_end,omitempty" url:"place_order_time_end,omitempty"` // [Optional] <p>End time (inclusive) of order placement, in timestamp format.</p>
	PlaceOrderTimeStart *int64 `json:"place_order_time_start,omitempty" url:"place_order_time_start,omitempty"` // [Optional] <p>Start time (inclusive) of order placement, in timestamp format.</p>
	SellerCampaignType *string `json:"seller_campaign_type,omitempty" url:"seller_campaign_type,omitempty"` // [Optional] <p>Seller Campaign Type. Applicable values:</p><p>TargetCampaign</p><p>OpenCampaign</p><p>MCNCampaign</p>
	VerifiedStatus *string `json:"verified_status,omitempty" url:"verified_status,omitempty"` // [Optional] <p>Verified Status. Applicable values:</p><p>Unverified</p><p>Valid</p><p>Invalid</p>
}
type GetConversionReportResponse struct {
	BaseResponse // Common response fields
	Response GetConversionReportResponseData `json:"response"` // <p>Response payload containing result data.</p>
}
type GetConversionReportResponseData struct {
	List []GetConversionReportResponseDataList `json:"list"` // [Required] <p>Array of order records. Each object contains order and commission details.</p>
	TotalCount int64 `json:"total_count"` // [Required] <p>Total number of entities that match the condition.</p>
	HasMore bool `json:"has_more"` // [Required] <p>This is to indicate whether the list is more than one page. If this value is true, you may want to continue to check next page to retrieve the rest of datas.</p>
}
type GetConversionReportResponseDataList struct {
	OrderSn string `json:"order_sn"` // [Required] <p>Unique identifier of the order.</p>
	OrderStatus OrderStatus `json:"order_status"` // [Required] <p>Current status of the order (e.g., Pending, Completed, Cancelled).</p>
	VerifiedStatus string `json:"verified_status"` // [Required] <p>Verification status of the order (Unverified, Verified).</p>
	PlaceOrderTime string `json:"place_order_time"` // [Required] <p>Time when the order was placed.</p>
	OrderCompletedTime string `json:"order_completed_time"` // [Required] <p>Time when the order was marked as completed.</p>
	ConversionCompletedTime string `json:"conversion_completed_time"` // [Required] <p>Time when the conversion (affiliate action) was completed.</p>
	AffiliateName string `json:"affiliate_name"` // [Required] <p>Display name of the affiliate who promoted the item.</p>
	AffiliateUsername string `json:"affiliate_username"` // [Required] <p>Login username of the affiliate.</p>
	LinkedMcn string `json:"linked_mcn"` // [Required] <p>MCN (Multi-Channel Network) linked with the affiliate, if any.</p>
	CampaignPartner string `json:"campaign_partner"` // [Required] <p>Partner identifier for the campaign.</p>
	OrderType string `json:"order_type"` // [Required] <p>Type of order: Direct Order or Indirect Order.</p>
	OrderBrandCommission string `json:"order_brand_commission"` // [Required] <p>Commission (amount) for the whole order, paid by the seller.</p>
	Channel string `json:"channel"` // [Required] <p>Traffic channel or platform where the promotion took place.</p>
	AffiliateId int64 `json:"affiliate_id"` // [Required] <p>Unique identifier of the affiliate.</p>
	BuyerStatus string `json:"buyer_status"` // [Required] <p>Buyer Status. Applicable values:</p><p>New</p><p>Existing</p>
	Items []Items `json:"items"` // [Required] 
}
type GetManagedAffiliateListRequest struct {
	PageNo int64 `json:"page_no" url:"page_no"` // [Required] <p>The start index of request.</p><p>The max managed affiliates of affiliate is 2000. Zero count will returned if offset &gt; 2000 or offset &gt; real managed count.</p>
	PageSize int64 `json:"page_size" url:"page_size"` // [Required] <p>The number of affiliate returned by this request, Max is 100, default is 20.</p><p>The max managed affiliates of affiliate is 2000.&nbsp;</p>
}
type GetManagedAffiliateListResponse struct {
	BaseResponse // Common response fields
	Response GetManagedAffiliateListResponseData `json:"response"` // <p>Detail informations you are querying.</p>
}
type GetManagedAffiliateListResponseData struct {
	TotalCount int64 `json:"total_count"` // [Required] <p>The total count of affiliates that managed by this seller.</p>
	AffiliateList []ResponseDataAffiliate `json:"affiliate_list"` // [Required] <p>Affiliate list managed by seller.</p><p>Not all return fields will have values.</p>
}
type GetOpenCampaignAddedProductRequest struct {
	Cursor *string `json:"cursor,omitempty" url:"cursor,omitempty"` // [Optional] <p>Specifies the starting entry of data to return in the current call. Default is "" or not passed. If data is more than one page, the cursor can be some entry to start next call.</p>
	PageSize int64 `json:"page_size" url:"page_size"` // [Required] <p>Each result set is returned as a page of entries. Use the "page_size" filters to control the maximum number of entries to retrieve per page (i.e., per call). This integer value is used to specify the maximum number of entries to return in a single "page" of data. The limit of page_size if between 1 and 100.</p>
	SearchContent *string `json:"search_content,omitempty" url:"search_content,omitempty"` // [Optional] <p>Search for item_name or item_id, item_id should be split by comma and at most 50 items.</p>
	SearchType *string `json:"search_type,omitempty" url:"search_type,omitempty"` // [Optional] <p>Search type: ITEM_NAME or ITEM_ID</p>
	SortBy *string `json:"sort_by,omitempty" url:"sort_by,omitempty"` // [Optional] <p>Use this field to specify which field to use to sort the returned item list. Sort by update_time and commission_id in descending order by default.&nbsp;Available values:</p><p>commission_rate: Sort by commission_rate in ascending order</p><p>-commission_rate: Sort by commission_rate in descending order</p>
}
type GetOpenCampaignAddedProductResponse struct {
	BaseResponse // Common response fields
	Response GetOpenCampaignAddedProductResponseData `json:"response"` // 
}
type GetOpenCampaignAddedProductResponseData struct {
	ItemList []ResponseDataItem `json:"item_list"` // [Required] 
	TotalCount int64 `json:"total_count"` // [Required] <p>Total number of items that match the condition</p>
	Cursor string `json:"cursor"` // [Required] <p>Pass the content in the next request as cursor to get the next page data</p>
	HasMore bool `json:"has_more"` // [Required] <p>This is to indicate whether the item list is more than one page. If this value is true, you may want to continue to check next page to retrieve orders.</p>
}
type GetOpenCampaignBatchTaskResultRequest struct {
	TaskId string `json:"task_id" url:"task_id"` // [Required] <p>Task id, used to query task progress</p>
}
type GetOpenCampaignBatchTaskResultResponse struct {
	BaseResponse // Common response fields
	Response GetOpenCampaignBatchTaskResultResponseData `json:"response"` // 
}
type GetOpenCampaignBatchTaskResultResponseData struct {
	Status string `json:"status"` // [Required] <p>Task status. Applicable values:</p><p>Doing</p><p>Done</p><p>Fail</p><p><br /></p><p>Note: Please&nbsp;note that task&nbsp;<b>Done</b>&nbsp;here refers to the completion of scanning all products in the shop, but not the successful execution of all products. Some products may fail, but due to the unpredictable huge volume of data, detailed information will not returned in the&nbsp;fail_reason. After the task is&nbsp;<b>Done</b>, you need to retrieve the list again by GET API and compare it with the before list to confirm the execution details</p>
	ProgressRate int64 `json:"progress_rate"` // [Required] <p>Progress rate, 80 means 80%</p>
	FailReason string `json:"fail_reason"` // [Required] <p>Error message, if it is not empty, it means there is an error</p><p>Will not return the detail error for each products, you can check the products detail by using GET API, or using the batch operate API</p>
}
type GetOpenCampaignNotAddedProductRequest struct {
	Cursor *string `json:"cursor,omitempty" url:"cursor,omitempty"` // [Optional] <p>Specifies the starting entry of data to return in the current call. Default is "" or not passed. If data is more than one page, the cursor can be some entry to start next call.</p>
	PageSize int64 `json:"page_size" url:"page_size"` // [Required] <p>Each result set is returned as a page of entries. Use the "page_size" filters to control the maximum number of entries to retrieve per page (i.e., per call). This integer value is used to specify the maximum number of entries to return in a single "page" of data. The limit of page_size if between 1 and 100.</p>
	SearchContent *string `json:"search_content,omitempty" url:"search_content,omitempty"` // [Optional] <p>Search for item name or item id. item id should be split by comma and at most 50 items. When search_content is passed, search_type is required.</p>
	SearchType *string `json:"search_type,omitempty" url:"search_type,omitempty"` // [Optional] <p>Search&nbsp;type: ITEM_ID or&nbsp;ITEM_NAME</p>
	SortBy *string `json:"sort_by,omitempty" url:"sort_by,omitempty"` // [Optional] <p>Use this field to specify which field to use to sort the returned item list.&nbsp;Available values:<br />-sales: Sort by sales in descending order (default value)<br />sales: Sort by sales in ascending order<br />-stock: Sort by inventory in descending order<br />stock: Sort by inventory in ascending order<br />-price: Sort by price in descending order<br />price: Sort by price in ascending order</p>
}
type GetOpenCampaignNotAddedProductResponse struct {
	BaseResponse // Common response fields
	Response GetOpenCampaignNotAddedProductResponseData `json:"response"` // 
}
type GetOpenCampaignNotAddedProductResponseData struct {
	ItemList []GetOpenCampaignNotAddedProductResponseDataItem `json:"item_list"` // [Required] 
	TotalCount int64 `json:"total_count"` // [Required] <p>Total number of items that match the condition</p>
	Cursor string `json:"cursor"` // [Required] <p>Pass the content in the next request as cursor to get the next page data</p>
	HasMore bool `json:"has_more"` // [Required] <p>This is to indicate whether the item list is more than one page. If this value is true, you may want to continue to check next page to retrieve orders.</p>
}
type GetOpenCampaignNotAddedProductResponseDataItem struct {
	ItemId int64 `json:"item_id"` // [Required] <p>Item ID</p>
	ItemName string `json:"item_name"` // [Required] <p>Item name</p>
	Sales int64 `json:"sales"` // [Required] <p>Item sales</p>
	DisplayPrice string `json:"display_price"` // [Required] <p>Item display price</p>
	Stock int64 `json:"stock"` // [Required] <p>Item stock</p>
	IsInBlacklist bool `json:"is_in_blacklist"` // [Required] <p>If item is in blacklist, it cannot set up&nbsp;open&nbsp;campaign</p>
	WithOpenCampaign bool `json:"with_open_campaign"` // [Required] <p>If item already has open campaign, it cannot set up another open campaign</p><p>The item list may be delayed, so it is used to further filter items that already have open campaigns</p>
}
type GetOpenCampaignPerformanceRequest struct {
	EndDate string `json:"end_date" url:"end_date"` // [Required] <p>The end_date must be:&nbsp;</p><p>- Equal to start_date for "Day" period type</p><p>- Saturday for "Week" period type</p><p>- The last day of a Month for "Month" period type. If the selected month is the current month, the end_date should be the latest data date</p><p>- The latest data date for "Last7d" period type</p><p>- The latest data date for "Last30d" period type</p><p><br /></p><p>Note:&nbsp;</p><p>- The end_date must be later than the start_date and earlier than the latest data date</p><p>- The latest data date can be obtained by using "AmsMarker" in the v2.ams.get_performance_data_update_time API.</p>
	ItemId *int64 `json:"item_id,omitempty" url:"item_id,omitempty"` // [Optional] <p>Item ID for query.</p>
	PageNo int64 `json:"page_no" url:"page_no"` // [Required] <p>Specifies the page number of data to return in the current call. Starting from 1. if data is more than one page, the page_no can be some entry to start next call.</p>
	PageSize int64 `json:"page_size" url:"page_size"` // [Required] <p>Each result set is returned as a page of entries. Use the "page_size" filters to control the maximum number of entries to retrieve per page (i.e., per call). This integer value is used to specify the maximum number of entries to return in a single "page" of data. The limit of page_size if between 1 and 20.</p>
	PeriodType string `json:"period_type" url:"period_type"` // [Required] <p>Period Type. Applicable values:<br />Day<br />Week<br />Month<br />Last7d<br />Last30d<br /><br /></p><p>Note: The start date and end date must align with the Period Type.</p>
	StartDate string `json:"start_date" url:"start_date"` // [Required] <p>The start_date must be:</p><p>- Any day in the past three calendar months for "Day" period type</p><p>- Sunday for "Week" period type</p><p>- The 1st day of a Month for "Month" period type</p><p>- The date that is 6 days prior to the latest data date for "Last7d" period type</p><p>- The date that is 29 days prior to the latest data date for "Last30d" period type</p><p><br /></p><p>Note: The latest data date can be obtained by using "AmsMarker" in the v2.ams.get_performance_data_update_time API.</p>
}
type GetOpenCampaignPerformanceResponse struct {
	BaseResponse // Common response fields
	Response GetOpenCampaignPerformanceResponseData `json:"response"` // 
}
type GetOpenCampaignPerformanceResponseData struct {
	List []GetOpenCampaignPerformanceResponseDataList `json:"list"` // [Required] 
	TotalCount int64 `json:"total_count"` // [Required] <p>This is to indicate the whole number of items.</p>
	HasMore bool `json:"has_more"` // [Required] <p>This is to indicate whether the list is more than one page. If this value is true, you may want to continue to check next page to retrieve the rest of datas.</p>
	FetchedDateRange string `json:"fetched_date_range"` // [Required] <p>Effective query date range. Invalid input ranges will be automatically shifted.</p>
}
type GetOpenCampaignPerformanceResponseDataList struct {
	ItemId int64 `json:"item_id"` // [Required] <p>Unique identifier of the promoted item within open campaign.</p>
	ItemName string `json:"item_name"` // [Required] <p>Name or title of the promoted item within open campaign.</p>
	Affiliates int64 `json:"affiliates"` // [Required] <p>Number of affiliates currently participating in the campaign for this item.</p>
	Sales string `json:"sales"` // [Required] <p>Total sales amount generated from the campaign, in the market's currency.</p>
	ItemSold int64 `json:"item_sold"` // [Required] <p>Total quantity of the item sold through the campaign.</p>
	EstCommission string `json:"est_commission"` // [Required] <p>Estimated commission amount payable to affiliates for this item, based on current campaign data.</p>
}
type GetOptimizationSuggestionProductRequest struct {
	PageNo int64 `json:"page_no" url:"page_no"` // [Required] <p>Specifies the page number of data to return in the current call. Starting from 1. if data is more than one page, the page_no can be some entry to start next call.&nbsp;</p>
	PageSize int64 `json:"page_size" url:"page_size"` // [Required] <p>Each result set is returned as a page of entries. Use the "page_size" filters to control the maximum number of entries to retrieve per page (i.e., per call). This integer value is used to specify the maximum number of entries to return in a single "page" of data.The limit of page_size if between 1 and 100.</p>
	RcmdReasonFilter string `json:"rcmd_reason_filter" url:"rcmd_reason_filter"` // [Required] <p>Recommended types. Applicable values:&nbsp;</p><p>product_opportunities</p><p>optimize_increase_commission_rate</p><p>optimize_extend_promotion_period</p>
}
type GetOptimizationSuggestionProductResponse struct {
	BaseResponse // Common response fields
	Response GetOptimizationSuggestionProductResponseData `json:"response"` // <p>Detail informations you are querying.</p>
}
type GetOptimizationSuggestionProductResponseData struct {
	ItemList []GetOptimizationSuggestionProductResponseDataItem `json:"item_list"` // [Required] 
	Total int64 `json:"total"` // [Required] <p>Total number of items that match the condition</p>
	HasMore bool `json:"has_more"` // [Required] <p>This is to indicate whether the item list is more than one page. If this value is true, you may want to continue to check next page to retrieve orders.</p>
}
type GetOptimizationSuggestionProductResponseDataItem struct {
	ItemId int64 `json:"item_id"` // [Required] <p>Item ID</p>
	ItemName string `json:"item_name"` // [Required] <p>Item Name</p>
	RcmdReason []string `json:"rcmd_reason"` // [Required] <p>Recommend reason. Applicable values:&nbsp;</p><p>severe_shortage</p><p>high_popularity<br />optimize_commission_rate</p><p>extend_time_period</p>
	CampaignId int64 `json:"campaign_id"` // [Required] <p>Campaign ID</p>
	CommissionRate float64 `json:"commission_rate"` // [Required] <p>Commission Rate, 1.11 means 1.11%, support two decimal places</p>
	PeriodStartTime int64 `json:"period_start_time"` // [Required] <p>Period Start Time</p>
	PeriodEndTime int64 `json:"period_end_time"` // [Required] <p>Period End Time, if get 32503651199 (2999-12-31 23:59:59), it means no limit</p>
	AffiliateCount int64 `json:"affiliate_count"` // [Required] <p>The total number of affiliates who have ever shared the product&nbsp;in the last 90 days</p>
	ItemSold int64 `json:"item_sold"` // [Required] <p>If the requested rcmd_reason_filter is product_opportunition, it is represented as item sales. Other scenarios are the total number of item sold for each product of the shop through AMS&nbsp;in the last 90 days</p>
	CampaignStatus CampaignStatus `json:"campaign_status"` // [Required] <p>Campaign Status:</p><p>Upcoming</p><p>Ongoing</p><p>Terminating</p>
	L2CategoryOrderCount int64 `json:"l2_category_order_count"` // [Required] <p>The total number of AMS orders for the product's L2 category in the last 30 days, only available when requested rcmd_reason_filter is product_opportunition</p>
	SuggestMinRate float64 `json:"suggest_min_rate"` // [Required] <p>Minimum suggested commission rate,&nbsp;1.1 means 1.1%, support two decimal places</p>
	SuggestMaxRate float64 `json:"suggest_max_rate"` // [Required] <p>Maximum suggested commission rate,&nbsp;1.2 means 1.2%, support two decimal places</p>
	PrefillRate float64 `json:"prefill_rate"` // [Required] <p>Prefill rate,&nbsp;1.1 means 1.1%, support two decimal places</p>
	PrefillSubsidyRate float64 `json:"prefill_subsidy_rate"` // [Required] <p>Prefill subsidy rate, platform commission rate calculated based on seller commission,&nbsp;1.2 means 1.2%, support two decimal places</p>
	DisplayPrice string `json:"display_price"` // [Required] <p>Display price</p>
	HasSubsidyData bool `json:"has_subsidy_data"` // [Required] <p>Has subsidy rate</p>
}
type GetPerformanceDataUpdateTimeRequest struct {
	MarkerType string `json:"marker_type" url:"marker_type"` // [Required] <p>Marker type. Applicable values:&nbsp;</p><p>- AmsMarker: Used to query the data update date for ams metrics.</p>
}
type GetPerformanceDataUpdateTimeResponse struct {
	BaseResponse // Common response fields
	Response GetPerformanceDataUpdateTimeResponseData `json:"response"` // 
}
type GetPerformanceDataUpdateTimeResponseData struct {
	LastReportDate string `json:"last_report_date"` // [Required] <p>The latest date of AMS dashboard data metrics update.</p>
}
type GetProductPerformanceRequest struct {
	Channel string `json:"channel" url:"channel"` // [Required] <p>Channel. Applicable values:&nbsp;<br />- AllChannel<br />- SocialMedia<br />- ShopeeVideo<br />- LiveStreaming</p>
	EndDate string `json:"end_date" url:"end_date"` // [Required] <p>The end_date must be:&nbsp;</p><p>- Equal to start_date for "Day" period type</p><p>- Saturday for "Week" period type</p><p>- The last day of a Month for "Month" period type. If the selected month is the current month, the end_date should be the latest data date</p><p>- The latest data date for "Last7d" period type</p><p>- The latest data date for "Last30d" period type</p><p><br /></p><p>Note:&nbsp;</p><p>- The end_date must be later than the start_date and earlier than the latest data date</p><p>- The latest data date can be obtained by using "AmsMarker" in the v2.ams.get_performance_data_update_time API.</p>
	ItemId *int64 `json:"item_id,omitempty" url:"item_id,omitempty"` // [Optional] <p>Item ID for query.</p>
	OrderType string `json:"order_type" url:"order_type"` // [Required] <p>Order Type. Applicable values:&nbsp;</p><p>PlacedOrder<br />ConfirmedOrder:&nbsp;<br /><br />Note:&nbsp;</p><p>- Placed orders are orders (COD and non-COD) that buyers have successfully placed, including paid and unpaid orders.</p><p>- Confirmed orders are either non-COD orders that have been paid for or COD orders that have been confirmed for shipping (usually 30 mins after placing the order).</p>
	PageNo int64 `json:"page_no" url:"page_no"` // [Required] <p>Specifies the page number of data to return in the current call. Starting from 1. If data is more than one page, the page_no can be some entry to start next call.</p>
	PageSize int64 `json:"page_size" url:"page_size"` // [Required] <p>Each result set is returned as a page of entries. Use the "page_size" filters to control the maximum number of entries to retrieve per page (i.e., per call). This integer value is used to specify the maximum number of entries to return in a single "page" of data. The limit of page_size if between 1 and 20.</p>
	PeriodType string `json:"period_type" url:"period_type"` // [Required] <p>Period Type. Applicable values:<br />Day<br />Week<br />Month<br />Last7d<br />Last30d<br /><br /></p><p>Note: The start date and end date must align with the Period Type.</p>
	StartDate string `json:"start_date" url:"start_date"` // [Required] <p>The start_date must be:</p><p>- Any day in the past three calendar months for "Day" period type</p><p>- Sunday for "Week" period type</p><p>- The 1st day of a Month for "Month" period type</p><p>- The date that is 6 days prior to the latest data date for "Last7d" period type</p><p>- The date that is 29 days prior to the latest data date for "Last30d" period type</p><p><br /></p><p>Note: The latest data date can be obtained by using "AmsMarker" in the v2.ams.get_performance_data_update_time API.</p>
}
type GetProductPerformanceResponse struct {
	BaseResponse // Common response fields
	Response GetProductPerformanceResponseData `json:"response"` // 
}
type GetProductPerformanceResponseData struct {
	List []GetProductPerformanceResponseDataList `json:"list"` // [Required] 
	TotalCount int64 `json:"total_count"` // [Required] <p>Total number of items that match the condition.</p>
	HasMore bool `json:"has_more"` // [Required] <p>This is to indicate whether the list is more than one page. If this value is true, you may want to continue to check next page to retrieve the rest of datas.</p>
	FetchedDateRange string `json:"fetched_date_range"` // [Required] <p>Effective query date range. Invalid input ranges will be automatically shifted.</p>
}
type GetProductPerformanceResponseDataList struct {
	ItemId int64 `json:"item_id"` // [Required] <p>Item ID.</p>
	ItemName string `json:"item_name"` // [Required] <p>Item Name.</p>
	Sales string `json:"sales"` // [Required] <p>Total value of the product sold through affiliate marketing.</p>
	ItemsSold int64 `json:"items_sold"` // [Required] <p>Total number of the product sold through affiliate marketing.</p>
	Orders int64 `json:"orders"` // [Required] <p>Total number of orders including the product generated through affiliate marketing.</p>
	Clicks int64 `json:"clicks"` // [Required] <p>Total number of clicks on your product links through affiliate marketing during the selected period.</p>
	EstCommission string `json:"est_commission"` // [Required] <p>Estimated payout of the product sold through affiliate marketing.</p>
	Roi string `json:"roi"` // [Required] <p>Return on Investment, equal to GMV divided by Estimated Commission. It can be used to evaluate the efficiency of your investment in affiliate marketing on the product.If it does not exist, the return value is --.</p>
	TotalBuyers int64 `json:"total_buyers"` // [Required] <p>Total number of buyers who have purchased the product through affiliate marketing.</p>
	NewBuyers int64 `json:"new_buyers"` // [Required] <p>Total number of new buyers who have purchased the product through affiliate marketing.</p>
}
type GetRecommendedAffiliateListRequest struct {
	PageSize int64 `json:"page_size" url:"page_size"` // [Required] <p>Each result set is returned as a page of entries. Use the "page_size" filters to control the maximum number of entries to retrieve per page (i.e., per call). This integer value is used to specify the maximum number of entries to return in a single "page" of data.&nbsp;</p><p><br /></p><p>Note: The response size will up to 200.</p>
}
type GetRecommendedAffiliateListResponse struct {
	BaseResponse // Common response fields
	Response GetRecommendedAffiliateListResponseData `json:"response"` // <p>Detail informations you are querying.</p>
}
type GetRecommendedAffiliateListResponseData struct {
	TotalCount int64 `json:"total_count"` // [Required] <p>The total count of affiliates that recommended for shop id.&nbsp;</p>
	AffiliateList []ResponseDataAffiliate `json:"affiliate_list"` // [Required] <p>Recommended Affiliate list. Not all return fields will have values.</p>
}
type GetShopPerformanceRequest struct {
	Channel string `json:"channel" url:"channel"` // [Required] <p>Channel. Applicable values:&nbsp;<br />- AllChannel<br />- SocialMedia<br />- ShopeeVideo<br />- LiveStreaming</p>
	EndDate string `json:"end_date" url:"end_date"` // [Required] <p>The end_date must be:&nbsp;</p><p>- Equal to start_date for "Day" period type</p><p>- Saturday for "Week" period type</p><p>- The last day of a Month for "Month" period type. If the selected month is the current month, the end_date should be the latest data date</p><p>- The latest data date for "Last7d" period type</p><p>- The latest data date for "Last30d" period type</p><p><br /></p><p>Note:&nbsp;</p><p>- The end_date must be later than the start_date and earlier than the latest data date</p><p>- The latest data date can be obtained by using "AmsMarker" in the v2.ams.get_performance_data_update_time API.</p>
	OrderType string `json:"order_type" url:"order_type"` // [Required] <p>Order Type. Applicable values:&nbsp;</p><p>PlacedOrder<br />ConfirmedOrder:&nbsp;<br /><br />Note:&nbsp;</p><p>- Placed orders are orders (COD and non-COD) that buyers have successfully placed, including paid and unpaid orders.</p><p>- Confirmed orders are either non-COD orders that have been paid for or COD orders that have been confirmed for shipping (usually 30 mins after placing the order).</p>
	PeriodType string `json:"period_type" url:"period_type"` // [Required] <p>Period Type. Applicable values: <br />Day<br />Week<br />Month<br />Last7d<br />Last30d<br /><br /></p><p>Note: The start date and end date must align with the Period Type.</p>
	StartDate string `json:"start_date" url:"start_date"` // [Required] <p>The start_date must be:</p><p>- Any day in the past three calendar months for "Day" period type</p><p>- Sunday for "Week" period type</p><p>- The 1st day of a Month for "Month" period type</p><p>- The date that is 6 days prior to the latest data date for "Last7d" period type</p><p>- The date that is 29 days prior to the latest data date for "Last30d" period type</p><p><br /></p><p>Note: The latest data date can be obtained by using "AmsMarker" in the v2.ams.get_performance_data_update_time API.</p>
}
type GetShopSuggestedRateResponse struct {
	BaseResponse // Common response fields
	Response GetShopSuggestedRateResponseData `json:"response"` // <p>Detail informations you are querying.</p>
}
type GetShopSuggestedRateResponseData struct {
	MinRate float64 `json:"min_rate"` // [Required] <p>Minimum suggested commission rate, 1.1 means 1.1%, support two decimal places</p>
	MaxRate float64 `json:"max_rate"` // [Required] <p>Maximum suggested commission rate,&nbsp;1.2 means 1.2%, support two decimal places</p>
}
type GetTargetedCampaignAddableProductListRequest struct {
	Cursor *string `json:"cursor,omitempty" url:"cursor,omitempty"` // [Optional] <p>Specifies the starting entry of data to return in the current call. Default is "" or not passed. If data is more than one page, the cursor can be some entry to start next call.</p>
	PageSize int64 `json:"page_size" url:"page_size"` // [Required] <p>Each result set is returned as a page of entries. Use the "page_size" filters to control the maximum number of entries to retrieve per page (i.e., per call). This integer value is used to specify the maximum number of entries to return in a single "page" of data. The limit of page_size if between 1 and 100.</p>
	SearchContent *string `json:"search_content,omitempty" url:"search_content,omitempty"` // [Optional] <p>Search by item name or item ID, item_id should be split by comma and at most 50 items.</p><p>Please specify search_type for it to be effective, otherwise search_content will be ignored.</p>
	SearchType *string `json:"search_type,omitempty" url:"search_type,omitempty"` // [Optional] <p>Search type: ITEM_NAME or ITEM_ID, used with search_content.</p>
	SortBy *string `json:"sort_by,omitempty" url:"sort_by,omitempty"` // [Optional] <p>Use this field to specify which field to use to sort the returned item list. Available values:<br />-sales: Sort by sales in descending order (default value)<br />sales: Sort by sales in ascending order<br />-stock: Sort by inventory in descending order<br />stock: Sort by inventory in ascending order<br />-price: Sort by price in descending order<br />price: Sort by price in ascending order</p>
}
type GetTargetedCampaignAddableProductListResponse struct {
	BaseResponse // Common response fields
	Response GetTargetedCampaignAddableProductListResponseData `json:"response"` // <p>Detail informations you are querying.</p>
}
type GetTargetedCampaignAddableProductListResponseData struct {
	ItemList []GetTargetedCampaignAddableProductListResponseDataItem `json:"item_list"` // [Required] <p>Item list.</p>
	TotalCount int64 `json:"total_count"` // [Required] <p>The total count of items that meet the query criteria.</p>
	Cursor string `json:"cursor"` // [Required] <p>Pass the content in the next request as cursor to get the next page data.</p>
}
type GetTargetedCampaignAddableProductListResponseDataItem struct {
	ItemId int64 `json:"item_id"` // [Required] <p>The unique key for item.</p>
	ItemName string `json:"item_name"` // [Required] <p>The name of the current item.</p>
	Sales int64 `json:"sales"` // [Required] <p>The sold of the current item.</p>
	DisplayPrice string `json:"display_price"` // [Required] <p>The display_price of the current item.</p>
	Stock int64 `json:"stock"` // [Required] <p>The stock of the current item.</p>
	IsInBlacklist bool `json:"is_in_blacklist"` // [Required] <p>Is the current item in the blacklist.</p>
}
type GetTargetedCampaignListRequest struct {
	CampaignIdList *string `json:"campaign_id_list,omitempty" url:"campaign_id_list,omitempty"` // [Optional] <p>The list of campaign_id for query, different campaign id should be split by comma and at most 50 campaigns.</p>
	CampaignName *string `json:"campaign_name,omitempty" url:"campaign_name,omitempty"` // [Optional] <p>Campaign name for query.</p>
	CampaignStatus *CampaignStatus `json:"campaign_status,omitempty" url:"campaign_status,omitempty"` // [Optional] <p>Campaign status for query. Applicable values:&nbsp;</p><p>Upcoming<br />Ongoing<br />Ended<br />Cancelled<br />Draft<br />Terminating<br />Terminated<br />Paused</p>
	ItemId *int64 `json:"item_id,omitempty" url:"item_id,omitempty"` // [Optional] <p>Item id for query.</p>
	ItemName *string `json:"item_name,omitempty" url:"item_name,omitempty"` // [Optional] <p>Item name for query.</p>
	PageNo int64 `json:"page_no" url:"page_no"` // [Required] <p>Specifies the page number of data to return in the current call. Starting from 1. if data is more than one page, the page_no can be some entry to start next call.</p>
	PageSize int64 `json:"page_size" url:"page_size"` // [Required] <p>Each result set is returned as a page of entries. Use the "page_size" filters to control the maximum number of entries to retrieve per page (i.e., per call), and the "page_no" to start next call. This integer value is used to specify the maximum number of entries to return in a single "page" of data.&nbsp;The limit of page_size if between 1 and 100.</p>
	PeriodEndTime *int64 `json:"period_end_time,omitempty" url:"period_end_time,omitempty"` // [Optional] <p>Campaign period end time for query.</p>
	PeriodStartTime *int64 `json:"period_start_time,omitempty" url:"period_start_time,omitempty"` // [Optional] <p>Campaign period start time for query.</p>
}
type GetTargetedCampaignListResponse struct {
	BaseResponse // Common response fields
	Response GetTargetedCampaignListResponseData `json:"response"` // <p>Detail informations you are querying.</p>
}
type GetTargetedCampaignListResponseData struct {
	TotalCount int64 `json:"total_count"` // [Required] <p>The total count of targeted campaigns that meet the query criteria.</p>
	CampaignList []GetTargetedCampaignListResponseDataCampaign `json:"campaign_list"` // [Required] <p>Targeted campaign list.</p>
}
type GetTargetedCampaignListResponseDataCampaign struct {
	CampaignId int64 `json:"campaign_id"` // [Required] <p>The unique key for the current campaign.</p>
	CampaignName string `json:"campaign_name"` // [Required] <p>The name of the current campaign.</p>
	CampaignStatus CampaignStatus `json:"campaign_status"` // [Required] <p>Campaign Status:&nbsp;</p><p>Upcoming<br />Ongoing<br />Ended<br />Cancelled<br />Draft<br />Terminating<br />Terminated<br />Paused</p>
	CampaignSource string `json:"campaign_source"` // [Required] <p>Source of campaign setup. Applicable values:&nbsp;<br />- ShopeeManaged (Note: You cannot view the details or edit this campaign. If you try to do so, an 'invalid campaign_id' error will occur.)<br />- Seller<br />- Unknown</p>
	PeriodStartTime int64 `json:"period_start_time"` // [Required] <p>The start time of the current campaign.</p>
	PeriodEndTime int64 `json:"period_end_time"` // [Required] <p>The end time of the current campaign, if get 32503651199 (2999-12-31 23:59:59), it means no limit</p>
	LastEditor string `json:"last_editor"` // [Required] <p>The last editor of the current campaign.</p>
	LastEditTime int64 `json:"last_edit_time"` // [Required] <p>The last edit time of the current campaign.</p>
	AffiliateCount int64 `json:"affiliate_count"` // [Required] <p>The total count of affiliates associated with the current campaign.</p>
	ItemCount int64 `json:"item_count"` // [Required] <p>The total count of items associated with the current campaign.</p>
	MinRate float64 `json:"min_rate"` // [Required] <p>The min commission rate of the current campaign.</p>
	MaxRate float64 `json:"max_rate"` // [Required] <p>The max commission rate of the current campaign.</p>
}
type GetTargetedCampaignPerformanceRequest struct {
	CampaignId *int64 `json:"campaign_id,omitempty" url:"campaign_id,omitempty"` // [Optional] <p>Campaign ID for query.</p>
	EndDate string `json:"end_date" url:"end_date"` // [Required] <p>The end_date must be:&nbsp;</p><p>- Equal to start_date for "Day" period type</p><p>- Saturday for "Week" period type</p><p>- The last day of a Month for "Month" period type. If the selected month is the current month, the end_date should be the latest data date</p><p>- The latest data date for "Last7d" period type</p><p>- The latest data date for "Last30d" period type</p><p><br /></p><p>Note:&nbsp;</p><p>- The end_date must be later than the start_date and earlier than the latest data date</p><p>- The latest data date can be obtained by using "AmsMarker" in the v2.ams.get_performance_data_update_time API.</p>
	PageNo int64 `json:"page_no" url:"page_no"` // [Required] <p>Specifies the page number of data to return in the current call. Starting from 1. if data is more than one page, the page_no can be some entry to start next call.</p>
	PageSize int64 `json:"page_size" url:"page_size"` // [Required] <p>Each result set is returned as a page of entries. Use the "page_size" filters to control the maximum number of entries to retrieve per page (i.e., per call). This integer value is used to specify the maximum number of entries to return in a single "page" of data. The limit of page_size if between 1 and 20.</p>
	PeriodType string `json:"period_type" url:"period_type"` // [Required] <p>Period Type. Applicable values:<br />Day<br />Week<br />Month<br />Last7d<br />Last30d<br /><br /></p><p>Note: The start date and end date must align with the Period Type.</p>
	StartDate string `json:"start_date" url:"start_date"` // [Required] <p>The start_date must be:</p><p>- Any day in the past three calendar months for "Day" period type</p><p>- Sunday for "Week" period type</p><p>- The 1st day of a Month for "Month" period type</p><p>- The date that is 6 days prior to the latest data date for "Last7d" period type</p><p>- The date that is 29 days prior to the latest data date for "Last30d" period type</p><p><br /></p><p>Note: The latest data date can be obtained by using "AmsMarker" in the v2.ams.get_performance_data_update_time API.</p>
}
type GetTargetedCampaignPerformanceResponse struct {
	BaseResponse // Common response fields
	Response GetTargetedCampaignPerformanceResponseData `json:"response"` // 
}
type GetTargetedCampaignPerformanceResponseData struct {
	List []GetTargetedCampaignPerformanceResponseDataList `json:"list"` // [Required] 
	TotalCount int64 `json:"total_count"` // [Required] <p>This is to indicate the whole number of target campaigns.</p>
	HasMore bool `json:"has_more"` // [Required] <p>This is to indicate whether the list is more than one page. If this value is true, you may want to continue to check next page to retrieve the rest of datas.</p>
	FetchedDateRange string `json:"fetched_date_range"` // [Required] <p>Effective query date range. Invalid input ranges will be automatically shifted.</p>
}
type GetTargetedCampaignPerformanceResponseDataList struct {
	CampaignId int64 `json:"campaign_id"` // [Required] <p>Unique identifier of the targeted campaign.</p>
	Affiliates int64 `json:"affiliates"` // [Required] <p>The number of affiliates ever brought sales for the targeted campaign.</p>
	Sales string `json:"sales"` // [Required] <p>Total sales amount generated from this targeted campaign, in the market's default currency.</p>
	ItemSold int64 `json:"item_sold"` // [Required] <p>Total quantity of the item sold through the targeted campaign.</p>
	EstCommission string `json:"est_commission"` // [Required] <p>The estimated commission amount payable to affiliates from this targeted campaign.</p>
	CampaignName string `json:"campaign_name"` // [Required] <p>Campaign name.</p>
}
type GetTargetedCampaignSettingsRequest struct {
	CampaignId int64 `json:"campaign_id" url:"campaign_id"` // [Required] <p>Campaign id for query.</p><p><br /></p><p>Note: For campaigns with campaign_source = ShopeeManaged, cannot be queried for details through this API.</p>
}
type GetTargetedCampaignSettingsResponse struct {
	BaseResponse // Common response fields
	Response GetTargetedCampaignSettingsResponseData `json:"response"` // <p>Detail informations you are querying.</p>
}
type GetTargetedCampaignSettingsResponseData struct {
	CampaignName string `json:"campaign_name"` // [Required] <p>The name of the current campaign.</p>
	CommissionStatus string `json:"commission_status"` // [Required] <p>Campaign Status:&nbsp;</p><p>Upcoming<br />Ongoing<br />Ended<br />Cancelled<br />Draft<br />Terminating<br />Terminated<br />Paused</p>
	PeriodStartTime int64 `json:"period_start_time"` // [Required] <p>The start time of the current campaign.</p>
	PeriodEndTime int64 `json:"period_end_time"` // [Required] <p>The end time of the current campaign, if get 32503651199 (2999-12-31 23:59:59), it means no limit.</p>
	IsSetBudget bool `json:"is_set_budget"` // [Required] <p>Has the current campaign set a budget.</p><p><br /></p><p>Note: TH not supported</p>
	Budget float64 `json:"budget"` // [Required] <p>The budget of the current campaign.</p><p><br /></p><p>Note: TH not supported</p>
	BudgetCost float64 `json:"budget_cost"` // [Required] <p>The budget already spent on the current campaign.</p><p><br /></p><p>Note: TH not supported</p>
	SellerMessage string `json:"seller_message"` // [Required] <p>The message displayed to affiliates.</p>
	PendingTerminatedTime int64 `json:"pending_terminated_time"` // [Required] <p>Pending Terminated Time.</p>
	AffiliateList []Affiliate `json:"affiliate_list"` // [Required] <p>The&nbsp;list of affiliates associated with the current campaign.</p>
	ItemList []GetTargetedCampaignSettingsResponseDataItem `json:"item_list"` // [Required] <p>The&nbsp;list of items associated with the current campaign.</p>
}
type GetTargetedCampaignSettingsResponseDataItem struct {
	ItemId int64 `json:"item_id"` // [Required] <p>Item ID.</p>
	ItemName string `json:"item_name"` // [Required] <p>Item Name.</p>
	Rate float64 `json:"rate"` // [Required] <p>Commission rate of current item, 1.1 means 1.1%, support two decimal places.</p>
	MaxCommissionRateCurrentDay float64 `json:"max_commission_rate_current_day"` // [Required] <p>Max Commission Rate Current Day,&nbsp;1.1 means 1.1%, support two decimal places.</p>
	CommissionProtectionList []CommissionProtection `json:"commission_protection_list"` // [Required] <p>Commission Protection List.</p>
}
type GetValidationListResponse struct {
	BaseResponse // Common response fields
	Response GetValidationListResponseData `json:"response"` // 
}
type GetValidationListResponseData struct {
	ValidationList []Validation `json:"validation_list"` // [Required] 
}
type GetValidationReportRequest struct {
	AttrCampaignId *int64 `json:"attr_campaign_id,omitempty" url:"attr_campaign_id,omitempty"` // [Optional] <p>ID referencing the campaign rule applied. (Ties to the campaign seller created).</p>
	CampaignSource string `json:"campaign_source" url:"campaign_source"` // [Required] <p>Source of campaign setup. Applicable values:</p><p>ShopeeManaged</p><p>Seller</p>
	ItemId *int64 `json:"item_id,omitempty" url:"item_id,omitempty"` // [Optional] <p>Unique identifier of the product.</p>
	ItemName *string `json:"item_name,omitempty" url:"item_name,omitempty"` // [Optional] <p>The product's name.</p>
	L1CategoryId *int64 `json:"l1_category_id,omitempty" url:"l1_category_id,omitempty"` // [Optional] <p>Hierarchical product category classification. (L1 Category)</p>
	L2CategoryId *int64 `json:"l2_category_id,omitempty" url:"l2_category_id,omitempty"` // [Optional] <p>Hierarchical product category classification. (L2 Category)</p>
	L3CategoryId *int64 `json:"l3_category_id,omitempty" url:"l3_category_id,omitempty"` // [Optional] <p>Hierarchical product category classification. (L3 Category)</p>
	OrderSn *string `json:"order_sn,omitempty" url:"order_sn,omitempty"` // [Optional] <p>Unique identifier of the order.</p>
	PageNo int64 `json:"page_no" url:"page_no"` // [Required] <p>Specifies the page number of data to return in the current call. Starting from 1. If data is more than one page, the page_no can be some entry to start next call.</p>
	PageSize int64 `json:"page_size" url:"page_size"` // [Required] <p>Number of records returned per page, the maximum limit is 500, and page_no * page_size must be &lt;= 10000.</p>
	PlaceOrderTimeEnd int64 `json:"place_order_time_end" url:"place_order_time_end"` // [Required] <p>End time (inclusive) of order placement, in timestamp format.</p>
	PlaceOrderTimeStart int64 `json:"place_order_time_start" url:"place_order_time_start"` // [Required] <p>Start time (inclusive) of order placement, in timestamp format.</p>
	ValidationId string `json:"validation_id" url:"validation_id"` // [Required] <p>Unique identifier of the billing entry.</p>
	ValidationMonth int64 `json:"validation_month" url:"validation_month"` // [Required] <p>Billing month in the format YYYYMM (e.g., 202405).</p>
	VerifiedStatus *string `json:"verified_status,omitempty" url:"verified_status,omitempty"` // [Optional] <p>Verified Status. Applicable values:</p><p>Valid</p><p>Invalid</p>
}
type GetValidationReportResponse struct {
	BaseResponse // Common response fields
	Response GetValidationReportResponseData `json:"response"` // <p>Response payload containing result data.</p>
}
type GetValidationReportResponseData struct {
	List []GetConversionReportResponseDataList `json:"list"` // [Required] <p>Array of order records. Each object contains order and commission details.</p>
	TotalCount int64 `json:"total_count"` // [Required] <p>Total number of entities that match the condition.</p>
	HasMore bool `json:"has_more"` // [Required] <p>This is to indicate whether the list is more than one page. If this value is true, you may want to continue to check next page to retrieve the rest of datas.</p>
}
type Item struct {
	ItemId int64 `json:"item_id"` // [Required] <p>Item ID.</p>
	Rate *float64 `json:"rate,omitempty"` // [Optional] <p>Commission rate of current item, 1.1 means 1.1%, support two decimal places.<br /></p>
}
type Items struct {
	ItemId int64 `json:"item_id"` // [Required] <p>Unique identifier of the item in the order.</p>
	ItemName string `json:"item_name"` // [Required] <p>Name of the item in the order.</p>
	ModelId int64 `json:"model_id"` // [Required] <p>SKU/model identifier for the item.</p>
	L1CategoryId int64 `json:"l1_category_id"` // [Required] <p>Level-1 global category id classification of the item.</p>
	L2CategoryId int64 `json:"l2_category_id"` // [Required] <p>Level-2 global category id classification of the item.</p>
	L3CategoryId int64 `json:"l3_category_id"` // [Required] <p>Level-3 global category id classification of the item.</p>
	PromotionId string `json:"promotion_id"` // [Required] <p>Identifier of the promotion campaign linked to the order.</p>
	Price int64 `json:"price"` // [Required] <p>Item price in cents (or smallest currency unit).</p>
	Qty int64 `json:"qty"` // [Required] <p>Quantity of the item purchased.</p>
	SellerCampaignType string `json:"seller_campaign_type"` // [Required] <p>Type of seller campaign:<br />1. Seller Open Campaign – Open to all affiliates.<br />2. Seller Target Campaign – Restricted to designated affiliates.</p>
	AttrCampaignId int64 `json:"attr_campaign_id"` // [Required] <p>ID referencing the campaign rule applied.</p>
	PurchaseValue int64 `json:"purchase_value"` // [Required] <p>Total purchase value of the order in cents (or smallest currency unit).</p>
	RefundAmount string `json:"refund_amount"` // [Required] <p>Amount refunded for the item.</p>
	ItemBrandCommission string `json:"item_brand_commission"` // [Required] <p>Commission (amount) for the item, paid by the seller.</p>
	ItemBrandCommissionRateToAffiliate string `json:"item_brand_commission_rate_to_affiliate"` // [Required] <p>Commission rate allocated to the affiliate for the item.</p>
	ItemBrandCommissionToAffiliate string `json:"item_brand_commission_to_affiliate"` // [Required] <p>Commission (amount) allocated to the affiliate for the item.</p>
	ItemBrandCommissionRateToMcn string `json:"item_brand_commission_rate_to_mcn"` // [Required] <p>Commission rate allocated to the MCN for the item.</p>
	ItemBrandCommissionToMcn string `json:"item_brand_commission_to_mcn"` // [Required] <p>Commission (amount) allocated to the MCN for the item.</p>
}
type List struct {
	AffiliateId int64 `json:"affiliate_id"` // [Required] <p>Unique identifier assigned to the affiliate. Used as a reference key in the system.</p>
	AffiliateName string `json:"affiliate_name"` // [Required] <p>Display name of the affiliate, typically the Shopee display name.</p>
	AffiliateUsername string `json:"affiliate_username"` // [Required] <p>Login or Shopee account username associated with the affiliate.</p><p>   </p><p><br /></p>
	Sales string `json:"sales"` // [Required] <p>Total value of the product sold through the affiliate's promotion.</p>
	ItemsSold int64 `json:"items_sold"` // [Required] <p>Total number of the product sold through the affiliate's promotion.<br /></p>
	Orders int64 `json:"orders"` // [Required] <p>Total number of orders generated through the affiliate's promotion.</p>
	Clicks int64 `json:"clicks"` // [Required] <p>Total number of clicks on your product links through affiliate marketing during the selected period.</p>
	EstCommission string `json:"est_commission"` // [Required] <p>Estimated payout through the affiliate's promotion.</p>
	Roi string `json:"roi"` // [Required] <p>Return on Investment, equal to GMV divided by Estimated Commission. It can be used to evaluate the efficiency of the affiliate's promotion. If it does not exist, the return value is --.</p>
	TotalBuyers int64 `json:"total_buyers"` // [Required] <p>Total number of buyers who have purchased the product through the affiliate's promotion.</p>
	NewBuyers int64 `json:"new_buyers"` // [Required] <p>Total number of new buyers who have purchased the product through the affiliate's promotion.</p>
}
type OfflineBills struct {
	OrderPlaceMonth int64 `json:"order_place_month"` // [Required] <p>Order placement month in the format YYYYMM.</p>
	TotalAmount float64 `json:"total_amount"` // [Required] <p>Total commission amount = commission_amount_after_tax + ams_credit_deducted_amount.</p>
	CommissionAmount float64 `json:"commission_amount"` // [Required] <p>Offline commission amount before tax.</p>
	CommissionAmountAfterTax float64 `json:"commission_amount_after_tax"` // [Required] <p>Offline commission amount including tax.</p>
	AmsCreditDeductedAmount float64 `json:"ams_credit_deducted_amount"` // [Required] <p>Commission amount already paid using AMS Credits.</p>
}
type OnlineBill struct {
	TotalAmount float64 `json:"total_amount"` // [Required] <p>Total commission amount for the billing month.</p>
	BillStatus int64 `json:"bill_status"` // [Required] <p>Billing Status. Applicable values:</p><p>1 = Pending<br />2 = Completed</p><p>3 = In process</p><p>4 = To pay via payment link</p><p>5 = Manual completed</p><p>6 = To Be Settled Offline</p>
	DeductedAmount float64 `json:"deducted_amount"` // [Required] <p>Commission amount already deducted.</p>
	AmsCreditDeductedAmount float64 `json:"ams_credit_deducted_amount"` // [Required] <p>Commission amount paid using AMS Credits</p>
	PendingAmount int64 `json:"pending_amount"` // [Required] <p>Commission amount pending deduction.</p>
}
type OpenCampaignKeyMetircs struct {
	Affiliates int64 `json:"affiliates"` // [Required] <p>Total number of affiliates who drove orders from Targeted Campaigns.</p>
	ItemsSold int64 `json:"items_sold"` // [Required] <p>Total number of items sold from Targeted Campaigns.</p>
	Sales string `json:"sales"` // [Required] <p>Total value of orders from Targeted Campaigns.</p>
	EstCommission string `json:"est_commission"` // [Required] <p>Total estimated commission for orders placed from Targeted Campaigns.</p>
}
type PopularSocialMedia struct {
	Platform string `json:"platform"` // [Required] <p>The platform of this social media account.&nbsp;</p>
	FollowerCount int64 `json:"follower_count"` // [Required] <p>The follower count of this account.</p>
}
type QueryAffiliateListRequest struct {
	AffiliateIdList *string `json:"affiliate_id_list,omitempty" url:"affiliate_id_list,omitempty"` // [Optional] <p>Query affiliate information by affiliate id list.</p><p>Max count of affiliate id is 200. Will return first 200 affiliates' information if length &gt; 200.</p>
	Name *string `json:"name,omitempty" url:"name,omitempty"` // [Optional] <p>Query affiliate information by name use fuzzy matching.</p><p>Will return first 200 affiliates' information is match number &gt; 200.</p>
	QueryType int64 `json:"query_type" url:"query_type"` // [Required] <p>Query type:&nbsp;</p><p>1: query affiliate information by id list</p><p>2: query affiliate id by name(fuzzy matching), only return affiliate id and affiliate name</p>
}
type QueryAffiliateListResponse struct {
	BaseResponse // Common response fields
	Response QueryAffiliateListResponseData `json:"response"` // 
}
type QueryAffiliateListResponseData struct {
	TotalCount int64 `json:"total_count"` // [Required] <p>The total count of affiliates by this query.</p><p>Max is 200.</p>
	AffiliateList *QueryAffiliateListResponseDataAffiliate `json:"affiliate_list"` // [Required] <p>Affiliate list by this query.</p><p>Not all return fields will have values.</p>
}
type QueryAffiliateListResponseDataAffiliate struct {
	AffiliateId int64 `json:"affiliate_id"` // [Required] <p>The unique key for the current affiliate.</p>
	AffiliateName string `json:"affiliate_name"` // [Required] <p>The name of the affiliate.</p>
	UserName string `json:"user_name"` // [Required] <p>The shopee user name or affiliate name for this affiliate.</p>
	PortraitUrl string `json:"portrait_url"` // [Required] <p>The url of affiliate's portrait</p>
	PopularSocialMedia []PopularSocialMedia `json:"popular_social_media"` // [Required] <p>The popular social media of this affiliate.</p>
	SocialMedias []SocialMedias `json:"social_medias"` // [Required] <p>Social media account list of this affiliate.</p>
	TotalClick int64 `json:"total_click"` // [Required] <p>Number of clicks in the last 30 days.</p>
	OrderRange []int64 `json:"order_range"` // [Required] <p>Range number of the orders in the last 30 days.</p>
	GmvRange []int64 `json:"gmv_range"` // [Required] <p>Range number of the gmv in the last 30 days.</p>
	IsOrangeTickKol bool `json:"is_orange_tick_kol"` // [Required] <p>Golden tick means affiliates create high quality contents with good sales conversion in Shopee Live or Shopee Video.</p>
	IsGoodFulfillment bool `json:"is_good_fulfillment"` // [Required] <p>Good sample fulfillment means that affiliates demonstrate better in free sample fulfillment compared to the majority of affiliates in recent180 days</p>
	PromoteCategoryIds []int64 `json:"promote_category_ids"` // [Required] <p>Three promote category ids for this affiliate</p>
	TopPopularContents []TopPopularContents `json:"top_popular_contents"` // [Required] <p>Top popular contents of this affiliate.&nbsp;</p>
	TopSellingProducts []TopSellingProducts `json:"top_selling_products"` // [Required] <p>Top selling items of the affiliate.</p>
}
type Rates struct {
	ItemId int64 `json:"item_id"` // [Required] <p>Item ID</p>
	MinRate float64 `json:"min_rate"` // [Required] <p>Minimum suggested commission rate, 1.1 means 1.1%, support two decimal places</p>
	MaxRate float64 `json:"max_rate"` // [Required] <p>Maximum suggested commission rate,&nbsp;1.2 means 1.2%, support two decimal places</p>
}
type RemoveAllProductsOpenCampaignSettingResponse struct {
	BaseResponse // Common response fields
	Response RemoveAllProductsOpenCampaignSettingResponseData `json:"response"` // 
}
type RemoveAllProductsOpenCampaignSettingResponseData struct {
	TaskType string `json:"task_type"` // [Required] <p>Task type. Applicable values:&nbsp;</p><p>batch_add_open_campaigns</p><p>batch_remove_open_campaigns</p><p>batch_update_open_campaigns</p><p><br /></p><p>For this API, task type will be&nbsp;batch_remove_open_campaigns</p>
	TaskId string `json:"task_id"` // [Required] <p>Task id, used to query task progress when calling v2.ams.get_open_campaign_batch_task_result API</p>
}
type ResponseDataAffiliate struct {
	AffiliateId int64 `json:"affiliate_id"` // [Required] <p>The unique key for the current affiliate.</p>
	AffiliateName string `json:"affiliate_name"` // [Required] <p>The name of the affiliate.</p>
	UserName string `json:"user_name"` // [Required] <p>The shopee user name or affiliate name for this affiliate.</p>
	PortraitUrl string `json:"portrait_url"` // [Required] <p>The portrait url of affiliate.</p>
	PopularSocialMedia *PopularSocialMedia `json:"popular_social_media"` // [Required] <p>The popular social media of this affiliate.</p>
	SocialMedias []SocialMedias `json:"social_medias"` // [Required] <p>Social media account list of this affiliate.</p>
	TotalClick int64 `json:"total_click"` // [Required] <p>Number of clicks in the last 30 days.</p>
	OrderRange []int64 `json:"order_range"` // [Required] <p>Range number of the orders in the last 30 days.</p>
	GmvRange []int64 `json:"gmv_range"` // [Required] <p>Range number of the GMV in the last 30 days.</p>
	IsOrangeTickKol bool `json:"is_orange_tick_kol"` // [Required] <p>Golden tick means affiliates create high quality contents with good sales conversion in Shopee Live or Shopee Video.</p>
	IsGoodFulfillment bool `json:"is_good_fulfillment"` // [Required] <p>Good sample fulfillment means that affiliates demonstrate better in free sample fulfillment compared to the majority of affiliates in recent180 days</p>
	PromoteCategoryIds []int64 `json:"promote_category_ids"` // [Required] <p>Three promote category ids for this affiliate.</p>
	TopPopularContents []TopPopularContents `json:"top_popular_contents"` // [Required] <p>Top popular contents of this affiliate.&nbsp;</p>
	TopSellingProducts []TopSellingProducts `json:"top_selling_products"` // [Required] <p>Top selling items of the affiliate.</p>
}
type ResponseDataFailed struct {
	CampaignId int64 `json:"campaign_id"` // [Required] <p>Campaign ID</p>
	FailError string `json:"fail_error"` // [Required] <p>Fail error</p>
	FailMessage string `json:"fail_message"` // [Required] <p>Fail message</p>
}
type ResponseDataItem struct {
	ItemId int64 `json:"item_id"` // [Required] <p>Item ID</p>
	ItemName string `json:"item_name"` // [Required] <p>Item Name</p>
	CampaignId int64 `json:"campaign_id"` // [Required] <p>Campaign ID</p>
	CampaignStatus CampaignStatus `json:"campaign_status"` // [Required] <p>Campaign Status:&nbsp;</p><p>Upcoming</p><p>Ongoing</p><p>Terminating</p>
	CommissionRate float64 `json:"commission_rate"` // [Required] <p>Commission Rate, 1.1 means 1.1%, support two decimal places</p>
	PeriodStartTime int64 `json:"period_start_time"` // [Required] <p>Period Start Time</p>
	PeriodEndTime int64 `json:"period_end_time"` // [Required] <p>Period End Time, if get 32503651199 (2999-12-31 23:59:59), it means no limit</p>
	PendingTerminatedTime int64 `json:"pending_terminated_time"` // [Required] <p>Pending Terminated Time</p>
	CommissionProtectionList []CommissionProtection `json:"commission_protection_list"` // [Required] <p>Commission Protection List</p>
	MaxCommissionRateCurrentDay float64 `json:"max_commission_rate_current_day"` // [Required] <p>Max Commission Rate Current Day,&nbsp;1.1 means 1.1%, support two decimal places</p>
}
type ResponseDataList struct {
	ContentId string `json:"content_id"` // [Required] <p>Unique identifier of the content where the product is placed.</p>
	ContentTitle string `json:"content_title"` // [Required] <p>Title or name of the content (e.g., video, livestream) associated with the product.</p>
	PostTime int64 `json:"post_time"` // [Required] <p>Livestream:&nbsp; The livestream start time.</p><p>Video: The video post time.</p>
	AffiliateName string `json:"affiliate_name"` // [Required] <p>Display name of the affiliate who posted the content, typically the Shopee name.</p>
	AffiliateUsername string `json:"affiliate_username"` // [Required] <p>Login or Shopee account username associated with the affiliate.</p>
	Products int64 `json:"products"` // [Required] <p>Number of products associated with the content.</p>
	Views int64 `json:"views"` // [Required] <p>The total viewed pv of the content of this shop within the selected time range</p>
	Likes int64 `json:"likes"` // [Required] <p>The total number of likes for the content of this shop within the selected time range</p>
	Comments int64 `json:"comments"` // [Required] <p>The total number of comments for the content of this shop within the selected time range</p>
	Sales string `json:"sales"` // [Required] <p>The total sales of the content associated with the shop orders within the selected time range</p>
	Orders int64 `json:"orders"` // [Required] <p>The total number of orders associated with the shop for the content in the selected time range</p>
	ItemsSold int64 `json:"items_sold"` // [Required] <p>The total number of items sold associated with the shop for the content in the selected time range</p>
	Channel string `json:"channel"` // [Required] <p>Channel. Applicable values:&nbsp;<br />- ShopeeVideo<br />- LiveStreaming</p>
}
type SocialMedias struct {
	Platform string `json:"platform"` // [Required] <p>The platform of this social media account.&nbsp;</p>
	FollowerCount int64 `json:"follower_count"` // [Required] <p>The follower count of this account.</p>
	SocialMediaUserName string `json:"social_media_user_name"` // [Required] <p>Social media name of this account.</p>
}
type TerminateTargetedCampaignRequest struct {
	CampaignId int64 `json:"campaign_id"` // [Required] <p>The unique key for the current campaign.</p>
}
type TerminateTargetedCampaignResponse struct {
	BaseResponse // Common response fields
}
type TopPopularContents struct {
	Platform string `json:"platform"` // [Required] <p>The platform of this affiliate's content.</p>
	CommentCount int64 `json:"comment_count"` // [Required] <p>The comment count of this affiliate's content.</p>
	LikeCount int64 `json:"like_count"` // [Required] <p>The like count of this affiliate's content.</p>
	ViewCount int64 `json:"view_count"` // [Required] <p>The view count of this affiliate's content.</p>
	CoverUrl string `json:"cover_url"` // [Required] <p>The cover link of this affiliate's content.</p>
	MediaUrl string `json:"media_url"` // [Required] <p>The media link of this affiliate's content.</p>
}
type UpdateAutoAddNewProductSettingRequest struct {
	CommissionRate *float64 `json:"commission_rate,omitempty"` // [Optional] <p>Commission rate, 1.1 means 1.1%, support two decimal places</p>
	Open bool `json:"open"` // [Required] <p>Enable or disable auto-add new product, if true is passed, it means enabled, if false is passed, it means disabled</p>
}
type UpdateAutoAddNewProductSettingResponse struct {
	BaseResponse // Common response fields
}
type UpdateBasicInfoOfTargetedCampaignRequest struct {
	Budget *float64 `json:"budget,omitempty"` // [Optional] <p>Budget value set for the current campaign.</p><p><br /></p><p>Note: TH not supported</p>
	CampaignId int64 `json:"campaign_id"` // [Required] <p>The unique key for the current campaign.</p>
	CampaignName *string `json:"campaign_name,omitempty"` // [Optional] <p>The name of the current campaign.</p>
	IsSetBudget *bool `json:"is_set_budget,omitempty"` // [Optional] <p>Budget allocation toggle for the current campaign.</p><p><br /></p><p>Note: TH not supported</p>
	PeriodEndTime *int64 `json:"period_end_time,omitempty"` // [Optional] <p>The period end time of campaign, in seconds.</p><p>Can set&nbsp;32503651199 (2999-12-31 23:59:59) represent of no limit.</p>
	PeriodStartTime *int64 `json:"period_start_time,omitempty"` // [Optional] <p>The start time of the designated campaign,&nbsp;in seconds.</p>
}
type UpdateBasicInfoOfTargetedCampaignResponse struct {
	BaseResponse // Common response fields
}
type Validation struct {
	ValidationId string `json:"validation_id"` // [Required] <p>Unique identifier of the billing entry.</p>
	PaymentMethod int64 `json:"payment_method"` // [Required] <p>Payment method. Applicable values:</p><p>1 = Online</p><p>2 = Offline</p>
	ValidationMonth int64 `json:"validation_month"` // [Required] <p>Billing month in the format YYYYMM (e.g., 202405).</p>
	CampaignSource string `json:"campaign_source"` // [Required] <p>Source of campaign setup. Applicable values:</p><p>ShopeeManaged</p><p>Seller</p>
	OnlineBill *OnlineBill `json:"online_bill"` // [Required] <p>Billing details when payment method is Online.</p>
	OfflineBills []OfflineBills `json:"offline_bills"` // [Required] <p>List of billing details when payment method is Offline, grouped by order placement month.<br /></p>
}
