package goshopee

type AddGlobalItemRequest struct {
	AttributeList   []Attribute      `json:"attribute_list,omitempty"`   // [Optional] Item attributes.
	Brand           *Brand           `json:"brand,omitempty"`            // [Optional]
	CategoryId      int64            `json:"category_id"`                // [Required] Category id of global item.
	Condition       *string          `json:"condition,omitempty"`        // [Optional] Condition of global item, "NEW" or "USED" is available.
	Description     string           `json:"description"`                // [Required] Description of global item.
	DescriptionInfo *DescriptionInfo `json:"description_info,omitempty"` // [Optional] New description field. New description field. Only whitelist sellers can use it. If you use the field, please upload the description_type=extended otherwise api will return error. If you don't use this field, you don't need to upload the description_type or upload description_type=normal
	DescriptionType *DescriptionType `json:"description_type,omitempty"` // [Optional] Values: See Data Definition- description_type (normal , extended). If you want to use extended_description, this field must be inputed
	Dimension       *Dimension       `json:"dimension,omitempty"`        // [Optional] Dimension information of global item.
	DsCatRcmdId     *string          `json:"ds_cat_rcmd_id,omitempty"`   // [Optional] <p>category recommendation service id<br /></p>
	GlobalItemName  string           `json:"global_item_name"`           // [Required] Name of global item.
	GlobalItemSku   *string          `json:"global_item_sku,omitempty"`  // [Optional] Sku of global item.
	Image           *Image           `json:"image,omitempty"`            // [Optional] Image information of global item.
	NormalStock     *int64           `json:"normal_stock,omitempty"`     // [Optional] Normal stock of global item.
	OriginalPrice   float64          `json:"original_price"`             // [Required] Original price of global item.
	PreOrder        *PreOrder        `json:"pre_order"`                  // [Required] Preorder information of global item.
	SellerStock     []SellerStock    `json:"seller_stock,omitempty"`     // [Optional] <p>seller_stock&nbsp;of global item.<br /></p>
	SizeChartInfo   *SizeChartInfo   `json:"size_chart_info,omitempty"`  // [Optional]
	VideoUploadId   []string         `json:"video_upload_id,omitempty"`  // [Optional] Video upload id of global item. Only accept one video_upload_id at most.
	Weight          float64          `json:"weight"`                     // [Required] Weight of global item.
}
type AddGlobalItemResponse struct {
	BaseResponse                           // Common response fields
	Response     AddGlobalItemResponseData `json:"response"` //
}
type AddGlobalItemResponseData struct {
	GlobalItemId int64 `json:"global_item_id"` // [Required] Id of added global item.
}
type AddGlobalModelRequest struct {
	GlobalItemId int64         `json:"global_item_id"` // [Required] ID of global item.
	GlobalModel  []GlobalModel `json:"global_model"`   // [Required] Global model setting list. Limit is  [1,50].
}
type AddGlobalModelResponse struct {
	BaseResponse // Common response fields
}
type AttributeInfo struct {
	InputType           int64    `json:"input_type"`            // [Required] <p>SINGLE_DROP_DOWN = 1</p><p>SINGLE_COMBO_BOX = 2</p><p>FREE_TEXT_FILED&nbsp; &nbsp; &nbsp; &nbsp; = 3</p><p>MULTI_DROP_DOWN&nbsp; &nbsp;= 4</p><p>MULTI_COMBO_BOX&nbsp; &nbsp;= 5<br /></p>
	InputValidationType int64    `json:"input_validation_type"` // [Required] <p>VALIDATOR_NO_VALIDATE_TYPE =&nbsp; 0</p><p>VALIDATOR_INT_TYPE = 1&nbsp;</p><p>VALIDATOR_STRING_TYPE = 2</p><p>VALIDATOR_FLOAT_TYPE = 3&nbsp;</p><p>VALIDATOR_DATE_TYPE = 4</p>
	FormatType          int64    `json:"format_type"`           // [Required] <p>FORMAT_NORMAL = 1</p><p>FORMAT_QUANTITATIVE_WITH_UNIT = 2<br /></p>
	DateFormatType      int64    `json:"date_format_type"`      // [Required] <p>YEAR_MONTH_DATE = 0 (DD/MM/YYYY)</p><p>YEAR_MONTH = 1 (MM/YYYY)<br /></p>
	AttributeUnitList   []string `json:"attribute_unit_list"`   // [Required] <p>Attribute's available units list<br /></p>
	MandatoryRegion     []string `json:"mandatory_region"`      // [Required] <p>Attribute is mandatory for these regions</p>
	MaxValueCount       int64    `json:"max_value_count"`       // [Required] <p>Max selected value count<br /></p>
	Introduction        string   `json:"introduction"`          // [Required] <p>introduction of special attribute<br /></p>
	IsOem               bool     `json:"is_oem"`                // [Required]
	SupportSearchValue  bool     `json:"support_search_value"`  // [Required] <p>Indicates whether this attribute has searchable values.</p><p>If yes, please call v2.global_product.search_global_attribute_value_list&nbsp;to get the default values</p>
}
type AttributeTree struct {
	AttributeId        int64                         `json:"attribute_id"`         // [Required] <p>Attribute ID<br /></p>
	Mandatory          bool                          `json:"mandatory"`            // [Required] <p>Is mandatory or not<br /></p>
	Name               string                        `json:"name"`                 // [Required] <p>Attribute Name<br /></p>
	AttributeValueList []AttributeTreeAttributeValue `json:"attribute_value_list"` // [Required] <p>All available values for this attribute<br /></p>
	AttributeInfo      *AttributeInfo                `json:"attribute_info"`       // [Required] <p>Attribute extra info<br /></p>
	MultiLang          []MultiLang                   `json:"multi_lang"`           // [Required] <p>Translate result for attribute name display<br /></p>
}
type AttributeTreeAttributeValue struct {
	ValueId            int64         `json:"value_id"`             // [Required] <p>Value ID<br /></p>
	Name               string        `json:"name"`                 // [Required] <p>Value name<br /></p>
	ValueUnit          string        `json:"value_unit"`           // [Required] <p>Value unit<br /></p>
	ChildAttributeList []interface{} `json:"child_attribute_list"` // [Required] <p>Child attributes for the value of parent attribute<br />The structure content is the same as attribute_tree<br /></p>
	MultiLang          *MultiLang    `json:"multi_lang"`           // [Required] <p>Translate results for value name display<br /></p>
}
type CategoryRecommendRequest struct {
	GlobalItemName          string  `json:"global_item_name" url:"global_item_name"`                                         // [Required] name of item
	GlobalProductCoverImage *string `json:"global_product_cover_image,omitempty" url:"global_product_cover_image,omitempty"` // [Optional] <p>Please use the image id returned by v2.media_space.upload_image api, we will ignore if this field is empty string<br /></p>
}
type CategoryRecommendResponse struct {
	BaseResponse                               // Common response fields
	Response     CategoryRecommendResponseData `json:"response"` //
}
type CategoryRecommendResponseData struct {
	CategoryId []int64 `json:"category_id"` // [Required] Shopee's unique identifier for a category.
}
type CreatePublishTaskRequest struct {
	GlobalItemId int64                         `json:"global_item_id"` // [Required] Id of global item.
	Item         *CreatePublishTaskRequestItem `json:"item,omitempty"` // [Optional] Item information.
	ShopRegion   string                        `json:"shop_region"`    // [Required] Region of shop.
}
type CreatePublishTaskRequestItem struct {
	ItemName                 *string                             `json:"item_name,omitempty"`                  // [Optional] <p>Name of item.&nbsp;If you upload this field, we will take your value, so you should pass the value in the local language, if you don't upload this field, Shopee will automatically translate your global product name into the local language.</p>
	Description              *string                             `json:"description,omitempty"`                // [Optional] <p>Description of item.&nbsp;If you upload this field, we will take your value, so you should pass the value in the local language, if you don't upload this field, Shopee will automatically translate your global product description into the local language.</p>
	ItemStatus               *ItemStatus                         `json:"item_status,omitempty"`                // [Optional] Status of item.
	OriginalPrice            *float64                            `json:"original_price,omitempty"`             // [Optional] <p>Original price of item.</p><p><b><font color="#c24f4a">For&nbsp;SG/MY/BR/MX/PL/ES/AR seller:</font></b>&nbsp;Sellers can set the price with two decimal place,&nbsp;other regions can only set the price as an integer.&nbsp;If you upload this field, we will take your value, so you should pass the value in local currency, if you don't upload this field, Shopee will automatically calculate the price.<br /></p>
	Image                    *Image                              `json:"image,omitempty"`                      // [Optional] Image information of item.
	Model                    []CreatePublishTaskRequestItemModel `json:"model,omitempty"`                      // [Optional] Model information of item.
	SizeChart                *string                             `json:"size_chart,omitempty"`                 // [Optional] <p>Size chart of item. Only support image_id for now</p>
	Logistic                 []Logistic                          `json:"logistic,omitempty"`                   // [Optional] Logistic information of item.
	PreOrder                 *ItemPreOrder                       `json:"pre_order,omitempty"`                  // [Optional] Preorder information of item.
	DescriptionInfo          *DescriptionInfo                    `json:"description_info,omitempty"`           // [Optional] <p>New description field. Only whitelist sellers can use it. If you use the field, please upload the description_type=extended otherwise api will return error. If you don't use this field, you don't need to upload the description_type or upload description_type=normal.&nbsp;If you upload this field, we will take your value, so you should pass the value in the local language, if you don't upload this field, Shopee will automatically translate your global product description into the local language.</p>
	StandardiseTierVariation []StandardiseTierVariation          `json:"standardise_tier_variation,omitempty"` // [Optional]
}
type CreatePublishTaskRequestItemModel struct {
	TierIndex     []int64 `json:"tier_index"`             // [Required] Tier index of model.
	OriginalPrice float64 `json:"original_price"`         // [Required] <p>Original price of model.&nbsp;If you upload this field, we will take your value, so you should pass the value in local currency, if you don't upload this field, Shopee will automatically calculate the price.</p>
	ModelStatus   *string `json:"model_status,omitempty"` // [Optional] <p>can be&nbsp;"NORMAL" or "UNAVAILABLE". Normal models can be sold on the buyer's side, and UNAVAILABLE models cannot be sold on the buyer's side. If you do not upload this field, the model status will be considered as "NORMAL".</p>
}
type CreatePublishTaskResponse struct {
	BaseResponse                               // Common response fields
	Response     CreatePublishTaskResponseData `json:"response"` //
}
type CreatePublishTaskResponseData struct {
	PublishTaskId int64 `json:"publish_task_id"` // [Required] The id of publish task.
}
type DeleteGlobalItemRequest struct {
	GlobalItemId int64 `json:"global_item_id"` // [Required] The id of global item to delete.
}
type DeleteGlobalItemResponse struct {
	BaseResponse                              // Common response fields
	Response     DeleteGlobalItemResponseData `json:"response"` //
}
type DeleteGlobalItemResponseData struct {
	FailureDeleteItem []FailureDeleteItem `json:"failure_delete_item"` // [Required] If delete failed, this field shows the details.
}
type DeleteGlobalModelRequest struct {
	GlobalItemId  int64 `json:"global_item_id"`  // [Required] Shopee's unique identifier for an global item.
	GlobalModelId int64 `json:"global_model_id"` // [Required] Shopee's unique identifier for an global model.
}
type DeleteGlobalModelResponse struct {
	BaseResponse                               // Common response fields
	Response     DeleteGlobalModelResponseData `json:"response"` //
}
type DeleteGlobalModelResponseData struct {
	GlobalModelId int64      `json:"global_model_id"` // [Required] Global model id.
	Failures      []Failures `json:"failures"`        // [Required]
}
type DtsLimit struct {
	DaysToShipRangeList []StockLimit `json:"days_to_ship_range_list"` // [Required] <p>Allowed limit scope for Pre order</p>
}
type FailureDeleteItem struct {
	ShopId int64 `json:"shop_id"` // [Required] The id of shop corresponding to the related item failed to delete.
	ItemId int64 `json:"item_id"` // [Required] The id of related item failed to delete.
}
type GetAttributeTreeRequest struct {
	CategoryIdList []int64 `json:"category_id_list" url:"category_id_list"`     // [Required] <p>Max count is 20</p>
	Language       *string `json:"language,omitempty" url:"language,omitempty"` // [Optional] <p>Language</p><p>Support Lanuage:</p><p>"SG": [ "en", "zh-Hans", "ms" ],&nbsp;</p><p>"MY": [ "en", "zh-Hans", "ms" ],</p><p>"PH": [ "en", "zh-Hans" ],</p><p>"VN": [ "vn", "en" ],</p><p>"ID": [ "id", "en" ],</p><p>"TH": [ "th", "en" ],</p><p>"BR": [ "pt-BR", "en" ],</p><p>"MX": [ "es-MX", "en" ],</p><p>"CO": [ "es-CO", "en" ],</p><p>"CL": [ "es-CL", "en" ],</p><p>"TW": [ "zh-Hant", "zh-Hans", "en" ],</p><p>"IN": [ "en", "hi" ]</p>
}
type GetAttributeTreeResponse struct {
	BaseResponse                              // Common response fields
	Response     GetAttributeTreeResponseData `json:"response"` // <p>Resopnse<br /></p>
}
type GetAttributeTreeResponseData struct {
	List []GetAttributeTreeResponseDataList `json:"list"` // [Required] <p>Each result corresponds to one category in category_ids<br /></p>
}
type GetAttributeTreeResponseDataList struct {
	AttributeTree []AttributeTree `json:"attribute_tree"` // [Required] <p>One category's attribute trees<br /></p>
	CategoryId    int64           `json:"category_id"`    // [Required] <p>Category ID<br /></p>
	Warning       string          `json:"warning"`        // [Required] <p>Warning msg</p>
}
type GetBrandListRequest struct {
	CategoryId int64 `json:"category_id" url:"category_id"` // [Required] ID of category.
	Offset     int64 `json:"offset" url:"offset"`           // [Required] Specifies the starting entry of data to return in the current call. Default is 0. if data is more than one page, the offset can be some entry to start next call.
	PageSize   int64 `json:"page_size" url:"page_size"`     // [Required] the size of one page.
	Status     int64 `json:"status" url:"status"`           // [Required] Brand status , 1: normal brand, 2: pending brand.
}
type GetBrandListResponse struct {
	BaseResponse                          // Common response fields
	Response     GetBrandListResponseData `json:"response"` //
}
type GetBrandListResponseData struct {
	BrandList   []ResponseDataBrand `json:"brand_list"`    // [Required]
	HasNextPage bool                `json:"has_next_page"` // [Required]  This is to indicate whether the item list is more than one page. If this value is true, you may want to continue to check next page to retrieve the rest of items.
	NextOffset  int64               `json:"next_offset"`   // [Required] If has_next_page is true, this value need set to next request.offset
	IsMandatory bool                `json:"is_mandatory"`  // [Required] Whether is mandatory.
	InputType   string              `json:"input_type"`    // [Required] <p>Input type: DROP_DOWN</p>
}
type GetCategoryRequest struct {
	Language *string `json:"language,omitempty" url:"language,omitempty"` // [Optional] <p>Display language. Language should be one of "zh-hans", "en"</p>
}
type GetCategoryResponse struct {
	BaseResponse                         // Common response fields
	Response     GetCategoryResponseData `json:"response"` //
}
type GetCategoryResponseData struct {
	CategoryList []Category `json:"category_list"` // [Required]
}
type GetGlobalItemIdRequest struct {
	ItemIdList []int64 `json:"item_id_list" url:"item_id_list"` // [Required] Item id list. Length limit is [1,20].
}
type GetGlobalItemIdResponse struct {
	BaseResponse                             // Common response fields
	Response     GetGlobalItemIdResponseData `json:"response"` //
}
type GetGlobalItemIdResponseData struct {
	ItemIdMap []ItemIdMap `json:"item_id_map"` // [Required]
}
type GetGlobalItemInfoRequest struct {
	GlobalItemIdList int64 `json:"global_item_id_list" url:"global_item_id_list"` // [Required] Global item id list. Length limit is [1,20].
}
type GetGlobalItemInfoResponse struct {
	BaseResponse                               // Common response fields
	Response     GetGlobalItemInfoResponseData `json:"response"` //
}
type GetGlobalItemInfoResponseData struct {
	GlobalItemList []GlobalItem `json:"global_item_list"` // [Required]
}
type GetGlobalItemLimitRequest struct {
	CategoryId *int64 `json:"category_id,omitempty" url:"category_id,omitempty"` // [Optional]
}
type GetGlobalItemLimitResponse struct {
	BaseResponse                                  // Common response fields
	Response       GetGlobalItemLimitResponseData `json:"response"`                   //
	SizeChartLimit *SizeChartLimit                `json:"size_chart_limit,omitempty"` //
}
type GetGlobalItemLimitResponseData struct {
	PriceLimit                       *PriceLimit               `json:"price_limit"`                          // [Required]
	StockLimit                       *StockLimit               `json:"stock_limit"`                          // [Required]
	GlobalItemNameLengthLimit        *StockLimit               `json:"global_item_name_length_limit"`        // [Required]
	GlobalItemImageCountLimit        *StockLimit               `json:"global_item_image_count_limit"`        // [Required]
	GlobalItemDescriptionLengthLimit *StockLimit               `json:"global_item_description_length_limit"` // [Required]
	TierVariationNameLengthLimit     *StockLimit               `json:"tier_variation_name_length_limit"`     // [Required]
	TierVariationOptionLengthLimit   *StockLimit               `json:"tier_variation_option_length_limit"`   // [Required]
	TextLengthMultiplier             float64                   `json:"text_length_multiplier"`               // [Required] Length ratio of Chinese characters to English characters in parameter verification. len(text)=len(Chinese characters)*text_length_multiplier+len(English characters )
	ExtendedDescriptionLimit         *ExtendedDescriptionLimit `json:"extended_description_limit"`           // [Required]
	DtsLimit                         *DtsLimit                 `json:"dts_limit"`                            // [Required]
	WeightLimit                      *WeightLimit              `json:"weight_limit"`                         // [Required]
	DimensionLimit                   *DimensionLimit           `json:"dimension_limit"`                      // [Required]
}
type GetGlobalItemListRequest struct {
	Offset         *string `json:"offset,omitempty" url:"offset,omitempty"`                     // [Optional] Specifies the starting entry of data to return in the current call. Default is null. if data is more than one page, the offset can be some entry to start next call.
	PageSize       int64   `json:"page_size" url:"page_size"`                                   // [Required] The size of one page. Limit is [1,50].
	UpdateTimeFrom *int64  `json:"update_time_from,omitempty" url:"update_time_from,omitempty"` // [Optional] The update_time_from and update_time_to fields specify a date range for retrieving orders (based on the item update time). The update_time_from field is the starting date range.
	UpdateTimeTo   *int64  `json:"update_time_to,omitempty" url:"update_time_to,omitempty"`     // [Optional] The update_time_from and update_time_to fields specify a date range for retrieving orders (based on the item update time). The update_time_to field is the ending date range
}
type GetGlobalItemListResponse struct {
	BaseResponse                               // Common response fields
	Response     GetGlobalItemListResponseData `json:"response"` //
}
type GetGlobalItemListResponseData struct {
	GlobalItemList []ResponseDataGlobalItem `json:"global_item_list"` // [Required]
	TotalCount     int64                    `json:"total_count"`      // [Required] Total global item count.
	HasNextPage    bool                     `json:"has_next_page"`    // [Required] This is to indicate whether the item list is more than one page. If this value is true, you may want to continue to check next page to retrieve the rest of items.
	Offset         string                   `json:"offset"`           // [Required] If has_next_page is true, this value need set to next request.offset.
}
type GetGlobalModelListRequest struct {
	GlobalItemId int64 `json:"global_item_id" url:"global_item_id"` // [Required] The id of global item.
}
type GetGlobalModelListResponse struct {
	BaseResponse                                // Common response fields
	Response     GetGlobalModelListResponseData `json:"response"` //
}
type GetGlobalModelListResponseData struct {
	TierVariation            []TierVariation                        `json:"tier_variation"`             // [Required] Tier variation information of global item.
	GlobalModel              []ResponseDataGlobalModel              `json:"global_model"`               // [Required] Global models.
	StandardiseTierVariation []ResponseDataStandardiseTierVariation `json:"standardise_tier_variation"` // [Required] <p>Standardise Tier variation information of global item.<br /></p>
}
type GetLocalAdjustmentRateResponse struct {
	BaseResponse                                    // Common response fields
	Response     GetLocalAdjustmentRateResponseData `json:"response"` //
}
type GetLocalAdjustmentRateResponseData struct {
	LocalAdjustmentRate float64 `json:"local_adjustment_rate"` // [Required] <p>The multiplier used to adjust the cross-border original price to local price</p>
}
type GetPublishableShopRequest struct {
	GlobalItemId int64   `json:"global_item_id" url:"global_item_id"`                 // [Required] Id of global item.
	ShopIdList   []int64 `json:"shop_id_list,omitempty" url:"shop_id_list,omitempty"` // [Optional] <p>Shop id list for checking if the shop is publishable.If not input the list, will return the first 300 publishable shop list in response<br /></p>
}
type GetPublishableShopResponse struct {
	BaseResponse                                // Common response fields
	Response     GetPublishableShopResponseData `json:"response"` //
}
type GetPublishableShopResponseData struct {
	PublishableShop []PublishableShop `json:"publishable_shop"` // [Required] Detail of publishable shops.
}
type GetPublishedListRequest struct {
	GlobalItemId int64   `json:"global_item_id" url:"global_item_id"`                 // [Required] Id of global item.
	ShopIdList   []int64 `json:"shop_id_list,omitempty" url:"shop_id_list,omitempty"` // [Optional] <p>Shop id list for checking if the shop is publishable.If not input the list, will return the first 300 publishable shop list in response after the&nbsp;migration period.<br /></p>
}
type GetPublishedListResponse struct {
	BaseResponse                              // Common response fields
	Response     GetPublishedListResponseData `json:"response"` //
}
type GetPublishedListResponseData struct {
	PublishedItem []PublishedItem `json:"published_item"` // [Required] Detail of published items.
}
type GetPublishTaskResultRequest struct {
	PublishTaskId int64 `json:"publish_task_id" url:"publish_task_id"` // [Required] Id of publish task.
}
type GetPublishTaskResultResponse struct {
	BaseResponse                                  // Common response fields
	Response     GetPublishTaskResultResponseData `json:"response"` //
}
type GetPublishTaskResultResponseData struct {
	PublishStatus string                                  `json:"publish_status"` // [Required] Status of publish task.
	Success       *ResponseDataSuccess                    `json:"success"`        // [Required] If publish task is successful, this field shows the published results.
	Failed        *GetPublishTaskResultResponseDataFailed `json:"failed"`         // [Required] If publish task is failed, this field shows the failed reason.
}
type GetPublishTaskResultResponseDataFailed struct {
	FailedReason string `json:"failed_reason"` // [Required] Failed reason.
}
type GetRecommendAttributeRequest struct {
	CategoryId     int64   `json:"category_id" url:"category_id"`                           // [Required] ID of category.
	CoverImageId   *string `json:"cover_image_id,omitempty" url:"cover_image_id,omitempty"` // [Optional] ID of image.
	GlobalItemName string  `json:"global_item_name" url:"global_item_name"`                 // [Required] Name of item.
}
type GetRecommendAttributeResponse struct {
	BaseResponse                                   // Common response fields
	Response     GetRecommendAttributeResponseData `json:"response"` //
}
type GetRecommendAttributeResponseData struct {
	AttributeList []ResponseDataAttribute `json:"attribute_list"` // [Required] Attribute info list.
}
type GetShopPublishableStatusRequest struct {
	GlobalItemId int64 `json:"global_item_id" url:"global_item_id"` // [Required] <p>Id of global item.<br /></p>
	Offset       int64 `json:"offset" url:"offset"`                 // [Required] <p>Specifies the starting entry of data to return in the current call. Default is 0. if data is more than one page, the offset can be some entry to start next call.<br /></p>
	PageSize     int64 `json:"page_size" url:"page_size"`           // [Required] <p>the size of one page.Max=100<br /></p>
}
type GetShopPublishableStatusResponse struct {
	BaseResponse                                      // Common response fields
	Response     GetShopPublishableStatusResponseData `json:"response"` // <p>Detail informations you are querying.<br /></p>
}
type GetShopPublishableStatusResponseData struct {
	ShopPublishableStatusList []ShopPublishableStatus `json:"shop_publishable_status_list"` // [Required] <p>Detail of publishable shops.<br /></p>
	HasNextPage               bool                    `json:"has_next_page"`                // [Required] <p>This is to indicate whether the item list is more than one page. If this value is true, you may want to continue to check next page to retrieve the rest of items.<br /></p>
	NextOffset                int64                   `json:"next_offset"`                  // [Required] <p>if has_next_page is true, this value need set to next request.offset<br /></p>
}
type GetSizeChartDetailRequest struct {
	Language    *string `json:"language,omitempty"` // [Optional] <p>language should be in the list: ["en", "zh-Hans"]<br /></p>
	SizeChartId int64   `json:"size_chart_id"`      // [Required]
}
type GetSizeChartDetailResponse struct {
	BaseResponse                                // Common response fields
	Response     GetSizeChartDetailResponseData `json:"response"` //
}
type GetSizeChartDetailResponseData struct {
	SizeChartId    int64           `json:"size_chart_id"`    // [Required] <p>ID of new size chart<br /></p>
	SizeChartName  string          `json:"size_chart_name"`  // [Required] <p>name of new size chart<br /></p>
	SizeChartTable *SizeChartTable `json:"size_chart_table"` // [Required] <p>new size chart is a table format which include multiple columns. each column has column header (measurement) and multiple values (measurement value) of this column.<br /></p>
}
type GetSizeChartListRequest struct {
	CategoryId int64  `json:"category_id"` // [Required]
	Cursor     string `json:"cursor"`      // [Required]
	PageSize   int64  `json:"page_size"`   // [Required]
}
type GetSizeChartListResponse struct {
	BaseResponse                              // Common response fields
	Response     GetSizeChartListResponseData `json:"response"` //
}
type GetSizeChartListResponseData struct {
	SizeChartList []SizeChart `json:"size_chart_list"` // [Required]
	TotalCount    int64       `json:"total_count"`     // [Required]
	NextCursor    string      `json:"next_cursor"`     // [Required]
}
type GetVariationsRequest struct {
	CategoryId int64 `json:"category_id" url:"category_id"` // [Required] <p>Leaf category id<br /></p>
}
type GetVariationsResponse struct {
	BaseResponse                                    // Common response fields
	Data                     interface{}            `json:"data,omitempty"`                       //
	StandardiseVariationList []StandardiseVariation `json:"standardise_variation_list,omitempty"` //
}
type GlobalItem struct {
	GlobalItemId          int64                      `json:"global_item_id"`           // [Required] Shopee's unique identifier for an global item.
	GlobalItemName        string                     `json:"global_item_name"`         // [Required] Name of the global item.
	Description           string                     `json:"description"`              // [Required] Description of the global item.
	GlobalItemSku         string                     `json:"global_item_sku"`          // [Required] An global item SKU (stock keeping unit) is an identifier defined by a seller, sometimes called parent SKU. Item SKU can be assigned to an item in Shopee Listings.
	GlobalItemStatus      string                     `json:"global_item_status"`       // [Required] The current status of the item. You can only query global product with normal status, otherwise api will return error.
	CreateTime            int64                      `json:"create_time"`              // [Required] Timestamp that indicates the date and time that the global item was created.
	UpdateTime            int64                      `json:"update_time"`              // [Required] Timestamp that indicates the last time that there was a change in value of the global item.
	StockInfo             []StockInfo                `json:"stock_info"`               // [Required] If the item has models, this field will not be returned, please get it through get_model_list api.
	PriceInfo             []PriceInfo                `json:"price_info"`               // [Required] If the item has models, price_info will not be returned. Please get the price of each model through the get_global_model_list api.
	Image                 *GlobalItemImage           `json:"image"`                    // [Required]
	Weight                string                     `json:"weight"`                   // [Required] <p>The weight of this global item, the unit is KG.</p><p>If set the weight of global models under this item, will return the max weight of all global models during the switching period to ensure system compatibility, please switch to call v2.global_product.get_global_model_list to get the weight of models.</p>
	Dimension             *Dimension                 `json:"dimension"`                // [Required] <p>The dimension of this global item.</p><p>If set the dimension of global models under this global item, will return the dimension with largest volume calculated by height*length*width during the switching period to ensure system compatibility, please switch to call v2.global_product.get_global_model_list to get the dimension of models.</p>
	PreOrder              *PreOrder                  `json:"pre_order"`                // [Required] <p>If set the DTS of global models under this item, will return the max DTS of all global models during the switching period to ensure system compatibility, please switch to call v2.global_product.get_global_model_list to get the DTS of models.<br /></p>
	SizeChart             string                     `json:"size_chart"`               // [Required] Url of size chart image.
	Condition             string                     `json:"condition"`                // [Required] Is it second-hand.
	HasModel              bool                       `json:"has_model"`                // [Required] Does it contain model.
	Video                 *Video                     `json:"video"`                    // [Required]
	CategoryId            int64                      `json:"category_id"`              // [Required] Shopee's unique identifier for a category.
	Brand                 *Brand                     `json:"brand"`                    // [Required]
	AttributeList         []GlobalItemAttribute      `json:"attribute_list"`           // [Required]
	DescriptionInfo       *GlobalItemDescriptionInfo `json:"description_info"`         // [Required] New description field.New description field. Only whitelist sellers can use it. If you use the field, please upload the description_type=extended otherwise api will return error. If you don't use this field, you don't need to upload the description_type or upload description_type=normal
	DescriptionType       DescriptionType            `json:"description_type"`         // [Required] Type of description : values: See Data Definition- description_type (normal , extended).
	IsFulfillmentByShopee bool                       `json:"is_fulfillment_by_shopee"` // [Required] <p>whether item is fulfillment by shopee</p>
	SizeChartId           int64                      `json:"size_chart_id"`            // [Required] <p>size_chart 模板ID</p>
}
type GlobalModel struct {
	OriginalPrice  float64       `json:"original_price"`             // [Required] Original price of global model.
	SellerStock    []SellerStock `json:"seller_stock,omitempty"`     // [Optional] <p>seller_stock of global item<br /></p>
	GlobalModelSku *string       `json:"global_model_sku,omitempty"` // [Optional] Sku of global model. model_sku length information needs to be no more than 100 characters.
	TierIndex      []int64       `json:"tier_index"`                 // [Required] <p>Tier index of global model. Index starts from 0.</p><p><br /></p><p>If you want to update one tier/two tier to no tier, can just pass the tier_variation and standardise_tier_variation as [], and pass the global_model &gt;&gt; tier_index as [], meanwhile pass the original_price, seller_stock, etc., to set the price and stock for the modified product with no tier structure.<br /></p>
	Weight         *float64      `json:"weight,omitempty"`           // [Optional] <p>The weight of this global model, the unit is KG.</p><p>If don't set the weight of this global model, will use the weight of global item by default.</p><p>If set the dimension of this global model, them must set the weight of this global model.</p>
	Dimension      *Dimension    `json:"dimension,omitempty"`        // [Optional] <p>The dimension of this global model.</p><p>If don't set the dimension of this global model, will use the dimension of global item by default.</p>
	PreOrder       *PreOrder     `json:"pre_order,omitempty"`        // [Optional] <p>Pre-order information of this global model.</p><p><br /></p><p>Notes:&nbsp;</p><p>If don't set the DTS of this global model, will use the DTS of the global item by default.</p>
}
type GlobalModelStockInfo struct {
	StockType       int64  `json:"stock_type"`        // [Required] Stock type. "1" means wms on hand, "2" means seller on hand.
	StockLocationId string `json:"stock_location_id"` // [Required] Stock location id.
	CurrentStock    int64  `json:"current_stock"`     // [Required] Current stock.
	NormalStock     int64  `json:"normal_stock"`      // [Required] Normal stock.
	ReservedStock   int64  `json:"reserved_stock"`    // [Required] Reserved stock.
}
type InitTierVariationRequest struct {
	GlobalItemId             int64                      `json:"global_item_id"`                       // [Required] ID of global item.
	GlobalModel              []GlobalModel              `json:"global_model"`                         // [Required] Model info list, model number at most 50
	StandardiseTierVariation []StandardiseTierVariation `json:"standardise_tier_variation,omitempty"` // [Optional] <p>There is at least one standardise_tier_variation and&nbsp;tier_variation.<br /></p><p><br /></p><p>If you want to update one tier/two tier to no tier, can just pass the tier_variation and standardise_tier_variation as [], and pass the global_model &gt;&gt; tier_index as [], meanwhile pass the original_price, seller_stock, etc., to set the price and stock for the modified product with no tier structure.<br /></p>
}
type InitTierVariationResponse struct {
	BaseResponse // Common response fields
}
type ItemIdMap struct {
	ItemId       int64 `json:"item_id"`        // [Required] Id of item.
	GlobalItemId int64 `json:"global_item_id"` // [Required] Id of global item.
}
type PreOrder struct {
	DaysToShip int64 `json:"days_to_ship"` // [Required] <p>Days to ship. Please get the days_to_ship range from the get_dts_limit API.</p>
}
type PriceInfo struct {
	Currency           string  `json:"currency"`              // [Required] The three-digit code representing the currency unit used for the item in Shopee Listings.
	OriginalPrice      float64 `json:"original_price"`        // [Required] The original price of the item in the listing currency.
	SipItemPrice       float64 `json:"sip_item_price"`        // [Required] SIP item price.
	SipItemPriceSource string  `json:"sip_item_price_source"` // [Required] source of sip' price. ( auto or manual).
}
type PublishableShop struct {
	ShopId     int64  `json:"shop_id"`     // [Required] Id of publishable shop.
	ShopRegion string `json:"shop_region"` // [Required] Region of published shop.
}
type PublishedItem struct {
	ShopId     int64      `json:"shop_id"`     // [Required] Shop id corresponding to the published item.
	ShopRegion string     `json:"shop_region"` // [Required] Region of shop.
	ItemId     int64      `json:"item_id"`     // [Required] Id of published item.
	ItemStatus ItemStatus `json:"item_status"` // [Required] <p>Status of published item.Applicable values: 0.DELETED(Item is deleted by seller himself),1.NORMAL, 2.BANNED,3.REVIEWING,4.INVALID(Shopee Admin deleted),5.INVALID_HIDE(Shopee Admin delete confirmed),6.BLACKLISTED(Offensive_hide),8.NORMAL_UNLIST</p>
}
type RequestBrand struct {
	BrandId *int64 `json:"brand_id,omitempty"` // [Optional] Id of brand.
}
type RequestGlobalModel struct {
	GlobalModelSku string     `json:"global_model_sku"`    // [Required] Sku of global model.
	GlobalModelId  int64      `json:"global_model_id"`     // [Required] ID of global model.
	Weight         *float64   `json:"weight,omitempty"`    // [Optional] <p>The weight of this global model, the unit is KG.</p><p>If don't set the weight of this global model, will use the weight of global item by default.</p><p>If set the dimension of this global model, them must set the weight of this global model.</p>
	Dimension      *Dimension `json:"dimension,omitempty"` // [Optional] <p>The dimension of this global model.</p><p>If don't set the dimension of this global model, will use the dimension of global item by default.</p>
	PreOrder       *PreOrder  `json:"pre_order,omitempty"` // [Optional] <p>Pre-order information of this global model.</p><p><br /></p><p>Notes:&nbsp;</p><p>If don't set the DTS of this global model, will use the DTS of the global item by default.</p>
}
type RequestModel struct {
	ModelId   int64 `json:"model_id"`   // [Required] <p>ID of model<br /></p>
	TierIndex int64 `json:"tier_index"` // [Required] <p>Model's tier_variation<br /></p>
}
type RequestPrice struct {
	GlobalModelId *int64  `json:"global_model_id,omitempty"` // [Optional] ID of global model.
	OriginalPrice float64 `json:"original_price"`            // [Required] Original price of global item.
}
type ResponseDataGlobalItem struct {
	GlobalItemId int64 `json:"global_item_id"` // [Required] Shopee's unique identifier for an global item.
	UpdateTime   int64 `json:"update_time"`    // [Required] Timestamp that indicates the last time that there was a change in value of the item, such as price/stock change.
}
type ResponseDataGlobalModel struct {
	GlobalModelId         int64                  `json:"global_model_id"`          // [Required] Id of global model.
	GlobalModelSku        string                 `json:"global_model_sku"`         // [Required] Sku of global model.
	PriceInfo             *GlobalModelPriceInfo  `json:"price_info"`               // [Required] Price info of global model.
	StockInfo             []GlobalModelStockInfo `json:"stock_info"`               // [Required] Stock info of global model.
	TierIndex             []int64                `json:"tier_index"`               // [Required] Tier index of global model.
	Weight                string                 `json:"weight"`                   // [Required] <p>The weight of this global model, the unit is KG.</p><p>If don't set the weight of this global model, will use the weight of global item by default.</p>
	Dimension             *Dimension             `json:"dimension"`                // [Required] <p>The dimension of this global model.</p><p>If don't set the dimension of this global model, will use the dimension of global item by default.</p>
	PreOrder              *PreOrder              `json:"pre_order"`                // [Required] <p>Pre-order information of this global model.</p><p><br /></p><p>Notes:&nbsp;</p><p>If don't set the DTS of this global model, will use the DTS of the global item by default.</p>
	IsFulfillmentByShopee bool                   `json:"is_fulfillment_by_shopee"` // [Required] <p>If it it a FBS model</p>
}
type ResponseDataSuccess struct {
	Region string `json:"region"`  // [Required] The region of published item.
	ShopId string `json:"shop_id"` // [Required] The shop id of published item.
	ItemId string `json:"item_id"` // [Required] The id of published item.
}
type SearchGlobalAttributeValueListRequest struct {
	AttributeId int64   `json:"attribute_id"`         // [Required]
	Cursor      int64   `json:"cursor"`               // [Required]
	Limit       int64   `json:"limit"`                // [Required] <p>The range is 1 to 100</p>
	ValueName   *string `json:"value_name,omitempty"` // [Optional]
}
type SearchGlobalAttributeValueListResponse struct {
	BaseResponse                                            // Common response fields
	DebugMessage string                                     `json:"debug_message,omitempty"` //
	Msg          string                                     `json:"msg,omitempty"`           //
	Response     SearchGlobalAttributeValueListResponseData `json:"response"`                //
}
type SearchGlobalAttributeValueListResponseData struct {
	ValueList []Value   `json:"value_list"` // [Required]
	PageInfo  *PageInfo `json:"page_info"`  // [Required]
}
type SetSyncFieldRequest struct {
	ShopSyncList []ShopSync `json:"shop_sync_list"` // [Required] Length limit is [1,50].
}
type SetSyncFieldResponse struct {
	BaseResponse // Common response fields
}
type ShopPublishableStatus struct {
	ShopId                int64  `json:"shop_id"`                 // [Required] <p>Id of publishable shop.<br /></p>
	Region                string `json:"region"`                  // [Required] <p>Region of published shop.<br /></p>
	ShopPublishableStatus bool   `json:"shop_publishable_status"` // [Required] <p>If the shop is publishable, ture means shop is publishable, fals means shop is unpublishable<br /></p>
	UnpublishableReason   string `json:"unpublishable_reason"`    // [Required] <p>Return the unpublishable reason. If the shop is publishable, will return empty for this field.<br /></p>
}
type ShopSync struct {
	ShopRegion                 string `json:"shop_region"`                    // [Required] TW TH MY BR IN SG VN
	NameAndDescription         bool   `json:"name_and_description"`           // [Required] sync name and description
	MediaInformation           bool   `json:"media_information"`              // [Required] sync media information
	TierVariationNameAndOption bool   `json:"tier_variation_name_and_option"` // [Required] sync tier variation
	Price                      bool   `json:"price"`                          // [Required] sync price
	DaysToShip                 bool   `json:"days_to_ship"`                   // [Required] sync days to ship info
}
type Stock struct {
	GlobalModelId *int64        `json:"global_model_id,omitempty"` // [Optional] ID of global model.
	SellerStock   []SellerStock `json:"seller_stock,omitempty"`    // [Optional]
}
type StockInfo struct {
	StockType       int64  `json:"stock_type"`        // [Required] The stock type.
	StockLocationId string `json:"stock_location_id"` // [Required] location_id of the stock.
	NormalStock     int64  `json:"normal_stock"`      // [Required] The normal stock quantity of the variation in the listing currency.
	ReservedStock   int64  `json:"reserved_stock"`    // [Required] The reserved stock quantity of the variation in the listing currency.
}
type SupportSizeChartRequest struct {
	CategoryId int64 `json:"category_id" url:"category_id"` // [Required] Id of category.
}
type SupportSizeChartResponse struct {
	BaseResponse                              // Common response fields
	Response     SupportSizeChartResponseData `json:"response"` //
}
type SupportSizeChartResponseData struct {
	SupportSizeChart bool `json:"support_size_chart"` // [Required] If category support size chart.
}
type UpdateGlobalItemRequest struct {
	AttributeList   []Attribute      `json:"attribute_list,omitempty"`   // [Optional] Item attributes.
	Brand           *RequestBrand    `json:"brand,omitempty"`            // [Optional]
	CategoryId      *int64           `json:"category_id,omitempty"`      // [Optional] Category id of global item.
	Condition       *string          `json:"condition,omitempty"`        // [Optional] Condition of global item, "NEW" or "USED" is available.
	Description     *string          `json:"description,omitempty"`      // [Optional] Description of global item.
	DescriptionInfo *DescriptionInfo `json:"description_info,omitempty"` // [Optional] New description field. New description field. Only whitelist sellers can use it. If you use the field, please upload the description_type=extended otherwise api will return error. If you don't use this field, you don't need to upload the description_type or upload description_type=normal
	DescriptionType *DescriptionType `json:"description_type,omitempty"` // [Optional] Values: See Data Definition- description_type (normal , extended). If you want to use extended_description or change description type ,this field must be inputed
	Dimension       *Dimension       `json:"dimension,omitempty"`        // [Optional] <p>The dimension of this global item.</p><p>Updating the dimension of this global item will overwrite the dimension of all global models under this global item.</p>
	GlobalItemId    int64            `json:"global_item_id"`             // [Required] Id of global item.
	GlobalItemName  *string          `json:"global_item_name,omitempty"` // [Optional] Name of global item.
	GlobalItemSku   *string          `json:"global_item_sku,omitempty"`  // [Optional] Sku of global item.
	Image           *Image           `json:"image,omitempty"`            // [Optional] Image information of global item.
	PreOrder        *PreOrder        `json:"pre_order,omitempty"`        // [Optional] <p>Preorder information of global item.</p><p><br /></p><p>Updating the DTS of global item will overwrite the DTS of all global models under the global item</p>
	SizeChartInfo   *SizeChartInfo   `json:"size_chart_info,omitempty"`  // [Optional]
	VideoUploadId   []string         `json:"video_upload_id,omitempty"`  // [Optional] Video upload id of global item.
	Weight          *float64         `json:"weight,omitempty"`           // [Optional] <p>The weight of this global item, the unit is KG.</p><p>Updating the weight of this&nbsp;global item will overwrite the weight of all global models under this global item.</p>
}
type UpdateGlobalItemResponse struct {
	BaseResponse                              // Common response fields
	Response     UpdateGlobalItemResponseData `json:"response"` //
}
type UpdateGlobalItemResponseData struct {
	GlobalItemId int64 `json:"global_item_id"` // [Required] Id of updated global item.
}
type UpdateGlobalModelRequest struct {
	GlobalItemId int64                `json:"global_item_id"` // [Required] ID of global item.
	GlobalModel  []RequestGlobalModel `json:"global_model"`   // [Required] Sku setting for global model. Limit is [1,50].
}
type UpdateGlobalModelResponse struct {
	BaseResponse // Common response fields
}
type UpdateLocalAdjustmentRateRequest struct {
	AdjustmentRate float64 `json:"adjustment_rate"` // [Required] <p>The multiplier used to adjust the cross-border original price to local price</p>
}
type UpdateLocalAdjustmentRateResponse struct {
	BaseResponse // Common response fields
}
type UpdatePriceRequest struct {
	GlobalItemId int64          `json:"global_item_id"` // [Required] ID of global item.
	PriceList    []RequestPrice `json:"price_list"`     // [Required] Price setting for global model. Limit is [1,50].
}
type UpdatePriceResponse struct {
	BaseResponse // Common response fields
}
type UpdateSizeChartRequest struct {
	GlobalItemId int64  `json:"global_item_id"` // [Required] Id of global item.
	SizeChart    string `json:"size_chart"`     // [Required] Image id of size chart.
}
type UpdateSizeChartResponse struct {
	BaseResponse // Common response fields
}
type UpdateStockRequest struct {
	GlobalItemId int64   `json:"global_item_id"` // [Required] ID of global item.
	StockList    []Stock `json:"stock_list"`     // [Required] Stock setting for global model. Limit is [1,50].
}
type UpdateStockResponse struct {
	BaseResponse // Common response fields
}
type UpdateTierVariationRequest struct {
	GlobalItemId             int64                      `json:"global_item_id"`                       // [Required] ID of global item.
	ModelList                []RequestModel             `json:"model_list,omitempty"`                 // [Optional]
	StandardiseTierVariation []StandardiseTierVariation `json:"standardise_tier_variation,omitempty"` // [Optional] <p>item standardise tier variation&nbsp;</p><p>There is at least one standardise_tier_variation and&nbsp;tier_variation<br /></p>
}
type UpdateTierVariationResponse struct {
	BaseResponse // Common response fields
}
