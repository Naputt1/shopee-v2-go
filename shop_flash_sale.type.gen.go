package goshopee

type AddShopFlashSaleItemsRequest struct {
	FlashSaleId int64                               `json:"flash_sale_id"` // [Required]
	Items       []AddShopFlashSaleItemsRequestItems `json:"items"`         // [Required]
}
type AddShopFlashSaleItemsRequestItems struct {
	ItemId              int64    `json:"item_id"`                          // [Required]
	PurchaseLimit       int64    `json:"purchase_limit"`                   // [Required] <p>min=0, 0 means no limit</p>
	Models              []Models `json:"models,omitempty"`                 // [Optional] <p>If the item has variation, this param is necessary.</p>
	ItemInputPromoPrice *float64 `json:"item_input_promo_price,omitempty"` // [Optional] <p>promotion price without tax of the item. If the item has no variation, this param is necessary, otherwise don't use this field<br /></p>
	ItemStock           *int64   `json:"item_stock,omitempty"`             // [Optional] <p>min=1, The campaign stock of the item. If the item has no variation, this param is necessary, otherwise don't use this field</p>
}
type AddShopFlashSaleItemsResponse struct {
	BaseResponse                                   // Common response fields
	Response     AddShopFlashSaleItemsResponseData `json:"response"` // Response data
}
type AddShopFlashSaleItemsResponseData struct {
	FailedItems []FailedItems `json:"failed_items"` // [Required]
}
type CreateShopFlashSaleRequest struct {
	TimeslotId int64 `json:"timeslot_id"` // [Required] <p>can get it from v2.shop_flash_sale.get_time_slot_id API, and you can only use the timeslot which start_time &gt; now</p>
}
type CreateShopFlashSaleResponse struct {
	BaseResponse                                 // Common response fields
	Response     CreateShopFlashSaleResponseData `json:"response"` // Response data
}
type CreateShopFlashSaleResponseData struct {
	TimeslotId  int64 `json:"timeslot_id"`   // [Required]
	FlashSaleId int64 `json:"flash_sale_id"` // [Required]
	Status      int64 `json:"status"`        // [Required] <p>the status of shop flash sale</p><p>0: deleted</p><p>1: enabled</p><p>2: disabled</p><p>3: system_rejected</p>
}
type Criteria struct {
	CriteriaId       int64   `json:"criteria_id"`        // [Required]
	MinProductRating float64 `json:"min_product_rating"` // [Required] <p>Product Rating(0.0-5.0),&nbsp;-1 means no limit</p>
	MinLikes         int64   `json:"min_likes"`          // [Required] <p>Likes(s), -1 means no limit<br /></p>
	MustNotPreOrder  bool    `json:"must_not_pre_order"` // [Required] <p>Pre-Order(s)<br /></p>
	MinOrderTotal    int64   `json:"min_order_total"`    // [Required] <p>Orders in the last 30 day(s), -1 means no limit<br /></p>
	MaxDaysToShip    int64   `json:"max_days_to_ship"`   // [Required] Days to Ship,&nbsp;<span style="font-size:16px;">-1 means no limit</span>
	MinRepetitionDay int64   `json:"min_repetition_day"` // [Required] Repetition Control (Same Product cannot Join ISFS within N Days)<br /><p>, -1 means no limit<br /></p>
	MinPromoStock    int64   `json:"min_promo_stock"`    // [Required] <p>Promo Stock,&nbsp;-1 means no limit</p>
	MaxPromoStock    int64   `json:"max_promo_stock"`    // [Required] <p>Promo Stock,&nbsp;-1 means no limit<br /></p>
	MinDiscount      int64   `json:"min_discount"`       // [Required] <p>Discount Limit, 10 means 10%, -1 means no limit</p><p><br /></p>
	MaxDiscount      int64   `json:"max_discount"`       // [Required] <p>Discount Limit, 100 means 100%,&nbsp;-1 means no limit</p>
	MinDiscountPrice int64   `json:"min_discount_price"` // [Required] <p>Discount Limit, -1 means no limit, real min discount price = min_discount_price / 100000<br /></p>
	MaxDiscountPrice int64   `json:"max_discount_price"` // [Required] <p>Discount Limit, -1 means no limit, real max discount price = max_discount_price / 100000</p>
	NeedLowestPrice  bool    `json:"need_lowest_price"`  // [Required] <p>lower than lowest price in last 7 days (exclude Shopee Flash Deals)<br /></p>
}
type DeleteShopFlashSaleItemsRequest struct {
	FlashSaleId int64   `json:"flash_sale_id"` // [Required]
	ItemIds     []int64 `json:"item_ids"`      // [Required] <p>if you delete a item, will delete all models of the item</p>
}
type DeleteShopFlashSaleItemsResponse struct {
	BaseResponse                                      // Common response fields
	Response     DeleteShopFlashSaleItemsResponseData `json:"response"` // Response data
}
type DeleteShopFlashSaleItemsResponseData struct {
	FailedItems []FailedItems `json:"failed_items"` // [Required]
}
type DeleteShopFlashSaleRequest struct {
	FlashSaleId int64 `json:"flash_sale_id"` // [Required] <p>cannot delete ongoing and expired shop flash sale</p>
}
type DeleteShopFlashSaleResponse struct {
	BaseResponse                                 // Common response fields
	Response     DeleteShopFlashSaleResponseData `json:"response"` // Response data
}
type DeleteShopFlashSaleResponseData struct {
	TimeslotId  int64 `json:"timeslot_id"`   // [Required]
	FlashSaleId int64 `json:"flash_sale_id"` // [Required]
	Status      int64 `json:"status"`        // [Required] <p>the status of shop flash sale</p><p>0: deleted</p><p>1: enabled</p><p>2: disabled</p><p>3: system_rejected</p>
}
type FailedItems struct {
	ItemId                int64                   `json:"item_id"`                // [Required]
	ModelId               int64                   `json:"model_id"`               // [Required] <p>If the item has no variation, this field will be empty</p>
	ErrCode               int64                   `json:"err_code"`               // [Required]
	ErrMsg                string                  `json:"err_msg"`                // [Required] <p>the reason why the model cannot be added</p>
	UnqualifiedConditions []UnqualifiedConditions `json:"unqualified_conditions"` // [Required] <p>if the model doesn't meet a criteria, will show the detail in this field<br /></p>
}
type FlashSale struct {
	TimeslotId       int64 `json:"timeslot_id"`        // [Required]
	FlashSaleId      int64 `json:"flash_sale_id"`      // [Required]
	Status           int64 `json:"status"`             // [Required] <p>the status of shop flash sale</p><p>0: deleted</p><p>1: enabled</p><p>2: disabled</p><p>3: system_rejected, you cannot edit the shop flash sale in 'system_rejected' status</p>
	StartTime        int64 `json:"start_time"`         // [Required] <p>the start time of shop flash sale<br /></p>
	EndTime          int64 `json:"end_time"`           // [Required] <p>the end time of shop flash sale<br /></p>
	EnabledItemCount int64 `json:"enabled_item_count"` // [Required] <p>the number of enabled items in shop flash sale</p>
	ItemCount        int64 `json:"item_count"`         // [Required] <p>the number of items in shop flash sale<br /></p>
	Type             int64 `json:"type"`               // [Required] <p>the state of shop flash sale</p><p>1: upcoming</p><p>2: ongoing</p><p>3: expired</p>
	RemindmeCount    int64 `json:"remindme_count"`     // [Required] <p>No. of Reminders Set<br /></p>
	ClickCount       int64 `json:"click_count"`        // [Required] <p>No. of Product Clicks<br /></p>
}
type GetItemCriteriaResponse struct {
	BaseResponse                             // Common response fields
	Response     GetItemCriteriaResponseData `json:"response"` // Response data
}
type GetItemCriteriaResponseData struct {
	Criteria                []Criteria `json:"criteria"`                   // [Required] <p>criteria detail</p>
	PairIds                 []PairIds  `json:"pair_ids"`                   // [Required] <p>the mapping relationship between criteria and category</p>
	OverlapBlockCategoryIds []int64    `json:"overlap_block_category_ids"` // [Required] <p>Due to regulations, the promotion of some products in these categories are prohibited in this region<br /></p>
}
type GetShopFlashSaleItemsRequest struct {
	FlashSaleId int64 `json:"flash_sale_id"` // [Required]
	Limit       int64 `json:"limit"`         // [Required] <pre>min=1,max=100</pre>
	Offset      int64 `json:"offset"`        // [Required] <pre>min=0,max=1000</pre>
}
type GetShopFlashSaleItemsResponse struct {
	BaseResponse                                   // Common response fields
	Response     GetShopFlashSaleItemsResponseData `json:"response"` // Response data
}
type GetShopFlashSaleItemsResponseData struct {
	TotalCount int64                `json:"total_count"` // [Required]
	Models     []ResponseDataModels `json:"models"`      // [Required] <p>If the item has variation, the infomation of model will be in this field</p>
	ItemInfo   []ItemInfo           `json:"item_info"`   // [Required]
}
type GetShopFlashSaleListRequest struct {
	EndTime   *int64 `json:"end_time,omitempty"`   // [Optional] <p>you should use start_time and end_time together, and start_time shoule be &lt;&nbsp;end_time<br /></p>
	Limit     int64  `json:"limit"`                // [Required] <pre>min=1,max=100</pre>
	Offset    int64  `json:"offset"`               // [Required] <pre>min=0,max=1000</pre>
	StartTime *int64 `json:"start_time,omitempty"` // [Optional] <p>you should use start_time and end_time together, and start_time shoule be &lt;&nbsp;end_time</p>
	Type      int64  `json:"type"`                 // [Required] <p>you can use this filed to search different state of shop flash sale</p><p><br /></p><p>0: all state</p><p>1: upcoming state</p><p>2: ongoing state</p><p>3: expired state</p>
}
type GetShopFlashSaleListResponse struct {
	BaseResponse                                   // Common response fields
	Response      GetShopFlashSaleListResponseData `json:"response"`                  // Response data
	FlashSaleList []FlashSale                      `json:"flash_sale_list,omitempty"` //
}
type GetShopFlashSaleListResponseData struct {
	TotalCount int64 `json:"total_count"` // [Required] <p>the number of shop flash sale that the shop has</p>
}
type GetShopFlashSaleRequest struct {
	FlashSaleId int64 `json:"flash_sale_id"` // [Required]
}
type GetShopFlashSaleResponse struct {
	BaseResponse                              // Common response fields
	Response     GetShopFlashSaleResponseData `json:"response"` // Response data
}
type GetShopFlashSaleResponseData struct {
	TimeslotId       int64 `json:"timeslot_id"`        // [Required]
	FlashSaleId      int64 `json:"flash_sale_id"`      // [Required]
	Status           int64 `json:"status"`             // [Required] <p>the status of shop flash sale</p><p>0: deleted</p><p>1: enabled</p><p>2: disabled</p>
	StartTime        int64 `json:"start_time"`         // [Required] <p>the start time of shop flash sale</p>
	EndTime          int64 `json:"end_time"`           // [Required] <p>the end time of shop flash sale<br /></p>
	EnabledItemCount int64 `json:"enabled_item_count"` // [Required] <p>the number of enabled items in shop flash sale</p>
	ItemCount        int64 `json:"item_count"`         // [Required] <p>the number of items in shop flash sale<br /></p>
	Type             int64 `json:"type"`               // [Required] <p>the state of shop flash sale</p><p>1: upcoming</p><p>2: ongoing</p><p>3: expired</p>
}
type GetTimeSlotIdRequest struct {
	EndTime   int64 `json:"end_time"`   // [Required] <pre><span style="font-family:Roboto, &quot;Helvetica Neue&quot;, Helvetica, &quot;Droid Sans&quot;, Arial, sans-serif;">should be &gt; start_time, max=2145887999</span><br /></pre>
	StartTime int64 `json:"start_time"` // [Required] <pre><span style="font-family:Roboto, &quot;Helvetica Neue&quot;, Helvetica, &quot;Droid Sans&quot;, Arial, sans-serif;">min = now, max=2145887999, should be &lt; end_time</span><br /></pre>
}
type GetTimeSlotIdResponse struct {
	BaseResponse                           // Common response fields
	Response     GetTimeSlotIdResponseData `json:"response"` // Response data
}
type GetTimeSlotIdResponseData struct {
	TimeslotId int64 `json:"timeslot_id"` // [Required]
	StartTime  int64 `json:"start_time"`  // [Required] <p>the start time of time slot</p>
	EndTime    int64 `json:"end_time"`    // [Required] <p>the end time of time slot<br /></p>
}
type ItemInfo struct {
	ItemId                int64                  `json:"item_id"`                  // [Required]
	ItemName              string                 `json:"item_name"`                // [Required]
	Status                int64                  `json:"status"`                   // [Required] <p>item status</p><p><br /></p><p>0: Deleted<br />1: Normal<br /></p><p>2: reviewing</p><p>3: banned</p><p>4: invalid</p><p>5: invalid hide</p><p>6: offensive hide</p><p>7: auditing</p><p>8: normal unlist</p>
	Image                 string                 `json:"image"`                    // [Required] <p>item image</p>
	ItemStatus            ItemStatus             `json:"item_status"`              // [Required] <p>the status of item in shop flash sale.&nbsp;If the item has variation, this field will be empty</p><p>0: disable<br /></p><p>1: enable</p><p>2: delete</p><p>4: system_rejected,&nbsp;the item is rejected by system</p><p>5: manual_rejected, the item is rejected manually</p>
	OriginalPrice         float64                `json:"original_price"`           // [Required] <p>original price of item, if the item has variation, this field will be empty</p>
	InputPromotionPrice   float64                `json:"input_promotion_price"`    // [Required] <p>promotion price without tax of item, if the item has variation, this field will be empty</p>
	PromotionPriceWithTax float64                `json:"promotion_price_with_tax"` // [Required] <p>promotion price with tax of item,&nbsp;if the item has no variation, this field will has value</p>
	PurchaseLimit         int64                  `json:"purchase_limit"`           // [Required] <p>0 means NO LIMIT</p><p><br /></p><p>purchase limit of item,&nbsp;if the item has variation, this field will be empty</p>
	CampaignStock         int64                  `json:"campaign_stock"`           // [Required] <p>campaign stock of item,&nbsp;if the item has no variation, this field will has value</p>
	Stock                 int64                  `json:"stock"`                    // [Required] <p>Active inventory&nbsp;of item,&nbsp;if the item has no variation, this field will has value</p>
	RejectReason          string                 `json:"reject_reason"`            // [Required] <p>if the item_status is 4 or 5, this field will show the reason why this item was rejected</p><p><br /></p><p>if the item has variation, this field will be empty</p>
	UnqualifiedConditions *UnqualifiedConditions `json:"unqualified_conditions"`   // [Required] <p>if the item doesn't meet a criteria, will show the detail in this field</p><p><br /></p><p>if the item has variation, this field will be empty</p>
}
type ItemsModels struct {
	ModelId         int64    `json:"model_id"`                    // [Required] <p>If the item has variation, this param is necessary.</p>
	Status          int64    `json:"status"`                      // [Required] <p>you can use this field to set the status of model</p><p><br /></p><p>0: disable<br /></p><p>1: enable</p>
	InputPromoPrice *float64 `json:"input_promo_price,omitempty"` // [Optional] <p>promotion price without tax</p><p><br /></p><p>if the model is enabled(status&nbsp; = 1) now, you can't set this field, you can only disable the model</p><p>if the model is disabled(status&nbsp; = 0) now and you want to set this field, you should also set status to 1</p>
	Stock           *int64   `json:"stock,omitempty"`             // [Optional] <p>min=1, Campaign Stock, Campaign stock can only be reserved from either Shopee stock or Seller stock<br /></p><p><br /></p><p>if the model is enabled(status&nbsp; = 1) now, you can't set this field,&nbsp;you can only disable the model</p><p>if the model is disabled(status&nbsp; = 0) now and you want to set this field, you should also set status to 1</p>
}
type Models struct {
	ModelId         int64   `json:"model_id"`          // [Required] <p>If the item has variation, this param is necessary.</p>
	InputPromoPrice float64 `json:"input_promo_price"` // [Required] <p>promotion price without tax<br /></p>
	Stock           int64   `json:"stock"`             // [Required] <p>min=1, Campaign Stock, Campaign stock can only be reserved from either Shopee stock or Seller stock<br /></p>
}
type PairIds struct {
	CriteriaId   int64             `json:"criteria_id"`   // [Required]
	CategoryList []PairIdsCategory `json:"category_list"` // [Required] <p>these are the categories that the shop has items, and the criteria will apply to these categories</p>
}
type PairIdsCategory struct {
	CategoryId int64  `json:"category_id"` // [Required] <p>o means this is All category</p>
	Name       string `json:"name"`        // [Required] <p>category name</p>
	ParentId   int64  `json:"parent_id"`   // [Required] <p>the parent category id, 0 means this category is L1 category</p>
}
type ResponseDataModels struct {
	ItemId                int64                  `json:"item_id"`                  // [Required]
	ModelId               int64                  `json:"model_id"`                 // [Required]
	ModelName             string                 `json:"model_name"`               // [Required]
	Status                int64                  `json:"status"`                   // [Required] <p>the status of model in shop flash sale</p><p>0: disable<br /></p><p>1: enable</p><p>2: delete</p><p>4: system_rejected,&nbsp;the model is rejected by system</p><p>5: manual_rejected, the model is rejected manually</p>
	OriginalPrice         float64                `json:"original_price"`           // [Required]
	InputPromotionPrice   float64                `json:"input_promotion_price"`    // [Required] <p>promotion price without tax</p>
	PromotionPriceWithTax float64                `json:"promotion_price_with_tax"` // [Required] <p>promotion price with tax<br /></p>
	PurchaseLimit         int64                  `json:"purchase_limit"`           // [Required] <p>0 means NO LIMIT</p>
	CampaignStock         int64                  `json:"campaign_stock"`           // [Required]
	Stock                 int64                  `json:"stock"`                    // [Required] <p>Active inventory<br /></p>
	RejectReason          string                 `json:"reject_reason"`            // [Required] <p>if the status is 4 or 5, this field will show the reason why this model was rejected</p>
	UnqualifiedConditions *UnqualifiedConditions `json:"unqualified_conditions"`   // [Required] <p>if the model doesn't meet a criteria, will show the detail in this field<br /></p>
}
type UnqualifiedConditions struct {
	UnqualifiedCode int64  `json:"unqualified_code"` // [Required]
	UnqualifiedMsg  string `json:"unqualified_msg"`  // [Required]
}
type UpdateShopFlashSaleItemsRequest struct {
	FlashSaleId int64                                  `json:"flash_sale_id"` // [Required]
	Items       []UpdateShopFlashSaleItemsRequestItems `json:"items"`         // [Required]
}
type UpdateShopFlashSaleItemsRequestItems struct {
	ItemId              int64         `json:"item_id"`                          // [Required]
	PurchaseLimit       *int64        `json:"purchase_limit,omitempty"`         // [Optional] <p>min=0, 0 means no limit</p><p><br /></p><p>if the item is in enabled status or the item has models in enabled status, you can't set this field</p>
	Models              []ItemsModels `json:"models,omitempty"`                 // [Optional] <p>If the item has variation, this param is necessary, otherwise please don't use this field</p>
	ItemStatus          *ItemStatus   `json:"item_status,omitempty"`            // [Optional] <p>The status of the item. If the item has no variation, this param is necessary, otherwise don't use this field</p><p><br /></p><p>you can use this field to set the status of item</p><p><br /></p><p>0: disable<br /></p><p>1: enable</p>
	ItemInputPromoPrice *float64      `json:"item_input_promo_price,omitempty"` // [Optional] <p>The promotion price of the item. If the item has no variation, you can use this field to update the promotion price of the item, otherwise don't use this field</p><p><br /></p><p>if the item is enabled(item_status&nbsp; = 1) now, you can't set this field, you can only disable the item</p><p>if the item is disabled(item_status&nbsp; = 0) now and you want to set this field, you should also set item_status to 1</p>
	ItemStock           *int64        `json:"item_stock,omitempty"`             // [Optional] <p>min=1, The campaign stock of the item. If the item has no variation, you can use this field to update the campaign stock of the item, otherwise don't use this field</p><p><br /></p><p>if the item is enabled(item_status&nbsp; = 1) now, you can't set this field,&nbsp;you can only disable the item</p><p>if the item is disabled(item_status&nbsp; = 0) now and you want to set this field, you should also set item_status to 1</p>
}
type UpdateShopFlashSaleItemsResponse struct {
	BaseResponse                                      // Common response fields
	Response     UpdateShopFlashSaleItemsResponseData `json:"response"` // Response data
}
type UpdateShopFlashSaleItemsResponseData struct {
	FailedItems []FailedItems `json:"failed_items"` // [Required]
}
type UpdateShopFlashSaleRequest struct {
	FlashSaleId int64 `json:"flash_sale_id"` // [Required]
	Status      int64 `json:"status"`        // [Required] <p>the status of shop flash sale you want to set, you cannot edit the shop flash sale in 'system_rejected' status</p><p><br /></p><p>Disabling this Flash Sale will disable all items in this session</p><p><br /></p><p>1: enable<br /></p><p>2: disbaled</p>
}
type UpdateShopFlashSaleResponse struct {
	BaseResponse                                 // Common response fields
	Response     UpdateShopFlashSaleResponseData `json:"response"` // Response data
}
type UpdateShopFlashSaleResponseData struct {
	TimeslotId  int64 `json:"timeslot_id"`   // [Required]
	FlashSaleId int64 `json:"flash_sale_id"` // [Required]
	Status      int64 `json:"status"`        // [Required] <p>the status of shop flash sale</p><p>0: deleted</p><p>1: enabled</p><p>2: disabled</p><p>3: system_rejected, you cannot edit the shop flash sale in 'system_rejected' status</p>
}
