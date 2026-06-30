package goshopee

type AllowInfo struct {
	AllowStitch bool `json:"allow_stitch"` // [Required] <p>Whether allow stitch.</p>
	AllowDuet bool `json:"allow_duet"` // [Required] <p>Whether allow duet.</p>
}
type Conversion struct {
	PlacedBuyers int64 `json:"placed_buyers"` // [Required] <p>Number of unique buyers who placed order from the video.</p>
	ConfirmedBuyers int64 `json:"confirmed_buyers"` // [Required] <p>Number of unique buyers who confirmed order from the video.</p>
	TotalAtc int64 `json:"total_atc"` // [Required] <p>Number of "Add To Cart" button clicked for all products in the orange bag during video viewing.</p>
	Ctr float64 `json:"ctr"` // [Required] <p>Number of products clicks divided by Number of video views.</p>
	PlacedCoRate float64 `json:"placed_co_rate"` // [Required] <p>Number of placed product orders from the video divided by Number of product clicks from the video.</p>
	ConfirmedCoRate float64 `json:"confirmed_co_rate"` // [Required] <p>Number of confirmed product orders from the video divided by Number of product clicks from the video.</p>
	PlacedAbs float64 `json:"placed_abs"` // [Required] <p>Total placed sales divided by Total placed orders.</p>
	ConfirmedAbs float64 `json:"confirmed_abs"` // [Required] <p>Total confirmed sales divided by Total confirmed orders.</p>
	PlacedGpm float64 `json:"placed_gpm"` // [Required] <p>The placed Sales generated for every 1,000 views.</p>
	ConfirmedGpm float64 `json:"confirmed_gpm"` // [Required] <p>The confirmed Sales generated for every 1,000 views.</p>
	VideoWithProducts int64 `json:"video_with_products"` // [Required] <p>Videos with at least one product in the orange bag.</p>
	PlacedRevenueGeneratingVideos int64 `json:"placed_revenue_generating_videos"` // [Required] <p>Videos that generates placed revenues.</p>
	ConfirmedRevenueGeneratingVideos int64 `json:"confirmed_revenue_generating_videos"` // [Required] <p>Videos that generates confirmed revenues.</p>
}
type DeleteVideoRequest struct {
	PostIdList []string `json:"post_id_list,omitempty"` // [Optional] <p>You can only select one from video_upload_id_list and post_id_list:&nbsp;</p><p>- If you want to delete video with draft status, please pass&nbsp;video_upload_id_list.</p><p>- If you want to delete video with post status, please pass&nbsp;post_id_list.</p>
	VideoUploadIdList []string `json:"video_upload_id_list,omitempty"` // [Optional] <p>You can only select one from video_upload_id_list and post_id_list:&nbsp;</p><p>- If you want to delete video with draft status, please pass&nbsp;video_upload_id_list.</p><p>- If you want to delete video with post status, please pass&nbsp;post_id_list.</p>
}
type DeleteVideoResponse struct {
	BaseResponse // Common response fields
	Response DeleteVideoResponseData `json:"response"` // <p>Detail informations you are querying.</p>
}
type DeleteVideoResponseData struct {
	SuccessList []DeleteVideoResponseDataSuccess `json:"success_list"` // [Required] <p>The list of video delete successfully.</p>
	FailureList []DeleteVideoResponseDataFailure `json:"failure_list"` // [Required] <p>The list of video delete deleted.</p>
}
type DeleteVideoResponseDataFailure struct {
	FailVideoUploadId string `json:"fail_video_upload_id"` // [Required] <p>Failed video_upload_id.</p>
	FailPostId string `json:"fail_post_id"` // [Required] <p>Failed post_id.</p>
	FailedReason string `json:"failed_reason"` // [Required] <p>Failed&nbsp;reason of the corresponding video_upload_id or post_id.</p>
}
type DeleteVideoResponseDataSuccess struct {
	SuccessVideoUploadId string `json:"success_video_upload_id"` // [Required] <p>The video_upload_id delete successfully.</p>
	SuccessPostId string `json:"success_post_id"` // [Required] <p>The post_id delete successfully.</p>
}
type EditVideoInfoRequest struct {
	VideoUploadList []VideoUpload `json:"video_upload_list"` // [Required] <p>Video information collection, no more than 5.</p>
}
type EditVideoInfoResponse struct {
	BaseResponse // Common response fields
	Response EditVideoInfoResponseData `json:"response"` // <p>Detail informations you are querying.</p>
}
type EditVideoInfoResponseData struct {
	SuccessList []string `json:"success_list"` // [Required] <p>The list of video_upload_id edit successfully.</p>
	FailureList []EditVideoInfoResponseDataFailure `json:"failure_list"` // [Required] <p>The list of video_upload_id edit failed.</p>
}
type EditVideoInfoResponseDataFailure struct {
	FailVideoUploadId string `json:"fail_video_upload_id"` // [Required] <p>Failed&nbsp;video_upload_id.</p>
	FailedReason string `json:"failed_reason"` // [Required] <p>Failed&nbsp;reason of the corresponding video_upload_id.</p>
}
type Engagement struct {
	TotalViews int64 `json:"total_views"` // [Required] <p>Number of views from all videos</p>
	TotalLikes int64 `json:"total_likes"` // [Required] <p>Number of likes from all videos</p>
	TotalShares int64 `json:"total_shares"` // [Required] <p>Number of shares from all videos</p>
	TotalComments int64 `json:"total_comments"` // [Required] <p>Number of comments from all videos</p>
	VideoNewFollowers int64 `json:"video_new_followers"` // [Required] <p>Number of new followers from all videos</p>
}
type GetCoverListRequest struct {
	VideoUploadId string `json:"video_upload_id" url:"video_upload_id"` // [Required] <p>ID of uploaded video. Obtain from v2.media.get_video_upload_result.</p>
}
type GetCoverListResponse struct {
	BaseResponse // Common response fields
	Response GetCoverListResponseData `json:"response"` // <p>Detail informations you are querying.</p>
}
type GetCoverListResponseData struct {
	ImageUrlList []string `json:"image_url_list"` // [Required] <p>List of image url for each frame of the uploaded video, you can select one as the video cover when calling v2.video.edit_video_info.</p>
}
type GetMetricTrendRequest struct {
	EndDate string `json:"end_date" url:"end_date"` // [Required] <p>The end_date format should be "YYYY-MM-DD".</p><p>- For Day, Last7d, Last15d, and Last30d, the end_date must before current day.</p><p>- For Week, the end_date must be Sunday and must be less than or equal to the current week.</p><p>- For Month, the end_date must be the end of the month and must be less than or equal to the current month.</p>
	PeriodType string `json:"period_type" url:"period_type"` // [Required] <p>Period Type. Applicable values:</p><p>Day<br />Week<br />Month<br /></p><p>Last7d</p><p>Last15d</p><p>Last30d</p><p><br /></p><p>Note: The end date must align with the Period Type.</p>
}
type GetMetricTrendResponse struct {
	BaseResponse // Common response fields
	Response GetMetricTrendResponseData `json:"response"` // <p>Detail informations you are querying.</p>
}
type GetMetricTrendResponseData struct {
	VideoTotalMetricList []VideoTotalMetric `json:"video_total_metric_list"` // [Required] 
}
type GetOverviewPerformanceRequest struct {
	EndDate string `json:"end_date" url:"end_date"` // [Required] <p>The end_date format should be "YYYY-MM-DD".</p><p>- For Day, Last7d, Last15d, and Last30d, the end_date must before current day.</p><p>- For Week, the end_date must be Sunday and must be less than or equal to the current week.</p><p>- For Month, the end_date must be the end of the month and must be less than or equal to the current month.</p>
	PeriodType string `json:"period_type" url:"period_type"` // [Required] <p>Period Type. Applicable values:</p><p>Day<br />Week<br />Month<br /></p><p>Last7d</p><p>Last15d</p><p>Last30d</p><p><br /></p><p>Note: The end date must align with the Period Type.</p>
}
type GetOverviewPerformanceResponse struct {
	BaseResponse // Common response fields
	Response GetOverviewPerformanceResponseData `json:"response"` // 
}
type GetOverviewPerformanceResponseData struct {
	KeyMetric *KeyMetric `json:"key_metric"` // [Required] 
	Conversion *Conversion `json:"conversion"` // [Required] 
	Engagement *Engagement `json:"engagement"` // [Required] 
	FetchedDateRange string `json:"fetched_date_range"` // [Required] <p>Data offline computation time.</p>
}
type GetProdcutPerformanceListRequest struct {
	EndDate string `json:"end_date" url:"end_date"` // [Required] <p>The end_date format should be "YYYY-MM-DD".</p><p>- For Day, Last7d, Last15d, and Last30d, the end_date must before current day.</p><p>- For Week, the end_date must be Sunday and must be less than or equal to the current week.</p><p>- For Month, the end_date must be the end of the month and must be less than or equal to the current month.</p>
	ItemId *int64 `json:"item_id,omitempty" url:"item_id,omitempty"` // [Optional] <p>Shopee's unique identifier for an item.</p>
	ItemName *string `json:"item_name,omitempty" url:"item_name,omitempty"` // [Optional] <p>Search by product name.</p>
	OrderBy string `json:"order_by" url:"order_by"` // [Required] <p>Use this field to specify which field to use to sort the returned list. Available values:</p><p>PlacedOrders</p><p>PlacedSales</p><p>PlacedUniqueBuyers</p><p>ConfirmedOrders</p><p>ConfirmedSales</p><p>ConfirmedUniqueBuyers</p>
	PageNo int64 `json:"page_no" url:"page_no"` // [Required] <p>The start index of request. Starting from 1.</p>
	PageSize int64 `json:"page_size" url:"page_size"` // [Required] <p>The number of item returned by this request. Max is 20.</p>
	PeriodType string `json:"period_type" url:"period_type"` // [Required] <p>Period Type. Applicable values:</p><p>Day<br />Week<br />Month<br /></p><p>Last7d</p><p>Last15d</p><p>Last30d</p><p><br /></p><p>Note: The end date must align with the Period Type.</p>
	Sort string `json:"sort" url:"sort"` // [Required] <p>Use this field to specify whether the returned list is sorted in ascending or descending order_by.&nbsp;Available values:</p><p>asc</p><p>desc</p>
}
type GetProdcutPerformanceListResponse struct {
	BaseResponse // Common response fields
	Response GetProdcutPerformanceListResponseData `json:"response"` // <p>Detail informations you are querying.</p>
}
type GetProdcutPerformanceListResponseData struct {
	TotalCount int64 `json:"total_count"` // [Required] <p>The total count of product that match the condition.</p>
	HasMore bool `json:"has_more"` // [Required] <p>This is to indicate whether the video list is more than one page. If this value is true, you may want to continue to check next page to retrieve the rest of data.</p>
	List []GetProdcutPerformanceListResponseDataList `json:"list"` // [Required] <p>The list of product that match the condition.</p>
}
type GetProdcutPerformanceListResponseDataList struct {
	ShopId int64 `json:"shop_id"` // [Required] <p>Shopee's unique identifier for a shop.</p>
	ItemId int64 `json:"item_id"` // [Required] <p>Shopee's unique identifier for an item.</p>
	ItemName string `json:"item_name"` // [Required] <p>Name of the item.</p>
	ItemCoverImageUrl string `json:"item_cover_image_url"` // [Required] <p>Cover image url of the item.</p>
	ItemDescription string `json:"item_description"` // [Required] <p>Description of the item.</p>
	PlacedOrders int64 `json:"placed_orders"` // [Required] <p>The number of placed orders for the item.</p>
	ConfirmedOrders int64 `json:"confirmed_orders"` // [Required] <p>The number of confirmed orders for the item.</p>
	PlacedSales float64 `json:"placed_sales"` // [Required] <p>The placed value of orders for the item.</p>
	ConfirmedSales float64 `json:"confirmed_sales"` // [Required] <p>The confirmed value of orders for the item.</p>
	PlacedUniqueBuyers int64 `json:"placed_unique_buyers"` // [Required] <p>Number of unique buyers who placed order for the item.</p>
	ConfirmedUniqueBuyers int64 `json:"confirmed_unique_buyers"` // [Required] <p>Number of unique buyers who confirmed order for the item.</p>
	FetchedDateRange string `json:"fetched_date_range"` // [Required] <p>Data Date Range.</p>
}
type GetUserDemographicsResponse struct {
	BaseResponse // Common response fields
	Response GetUserDemographicsResponseData `json:"response"` // 
}
type GetUserDemographicsResponseData struct {
	Age interface{} `json:"age"` // [Required] <p>The age distribution of your viewers.</p><p><br /></p><p>Note: The type of age is a map. The key is an enumerated value corresponding to an age range:&nbsp;</p><p>-1: Unknown</p><p>1: 18-24 years old</p><p>2: 25-34 years old</p><p>3: 35-44 years old</p><p>4: 45+ years old</p><p><br /></p><p>The value is the number of viewers in each age group.</p>
	Gender interface{} `json:"gender"` // [Required] <p>The gender distribution of your viewers.</p><p><br /></p><p>Note: The type of gender is a map. The key is one of:&nbsp;</p><p>Male</p><p>Female</p><p>Predicted Male</p><p>Predicted Female</p><p><br /></p><p>The value is the number of viewers for each gender type.</p>
	Location interface{} `json:"location"` // [Required] <p>The geographic distribution of your viewers.&nbsp;</p><p><br /></p><p>Note: The type of location is a map. The key is top 10 city, and the value is the number of viewers in each city.</p>
	Identity interface{} `json:"identity"` // [Required] <p>The distribution of viewers based on whether they follow your Shopee Video profile.</p><p><br /></p><p>Note: The type of identity is a map. The key is either "follow" or "unfollow", indicating followers and non-followers respectively, and the value is number of page views generated by each group.</p>
	Activity interface{} `json:"activity"` // [Required] <p>The distribution of video views across different hours of the day.</p><p><br /></p><p>Note: The type of activity is a map. The key is the hour of the day (ranging from 0 to 23), and the value is the number of video views generated during that specific hour.</p>
	Content interface{} `json:"content"` // [Required] <p>The types of videos that your viewer is most interested in.<br /></p><p><br /></p><p>Note: The type of content is a map. The key is top 10 content category, and the value is the number of video views corresponding to that content category.</p>
	Shopping interface{} `json:"shopping"` // [Required] <p>The types of products that your viewers is most interested in.<br /></p><p><br /></p><p>Note: The type of shopping is a map. The key is top 10 product category, and the value is the number of video views corresponding to that product category.</p>
}
type GetVideoDetailAudienceDistributionRequest struct {
	PostId string `json:"post_id" url:"post_id"` // [Required] <p>A unique identifier for Shopee videos.</p>
}
type GetVideoDetailAudienceDistributionResponse struct {
	BaseResponse // Common response fields
	Response GetVideoDetailAudienceDistributionResponseData `json:"response"` // <p>Detail informations you are querying.</p>
}
type GetVideoDetailAudienceDistributionResponseData struct {
	Age interface{} `json:"age"` // [Required] <p>The age distribution of your viewers.</p><p><br /></p><p>Note: The type of age is a map. The key is an enumerated value corresponding to an age range:&nbsp;</p><p>-1: Unknown</p><p>1: 18-24 years old</p><p>2: 25-34 years old</p><p>3: 35-44 years old</p><p>4: 45+ years old</p><p><br /></p><p>The value is the number of viewers in each age group.</p>
	Gender interface{} `json:"gender"` // [Required] <p>The gender distribution of your viewers.</p><p><br /></p><p>Note: The type of gender is a map. The key is one of:&nbsp;</p><p>male</p><p>female</p><p>predictedMale</p><p>predictedFemale</p><p>unknown</p><p><br /></p><p>The value is the number of viewers for each gender type.</p>
	Location interface{} `json:"location"` // [Required] <p>The geographic distribution of your viewers.&nbsp;</p><p><br /></p><p>Note: The type of location is a map. The key is top 10 city, and the value is the number of viewers in each city.</p>
	Identity interface{} `json:"identity"` // [Required] <p>The distribution of viewers based on whether they follow your Shopee Video profile.</p><p><br /></p><p>Note: The type of identity is a map. The key is one of:&nbsp;</p><p>0: Non-follower</p><p>1: Follower</p><p><br /></p><p>The value is number of user views generated by each group.</p>
	Activity interface{} `json:"activity"` // [Required] <p>The distribution of video views across different hours of the day.</p><p><br /></p><p>Note: The type of activity is a map. The key is the hour of the day (ranging from 0 to 23), and the value is the number of video views generated during that specific hour.</p>
	Content interface{} `json:"content"` // [Required] <p>The types of videos that your viewer is most interested in.<br /></p><p><br /></p><p>Note: The type of content is a map. The key is content category, and the value is the number of video views corresponding to that content category.</p>
	Shopping interface{} `json:"shopping"` // [Required] <p>The types of products that your viewers is most interested in.<br /></p><p><br /></p><p>Note: The type of shopping is a map. The key is product category, and the value is the number of video views corresponding to that product category.</p>
}
type GetVideoDetailMetricTrendRequest struct {
	MetricName string `json:"metric_name" url:"metric_name"` // [Required] <p>The name of metric that require obtaining trend data. Applicable values:&nbsp;</p><p>Views, Likes, Comments, Shares, FollowersGrowth, PlacedOrders, PlacedSales, UniqueBuyers, ConversionRate, SoldItems, SalesPerOrder, SalesPerBuyer</p>
	PostId string `json:"post_id" url:"post_id"` // [Required] <p>A unique identifier for Shopee videos.</p>
}
type GetVideoDetailMetricTrendResponse struct {
	BaseResponse // Common response fields
	Response GetVideoDetailMetricTrendResponseData `json:"response"` // <p>Detail informations you are querying.</p>
}
type GetVideoDetailMetricTrendResponseData struct {
	MetricTrend interface{} `json:"metric_trend"` // [Required] <p>The type of metric_trend is a map. The key is date (in millisecond timestamp format), and the value is the number corresponding to metric.</p>
}
type GetVideoDetailPerformanceRequest struct {
	PostId string `json:"post_id" url:"post_id"` // [Required] <p>A unique identifier for Shopee videos.</p>
}
type GetVideoDetailPerformanceResponse struct {
	BaseResponse // Common response fields
	Response GetVideoDetailPerformanceResponseData `json:"response"` // <p>Detail informations you are querying.</p>
}
type GetVideoDetailPerformanceResponseData struct {
	VideoInfo *GetVideoDetailPerformanceResponseDataVideoInfo `json:"video_info"` // [Required] <p>Video post detail informations you are querying.</p>
	VideoPerformance *VideoPerformance `json:"video_performance"` // [Required] <p>Overall performance data of the video you are querying.</p>
}
type GetVideoDetailPerformanceResponseDataVideoInfo struct {
	PostId string `json:"post_id"` // [Required] <p>A unique identifier for Shopee videos.</p>
	PostTime int64 `json:"post_time"` // [Required] <p>The time when the video post to Shopee Video.</p>
	VideoUrl string `json:"video_url"` // [Required] <p>Video play url.</p>
	CoverImageUrl string `json:"cover_image_url"` // [Required] <p>Cover image url of the Shopee Video.</p>
	Caption string `json:"caption"` // [Required] <p>Description of the Shopee Video.</p>
	Duration int64 `json:"duration"` // [Required] <p>Video duration time in millisecond.</p>
	RelatedItemCount int64 `json:"related_item_count"` // [Required] <p>Number of products linked with the Shopee Video.</p>
}
type GetVideoDetailProductPerformanceRequest struct {
	ItemId *int64 `json:"item_id,omitempty" url:"item_id,omitempty"` // [Optional] <p>Shopee's unique identifier for an item.</p>
	ItemName *string `json:"item_name,omitempty" url:"item_name,omitempty"` // [Optional] <p>Name of the item.</p>
	PageNo int64 `json:"page_no" url:"page_no"` // [Required] <p>The start index of request. Starting from 1.</p>
	PageSize int64 `json:"page_size" url:"page_size"` // [Required] <p>The number of item returned by this request. Max is 20.</p>
	PostId string `json:"post_id" url:"post_id"` // [Required] <p>The unique identifier for post Shopee Video.</p>
}
type GetVideoDetailProductPerformanceResponse struct {
	BaseResponse // Common response fields
	Response GetVideoDetailProductPerformanceResponseData `json:"response"` // <p>Detail informations you are querying.</p>
}
type GetVideoDetailProductPerformanceResponseData struct {
	List []GetVideoDetailProductPerformanceResponseDataList `json:"list"` // [Required] <p>The list of item that match the condition.</p>
	TotalCount int64 `json:"total_count"` // [Required] <p>The total count of video that match the condition.</p>
	HasMore bool `json:"has_more"` // [Required] <p>This is to indicate whether the video list is more than one page. If this value is true, you may want to continue to check next page to retrieve the rest of data.</p>
}
type GetVideoDetailProductPerformanceResponseDataList struct {
	ShopId int64 `json:"shop_id"` // [Required] <p>Shopee's unique identifier for a shop.&nbsp;</p>
	ItemId int64 `json:"item_id"` // [Required] <p>Shopee's unique identifier for an item.</p>
	ItemName string `json:"item_name"` // [Required] <p>Name of the item.</p>
	ItemCoverImageUrl string `json:"item_cover_image_url"` // [Required] <p>Cover image url of the item.</p>
	ItemDescription string `json:"item_description"` // [Required] <p>Description of the item.</p>
	Likes int64 `json:"likes"` // [Required] <p>Like count the post Shopee Video.</p>
	Comments int64 `json:"comments"` // [Required] <p>Comment count the post Shopee Video.</p>
	PlacedOrders int64 `json:"placed_orders"` // [Required] <p>Amount of product orders from the video.</p>
	PlacedSales float64 `json:"placed_sales"` // [Required] <p>Amount of product sales from the video.</p>
	UniqueBuyers int64 `json:"unique_buyers"` // [Required] <p>Buyers of the product in the video.</p>
	SoldItems int64 `json:"sold_items"` // [Required] <p>Amount of products sold from the video.</p>
	ProductClicks int64 `json:"product_clicks"` // [Required] <p>Amount of product clicks from the video.</p>
	ProductClickRate float64 `json:"product_click_rate"` // [Required] <p>Amount of product clicks from the video/Product view from video.</p>
	ConversionRate float64 `json:"conversion_rate"` // [Required] <p>Amount of products sold from the video/amount of views from the video.</p>
	SalesPerOrder float64 `json:"sales_per_order"` // [Required] <p>Amount of product sales from the video/amount of product orders from the video.</p>
	SalesPerBuyer float64 `json:"sales_per_buyer"` // [Required] <p>Amount of product sales from the video/amount of buyers from the video.</p>
}
type GetVideoDetailRequest struct {
	PostId *string `json:"post_id,omitempty" url:"post_id,omitempty"` // [Optional] <p>You can only select one from video_upload_id and post_id:&nbsp;</p><p>- If you want to get detail information of video with draft status, please pass&nbsp;video_upload_id.</p><p>- If you want to get detail information of video with post status, please pass&nbsp;post_id.</p>
	VideoUploadId *string `json:"video_upload_id,omitempty" url:"video_upload_id,omitempty"` // [Optional] <p>You can only select one from video_upload_id and post_id:&nbsp;</p><p>- If you want to get detail information of video with draft status, please pass&nbsp;video_upload_id.</p><p>- If you want to get detail information of video with post status, please pass&nbsp;post_id.</p>
}
type GetVideoDetailResponse struct {
	BaseResponse // Common response fields
	Response GetVideoDetailResponseData `json:"response"` // <p>Detail informations you are querying.</p>
}
type GetVideoDetailResponseData struct {
	VideoUploadId string `json:"video_upload_id"` // [Required] <p>ID of uploaded video.</p>
	PostId string `json:"post_id"` // [Required] <p>The unique identifier for post Shopee Video. Only have value when the video status is 300 (POSTED).</p>
	PostTime int64 `json:"post_time"` // [Required] <p>The time when the video post to Shopee Video. Only have value when the video status is 300 (POSTED).</p>
	VideoUrl string `json:"video_url"` // [Required] <p>Video play url.</p>
	Status int64 `json:"status"` // [Required] <p>Video current status. Applicable values:</p><p>200: DRAFT<br />300: POSTED</p><p>400: DELETED</p><p>500: SCHEDULED<br />600: SCHEDULED_FAILED</p>
	CoverImageUrl string `json:"cover_image_url"` // [Required] <p>Cover image url of the Shopee Video.</p>
	Caption string `json:"caption"` // [Required] <p>Description of the Shopee Video.</p>
	Duration int64 `json:"duration"` // [Required] <p>Video duration time in millisecond.</p>
	Views int64 `json:"views"` // [Required] <p>View count of post Shopee Video. Only have value when the video status is 300 (POSTED).</p>
	Likes int64 `json:"likes"` // [Required] <p>Like count of post Shopee Video. Only have value when the video status is 300 (POSTED).</p>
	Comments int64 `json:"comments"` // [Required] <p>Comment count the post Shopee Video.&nbsp;Only have value when the video status is 300 (POSTED).</p>
	HasPerformance bool `json:"has_performance"` // [Required] <p>Whether there is video metric data.</p>
	ItemList []GetVideoDetailResponseDataItem `json:"item_list"` // [Required] <p>List of products linked with the Shopee Video.</p>
	AllowInfo *AllowInfo `json:"allow_info"` // [Required] <p>Whether allow stitch and duet.</p>
	ScheduledInfo *ScheduledInfo `json:"scheduled_info"` // [Required] <p>When&nbsp;scheduled_post is true, scheduled_post_time must not empty.</p><p>When&nbsp;scheduled_post is false, scheduled_post_time must empty.</p>
	UpdateTime int64 `json:"update_time"` // [Required] <p>The lasted update time the video.</p>
}
type GetVideoDetailResponseDataItem struct {
	ShopId int64 `json:"shop_id"` // [Required] <p>Shopee's unique identifier for a shop of the item.</p>
	ItemId int64 `json:"item_id"` // [Required] <p>Shopee's unique identifier for an item.</p>
	ItemName string `json:"item_name"` // [Required] <p>Name of the item.</p>
	CustomItemName string `json:"custom_item_name"` // [Required] <p>Name of the item displayed on Shopee Video (max 255 characters).</p>
	ItemCoverImageUrl string `json:"item_cover_image_url"` // [Required] <p>Cover image url of the item.</p>
	MinPrice float64 `json:"min_price"` // [Required] <p>Min price of the item.</p>
	MaxPrice float64 `json:"max_price"` // [Required] <p>Max price of the item.</p>
	Stock int64 `json:"stock"` // [Required] <p>Stock of the item.</p>
}
type GetVideoListRequest struct {
	ListType int64 `json:"list_type" url:"list_type"` // [Required] <p>Search tpye for video in draft status or video already post to Shopee Video.</p><p>1: draft</p><p>2: post</p>
	PageNo int64 `json:"page_no" url:"page_no"` // [Required] <p>The start index of request. Starting from 1.</p>
	PageSize int64 `json:"page_size" url:"page_size"` // [Required] <p>The number of affiliate returned by this request, Max is 20.</p>
}
type GetVideoListResponse struct {
	BaseResponse // Common response fields
	Response GetVideoListResponseData `json:"response"` // <p>Detail informations you are querying.</p>
}
type GetVideoListResponseData struct {
	TotalCount int64 `json:"total_count"` // [Required] <p>The total count of video that match the condition.</p>
	HasMore bool `json:"has_more"` // [Required] <p>This is to indicate whether the video list is more than one page. If this value is true, you may want to continue to check next page to retrieve the rest of data.</p>
	List *GetVideoListResponseDataList `json:"list"` // [Required] <p>The list of video that match the condition.</p>
}
type GetVideoListResponseDataList struct {
	VideoUploadId string `json:"video_upload_id"` // [Required] <p>ID of uploaded video.</p>
	PostId string `json:"post_id"` // [Required] <p>The unique identifier for post Shopee Video. Only have value when the video status is 300 (POSTED).</p>
	PostTime int64 `json:"post_time"` // [Required] <p>The time when the video post to Shopee Video. Only have value when the video status is 300 (POSTED).</p>
	VideoUrl string `json:"video_url"` // [Required] <p>Video play url.</p>
	Status int64 `json:"status"` // [Required] <p>Video current status. Applicable values:</p><p>200: DRAFT<br />300: POSTED</p><p>400: DELETED</p><p>500: SCHEDULED<br />600: SCHEDULED_FAILED</p>
	CoverImageUrl string `json:"cover_image_url"` // [Required] <p>Cover image url of the Shopee Video.</p>
	Caption string `json:"caption"` // [Required] <p>Description of the Shopee Video.</p>
	Duration int64 `json:"duration"` // [Required] <p>Video duration time in millisecond.</p>
	Views int64 `json:"views"` // [Required] <p>View count of post Shopee Video. Only have value when the video status is 300 (POSTED).</p>
	Likes int64 `json:"likes"` // [Required] <p>Like count the post Shopee Video.&nbsp;Only have value when the video status is 300 (POSTED).</p>
	Comments int64 `json:"comments"` // [Required] <p>Comment count the post Shopee Video.&nbsp;Only have value when the video status is 300 (POSTED).</p>
	HasPerformance bool `json:"has_performance"` // [Required] <p>Whether there is video metric data.</p>
	ItemList *GetVideoDetailResponseDataItem `json:"item_list"` // [Required] <p>List of products linked with the Shopee Video.</p>
	AllowInfo *AllowInfo `json:"allow_info"` // [Required] <p>Whether allow stitch and duet.</p>
	ScheduledInfo *ScheduledInfo `json:"scheduled_info"` // [Required] <p>When&nbsp;scheduled_post is true, scheduled_post_time must not empty.</p><p>When&nbsp;scheduled_post is false, scheduled_post_time must empty.</p>
	UpdateTime int64 `json:"update_time"` // [Required] <p>The lasted update time the video.</p>
}
type GetVideoPerformanceListRequest struct {
	Caption *string `json:"caption,omitempty" url:"caption,omitempty"` // [Optional] <p>Description of the Shopee Video.</p>
	EndDate string `json:"end_date" url:"end_date"` // [Required] <p>The end_date format should be "YYYY-MM-DD".</p><p>- For Day, Last7d, Last15d, and Last30d, the end_date must before current day.</p><p>- For Week, the end_date must be Sunday and must be less than or equal to the current week.</p><p>- For Month, the end_date must be the end of the month and must be less than or equal to the current month.</p>
	OrderBy string `json:"order_by" url:"order_by"` // [Required] <p>Use this field to specify which field to use to sort the returned list. Available values:</p><p>Views</p><p>Likes</p><p>Comments</p><p>AvgViewsDuration</p>
	PageNo int64 `json:"page_no" url:"page_no"` // [Required] <p>The start index of request. Starting from 1.</p>
	PageSize int64 `json:"page_size" url:"page_size"` // [Required] <p>The number of video returned by this request. Max is 20.</p>
	PeriodType string `json:"period_type" url:"period_type"` // [Required] <p>Period Type. Applicable values:</p><p>Day<br />Week<br />Month<br /></p><p>Last7d</p><p>Last15d</p><p>Last30d</p><p><br /></p><p>Note: The end date must align with the Period Type.</p>
	Sort string `json:"sort" url:"sort"` // [Required] <p>Use this field to specify whether the returned list is sorted in ascending or descending order_by.&nbsp;Available values:</p><p>asc</p><p>desc</p>
}
type GetVideoPerformanceListResponse struct {
	BaseResponse // Common response fields
	Response GetVideoPerformanceListResponseData `json:"response"` // <p>Detail informations you are querying.</p>
}
type GetVideoPerformanceListResponseData struct {
	TotalCount int64 `json:"total_count"` // [Required] <p>The total count of video that match the condition.</p>
	HasMore bool `json:"has_more"` // [Required] <p>This is to indicate whether the video list is more than one page. If this value is true, you may want to continue to check next page to retrieve the rest of data.</p>
	List []GetVideoPerformanceListResponseDataList `json:"list"` // [Required] <p>The list of video that match the condition.</p>
}
type GetVideoPerformanceListResponseDataList struct {
	VideoUploadId string `json:"video_upload_id"` // [Required] <p>ID of uploaded video.</p>
	PostId string `json:"post_id"` // [Required] <p>The unique identifier for post Shopee Video.&nbsp;Only have value when the video status is 300 (POSTED).</p>
	PostTime int64 `json:"post_time"` // [Required] <p>The time when the video post to Shopee Video.&nbsp;Only have value when the video status is 300 (POSTED).</p>
	VideoUrl string `json:"video_url"` // [Required] <p>Video play url.</p>
	Status int64 `json:"status"` // [Required] <p>Video current status. Applicable values:</p><p>300: POSTED</p><p>400: DELETED</p>
	CoverImageUrl string `json:"cover_image_url"` // [Required] <p>Cover image url of the Shopee Video.</p>
	Caption string `json:"caption"` // [Required] <p>Description of the Shopee Video.</p>
	Duration string `json:"duration"` // [Required] <p>Video duration time in millisecond.</p>
	Views int64 `json:"views"` // [Required] <p>View count of post Shopee Video.&nbsp;Only have value when the video status is 300 (POSTED).</p>
	Likes int64 `json:"likes"` // [Required] <p>Like count the post Shopee Video.&nbsp;Only have value when the video status is 300 (POSTED).</p>
	Comments int64 `json:"comments"` // [Required] <p>Comment count the post Shopee Video.&nbsp;Only have value when the video status is 300 (POSTED).</p>
	Shares int64 `json:"shares"` // [Required] <p>Share count the post Shopee Video.&nbsp;Only have value when the video status is 300 (POSTED).</p>
	AvgViewsDuration int64 `json:"avg_views_duration"` // [Required] <p>Total watch duration per video.</p>
	CompletionRate float64 `json:"completion_rate"` // [Required] <p>Video completion rate.</p>
	PlacedOrders int64 `json:"placed_orders"` // [Required] <p>The number of placed orders for the video.</p>
	ConfirmedOrders int64 `json:"confirmed_orders"` // [Required] <p>The number of confirmed orders for the video.</p>
	PlacedSales float64 `json:"placed_sales"` // [Required] <p>The placed value of orders for the video.</p>
	ConfirmedSales float64 `json:"confirmed_sales"` // [Required] <p>The confirmed value of orders for the video.</p>
	PlacedItemSold int64 `json:"placed_item_sold"` // [Required] <p>Number of item sold from placed orders in the video.</p>
	ConfirmedItemSold int64 `json:"confirmed_item_sold"` // [Required] <p>Number of item sold from confirmed orders in the video.</p>
	FetchedDateRange string `json:"fetched_date_range"` // [Required] <p>Data Date Range.</p>
}
type KeyMetric struct {
	PlacedSales float64 `json:"placed_sales"` // [Required] <p>The placed value of orders from all videos in the period selected.</p>
	ConfirmedSales float64 `json:"confirmed_sales"` // [Required] <p>The confirmed value of orders from all videos in the period selected.</p>
	PlacedOrders int64 `json:"placed_orders"` // [Required] <p>The number of placed orders from all videos in the period selected.&nbsp;</p>
	ConfirmedOrders int64 `json:"confirmed_orders"` // [Required] <p>The number of confirmed orders from all videos in the period selected.</p>
	PlacedItemSold int64 `json:"placed_item_sold"` // [Required] <p>Number of item sold from placed orders in the video.</p>
	ConfirmedItemSold int64 `json:"confirmed_item_sold"` // [Required] <p>Number of item sold from confirmed orders in the video.</p>
	TotalViewers int64 `json:"total_viewers"` // [Required] <p>Number of viewers of the video.</p>
	EffectiveViews int64 `json:"effective_views"` // [Required] <p>Number of views for the video that lasted for more than 3 seconds.</p>
	AvgViewDuration int64 `json:"avg_view_duration"` // [Required] <p>Total watch duration per video.</p>
}
type PostVideoRequest struct {
	VideoUploadIdList []string `json:"video_upload_id_list"` // [Required] <p>ID of uploaded video. Obtain from v2.media.get_video_upload_result. No more than 5.</p>
}
type PostVideoResponse struct {
	BaseResponse // Common response fields
	Response PostVideoResponseData `json:"response"` // <p>Detail informations you are querying.</p>
}
type PostVideoResponseData struct {
	SuccessList []PostVideoResponseDataSuccess `json:"success_list"` // [Required] <p>The list of video post successfully.</p>
	FailureList []EditVideoInfoResponseDataFailure `json:"failure_list"` // [Required] <p>The list of video post failed.</p>
}
type PostVideoResponseDataSuccess struct {
	SuccessVideoUploadId string `json:"success_video_upload_id"` // [Required] <p>The video_upload_id post successfully.</p>
	PostId string `json:"post_id"` // [Required] <p>The unique identifier for post Shopee Video.</p>
}
type ScheduledInfo struct {
	ScheduledPost bool `json:"scheduled_post"` // [Required] <p>Whether post it to Shopee Video at scheduled time.</p>
	ScheduledPostTime int64 `json:"scheduled_post_time"` // [Required] <p>Scheduled post time, millisecond timestamp.</p>
}
type VideoPerformance struct {
	Views int64 `json:"views"` // [Required] <p>Amount of views from the video.</p>
	Likes int64 `json:"likes"` // [Required] <p>Total likes from the video.</p>
	Comments int64 `json:"comments"` // [Required] <p>Total comments from the video.</p>
	Shares int64 `json:"shares"` // [Required] <p>Total shares from the video.</p>
	FollowersGrowth int64 `json:"followers_growth"` // [Required] <p>Amount of new followers from the Video.</p>
	PlacedOrders int64 `json:"placed_orders"` // [Required] <p>Amount of product orders from the video.</p>
	PlacedSales float64 `json:"placed_sales"` // [Required] <p>Amount of product sales from the video.</p>
	UniqueBuyers int64 `json:"unique_buyers"` // [Required] <p>Buyers of the products in the video.</p>
	ConversionRate float64 `json:"conversion_rate"` // [Required] <p>Amount of products sold from the video/amount of views from the video.</p>
	SoldItems int64 `json:"sold_items"` // [Required] <p>Amount of products sold from the video.</p>
	ProductClicks int64 `json:"product_clicks"` // [Required] <p>The product click value of orders for item.</p>
	ProductClickRate float64 `json:"product_click_rate"` // [Required] <p>The product click rate value of orders for item.</p>
	SalesPerOrder float64 `json:"sales_per_order"` // [Required] <p>Amount of product sales from the video/amount of product orders from the video.</p>
	SalesPerBuyer float64 `json:"sales_per_buyer"` // [Required] <p>Amount of product sales from the video/amount of buyers from the video.</p>
}
type VideoTotalMetric struct {
	PlacedSales float64 `json:"placed_sales"` // [Required] <p>The placed value of orders from all videos in the period selected.</p>
	ConfirmedSales float64 `json:"confirmed_sales"` // [Required] <p>The confirmed value of orders from all videos in the period selected.</p>
	PlacedOrders int64 `json:"placed_orders"` // [Required] <p>The number of placed orders from all videos in the period selected.</p>
	ConfirmedOrders int64 `json:"confirmed_orders"` // [Required] <p>The number of confirmed orders from all videos in the period selected.</p>
	PlacedItemSold int64 `json:"placed_item_sold"` // [Required] <p>Number of item sold from placed orders in the video.</p>
	ConfirmedItemSold int64 `json:"confirmed_item_sold"` // [Required] <p>Number of item sold from confirmed orders in the video.</p>
	TotalViewers int64 `json:"total_viewers"` // [Required] <p>Number of viewers in the video.</p>
	EffectiveViews int64 `json:"effective_views"` // [Required] <p>Number of views from the video that lasted for more than 3 seconds.</p>
	AvgViewDuration int64 `json:"avg_view_duration"` // [Required] <p>Total watch duration per video.</p>
	PlacedBuyers int64 `json:"placed_buyers"` // [Required] <p>Number of unique buyers who placed order from the video.</p>
	ConfirmedBuyers int64 `json:"confirmed_buyers"` // [Required] <p>Number of unique buyers who confirmed order from the video.</p>
	TotalAtc int64 `json:"total_atc"` // [Required] <p>Number of "Add To Cart" button clicked for all products in the orange bag during video viewing.</p>
	Ctr float64 `json:"ctr"` // [Required] <p>Number of products clicks divided by Number of video views.</p>
	PlacedCoRate float64 `json:"placed_co_rate"` // [Required] <p>Number of placed product orders from the video divided by Number of product clicks from the video.</p>
	ConfirmedCoRate float64 `json:"confirmed_co_rate"` // [Required] <p>Number of confirmed product orders from the video divided by Number of product clicks from the video.</p>
	PlacedAbs float64 `json:"placed_abs"` // [Required] <p>Total placed sales divided by Total placed orders.</p>
	ConfirmedAbs float64 `json:"confirmed_abs"` // [Required] <p>Total confirmed sales divided by Total confirmed orders.</p>
	PlacedGpm float64 `json:"placed_gpm"` // [Required] <p>The placed Sales generated for every 1,000 views.</p>
	ConfirmedGpm float64 `json:"confirmed_gpm"` // [Required] <p>The confirmed Sales generated for every 1,000 views.</p>
	VideoWithProducts int64 `json:"video_with_products"` // [Required] <p>Videos with at least one product in the orange bag</p>
	PlacedRevenueGeneratingVideos int64 `json:"placed_revenue_generating_videos"` // [Required] <p>Videos that generates placed revenues.</p>
	ConfirmedRevenueGeneratingVideos int64 `json:"confirmed_revenue_generating_videos"` // [Required] <p>Videos that generates confirmed revenues.</p>
	TotalViews int64 `json:"total_views"` // [Required] <p>Number of views from all videos.</p>
	TotalLikes int64 `json:"total_likes"` // [Required] <p>Number of likes from all videos.</p>
	TotalShares int64 `json:"total_shares"` // [Required] <p>Number of shares from all videos.</p>
	TotalComments int64 `json:"total_comments"` // [Required] <p>Number of comments from all videos.</p>
	VideoNewFollowers int64 `json:"video_new_followers"` // [Required] <p>Number of new followers from all videos.</p>
	DataPeriod string `json:"data_period"` // [Required] <p>Data offline computation time.</p>
}
type VideoUpload struct {
	VideoUploadId string `json:"video_upload_id"` // [Required] <p>ID of uploaded video. Obtain from v2.media.get_video_upload_result.</p>
	Caption *string `json:"caption,omitempty"` // [Optional] <p>Description of the Shopee Video.</p>
	CoverImageUrl string `json:"cover_image_url"` // [Required] <p>Selected cover image url of the Shopee Video. Obtain from v2.video.get_cover_list.</p>
	ItemInfo []VideoUploadItemInfo `json:"item_info,omitempty"` // [Optional] <p>List of products to be linked with the Shopee Video, no more than 6.</p>
	AllowInfo *AllowInfo `json:"allow_info"` // [Required] <p>Whether allow stitch and duet.</p>
	ScheduledInfo *ScheduledInfo `json:"scheduled_info"` // [Required] <p>When&nbsp;scheduled_post is true, scheduled_post_time must not empty.</p><p>When&nbsp;scheduled_post is false, scheduled_post_time must empty.</p>
}
type VideoUploadItemInfo struct {
	ItemId int64 `json:"item_id"` // [Required] <p>Shopee's unique identifier for an item.</p>
	CustomItemName *string `json:"custom_item_name,omitempty"` // [Optional] <p>Product display name in Shopee Video.</p>
}
