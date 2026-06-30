package goshopee

type QueryBrShopBlockStatusResponse struct {
	BaseResponse // Common response fields
	Response QueryBrShopBlockStatusResponseData `json:"response"` // 
}
type QueryBrShopBlockStatusResponseData struct {
	ShopId int64 `json:"shop_id"` // [Required] <p>Shopee's unique identifier for a shop.</p>
	IsBlock bool `json:"is_block"` // [Required] <p>shop blocked status</p>
}
type QueryBrShopEnrollmentStatusResponse struct {
	BaseResponse // Common response fields
	Response QueryBrShopEnrollmentStatusResponseData `json:"response"` // 
}
type QueryBrShopEnrollmentStatusResponseData struct {
	ShopId int64 `json:"shop_id"` // [Required] <p>Shopee's unique identifier for a shop</p>
	EnrollmentStatus int64 `json:"enrollment_status"` // [Required] <p>1: enable enrollment</p><p>2: disable enrollment</p><p>3: already enrollment</p>
	EnableEnrollmentTime int64 `json:"enable_enrollment_time"` // [Required] <p>The time of this shop able to enroll FBS.</p>
}
type QueryBrShopInvoiceErrorRequest struct {
	PageNo *int64 `json:"page_no,omitempty"` // [Optional] 
	PageSize *int64 `json:"page_size,omitempty"` // [Optional] <p>max: 100</p>
}
type QueryBrShopInvoiceErrorResponse struct {
	BaseResponse // Common response fields
	Response QueryBrShopInvoiceErrorResponseData `json:"response"` // 
}
type QueryBrShopInvoiceErrorResponseData struct {
	Total int64 `json:"total"` // [Required] 
	List []QueryBrShopInvoiceErrorResponseDataList `json:"list"` // [Required] 
}
type QueryBrShopInvoiceErrorResponseDataList struct {
	ShopId int64 `json:"shop_id"` // [Required] <p>Shopee's unique identifier for a shop.</p>
	BizRequestType int64 `json:"biz_request_type"` // [Required] <p>1: Inbound</p><p>2: Return From Warehouse</p><p>3:&nbsp;Sales order invoice</p><p>4: Move Transfer</p><p>5：IA</p>
	BizRequestId string `json:"biz_request_id"` // [Required] <p>Return by default. The business FBS request order ID.</p>
	FailReason string `json:"fail_reason"` // [Required] <p>Invoice issuance failed reason.</p>
	FailType int64 `json:"fail_type"` // [Required] <p>1: sku tax info error</p><p>2: seller tax info error</p>
	InvoiceDeadlineTime int64 `json:"invoice_deadline_time"` // [Required] <p>The expired time of this failed invoice. If expired, then this request order would be cancelled.</p>
	ShopSkuList []ShopSku `json:"shop_sku_list"` // [Required] 
	InvoiceId string `json:"invoice_id"` // [Required] <p>Invoice ID</p>
	ReminderDesc string `json:"reminder_desc"` // [Required] <p>remind seller if this block issue is not solved , it will block the shop or item</p>
}
type QueryBrSkuBlockStatusRequest struct {
	ShopSkuId string `json:"shop_sku_id"` // [Required] 
}
type QueryBrSkuBlockStatusResponse struct {
	BaseResponse // Common response fields
	Response QueryBrSkuBlockStatusResponseData `json:"response"` // 
}
type QueryBrSkuBlockStatusResponseData struct {
	ShopSkuId string `json:"shop_sku_id"` // [Required] <p>itemID_modelID</p>
	IsBlock bool `json:"is_block"` // [Required] <p>product is blocked and warehouse stock cannot be sold</p>
	ShopItemId int64 `json:"shop_item_id"` // [Required] <p>ID of item</p>
	ShopModelId int64 `json:"shop_model_id"` // [Required] <p>ID of model</p>
	ShopItemName string `json:"shop_item_name"` // [Required] <p>Name of Item</p>
	ShopModelName string `json:"shop_model_name"` // [Required] <p>Name of model</p>
}
type ShopSku struct {
	ShopItemId int64 `json:"shop_item_id"` // [Required] <p>ID of item</p>
	ShopModelId int64 `json:"shop_model_id"` // [Required] <p>ID of model</p>
	ShopItemName string `json:"shop_item_name"` // [Required] <p>Name of item</p>
	ShopModelName string `json:"shop_model_name"` // [Required] <p>Name of model</p>
	FailReason string `json:"fail_reason"` // [Required] <p>Invoice issuance failed reason.</p>
}
