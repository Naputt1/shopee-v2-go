package goshopee

type AddTopPicksRequest struct {
	IsActivated bool `json:"is_activated"` // [Required] 
	ItemIdList []int64 `json:"item_id_list"` // [Required] 
	Name string `json:"name"` // [Required] 
}
type AddTopPicksResponse struct {
	BaseResponse // Common response fields
	Response AddTopPicksResponseData `json:"response"` // Detail informations you are querying.
}
type AddTopPicksResponseData struct {
	CollectionList []Collection `json:"collection_list"` // [Required] The top picks list in this shop.
}
type Collection struct {
	IsActivated bool `json:"is_activated"` // [Required] whether is activated
	ItemList []CollectionItem `json:"item_list"` // [Required] a list of item
	TopPicksId int64 `json:"top_picks_id"` // [Required] collection id
	Name string `json:"name"` // [Required] collection name
}
type CollectionItem struct {
	ItemName string `json:"item_name"` // [Required] The name of item.
	ItemId int64 `json:"item_id"` // [Required] The id of item.
	CurrentPrice float64 `json:"current_price"` // [Required] The price before tax of item.
	InflatedPriceOfCurrentPrice float64 `json:"inflated_price_of_current_price"` // [Required] The price after tax of item.
	Sales int64 `json:"sales"` // [Required] The sales of item.
}
type DeleteTopPicksRequest struct {
	TopPicksId int64 `json:"top_picks_id"` // [Required] collection id
}
type DeleteTopPicksResponse struct {
	BaseResponse // Common response fields
	Response DeleteTopPicksResponseData `json:"response"` // Detail informations you are querying.
}
type DeleteTopPicksResponseData struct {
	TopPicksId int64 `json:"top_picks_id"` // [Required] collection id
}
type GetTopPicksListResponse struct {
	BaseResponse // Common response fields
	Response GetTopPicksListResponseData `json:"response"` // Detail informations you are querying.
}
type GetTopPicksListResponseData struct {
	CollectionList []Collection `json:"collection_list"` // [Required] The top picks list in this shop.
}
type UpdateTopPicksRequest struct {
	IsActivated *bool `json:"is_activated,omitempty"` // [Optional] if true, we will close other collection and open this collection
	ItemIdList []int64 `json:"item_id_list,omitempty"` // [Optional] a list of item id, and we will cover old item_ids by new_item_ids
	Name *string `json:"name,omitempty"` // [Optional] collection name
	TopPicksId int64 `json:"top_picks_id"` // [Required] collection id
}
type UpdateTopPicksResponse struct {
	BaseResponse // Common response fields
	Response UpdateTopPicksResponseData `json:"response"` // Detail informations you are querying.
}
type UpdateTopPicksResponseData struct {
	CollectionList []Collection `json:"collection_list"` // [Required] The top picks list in this shop.
}
