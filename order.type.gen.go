package goshopee

type AddressBreakdown struct {
	Region          string `json:"region"`           // [Required] <p>Return region value</p><p>- PH, TH only<br /></p>
	State           string `json:"state"`            // [Required] <p>Return value</p><p>- TH: Province<br /></p>
	City            string `json:"city"`             // [Required] <p>Return value</p><p>- TH: District<br /></p>
	Town            string `json:"town"`             // [Required] <p>Return value</p><p>- TH: Sub district<br /></p>
	Postcode        string `json:"postcode"`         // [Required] <p>Return value</p><p>- TH: Postal code</p><p>- PH: Postal code<br /></p>
	DetailedAddress string `json:"detailed_address"` // [Required] <p>Return value</p><p>- PH: Additional details, i.e. street name, building</p><p>- TH: Additional details, i.e. house number<br /></p>
	AdditionalInfo  string `json:"additional_info"`  // [Required] <p>Return value:</p><p>- Empty for PH, TH<br /></p>
	FullAddress     string `json:"full_address"`     // [Required] <p>- only has value when invoice_type is personal</p><p>- Buyer address in format "detailed_address, town, district, state, postcode, additional_info" for all regions</p><p>--- for TH: leave the 'additional_info' as empty</p>
}
type BatchDownload struct {
	Start          int64  `json:"start"`                     // [Required] <p>Format YYYYMMDD</p><p>e.g. 20240101</p>
	End            int64  `json:"end"`                       // [Required] <p>Format YYYYMMDD</p><p>e.g. 20240101</p>
	DocumentType   int64  `json:"document_type"`             // [Required] <p>1 = Remessa</p><p>2 = Return</p><p>3 = Symbolic Return</p><p>4 = Sale</p><p>5 = Entrada</p><p>6 = Symbolic Remessa</p><p>7 = all</p>
	FileType       int64  `json:"file_type"`                 // [Required] <p>1 = xml only</p><p>2 = pdf only</p><p>3 = both</p>
	DocumentStatus *int64 `json:"document_status,omitempty"` // [Optional] <p>1= authorized only</p><p>2= cancelled</p><p><br /></p><p>Default:&nbsp;If document_status not passed or passed empty, means documents under ALL status (both authorized and cancelled) must be included</p>
}
type BookingItem struct {
	ItemId                 int64          `json:"item_id"`                  // [Required] <p>Shopee's unique identifier for an item.<br /></p>
	ItemName               string         `json:"item_name"`                // [Required] <p>The name of the item.<br /></p>
	ItemSku                string         `json:"item_sku"`                 // [Required] <p>A item SKU (stock keeping unit) is an identifier defined by a seller, sometimes called parent SKU. Item SKU can be assigned to an item in Shopee Listings.<br /></p>
	ModelId                int64          `json:"model_id"`                 // [Required] <p>ID of the model that belongs to the same item.<br /></p>
	ModelName              string         `json:"model_name"`               // [Required] <p>Name of the model that belongs to the same item. A seller can offer models of the same item. For example, the seller could create a fixed-priced listing for a t-shirt design and offer the shirt in different colors and sizes. In this case, each color and size combination is a separate model. Each model can have a different quantity and price.<br /></p>
	ModelSku               string         `json:"model_sku"`                // [Required] <p>A model SKU (stock keeping unit) is an identifier defined by a seller. It is only intended for the seller's use. Many sellers assign a SKU to an item of a specific type, size, and color, which are models of one item in Shopee Listings.<br /></p>
	ModelQuantityPurchased int64          `json:"model_quantity_purchased"` // [Required] <p>The number of identical items from one listing/item in the same booking.<br /></p>
	Weight                 float64        `json:"weight"`                   // [Required] <p>The weight of the item<br /></p>
	ProductLocationId      StringSlice    `json:"product_location_id"`      // [Required] <p>The fulfilment warehouse ID(s) of the items in the booking. (Multi-Warehouse sellers only)</p>
	ImageInfo              *ItemImageInfo `json:"image_info"`               // [Required] <p>Image info of the product.<br /></p>
}
type CancelOrderRequest struct {
	CancelReason          string              `json:"cancel_reason"`                      // [Required] <p>The reason seller want to cancel this order. Applicable values:&nbsp;</p><p>OUT_OF_STOCK</p><p>CUSTOMER_REQUEST</p><p>UNDELIVERABLE_AREA (Note: Only apply for TW and MY)</p><p>COD_NOT_SUPPORTED</p>
	ItemList              []RequestSubItem    `json:"item_list,omitempty"`                // [Optional] Required when cancel_reason is OUT_OF_STOCK.
	OrderSn               string              `json:"order_sn"`                           // [Required] Shopee's unique identifier for an order.
	PartialCancelItemList []PartialCancelItem `json:"partial_cancel_item_list,omitempty"` // [Optional] <p>The list of item models and quantities that the seller wants to partially cancel. This field should be provided when the seller intends to cancel only part of the order due to unavailable items while continuing to fulfill the remaining items.</p>
}
type CancelOrderResponse struct {
	BaseResponse                         // Common response fields
	Response     CancelOrderResponseData `json:"response"` // Response data
}
type CancelOrderResponseData struct {
	UpdateTime int64 `json:"update_time"` // [Required] <p>The time when the order is updated.</p>
}
type CompanyAddressBreakdown struct {
	CompanyRegion          string `json:"company_region"`           // [Required] <p>Return region value</p><p>- PH, TH only<br /></p>
	CompanyState           string `json:"company_state"`            // [Required] <p>Return value</p><p>- PH: Province</p><p>- TH: Province<br /></p>
	CompanyCity            string `json:"company_city"`             // [Required] <p>Return value</p><p>- PH: City<br /></p>
	CompanyDistrict        string `json:"company_district"`         // [Required] <p>Return value</p><p>- PH: Barangay</p><p>- TH: District<br /></p>
	CompanyTown            string `json:"company_town"`             // [Required] <p>Return value</p><p>- TH: Sub district<br /></p>
	CompanyPostcode        string `json:"company_postcode"`         // [Required] <p>Return postal code</p><p>- TH, PH only<br /></p>
	CompanyDetailedAddress string `json:"company_detailed_address"` // [Required] <p>Return value</p><p>- PH: Detailed address</p><p>- TH: Detailed address<br /></p>
	CompanyAdditionalInfo  string `json:"company_additional_info"`  // [Required] <p>Return value:</p><p>- Empty for PH, TH<br /></p>
	CompanyFullAddress     string `json:"company_full_address"`     // [Required] <p>Concatenation of company address breakdown<br /></p><p>- only has value when invoice_type is company</p><p><br /></p><p><br /></p>
}
type DownloadFbsInvoicesRequest struct {
	RequestIdList *RequestId `json:"request_id_list,omitempty"` // [Optional] <p>list of request id (task identifiers)</p>
}
type DownloadFbsInvoicesResponse struct {
	BaseResponse                                 // Common response fields
	Response     DownloadFbsInvoicesResponseData `json:"response"`            // Response data
	ErrorMsg     string                          `json:"error_msg,omitempty"` //
	Timestamp    int64                           `json:"timestamp,omitempty"` //
}
type DownloadFbsInvoicesResponseData struct {
	RequestId int64  `json:"request_id"` // [Required]
	FileLink  string `json:"file_link"`  // [Required]
}
type DownloadInvoiceDocRequest struct {
	OrderSn string `json:"order_sn" url:"order_sn"` // [Required] Shopee's unique identifier for an order.
}
type DownloadInvoiceDocResponse struct {
	BaseResponse             // Common response fields
	InvoiceDoc   interface{} `json:"invoice_doc,omitempty"` //
}
type DriverInfo struct {
	DriverName   string `json:"driver_name"`    // [Required] <p>Driver Name</p>
	DriverPhone  string `json:"driver_phone"`   // [Required] <p>Driver phone number</p>
	VehicleType  string `json:"vehicle_type"`   // [Required] <p>Delivery vehicle type</p>
	LicensePlate string `json:"license_plate"`  // [Required] <p>License plate number</p>
	CourierPhoto string `json:"courier_photo"`  // [Required] <p>URL of the driver's photo</p>
	EtaStartTime int64  `json:"eta_start_time"` // [Required] <p>Earliest estimated arrival time at pickup address</p>
	EtaEndTime   int64  `json:"eta_end_time"`   // [Required] <p>Latest estimated arrival time at pickup address</p>
	DriverStatus string `json:"driver_status"`  // [Required] <p>Driver status. Applicable values:</p><p>- Allocating Driver<br />- Driver assigned<br />- Driver is on the way<br />- Driver is arrived</p><p>-&nbsp;Driver should arrive by {starting_time} - {end_time}</p>
}
type Filter struct {
	PackageStatus       *int64   `json:"package_status,omitempty"`        // [Optional] <p>Use this field to filter the packages of specific status. Applicable values:</p><p>0: All</p><p>1: Pending</p><p>2: ToProcess</p><p>3: Processed</p><p><br /></p><p><font color="#c24f4a">Default value = 2 (ToProcess)</font></p>
	ProductLocationIds  []string `json:"product_location_ids,omitempty"`  // [Optional] <p>List of product_location_id. Use this field to filter the packages under specific warehouses.</p>
	LogisticsChannelIds []int64  `json:"logistics_channel_ids,omitempty"` // [Optional] <p>List of logistics_channel_id. Use this field to filter the packages under specific logistics channels.</p>
	FulfillmentType     *int64   `json:"fulfillment_type,omitempty"`      // [Optional] <p>Use this field to filter the packages fulfilled by shopee or seller. Applicable values:&nbsp;</p><p>0: None (not apply filter)</p><p>1: Shopee</p><p>2: Seller</p><p><br /></p><p><font color="#c24f4a">Default value = 2 (Seller)</font></p>
	InvoicePending      *bool    `json:"invoice_pending,omitempty"`       // [Optional] <p>Use this field to filter the packages under invoice_pending.</p><p><font color="#c24f4a"><br /></font></p><p><font color="#c24f4a">Default value = false</font></p>
	SortingGroup        *int64   `json:"sorting_group,omitempty"`         // [Optional] <p>[Only for TW 30029 channel] Use this field to filter the sorting group of parcel. This field is only available for package with package_status = 3 and logistics_channel_id = 30029. Applicable values:&nbsp;<br />0: All</p><p>1: North<br />2: South</p>
	OrderType           *int64   `json:"order_type,omitempty"`            // [Optional] <p>Use this field to filter packages by order type. Applicable values:</p><p>0: All</p><p>1: Regular Order</p><p>2: Instant Order</p><p><br /></p><p><font color="#c24f4a">Default value = 0 (All)</font></p><p><br /></p><p>Note:&nbsp;For VN shops, using 2: Instant Order will return both Instant Delivery and Same-day Delivery packages.</p>
	IsPreOrder          *int64   `json:"is_pre_order,omitempty"`          // [Optional] <p>Use this field to filter packages by pre-order status. Applicable values:</p><p>0: All</p><p>1:&nbsp;Pre-Order</p><p>2:&nbsp;Non Pre-Order</p><p><br /></p><p><font color="#c24f4a">Default value = 0 (All)</font></p>
	ShippingPriority    *int64   `json:"shipping_priority,omitempty"`     // [Optional] <p>Use this field to filter packages by shipping priority. Applicable values:&nbsp;</p><p><br /></p><p>- For MY/PH/TW/TH shops, and VN Preferred/Preferred Plus/Shopee Mall shops:</p><p>0: All</p><p>1: Overdue</p><p>2: Ship by Today</p><p>3: Ship by Tomorrow</p><p><br /></p><p>- For other shops：<br />0: All<br />1: Overdue<br />2: Within 24h<br />3: Beyond 24h<br /></p><p><br /></p><p><font color="#c24f4a">Default value = 0 (All)</font></p>
}
type GenerateFbsInvoicesRequest struct {
	BatchDownload *BatchDownload `json:"batch_download,omitempty"` // [Optional]
}
type GenerateFbsInvoicesResponse struct {
	BaseResponse                                         // Common response fields
	ErrorMsg     string                                  `json:"error_msg,omitempty"`   // <p>Error messages</p>
	ResultList   []GenerateFbsInvoicesResponseDataResult `json:"result_list,omitempty"` //
}
type GenerateFbsInvoicesResponseDataResult struct {
	RequestId   int64  `json:"request_id"`   // [Required] <p>Unique task identifier that includes one or more tax documents to be downloaded according to the filters sent in the request.<br /></p>
	FailError   string `json:"fail_error"`   // [Required] <p>Indicate error type if one element hit error. Empty if no error happened.</p>
	FailMessage string `json:"fail_message"` // [Required] <p>Indicate error details if one element hit error. Empty if no error happened.<br /></p>
}
type Geolocation struct {
	Latitude  float64 `json:"latitude"`  // [Required] <p>Latitude.</p>
	Longitude float64 `json:"longitude"` // [Required] <p>Longitude.</p>
}
type GetBookingDetailRequest struct {
	BookingSnList          string  `json:"booking_sn_list" url:"booking_sn_list"`                                       // [Required] <p>The set of booking_sn. If there are multiple booking_sn, you need to use English comma to connect them. limit [1,50]<br /></p>
	ResponseOptionalFields *string `json:"response_optional_fields,omitempty" url:"response_optional_fields,omitempty"` // [Optional] <p>The response fields you want to get. Please select from the below response parameters. If you input an object field, all the params under it will be included automatically in the response. If there are multiple response fields you want to get, you need to use English comma to connect them. Available values:&nbsp;item_list,cancel_by,cancel_reason,fulfillment_flag,pickup_done_time,shipping_carrier, recipient_address, dropshipper, dropshipper_phone<br /></p>
}
type GetBookingDetailResponse struct {
	BaseResponse                              // Common response fields
	Response     GetBookingDetailResponseData `json:"response"` // Response data
}
type GetBookingDetailResponseData struct {
	BookingList []ResponseDataBooking `json:"booking_list"` // [Required] <p>The list of bookings.<br /></p>
}
type GetBookingListRequest struct {
	BookingStatus  *BookingStatus `json:"booking_status,omitempty" url:"booking_status,omitempty"` // [Optional] <p>The booking_status filter for retrieving bookings and each one only every request. Available value: READY_TO_SHIP/PROCESSED/SHIPPED/CANCELLED/MATCHED<br /></p>
	Cursor         *string        `json:"cursor,omitempty" url:"cursor,omitempty"`                 // [Optional] <p>Specifies the starting entry of data to return in the current call. Default is "". If data is more than one page, the offset can be some entry to start next call.<br /></p>
	PageSize       int64          `json:"page_size" url:"page_size"`                               // [Required] <p>Each result set is returned as a page of entries. Use the "page_size" filters to control the maximum number of entries to retrieve per page (i.e., per call). This integer value is used to specify the maximum number of entries to return in a single "page" of data.The limit of page_size if between 1 and 100.<br /></p>
	TimeFrom       int64          `json:"time_from" url:"time_from"`                               // [Required] <p>The time_from and time_to fields specify a date range for retrieving bookings (based on the time_range_field). The time_from field is the starting date range. The maximum date range that may be specified with the time_from and time_to fields is 15 days.<br /></p>
	TimeRangeField string         `json:"time_range_field" url:"time_range_field"`                 // [Required] <p>The kind of time_from and time_to.&nbsp;Available value: create_time, update_time.<br /></p>
	TimeTo         int64          `json:"time_to" url:"time_to"`                                   // [Required] <p>The time_from and time_to fields specify a date range for retrieving bookings (based on the time_range_field). The time_from field is the starting date range. The maximum date range that may be specified with the time_from and time_to fields is 15 days.</p>
}
type GetBookingListResponse struct {
	BaseResponse                            // Common response fields
	Response     GetBookingListResponseData `json:"response"` // Response data
}
type GetBookingListResponseData struct {
	More        bool                                `json:"more"`         // [Required] <p>This is to indicate whether the booking list is more than one page. If this value is true, you may want to continue to check next page to retrieve bookings.<br /></p>
	BookingList []GetBookingListResponseDataBooking `json:"booking_list"` // [Required]
}
type GetBookingListResponseDataBooking struct {
	BookingSn     string        `json:"booking_sn"`     // [Required] <p>Shopee's unique identifier for a booking.<br /></p>
	OrderSn       string        `json:"order_sn"`       // [Required] <p>Shopee's unique identifier for an order. Only return if booking_status is MATCHED.<br /></p>
	BookingStatus BookingStatus `json:"booking_status"` // [Required] <p>The booking_status filter for retrieving booking and each one only every request. Available value: READY_TO_SHIP/PROCESSED/SHIPPED/CANCELLED/MATCHED<br /></p>
	NextCursor    string        `json:"next_cursor"`    // [Required] <p>If more is true, you should pass the next_cursor in the next request as cursor. The value of next_cursor will be empty string when more is false.<br /></p>
}
type GetBuyerInvoiceInfoRequest struct {
	Queries []Queries `json:"queries"` // [Required]
}
type GetBuyerInvoiceInfoResponse struct {
	BaseResponse                  // Common response fields
	InvoiceInfoList []InvoiceInfo `json:"invoice_info_list,omitempty"` //
}
type GetEstimateCancelValueRequest struct {
	OrderSn               string              `json:"order_sn"`                 // [Required] <p>Shopee's unique identifier for an order.</p>
	PartialCancelItemList []PartialCancelItem `json:"partial_cancel_item_list"` // [Required] <p>The list of item models and quantities for which the seller wants to estimate the cancellation value before submitting the actual partial cancellation request.</p>
}
type GetEstimateCancelValueResponse struct {
	BaseResponse            // Common response fields
	CancelValuePrice string `json:"cancel_value_price,omitempty"` // <p>The estimated cancellation value for the selected item quantities. This value is calculated before the actual cancellation is submitted and can be used by sellers to preview the expected cancellation amount and support partial cancellation confirmation.</p>
}
type GetFbsInvoicesResultRequest struct {
	RequestIdList *RequestId `json:"request_id_list"` // [Required] <p>-</p>
}
type GetFbsInvoicesResultResponse struct {
	BaseResponse                                          // Common response fields
	ErrorMsg     string                                   `json:"error_msg,omitempty"`   // <p>Indicate error details if hit error. Empty if no error happened.</p>
	ResultList   []GetFbsInvoicesResultResponseDataResult `json:"result_list,omitempty"` //
}
type GetFbsInvoicesResultResponseDataResult struct {
	RequestId int64  `json:"request_id"` // [Required] <p>Represents the current status of the request</p>
	FileName  string `json:"file_name"`  // [Required] <p>Name of the file&nbsp;to be downloaded</p>
	Status    string `json:"status"`     // [Required] <p>Represents the current status of the request</p>
}
type GetOrderDetailRequest struct {
	OrderSnList               string  `json:"order_sn_list" url:"order_sn_list"`                                                   // [Required] The set of order_sn. If there are multiple order_sn, you need to use English comma to connect them. limit [1,50]
	RequestOrderStatusPending *bool   `json:"request_order_status_pending,omitempty" url:"request_order_status_pending,omitempty"` // [Optional] <p>Compatible parameter during migration period, send True will let API support PENDING status and return&nbsp; pending_terms, send False or don’t send will fallback to old logic</p>
	ResponseOptionalFields    *string `json:"response_optional_fields,omitempty" url:"response_optional_fields,omitempty"`         // [Optional] <p>a response fields you want to get. Please select from the below response parameters. If you input an object field, all the params under it will be included automatically in the response. If there are multiple response fields you want to get, you need to use English comma to connect them.  Available values: buyer_user_id,buyer_username,estimated_shipping_fee,recipient_address,actual_shipping_fee ,goods_to_declare,note,note_update_time,item_list,pay_time,dropshipper, dropshipper_phone,split_up,buyer_cancel_reason,cancel_by,cancel_reason,actual_shipping_fee_confirmed,buyer_cpf_id,fulfillment_flag,pickup_done_time,package_list,shipping_carrier,payment_method,total_amount,buyer_username,invoice_data,order_chargeable_weight_gram,return_request_due_date,edt,payment_info,international_label</p>
}
type GetOrderDetailResponse struct {
	BaseResponse                            // Common response fields
	Response     GetOrderDetailResponseData `json:"response"` // Response data
}
type GetOrderDetailResponseData struct {
	OrderList []GetOrderDetailResponseDataOrder `json:"order_list"` // [Required] The list of orders.
}
type GetOrderDetailResponseDataOrder struct {
	OrderSn                               string                 `json:"order_sn"`                                  // [Required] Return by default. Shopee's unique identifier for an order.
	Region                                string                 `json:"region"`                                    // [Required] Return by default. The two-digit code representing the region where the order was made.
	Currency                              string                 `json:"currency"`                                  // [Required] Return by default. The three-digit code representing the currency unit for which the order was paid.
	Cod                                   bool                   `json:"cod"`                                       // [Required] Return by default. This value indicates whether the order was a COD (cash on delivery) order.
	TotalAmount                           float64                `json:"total_amount"`                              // [Required] The total amount paid by the buyer for the order. This amount includes the total sale price of items, shipping cost beared by buyer; and offset by Shopee promotions if applicable. This value will only return after the buyer has completed payment for the order.
	PendingTerms                          []string               `json:"pending_terms"`                             // [Required] <p>The list of pending terms. Applicable values:</p><p>- SYSTEM_PENDING: Under Shopee internal processing.</p><p>- KYC_PENDING: Under KYC checking (TW CB order only).</p><p>- ARRANGE_SHIPMENT_PENDING: Temporarily held due to 3PL capacity constraints.</p>
	PendingDescription                    []string               `json:"pending_description"`                       // [Required] <p>The value of this field is the description of pending reason corresponding with pending terms. Applicable values:&nbsp;</p><p>- For SYSTEM_PENDING: Order is being processed by Shopee.</p><p>- For KYC_PENDING: Order is pending buyer TW KYC pre-authorization.</p><p>- For ARRANGE_SHIPMENT_PENDING: Allocating delivery resources due to high order volume. Label print will be available within 4 days after buyer paid.</p>
	OrderStatus                           OrderStatus            `json:"order_status"`                              // [Required] Return by default. Enumerated type that defines the current status of the order.
	ShippingCarrier                       string                 `json:"shipping_carrier"`                          // [Required] <p>The logistics service provider that the buyer selected for the order to deliver items.&nbsp;</p><p><br /></p><p>Note: If logistics_channel_id is 90021, 90025 or 90026, service_code will be appended,&nbsp;e.g., Entrega Turbo - M1020.</p>
	PaymentMethod                         string                 `json:"payment_method"`                            // [Required] The payment method that the buyer selected to pay for the order. Applicable values: See Data Definition- Payment Methods.
	EstimatedShippingFee                  float64                `json:"estimated_shipping_fee"`                    // [Required] The estimated shipping fee is an estimation calculated by Shopee based on specific logistics courier's standard.
	MessageToSeller                       string                 `json:"message_to_seller"`                         // [Required] Return by default. Message to seller.
	CreateTime                            int64                  `json:"create_time"`                               // [Required] Return by default. Timestamp that indicates the date and time that the order was created.
	UpdateTime                            int64                  `json:"update_time"`                               // [Required] Return by default. Timestamp that indicates the last time that there was a change in value of order, such as order status changed from 'Paid' to 'Completed'.
	DaysToShip                            int64                  `json:"days_to_ship"`                              // [Required] Return by default. Shipping preparation time set by the seller when listing item on Shopee.
	ShipByDate                            int64                  `json:"ship_by_date"`                              // [Required] Return by default. The deadline to ship out the parcel.
	BuyerUserId                           int64                  `json:"buyer_user_id"`                             // [Required] <p>The user id of buyer of this order, will be empty if it is a non-integrated order in TW region.</p>
	BuyerUsername                         string                 `json:"buyer_username"`                            // [Required] <p>The name of buyer, will be masked as "****" if it is a non-integrated order in TW region.</p>
	RecipientAddress                      *OrderRecipientAddress `json:"recipient_address"`                         // [Required] <p>This object contains detailed breakdown for the recipient address.<br />Different parameters might be masked according to each market and kind of seller.<br /><br />For TW region integrated channel orders will be all masked as "****". More details may refer the announcement.<br /></p>
	ActualShippingFee                     float64                `json:"actual_shipping_fee"`                       // [Required] The actual shipping fee of the order if available from external logistics partners.
	GoodsToDeclare                        bool                   `json:"goods_to_declare"`                          // [Required] Only work for cross-border order.This value indicates whether the order contains goods that are required to declare at customs. "T" means true and it will mark as "T" on the shipping label; "F" means false and it will mark as "P" on the shipping label. This value is accurate ONLY AFTER the order trackingNo is generated, please capture this value AFTER your retrieve the trackingNo.
	Note                                  string                 `json:"note"`                                      // [Required] The note seller made for own reference.
	NoteUpdateTime                        int64                  `json:"note_update_time"`                          // [Required] Update time for the note.
	ItemList                              []OrderItem            `json:"item_list"`                                 // [Required] This object contains the detailed breakdown for the result of this API call.
	PayTime                               int64                  `json:"pay_time"`                                  // [Required] The time when the order status is updated from UNPAID to PAID. This value is NULL when order is not paid yet.
	Dropshipper                           string                 `json:"dropshipper"`                               // [Required] For Indonesia orders only. The name of the dropshipper.
	DropshipperPhone                      string                 `json:"dropshipper_phone"`                         // [Required] The phone number of dropshipper, could be empty.
	SplitUp                               bool                   `json:"split_up"`                                  // [Required] To indicate whether this order is split to fullfil order(forder) level. Call GetForderInfo if it's "true".
	BuyerCancelReason                     string                 `json:"buyer_cancel_reason"`                       // [Required] Cancel reason from buyer, could be empty.
	CancelBy                              string                 `json:"cancel_by"`                                 // [Required] Could be one of buyer, seller, system or Ops.
	CancelReason                          string                 `json:"cancel_reason"`                             // [Required] Use this field to get reason for buyer, seller, and system cancellation.
	ActualShippingFeeConfirmed            bool                   `json:"actual_shipping_fee_confirmed"`             // [Required] Use this filed to judge whether the actual_shipping_fee is confirmed.
	BuyerCpfId                            string                 `json:"buyer_cpf_id"`                              // [Required] Buyer's CPF number for taxation and invoice purposes. Only for Brazil order.
	FulfillmentFlag                       string                 `json:"fulfillment_flag"`                          // [Required] Use this field to indicate the order is fulfilled by shopee or seller. Applicable values: fulfilled_by_shopee, fulfilled_by_cb_seller, fulfilled_by_local_seller.
	PickupDoneTime                        int64                  `json:"pickup_done_time"`                          // [Required] The timestamp when pickup is done.
	PackageList                           []OrderPackage         `json:"package_list"`                              // [Required] The list of package under an order
	InvoiceData                           *InvoiceData           `json:"invoice_data"`                              // [Required] The invoice data of the order.
	CheckoutShippingCarrier               string                 `json:"checkout_shipping_carrier"`                 // [Required] For non masking order, the logistics service provider that the buyer selected for the order to deliver items.  For masking order, the logistics service type that the buyer selected for the order to deliver items.
	ReverseShippingFee                    float64                `json:"reverse_shipping_fee"`                      // [Required] Shopee charges the reverse shipping fee for the returned order.The value of this field will be non-negative.
	OrderChargeableWeightGram             int64                  `json:"order_chargeable_weight_gram"`              // [Required] display weight used to calculate ASF for this order
	PrescriptionCheckStatus               int64                  `json:"prescription_check_status"`                 // [Required] <p>Prescription check status. For ID, PH whitelisted sellers, the applicable values:</p><p>0: NONE</p><p>1: PASSED</p><p>2: FAILED</p><p>For TH whitelisted sellers, the applicable values:</p><p>0: NONE</p><p>1: PASSED<strong></strong></p>
	PharmacistName                        string                 `json:"pharmacist_name"`                           // [Required] <p>Name of the Pharmacist for Prescription Order.</p>
	PrescriptionImages                    []string               `json:"prescription_images"`                       // [Required] <p>Return prescription images of this order, only for ID and PH whitelist sellers.</p><p><br /></p><p>Please add the prefix to review:</p><p>for ID:&nbsp;<a href="https://cf.shopee.co.id/file/+prescription_image" target="_blank" style="font-size:14px;">https://cf.shopee.co.id/file/+prescription_image</a></p><p>for PH:<a href="https://cf.shopee.co.id/file/+prescription_image" target="_blank" style="font-size:14px;">https://cf.shopee.ph/file/+prescription_image</a></p>
	PrescriptionApprovalTime              int64                  `json:"prescription_approval_time"`                // [Required] <p>Time of when the prescription is approved.</p>
	PrescriptionRejectionTime             int64                  `json:"prescription_rejection_time"`               // [Required] <p>Time of when the prescription is rejected.</p>
	PrescriptionRejectReason              string                 `json:"prescription_reject_reason"`                // [Required] <p>Return the reason why a prescription is rejected. If there is no rejection reason, return empty.Only for ID and PH whitelist sellers</p>
	IsBuyerShopCollection                 bool                   `json:"is_buyer_shop_collection"`                  // [Required] <p>To indicate if this order is buyer self collection at store order</p>
	BuyerProofOfCollection                []string               `json:"buyer_proof_of_collection"`                 // [Required] <p>The image url of the buyer self collection at the store.</p>
	EdtFrom                               int64                  `json:"edt_from"`                                  // [Required] <p>Earliest estimated delivery date of orders (only available for BR region)<br /></p>
	EdtTo                                 int64                  `json:"edt_to"`                                    // [Required] <p>Latest estimated delivery time of orders (only available for BR region)<br /></p>
	BookingSn                             string                 `json:"booking_sn"`                                // [Required] <p>Return by default. Shopee's unique identifier for a booking.</p><p>Only returned for advance fulfilment matched order only.</p>
	AdvancePackage                        bool                   `json:"advance_package"`                           // [Required] <p>Indicate whether order will be fulfilled using advance fulfilment stock or not. If value is true, order will be matched with a booking and seller should not arrange shipment.</p>
	ReturnRequestDueDate                  int64                  `json:"return_request_due_date"`                   // [Required] <p>This field represents the deadline for buyers to initiate returns and refunds after order is completed.</p><p><br /></p><p>The “return_request_due_date” response parameter will be returned if the requested order meets&nbsp;<b>ALL&nbsp;the conditions</b>&nbsp;below:</p><p>- The status of the order is COMPLETED</p><p>- The return refund eligibility of the order is true</p><p><br /></p><p>If you have any questions related to the function of "returns and refunds after order is completed," please refer to the following link:&nbsp;https://seller.shopee.tw/edu/article/18474</p>
	PaymentInfo                           []PaymentInfo          `json:"payment_info"`                              // [Required] <p>[Only for BR] List of payment information, to follow&nbsp;<a href="https://drive.google.com/file/d/1VfqlbmXr3XR6BkpKOPUbLCgjiqvsBbLd/view?usp=sharing" target="_blank">NT 2025.001</a>&nbsp;(BR government invoice rules).</p>
	HotListingOrder                       bool                   `json:"hot_listing_order"`                         // [Required] <p>[Only for PH,TH,VN,MY,BR,TW] True if the order includes hot listing item.</p>
	IsInternational                       bool                   `json:"is_international"`                          // [Required] <p>[Only for BR] Indicate if the order is SIP order. This field will only be returned if international_label is included in response_optional_field in the request.</p>
	CanFullCancelOrder                    bool                   `json:"can_full_cancel_order"`                     // [Required] <p>Indicates whether the order can be full cancelled:&nbsp;</p><p>-&nbsp;If this value is true, seller can cancel the entire order</p><p>- If the value is false, full order cancellation is not available for the order</p>
	CanPartialCancelOrder                 bool                   `json:"can_partial_cancel_order"`                  // [Required] <p>Indicates whether the order is eligible for partial cancellation. This value is determined by both the system eligibility check and the buyer’s out-of-stock handling preference.&nbsp;</p><p>- If this value is true, seller can cancel selected out-of-stock item quantities while continuing to fulfill the remaining items.&nbsp;</p><p>- If this value is false, partial cancellation is not allowed.</p>
	BuyerPreferenceForPartialCancellation int64                  `json:"buyer_preference_for_partial_cancellation"` // [Required] <p>Indicates the buyer’s preference for handling out-of-stock items in the order. Applicable values:</p><p>0 = Ship Available Items Only (The buyer allows the seller to cancel unavailable items and continue shipping the remaining available items)</p><p>1 = Cancel The Entire Order (The buyer does not allow partial cancellation. If any item is unavailable, the seller should cancel the entire order instead)</p>
	AffiliateSampleType                   int64                  `json:"affiliate_sample_type"`                     // [Required] <p>Indicates that this order is a refundable sample  order. Applicable values:</p><p>0 = Order is not a refundable sample order</p><p>1 = Order is a refundable sample order</p>
}
type GetOrderListRequest struct {
	Cursor                    *string      `json:"cursor,omitempty" url:"cursor,omitempty"`                                             // [Optional] <p>Specifies the starting entry of data to return in the current call. The default is empty. If the data is more than one page, the offset can be some entry to start the next call.</p>
	LogisticsChannelId        *int64       `json:"logistics_channel_id,omitempty" url:"logistics_channel_id,omitempty"`                 // [Optional] <p>The identity of logistic channel. Valid only for BR.<br /></p>
	OrderStatus               *OrderStatus `json:"order_status,omitempty" url:"order_status,omitempty"`                                 // [Optional] The order_status filter for retriveing orders and each one only every request. Available value: UNPAID/READY_TO_SHIP/PROCESSED/SHIPPED/COMPLETED/IN_CANCEL/CANCELLED/INVOICE_PENDING
	PageSize                  int64        `json:"page_size" url:"page_size"`                                                           // [Required] Each result set is returned as a page of entries. Use the "page_size" filters to control the maximum number of entries to retrieve per page (i.e., per call). This integer value is used to specify the maximum number of entries to return in a single "page" of data.The limit of page_size if between 1 and 100.
	RequestOrderStatusPending *bool        `json:"request_order_status_pending,omitempty" url:"request_order_status_pending,omitempty"` // [Optional] <p>Compatible parameter during migration period, send True will let API support PENDING status, send False or don’t send will fallback to old logic.<br /></p>
	ResponseOptionalFields    *string      `json:"response_optional_fields,omitempty" url:"response_optional_fields,omitempty"`         // [Optional] Optional fields in response. Available value: order_status.
	TimeFrom                  int64        `json:"time_from" url:"time_from"`                                                           // [Required] The time_from and time_to fields specify a date range for retrieving orders (based on the time_range_field). The time_from field is the starting date range. The maximum date range that may be specified with the time_from and time_to fields is 15 days.
	TimeRangeField            string       `json:"time_range_field" url:"time_range_field"`                                             // [Required] The kind of time_from and time_to. Available value: create_time, update_time.
	TimeTo                    int64        `json:"time_to" url:"time_to"`                                                               // [Required] The time_from and time_to fields specify a date range for retrieving orders (based on the time_range_field). The time_from field is the starting date range. The maximum date range that may be specified with the time_from and time_to fields is 15 days.
}
type GetOrderListResponse struct {
	BaseResponse                          // Common response fields
	Response     GetOrderListResponseData `json:"response"` // Response data
}
type GetOrderListResponseData struct {
	More       bool                            `json:"more"`        // [Required] This is to indicate whether the order list is more than one page. If this value is true, you may want to continue to check next page to retrieve orders.
	OrderList  []GetOrderListResponseDataOrder `json:"order_list"`  // [Required]
	NextCursor string                          `json:"next_cursor"` // [Required] If  more is true, you should pass the next_cursor in the next request as cursor. The value of next_cursor will be empty string when more is false.
}
type GetOrderListResponseDataOrder struct {
	OrderSn     string      `json:"order_sn"`     // [Required]  Shopee's unique identifier for an order.
	OrderStatus OrderStatus `json:"order_status"` // [Required] The order_status filter for retriveing orders and each one only every request. Available value: UNPAID/READY_TO_SHIP/PROCESSED/SHIPPED/COMPLETED/IN_CANCEL/CANCELLED
	BookingSn   string      `json:"booking_sn"`   // [Required] <p>Return by default. Shopee's unique identifier for a booking.</p><p>Only returned for advance fulfilment matched order only.<br /></p>
}
type GetPackageDetailRequest struct {
	PackageNumberList string `json:"package_number_list" url:"package_number_list"` // [Required] <p>The set of package_number. If there are multiple package_number, you need to use English comma to connect them. limit [1,50]</p>
}
type GetPackageDetailResponse struct {
	BaseResponse                              // Common response fields
	Response     GetPackageDetailResponseData `json:"response"` // Response data
}
type GetPackageDetailResponseData struct {
	PackageList []ResponseDataPackage `json:"package_list"` // [Required] <p>The list of packages.</p>
}
type GetPendingBuyerInvoiceOrderListRequest struct {
	Cursor   *string `json:"cursor,omitempty" url:"cursor,omitempty"` // [Optional] Specifies the starting entry of data to return in the current call. Default is "". If data is more than one page, the offset can be some entry to start next call.
	PageSize int64   `json:"page_size" url:"page_size"`               // [Required] Each result set is returned as a page of entries. Use the "page_size" filters to control the maximum number of entries to retrieve per page (i.e., per call). This integer value is used to specify the maximum number of entries to return in a single "page" of data.The limit of page_size if between 1 and 100.
}
type GetPendingBuyerInvoiceOrderListResponse struct {
	BaseResponse                                             // Common response fields
	Response     GetPendingBuyerInvoiceOrderListResponseData `json:"response"` // Response data
}
type GetPendingBuyerInvoiceOrderListResponseData struct {
	More       bool      `json:"more"`        // [Required] This is to indicate whether the order list is more than one page. If this value is true, you may want to continue to check next page to retrieve orders.
	NextCursor string    `json:"next_cursor"` // [Required] If more is true, you should pass the next_cursor in the next request as cursor. The value of next_cursor will be empty string when more is false.
	OrderList  []Queries `json:"order_list"`  // [Required]
}
type GetShipmentListRequest struct {
	Cursor   *string `json:"cursor,omitempty" url:"cursor,omitempty"` // [Optional] Specifies the starting entry of data to return in the current call. Default is "". If data is more than one page, the offset can be some entry to start next call.
	PageSize int64   `json:"page_size" url:"page_size"`               // [Required] Each result set is returned as a page of entries. Use the "page_size" filters to control the maximum number of entries to retrieve per page (i.e., per call). This integer value is used to specify the maximum number of entries to return in a single "page" of data.The limit of page_size if between 1 and 100.
}
type GetShipmentListResponse struct {
	BaseResponse                             // Common response fields
	Response     GetShipmentListResponseData `json:"response"` // Response data
}
type GetShipmentListResponseData struct {
	OrderList  []Order `json:"order_list"`  // [Required] The list of  shipment orders
	More       bool    `json:"more"`        // [Required] This is to indicate whether the order list is more than one page. If this value is true, you may want to continue to check next page to retrieve orders.
	NextCursor string  `json:"next_cursor"` // [Required] If more is true, you should pass the next_cursor in the next request as cursor. The value of next_cursor will be empty string when more is false.
}
type GetWarehouseFilterConfigResponse struct {
	BaseResponse                                      // Common response fields
	Response     GetWarehouseFilterConfigResponseData `json:"response"` // Response data
}
type GetWarehouseFilterConfigResponseData struct {
	WarehouseFilters []WarehouseFilters `json:"warehouse_filters"` // [Required]
}
type HandleBuyerCancellationRequest struct {
	Operation string `json:"operation"` // [Required] The operation you want to handle.Avaiable value: ACCEPT, REJECT
	OrderSn   string `json:"order_sn"`  // [Required] Shopee's unique identifier for an order.
}
type HandleBuyerCancellationResponse struct {
	BaseResponse                                     // Common response fields
	Response     HandleBuyerCancellationResponseData `json:"response"` // Response data
}
type HandleBuyerCancellationResponseData struct {
	UpdateTime int64 `json:"update_time"` // [Required] The time when the order is updated.
}
type HandlePrescriptionCheckRequest struct {
	FreeText         *string        `json:"free_text,omitempty"`          // [Optional] <p>The reason for rejecting the prescription. Only usable when the reject_reason_code = 5.</p>
	IsApproved       bool           `json:"is_approved"`                  // [Required] <p>Approve or reject the prescription. Available values: TRUE, FALSE.<br /></p>
	Items            []RequestItems `json:"items,omitempty"`              // [Optional] <p>The list of invalid items that make the prescription get rejected<br /></p>
	OrderSn          string         `json:"order_sn"`                     // [Required] <p>Shopee's unique identifier for an order.<br /></p>
	PharmacistName   *string        `json:"pharmacist_name,omitempty"`    // [Optional] <p>Full name of the pharmacist.&nbsp;Required for PH and ID Prescription Orders.</p>
	RejectReasonCode *int64         `json:"reject_reason_code,omitempty"` // [Optional] <p>Reject reason code. Available values:&nbsp;<br /></p><p>1 = Invalid Prescription (counterfeit/incorrect format)</p><p>2 = Incorrect Dosage</p><p>3 = No Prescription</p><p>4 = Unclear Image</p><p>5 = Free Text</p>
}
type HandlePrescriptionCheckResponse struct {
	BaseResponse                                     // Common response fields
	Response     HandlePrescriptionCheckResponseData `json:"response"` // Response data
}
type HandlePrescriptionCheckResponseData struct {
	IsSuccess bool `json:"is_success"` // [Required] <p>This is to indicate whether the request has been executed successfully.<br /></p>
}
type HouseholdAddressBreakdown struct {
	HouseholdRegion          string `json:"household_region"`           // [Required] <p>Region of the household address.</p>
	HouseholdState           string `json:"household_state"`            // [Required] <p>State of the household address.</p>
	HouseholdCity            string `json:"household_city"`             // [Required] <p>City of the household address.</p>
	HouseholdProvince        string `json:"household_province"`         // [Required] <p>Province of the household address.</p>
	HouseholdDistrict        string `json:"household_district"`         // [Required] <p>District of the household address.</p>
	HouseholdTown            string `json:"household_town"`             // [Required] <p>Town of the household address.</p>
	HouseholdBarangay        string `json:"household_barangay"`         // [Required] <p>Barangay of the household address.</p>
	HouseholdPostcode        string `json:"household_postcode"`         // [Required] <p>Postal code of the household address.</p>
	HouseholdDetailedAddress string `json:"household_detailed_address"` // [Required] <p>Detailed street address of the household.</p>
	HouseholdAdditionalInfo  string `json:"household_additional_info"`  // [Required] <p>Additional address information provided by the buyer.</p>
	HouseholdFullAddress     string `json:"household_full_address"`     // [Required] <p>Full formatted household address.</p>
}
type InvoiceData struct {
	Number             string  `json:"number"`               // [Required] The number of the invoice.
	SeriesNumber       string  `json:"series_number"`        // [Required] The series number of the invoice.
	AccessKey          string  `json:"access_key"`           // [Required] The access key of the invoice.
	IssueDate          int64   `json:"issue_date"`           // [Required] The issue date of the invoice.
	TotalValue         float64 `json:"total_value"`          // [Required] The total value of the invoice.
	ProductsTotalValue float64 `json:"products_total_value"` // [Required] The products total value of the invoice.
	TaxCode            string  `json:"tax_code"`             // [Required] The tax code for the invoice.
	Status             string  `json:"status"`               // [Required] <p>The invoice statuses should be:</p><p>- valid (The invoice sent is valid)</p><p>-&nbsp;pending (System is waiting to receive the invoice)</p>
	PendingReason      string  `json:"pending_reason"`       // [Required] <p>It's the failed reason if status is pending<strong></strong></p>
}
type InvoiceDetail struct {
	Name                      string                     `json:"name"`                        // [Required] <p>Buyer name (has value when invoice_type is personal, household, or company)<br /></p><p>- VN, TH, PH only<br /></p>
	Email                     string                     `json:"email"`                       // [Required] <p>Buyer email address&nbsp;(has value when invoice_type is personal and household)<br /></p><p>- VN, TH, PH only<br /></p>
	PhoneNumber               string                     `json:"phone_number"`                // [Required] <p>Buyer phone number<br /></p><p>- TH only<br /></p>
	TaxId                     string                     `json:"tax_id"`                      // [Required] <p>has value when invoice_type is personal and household.&nbsp;<br /></p><p>- VN, TH, PH only<br /></p>
	Address                   string                     `json:"address"`                     // [Required] <p>Buyer address in format "Street &amp; number, city, zipcode, any additional info provided by buyer"&nbsp;(has value when invoice_type is personal and household)<br /></p><p>- PH, VN only<br /></p>
	IdCardAddress             string                     `json:"id_card_address"`             // [Required] <p>Same function as the address, only having a different field name for TH.Buyer address in format "Street &amp; number, city, zipcode, any additional info provided by buyer"&nbsp;(only has value when invoice_type is personal).<br /></p>
	AddressBreakdown          *AddressBreakdown          `json:"address_breakdown"`           // [Required] <p>Buyer address breakdown.<br /></p><p>- TH, PH only<br /></p>
	CompanyHeadOffice         string                     `json:"company_head_office"`         // [Required] <p>- return value for TH only&nbsp;(only has value when invoice_type is company)<br /></p>
	CompanyName               string                     `json:"company_name"`                // [Required] <p>- Only return value when invoice type is company</p><p>- VN, TH, PH only<br /></p>
	CompanyBranchName         string                     `json:"company_branch_name"`         // [Required] <p>- Only return value when invoice type is company</p><p>- TH only<br /></p>
	CompanyBranchId           string                     `json:"company_branch_id"`           // [Required] <p>- Only return value when invoice type is company</p><p>- TH only<br /></p>
	CompanyType               string                     `json:"company_type"`                // [Required] <p>- Only return value when invoice type is company</p><p>- TH only<br /></p>
	CompanyEmail              string                     `json:"company_email"`               // [Required] <p>- Only return value when invoice type is company</p><p>- VN, TH, PH only<br /></p>
	CompanyTaxId              string                     `json:"company_tax_id"`              // [Required] <p>- Only return value when invoice type is company</p><p>- VN, TH, PH only<br /></p>
	CompanyAddress            string                     `json:"company_address"`             // [Required] <p>Buyer address in format "Street &amp; number,city, zipcode, any additional info provided by buyer"&nbsp;(only has value when invoice_type is company)<br /></p><p>- VN, TH only<br /></p>
	CompanyAddressBreakdown   *CompanyAddressBreakdown   `json:"company_address_breakdown"`   // [Required] <p>Company address breakdown</p><p>- PH, TH only<br /></p>
	HouseholdAddressBreakdown *HouseholdAddressBreakdown `json:"household_address_breakdown"` // [Required] <p>Household address breakdown</p><p>-Only for VN</p>
	NationalId                string                     `json:"national_id"`                 // [Required] <p>National ID information provided by the buyer.</p><p>- Only return value when invoice_type is personal<br />- VN only</p>
}
type InvoiceInfo struct {
	OrderSn       string         `json:"order_sn"`       // [Required] <p>Shopee's unique identifier for an order.<br /></p>
	InvoiceType   string         `json:"invoice_type"`   // [Required] <p>Type of invoice requested: {1: personal, 2: company, 3: household}.<br /></p>
	InvoiceDetail *InvoiceDetail `json:"invoice_detail"` // [Required] <p>Invoice info submitted by buyer. Might be masked, e.g. A****b, depending on order status.<br /></p>
	Error         string         `json:"error"`          // [Required] <p>Error in retrieving the receipt setting of a particular order.<br /></p>
	IsRequested   bool           `json:"is_requested"`   // [Required] <p>To identify order with and without buyer request, applicable to PL.<br /></p>
}
type InvoicePending struct {
	Status        string `json:"status"`         // [Required] <p>The invoice statuses should be:</p><p>- valid (The invoice sent is valid)</p><p>-&nbsp;pending (System is waiting to receive the invoice)</p>
	PendingReason string `json:"pending_reason"` // [Required] <p>It's the failed reason if status is pending</p>
}
type Order struct {
	OrderSn       string `json:"order_sn"`       // [Required] Shopee's unique identifier for an order.
	PackageNumber string `json:"package_number"` // [Required] Shopee's unique identifier for the package under an order
}
type OrderItem struct {
	ItemId                            int64          `json:"item_id"`                                // [Required] Shopee's unique identifier for an item.
	ItemName                          string         `json:"item_name"`                              // [Required] The name of the item.
	ItemSku                           string         `json:"item_sku"`                               // [Required]  A item SKU (stock keeping unit) is an identifier defined by a seller, sometimes called parent SKU. Item SKU can be assigned to an item in Shopee Listings.
	ModelId                           int64          `json:"model_id"`                               // [Required] ID of the model that belongs to the same item.
	ModelName                         string         `json:"model_name"`                             // [Required] Name of the model that belongs to the same item. A seller can offer models of the same item. For example, the seller could create a fixed-priced listing for a t-shirt design and offer the shirt in different colors and sizes. In this case, each color and size combination is a separate model. Each model can have a different quantity and price.
	ModelSku                          string         `json:"model_sku"`                              // [Required] A model SKU (stock keeping unit) is an identifier defined by a seller. It is only intended for the seller's use. Many sellers assign a SKU to an item of a specific type, size, and color, which are models of one item in Shopee Listings.
	ModelQuantityPurchased            int64          `json:"model_quantity_purchased"`               // [Required] The number of identical items purchased at the same time by the same buyer from one listing/item.
	ModelOriginalPrice                float64        `json:"model_original_price"`                   // [Required] The original price of the item in the listing currency.
	ModelDiscountedPrice              float64        `json:"model_discounted_price"`                 // [Required] The after-discount price of the item in the listing currency. If there is no discount, this value will be same as that of model_original_price. In case of bundle deal item, this value will return 0 as by design bundle deal discount will not be breakdown to item/model level. Due to technical restriction, the value will return the price before bundle deal if we don't configure it to 0. Please call GetEscrowDetails if you want to calculate item-level discounted price for bundle deal item.
	Wholesale                         bool           `json:"wholesale"`                              // [Required] This value indicates whether buyer buy the order item in wholesale price.
	Weight                            float64        `json:"weight"`                                 // [Required] The weight of the item
	AddOnDeal                         bool           `json:"add_on_deal"`                            // [Required] To indicate if this item belongs to an addon deal.
	MainItem                          bool           `json:"main_item"`                              // [Required] To indicate if this item is main item or sub item. True means main item, false means sub item.
	AddOnDealId                       int64          `json:"add_on_deal_id"`                         // [Required] A unique ID to distinguish groups of items in Cart, and Order. (e.g. AddOnDeal)
	PromotionType                     string         `json:"promotion_type"`                         // [Required] <p>Available type：product_promotion, flash_sale, bundle_deal, add_on_deal_main, add_on_deal_sub.</p><p><br /></p><p>For items which attend multiple promotions will only show one promotion, the order of priority is:&nbsp;</p><p>bundle_deal &gt; add_on_deal_main &gt; add_on_deal_sub &gt; product_promotion &gt;flash_sale</p>
	PromotionId                       int64          `json:"promotion_id"`                           // [Required] The ID of the promotion.
	OrderItemId                       int64          `json:"order_item_id"`                          // [Required] The identify of order item.
	LineItemId                        int64          `json:"line_item_id"`                           // [Required] <p>The identity of order item. In case the order item is a bundle deal, this value will be unique to distinguish the order item</p>
	PromotionGroupId                  int64          `json:"promotion_group_id"`                     // [Required] The identify of product promotion.
	ImageInfo                         *ItemImageInfo `json:"image_info"`                             // [Required] Image info of the product.
	ProductLocationId                 StringSlice    `json:"product_location_id"`                    // [Required] The fulfilment warehouse ID(s) of the items in the order. (Multi-Warehouse sellers only)
	IsPrescriptionItem                bool           `json:"is_prescription_item"`                   // [Required] <p>To indicate if this item is prescription item. Only for PH, TH, ID local shop.<br /></p>
	ErrorInFetchingIsPrescriptionItem bool           `json:"error_in_fetching_is_prescription_item"` // [Required] <p>To indicate if there was an error when validating whether this item is prescription. Default false. If is_prescription_item=false and this field is true, the item's prescription status is uncertain (label service call failed). Only for TH, PH, ID local shop.</p>
	ConsultationId                    string         `json:"consultation_id"`                        // [Required] <p>An identifier of teleconsultation session which buyer did to order this item. Empty if item is not ordered through teleconsultation session</p>
	IsB2cOwnedItem                    bool           `json:"is_b2c_owned_item"`                      // [Required] <p>determine if item is B2C_shop_item</p><p>It should be `<b>is_b2c_shop_item</b>` but it was a bug from dev. Then now it's <b>is_b2c_owned_item</b></p>
	PromotionList                     []Promotion    `json:"promotion_list"`                         // [Required]
	HotListingItem                    bool           `json:"hot_listing_item"`                       // [Required] <p>[Only for PH,TH,VN,MY,BR,TW] True if the item is hot listing.</p>
	ActiveQty                         int64          `json:"active_qty"`                             // [Required] <p>The quantity of the item model that remains active in the order and is still expected to be fulfilled.</p>
	CancelRequestedQty                int64          `json:"cancel_requested_qty"`                   // [Required] <p>The quantity of the item model that is currently under a cancellation request but has not yet reached the final cancelled status.</p>
	CancelledQty                      int64          `json:"cancelled_qty"`                          // [Required] <p>The quantity of the item model that has already been successfully cancelled.</p>
	ReturnRequestedQty                int64          `json:"return_requested_qty"`                   // [Required] <p>The quantity of the item model that is currently under a return/refund request.</p>
	ReturnedQty                       int64          `json:"returned_qty"`                           // [Required] <p>The quantity of the item model that has already been successfully returned through the return/refund process.</p>
}
type OrderPackage struct {
	PackageNumber          string          `json:"package_number"`           // [Required] Shopee's unique identifier for the package under an order.
	LogisticsStatus        LogisticsStatus `json:"logistics_status"`         // [Required] The Shopee logistics status for the order. Applicable values: See Data Definition-LogisticsStatus.
	LogisticsChannelId     int64           `json:"logistics_channel_id"`     // [Required] <p>The identity of logistic channel.</p>
	ShippingCarrier        string          `json:"shipping_carrier"`         // [Required] <p>The logistics service provider that the buyer selected for the order to deliver items.&nbsp;</p><p><br /></p><p>Note: If logistics_channel_id is 90021, 90025 or 90026, service_code will be appended,&nbsp;e.g., Entrega Turbo - M1020.</p>
	AllowSelfDesignAwb     bool            `json:"allow_self_design_awb"`    // [Required] <p>To indicate whether the package allows for self-designed AWB, if allow_self_design_awb returns false, it means that the package does not allow for self-designed AWB and only the system-AWB can be used.</p>
	ItemList               []PackageItem   `json:"item_list"`                // [Required] The lis of items.
	ParcelChargeableWeight int64           `json:"parcel_chargeable_weight"` // [Required] display weight used to calculate ASF for this parcel
	GroupShipmentId        int64           `json:"group_shipment_id"`        // [Required] <p>The common identifier for multiple orders combined in the same parcel.<br /></p>
	VirtualContactNumber   string          `json:"virtual_contact_number"`   // [Required] <p>[Only for TW non-integrated channel] The virtual phone number to contact the recipient.<br /></p>
	PackageQueryNumber     string          `json:"package_query_number"`     // [Required] <p>[Only for TW non-integrated channel] The query number used in virtual phone number calls to contact the recipient of this package.<br /></p>
	SortingGroup           string          `json:"sorting_group"`            // [Required] <p>[Only for TW 30029 channel] This field indicate the sorting group value of the package. This field is only available for logistics_channel_id = 30029 and after the package has been arranged for shipment.</p>
}
type OrderRecipientAddress struct {
	Name        string       `json:"name"`         // [Required] <p>Recipient's name for the address.<br /></p>
	Phone       string       `json:"phone"`        // [Required] <p>Recipient's phone number input when order was placed.</p><p>[Only for TW non-integrated channel] Will return "****" when the "virtual_contact_number" is available</p>
	Town        string       `json:"town"`         // [Required] <p>The town of the recipient's address. Whether there is a town will depend on the region and/or country.<br /></p>
	District    string       `json:"district"`     // [Required] <p>The district of the recipient's address. Whether there is a district will depend on the region and/or country.<br /></p>
	City        string       `json:"city"`         // [Required] <p>The city of the recipient's address. Whether there is a city will depend on the region and/or country.<br /></p>
	State       string       `json:"state"`        // [Required] <p>The state/province of the recipient's address. Whether there is a state/province will depend on the region and/or country.<br /></p>
	Region      string       `json:"region"`       // [Required] <p>The two-digit code representing the region of the Recipient.<br /></p>
	Zipcode     string       `json:"zipcode"`      // [Required] <p>Recipient's postal code.<br /></p>
	FullAddress string       `json:"full_address"` // [Required] <p>The full address of the recipient, including country, state, even street, and etc.<br /></p>
	Geolocation *Geolocation `json:"geolocation"`  // [Required] <p>Geolocation info. Only available for logistics_channel_id 90026.</p>
}
type PackageItem struct {
	ItemId            int64       `json:"item_id"`             // [Required] Shopee's unique identifier for an item.
	ModelId           int64       `json:"model_id"`            // [Required] Shopee's unique identifier for a model.
	ModelQuantity     int64       `json:"model_quantity"`      // [Required] <p>The number of identical items/variations purchased at the same time by the same buyer from one listing/item.</p>
	OrderItemId       int64       `json:"order_item_id"`       // [Required] <p>The identify of order item. For items in one same bundle deal promotion, the order_item_id should share the same id, such as 1,2. For items not in bundle deal promotion, the order_item_id should be the same as item_id.<br /></p>
	PromotionGroupId  int64       `json:"promotion_group_id"`  // [Required] <p>The identify of product promotion.<br /></p>
	ProductLocationId StringSlice `json:"product_location_id"` // [Required] <p>The warehouse ID of the item.<br /></p>
}
type Packages struct {
	OrderSn            string `json:"order_sn"`             // [Required] <p>Shopee's unique identifier for an order.</p>
	PackageNumber      string `json:"package_number"`       // [Required] <p>Shopee's unique identifier for the package under an order</p>
	LogisticsChannelId int64  `json:"logistics_channel_id"` // [Required] <p>The identity of logistic channel. </p><p></p>
	ProductLocationId  string `json:"product_location_id"`  // [Required] Just use this field to pass the next step of Mass ArrangeShipment<br />
	SortingGroup       string `json:"sorting_group"`        // [Required] <p>[Only for TW 30029 channel] This field indicate the sorting group value of the package. This field is only available for logistics_channel_id = 30029 and after the package has been arranged for shipment.</p>
	IsShipmentArranged bool   `json:"is_shipment_arranged"` // [Required] <p>Only effective when the package's logistics_status/fulfillment_status is LOGISTICS_READY.&nbsp;</p><p><br /></p><p>This parameter further distinguishes between two scenarios:</p><p>- true: Package shipment has been arranged (Seller has processed shipment, system is generating tracking number, not yet updated to LOGISTICS_REQUEST_CREATED, no duplicate action needed)</p><p>- false: Package awaiting shipment arrangement (Seller hasn't processed shipment yet, shipping arrangement required)</p>
}
type PartialCancelItem struct {
	ItemId           int64 `json:"item_id"`            // [Required] Shopee's unique identifier for an item.
	ModelId          int64 `json:"model_id"`           // [Required] Shopee's unique identifier for a model.
	OrderItemId      int64 `json:"order_item_id"`      // [Required] <p>The identify of order item. For items in one same bundle deal promotion, the order_item_id should share the same id, such as 1,2. For items not in bundle deal promotion, the order_item_id should be the same as item_id.</p>
	PromotionGroupId int64 `json:"promotion_group_id"` // [Required] <p>The identify of product promotion. For items in one same add on deal promotion, the promotion_group_id should share the same id. For items not in add on deal promotion, the promotion_group_id should be 0. And the data is from group_id of shopee.orders.GetOrderDetails.</p>
	ModelQuantity    int64 `json:"model_quantity"`     // [Required] <p>The number of identical items put in the package.<br /></p>
}
type PaymentInfo struct {
	PaymentMethod            string  `json:"payment_method"`             // [Required] <p>[Only for BR] Payment method used in the order, such as Credit Card, Debit Card, Pix, etc.</p>
	PaymentProcessorRegister string  `json:"payment_processor_register"` // [Required] <p>[Only for BR] CNPJ of the payment processor handling the transaction.</p>
	CardBrand                string  `json:"card_brand"`                 // [Required] <p>[Only for BR] Card brand for credit or debit transactions, such as VISA, MASTER, etc. Empty string for Pix payments.</p>
	TransactionId            string  `json:"transaction_id"`             // [Required] <p>[Only for BR] Payment authorization code generated by the bank or payment processor to validate the transaction.</p>
	PaymentAmount            float64 `json:"payment_amount"`             // [Required] <p>[Only for BR] Amount paid by the corresponding payment method.<br /></p>
}
type Promotion struct {
	PromotionType string `json:"promotion_type"` // [Required] <p>Indicates the type of item or package level promotion applied to a product. Each item can be associated with at most one item promotion and one package promotion at a time.<br /><br />Item Promotions:</p><p>low_price_promotion</p><p>deep_discount</p><p>platform_sale</p><p>seller_discount</p><p>flash_sale</p><p>wholesale</p><p>welcome_package_free_gift</p><p>brand_flash_sale</p><p>in_shop_flash_sale</p><p>synced_promo</p><p>platform_streaming_price</p><p>seller_streaming_price</p><p>exclusive_streamer_price</p><p>price_bidding_with_rebate</p><p>price_bidding_without_rebate</p><p>seller_advisor_price</p><p>selling_price</p><p>settlement_price</p><p>campaign_settlement_price</p><p>local_sip_settlement_price</p><p>platform_exclusive_price</p><p>seller_exclusive_price</p><p>seller_member_exclusive_sku</p><p>item_price</p><p>order_sync_price</p><p><br />Package Promotions:<br />bundle_deal<br />add_on_deal_main<br />add_on_deal_sub</p>
	PromotionId   int64  `json:"promotion_id"`   // [Required] <p>Represents the unique identifier of a specific promotion applied to an item. Each promotion_id corresponds to a distinct promotion rule or campaign, defined under a particular promotion_type. The value is expressed in a numeric string format.</p>
}
type Queries struct {
	OrderSn string `json:"order_sn"` // [Required] Shopee's unique identifier for an order.
}
type RecipientAddress struct {
	Name        string `json:"name"`         // [Required] <p>Recipient's name for the address.<br /></p>
	Phone       string `json:"phone"`        // [Required] <p>Recipient's phone number input when booking was placed.<br /></p>
	Town        string `json:"town"`         // [Required] <p>The town of the recipient's address. Whether there is a town will depend on the region and/or country.<br /></p>
	District    string `json:"district"`     // [Required] <p>The district of the recipient's address. Whether there is a district will depend on the region and/or country.<br /></p>
	City        string `json:"city"`         // [Required] <p>The city of the recipient's address. Whether there is a city will depend on the region and/or country.<br /></p>
	State       string `json:"state"`        // [Required] <p>The state/province of the recipient's address. Whether there is a state/province will depend on the region and/or country.<br /></p>
	Region      string `json:"region"`       // [Required] <p>The two-digit code representing the region of the Recipient.<br /></p>
	Zipcode     string `json:"zipcode"`      // [Required] <p>Recipient's postal code.<br /></p>
	FullAddress string `json:"full_address"` // [Required] <p>The full address of the recipient, including country, state, even street, and etc.<br /></p>
}
type RequestId struct {
	RequestId []int64 `json:"request_id"` // [Required] <p>A list of integers representing the request IDs to be queried.</p>
}
type RequestItems struct {
	ItemId  int64 `json:"item_id"`  // [Required] <p>Shopee's unique identifier for an item.<br /></p>
	ModelId int64 `json:"model_id"` // [Required] <p>Shopee's unique identifier for a model of an item.<br /></p>
	GroupId int64 `json:"group_id"` // [Required] <p>The identify of product promotion. For items in one same add on deal promotion, the group_id should share the same id. For items not in add on deal promotion, the group_id should be 0. And the data is from group_id of shopee.orders.GetOrderDetails.<br /></p>
}
type RequestSubItem struct {
	ItemId  int64 `json:"item_id"`  // [Required] Id of item.
	ModelId int64 `json:"model_id"` // [Required] Id of the model that belongs to the same item.
}
type ResponseDataBooking struct {
	BookingSn        string            `json:"booking_sn"`        // [Required] <p>Return by default. Shopee's unique identifier for a booking.<br /></p>
	OrderSn          string            `json:"order_sn"`          // [Required] <p>Shopee's unique identifier for an order. Only return if booking_status is MATCHED.<br /></p>
	Region           string            `json:"region"`            // [Required] <p>Return by default. The two-digit code representing the region where the booking was made.<br /></p>
	BookingStatus    BookingStatus     `json:"booking_status"`    // [Required] <p>Return by default.&nbsp;Enumerated type that defines the current status of the booking. Available value:&nbsp;READY_TO_SHIP/PROCESSED/SHIPPED/CANCELLED/MATCHED<br /></p>
	MatchStatus      string            `json:"match_status"`      // [Required] <p>MATCH_PENDING/MATCH_SUCCESSFUL/MATCH_FAILED<br /></p>
	ShippingCarrier  string            `json:"shipping_carrier"`  // [Required] <p>The logistics service provider&nbsp;that will deliver the booking.<br /></p>
	CreateTime       int64             `json:"create_time"`       // [Required] <p>Return by default. Timestamp that indicates the date and time that the booking was created.<br /></p>
	UpdateTime       int64             `json:"update_time"`       // [Required] <p>Return by default. Timestamp that indicates the last time that there was a change in value of booking, such as booking status changed from 'Processed' to 'Shipped'.<br /></p>
	ShipByDate       int64             `json:"ship_by_date"`      // [Required] <p>Return by default. The deadline to ship out the parcel.<br /></p>
	RecipientAddress *RecipientAddress `json:"recipient_address"` // [Required] <p>This object contains detailed breakdown for the recipient address.<br /></p>
	ItemList         []BookingItem     `json:"item_list"`         // [Required] <p>This object contains the detailed breakdown for the result of this API call.<br /></p>
	Dropshipper      string            `json:"dropshipper"`       // [Required] <p>For Indonesia bookings only. The name of the dropshipper.<br /></p>
	DropshipperPhone string            `json:"dropshipper_phone"` // [Required] <p>The phone number of dropshipper, could be empty.<br /></p>
	CancelBy         string            `json:"cancel_by"`         // [Required] <p>Could be one of buyer, seller, system or Ops.<br /></p>
	CancelReason     string            `json:"cancel_reason"`     // [Required] <p>Use this field to get reason for buyer, seller, and system cancellation.<br /></p>
	FulfillmentFlag  string            `json:"fulfillment_flag"`  // [Required] <p>Use this field to indicate the booking is fulfilled by shopee or seller. Applicable values:&nbsp;fulfilled_by_shopee, fulfilled_by_cb_seller, fulfilled_by_local_seller.<br /></p>
	PickupDoneTime   int64             `json:"pickup_done_time"`  // [Required] <p>The timestamp when pickup is done.<br /></p>
}
type ResponseDataPackage struct {
	OrderSn                               string                    `json:"order_sn"`                                  // [Required] <p>Shopee's unique identifier for an order.</p>
	PackageNumber                         string                    `json:"package_number"`                            // [Required] <p>Shopee's unique identifier for the package under an order.</p>
	FulfillmentStatus                     string                    `json:"fulfillment_status"`                        // [Required] <p>The Shopee fulfillment status for the package. Applicable values: See V2.0 Data Definition - PackageFulfillmentStatus.</p>
	UpdateTime                            int64                     `json:"update_time"`                               // [Required] <p>Timestamp that indicates the last time that there was a change in value of package.</p>
	LogisticsChannelId                    int64                     `json:"logistics_channel_id"`                      // [Required] <p>The identity of logistic channel.</p>
	ShippingCarrier                       string                    `json:"shipping_carrier"`                          // [Required] <p>The logistics service provider that the buyer selected for the package to deliver items.&nbsp;</p><p><br /></p><p>Note: If logistics_channel_id is 90021, 90025 or 90026, service_code will be appended, e.g., Entrega Turbo - M1020.</p>
	AllowSelfDesignAwb                    bool                      `json:"allow_self_design_awb"`                     // [Required] <p>To indicate whether the package allows for self-designed AWB, if allow_self_design_awb returns false, it means that the package does not allow for self-designed AWB and only the system-AWB can be used.</p>
	DaysToShip                            int64                     `json:"days_to_ship"`                              // [Required] <p>Shipping preparation time set by the seller when listing item on Shopee.<br /></p>
	ShipByDate                            int64                     `json:"ship_by_date"`                              // [Required] <p>The deadline to ship out the package.</p>
	PendingTerms                          []string                  `json:"pending_terms"`                             // [Required] <p>The list of pending terms. Applicable values:</p><p>- SYSTEM_PENDING: Under Shopee internal processing.</p><p>- KYC_PENDING: Under KYC checking (TW CB order only).</p><p>- ARRANGE_SHIPMENT_PENDING: Temporarily held due to 3PL capacity constraints.</p>
	PendingDescription                    []string                  `json:"pending_description"`                       // [Required] <p>The value of this field is the description of pending reason corresponding with pending terms. Applicable values:&nbsp;</p><p>- For SYSTEM_PENDING: Order is being processed by Shopee.</p><p>- For KYC_PENDING: Order is pending buyer TW KYC pre-authorization.</p><p>- For ARRANGE_SHIPMENT_PENDING: Allocating delivery resources due to high order volume. Label print will be available within 4 days after buyer paid.</p>
	TrackingNumber                        string                    `json:"tracking_number"`                           // [Required] <p>The tracking number of this package.<br /></p>
	TrackingNumberExpirationDate          int64                     `json:"tracking_number_expiration_date"`           // [Required] <p>[TW only] Tracking number expiration date</p>
	PickupDoneTime                        int64                     `json:"pickup_done_time"`                          // [Required] <p>The timestamp when pickup is done.<br /></p>
	IsSplitUp                             bool                      `json:"is_split_up"`                               // [Required] <p>To indicate whether this parcel is split.<br /></p>
	ItemList                              []ResponseDataPackageItem `json:"item_list"`                                 // [Required] <p>The lis of items in the package.</p>
	RecipientAddress                      *OrderRecipientAddress    `json:"recipient_address"`                         // [Required] <p>This object contains detailed breakdown for the recipient address.<br />Different parameters might be masked according to each market and kind of seller.<br /><br />For TW region integrated channel orders will be all masked as "****". More details may refer the announcement.</p>
	ParcelChargeableWeightGram            int64                     `json:"parcel_chargeable_weight_gram"`             // [Required] <p>display weight used to calculate ASF for this parcel<br /></p>
	GroupShipmentId                       int64                     `json:"group_shipment_id"`                         // [Required] <p>The common identifier for multiple orders combined in the same parcel.<br /></p>
	VirtualContactNumber                  string                    `json:"virtual_contact_number"`                    // [Required] <p>[Only for TW non-integrated channel] The virtual phone number to contact the recipient.<br /></p>
	PackageQueryNumber                    string                    `json:"package_query_number"`                      // [Required] <p>[Only for TW non-integrated channel] The query number used in virtual phone number calls to contact the recipient of this package.<br /></p>
	SortingGroup                          string                    `json:"sorting_group"`                             // [Required] <p>[Only for TW 30029 channel] This field indicate the sorting group value of the package. This field is only available for logistics_channel_id = 30029 and after the package has been arranged for shipment.</p>
	IsShipmentArranged                    bool                      `json:"is_shipment_arranged"`                      // [Required] <p>Only effective when the package's logistics_status/fulfillment_status is LOGISTICS_READY.&nbsp;</p><p><br /></p><p>This parameter further distinguishes between two scenarios:</p><p>- true: Package shipment has been arranged (Seller has processed shipment, system is generating tracking number, not yet updated to LOGISTICS_REQUEST_CREATED, no duplicate action needed)</p><p>- false: Package awaiting shipment arrangement (Seller hasn't processed shipment yet, shipping arrangement required)</p>
	StatusInfoTag                         *StatusInfoTag            `json:"status_info_tag"`                           // [Required] <p>Package shipping urgency tag information.</p>
	CanSplitOrder                         bool                      `json:"can_split_order"`                           // [Required] <p>This field indicates whether this order can be split into multiple packages for separate shipment.</p><p>- true: Support splitting, can call v2.order.split_order to execute</p><p>- false: Does not support splitting</p>
	CanUnsplitOrder                       bool                      `json:"can_unsplit_order"`                         // [Required] <p>This field indicates whether this order can be unsplit.</p><p>- true: Support unsplitting, can call v2.order.unsplit_order to execute</p><p>- false: Does not support unsplitting</p>
	IsPreOrder                            bool                      `json:"is_pre_order"`                              // [Required] <p>This field indicates whether this order is a pre-order.</p><p>- true: Pre-order</p><p>- false: Non pre-order</p>
	PharmacistName                        string                    `json:"pharmacist_name"`                           // [Required] <p>Name of the Pharmacist for Prescription Order.</p>
	PrescriptionImages                    []string                  `json:"prescription_images"`                       // [Required] <p>Return prescription images of this order, only for ID and PH whitelist sellers.</p><p><br /></p><p>Please add the prefix to review:</p><p>for ID:&nbsp;https://cf.shopee.co.id/file/+prescription_image</p><p>for PH: https://cf.shopee.ph/file/+prescription_image</p>
	PrescriptionApprovalTime              int64                     `json:"prescription_approval_time"`                // [Required] <p>Time of when the prescription is approved.</p>
	PrescriptionRejectionTime             int64                     `json:"prescription_rejection_time"`               // [Required] <p>Time of when the prescription is rejected.</p>
	IsBuyerShopCollection                 bool                      `json:"is_buyer_shop_collection"`                  // [Required] <p>To indicate if this order is buyer self collection at store order.</p>
	BuyerProofOfCollection                []string                  `json:"buyer_proof_of_collection"`                 // [Required] <p>The image url of the proof for buyer self collection at the store.</p>
	PreparationEndTime                    int64                     `json:"preparation_end_time"`                      // [Required] <p>The system-calculated deadline for package preparation. When the package fulfillment_status/logistics_status changes to "LOGISTICS_READY", the system calculates this time based on the "Preparation Time" configured for the logistics channel of this package.&nbsp;</p><p><br /></p><p>Notes:&nbsp;</p><p>1) Only effective for logistics channels that have Auto Call Driver enabled and Preparation Time configured.</p><p>2) Seller needs to complete packing and waybill printing before this time to ensure the package is ready when the driver arrives.</p><p>3) When this time is reached, the system will automatically arrange shipment and trigger driver dispatch:</p><p>- If driver calling is successful, the package fulfillment_status/logistics_status will change from “LOGISTICS_READY” to “LOGISTICS_REQUEST_CREATED”.</p><p>- If driver calling fails, the package fulfillment_status/logistics_status will remain unchanged, and the seller must arrange shipment manually.</p>
	DriverInfo                            *DriverInfo               `json:"driver_info"`                               // [Required] <p>After the driver is successfully called, the driver's information will be returned.</p><p><br /></p><p>Note: Data availability depends on the specific 3PL provider; certain fields may be omitted due to provider policies, PII restrictions, or data unavailability.</p>
	CanFullCancelOrder                    bool                      `json:"can_full_cancel_order"`                     // [Required] <p>Indicates whether the order can be full cancelled:&nbsp;</p><p>-&nbsp;If this value is true, seller can cancel the entire order</p><p>- If the value is false, full order cancellation is not available for the order</p>
	CanPartialCancelOrder                 bool                      `json:"can_partial_cancel_order"`                  // [Required] <p>Indicates whether the order is eligible for partial cancellation. This value is determined by both the system eligibility check and the buyer’s out-of-stock handling preference.&nbsp;</p><p>- If this value is true, seller can cancel selected out-of-stock item quantities while continuing to fulfill the remaining items.&nbsp;</p><p>- If this value is false, partial cancellation is not allowed.</p>
	BuyerPreferenceForPartialCancellation int64                     `json:"buyer_preference_for_partial_cancellation"` // [Required] <p>Indicates the buyer’s preference for handling out-of-stock items in the order. Applicable values:</p><p>0 = Ship Available Items Only (The buyer allows the seller to cancel unavailable items and continue shipping the remaining available items)</p><p>1 = Cancel The Entire Order (The buyer does not allow partial cancellation. If any item is unavailable, the seller should cancel the entire order instead)</p>
	InvoicePending                        *InvoicePending           `json:"invoice_pending"`                           // [Required]
}
type ResponseDataPackageItem struct {
	ItemId                            int64       `json:"item_id"`                                // [Required] <p>Shopee's unique identifier for an item.</p>
	ModelId                           int64       `json:"model_id"`                               // [Required] <p>Shopee's unique identifier for a model.</p>
	ItemSku                           string      `json:"item_sku"`                               // [Required] <p>A item SKU (stock keeping unit) is an identifier defined by a seller, sometimes called parent SKU. Item SKU can be assigned to an item in Shopee Listings.<br /></p>
	ModelSku                          string      `json:"model_sku"`                              // [Required] <p>ID of the model that belongs to the same item.<br /></p>
	ModelQuantity                     int64       `json:"model_quantity"`                         // [Required] <p>The number of identical items/variations purchased at the same time by the same buyer from one listing/item.</p>
	OrderItemId                       int64       `json:"order_item_id"`                          // [Required] <p>The identify of order item. For items in one same bundle deal promotion, the order_item_id should share the same id, such as 1,2. For items not in bundle deal promotion, the order_item_id should be the same as item_id.</p>
	PromotionGroupId                  int64       `json:"promotion_group_id"`                     // [Required] <p>The identify of product promotion.</p>
	ProductLocationId                 StringSlice `json:"product_location_id"`                    // [Required] <p>The warehouse ID of the item.</p>
	ConsultationId                    string      `json:"consultation_id"`                        // [Required] <p>An identifier of teleconsultation session which buyer did to order this item. Empty if item is not ordered through teleconsultation session</p>
	IsPrescriptionItem                bool        `json:"is_prescription_item"`                   // [Required] <p>To indicate if this item is a prescription item. Default false. Only for PH, TH, ID whitelist shops.</p>
	ErrorInFetchingIsPrescriptionItem bool        `json:"error_in_fetching_is_prescription_item"` // [Required] <p>To indicate if there was an error when validating whether this item is prescription. Default false. If is_prescription_item=false and this field is true, the item's prescription status is uncertain (label service call failed).&nbsp;Only for PH, TH, ID whitelist shops.</p>
	PrescriptionCheckStatus           int64       `json:"prescription_check_status"`              // [Required] <p>Prescription check status. For ID, PH whitelisted sellers, the applicable values:</p><p>0: NONE</p><p>1: PASSED</p><p>2: FAILED</p><p>For TH whitelisted sellers, the applicable values:</p><p>0: NONE</p><p>1: PASSED</p>
	PrescriptionRejectReason          string      `json:"prescription_reject_reason"`             // [Required] <p>Return the reason why a prescription is rejected. If no rejection reason, return empty. Only for ID and PH whitelist sellers.</p>
}
type ResponseDataPagination struct {
	TotalCount int64  `json:"total_count"` // [Required] <p>Total orders can be returned with your query<br /></p>
	NextCursor string `json:"next_cursor"` // [Required] <p>if packages is not empty or length of packages &lt;= page_size. You should pass the next_cursor in the next request as page_sentinel. <br /></p>
	More       bool   `json:"more"`        // [Required] <p>To indicate, it's a the last page or not<br /></p>
}
type ResponseDataSort struct {
	SortType int64 `json:"sort_type"` // [Required] <p>As same as request param</p><p></p>
	IsAsc    bool  `json:"is_asc"`    // [Required] <p>As same as request param</p><p></p>
}
type SearchPackageListRequest struct {
	Filter     *Filter     `json:"filter,omitempty"` // [Optional]
	Pagination *Pagination `json:"pagination"`       // [Required]
	Sort       *Sort       `json:"sort,omitempty"`   // [Optional]
}
type SearchPackageListResponse struct {
	BaseResponse                               // Common response fields
	Response     SearchPackageListResponseData `json:"response"`         // Response data
	Mesage       string                        `json:"mesage,omitempty"` // <p>Indicate error details if hit error. Empty if no error happened.</p>
}
type SearchPackageListResponseData struct {
	PackagesList []Packages              `json:"packages_list"` // [Required]
	Pagination   *ResponseDataPagination `json:"pagination"`    // [Required]
	Sort         *ResponseDataSort       `json:"sort"`          // [Required] <p>As same as request param</p>
}
type SetNoteRequest struct {
	Note    string `json:"note"`     // [Required] The note seller add for reference.
	OrderSn string `json:"order_sn"` // [Required] Shopee's unique identifier for an order.
}
type SetNoteResponse struct {
	BaseResponse // Common response fields
}
type Sort struct {
	SortType  *int64 `json:"sort_type,omitempty"` // [Optional] <p>Use this field to specify which field to use to sort the returned package list. Available values:&nbsp;</p><p>1: ShipByDate&nbsp;&nbsp;<br /></p><p>2: CreateDate<br /></p><p>3: ConfirmedDate<br /></p><p><br /></p><p><font color="#c24f4a">Default value = 1 (ShipByDate)</font><br /></p>
	Ascending *bool  `json:"ascending,omitempty"` // [Optional] <p>Use this field to specify whether the returned package list is sorted in ascending or descending sort_type.</p><p><font color="#c24f4a"><br /></font></p><p><font color="#c24f4a">Default value = true</font></p>
}
type SplitOrderRequest struct {
	OrderSn     string                     `json:"order_sn"`     // [Required] Shopee's unique identifier for an order.
	PackageList []SplitOrderRequestPackage `json:"package_list"` // [Required] <p>The list of packages that you want to split.&nbsp;</p><p><br /></p><p>Note:&nbsp;</p><p>- Orders that include installation services cannot be split by quantity.</p><p>- When splitting the order, must contain all items in the order in one request.</p><p>- You can split the order into 30 parcels at most in TW and 5 parcels at most in other regions.</p>
}
type SplitOrderRequestPackage struct {
	ItemList []PartialCancelItem `json:"item_list"` // [Required] <p>The list of items under the same package.</p>
}
type SplitOrderResponse struct {
	BaseResponse                        // Common response fields
	Response     SplitOrderResponseData `json:"response"` // Response data
}
type SplitOrderResponseData struct {
	OrderSn     string                          `json:"order_sn"`     // [Required] Shopee's unique identifier for an order.
	PackageList []SplitOrderResponseDataPackage `json:"package_list"` // [Required] The list of package under this order you have split.
}
type SplitOrderResponseDataPackage struct {
	PackageNumber string              `json:"package_number"` // [Required] Shopee's unique identifier for the package under an order.
	ItemList      []PartialCancelItem `json:"item_list"`      // [Required] The list of items under this package.
}
type StatusInfoTag struct {
	TagId     int64 `json:"tag_id"`    // [Required] <p>Shipping urgency tag type, applicable values below:</p><p>0: No tag</p><p>1: Will be cancelled within 1 day</p><p>2: Must ship before the specified timestamp</p><p>3: Shipment delayed</p><p>4: Must ship within the current hour</p><p>5: Will be cancelled at the specified timestamp</p>
	Timestamp int64 `json:"timestamp"` // [Required] <p>When tag_id is 2 or 5, returns specific timestamp (e.g., cancel time, shipment deadline); otherwise returns 0.</p>
}
type UnsplitOrderRequest struct {
	OrderSn string `json:"order_sn"` // [Required] Shopee's unique identifier for an order.
}
type UnsplitOrderResponse struct {
	BaseResponse // Common response fields
}
type UploadInvoiceDocRequest struct {
	File     interface{} `json:"file"`      // [Required] <p>invoice file.&nbsp;File size limit to 1MB.</p>
	FileType int64       `json:"file_type"` // [Required] <p>the type of invoice file. 1:pdf 2.jpeg 3.png. 4.xml</p>
	OrderSn  string      `json:"order_sn"`  // [Required] Shopee's unique identifier for an order.
}
type UploadInvoiceDocResponse struct {
	BaseResponse // Common response fields
}
type WarehouseFilters struct {
	WarehouseName     string `json:"warehouse_name"`      // [Required] <p>The warehouse name filled in when creating the warehouse address.</p>
	WarehouseType     int64  `json:"warehouse_type"`      // [Required] <p>Type of warehouse. Applicable values:</p><p>- 1: Local Warehouse</p><p>- 2: CB Warehouse</p>
	ProductLocationId string `json:"product_location_id"` // [Required] <p>Location identifier for stocks. Different location_ids represent that your addresses are in different item stocks.</p>
	AddressId         int64  `json:"address_id"`          // [Required] <p>Identity of address.</p>
	Address           string `json:"address"`             // [Required] <p>Detail address of your warehouse.</p>
}
