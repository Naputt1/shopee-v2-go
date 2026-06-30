package goshopee

type AddItemListRequest struct {
	ItemList []TopSellingProducts `json:"item_list"` // [Required] <p>The list of item to add.</p>
	SessionId int64 `json:"session_id"` // [Required] <p>The identifier of livestream session.</p>
}
type AddItemListResponse struct {
	BaseResponse // Common response fields
	Response interface{} `json:"response"` // Actual response data
}
type AffiliateInfo struct {
	CommissionRate float64 `json:"commission_rate"` // [Required] <p>The commission rate that the streamer can get, for example, 0.1 means 10%.</p>
	IsCampaign bool `json:"is_campaign"` // [Required] <p>Whether participate in a campaign project (generally, the commission will be higher)</p>
	CampaignMcnName string `json:"campaign_mcn_name"` // [Required] <p>MCN agency that initiated this campaign</p>
	CampaignStartTime int64 `json:"campaign_start_time"` // [Required] <p>Campaign start time, it's unix timestamp in seconds.</p>
	CampaignEndTime int64 `json:"campaign_end_time"` // [Required] <p>Campaign end time, it's unix timestamp in seconds.</p>
}
type ApplyItemSetRequest struct {
	ItemSetIds []int64 `json:"item_set_ids"` // [Required] <p>List of item set id to apply.</p>
	SessionId int64 `json:"session_id"` // [Required] <p>The identifier of livestream session.</p>
}
type ApplyItemSetResponse struct {
	BaseResponse // Common response fields
	Response interface{} `json:"response"` // Actual response data
}
type BanUserCommentRequest struct {
	BanUserId int64 `json:"ban_user_id"` // [Required] <p>The user id that will be banned from posting comments.</p>
	SessionId int64 `json:"session_id"` // [Required] <p>The identifier of livestream session.</p>
}
type BanUserCommentResponse struct {
	BaseResponse // Common response fields
	Response interface{} `json:"response"` // Actual response data
}
type CreateSessionRequest struct {
	CoverImageUrl string `json:"cover_image_url"` // [Required] <p>The cover image URL of livestream session.</p><p>Please call the v2.livestream.upload_image to upload the cover image file and get the cover_image_url.</p>
	Description *string `json:"description,omitempty"` // [Optional] <p>The description of livestream session, cannot exceed 200 characters.</p>
	IsTest *bool `json:"is_test,omitempty"` // [Optional] <p>Indicate whether the livestream session is for testing purpose only.</p>
	Title string `json:"title"` // [Required] <p>The title of livestream session, cannot exceed 200 characters.</p>
}
type CreateSessionResponse struct {
	BaseResponse // Common response fields
	Response CreateSessionResponseData `json:"response"` // <p>Detail informations you are querying.</p>
}
type CreateSessionResponseData struct {
	SessionId int64 `json:"session_id"` // [Required] <p>The identifier of livestream session.</p>
}
type DeleteItemListRequest struct {
	ItemList []TopSellingProducts `json:"item_list"` // [Required] <p>The list of item to delete.</p>
	SessionId int64 `json:"session_id"` // [Required] <p>The identifier of livestream session.</p>
}
type DeleteItemListResponse struct {
	BaseResponse // Common response fields
	Response interface{} `json:"response"` // Actual response data
}
type DeleteShowItemRequest struct {
	SessionId int64 `json:"session_id"` // [Required] <p>The identifier of livestream session.</p>
}
type DeleteShowItemResponse struct {
	BaseResponse // Common response fields
	Response interface{} `json:"response"` // Actual response data
}
type EndSessionRequest struct {
	SessionId int64 `json:"session_id"` // [Required] <p>The identifier of livestream session.</p>
}
type EndSessionResponse struct {
	BaseResponse // Common response fields
	Response interface{} `json:"response"` // Actual response data
}
type GetItemCountRequest struct {
	SessionId int64 `json:"session_id" url:"session_id"` // [Required] <p>The identifier of livestream session.</p>
}
type GetItemCountResponse struct {
	BaseResponse // Common response fields
	Response GetItemCountResponseData `json:"response"` // <p>Detail informations you are querying.</p>
}
type GetItemCountResponseData struct {
	ItemCount int64 `json:"item_count"` // [Required] <p>The number of items in the shopping bag of this session.</p>
	MaxItemCount int64 `json:"max_item_count"` // [Required] <p>The maximum number of items allowed in the shopping bag of this session.</p>
}
type GetItemListRequest struct {
	Offset int64 `json:"offset" url:"offset"` // [Required] <p>Specifies the starting entry of data to return in the current call. Default is 0, if data is more than one page, the offset can be some entry to start next call.</p>
	PageSize int64 `json:"page_size" url:"page_size"` // [Required] <p>Each result set is returned as a page of entries. Use the "page_size" filters to control the maximum number of entries to retrieve per page (i.e., per call). This integer value is used to specify the maximum number of entries to return in a single "page" of data. The limit of page_size if between 1 and 100.</p>
	SessionId int64 `json:"session_id" url:"session_id"` // [Required] <p>The identifier of livestream session.</p>
}
type GetItemListResponse struct {
	BaseResponse // Common response fields
	Response GetItemListResponseData `json:"response"` // <p>Detail informations you are querying.</p>
}
type GetItemListResponseData struct {
	More bool `json:"more"` // [Required] <p>This is to indicate whether the list is more than one page. If this value is true, you may want to continue to check next page to retrieve the rest of data.</p>
	NextOffset int64 `json:"next_offset"` // [Required] <p>If more is true, this value need set to next request offset.</p>
	List []GetItemListResponseDataList `json:"list"` // [Required] 
}
type GetItemListResponseDataList struct {
	ItemNo int64 `json:"item_no"` // [Required] <p>The order of this item in the shopping bag of current session, start from 1.</p>
	ItemId int64 `json:"item_id"` // [Required] <p>Shopee's unique identifier for an item.</p>
	ShopId int64 `json:"shop_id"` // [Required] <p>The shop id of this item.</p>
	Name string `json:"name"` // [Required] <p>Name of the item in local language.<br /></p>
	ImageUrl string `json:"image_url"` // [Required] <p>The image url of this item.</p>
	PriceInfo *ListPriceInfo `json:"price_info"` // [Required] 
	AffiliateInfo *AffiliateInfo `json:"affiliate_info"` // [Required] 
}
type GetItemSetItemListRequest struct {
	ItemSetId int64 `json:"item_set_id" url:"item_set_id"` // [Required] <p>The identifier of the item set.</p>
	Offset int64 `json:"offset" url:"offset"` // [Required] <p>Specifies the starting entry of data to return in the current call. Default is 0, if data is more than one page, the offset can be some entry to start next call.</p>
	PageSize int64 `json:"page_size" url:"page_size"` // [Required] <p>Each result set is returned as a page of entries. Use the "page_size" filters to control the maximum number of entries to retrieve per page (i.e., per call). This integer value is used to specify the maximum number of entries to return in a single "page" of data. The limit of page_size if between 1 and 100.</p>
}
type GetItemSetItemListResponse struct {
	BaseResponse // Common response fields
	Response GetItemSetItemListResponseData `json:"response"` // <p>Detail informations you are querying.</p>
}
type GetItemSetItemListResponseData struct {
	More bool `json:"more"` // [Required] <p>This is to indicate whether the list is more than one page. If this value is true, you may want to continue to check next page to retrieve the rest of data.</p>
	NextOffset int64 `json:"next_offset"` // [Required] <p>If more is true, this value need set to next request offset.</p>
	List []GetItemSetItemListResponseDataList `json:"list"` // [Required] 
}
type GetItemSetItemListResponseDataList struct {
	ItemId int64 `json:"item_id"` // [Required] <p>Shopee's unique identifier for an item.</p>
	ShopId int64 `json:"shop_id"` // [Required] <p>The shop id of this item.</p>
	Name string `json:"name"` // [Required] <p>Name of the item in local language.<br /></p>
	ImageUrl string `json:"image_url"` // [Required] <p>The image url of this item.</p>
	PriceInfo *ListPriceInfo `json:"price_info"` // [Required] 
	AffiliateInfo *AffiliateInfo `json:"affiliate_info"` // [Required] 
}
type GetItemSetListRequest struct {
	Keyword *string `json:"keyword,omitempty" url:"keyword,omitempty"` // [Optional] <p>Search the item set with it's name matching the keyword.</p>
	Offset int64 `json:"offset" url:"offset"` // [Required] <p>Specifies the starting entry of data to return in the current call. Default is 0, if data is more than one page, the offset can be some entry to start next call.</p>
	PageSize int64 `json:"page_size" url:"page_size"` // [Required] <p>Each result set is returned as a page of entries. Use the "page_size" filters to control the maximum number of entries to retrieve per page (i.e., per call). This integer value is used to specify the maximum number of entries to return in a single "page" of data. The limit of page_size if between 1 and 100.</p>
}
type GetItemSetListResponse struct {
	BaseResponse // Common response fields
	Response GetItemSetListResponseData `json:"response"` // <p>Detail informations you are querying.</p>
}
type GetItemSetListResponseData struct {
	More bool `json:"more"` // [Required] <p>This is to indicate whether the list is more than one page. If this value is true, you may want to continue to check next page to retrieve the rest of data.</p>
	NextOffset int64 `json:"next_offset"` // [Required] <p>If more is true, this value need set to next request offset.</p>
	List []GetItemSetListResponseDataList `json:"list"` // [Required] 
}
type GetItemSetListResponseDataList struct {
	ItemSetId int64 `json:"item_set_id"` // [Required] <p>The identifier of the item set.</p>
	ItemSetName string `json:"item_set_name"` // [Required] <p>The name of the item set.</p>
	ItemCount int64 `json:"item_count"` // [Required] <p>The number of items in this item set.</p>
}
type GetLatestCommentListRequest struct {
	Offset *int64 `json:"offset,omitempty" url:"offset,omitempty"` // [Optional] <p>Specifies the starting entry of data to return in the current call. Default is 0, if data is more than one page, the offset can be some entry to start next call.</p>
	SessionId int64 `json:"session_id" url:"session_id"` // [Required] <p>The identifier of livestream session.</p>
}
type GetLatestCommentListResponse struct {
	BaseResponse // Common response fields
	Response GetLatestCommentListResponseData `json:"response"` // <p>Detail informations you are querying.</p>
}
type GetLatestCommentListResponseData struct {
	NextOffset int64 `json:"next_offset"` // [Required] <p>The offset for next page request.</p>
	List []GetLatestCommentListResponseDataList `json:"list"` // [Required] 
}
type GetLatestCommentListResponseDataList struct {
	CommentId int64 `json:"comment_id"` // [Required] <p>The identifier of comment.</p>
	Content string `json:"content"` // [Required] <p>The content of comment.</p>
	Timestamp int64 `json:"timestamp"` // [Required] <p>Timestamp for posting comment.&nbsp;It's unix timestamp in seconds.</p>
	UserId int64 `json:"user_id"` // [Required] <p>The user id of the one who posted the comment.</p>
	Username string `json:"username"` // [Required] <p>The username of the one who posted comment.</p>
}
type GetLikeItemListRequest struct {
	Keyword *string `json:"keyword,omitempty" url:"keyword,omitempty"` // [Optional] <p>Search items with name matching this keyword.</p>
	Offset int64 `json:"offset" url:"offset"` // [Required] <p>Specifies the starting entry of data to return in the current call. Default is 0, if data is more than one page, the offset can be some entry to start next call.</p>
	PageSize int64 `json:"page_size" url:"page_size"` // [Required] <p>Each result set is returned as a page of entries. Use the "page_size" filters to control the maximum number of entries to retrieve per page (i.e., per call). This integer value is used to specify the maximum number of entries to return in a single "page" of data. The limit of page_size if between 1 and 100.</p>
}
type GetLikeItemListResponse struct {
	BaseResponse // Common response fields
	Response GetLikeItemListResponseData `json:"response"` // <p>Detail informations you are querying.</p>
}
type GetLikeItemListResponseData struct {
	More bool `json:"more"` // [Required] <p>This is to indicate whether the list is more than one page. If this value is true, you may want to continue to check next page to retrieve the rest of data.</p>
	NextOffset int64 `json:"next_offset"` // [Required] <p>If more is true, this value need set to next request offset.</p>
	List []GetItemSetItemListResponseDataList `json:"list"` // [Required] 
}
type GetRecentItemListRequest struct {
	Offset int64 `json:"offset" url:"offset"` // [Required] <p>Specifies the starting entry of data to return in the current call. Default is 0, if data is more than one page, the offset can be some entry to start next call.</p>
	PageSize int64 `json:"page_size" url:"page_size"` // [Required] <p>Each result set is returned as a page of entries. Use the "page_size" filters to control the maximum number of entries to retrieve per page (i.e., per call). This integer value is used to specify the maximum number of entries to return in a single "page" of data. The limit of page_size if between 1 and 100.</p>
}
type GetRecentItemListResponse struct {
	BaseResponse // Common response fields
	Response GetRecentItemListResponseData `json:"response"` // <p>Detail informations you are querying.</p>
}
type GetRecentItemListResponseData struct {
	More bool `json:"more"` // [Required] <p>This is to indicate whether the list is more than one page. If this value is true, you may want to continue to check next page to retrieve the rest of data.</p>
	NextOffset int64 `json:"next_offset"` // [Required] <p>If more is true, this value need set to next request offset.</p>
	List []GetItemSetItemListResponseDataList `json:"list"` // [Required] 
}
type GetSessionDetailRequest struct {
	SessionId int64 `json:"session_id" url:"session_id"` // [Required] <p>The identifier of livestream session.</p>
}
type GetSessionDetailResponse struct {
	BaseResponse // Common response fields
	Response GetSessionDetailResponseData `json:"response"` // <p>Detail informations you are querying.</p>
}
type GetSessionDetailResponseData struct {
	SessionId int64 `json:"session_id"` // [Required] <p>The identifier of livestream session.</p>
	Title string `json:"title"` // [Required] <p>The title of the livestream session.</p>
	Description string `json:"description"` // [Required] <p>The description of the livestream session.</p>
	CoverImageUrl string `json:"cover_image_url"` // [Required] <p>The cover image URL of the livestream session.</p>
	Status int64 `json:"status"` // [Required] <p>The status of the livestream session, the enumeration values are as follows:</p><p>0 - Initial</p><p>1 - Ongoing</p><p>2 - Ended</p>
	ShareUrl string `json:"share_url"` // [Required] <p>The share link of the livestream session.</p>
	IsTest bool `json:"is_test"` // [Required] <p>Indicate whether this livestream session if for testing purpose only.</p>
	CreateTime int64 `json:"create_time"` // [Required] <p>The creation time of the livestream session. It's unix timestamp in seconds.</p>
	UpdateTime int64 `json:"update_time"` // [Required] <p>The update time of the livestream session. It's unix timestamp in seconds.</p>
	StartTime int64 `json:"start_time"` // [Required] <p>The start time of the livestream session, 0 if session is not started yet.&nbsp;It's unix timestamp in seconds.</p>
	EndTime int64 `json:"end_time"` // [Required] <p>The end time of livestream session, 0 if session is not ended yet.&nbsp;It's unix timestamp in seconds.</p>
	StreamUrlList *StreamUrl `json:"stream_url_list"` // [Required] 
}
type GetSessionItemMetricRequest struct {
	Offset int64 `json:"offset" url:"offset"` // [Required] <p>Specifies the starting entry of data to return in the current call. Default is 0, if data is more than one page, the offset can be some entry to start next call.</p>
	PageSize int64 `json:"page_size" url:"page_size"` // [Required] <p>Each result set is returned as a page of entries. Use the "page_size" filters to control the maximum number of entries to retrieve per page (i.e., per call). This integer value is used to specify the maximum number of entries to return in a single "page" of data. The limit of page_size if between 1 and 100.</p>
	SessionId int64 `json:"session_id" url:"session_id"` // [Required] <p>The identifier of livestream session.</p>
}
type GetSessionItemMetricResponse struct {
	BaseResponse // Common response fields
	Response GetSessionItemMetricResponseData `json:"response"` // <p>Detail informations you are querying.</p>
}
type GetSessionItemMetricResponseData struct {
	More bool `json:"more"` // [Required] <p>This is to indicate whether the list is more than one page. If this value is true, you may want to continue to check next page to retrieve the rest of data.</p>
	NextOffset int64 `json:"next_offset"` // [Required] <p>If more is true, this value need set to next request offset.</p>
	List []GetSessionItemMetricResponseDataList `json:"list"` // [Required] 
}
type GetSessionItemMetricResponseDataList struct {
	Item *ListItem `json:"item"` // [Required] 
	Metric *ListMetric `json:"metric"` // [Required] 
}
type GetSessionMetricRequest struct {
	SessionId int64 `json:"session_id" url:"session_id"` // [Required] <p>The identifier of livestream session.</p>
}
type GetSessionMetricResponse struct {
	BaseResponse // Common response fields
	Response GetSessionMetricResponseData `json:"response"` // <p>Detail informations you are querying.</p>
}
type GetSessionMetricResponseData struct {
	Gmv float64 `json:"gmv"` // [Required] <p>Value of placed orders (paid and unpaid) during Livestream, including sales from cancelled orders.<br /></p>
	Atc int64 `json:"atc"` // [Required] <p>Number of "Add To Cart" button clicked for all products in the orange bag during livestream.<br /></p>
	Ctr float64 `json:"ctr"` // [Required] <p>Number of products clicks divided by Number of Livestream views.<br /></p>
	Co float64 `json:"co"` // [Required] <p>Amount of product orders from the stream divided by Amount of product clicks from the stream.<br /></p>
	Orders int64 `json:"orders"` // [Required] <p>Number of placed orders (paid and unpaid) during Livestream, including cancelled orders.<br /></p>
	Ccu int64 `json:"ccu"` // [Required] <p>Number of viewers during stream.</p>
	EngageCcu1m int64 `json:"engage_ccu_1m"` // [Required] <p>Number of Concurrent viewers in the stream that have watched for more than 1 minute.<br /></p>
	PeakCcu int64 `json:"peak_ccu"` // [Required] <p>Highest number of viewers during stream.<br /></p>
	Likes int64 `json:"likes"` // [Required] <p>Number of "Like" clicked during livestream.<br /></p>
	Comments int64 `json:"comments"` // [Required] <p>Number of comments acquired during the stream.<br /></p>
	Shares int64 `json:"shares"` // [Required] <p>Number of shares created during the stream.<br /></p>
	Views int64 `json:"views"` // [Required] <p>Number of views from the stream.<br /></p>
	AvgViewingDuration int64 `json:"avg_viewing_duration"` // [Required] <p>Average of Viewer duration watching in the stream.<br /></p>
}
type GetShowItemRequest struct {
	SessionId int64 `json:"session_id" url:"session_id"` // [Required] <p>The identifier of livestream session.</p>
}
type GetShowItemResponse struct {
	BaseResponse // Common response fields
	Response GetShowItemResponseData `json:"response"` // <p>Detail informations you are querying.</p>
}
type GetShowItemResponseData struct {
	HasShowItem bool `json:"has_show_item"` // [Required] <p>Whether has the showing item.</p>
	Item *GetShowItemResponseDataItem `json:"item"` // [Required] 
}
type GetShowItemResponseDataItem struct {
	ItemNo int64 `json:"item_no"` // [Required] <p>The order of this item in the shopping bag of current session, start from 1. Only return item_no when showing item is in the shopping bag of current session.</p>
	ItemId int64 `json:"item_id"` // [Required] <p>Shopee's unique identifier for an item.</p>
	ShopId int64 `json:"shop_id"` // [Required] <p>The shop id of this item.</p>
	Name string `json:"name"` // [Required] <p>Name of the item in local language.<br /></p>
	ImageUrl string `json:"image_url"` // [Required] <p>The image url of this item.</p>
	PriceInfo *ListPriceInfo `json:"price_info"` // [Required] 
}
type ListItem struct {
	ItemId int64 `json:"item_id"` // [Required] <p>Shopee's unique identifier for an item.<br /></p>
	ShopId int64 `json:"shop_id"` // [Required] <p>The shop id of the item.</p>
	Name string `json:"name"` // [Required] <p>Name of the item in local language.<br /></p>
	ImageUrl string `json:"image_url"` // [Required] <p>The image url of the item.</p>
	PriceInfo *ListPriceInfo `json:"price_info"` // [Required] 
}
type ListMetric struct {
	ItemClicks int64 `json:"item_clicks"` // [Required] <p>Number of product clicks.</p>
	Atc int64 `json:"atc"` // [Required] <p>Number of "Add To Cart" button clicked for all products in the orange bag during livestream.<br /></p>
	SoldItems int64 `json:"sold_items"` // [Required] <p>Number of product sold.</p>
}
type ListPriceInfo struct {
	Currency string `json:"currency"` // [Required] <p>The three-digit code representing the currency unit used for the item.</p>
	CurrentPrice float64 `json:"current_price"` // [Required] <p>The current price of the item in the listing currency. If product under an ongoing promotion, current_price will be the promotion price.</p>
	OriginalPrice float64 `json:"original_price"` // [Required] <p>The original price of the item in the listing currency.</p>
}
type PostCommentRequest struct {
	Content string `json:"content"` // [Required] <p>The comment content, cannot exceed 150 characters.</p>
	SessionId int64 `json:"session_id"` // [Required] <p>The identifier of livestream session.</p>
}
type PostCommentResponse struct {
	BaseResponse // Common response fields
	Response PostCommentResponseData `json:"response"` // <p>Detail informations you are querying.</p>
}
type PostCommentResponseData struct {
	CommentId int64 `json:"comment_id"` // [Required] <p>The identifier of the comment.</p>
}
type StartSessionRequest struct {
	DomainId int64 `json:"domain_id"` // [Required] <p>The identifier of the stream domain.</p>
	SessionId int64 `json:"session_id"` // [Required] <p>The identifier of livestream session.</p>
}
type StartSessionResponse struct {
	BaseResponse // Common response fields
	Response interface{} `json:"response"` // Actual response data
}
type StreamUrl struct {
	PushUrl string `json:"push_url"` // [Required] <p>The push stream url for the livestream session.</p>
	PushKey string `json:"push_key"` // [Required] <p>The push stream key for the livestream session.</p>
	PlayUrl string `json:"play_url"` // [Required] <p>The pull stream url of the livestream session.</p>
	DomainId int64 `json:"domain_id"` // [Required] <p>The identifier of the stream domain, need to be passed in request for v2.livestream.start_session.</p>
}
type TopSellingProducts struct {
	ItemId int64 `json:"item_id"` // [Required] <p>Shopee's unique identifier for an item.</p>
}
type UnbanUserCommentRequest struct {
	SessionId int64 `json:"session_id"` // [Required] <p>The identifier of livestream session.</p>
	UnbanUserId int64 `json:"unban_user_id"` // [Required] <p>The user ID that will be unbanned from posting comments.</p>
}
type UnbanUserCommentResponse struct {
	BaseResponse // Common response fields
	Response interface{} `json:"response"` // Actual response data
}
type UpdateItemListRequest struct {
	ItemList []TopSellingProducts `json:"item_list"` // [Required] <p>The list of item with updated order.</p>
	SessionId int64 `json:"session_id"` // [Required] <p>The identifier of livestream session.</p>
}
type UpdateItemListResponse struct {
	BaseResponse // Common response fields
	Response interface{} `json:"response"` // Actual response data
}
type UpdateSessionRequest struct {
	CoverImageUrl string `json:"cover_image_url"` // [Required] <p>The cover image url of the livestream session.</p><p>Please call the v2.livestream.upload_image to upload the cover image file and get the cover_image_url.</p>
	Description *string `json:"description,omitempty"` // [Optional] <p>The description of the livestream session, cannot exceed 200 characters.</p>
	IsTest bool `json:"is_test"` // [Required] <p>Indicate whether this livestream session if for testing purpose only.</p>
	SessionId int64 `json:"session_id"` // [Required] <p>The identifier of livestream session.</p>
	Title string `json:"title"` // [Required] <p>The title of the livestream session, cannot exceed 200 characters.</p>
}
type UpdateSessionResponse struct {
	BaseResponse // Common response fields
	Response interface{} `json:"response"` // Actual response data
}
type UpdateShowItemRequest struct {
	ItemId int64 `json:"item_id"` // [Required] <p>Shopee's unique identifier for an item.</p>
	SessionId int64 `json:"session_id"` // [Required] <p>The identifier of livestream session.</p>
}
type UpdateShowItemResponse struct {
	BaseResponse // Common response fields
	Response interface{} `json:"response"` // Actual response data
}
type UploadImageRequest struct {
	Image interface{} `json:"image"` // [Required] <p>The image file to upload.</p>
}
type UploadImageResponse struct {
	BaseResponse // Common response fields
	Response UploadImageResponseData `json:"response"` // <p>Detail informations you are querying.</p>
}
type UploadImageResponseData struct {
	ImageUrl string `json:"image_url"` // [Required] <p>The image URL</p>
}
