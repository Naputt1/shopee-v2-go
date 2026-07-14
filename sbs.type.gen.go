package goshopee

type AdjustQty struct {
	AdjustTotal     int64 `json:"adjust_total"`      // [Required] <p>"Total number of SKU changes during the selected time period."</p>
	AdjustLostFound int64 `json:"adjust_lost_found"` // [Required] <p>"Total quantity of lost or recovered items during the selected time period."</p>
	AdjustTransWhs  int64 `json:"adjust_trans_whs"`  // [Required] <p>Total quantity of transfer orders created by the warehouse during the selected time period<br /></p>
}
type BoundWhs struct {
	WhsRegion string `json:"whs_region"` // [Required] <p>the warehouse region bound with the shop</p>
	WhsIds    string `json:"whs_ids"`    // [Required] <p>the warehouse id bound with the shop</p>
}
type EndQty struct {
	EndOnHandTotal int64 `json:"end_on_hand_total"` // [Required] <p>Total inventory at the end time.<br /></p>
	EndSellable    int64 `json:"end_sellable"`      // [Required]
	EndReserved    int64 `json:"end_reserved"`      // [Required]
	EndUnsellable  int64 `json:"end_unsellable"`    // [Required]
}
type GetBoundWhsInfoResponse struct {
	BaseResponse                             // Common response fields
	Response     GetBoundWhsInfoResponseData `json:"response"` // Response data
}
type GetBoundWhsInfoResponseData struct {
	List []GetBoundWhsInfoResponseDataList `json:"list"` // [Required]
}
type GetBoundWhsInfoResponseDataList struct {
	ShopId   int64      `json:"shop_id"`   // [Required]
	BoundWhs []BoundWhs `json:"bound_whs"` // [Required]
}
type GetCurrentInventoryRequest struct {
	CategoryId             *int64  `json:"category_id,omitempty"`              // [Optional] <p>Category id. Here you need to call the&nbsp;get_category&nbsp;API to retrieve the first-tier category_id.<br /></p>
	InboundPendingApproval *int64  `json:"inbound_pending_approval,omitempty"` // [Optional] <p>Blank-All；0-No；1-Yes</p>
	Keyword                *string `json:"keyword,omitempty"`                  // [Optional] <p>Bind Value and Search_type</p>
	NotMovingTag           *int64  `json:"not_moving_tag,omitempty"`           // [Optional] <p>Blank-All；0-No；1-Yes</p>
	PageNo                 *int64  `json:"page_no,omitempty"`                  // [Optional] <p>Specifies the page number of data to return in the current call. Starting from 1. if data is more than one page, the page_no can be some entry to start next call. If empty, the default value is 1.</p>
	PageSize               *int64  `json:"page_size,omitempty"`                // [Optional] <p>Each result set is returned as a page of entries. Use the "page_size" filters to control the maximum number of entries to retrieve per page (i.e., per call), and the "page_no" to start next call. This integer value is used to specify the maximum number of entries to return in a single "page" of data.</p><p>If empty, the default value is 10. The value should be between 1 and 100.</p>
	ProductsWithInventory  *int64  `json:"products_with_inventory,omitempty"`  // [Optional] <p>Blank-All；0-No；1-Yes</p>
	SearchType             *int64  `json:"search_type,omitempty"`              // [Optional] <p>0-All data；1-Product Name；2-SKU ID；3-Variations；4-Item ID</p>
	StockLevels            *string `json:"stock_levels,omitempty"`             // [Optional] <p>1-Low Stock &amp; No Sellable stock; 2-Low Stock &amp; To replenish; 3-Low Stock &amp; Replenished; 4-Excess<br /></p>
	WhsIds                 *string `json:"whs_ids,omitempty"`                  // [Optional] <p>Whs ID list, comma-separated</p>
	WhsRegion              string  `json:"whs_region"`                         // [Required] <p>The warehouse region you want to query, can only query one region in a request<br />Optional value: BR、CN、ID、MY、MX、TH、TW、PH、VN、SG<br /><br /><b>If do not pass, will get error "block by gateway due to invalid cid"</b></p>
}
type GetCurrentInventoryResponse struct {
	BaseResponse                                 // Common response fields
	Response     GetCurrentInventoryResponseData `json:"response"` // Response data
}
type GetCurrentInventoryResponseData struct {
	ItemList []GetCurrentInventoryResponseDataItem `json:"item_list"` // [Required] <p>Data list of item sku</p>
}
type GetCurrentInventoryResponseDataItem struct {
	WarehouseItemId string `json:"warehouse_item_id"` // [Required] <p>Warehouse item id; To indicate an unique item in a warehouse<br /><br />one warehouse item id can match with multiple shop_item_id<br /><br />For Global Item,&nbsp;warehouse_item_id=Global Item id<br />For Local Item, shop_item_id=item_id</p>
	ItemName        string `json:"item_name"`         // [Required] <p>item name</p>
	ItemImage       string `json:"item_image"`        // [Required] <p>item image</p>
	SkuList         []Sku  `json:"sku_list"`          // [Required] <p>Data list of mtsku</p>
}
type GetExpiryReportRequest struct {
	CategoryIdL1 *int64  `json:"category_id_l1,omitempty"` // [Optional] <p>Only <b>Level 1 Category</b> can be filtered</p>
	ExpiryStatus *string `json:"expiry_status,omitempty"`  // [Optional] <p>0-Expired，2-Expiring，4-expiry_blocked，5-damaged，6-normal。Multiple selections allowed, separated by commas.<br /></p>
	ItemId       *string `json:"item_id,omitempty"`        // [Optional]
	ItemName     *string `json:"item_name,omitempty"`      // [Optional]
	PageNo       *int64  `json:"page_no,omitempty"`        // [Optional] <p>Specifies the page number of data to return in the current call. Starting from 1. if data is more than one page, the page_no can be some entry to start next call. If empty, the default value is 1.<br /></p>
	PageSize     *int64  `json:"page_size,omitempty"`      // [Optional] <p>Each result set is returned as a page of entries. Use the "page_size" filters to control the maximum number of entries to retrieve per page (i.e., per call), and the "page_no" to start next call. This integer value is used to specify the maximum number of entries to return in a single "page" of data.</p><p>If empty, the default value is 10. The value should be between 1 and 40.</p>
	SkuId        *string `json:"sku_id,omitempty"`         // [Optional]
	Variation    *string `json:"variation,omitempty"`      // [Optional]
	WhsIds       *string `json:"whs_ids,omitempty"`        // [Optional]
	WhsRegion    string  `json:"whs_region"`               // [Required] <p>Num value: BR、CN、ID、MY、MX、TH、TW、PH、VN、SG<br /><br /></p><p><b>If do not pass, will get error "block by gateway due to invalid cid"</b></p>
}
type GetExpiryReportResponse struct {
	BaseResponse                             // Common response fields
	Response     GetExpiryReportResponseData `json:"response"` // Response data
}
type GetExpiryReportResponseData struct {
	ItemList []GetExpiryReportResponseDataItem `json:"item_list"` // [Required]
}
type GetExpiryReportResponseDataItem struct {
	WarehouseItemId string    `json:"warehouse_item_id"` // [Required] <p>Warehouse item id; To indicate an unique item in a warehouse<br />one warehouse item id can match with multiple shop_item_id<br /><br />For Global Item,&nbsp;warehouse_item_id=Global Item id<br />For Local Item, shop_item_id=item_id</p>
	ItemName        string    `json:"item_name"`         // [Required]
	ItemImage       string    `json:"item_image"`        // [Required]
	SkuList         []ItemSku `json:"sku_list"`          // [Required]
}
type GetStockAgingRequest struct {
	AgingStorageTag  *int64  `json:"aging_storage_tag,omitempty"`  // [Optional] <p>0-false；1-true<br /></p>
	CategoryId       *int64  `json:"category_id,omitempty"`        // [Optional] <p>L1-level product category ID. You need to call the get_category API to obtain the first-level category_id<br /></p>
	ExcessStorageTag *int64  `json:"excess_storage_tag,omitempty"` // [Optional] <p>0-false；1-true<br /></p>
	Keyword          *string `json:"keyword,omitempty"`            // [Optional] <p>bound with search_type<br /></p>
	PageNo           *int64  `json:"page_no,omitempty"`            // [Optional] <p>Specifies the page number of data to return in the current call. Starting from 1. if data is more than one page, the page_no can be some entry to start next call. If empty, the default value is 1.<br /></p>
	PageSize         *int64  `json:"page_size,omitempty"`          // [Optional] <p>Each result set is returned as a page of entries. Use the "page_size" filters to control the maximum number of entries to retrieve per page (i.e., per call), and the "page_no" to start next call. This integer value is used to specify the maximum number of entries to return in a single "page" of data.</p><p>If empty, the default value is 10. The value should be between 1 and 100.</p>
	SearchType       *int64  `json:"search_type,omitempty"`        // [Optional] <p>1-Product Name；2-SKU ID；3-Variations；4-Item ID<br /></p>
	WhsIds           *string `json:"whs_ids,omitempty"`            // [Optional] <p>split by comma<br /></p>
	WhsRegion        string  `json:"whs_region"`                   // [Required] <p>BR、CN、ID、MY、MX、TH、TW、PH、VN、SG<br /></p><p><b>If do not pass, will get error "block by gateway due to invalid cid"</b></p>
}
type GetStockAgingResponse struct {
	BaseResponse                           // Common response fields
	Response     GetStockAgingResponseData `json:"response"` // Response data
}
type GetStockAgingResponseData struct {
	ItemList []GetStockAgingResponseDataItem `json:"item_list"` // [Required] <p>Data list of item sku<br /></p>
}
type GetStockAgingResponseDataItem struct {
	WarehouseItemId string                `json:"warehouse_item_id"` // [Required] <p>Warehouse item id; To indicate an unique item in a warehouse<br />one warehouse item id can match with multiple shop_item_id<br /><br /><b>For Global Item,&nbsp;warehouse_item_id=Global Item id<br />For Local Item, shop_item_id=item_id</b></p>
	ItemName        string                `json:"item_name"`         // [Required] <p>item name<br /></p>
	ItemImage       string                `json:"item_image"`        // [Required] <p>item image<br /></p>
	SkuList         []ResponseDataItemSku `json:"sku_list"`          // [Required] <p>Data list of mtsku<br /></p>
}
type GetStockMovementRequest struct {
	CategoryIdL1 *int64  `json:"category_id_l1,omitempty"` // [Optional] <p>L1-level category_id. You need to call the get_category API to retrieve the first-level category_id.<br /></p>
	EndTime      string  `json:"end_time"`                 // [Required] <p>End date in YYYY-MM-DD format. Only data within the past 1 year can be queried, and the time range must not exceed 90 days.<br /></p>
	ItemId       *string `json:"item_id,omitempty"`        // [Optional]
	ItemName     *string `json:"item_name,omitempty"`      // [Optional] <p>Product Name Filter<br /></p>
	PageNo       *int64  `json:"page_no,omitempty"`        // [Optional] <p>Specifies the page number of data to return in the current call. Starting from 1. if data is more than one page, the page_no can be some entry to start next call. If empty, the default value is 1.<br /></p>
	PageSize     *int64  `json:"page_size,omitempty"`      // [Optional] <p>Each result set is returned as a page of entries. Use the "page_size" filters to control the maximum number of entries to retrieve per page (i.e., per call), and the "page_no" to start next call. This integer value is used to specify the maximum number of entries to return in a single "page" of data.</p><p>If empty, the default value is 10. The value should be between 1 and 20.<br /></p>
	SkuId        *string `json:"sku_id,omitempty"`         // [Optional]
	StartTime    string  `json:"start_time"`               // [Required] <p>Start date in YYYY-MM-DD format. Only data within the past 1 year can be queried, and the time range must not exceed 90 days.<br /></p>
	Variation    *string `json:"variation,omitempty"`      // [Optional]
	WhsIds       *string `json:"whs_ids,omitempty"`        // [Optional] <p>Multiple warehouse_id values should be separated by commas.<br /></p>
	WhsRegion    string  `json:"whs_region"`               // [Required] <p>Warehouse Region. Enum values: BR, CN, ID, MY, MX, TH, TW, PH, VN, SG<br /></p><p><b>If do not pass, will get error "block by gateway due to invalid cid"</b></p>
}
type GetStockMovementResponse struct {
	BaseResponse                              // Common response fields
	Response     GetStockMovementResponseData `json:"response"` // Response data
}
type GetStockMovementResponseData struct {
	Total        int64                              `json:"total"`          // [Required]
	StartTime    string                             `json:"start_time"`     // [Required]
	EndTime      string                             `json:"end_time"`       // [Required]
	QueryEndTime string                             `json:"query_end_time"` // [Required]
	ItemList     []GetStockMovementResponseDataItem `json:"item_list"`      // [Required] <p>Data list of item sku<br /></p>
}
type GetStockMovementResponseDataItem struct {
	WarehouseItemId string                                `json:"warehouse_item_id"` // [Required] <p>Warehouse item id; To indicate an unique item in a warehouse<br /><br />one warehouse item id can match with multiple shop_item_id<br /></p>
	ItemName        string                                `json:"item_name"`         // [Required] <p>item name<br /></p>
	ItemImage       string                                `json:"item_image"`        // [Required] <p>item image<br /></p>
	SkuList         []GetStockMovementResponseDataItemSku `json:"sku_list"`          // [Required] <p>Data list of mtsku<br /></p>
}
type GetStockMovementResponseDataItemSku struct {
	MtskuId            string                   `json:"mtsku_id"`             // [Required] <p>mtsku id<br /></p>
	ModelId            string                   `json:"model_id"`             // [Required] <p>model sku id<br /></p>
	Variation          string                   `json:"variation"`            // [Required]
	FulfillMappingMode int64                    `json:"fulfill_mapping_mode"` // [Required] <p>0-Null；1-Bundle SKU；2-Parent SKU<br /></p>
	Barcode            string                   `json:"barcode"`              // [Required]
	WhsList            []ResponseDataItemSkuWhs `json:"whs_list"`             // [Required] <p>Info of whs<br /></p>
	StartQty           *StartQty                `json:"start_qty"`            // [Required] <p>Inventory information at the start time.<br /></p>
	EndQty             *EndQty                  `json:"end_qty"`              // [Required]
	InboundQty         *InboundQty              `json:"inbound_qty"`          // [Required] <p>Inbound information during the selected time period<br /></p>
	OutboundQty        *OutboundQty             `json:"outbound_qty"`         // [Required]
	AdjustQty          *AdjustQty               `json:"adjust_qty"`           // [Required] <p>"SKU change information during the selected time period."</p><p> </p><p><br /></p>
	ShopSkuList        []SkuShopSku             `json:"shop_sku_list"`        // [Required]
}
type InboundQty struct {
	InboundTotal    int64 `json:"inbound_total"`    // [Required] <p>Total inbound quantity during the selected time period<br /></p>
	InboundMy       int64 `json:"inbound_my"`       // [Required] <p>Total merchant procurement quantity during the selected time period.<br /></p>
	InboundReturned int64 `json:"inbound_returned"` // [Required] <p>Total number of SKUs returned by buyers and received into the warehouse during the selected time period.<br /></p>
}
type ItemSku struct {
	MtskuId            string       `json:"mtsku_id"`             // [Required] <p>Unique ID for a warehouse SKU<br />"warehouse_item_id"_"warehouse_model_id"</p>
	ModelId            string       `json:"model_id"`             // [Required] <p>Warehouse model SKU ID</p><p>For CB global items, this is equal to the global model_id.</p><p><br /></p><p>For local items, it differs from model_id; use shop_model_id to match the model_id</p>
	FulfillMappingMode int64        `json:"fulfill_mapping_mode"` // [Required] <p>0-Null；1-Bundle SKU；2-Parent SKU<br /></p>
	Variation          string       `json:"variation"`            // [Required]
	ShopSkuList        []SkuShopSku `json:"shop_sku_list"`        // [Required]
	WhsList            []SkuWhs     `json:"whs_list"`             // [Required]
}
type ItemSkuWhs struct {
	WhsId              string `json:"whs_id"`                 // [Required] <p>Whs id<br /></p>
	QtyOfStockAgeOne   int64  `json:"qty_of_stock_age_one"`   // [Required] <p>0-30Days<br /></p>
	QtyOfStockAgeTwo   int64  `json:"qty_of_stock_age_two"`   // [Required] <p>31-60Days<br /></p>
	QtyOfStockAgeThree int64  `json:"qty_of_stock_age_three"` // [Required] <p>61-90Days<br /></p>
	QtyOfStockAgeFour  int64  `json:"qty_of_stock_age_four"`  // [Required] <p>91-120Days<br /></p>
	QtyOfStockAgeFive  int64  `json:"qty_of_stock_age_five"`  // [Required] <p>121-180Days<br /></p>
	QtyOfStockAgeSix   int64  `json:"qty_of_stock_age_six"`   // [Required] <p>&gt;180Days<br /></p>
	ExcessStock        int64  `json:"excess_stock"`           // [Required] <p>expired stock<br /></p>
	AgingStorageTag    int64  `json:"aging_storage_tag"`      // [Required]
}
type OutboundQty struct {
	OutboundTotal    int64 `json:"outbound_total"`    // [Required] <p>Total outbound quantity during the selected time period.</p>
	OutboundSold     int64 `json:"outbound_sold"`     // [Required] <p>"Total sold quantity during the selected time period."</p>
	OutboundReturned int64 `json:"outbound_returned"` // [Required] <p>Total merchant return quantity during the selected time period.</p>
	OutboundDisposed int64 `json:"outbound_disposed"` // [Required] <p>Total disposal quantity during the selected time period.</p>
}
type ResponseDataItemSku struct {
	MtskuId            string       `json:"mtsku_id"`             // [Required] <p>mtsku id<br /></p>
	ModelId            string       `json:"model_id"`             // [Required] <p>Warehouse model SKU ID</p><p>For CB global items, this is equal to the global model_id.</p><p>For local items, it differs from model_id; use shop_model_id to match the model_id</p>
	FulfillMappingMode int64        `json:"fulfill_mapping_mode"` // [Required] <p>0-Null；1-Bundle SKU；2-Parent SKU<br /></p>
	ModelName          string       `json:"model_name"`           // [Required] <p>model name<br /></p>
	Barcode            string       `json:"barcode"`              // [Required]
	WhsList            []ItemSkuWhs `json:"whs_list"`             // [Required] <p>Info of whs<br /></p>
	ShopSkuList        []SkuShopSku `json:"shop_sku_list"`        // [Required]
}
type ResponseDataItemSkuWhs struct {
	WhsId            string `json:"whs_id"`              // [Required] <p>Whs id<br /></p>
	StartOnHandTotal int64  `json:"start_on_hand_total"` // [Required] <p>Total warehouse inventory at the start time<br /></p>
	InboundTotal     int64  `json:"inbound_total"`       // [Required] <p>Inbound quantity to the warehouse during the selected time period.<br /></p>
	OutboundTotal    int64  `json:"outbound_total"`      // [Required] <p>Outbound quantity from the warehouse during the selected time period<br /></p>
	AdjustTotal      int64  `json:"adjust_total"`        // [Required] <p>Inventory adjustment quantity in the warehouse during the selected time period<br /></p>
	EndOnHandTotal   int64  `json:"end_on_hand_total"`   // [Required] <p>Total warehouse inventory at the end time.<br /></p>
}
type Sku struct {
	MtskuId            string       `json:"mtsku_id"`             // [Required] <p>mtsku id</p>
	ModelId            string       `json:"model_id"`             // [Required] <p>Warehouse model SKU ID</p><p>For CB global items, this is equal to the global model_id.</p><p>      </p><p>For local items, it differs from model_id; use shop_model_id to match the model_id</p>
	FulfillMappingMode int64        `json:"fulfill_mapping_mode"` // [Required] <p>0-Null；1-Bundle SKU；2-Parent SKU</p>
	ModelName          string       `json:"model_name"`           // [Required] <p>model name</p>
	NotMovingTag       int64        `json:"not_moving_tag"`       // [Required]
	WhsList            []Whs        `json:"whs_list"`             // [Required] <p>Info of whs</p>
	ShopSkuList        []SkuShopSku `json:"shop_sku_list"`        // [Required]
}
type SkuShopSku struct {
	ShopSkuId   string `json:"shop_sku_id"`   // [Required] <p>shop level sku_id  shop_sku_id="item_id" _ "model_id"</p>
	ShopItemId  string `json:"shop_item_id"`  // [Required] <p>shop_item_id="item_id" in Product Module</p>
	ShopModelId string `json:"shop_model_id"` // [Required] <p>shop_model_id= item level model_id</p>
}
type SkuWhs struct {
	WhsId            string `json:"whs_id"`             // [Required] <p>warehouse ID</p>
	ExpiringQty      int64  `json:"expiring_qty"`       // [Required] <p>Stocks that are expiring soon and should be sold as soon as possible.<br /></p>
	ExpiredQty       int64  `json:"expired_qty"`        // [Required] <p>Stock past expiry date.<br /></p>
	ExpiryBlockedQty int64  `json:"expiry_blocked_qty"` // [Required] <p>Stocks that are too near to expiry and cannot be sold any more.<br /></p>
	DamagedQty       int64  `json:"damaged_qty"`        // [Required] <p>Stock in damaged condition and cannot be sold.<br /></p>
	NormalQty        int64  `json:"normal_qty"`         // [Required] <p>Stocks that are normal.<br /></p>
	TotalQty         int64  `json:"total_qty"`          // [Required] <p>Total stocks on hand.<br /></p>
}
type StartQty struct {
	StartOnHandTotal int64 `json:"start_on_hand_total"` // [Required] <p>sku number at the start time<br /></p>
	StartSellable    int64 `json:"start_sellable"`      // [Required] <p>Number of sellable SKUs at the start time<br /></p>
	StartReserved    int64 `json:"start_reserved"`      // [Required] <p>Number of reserved SKUs at the start time.<br /></p>
	StartUnsellable  int64 `json:"start_unsellable"`    // [Required]
}
type Whs struct {
	WhsId                      string  `json:"whs_id"`                         // [Required] <p>&nbsp;Warehouse ID</p>
	StockLevel                 int64   `json:"stock_level"`                    // [Required] <p>-1-No need to calculate stock level；0-None；1-Low Stock &amp; No Sellable stock; 2-Low Stock &amp; To replenish; 3-Low Stock &amp; Replenished; 4-Excess<br /></p>
	IrApprovalQty              int64   `json:"ir_approval_qty"`                // [Required] <p>IR approval but no ASN generated will be included</p>
	InTransitPendingPutawayQty int64   `json:"in_transit_pending_putaway_qty"` // [Required] <p>ASN in-transit，ASN pending putaway, Move transfer in transit and Move transfer pending putaway will be included</p>
	SellableQty                int64   `json:"sellable_qty"`                   // [Required] <p>Stocks that are available for sale. This includes warehouse sellable stock, move transfer &amp; rack transfer reserved stock that available for sale, pre-order stock.This quantity may be greater than qty displayed to buyers as it includes stock reserved for upcoming promotions.<br /></p>
	ReservedQty                int64   `json:"reserved_qty"`                   // [Required] <p>Stocks reserved by buyer order, RTS. And also includes rack transfer reserved stock that are not available for sale<br /></p>
	UnsellableQty              int64   `json:"unsellable_qty"`                 // [Required] <p>Stocks in the warehouse that are not available for sale. This includes damaged stocks, isolated stock due to expired/near expiring, in warehouse staging location, missing, etc.<br /></p>
	ExcessStock                int64   `json:"excess_stock"`                   // [Required] <p>Number of units that are above 6 days of sales coverage Days.</p>
	CoverageDays               float64 `json:"coverage_days"`                  // [Required] <p>Days that the current sellable and pending inbound inventory are expected to last based on current selling speed.<br /></p>
	InWhsCoverageDays          float64 `json:"in_whs_coverage_days"`           // [Required] <p>Days that the current sellable inventory are expected to last based on current selling speed.<br /></p>
	SellingSpeed               float64 `json:"selling_speed"`                  // [Required] <p>Average daily sold quantity</p>
	Last7Sold                  int64   `json:"last_7_sold"`                    // [Required] <p>Sales qty last 7 days</p>
	Last15Sold                 int64   `json:"last_15_sold"`                   // [Required] <p>Sales qty last 15 days<br /></p>
	Last30Sold                 int64   `json:"last_30_sold"`                   // [Required] <p>Sales qty last 30 days<br /></p>
	Last60Sold                 int64   `json:"last_60_sold"`                   // [Required] <p>Sales qty last 60 days<br /></p>
	Last90Sold                 int64   `json:"last_90_sold"`                   // [Required] <p>Sales qty last 90 days<br /></p>
}
