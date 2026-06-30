package goshopee

type AddBundleDealItemRequest struct {
	BundleDealId int64 `json:"bundle_deal_id"` // [Required] Shopee's unique identifier for a bundle deal activity.
	ItemList []MainItem `json:"item_list"` // [Required] The items added in this bundle deal promotion.
}
type AddBundleDealItemResponse struct {
	BaseResponse // Common response fields
	Response AddBundleDealItemResponseData `json:"response"` // Detail informations you are querying.
}
type AddBundleDealItemResponseData struct {
	FailedList []Failed `json:"failed_list"` // [Required] Indicate error details.
	SuccessList []interface{} `json:"success_list"` // [Required] The list of succeed added items
}
type AddBundleDealRequest struct {
	AdditionalTiers []AdditionalTiers `json:"additional_tiers,omitempty"` // [Optional] <p>Use to create tiered discount for bundle deals, a max of 2 additional tiers are allowed to create.</p><p>the rule of multiple tiers needs to follow this faq&nbsp;<a href="https://open.shopee.com/faq/53" target="_blank" style="font-size:14px;">https://open.shopee.com/faq/53</a></p><p>For additional tiers, the fix price, discount_percentage, discount_value should be consistent with tier 1</p>
	DiscountPercentage int64 `json:"discount_percentage"` // [Required] The discount that the buyer can get when buying a bundle deal. Need to input it when the bundle deal rule type is 2
	DiscountValue float64 `json:"discount_value"` // [Required] The deducted price when when buying a bundle deal. Need to input it when the bundle deal rule type is 3
	EndTime int64 `json:"end_time"` // [Required] The time when bundle deal activity end. The end time must be 1 hour later than start time.
	FixPrice float64 `json:"fix_price"` // [Required] The amount of the buyer needs to spend to purchase a bundle deal. Need to input it when the bundle deal rule type is 1
	MinAmount int64 `json:"min_amount"` // [Required] The quantity of items that need buyer to combine purchased
	Name string `json:"name"` // [Required] Title of the bundle deal
	PurchaseLimit int64 `json:"purchase_limit"` // [Required] Maximum number of bundle deals that can be bought by a buyer.
	RuleType int64 `json:"rule_type"` // [Required] The bundle deal rule type：FIX_PRICE = 1 ；DISCOUNT_PERCENTAGE = 2； DISCOUNT_VALUE = 3
	StartTime int64 `json:"start_time"` // [Required] The time when bundle deal activity start.The start time must be later than current time.
}
type AddBundleDealResponse struct {
	BaseResponse // Common response fields
	Response AddBundleDealResponseData `json:"response"` // 
}
type AddBundleDealResponseData struct {
	BundleDealId int64 `json:"bundle_deal_id"` // [Required] Shopee's unique identifier for a bundle deal activity.
}
type AdditionalTiers struct {
	MinAmount int64 `json:"min_amount"` // [Required] <p>The quantity of items that the buyers need to purchase for additional tier<br /></p>
	FixPrice float64 `json:"fix_price"` // [Required] <p>The bundle price when the buyers purchase a bundle deal for additional tiers. Need to input it when the bundle deal rule type is 1.<br /></p>
	DiscountValue float64 `json:"discount_value"` // [Required] <p>The bundle deal discount amount the buyer can save when purchasing a bundle deal. Need to input it when the bundle deal rule type is 3<br /></p>
	DiscountPercentage int64 `json:"discount_percentage"` // [Required] <p>The bundle deal discount% that the buyer can get when buying a bundle deal for additional tiers. Need to input it when the bundle deal rule type is 2<br /></p>
}
type BundleDeal struct {
	BundleDealId int64 `json:"bundle_deal_id"` // [Required] Shopee's unique identifier for a bundle deal activity.
	Name string `json:"name"` // [Required] Title of the bundle deal
	StartTime int64 `json:"start_time"` // [Required] The time when bundle deal activity start.
	EndTime int64 `json:"end_time"` // [Required] The time when bundle deal activity end.
	BundleDealRule *BundleDealBundleDealRule `json:"bundle_deal_rule"` // [Required] 
	PurchaseLimit int64 `json:"purchase_limit"` // [Required] Maximum number of bundle deals that can be bought by a buyer.
}
type BundleDealBundleDealRule struct {
	RuleType int64 `json:"rule_type"` // [Required] The bundle deal rule type：FIX_PRICE = 1 ；DISCOUNT_PERCENTAGE = 2； DISCOUNT_VALUE = 3
	DiscountValue float64 `json:"discount_value"` // [Required] The deducted price when when buying a bundle deal. Need to input it when the bundle deal rule type is 3
	FixPrice float64 `json:"fix_price"` // [Required] The amount of the buyer needs to spend to purchase a bundle deal. Need to input it when the bundle deal rule type is 1
	DiscountPercentage int64 `json:"discount_percentage"` // [Required]  The discount that the buyer can get when buying a bundle deal. Need to input it when the bundle deal rule type is 2
	MinAmount int64 `json:"min_amount"` // [Required]  The quantity of items that need buyer to combine purchased
	AdditionalTiers []AdditionalTiers `json:"additional_tiers"` // [Required] <p>Use to create tiered discount for bundle deals, a max of 2 additional tiers are allowed to create bundle deals like buy 2 get 10% off, buy 3 for 15% off, buy 4 for 20% off;&nbsp;For each tier, we will need to set the following 4 values based on bundle deal type&nbsp;+&nbsp; &nbsp; min_amount = IntAttribute()&nbsp;+&nbsp;&nbsp;&nbsp; fix_price = IntAttribute()&nbsp;+&nbsp;&nbsp;&nbsp; discount_percentage = IntAttribute()&nbsp;+&nbsp;&nbsp;&nbsp; discount_value = IntAttribute()Note: for additional tiers, the fix price, discount_percentage, discount_value should be consistent with tier 1<br /></p>
}
type BundleDealRule struct {
	RuleType int64 `json:"rule_type"` // [Required] The bundle deal rule type：FIX_PRICE = 1 ；DISCOUNT_PERCENTAGE = 2； DISCOUNT_VALUE = 3
	DiscountValue float64 `json:"discount_value"` // [Required] The deducted price when when buying a bundle deal.Need to input it when the bundle deal rule type is 3
	FixPrice float64 `json:"fix_price"` // [Required] The amount of the buyer needs to spend to purchase a bundle deal. Need to input it when the bundle deal rule type is 1
	DiscountPercentage int64 `json:"discount_percentage"` // [Required] The discount that the buyer can get when buying a bundle deal. Need to input it when the bundle deal rule type is 2
	MinAmount int64 `json:"min_amount"` // [Required] The quantity of items that need buyer to combine purchased
	AdditionalTiers *AdditionalTiers `json:"additional_tiers"` // [Required] <p>Use to create tiered discount for bundle deals, a max of 2 additional tiers are allowed to create bundle deals like buy 2 get 10% off, buy 3 for 15% off, buy 4 for 20% off;&nbsp;For each tier, we will need to set the following 4 values based on bundle deal type&nbsp;+&nbsp; &nbsp; min_amount = IntAttribute()&nbsp;+&nbsp;&nbsp;&nbsp; fix_price = IntAttribute()&nbsp;+&nbsp;&nbsp;&nbsp; discount_percentage = IntAttribute()&nbsp;+&nbsp;&nbsp;&nbsp; discount_value = IntAttribute()&nbsp;&nbsp;Note: for additional tiers, the fix price, discount_percentage, discount_value should be consistent with tier 1<br /></p>
}
type DeleteBundleDealItemRequest struct {
	BundleDealId int64 `json:"bundle_deal_id"` // [Required] Shopee's unique identifier for a bundle deal activity.
	ItemList []TopSellingProducts `json:"item_list"` // [Required] The items deleted in this bundle deal promotion.
}
type DeleteBundleDealItemResponse struct {
	BaseResponse // Common response fields
	Response DeleteBundleDealItemResponseData `json:"response"` // Detail informations you are querying.
}
type DeleteBundleDealItemResponseData struct {
	FailedList []Failed `json:"failed_list"` // [Required] Indicate error details.
	SuccessList []interface{} `json:"success_list"` // [Required] The list of succeed added items
}
type DeleteBundleDealRequest struct {
	BundleDealId int64 `json:"bundle_deal_id"` // [Required] Shopee's unique identifier for a bundle deal activity.
}
type DeleteBundleDealResponse struct {
	BaseResponse // Common response fields
	Response DeleteBundleDealResponseData `json:"response"` // Detail informations you are querying.
}
type DeleteBundleDealResponseData struct {
	BundleDealId int64 `json:"bundle_deal_id"` // [Required] Shopee's unique identifier for a bundle deal activity.
}
type EndBundleDealRequest struct {
	BundleDealId int64 `json:"bundle_deal_id"` // [Required] Shopee's unique identifier for a bundle deal activity.
}
type EndBundleDealResponse struct {
	BaseResponse // Common response fields
	Response EndBundleDealResponseData `json:"response"` // Detail informations you are querying.
}
type EndBundleDealResponseData struct {
	BundleDealId int64 `json:"bundle_deal_id"` // [Required] Shopee's unique identifier for a bundle deal activity.
}
type GetBundleDealItemRequest struct {
	BundleDealId int64 `json:"bundle_deal_id"` // [Required] Shopee's unique identifier for a bundle deal activity.
}
type GetBundleDealItemResponse struct {
	BaseResponse // Common response fields
	Response GetBundleDealItemResponseData `json:"response"` // 
}
type GetBundleDealItemResponseData struct {
	ItemList []MainItem `json:"item_list"` // [Required] The list of bundle deal item
	TotalCount int64 `json:"total_count"` // [Required] The number of  items in this bundle deal
}
type GetBundleDealListRequest struct {
	PageNo *int64 `json:"page_no,omitempty" url:"page_no,omitempty"` // [Optional] Data paging, represents the page number, starting from 1, the default is 1
	PageSize *int64 `json:"page_size,omitempty" url:"page_size,omitempty"` // [Optional] Data paging, representing the data size of each page, the maximum is 1000, the default is 20
	TimeStatus *int64 `json:"time_status,omitempty" url:"time_status,omitempty"` // [Optional] The Status of bundle deal，all=1；upcoming=2；ongoing=3，expired=4 , the default is 1
}
type GetBundleDealListResponse struct {
	BaseResponse // Common response fields
	Response GetBundleDealListResponseData `json:"response"` // 
}
type GetBundleDealListResponseData struct {
	BundleDealList []BundleDeal `json:"bundle_deal_list"` // [Required] The list of bundle deal id
	More bool `json:"more"` // [Required] this field shows whether there are more bundle deals in next page or not
}
type GetBundleDealRequest struct {
	BundleDealId int64 `json:"bundle_deal_id" url:"bundle_deal_id"` // [Required] Shopee's unique identifier for a bundle deal activity.
}
type GetBundleDealResponse struct {
	BaseResponse // Common response fields
	Response GetBundleDealResponseData `json:"response"` // Detail informations you are querying.
}
type GetBundleDealResponseData struct {
	BundleDealId int64 `json:"bundle_deal_id"` // [Required] Shopee's unique identifier for a bundle deal activity.
	Name string `json:"name"` // [Required] Title of the bundle deal
	StartTime int64 `json:"start_time"` // [Required] The time when bundle deal activity start.
	EndTime int64 `json:"end_time"` // [Required] The time when bundle deal activity end. 
	BundleDealRule *BundleDealRule `json:"bundle_deal_rule"` // [Required] 
	PurchaseLimit int64 `json:"purchase_limit"` // [Required] Maximum number of bundle deals that can be bought by a buyer.
}
type MainItem struct {
	ItemId int64 `json:"item_id"` // [Required] Shopee's unique identifier for an item.
	Status int64 `json:"status"` // [Required] The status of bundle deal item：enable = 1；disable =0
}
type UpdateBundleDealItemRequest struct {
	BundleDealId int64 `json:"bundle_deal_id"` // [Required] Shopee's unique identifier for a bundle deal activity.
	ItemList []MainItem `json:"item_list"` // [Required] The items added in this bundle deal promotion.
}
type UpdateBundleDealItemResponse struct {
	BaseResponse // Common response fields
	Response UpdateBundleDealItemResponseData `json:"response"` //  Detail informations you are querying.
}
type UpdateBundleDealItemResponseData struct {
	FailedList []Failed `json:"failed_list"` // [Required] Indicate error details.
	SuccessList []interface{} `json:"success_list"` // [Required] The list of succeed added items
}
type UpdateBundleDealRequest struct {
	AdditionalTiers *AdditionalTiers `json:"additional_tiers,omitempty"` // [Optional] <p>Use to create tiered discount for bundle deals, a max of 2 additional tiers are allowed to create bundle deals like buy 2 get 10% off, buy 3 for 15% off, buy 4 for 20% off;&nbsp;For each tier, we will need to set the following 4 values based on bundle deal type&nbsp;+&nbsp; &nbsp; min_amount = IntAttribute()&nbsp;+&nbsp;&nbsp;&nbsp; fix_price = IntAttribute()&nbsp;+&nbsp;&nbsp;&nbsp; discount_percentage = IntAttribute()&nbsp;+&nbsp;&nbsp;&nbsp; discount_value = IntAttribute()Note: for additional tiers, the fix price, discount_percentage, discount_value should be consistent with tier 1<br /></p>
	BundleDealId int64 `json:"bundle_deal_id"` // [Required] Shopee's unique identifier for a bundle deal activity.
	DiscountPercentage *int64 `json:"discount_percentage,omitempty"` // [Optional] The discount that the buyer can get when buying a bundle deal. Need to input it when the bundle deal rule type is 2
	DiscountValue *float64 `json:"discount_value,omitempty"` // [Optional]  The deducted price when when buying a bundle deal. Need to input it when the bundle deal rule type is 3
	EndTime *int64 `json:"end_time,omitempty"` // [Optional] The time when bundle deal activity end. The end time must be later than current time.
	FixPrice *float64 `json:"fix_price,omitempty"` // [Optional] The amount of the buyer needs to spend to purchase a bundle deal.Need to input it when the bundle deal rule type is 1
	MinAmount *int64 `json:"min_amount,omitempty"` // [Optional] The quantity of items that need buyer to combine purchased
	Name *string `json:"name,omitempty"` // [Optional] Title of the bundle deal
	PurchaseLimit *int64 `json:"purchase_limit,omitempty"` // [Optional] Maximum number of bundle deals that can be bought by a buyer.
	RuleType *int64 `json:"rule_type,omitempty"` // [Optional] The bundle deal rule type：FIX_PRICE = 1 ；DISCOUNT_PERCENTAGE = 2； DISCOUNT_VALUE = 3
	StartTime *int64 `json:"start_time,omitempty"` // [Optional] The time when bundle deal activity start.The start time must be later than current time.
}
type UpdateBundleDealResponse struct {
	BaseResponse // Common response fields
	Response UpdateBundleDealResponseData `json:"response"` // 
}
type UpdateBundleDealResponseData struct {
	BundleDealId int64 `json:"bundle_deal_id"` // [Required] Shopee's unique identifier for a bundle deal activity.
	Name string `json:"name"` // [Required] Title of the bundle deal
	StartTime int64 `json:"start_time"` // [Required] The time when bundle deal activity start.
	EndTime int64 `json:"end_time"` // [Required] The time when bundle deal activity end. 
	BundleDealRule *BundleDealBundleDealRule `json:"bundle_deal_rule"` // [Required] 
	PurchaseLimit int64 `json:"purchase_limit"` // [Required] Maximum number of bundle deals that can be bought by a buyer.
}
