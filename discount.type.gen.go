package goshopee

type AddDiscountItemRequest struct {
	DiscountId int64         `json:"discount_id"` // [Required] Shopee's unique identifier for a discount activity.
	ItemList   []RequestItem `json:"item_list"`   // [Required] The items added in this discount promotion.
}
type AddDiscountItemResponse struct {
	BaseResponse                             // Common response fields
	Response     AddDiscountItemResponseData `json:"response"` // Detail informations you are querying.
}
type AddDiscountItemResponseData struct {
	DiscountId int64                                       `json:"discount_id"` // [Required] Shopee's unique identifier for a discount activity.
	Count      int64                                       `json:"count"`       // [Required] The number of items that add successfully.
	ErrorList  []DeleteAddOnDealSubItemResponseDataSubItem `json:"error_list"`  // [Required] Indicate error details.
}
type AddDiscountRequest struct {
	DiscountName string `json:"discount_name"` // [Required] Title of the discount.
	EndTime      int64  `json:"end_time"`      // [Required] <p>The time when discount activity end.The end time must be 1 hour later than start time,and the discount period must be less than 180 days.</p>
	StartTime    int64  `json:"start_time"`    // [Required] The time when discount activity start.The start time must be 1 hour later than current time.
}
type AddDiscountResponse struct {
	BaseResponse                         // Common response fields
	Response     AddDiscountResponseData `json:"response"` //
}
type AddDiscountResponseData struct {
	DiscountId int64 `json:"discount_id"` // [Required] Shopee's unique identifier for a discount activity.
}
type DeleteAddOnDealSubItemResponseDataSubItem struct {
	ItemId      int64  `json:"item_id"`      // [Required] The items which have something error.
	ModelId     int64  `json:"model_id"`     // [Required] The models which have something error.
	FailMessage string `json:"fail_message"` // [Required] Indicate error details if one element hit error.
	FailError   string `json:"fail_error"`   // [Required] Indicate error type if one element hit error.
}
type DeleteDiscountItemRequest struct {
	DiscountId int64  `json:"discount_id"`        // [Required] Shopee's unique identifier for a discount activity.
	ItemId     int64  `json:"item_id"`            // [Required] Shopee's unique identifier for an item.
	ModelId    *int64 `json:"model_id,omitempty"` // [Optional] Shopee's unique identifier for a variation of an item. If there is no variation of this item, you don't need to input this param. Dafault is 0.
}
type DeleteDiscountItemResponse struct {
	BaseResponse                                // Common response fields
	Response     DeleteDiscountItemResponseData `json:"response"` // Detail informations you are querying.
}
type DeleteDiscountItemResponseData struct {
	DiscountId int64                                       `json:"discount_id"` // [Required] Shopee's unique identifier for a discount activity.
	ErrorList  []DeleteAddOnDealSubItemResponseDataSubItem `json:"error_list"`  // [Required] Detail informations about error.
}
type DeleteDiscountRequest struct {
	DiscountId int64 `json:"discount_id"` // [Required] Shopee's unique identifier for a discount activity.
}
type DeleteDiscountResponse struct {
	BaseResponse                            // Common response fields
	Response     DeleteDiscountResponseData `json:"response"` //
}
type DeleteDiscountResponseData struct {
	DiscountId int64 `json:"discount_id"` // [Required] Shopee's unique identifier for a discount activity.
	ModifyTime int64 `json:"modify_time"` // [Required] The time when discount has been deleted.
}
type DeleteSipDiscountRequest struct {
	Region string `json:"region"` // [Required] <p>The region of SIP affiliate shop that needs to delete discount.</p>
}
type DeleteSipDiscountResponse struct {
	BaseResponse                               // Common response fields
	Response     DeleteSipDiscountResponseData `json:"response"` //
}
type DeleteSipDiscountResponseData struct {
	Region string `json:"region"` // [Required] <p>The region of SIP affiliate shop that needs to delete discount.<br /></p>
}
type Discount struct {
	Status       string `json:"status"`        // [Required] The status of discount.
	DiscountName string `json:"discount_name"` // [Required] Title of the discount.
	StartTime    int64  `json:"start_time"`    // [Required] The time when discount activity start.
	EndTime      int64  `json:"end_time"`      // [Required] The time when discount activity end.
	DiscountId   int64  `json:"discount_id"`   // [Required] Shopee's unique identifier for a discount activity.
	Source       int64  `json:"source"`        // [Required] Source of the discount. 7: live stream, 1: admin, 0: others
}
type EndDiscountRequest struct {
	DiscountId int64 `json:"discount_id"` // [Required] Shopee's unique identifier for a discount activity.
}
type EndDiscountResponse struct {
	BaseResponse                         // Common response fields
	Response     EndDiscountResponseData `json:"response"` // Detail informations you are querying.
}
type EndDiscountResponseData struct {
	DiscountId int64 `json:"discount_id"` // [Required] Shopee's unique identifier for a discount activity.
	ModifyTime int64 `json:"modify_time"` // [Required] The time to track the modified time.
}
type GetDiscountListRequest struct {
	DiscountStatus string `json:"discount_status"`            // [Required] The status filter for retriveing discount list. Available value: upcoming/ongoing/expired/all.
	PageNo         int64  `json:"page_no"`                    // [Required] <p>Specifies the page number of data to return in the current call. Starting from 1. if data is more than one page, the page_no can be some entry to start next call.</p>
	PageSize       int64  `json:"page_size"`                  // [Required] If many items are available to retrieve, you may need to call GetDiscountsList multiple times to retrieve all the data. Each result set is returned as a page of entries. Use the Pagination filters to control the maximum number of entries (<= 100) to retrieve per page (i.e., per call), the offset number to start next call. This integer value is used to specify the maximum number of entries to return in a single "page" of data.
	UpdateTimeFrom *int64 `json:"update_time_from,omitempty"` // [Optional] The update_time_from and update_time_to fields specify a date range for retrieving orders (based on the discount update time). The maximum date range that may be specified with the update_time_from and update_time_to fields is 30 days.
	UpdateTimeTo   *int64 `json:"update_time_to,omitempty"`   // [Optional] The update_time_from and update_time_to fields specify a date range for retrieving orders (based on the discount update time). The maximum date range that may be specified with the update_time_from and update_time_to fields is 30 days.
}
type GetDiscountListResponse struct {
	BaseResponse                             // Common response fields
	Response     GetDiscountListResponseData `json:"response"` // Detail informations you are querying.
}
type GetDiscountListResponseData struct {
	DiscountList []Discount `json:"discount_list"` // [Required] The discounts created in this shop.
	More         bool       `json:"more"`          // [Required] <p>This is to indicate whether the item list is more than one page. If this value is true, you may want to continue to check next page to retrieve the rest of items.<br /></p>
}
type GetDiscountRequest struct {
	DiscountId int64 `json:"discount_id" url:"discount_id"` // [Required] Shopee's unique identifier for a discount activity.
	PageNo     int64 `json:"page_no" url:"page_no"`         // [Required] Specifies the page number of data to return in the current call. Starting from 1. if data is more than one page, the page_no can be some entry to start next call.
	PageSize   int64 `json:"page_size" url:"page_size"`     // [Required] Each result set is returned as a page of entries. Use the "page_size" filters to control the maximum number of entries to retrieve per page (i.e., per call), and the "page_no" to start next call. This integer value is used to specify the maximum number of entries to return in a single "page" of data.
}
type GetDiscountResponse struct {
	BaseResponse                         // Common response fields
	Response     GetDiscountResponseData `json:"response"` // Detail informations you are querying.
}
type GetDiscountResponseData struct {
	Status       string                        `json:"status"`        // [Required] The status of discount promotion
	DiscountName string                        `json:"discount_name"` // [Required] Title of the discount.
	ItemList     []GetDiscountResponseDataItem `json:"item_list"`     // [Required] The items selected in this discount.
	StartTime    int64                         `json:"start_time"`    // [Required] The time when discount activity start.
	DiscountId   int64                         `json:"discount_id"`   // [Required] Shopee's unique identifier for a discount activity.
	EndTime      int64                         `json:"end_time"`      // [Required] The time when discount activity end.
	More         bool                          `json:"more"`          // [Required] This is to indicate whether the item list is more than one page. If this value is true, you may want to continue to check next page to retrieve the rest of items.
}
type GetDiscountResponseDataItem struct {
	ItemId                            int64       `json:"item_id"`                                // [Required] Shopee's unique identifier for an item.
	ItemName                          string      `json:"item_name"`                              // [Required] Name of the item in local language.
	NormalStock                       int64       `json:"normal_stock"`                           // [Required] The current stock quantity of the item.
	ItemPromotionStock                int64       `json:"item_promotion_stock"`                   // [Required] <p>The reserved stock of the item. If the item has no variation, this param is necessary.</p>
	ItemOriginalPrice                 float64     `json:"item_original_price"`                    // [Required] The original price before discount of the item. If there is variation, this value is 0.
	ItemPromotionPrice                float64     `json:"item_promotion_price"`                   // [Required] The discount price of the item. If there is variation, this value is 0.
	ItemInflatedPriceOfOriginalPrice  float64     `json:"item_inflated_price_of_original_price"`  // [Required] <p>The original price after tax of item (Only for taxable Shop).</p>
	ItemInflatedPriceOfPromotionPrice float64     `json:"item_inflated_price_of_promotion_price"` // [Required] <p>The discount price after tax of item (Only for taxable Shop).</p>
	ItemLocalPrice                    float64     `json:"item_local_price"`                       // [Required] <p>The local price of item calculated as: Local Price = CB Original Price × Local Adjustment Rate.<br />Reflects the final local price derived from shop-level adjustment rules and is denominated in local currency.</p>
	ItemLocalPromotionPrice           float64     `json:"item_local_promotion_price"`             // [Required] <p>The local discount price of item calculated as: Local Discount Price = Local Price × Discount Rate.<br />Reflects the final local seller discount price derived from setting a seller discount&nbsp;and is denominated in local currency.</p>
	ItemLocalPriceInflated            float64     `json:"item_local_price_inflated"`              // [Required] <p>The local price after tax of item (Only for taxable Shop).</p>
	ItemLocalPromotionPriceInflated   float64     `json:"item_local_promotion_price_inflated"`    // [Required] <p>The local discount price after tax of item (Only for taxable Shop).</p>
	ModelList                         []ItemModel `json:"model_list"`                             // [Required] The models belong to this item.
	PurchaseLimit                     int64       `json:"purchase_limit"`                         // [Required] The max number of this product in the promotion price.
}
type GetSipDiscountsRequest struct {
	Region *string `json:"region,omitempty" url:"region,omitempty"` // [Optional] <p>The region of SIP affiliate shop that needs to get discount information.</p><p>If do not pass, will return the discount information set for all SIP affiliate shops.</p>
}
type GetSipDiscountsResponse struct {
	BaseResponse                             // Common response fields
	Response     GetSipDiscountsResponseData `json:"response"` //
}
type GetSipDiscountsResponseData struct {
	DiscountList []ResponseDataDiscount `json:"discount_list"` // [Required] <p>List of discounts in each region. Will be filtered based on the "region" request parameter.<br /></p>
}
type ItemModel struct {
	ModelId                            int64   `json:"model_id"`                                // [Required] Shopee's unique identifier for a variation of an item.
	ModelName                          string  `json:"model_name"`                              // [Required] Name of the variation that belongs to the same item.
	ModelNormalStock                   int64   `json:"model_normal_stock"`                      // [Required] The current stock quantity of the variation.
	ModelPromotionStock                int64   `json:"model_promotion_stock"`                   // [Required] The reserved stock of the model.
	ModelOriginalPrice                 float64 `json:"model_original_price"`                    // [Required] The original price before discount of the variation.
	ModelPromotionPrice                float64 `json:"model_promotion_price"`                   // [Required] The discount price of the variation.
	ModelInflatedPriceOfOriginalPrice  float64 `json:"model_inflated_price_of_original_price"`  // [Required] <p>The original price after tax of model (Only for taxable Shop).</p>
	ModelInflatedPriceOfPromotionPrice float64 `json:"model_inflated_price_of_promotion_price"` // [Required] <p>The discount price after tax of model (Only for taxable Shop).</p>
	ModelLocalPrice                    float64 `json:"model_local_price"`                       // [Required] <p>The local price of model calculated as:&nbsp;Local Price = CB Original Price × Local Adjustment Rate.<br />Reflects the final local price derived from shop-level adjustment rules&nbsp;and is denominated in local currency.</p>
	ModelLocalPromotionPrice           float64 `json:"model_local_promotion_price"`             // [Required] <p>The local discount price of model calculated as: Local Discount Price = Local Price × Discount Rate.<br />Reflects the final local seller discount price derived from setting a seller discount&nbsp;and is denominated in local currency.</p>
	ModelLocalPriceInflated            float64 `json:"model_local_price_inflated"`              // [Required] <p>The local price after tax of model (Only for taxable Shop).</p>
	ModelLocalPromotionPriceInflated   float64 `json:"model_local_promotion_price_inflated"`    // [Required] <p>The local discount price after tax of model (Only for taxable Shop).</p>
}
type Model struct {
	ModelId             int64   `json:"model_id"`                        // [Required] Shopee's unique identifier for a variation of an item. If there is no variation of this item, you don't need to input this param. Dafault is 0.
	ModelPromotionPrice float64 `json:"model_promotion_price"`           // [Required] The discount price of the item.
	ModelPromotionStock *int64  `json:"model_promotion_stock,omitempty"` // [Optional] <p>The reserved stock of the model, default is no limit, and can not update. To edit the promotion stock, you need to delete the exist discount and re-add again.</p>
}
type RequestItem struct {
	ItemId             int64    `json:"item_id"`                        // [Required] Shopee's unique identifier for an item.
	ItemPromotionPrice *float64 `json:"item_promotion_price,omitempty"` // [Optional] The discount price of the item. If the item has no variation, this param is necessary.
	ItemPromotionStock *int64   `json:"item_promotion_stock,omitempty"` // [Optional] The reserved stock of the item.
	ModelList          []Model  `json:"model_list,omitempty"`           // [Optional] The models which belongs to this item.
	PurchaseLimit      int64    `json:"purchase_limit"`                 // [Required] The max number of this product in the promotion price. If it's No Limit, please input the 0 for this request data.
}
type RequestItemModel struct {
	ModelId             int64   `json:"model_id"`              // [Required] Shopee's unique identifier for a variation of an item. If there is no variation of this item, you don't need to input this param. Dafault is 0.
	ModelPromotionPrice float64 `json:"model_promotion_price"` // [Required] The discount price of the item.
}
type ResponseDataDiscount struct {
	Region          string `json:"region"`            // [Required] <p>The region of SIP affiliate shop.</p>
	Status          string `json:"status"`            // [Required] <p>The status of discount for SIP affiliate shop in current region, can be upcoming/ongoing, excluding expired discounts.</p>
	SipDiscountRate int64  `json:"sip_discount_rate"` // [Required] <p>The discount rate set for SIP affiliate shop in current region.</p>
	StartTime       int64  `json:"start_time"`        // [Required] <p>The start time of discount for SIP affiliate shop in current region,&nbsp;in UNIX seconds.</p>
	EndTime         int64  `json:"end_time"`          // [Required] <p>The end time of discount for SIP affiliate shop in current region,&nbsp;in UNIX seconds.</p>
	CreateTime      int64  `json:"create_time"`       // [Required] <p>The create time of discount for SIP affiliate shop in current region,&nbsp;in UNIX seconds.</p>
	UpdateTime      int64  `json:"update_time"`       // [Required] <p>The latest update time of discount for SIP affiliate shop in current region,&nbsp;in UNIX seconds.</p>
}
type SetSipDiscountRequest struct {
	Region          string `json:"region"`            // [Required] <p>The region of SIP affiliate shop that needs to set discount.</p>
	SipDiscountRate int64  `json:"sip_discount_rate"` // [Required] <p>The overall market discount rate that will apply to all items for SIP affiliate shop in current region.</p>
}
type SetSipDiscountResponse struct {
	BaseResponse                            // Common response fields
	Response     SetSipDiscountResponseData `json:"response"` //
}
type SetSipDiscountResponseData struct {
	Region          string `json:"region"`            // [Required] <p>The region of SIP affiliate shop.<br /></p>
	Status          string `json:"status"`            // [Required] <p>The status of discount for SIP affiliate shop in current region, can be upcoming/ongoing, excluding expired discounts.</p>
	SipDiscountRate int64  `json:"sip_discount_rate"` // [Required] <p>The discount rate set for SIP affiliate shop in current region.<br /></p>
	StartTime       int64  `json:"start_time"`        // [Required] <p>The start time of discount for SIP affiliate shop in current region,&nbsp;in UNIX seconds.</p><p><br /></p><p>Note: The start time is 30 minutes after the sellers set up the sip_discount_rate.</p>
	EndTime         int64  `json:"end_time"`          // [Required] <p>The end time of discount for SIP affiliate shop in current region,&nbsp;in UNIX seconds.<br /></p><p><br /></p><p>Note:&nbsp;The end time is 180 days after the start time.</p>
	CreateTime      int64  `json:"create_time"`       // [Required] <p>The create time of discount for SIP affiliate shop in current region,&nbsp;in UNIX seconds.<br /></p>
	UpdateTime      int64  `json:"update_time"`       // [Required] <p>The latest update time of discount for SIP affiliate shop in current region,&nbsp;in UNIX seconds.<br /></p>
}
type UpdateDiscountItemRequest struct {
	DiscountId int64                           `json:"discount_id"` // [Required] Shopee's unique identifier for a discount activity.
	ItemList   []UpdateDiscountItemRequestItem `json:"item_list"`   // [Required] The items selected to this discount. You can update at most 50 items per call.
}
type UpdateDiscountItemRequestItem struct {
	ItemId             int64              `json:"item_id"`                        // [Required] Shopee's unique identifier for an item.
	ItemPromotionPrice *float64           `json:"item_promotion_price,omitempty"` // [Optional] The discount price of the item.
	ModelList          []RequestItemModel `json:"model_list,omitempty"`           // [Optional] The models selected to this discount.
	PurchaseLimit      *int64             `json:"purchase_limit,omitempty"`       // [Optional] The max number of this product in the promotion price.
}
type UpdateDiscountItemResponse struct {
	BaseResponse                                // Common response fields
	Response     UpdateDiscountItemResponseData `json:"response"` // Detail informations you are querying.
}
type UpdateDiscountItemResponseData struct {
	DiscountId int64                                       `json:"discount_id"` // [Required] Shopee's unique identifier for a discount activity.
	Count      int64                                       `json:"count"`       // [Required] The number of items that modify successfully.
	ErrorList  []DeleteAddOnDealSubItemResponseDataSubItem `json:"error_list"`  // [Required] Error list of this discount.
}
type UpdateDiscountRequest struct {
	DiscountId   int64   `json:"discount_id"`             // [Required] Shopee's unique identifier for a discount activity.
	DiscountName *string `json:"discount_name,omitempty"` // [Optional] Title of the discount.
	EndTime      *int64  `json:"end_time,omitempty"`      // [Optional] The time when discount activity end. The end time must be 1 hour later than start time.
	StartTime    *int64  `json:"start_time,omitempty"`    // [Optional] The time when discount activity start. The new start time must later than original start time.
}
type UpdateDiscountResponse struct {
	BaseResponse                            // Common response fields
	Response     UpdateDiscountResponseData `json:"response"` // Detail informations you are querying.
}
type UpdateDiscountResponseData struct {
	DiscountId int64 `json:"discount_id"` // [Required] Shopee's unique identifier for a discount activity.
	ModifyTime int64 `json:"modify_time"` // [Required] The time when discount is updated.
}
