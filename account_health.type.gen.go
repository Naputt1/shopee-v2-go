package goshopee

type AptOrder struct {
	OrderSn string `json:"order_sn"` // [Required] <p>Order SN.</p>
	OrderCreateTime int64 `json:"order_create_time"` // [Required] <p>Order Paid Time.</p>
	ArrangePickUpTime int64 `json:"arrange_pick_up_time"` // [Required] <p>Seller arrange pick up time.</p>
	ActualPickUpTime int64 `json:"actual_pick_up_time"` // [Required] <p>Courier actual pick up time.</p>
	PreparationDays float64 `json:"preparation_days"` // [Required] <p>Preparation Days.</p>
	ShippingChannel string `json:"shipping_channel"` // [Required] <p>Logistics Company.</p>
	FirstMileType string `json:"first_mile_type"` // [Required] <p>First mile shipping type. Applicable values:</p><p>Pickup</p><p>Drop off</p>
	FirstMileTrackingNo string `json:"first_mile_tracking_no"` // [Required] <p>First Mile Tracking No.</p>
}
type CancellationOrder struct {
	OrderSn string `json:"order_sn"` // [Required] <p>Order SN.</p>
	CancellationType int64 `json:"cancellation_type"` // [Required] <p>Cancellation Type.&nbsp;Applicable values:&nbsp;</p><p>1: System Cancellation</p><p>2: Seller Cancellation<br /></p>
	DetailedReason int64 `json:"detailed_reason"` // [Required] <p>Reason. Applicable values:&nbsp;</p><p>1001: Return Refund</p><p>1002: Parcel Split Cancellation</p><p>1003: First Mile Pick up fail</p><p>1004: Order inclusion<br /></p><p>10005: Out of Stock<br />10006: Undeliverable area<br />10007: Cannot support COD<br />10008: Logistics request cancelled<br />10009: Logistics pickup failed<br />10010: Logistics not ready<br />10011: Inactive seller<br />10012: Seller did not ship order<br />10013: Order did not reach warehouse<br /></p><p>10014: Seller asked to cancel<br /></p><p>10015: Non-receipt<br />10016: Wrong item<br />10017: Damaged item<br />10018: Incomplete product<br />10019: Fake item<br />10020: Functional Damage<br />10021: Return Refund</p>
}
type FhrOrder struct {
	OrderSn string `json:"order_sn"` // [Required] <p>Order SN.</p>
	ParcelId int64 `json:"parcel_id"` // [Required] <p>Parcel ID.</p>
	ParcelDisplayId string `json:"parcel_display_id"` // [Required] <p>Display Parcel ID.</p>
	ConfirmTime int64 `json:"confirm_time"` // [Required] <p>Confirmed Date.</p>
	HandoverDeadline int64 `json:"handover_deadline"` // [Required] <p>Handover Deadline.<br /></p>
	FastHandoverDueDate int64 `json:"fast_handover_due_date"` // [Required] <p>Fast Handover Due Date.</p>
	ArrangePickUpTime int64 `json:"arrange_pick_up_time"` // [Required] <p>Seller arrange pick up time.</p>
	HandoverTime int64 `json:"handover_time"` // [Required] <p>Parcel drop off / pickup time.<br /></p>
	ShippingChannel string `json:"shipping_channel"` // [Required] <p>Logistics Company.</p>
	FirstMileType string `json:"first_mile_type"` // [Required] <p>First mile shipping type. Applicable values:</p><p>Pickup</p><p>Drop off</p>
	FirstMileTrackingNo string `json:"first_mile_tracking_no"` // [Required] <p>First Mile Tracking No.</p>
	DiagnosisScenario []string `json:"diagnosis_scenario"` // [Required] <p>Diagnosis of the issue.</p>
}
type GetLateOrdersRequest struct {
	PageNo *int64 `json:"page_no,omitempty" url:"page_no,omitempty"` // [Optional] <p>Specifies the page number of data to return in the current call. Starting from 1. if data is more than one page, the page_no can be some entry to start next call.&nbsp;</p><p><br /></p><p>Default is 1.</p>
	PageSize *int64 `json:"page_size,omitempty" url:"page_size,omitempty"` // [Optional] <p>Each result set is returned as a page of entries. Use the "page_size" filters to control the maximum number of entries to retrieve per page (i.e., per call), and the "page_no" to start next call. This integer value is used to specify the maximum number of entries to return in a single "page" of data.&nbsp;</p><p><br /></p><p>The limit of page_size if between 1 and 100. Default is 10.</p>
}
type GetLateOrdersResponse struct {
	BaseResponse // Common response fields
	Response GetLateOrdersResponseData `json:"response"` // 
}
type GetLateOrdersResponseData struct {
	LateOrderList []LateOrder `json:"late_order_list"` // [Required] <p>Late Orders.</p>
	TotalCount int64 `json:"total_count"` // [Required] <p>Total number of late orders.<br /></p>
}
type GetListingsWithIssuesRequest struct {
	PageNo *int64 `json:"page_no,omitempty" url:"page_no,omitempty"` // [Optional] <p>Specifies the page number of data to return in the current call. Starting from 1. if data is more than one page, the page_no can be some entry to start next call.&nbsp;</p><p><br /></p><p>Default is 1.</p>
	PageSize *int64 `json:"page_size,omitempty" url:"page_size,omitempty"` // [Optional] <p>Each result set is returned as a page of entries. Use the "page_size" filters to control the maximum number of entries to retrieve per page (i.e., per call), and the "page_no" to start next call. This integer value is used to specify the maximum number of entries to return in a single "page" of data.&nbsp;</p><p><br /></p><p>The limit of page_size if between 1 and 100. Default is 10.</p>
}
type GetListingsWithIssuesResponse struct {
	BaseResponse // Common response fields
	Response GetListingsWithIssuesResponseData `json:"response"` // 
}
type GetListingsWithIssuesResponseData struct {
	ListingList []Listing `json:"listing_list"` // [Required] <p>Listing with issues.</p>
	TotalCount int64 `json:"total_count"` // [Required] <p>Total number of listing with issues.</p>
}
type GetMetricSourceDetailRequest struct {
	MetricId int64 `json:"metric_id" url:"metric_id"` // [Required] <p>ID of metric. Supported values:&nbsp;</p><p>1: Late Shipment Rate (All Channels)</p><p>3:&nbsp;Non-Fulfilment Rate (All Channels)</p><p>4:&nbsp;Preparation Time</p><p>12:&nbsp;Pre-order Listing %</p><p>15:&nbsp;Days of Pre-order Listing Violation</p><p>25:&nbsp;Fast Handover Rate</p><p>28:&nbsp;On-time Pickup Failure Rate Violation Value<br /></p><p>42:&nbsp;Cancellation Rate (All Channels)</p><p>43:&nbsp;Return-refund Rate (All Channels)</p><p>52:&nbsp;Severe Listing Violations</p><p>53:&nbsp;Other Listing Violations</p><p>85: Late Shipment Rate (NDD)<br /></p><p>88:&nbsp;Non-fulfilment Rate (NDD<br /></p><p>91:&nbsp;Cancellation Rate (NDD)<br /></p><p>92:&nbsp;Return-refund Rate (NDD)</p><p>96: % SDD Listings</p><p>97: % NDD Listings</p><p>2001: Fast Handover Rate - SLS</p><p>2002: Fast Handover Rate - FBS<br /></p><p>2003: Fast Handover Rate - 3PF</p><p>2030: % HD Listings</p><p>2031: % HD Free Shipping Enabled</p><p>2032:&nbsp;Saturday Shipment</p><p>2033:&nbsp;Preparation Time PS</p><p>2033:&nbsp;Preparation Time PS</p><p>2036: OTDR Logistic Rate</p><p>2037:&nbsp;OTDR DD Rate</p>
	PageNo *int64 `json:"page_no,omitempty" url:"page_no,omitempty"` // [Optional] <p>Specifies the page number of data to return in the current call. Starting from 1. if data is more than one page, the page_no can be some entry to start next call.&nbsp;</p><p><br /></p><p>Default is 1.<br /></p>
	PageSize *int64 `json:"page_size,omitempty" url:"page_size,omitempty"` // [Optional] <p>Each result set is returned as a page of entries. Use the "page_size" filters to control the maximum number of entries to retrieve per page (i.e., per call), and the "page_no" to start next call. This integer value is used to specify the maximum number of entries to return in a single "page" of data.&nbsp;</p><p><br /></p><p>The limit of page_size if between 1 and 100. Default is 10.</p>
}
type GetMetricSourceDetailResponse struct {
	BaseResponse // Common response fields
	Response GetMetricSourceDetailResponseData `json:"response"` // 
}
type GetMetricSourceDetailResponseData struct {
	MetricId int64 `json:"metric_id"` // [Required] <p>ID of metric.</p>
	NfrOrderList []NfrOrder `json:"nfr_order_list"` // [Required] <p>Affected Orders for Non-fulfilment Rate.</p><p><br /></p><p>Supported metric_id:&nbsp;</p><p>3:&nbsp;Non-Fulfilment Rate (All Channels)</p><p>88:&nbsp;Non-fulfilment Rate (NDD)</p>
	CancellationOrderList []CancellationOrder `json:"cancellation_order_list"` // [Required] <p>Affected Orders for Cancellation Rate.&nbsp;</p><p><br /></p><p>Supported metric_id:&nbsp;</p><p>42:&nbsp;Cancellation Rate (All Channels)</p><p>91:&nbsp;Cancellation Rate (NDD)</p>
	ReturnRefundOrderList []ReturnRefundOrder `json:"return_refund_order_list"` // [Required] <p>Affected Orders for Return-refund Rate.</p><p><br /></p><p>Supported metric_id:&nbsp;</p><p>43:&nbsp;Return-refund Rate (All Channels)</p><p>92:&nbsp;Return-refund Rate (NDD)</p>
	LsrOrderList []LsrOrder `json:"lsr_order_list"` // [Required] <p>Affected Orders for Late Shipment Rate.</p><p><br /></p><p>Supported metric_id:&nbsp;</p><p>1:&nbsp;Late Shipment Rate (All Channels)</p><p>85:&nbsp;Late Shipment Rate (NDD)</p>
	FhrOrderList []FhrOrder `json:"fhr_order_list"` // [Required] <p>Affected Orders for Fast Handover Rate.</p><p><br /></p><p>Supported metric_id:&nbsp;</p><p>25:&nbsp;Fast Handover Rate</p><p>2001: Fast Handover Rate - SLS</p><p>2002: Fast Handover Rate - FBS<br /></p><p>2003: Fast Handover Rate - 3PF</p>
	OpfrDayDetailDataList []OpfrDayDetailData `json:"opfr_day_detail_data_list"` // [Required] <p>Relevant Violations for OPFR Violation Value.</p><p><br /></p><p>Supported metric_id:&nbsp;</p><p>28:&nbsp;On-time Pickup Failure Rate Violation Value</p>
	ViolationListingList []ViolationListing `json:"violation_listing_list"` // [Required] <p>Relevant Listings for Severe Listing Violations and Other Listing Violations.</p><p><br /></p><p>Supported metric_id:&nbsp;</p><p>52:&nbsp;Severe Listing Violations</p><p>53:&nbsp;Other Listing Violations</p>
	PreOrderListingViolationDataList []PreOrderListingViolationData `json:"pre_order_listing_violation_data_list"` // [Required] <p>Relevant Listings for Days of Pre-order Listing Violation.</p><p><br /></p><p>Supported metric_id:&nbsp;</p><p>15:&nbsp;Days of Pre-order Listing Violation</p>
	PreOrderListingList []PreOrderListing `json:"pre_order_listing_list"` // [Required] <p>Relevant Listings for Pre-order Listing.</p><p><br /></p><p>Supported metric_id:&nbsp;</p><p>12:&nbsp;Pre-order Listing %</p>
	SddListingList []SddListing `json:"sdd_listing_list"` // [Required] <p>Relevant Listings for % SDD Listings.</p><p><br /></p><p>Supported metric_id:&nbsp;</p><p>96: % SDD Listings.</p>
	NddListingList []NddListing `json:"ndd_listing_list"` // [Required] <p>Relevant Listings for % NDD Listings.</p><p><br /></p><p>Supported metric_id:&nbsp;</p><p>97: % NDD Listings.</p>
	AptOrderList []AptOrder `json:"apt_order_list"` // [Required] <p>Affected Parcels for Preparation Time.</p><p><br /></p><p>Supported metric_id:&nbsp;</p><p>4:&nbsp;Preparation Time</p>
	HdListingList *HdListing `json:"hd_listing_list"` // [Required] <p>Relevant Listings for % HD Listings and % HD Free Shipping Enabled.</p><p><br /></p><p>Supported metric_id:&nbsp;</p><p>2030: % HD Listings</p><p>2031: % HD Free Shipping Enabled</p>
	SaturdayShipmentList []AptOrder `json:"saturday_shipment_list"` // [Required] <p>Affected Parcels for Saturday Shipment</p><p><br />Supported metric_id:<br />2032: Saturday Shipment</p>
	OtdrOrderList []OtdrOrder `json:"otdr_order_list"` // [Required] 
	TotalCount int64 `json:"total_count"` // [Required] <p>Total number of Affected Orders or Relevant Listings.</p>
}
type GetPenaltyPointHistoryRequest struct {
	PageNo *int64 `json:"page_no,omitempty" url:"page_no,omitempty"` // [Optional] <p>Specifies the page number of data to return in the current call. Starting from 1. if data is more than one page, the page_no can be some entry to start next call.&nbsp;</p><p><br /></p><p>Default is 1.</p>
	PageSize *int64 `json:"page_size,omitempty" url:"page_size,omitempty"` // [Optional] <p>Each result set is returned as a page of entries. Use the "page_size" filters to control the maximum number of entries to retrieve per page (i.e., per call), and the "page_no" to start next call. This integer value is used to specify the maximum number of entries to return in a single "page" of data.&nbsp;</p><p><br /></p><p>The limit of page_size if between 1 and 100. Default is 10.</p>
	ViolationType *int64 `json:"violation_type,omitempty" url:"violation_type,omitempty"` // [Optional] <p>Applicable values:&nbsp;</p><p>5: High Late Shipment Rate</p><p>6: High Non-fulfilment Rate</p><p>7: High number of non-fulfilled orders</p><p>8: High number of late shipped orders</p><p>9: Prohibited Listings</p><p>10: Counterfeit / IP infringement</p><p>11: Spam</p><p>12: Copy/Steal images</p><p>13: Re-uploading deleted listings with no change</p><p>14: Bought counterfeit from mall</p><p>15: Counterfeit caught by Shopee</p><p>16: High percentage of pre-order listings</p><p>17: Confirmed Fraud attempts (total)</p><p>18: Confirmed Fraud attempts per week (All with vouchers only)</p><p>19: Fake return address</p><p>20: Shipping fraud/abuse</p><p>21: High No. of Non-responded Chat</p><p>22: Rude chat replies</p><p>23: Request buyer to cancel order</p><p>24: Rude reply to buyer's review</p><p>25: Violate Return/Refund policy</p><p>101: Tier Reason</p><p>3026: Misuse of Shopee’s IP</p><p>3028: Violate Shop Name Regulations</p><p>3030: Direct transactions outside of the Shopee platform</p><p>3032: Shipping empty / incomplete parcels</p><p>3034: Severe Violations on Shopee Feed</p><p>3036: Severe Violations on Shopee LIVE</p><p>3038: Misuse of Local Vendor Tag</p><p>3040: Use of misleading shop tag in listing image</p><p>3042: Counterfeit / IP Infringement test</p><p>3044: Repeat Offender - IP infringement and Counterfeit listings</p><p>3046: Violation of Live Animals Selling Policy</p><p>3048: Chat Spam</p><p>3050: High Overseas Return Refunds Rate</p><p>3052: Privacy breach in buyer's review reply</p><p>3054: Order Brushing</p><p>3056: porn image</p><p>3058: Incorrect Product Categories</p><p>3060: Extremely High Non-Fulfilment Rate</p><p>3062: Penalty of Affiliate Marketing Solution (AMS) Overdue Invoice Payment</p><p>3064: Government-related listing</p><p>3066: Listing invalid gifted items</p><p>3068: High non-fulfilment rate (Next Day Delivery Orders)</p><p>3070: High Late Shipment Rate (Next Day Delivery Orders)</p><p>3072: OPFR Violation Value</p><p>3074: Direct transactions outside Shopee platform via chat</p><p>3090: Prohibited Listings-Extreme Violations</p><p>3091: Prohibited Listings-High Violations</p><p>3092: Prohibited Listings-Mid Violations</p><p>3093: Prohibited Listings-Low Violations</p><p>3094: Counterfeit Listings-Extreme Violations</p><p>3095: Counterfeit Listings-High Violations</p><p>3096: Counterfeit Listings-Mid Violations</p><p>3097: Counterfeit Listings-Low Violations</p><p>3098: Spam Listings-Extreme Violations</p><p>3099: Spam Listings-High Violations</p><p>3100: Spam Listings-Mid Violations</p><p>3101: Spam Listings-Low Violations</p><p>3145: Return/Refund Rate (Non-integrated Channel)</p><p>4130: Poor Product Quality</p>
}
type GetPenaltyPointHistoryResponse struct {
	BaseResponse // Common response fields
	Response GetPenaltyPointHistoryResponseData `json:"response"` // 
}
type GetPenaltyPointHistoryResponseData struct {
	PenaltyPointList []PenaltyPoint `json:"penalty_point_list"` // [Required] <p>The penalty point records generated in the current quarter.<br /></p>
	TotalCount int64 `json:"total_count"` // [Required] <p>Total number of penalty point records.</p>
}
type GetPunishmentHistoryRequest struct {
	PageNo *int64 `json:"page_no,omitempty" url:"page_no,omitempty"` // [Optional] <p>Specifies the page number of data to return in the current call. Starting from 1. if data is more than one page, the page_no can be some entry to start next call.&nbsp;</p><p><br /></p><p>Default is 1.</p>
	PageSize *int64 `json:"page_size,omitempty" url:"page_size,omitempty"` // [Optional] <p>Each result set is returned as a page of entries. Use the "page_size" filters to control the maximum number of entries to retrieve per page (i.e., per call), and the "page_no" to start next call. This integer value is used to specify the maximum number of entries to return in a single "page" of data.&nbsp;</p><p><br /></p><p>The limit of page_size if between 1 and 100. Default is 10.</p>
	PunishmentStatus int64 `json:"punishment_status" url:"punishment_status"` // [Required] <p>The status of punishment. Applicable values:&nbsp;</p><p>1: Ongoing</p><p>2: Ended<br /></p>
}
type GetPunishmentHistoryResponse struct {
	BaseResponse // Common response fields
	Response GetPunishmentHistoryResponseData `json:"response"` // 
}
type GetPunishmentHistoryResponseData struct {
	PunishmentList []Punishment `json:"punishment_list"` // [Required] <p>The punishment records generated in the current quarter.<br /></p>
	TotalCount int64 `json:"total_count"` // [Required] <p>Total number of punishment records.<br /></p>
}
type GetShopPerformanceResponse struct {
	BaseResponse // Common response fields
	Response GetShopPerformanceResponseData `json:"response"` // 
}
type GetShopPerformanceResponseData struct {
	OverallPerformance *OverallPerformance `json:"overall_performance"` // [Required] 
	MetricList []Metric `json:"metric_list"` // [Required] 
}
type HdListing struct {
	ItemId int64 `json:"item_id"` // [Required] <p>Item ID.</p>
	CurrentStatus int64 `json:"current_status"` // [Required] <p>For 2030: % HD Listings, it refer to Current HD Status.</p><p>For 2031: % HD Free Shipping Enabled, it refer to Free Shipping Enabled Status.</p><p><br /></p><p>Applicable values:&nbsp;</p><p>1: Yes</p><p>2: No</p>
}
type LateOrder struct {
	OrderSn string `json:"order_sn"` // [Required] <p>Order SN.</p>
	ShippingDeadline int64 `json:"shipping_deadline"` // [Required] <p>Shipping Deadline of this order.</p>
	LateByDays int64 `json:"late_by_days"` // [Required] <p>Late-by Days of this order.</p>
}
type Listing struct {
	ItemId int64 `json:"item_id"` // [Required] <p>Item ID.</p>
	Reason int64 `json:"reason"` // [Required] <p>Reason of this item. Applicable values:&nbsp;</p><p>1: Prohibited</p><p>2: Counterfeit</p><p>3: Spam</p><p>4: Inappropriate Image</p><p>5: Insufficient Info</p><p>6: Mall Listing Improvement</p><p>7: Other Listing Improvement<br /></p>
}
type LsrOrder struct {
	OrderSn string `json:"order_sn"` // [Required] <p>Order SN.</p>
	ShippingDeadline int64 `json:"shipping_deadline"` // [Required] <p>Ship by date.</p>
	ActualShippingTime int64 `json:"actual_shipping_time"` // [Required] <p>Seller arrange shipment time.</p>
	LateByDays int64 `json:"late_by_days"` // [Required] <p>Late-by Days.</p>
	ActualPickUpTime int64 `json:"actual_pick_up_time"` // [Required] <p>Courier actual pick up time.</p>
	ShippingChannel string `json:"shipping_channel"` // [Required] <p>Logistics Company.</p>
	FirstMileType string `json:"first_mile_type"` // [Required] <p>First mile shipping type. Applicable values:</p><p>Pickup</p><p>Drop off</p>
	DiagnosisScenario []string `json:"diagnosis_scenario"` // [Required] <p>Diagnosis of the issue.</p>
}
type Metric struct {
	MetricType int64 `json:"metric_type"` // [Required] <p>Type of metric:&nbsp;</p><p>Fulfillment Performance = 1</p><p>Listing Performance = 2</p><p>Customer Service Performance = 3</p>
	MetricId int64 `json:"metric_id"` // [Required] <p>ID of metric.</p><p>If metric_id &lt; 0, it means that this is not a real metric, but a group of metrics.</p><p><br /></p><p>Non-Responded Chats = -1</p><p>Late Shipment Rate (All Channels) = 1</p><p>Non-Fulfilment Rate (All Channels) = 3</p><p>Preparation Time = 4</p><p>Chat Response Rate = 11</p><p>Pre-order Listing % = 12<br /></p><p>Days of Pre-order Listing Violation = 15</p><p>Response Time = 21</p><p>Shop Rating = 22</p><p>No. of Non-Responded Chats = 23</p><p>Fast Handover Rate = 25</p><p>On-time Pickup Failure Rate = 27</p><p>On-time Pickup Failure Rate Violation Value = 28</p><p>Average Response Time = 29</p><p>Cancellation Rate (All Channels) = 42</p><p>Return-refund Rate (All Channels) = 43</p><p>Severe Listing Violations = 52</p><p>Other Listing Violations = 53</p><p>Prohibited Listings = 54</p><p>Counterfeit/IP infringement = 55</p><p>Spam Listings = 56</p><p>Late Shipment Rate (NDD) = 85</p><p>Non-fulfilment Rate (NDD) = 88</p><p>Cancellation Rate (NDD) = 91</p><p>Return-refund Rate (NDD) = 92<br /></p><p>Customer Satisfaction = 95</p><p>% SDD Listings = 96</p><p>% NDD Listings = 97</p><p>Fast Handover Rate - SLS = 2001</p><p>Fast Handover Rate - FBS = 2002<br /></p><p>Fast Handover Rate - 3PF = 2003</p><p>Poor Quality Products = 2011</p><p>% HD Listings = 2030</p><p>% HD Free Shipping Enabled = 2031</p><p>Saturday Shipment = 2032<br />Preparation Time PS = 2033</p><p>OTDR Logistic Rate = 2036</p><p>OTDR DD Rate = 2037</p>
	ParentMetricId int64 `json:"parent_metric_id"` // [Required] <p>ID of parent metric.<br /></p>
	MetricName string `json:"metric_name"` // [Required] <p>Default name of metric.</p>
	CurrentPeriod float64 `json:"current_period"` // [Required] <p>The performance of the metric at current period.</p>
	LastPeriod float64 `json:"last_period"` // [Required] <p>The performance of the metric at last period.<br /></p>
	Unit int64 `json:"unit"` // [Required] <p>Unit of metric:&nbsp;</p><p>Number = 1</p><p>Percentage = 2</p><p>Second = 3</p><p>Day = 4</p><p>Hour = 5</p>
	Target *Target `json:"target"` // [Required] 
	ExemptionEndDate string `json:"exemption_end_date"` // [Required] <p>(Only for whitelist TW sellers) The exemption_end_date value will not be empty if ALL conditions are met:&nbsp;</p><p>- The shop is in the "POL Shop Whitelist"</p><p>- Within the "Exemption Period"</p><p>- The metric_id is 12 (Pre-order Listing %) or 15 (Days of Pre-order Listing Violation)</p>
}
type NddListing struct {
	ItemId int64 `json:"item_id"` // [Required] <p>Item ID.</p>
	CurrentNddStatus int64 `json:"current_ndd_status"` // [Required] <p>Current NDD Status. Applicable values:&nbsp;</p><p>1: Yes</p><p>0: No</p>
}
type NfrOrder struct {
	OrderSn string `json:"order_sn"` // [Required] <p>Order SN.</p>
	NonFulfillmentType int64 `json:"non_fulfillment_type"` // [Required] <p>Non-fulfilment type. Applicable values:&nbsp;<br /></p><p>1: System Cancellation</p><p>2: Seller Cancellation</p><p>3: Return Refunds<br /></p>
	DetailedReason int64 `json:"detailed_reason"` // [Required] <p>Reason. Applicable values:&nbsp;</p><p>1001: Return Refund</p><p>1002: Parcel Split Cancellation</p><p>1003: First Mile Pick up fail</p><p>1004: Order inclusion<br /></p><p>10005: Out of Stock<br />10006: Undeliverable area<br />10007: Cannot support COD<br />10008: Logistics request cancelled<br />10009: Logistics pickup failed<br />10010: Logistics not ready<br />10011: Inactive seller<br />10012: Seller did not ship order<br />10013: Order did not reach warehouse<br /></p><p>10014: Seller asked to cancel<br /></p><p>10015: Non-receipt<br />10016: Wrong item<br />10017: Damaged item<br />10018: Incomplete product<br />10019: Fake item<br />10020: Functional Damage<br />10021: Return Refund<br /></p>
}
type OpfrDayDetailData struct {
	Date string `json:"date"` // [Required] <p>Date.</p>
	ScheduledPickupNum int64 `json:"scheduled_pickup_num"` // [Required] <p>Number of scheduled pickups.<br /></p>
	FailedPickupNum int64 `json:"failed_pickup_num"` // [Required] <p>Number of failed pickups.<br /></p>
	Opfr int64 `json:"opfr"` // [Required] <p>OPFR.<br /></p>
	Target string `json:"target"` // [Required] <p>Target.</p>
}
type OtdrOrder struct {
	OrderId string `json:"order_id"` // [Required] <p>Order ID.</p>
	OrderSn string `json:"order_sn"` // [Required] <p>Order SN.</p>
	PaidTime int64 `json:"paid_time"` // [Required] <p>Order Paid Time.</p>
	EstimatedDeliveryDate int64 `json:"estimated_delivery_date"` // [Required] <p>Estimated delivery date</p>
	ActualPickUpTime int64 `json:"actual_pick_up_time"` // [Required] <p>Actual pick up time</p>
	RealDeliveryTime int64 `json:"real_delivery_time"` // [Required] <p>Real delivery time</p>
	DifferenceBetweenEddRdt string `json:"difference_between_edd_rdt"` // [Required] <p>Difference days between estimated delivery date and real delivery time</p>
}
type OverallPerformance struct {
	Rating int64 `json:"rating"` // [Required] <p>Overall Performance:&nbsp;</p><p>Poor = 1<br /></p><p>ImprovementNeeded = 2<br /></p><p>Good = 3<br /></p><p>Excellent = 4</p>
	FulfillmentFailed int64 `json:"fulfillment_failed"` // [Required] <p>The number of metrics that did not meet target under Fulfillment Performance type.</p>
	ListingFailed int64 `json:"listing_failed"` // [Required] <p>The number of metrics that did not meet target under Listing Performance type.<br /></p>
	CustomServiceFailed int64 `json:"custom_service_failed"` // [Required] <p>The number of metrics that did not meet target under Customer Service Performance type.</p>
}
type PenaltyPoint struct {
	IssueTime int64 `json:"issue_time"` // [Required] <p>The time when penalty points are issued.</p>
	LatestPointNum int64 `json:"latest_point_num"` // [Required] <p>The latest penalty points issued under current penalty point record.&nbsp;</p><p><br /></p><p>If seller raised appeal for this penalty point record and the appeal has been approved and Shopee adjusted the penalty point, then the original_point_num returns the penalty point before the adjustment, and latest_point_num returns the penalty point after the adjustment.<br /></p>
	OriginalPointNum int64 `json:"original_point_num"` // [Required] <p>The original penalty points&nbsp;issued under current penalty point record.</p><p><br /></p><p>If seller raised appeal for this penalty point record and the appeal has been approved and Shopee adjusted the penalty point, then the original_point_num returns the penalty point before the adjustment, and latest_point_num returns the penalty point after the adjustment.</p>
	ReferenceId int64 `json:"reference_id"` // [Required] <p>Reference ID for this penalty point record.</p>
	ViolationType int64 `json:"violation_type"` // [Required] <p>Applicable values:&nbsp;</p><p>5: High Late Shipment Rate</p><p>6: High Non-fulfilment Rate</p><p>7: High number of non-fulfilled orders</p><p>8: High number of late shipped orders</p><p>9: Prohibited Listings</p><p>10: Counterfeit / IP infringement</p><p>11: Spam</p><p>12: Copy/Steal images</p><p>13: Re-uploading deleted listings with no change</p><p>14: Bought counterfeit from mall</p><p>15: Counterfeit caught by Shopee</p><p>16: High percentage of pre-order listings</p><p>17: Confirmed Fraud attempts (total)</p><p>18: Confirmed Fraud attempts per week (All with vouchers only)</p><p>19: Fake return address</p><p>20: Shipping fraud/abuse</p><p>21: High No. of Non-responded Chat</p><p>22: Rude chat replies</p><p>23: Request buyer to cancel order</p><p>24: Rude reply to buyer's review</p><p>25: Violate Return/Refund policy</p><p>101: Tier Reason</p><p>3026: Misuse of Shopee’s IP</p><p>3028: Violate Shop Name Regulations</p><p>3030: Direct transactions outside of the Shopee platform</p><p>3032: Shipping empty / incomplete parcels</p><p>3034: Severe Violations on Shopee Feed</p><p>3036: Severe Violations on Shopee LIVE</p><p>3038: Misuse of Local Vendor Tag</p><p>3040: Use of misleading shop tag in listing image</p><p>3042: Counterfeit / IP Infringement test</p><p>3044: Repeat Offender - IP infringement and Counterfeit listings</p><p>3046: Violation of Live Animals Selling Policy</p><p>3048: Chat Spam</p><p>3050: High Overseas Return Refunds Rate</p><p>3052: Privacy breach in buyer's review reply</p><p>3054: Order Brushing</p><p>3056: porn image</p><p>3058: Incorrect Product Categories</p><p>3060: Extremely High Non-Fulfilment Rate</p><p>3062: Penalty of Affiliate Marketing Solution (AMS) Overdue Invoice Payment</p><p>3064: Government-related listing</p><p>3066: Listing invalid gifted items</p><p>3068: High non-fulfilment rate (Next Day Delivery Orders)</p><p>3070: High Late Shipment Rate (Next Day Delivery Orders)</p><p>3072: OPFR Violation Value</p><p>3074: Direct transactions outside Shopee platform via chat</p><p>3090: Prohibited Listings-Extreme Violations</p><p>3091: Prohibited Listings-High Violations</p><p>3092: Prohibited Listings-Mid Violations</p><p>3093: Prohibited Listings-Low Violations</p><p>3094: Counterfeit Listings-Extreme Violations</p><p>3095: Counterfeit Listings-High Violations</p><p>3096: Counterfeit Listings-Mid Violations</p><p>3097: Counterfeit Listings-Low Violations</p><p>3098: Spam Listings-Extreme Violations</p><p>3099: Spam Listings-High Violations</p><p>3100: Spam Listings-Mid Violations</p><p>3101: Spam Listings-Low Violations</p><p>3145: Return/Refund Rate (Non-integrated Channel)</p><p>4130: Poor Product Quality</p>
}
type PreOrderListing struct {
	ItemId int64 `json:"item_id"` // [Required] <p>Item ID.</p>
	CurrentPreOrderStatus int64 `json:"current_pre_order_status"` // [Required] <p>Current Pre-order Status. Applicable values:&nbsp;</p><p>1: Yes</p><p>2: No<br /></p>
}
type PreOrderListingViolationData struct {
	Date string `json:"date"` // [Required] <p>Date.</p>
	LiveListingCount int64 `json:"live_listing_count"` // [Required] <p>Number of Live Listings.<br /></p>
	PreOrderListingCount int64 `json:"pre_order_listing_count"` // [Required] <p>Number of pre-order Listings.<br /></p>
	PreOrderListingRate int64 `json:"pre_order_listing_rate"` // [Required] <p>Pre-order Listing %.<br /></p>
	Target string `json:"target"` // [Required] <p>Target.<br /></p>
}
type Punishment struct {
	IssueTime int64 `json:"issue_time"` // [Required] <p>The time when punishment are issued.<br /></p>
	StartTime int64 `json:"start_time"` // [Required] <p>Start time in the duration of this punishment record.</p>
	EndTime int64 `json:"end_time"` // [Required] <p>End time in the duration of this punishment record.<br /></p>
	PunishmentType int64 `json:"punishment_type"` // [Required] <p>Punishment Type of this punishment record. Applicable values:&nbsp;</p><p>103: Listings not displayed in category browsing<br />104: Listings not displayed in search<br />105: Unable to create new listings<br />106: Unable to edit listings<br />107: Unable to join marketing campaigns<br />108: No shipping subsidies<br />109: Account is suspended<br />600: Listings not displayed in search<br />601: Shop listings hide from recommendation<br />602: Listings not displayed in category browsing<br />1109: Listing Limit is reduced<br />1110: Listing Limit is reduced<br />1111: Listing Limit is reduced<br />1112: Listing Limit is reduced<br /></p><p>2008: Order Limit</p>
	Reason int64 `json:"reason"` // [Required] <p>Reason of this punishment record. Applicable values:&nbsp;</p><p>1: Tier 1</p><p>2: Tier 2</p><p>3: Tier 3</p><p>4: Tier 4</p><p>5: Tier 5</p><p>1109: Listing Limit Tier 1</p><p>1110: Listing Limit Tier 2</p><p>1111: Listing Limit POL<br /></p>
	ReferenceId int64 `json:"reference_id"` // [Required] <p>Reference ID for this punishment record.<br /></p>
	ListingLimit int64 `json:"listing_limit"` // [Required] <p>Return the specific value of listing limit when punishment_type is:&nbsp;</p><p>1109: Listing Limit is reduced<br />1110: Listing Limit is reduced<br />1111: Listing Limit is reduced<br />1112: Listing Limit is reduced</p>
	OrderLimit string `json:"order_limit"` // [Required] <p>Return the specific percentage of order limit when punishment_type is:&nbsp;</p><p>2008: Order Limit</p><p><br /></p><p>Daily Order Limit = X % * L28D ADO (Average Daily Order of this Shop in Past 28 Days)</p>
}
type ReturnRefundOrder struct {
	OrderSn string `json:"order_sn"` // [Required] <p>Order SN.</p>
	DetailedReason int64 `json:"detailed_reason"` // [Required] <p>Reason. Applicable values:&nbsp;</p><p>1001: Return Refund</p><p>1002: Parcel Split Cancellation</p><p>1003: First Mile Pick up fail</p><p>1004: Order inclusion<br /></p><p>10005: Out of Stock<br />10006: Undeliverable area<br />10007: Cannot support COD<br />10008: Logistics request cancelled<br />10009: Logistics pickup failed<br />10010: Logistics not ready<br />10011: Inactive seller<br />10012: Seller did not ship order<br />10013: Order did not reach warehouse<br /></p><p>10014: Seller asked to cancel<br /></p><p>10015: Non-receipt<br />10016: Wrong item<br />10017: Damaged item<br />10018: Incomplete product<br />10019: Fake item<br />10020: Functional Damage<br />10021: Return Refund</p>
}
type SddListing struct {
	ItemId int64 `json:"item_id"` // [Required] <p>Item ID.</p>
	CurrentSddStatus int64 `json:"current_sdd_status"` // [Required] <p>Current SDD Status. Applicable values:&nbsp;</p><p>1: Yes</p><p>0: No</p>
}
type Target struct {
	Value float64 `json:"value"` // [Required] <p>Value of target.</p>
	Comparator string `json:"comparator"` // [Required] <p>Comparator of target: &lt;, &lt;=, &gt;, &gt;=, =</p>
}
type ViolationListing struct {
	ItemId int64 `json:"item_id"` // [Required] <p>Item ID.</p>
	DetailedReason int64 `json:"detailed_reason"` // [Required] <p>Reason. Applicable values:&nbsp;</p><p>1: Prohibited</p><p>2: Counterfeit</p><p>3: Spam</p><p>4: Inappropriate Image</p><p>5: Insufficient Info</p><p>6: Mall Listing Improvement</p><p>7: Other Listing Improvement<br /></p><p>8: PQR Products</p>
	UpdateTime int64 `json:"update_time"` // [Required] <p>Updated on.</p>
}
