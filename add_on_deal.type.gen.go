package goshopee

type AddAddOnDealMainItemRequest struct {
	AddOnDealId int64 `json:"add_on_deal_id"` // [Required] Shopee's unique identifier for add on deal activity.
	MainItemList []MainItem `json:"main_item_list"` // [Required] The main items added in this add on deal promotion.
}
type AddAddOnDealMainItemResponse struct {
	BaseResponse // Common response fields
	Response AddAddOnDealMainItemResponseData `json:"response"` // 
}
type AddAddOnDealMainItemResponseData struct {
	MainItemList []MainItem `json:"main_item_list"` // [Required] The main items added in this add on deal promotion.
	AddOnDealId int64 `json:"add_on_deal_id"` // [Required] Shopee's unique identifier for add on deal activity.
}
type AddAddOnDealRequest struct {
	AddOnDealName string `json:"add_on_deal_name"` // [Required] Title of the add on deal
	EndTime int64 `json:"end_time"` // [Required] The time when add on deal activity end
	PerGiftNum *int64 `json:"per_gift_num,omitempty"` // [Optional] Number of gifts that buyers can get
	PromotionPurchaseLimit *int64 `json:"promotion_purchase_limit,omitempty"` // [Optional] promotion_purchase_limit
	PromotionType int64 `json:"promotion_type"` // [Required] The type of add on deal：add on discount =0；gift with mini spend=1
	PurchaseMinSpend *float64 `json:"purchase_min_spend,omitempty"` // [Optional] The minimum purchase amount that needs to be met to buy the gift with min.Spend
	StartTime int64 `json:"start_time"` // [Required] The time when add on deal activity start.
}
type AddAddOnDealResponse struct {
	BaseResponse // Common response fields
	Response AddAddOnDealResponseData `json:"response"` // 
}
type AddAddOnDealResponseData struct {
	AddOnDealId int64 `json:"add_on_deal_id"` // [Required] Shopee's unique identifier for an add on deal activity.
}
type AddAddOnDealSubItemRequest struct {
	AddOnDealId int64 `json:"add_on_deal_id"` // [Required] Shopee's unique identifier for add on deal activity.
	SubItemList []SubItem `json:"sub_item_list"` // [Required] The sub items added in this add on deal promotion.
}
type AddAddOnDealSubItemResponse struct {
	BaseResponse // Common response fields
	Response AddAddOnDealSubItemResponseData `json:"response"` // 
}
type AddAddOnDealSubItemResponseData struct {
	SubItemList []ResponseDataSubItem `json:"sub_item_list"` // [Required] The sub items added in this add on deal promotion.
	AddOnDealId int64 `json:"add_on_deal_id"` // [Required] Shopee's unique identifier for add on deal activity.
}
type AddOnDeal struct {
	StartTime int64 `json:"start_time"` // [Required] The time when add on deal activity start.
	EndTime int64 `json:"end_time"` // [Required] The time when add on deal activity end
	PromotionType int64 `json:"promotion_type"` // [Required] The type of add on deal：add on discount =0；gift with mini spend=1
	PurchaseMinSpend float64 `json:"purchase_min_spend"` // [Required] The minimum purchase amount that needs to be met to buy the gift with min.Spend
	AddOnDealId int64 `json:"add_on_deal_id"` // [Required] Shopee's unique identifier for an add on deal activity.
	PerGiftNum int64 `json:"per_gift_num"` // [Required] Number of gifts that buyers can get
	PromotionPurchaseLimit int64 `json:"promotion_purchase_limit"` // [Required] Max. number of add-on products that a customer can purchase per order.
	AddOnDealName string `json:"add_on_deal_name"` // [Required] Title of the add on deal
	Source int64 `json:"source"` // [Required] The create source of bundle deal：Seller=1，shopee admin=0
	SubItemPrioriry []int64 `json:"sub_item_prioriry"` // [Required] The display sequence of sub item in buyer side
}
type DeleteAddOnDealMainItemRequest struct {
	AddOnDealId int64 `json:"add_on_deal_id"` // [Required] Shopee's unique identifier for add on deal activity.
	MainItemList []int64 `json:"main_item_list"` // [Required]  The main items added in this add on deal promotion.
}
type DeleteAddOnDealMainItemResponse struct {
	BaseResponse // Common response fields
	Response DeleteAddOnDealMainItemResponseData `json:"response"` // 
}
type DeleteAddOnDealMainItemResponseData struct {
	MainItemList []int64 `json:"main_item_list"` // [Required] The main items added in this add on deal promotion.
	AddOnDealId int64 `json:"add_on_deal_id"` // [Required] Shopee's unique identifier for add on deal activity.
}
type DeleteAddOnDealRequest struct {
	AddOnDealId int64 `json:"add_on_deal_id"` // [Required] Shopee's unique identifier for an add on deal activity.
}
type DeleteAddOnDealResponse struct {
	BaseResponse // Common response fields
	Response DeleteAddOnDealResponseData `json:"response"` // 
}
type DeleteAddOnDealResponseData struct {
	AddOnDealId int64 `json:"add_on_deal_id"` // [Required] Shopee's unique identifier for an add on deal activity.
}
type DeleteAddOnDealSubItemRequest struct {
	AddOnDealId int64 `json:"add_on_deal_id"` // [Required] Shopee's unique identifier for add on deal activity.
	SubItemList []RequestSubItem `json:"sub_item_list"` // [Required] The sub items added in this add on deal promotion.
}
type DeleteAddOnDealSubItemResponse struct {
	BaseResponse // Common response fields
	Response DeleteAddOnDealSubItemResponseData `json:"response"` // 
}
type DeleteAddOnDealSubItemResponseData struct {
	SubItemList []DeleteAddOnDealSubItemResponseDataSubItem `json:"sub_item_list"` // [Required] The sub items added in this add on deal promotion.
	AddOnDealId int64 `json:"add_on_deal_id"` // [Required] Shopee's unique identifier for add on deal activity.
}
type EndAddOnDealRequest struct {
	AddOnDealId int64 `json:"add_on_deal_id"` // [Required] The identifier of the API request for error tracking
}
type EndAddOnDealResponse struct {
	BaseResponse // Common response fields
	Response EndAddOnDealResponseData `json:"response"` // 
}
type EndAddOnDealResponseData struct {
	AddOnDealId int64 `json:"add_on_deal_id"` // [Required] The identifier of the API request for error tracking
}
type GetAddOnDealListRequest struct {
	PageNo *int64 `json:"page_no,omitempty"` // [Optional] The default page number is 1
	PageSize *int64 `json:"page_size,omitempty"` // [Optional] The default page size is 100
	PromotionStatus PromotionStatus `json:"promotion_status"` // [Required] The Status of add on deal，default status is all
}
type GetAddOnDealListResponse struct {
	BaseResponse // Common response fields
	Response GetAddOnDealListResponseData `json:"response"` // 
}
type GetAddOnDealListResponseData struct {
	AddOnDealList []AddOnDeal `json:"add_on_deal_list"` // [Required] The list of add on deal id
	More bool `json:"more"` // [Required] This is to indicate whether the promotion list is more than one page. If this value is true, you may want to continue to check next page to retrieve the rest of promotions.
}
type GetAddOnDealMainItemRequest struct {
	AddOnDealId int64 `json:"add_on_deal_id"` // [Required] Shopee's unique identifier for add on deal activity.
}
type GetAddOnDealMainItemResponse struct {
	BaseResponse // Common response fields
	Response GetAddOnDealMainItemResponseData `json:"response"` // 
}
type GetAddOnDealMainItemResponseData struct {
	MainItemList []MainItem `json:"main_item_list"` // [Required] The main items added in this add on deal promotion.
	AddOnDealId int64 `json:"add_on_deal_id"` // [Required] Shopee's unique identifier for add on deal activity.
}
type GetAddOnDealRequest struct {
	AddOnDealId int64 `json:"add_on_deal_id"` // [Required] Shopee's unique identifier for an add on deal activity.
}
type GetAddOnDealResponse struct {
	BaseResponse // Common response fields
	Response GetAddOnDealResponseData `json:"response"` // 
}
type GetAddOnDealResponseData struct {
	StartTime int64 `json:"start_time"` // [Required] The time when add on deal activity start.
	EndTime int64 `json:"end_time"` // [Required] The time when add on deal activity end
	PromotionType int64 `json:"promotion_type"` // [Required] The type of add on deal：add on discount =0；gift with mini spend=1
	PurchaseMinSpend float64 `json:"purchase_min_spend"` // [Required] The minimum purchase amount that needs to be met to buy the gift with min.Spend
	AddOnDealId int64 `json:"add_on_deal_id"` // [Required] Shopee's unique identifier for an add on deal activity.
	PerGiftNum int64 `json:"per_gift_num"` // [Required] Number of gifts that buyers can get
	SubItemPriority []int64 `json:"sub_item_priority"` // [Required] The order of the sub item
	PromotionPurchaseLimit int64 `json:"promotion_purchase_limit"` // [Required] Max. number of add-on products that a customer can purchase per order.
	AddOnDealName string `json:"add_on_deal_name"` // [Required] Title of the add on deal
	Source int64 `json:"source"` // [Required] 
}
type GetAddOnDealSubItemRequest struct {
	AddOnDealId int64 `json:"add_on_deal_id"` // [Required] Shopee's unique identifier for add on deal activity.
}
type GetAddOnDealSubItemResponse struct {
	BaseResponse // Common response fields
	Response GetAddOnDealSubItemResponseData `json:"response"` // 
}
type GetAddOnDealSubItemResponseData struct {
	SubItemList []GetAddOnDealSubItemResponseDataSubItem `json:"sub_item_list"` // [Required] The sub items added in this add on deal promotion.
	AddOnDealId int64 `json:"add_on_deal_id"` // [Required] Shopee's unique identifier for add on deal activity.
}
type GetAddOnDealSubItemResponseDataSubItem struct {
	ItemId int64 `json:"item_id"` // [Required] Shopee's unique identifier for an item.
	Status int64 `json:"status"` // [Required] The status of add on deal item：enable = 1；disable =2
	SubItemLimit int64 `json:"sub_item_limit"` // [Required] The purchase limit of each sub item. Only the add on discount can be set and the default limit of gift with mini.spend is 1
	ModelId int64 `json:"model_id"` // [Required] Shopee's unique identifier for a model.
	Price *Price `json:"price"` // [Required] 
}
type Price struct {
	PromoInputPrice float64 `json:"promo_input_price"` // [Required] Add-on discount price before tax
	PromoPrice float64 `json:"promo_price"` // [Required] <p>Add-on discount price after tax<br /></p>
}
type ResponseDataSubItem struct {
	ItemId int64 `json:"item_id"` // [Required] Shopee's unique identifier for an item.
	Status int64 `json:"status"` // [Required] The status of add on deal item：enable = 1；disable =2
	ModelId int64 `json:"model_id"` // [Required] Shopee's unique identifier for a model.
	FailError string `json:"fail_error"` // [Required] 
	FailMessage string `json:"fail_message"` // [Required] 
}
type SubItem struct {
	ItemId *int64 `json:"item_id,omitempty"` // [Optional] Shopee's unique identifier for an item.
	ModelId *int64 `json:"model_id,omitempty"` // [Optional] Shopee's unique identifier for a model.
	Status *int64 `json:"status,omitempty"` // [Optional] The status of add on deal item：enable = 1；disable =2
	SubItemInputPrice *float64 `json:"sub_item_input_price,omitempty"` // [Optional] Add-on discount price before tax
	SubItemLimit *int64 `json:"sub_item_limit,omitempty"` // [Optional] The purchase limit of sub item.The purchase limit of each sub item. Only the add on discount can be set and the default limit of gift with mini.spend is 1
}
type UpdateAddOnDealMainItemRequest struct {
	AddOnDealId int64 `json:"add_on_deal_id"` // [Required] Shopee's unique identifier for add on deal activity.
	MainItemList []MainItem `json:"main_item_list"` // [Required] The main items added in this add on deal promotion.
}
type UpdateAddOnDealMainItemResponse struct {
	BaseResponse // Common response fields
	AddOnDealId int64 `json:"add_on_deal_id,omitempty"` // Shopee's unique identifier for add on deal activity.
	Response UpdateAddOnDealMainItemResponseData `json:"response"` // 
}
type UpdateAddOnDealMainItemResponseData struct {
	MainItemList []MainItem `json:"main_item_list"` // [Required] The main items added in this add on deal promotion.
}
type UpdateAddOnDealRequest struct {
	AddOnDealId int64 `json:"add_on_deal_id"` // [Required] Shopee's unique identifier for an add on deal activity.
	AddOnDealName *string `json:"add_on_deal_name,omitempty"` // [Optional] Title of the add on deal
	EndTime *int64 `json:"end_time,omitempty"` // [Optional] The time when bundle deal activity end. The end time must be later than start time.
	PerGiftNum *int64 `json:"per_gift_num,omitempty"` // [Optional] Number of gifts that buyers can get
	PromotionPurchaseLimit *int64 `json:"promotion_purchase_limit,omitempty"` // [Optional] Max. number of add-on products that a customer can purchase per order.
	PurchaseMinSpend *float64 `json:"purchase_min_spend,omitempty"` // [Optional] The minimum purchase amount that needs to be met to buy the gift with min.Spend
	StartTime *int64 `json:"start_time,omitempty"` // [Optional] The time when bundle deal activity start.The start time must be 1 hour than current time.
	SubItemPriority []int64 `json:"sub_item_priority,omitempty"` // [Optional] The order of sub item
}
type UpdateAddOnDealResponse struct {
	BaseResponse // Common response fields
	Response UpdateAddOnDealResponseData `json:"response"` // 
}
type UpdateAddOnDealResponseData struct {
	StartTime int64 `json:"start_time"` // [Required] The time when add on deal activity start.
	EndTime int64 `json:"end_time"` // [Required] The time when add on deal activity end
	PromotionType int64 `json:"promotion_type"` // [Required] The type of add on deal：add on discount =0；gift with mini spend=1
	PurchaseMinSpend float64 `json:"purchase_min_spend"` // [Required] The minimum purchase amount that needs to be met to buy the gift with min.Spend
	AddOnDealId int64 `json:"add_on_deal_id"` // [Required] Shopee's unique identifier for an add on deal activity.
	PerGiftNum int64 `json:"per_gift_num"` // [Required] Number of gifts that buyers can get
	PromotionPurchaseLimit int64 `json:"promotion_purchase_limit"` // [Required] Max. number of add-on products that a customer can purchase per order.
	AddOnDealName string `json:"add_on_deal_name"` // [Required] Title of the add on deal
}
type UpdateAddOnDealSubItemRequest struct {
	AddOnDealId int64 `json:"add_on_deal_id"` // [Required] Shopee's unique identifier for add on deal activity.
	SubItemList []SubItem `json:"sub_item_list"` // [Required] The sub items added in this add on deal promotion.
}
type UpdateAddOnDealSubItemResponse struct {
	BaseResponse // Common response fields
	AddOnDealId int64 `json:"add_on_deal_id,omitempty"` // Shopee's unique identifier for add on deal activity.
	Response UpdateAddOnDealSubItemResponseData `json:"response"` // 
}
type UpdateAddOnDealSubItemResponseData struct {
	SubItemList []UpdateAddOnDealSubItemResponseDataSubItem `json:"sub_item_list"` // [Required] The sub items added in this add on deal promotion.
}
type UpdateAddOnDealSubItemResponseDataSubItem struct {
	ItemId int64 `json:"item_id"` // [Required] Shopee's unique identifier for an item.
	Status int64 `json:"status"` // [Required] The status of add on deal item：enable = 1；disable =2
	ModelId int64 `json:"model_id"` // [Required] Shopee's unique identifier for a model.
	FailError string `json:"fail_error"` // [Required] 
	FailMessage string `json:"fail_message"` // [Required] 
	SubItemInputPrice float64 `json:"sub_item_input_price"` // [Required] The discounted price of sub item
	SubItemLimit int64 `json:"sub_item_limit"` // [Required] The purchase limit of sub item.The purchase limit of each sub item. Only the add on discount can be set and the default limit of gift with mini.spend is 1
}
