package goshopee

type AddItemListResponseData struct {
	InvalidItemIdList []Failed `json:"invalid_item_id_list"` // [Required] List of invalid item ids.
	ShopCategoryId    int64    `json:"shop_category_id"`     // [Required] ShopCategory's unique identifier.
	CurrentCount      int64    `json:"current_count"`        // [Required] Count of items under this shop category after deletion.
}
type AddShopCategoryRequest struct {
	Name       string `json:"name"`                  // [Required] ShopCategory's name.
	SortWeight *int64 `json:"sort_weight,omitempty"` // [Optional] ShopCategory's sort weight. The maximum number should be 2147483546.
}
type AddShopCategoryResponse struct {
	BaseResponse                             // Common response fields
	Response     AddShopCategoryResponseData `json:"response"` // Response data
}
type AddShopCategoryResponseData struct {
	ShopCategoryId int64 `json:"shop_category_id"` // [Required] ShopCategory's unique identifier.
}
type DeleteItemListResponseData struct {
	ShopCategoryId int64   `json:"shop_category_id"` // [Required] ShopCategory's unique identifier.
	InvalidItemId  []int64 `json:"invalid_item_id"`  // [Required] The list of item ids which are invalid; In other words, the item ids not being under the category.
	CurrentCount   int64   `json:"current_count"`    // [Required] count of items under this shop category after deleting
}
type DeleteShopCategoryRequest struct {
	ShopCategoryId int64 `json:"shop_category_id"` // [Required] ShopCategory's unique identifier.
}
type DeleteShopCategoryResponse struct {
	BaseResponse                                // Common response fields
	Response     DeleteShopCategoryResponseData `json:"response"` // Response data
}
type DeleteShopCategoryResponseData struct {
	ShopCategoryId int64  `json:"shop_category_id"` // [Required] ShopCategory's unique identifier.
	Msg            string `json:"msg"`              // [Required] The return message of the operation result
}
type Failed struct {
	ItemId      int64  `json:"item_id"`      // [Required] The invalid item id.
	FailError   string `json:"fail_error"`   // [Required] The reason of the fail.
	FailMessage string `json:"fail_message"` // [Required] The detailed reason of the failure and the hints of error fixing
}
type GetShopCategoryListRequest struct {
	PageNo   int64 `json:"page_no"`   // [Required] Specifies the total returned data per entry. The parameter range of page_no should be [1, 100]
	PageSize int64 `json:"page_size"` // [Required] Specifies the starting entry of data to return in the current call. The parameter range of page_size should be [1, 2147483647]
}
type GetShopCategoryListResponse struct {
	BaseResponse                                 // Common response fields
	Response     GetShopCategoryListResponseData `json:"response"` // Response data
}
type GetShopCategoryListResponseData struct {
	ShopCategorys []ShopCategorys `json:"shop_categorys"` // [Required] ShopCategory's unique identifier.
	TotalCount    int64           `json:"total_count"`    // [Required] This is to indicate the whole number of  in-shop categories under the shop.
	More          bool            `json:"more"`           // [Required] This is to indicate whether the list is more than one page. If this value is true, you may want to continue to check next page to retrieve the rest.
}
type ShopCategoryAddItemListRequest struct {
	ItemList       []int64 `json:"item_list"`        // [Required] Shopee's unique identifiers list for an item. Max. 100 items to be deleted per request.
	ShopCategoryId int64   `json:"shop_category_id"` // [Required] ShopCategory's unique identifier.
}
type ShopCategoryAddItemListResponse struct {
	BaseResponse                         // Common response fields
	Response     AddItemListResponseData `json:"response"` // Response data
}
type ShopCategoryDeleteItemListRequest struct {
	ItemList       []int64 `json:"item_list"`        // [Required] ShopCategory's unique identifier.
	ShopCategoryId int64   `json:"shop_category_id"` // [Required] The list of items need to be deleted. To note that the items which can be deleted successfully should be under this category.
}
type ShopCategoryDeleteItemListResponse struct {
	BaseResponse                            // Common response fields
	Response     DeleteItemListResponseData `json:"response"` // Response data
}
type ShopCategoryGetItemListRequest struct {
	PageNo         *int64 `json:"page_no,omitempty"`   // [Optional] If many items are available to retrieve, you may need to call this api multiple times to retrieve all the data. And the default will be 0. page_size*page_no should be [0, 2147483446].
	PageSize       *int64 `json:"page_size,omitempty"` // [Optional] Specifies the starting entry of data to return in the current call. Default is 1000. The input range of page_size is [0, 1000]
	ShopCategoryId int64  `json:"shop_category_id"`    // [Required] ShopCategory's unique identifier.
}
type ShopCategoryGetItemListResponse struct {
	BaseResponse                                     // Common response fields
	Response     ShopCategoryGetItemListResponseData `json:"response"` // Response data
}
type ShopCategoryGetItemListResponseData struct {
	ItemList   []int64 `json:"item_list"`   // [Required] A list of Shopee's unique identifiers for items.
	TotalCount int64   `json:"total_count"` // [Required] This is to indicate the whole number of items under the shop category.
	More       bool    `json:"more"`        // [Required] This is to indicate whether the item list is more than one page. If this value is true, you may want to continue to check next page to retrieve the rest of items.
}
type ShopCategorys struct {
	ShopCategoryId int64  `json:"shop_category_id"` // [Required] ShopCategory's unique identifier.
	Status         int64  `json:"status"`           // [Required] ShopCategory's status. Applicable values--1: 'NORMAL', 2: 'INACTIVE', 0: 'DELETED'
	Name           string `json:"name"`             // [Required] ShopCategory's name.
	SortWeight     int64  `json:"sort_weight"`      // [Required] ShopCategory's sort weight.
}
type UpdateShopCategoryRequest struct {
	Name           *string `json:"name,omitempty"`        // [Optional] ShopCategory's name.
	ShopCategoryId int64   `json:"shop_category_id"`      // [Required] ShopCategory's unique identifier.
	SortWeight     *int64  `json:"sort_weight,omitempty"` // [Optional] ShopCategory's sort weight.
	Status         *string `json:"status,omitempty"`      // [Optional] ShopCategory's status. Applicable values: NORMAL, INACTIVE, DELETED.
}
type UpdateShopCategoryResponse struct {
	BaseResponse                                // Common response fields
	Response     UpdateShopCategoryResponseData `json:"response"` // Response data
}
type UpdateShopCategoryResponseData struct {
	ShopCategoryId int64  `json:"shop_category_id"` // [Required] This is to indicate whether the shop categories list is more than one page. If this value is true, you may want to continue to check next page to retrieve the rest of shop categories
	Name           string `json:"name"`             // [Required] ShopCategory's name.
	SortWeight     int64  `json:"sort_weight"`      // [Required] ShopCategory's sort weight.
	Status         string `json:"status"`           // [Required] ShopCategory's status. Applicable values: NORMAL, INACTIVE, DELETED.
}
