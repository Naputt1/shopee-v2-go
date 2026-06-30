package goshopee

type Cursor struct {
	NextId   int64 `json:"next_id"`   // [Required]
	PrevId   int64 `json:"prev_id"`   // [Required]
	PageSize int64 `json:"page_size"` // [Required]
}
type EnterpriseInfo struct {
	CompanyName             string `json:"company_name"`              // [Required]
	Cnpj                    string `json:"cnpj"`                      // [Required]
	StateRegistrationNumber string `json:"state_registration_number"` // [Required]
	IsFreightPayer          bool   `json:"is_freight_payer"`          // [Required]
}
type GetMerchantInfoResponse struct {
	BaseResponse            // Common response fields
	AuthTime         int64  `json:"auth_time,omitempty"`         // The timestamp when the merchant was authorized to the partner.
	ExpireTime       int64  `json:"expire_time,omitempty"`       // Use this field to indicate the expiration date for merchant authorization.
	IsUpgradedCbsc   bool   `json:"is_upgraded_cbsc,omitempty"`  // <p>Use this filed to indicate whether this merchant is upgraded to cbsc.<br /></p>
	MerchantCurrency string `json:"merchant_currency,omitempty"` // <p>The three-digit code representing the currency unit used for the item in this merchant. If currency haven't been setting in CNSC/KRSC, it will be empty.</p><p>China merchant support CNY and USD currently.</p><p>Korea merchant support KRW and USD currently.&nbsp;</p><p>Hong kong merchant support USD currently, will support HKD later.</p>
	MerchantName     string `json:"merchant_name,omitempty"`     // Name of the merchant.
	MerchantRegion   string `json:"merchant_region,omitempty"`   // <p>Region of the merchant. Including KR, HK and CN.<br /></p>
}
type GetMerchantPrepaidAccountListRequest struct {
	PageNo   int64 `json:"page_no" url:"page_no"`     // [Required] <p>Specifies the page number of data to return in the current call. Starting from 1. if data is more than one page, the page_no can be some entry to start next call.</p>
	PageSize int64 `json:"page_size" url:"page_size"` // [Required] <p>Each result set is returned as a page of entries. Use the "page_size" filters to control the maximum number of entries to retrieve per page (i.e., per call), and the "page_no" to start next call. This integer value is used to specify the maximum number of entries to return in a single "page" of data.<br /></p><p>Min: 1 Max:10<br /></p>
}
type GetMerchantPrepaidAccountListResponse struct {
	BaseResponse                                           // Common response fields
	Response     GetMerchantPrepaidAccountListResponseData `json:"response"` // <p>Detail informations you are querying.</p>
}
type GetMerchantPrepaidAccountListResponseData struct {
	Total int64                                           `json:"total"` // [Required]
	List  []GetMerchantPrepaidAccountListResponseDataList `json:"list"`  // [Required]
	More  bool                                            `json:"more"`  // [Required] <p>This is to indicate whether the list is more than one page. If this value is true, you may want to continue to check next page to retrieve the rest of datas.<br /></p>
}
type GetMerchantPrepaidAccountListResponseDataList struct {
	PrepaidAccountId            int64  `json:"prepaid_account_id"`             // [Required] <p>Record ID</p>
	PrepaidAccountCourierKey    string `json:"prepaid_account_courier_key"`    // [Required] <p>Courier Company Key (快递公司编码).</p>
	PrepaidAccountCourierName   string `json:"prepaid_account_courier_name"`   // [Required] <p>Courier Company Name (快递公司名称).</p>
	PrepaidAccountPartnerId     string `json:"prepaid_account_partner_id"`     // [Required] <p>Prepaid Account Number (电子面单账户号码).</p>
	PrepaidAccountPartnerKey    string `json:"prepaid_account_partner_key"`    // [Required] <p>Prepaid Account Password (电子面单账户密码).</p>
	PrepaidAccountPartnerSecret string `json:"prepaid_account_partner_secret"` // [Required] <p>Partner Secret (电子面单密钥).</p>
	PrepaidAccountPartnerName   string `json:"prepaid_account_partner_name"`   // [Required] <p>Partner Name (电子面单客户账户名称).</p>
	PrepaidAccountPartnerNet    string `json:"prepaid_account_partner_net"`    // [Required] <p>Branch Name (网点名称).</p>
	PrepaidAccountPartnerCode   string `json:"prepaid_account_partner_code"`   // [Required] <p>Partner Code (电子面单承载编号).</p>
	PrepaidAccountCheckMan      string `json:"prepaid_account_check_man"`      // [Required] <p>Delivery Agent Name (电子面单承载快递员名).</p>
	PrepaidAccountIsDefault     bool   `json:"prepaid_account_is_default"`     // [Required] <p>This is to indicate whether the prepaid account is Default Prepaid Account.</p>
}
type GetMerchantWarehouseListRequest struct {
	Cursor        *Cursor `json:"cursor"`         // [Required] <p>// how to use DoubleSidedCursor</p><p>// Get data for the first page: Please pass next_id = 0 or nil, page_size = {your page size}.</p><p>// Get data for the next page: Please pass the Cursor from the previous response, and set prev_id=nil;</p><p>// Get data for the prev page: Please pass the Cursor from the previous response, and set next_id=nil;</p><p>// Stop fetching next data: The Cursor.next_id in the previous response is nil.</p><p>// Stop fetching prev data: The Cursor.prev_id in the previous response is nil.<br /></p>
	WarehouseType int64   `json:"warehouse_type"` // [Required] <p>1 means pickup warehouse</p><p>2 means return warehouse<br /></p>
}
type GetMerchantWarehouseListResponse struct {
	BaseResponse                                      // Common response fields
	Response     GetMerchantWarehouseListResponseData `json:"response"` //
}
type GetMerchantWarehouseListResponseData struct {
	TotalCount    int64       `json:"total_count"`    // [Required] <p>Total count of all warehouses.<br /></p>
	WarehouseList []Warehouse `json:"warehouse_list"` // [Required]
	Cursor        *Cursor     `json:"cursor"`         // [Required]
}
type GetMerchantWarehouseLocationListResponse struct {
	BaseResponse                                              // Common response fields
	Response     GetMerchantWarehouseLocationListResponseData `json:"response"` //
}
type GetMerchantWarehouseLocationListResponseData struct {
	LocationId    string `json:"location_id"`    // [Required] <p>Location identifier for stocks. Different location_ids represent that your addresses are in different item stocks<br /></p>
	WarehouseName string `json:"warehouse_name"` // [Required] <p>The warehouse name filled in when creating the warehouse address<br /></p>
}
type GetShopListByMerchantRequest struct {
	PageNo   int64 `json:"page_no" url:"page_no"`     // [Required] Specifies the page number of data to return in the current call. Starting from 1. if data is more than one page, the page_no can be some entry to start next call.
	PageSize int64 `json:"page_size" url:"page_size"` // [Required] Each result set is returned as a page of entries. Use the "page_size" filters to control the maximum number of entries to retrieve per page (i.e., per call), and the "page_no" to start next call. This integer value is used to specify the maximum number of entries to return in a single "page" of data.No more than 500.
}
type GetShopListByMerchantResponse struct {
	BaseResponse        // Common response fields
	More         bool   `json:"more,omitempty"`      // This is to indicate whether the list is more than one page. If this value is true, you may want to continue to check next page to retrieve the rest of datas.
	ShopList     []Shop `json:"shop_list,omitempty"` // list of shop authorized to the partner and bound to the merchant.
}
type GetWarehouseEligibleShopListRequest struct {
	Cursor        *Cursor `json:"cursor"`         // [Required] <p>// how to use DoubleSidedCursor</p><p>// Get data for the first page: Please pass next_id = 0 or nil, page_size = {your page size}.</p><p>// Get data for the next page: Please pass the Cursor from the previous response, and set prev_id=nil;</p><p>// Get data for the prev page: Please pass the Cursor from the previous response, and set next_id=nil;</p><p>// Stop fetching next data: The Cursor.next_id in the previous response is nil.</p><p>// Stop fetching prev data: The Cursor.prev_id in the previous response is nil.<br /></p>
	WarehouseId   int64   `json:"warehouse_id"`   // [Required] <p>Warehouse address identifier.<br /></p>
	WarehouseType int64   `json:"warehouse_type"` // [Required] <p>1 means pickup warehouse</p><p>2 means return warehouse<br /></p>
}
type GetWarehouseEligibleShopListResponse struct {
	BaseResponse                                          // Common response fields
	Response     GetWarehouseEligibleShopListResponseData `json:"response"` //
}
type GetWarehouseEligibleShopListResponseData struct {
	ShopList []ResponseDataShop `json:"shop_list"` // [Required] <p>Eligible shop list of the warehouse</p>
	Cursor   *Cursor            `json:"cursor"`    // [Required]
}
type ResponseDataShop struct {
	ShopId   int64  `json:"shop_id"`   // [Required] <p>Shopee's unique identifier for a shop.<br /></p>
	ShopName string `json:"shop_name"` // [Required] <p>Name of the shop.<br /></p>
}
type Shop struct {
	ShopId       int64          `json:"shop_id"`        // [Required] Shopee's unique identifier for a shop.
	SipAffiShops []SipAffiShops `json:"sip_affi_shops"` // [Required] List of SIP affiliate shops.Only primary shop will return this parameter
}
type SipAffiShops struct {
	AffiShopId int64 `json:"affi_shop_id"` // [Required]  Affiliate shop's id.
}
type Warehouse struct {
	WarehouseId     int64             `json:"warehouse_id"`     // [Required] <p>Warehouse address identifier.<br /></p>
	WarehouseName   string            `json:"warehouse_name"`   // [Required] <p>The warehouse name filled in when creating the warehouse.<br /></p>
	WarehouseType   int64             `json:"warehouse_type"`   // [Required] <p>1 means pickup warehouse</p><p>2 means return warehouse<br /></p>
	WarehouseRegion string            `json:"warehouse_region"` // [Required] <p>Region of your warehouse.<br /></p>
	LocationId      string            `json:"location_id"`      // [Required] <p>Location identifier for stocks. Different location_ids represent that your addresses are in different item stocks.<br /></p>
	Address         *WarehouseAddress `json:"address"`          // [Required]
	EnterpriseInfo  *EnterpriseInfo   `json:"enterprise_info"`  // [Required]
}
type WarehouseAddress struct {
	AddressName string `json:"address_name"` // [Required] <p>The address name filled in when creating the warehouse.<br /></p>
	Region      string `json:"region"`       // [Required] <p>Region of your warehouse address.<br /></p>
	Address     string `json:"address"`      // [Required] <p>Detail address of your warehouse address.<br /></p>
	City        string `json:"city"`         // [Required] <p>City of your warehouse address.<br /></p>
	District    string `json:"district"`     // [Required] <p>Distinct of your warehouse address.<br /></p>
	State       string `json:"state"`        // [Required] <p>State of your warehouse address.<br /></p>
	Town        string `json:"town"`         // [Required] <p>Town of your warehouse address.<br /></p>
	ZipCode     string `json:"zip_code"`     // [Required] <p>Zipcode of your warehouse address.<br /></p>
}
