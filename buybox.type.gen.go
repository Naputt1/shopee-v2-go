package goshopee

type GetBuyboxModelPerformanceonlyForBRNowRequest struct {
	ModelIdList []int64 `json:"model_id_list"` // [Required]
}
type GetBuyboxModelPerformanceonlyForBRNowResponse struct {
	BaseResponse                                                   // Common response fields
	Response     GetBuyboxModelPerformanceonlyForBRNowResponseData `json:"response"` // Response data
}
type GetBuyboxModelPerformanceonlyForBRNowResponseData struct {
	PerformanceList []Performance `json:"performance_list"`  // [Required]
	DataUpdatedTime string        `json:"data_updated_time"` // [Required]
}
type GetBuyboxModelsByModelIdonlyForBRNowRequest struct {
	ModelIdList []int64 `json:"model_id_list"` // [Required]
}
type GetBuyboxModelsByModelIdonlyForBRNowResponse struct {
	BaseResponse                                                  // Common response fields
	Response     GetBuyboxModelsByModelIdonlyForBRNowResponseData `json:"response"` // Response data
}
type GetBuyboxModelsByModelIdonlyForBRNowResponseData struct {
	ModelList []Model `json:"model_list"` // [Required]
}
type GetBuyboxModelsByShopIdonlyForBRNowRequest struct {
	Limit  *int64 `json:"limit,omitempty"`  // [Optional] <p>Page size. Valid range is 1 to 250.</p>
	Offset *int64 `json:"offset,omitempty"` // [Optional] <p>Pagination offset. Must be non-negative.</p>
}
type GetBuyboxModelsByShopIdonlyForBRNowResponse struct {
	BaseResponse                                                 // Common response fields
	Response     GetBuyboxModelsByShopIdonlyForBRNowResponseData `json:"response"` // Response data
}
type GetBuyboxModelsByShopIdonlyForBRNowResponseData struct {
	ModelList []Model   `json:"model_list"` // [Required]
	PageInfo  *PageInfo `json:"page_info"`  // [Required]
	ShopId    int64     `json:"shop_id"`    // [Required]
}
type GetBuyboxShopPerformanceonlyForBRNowResponse struct {
	BaseResponse                                                  // Common response fields
	Response     GetBuyboxShopPerformanceonlyForBRNowResponseData `json:"response"` // Response data
}
type GetBuyboxShopPerformanceonlyForBRNowResponseData struct {
	Past7DaysSoldCount int64   `json:"past_7_days_sold_count"` // [Required]
	TotalSoldCount     int64   `json:"total_sold_count"`       // [Required]
	Sales              float64 `json:"sales"`                  // [Required]
	DataUpdatedTime    string  `json:"data_updated_time"`      // [Required]
	ShopId             int64   `json:"shop_id"`                // [Required]
}
type Model struct {
	ModelId             int64  `json:"model_id"`               // [Required]
	IsModelMapped       bool   `json:"is_model_mapped"`        // [Required]
	IsEligible          bool   `json:"is_eligible"`            // [Required]
	ModelToggleOnStatus bool   `json:"model_toggle_on_status"` // [Required]
	LastUpdatedTime     string `json:"last_updated_time"`      // [Required]
}
type PageInfo struct {
	Offset     int64 `json:"offset"`      // [Required]
	Limit      int64 `json:"limit"`       // [Required]
	TotalCount int64 `json:"total_count"` // [Required]
}
type Performance struct {
	ModelId            int64 `json:"model_id"`               // [Required]
	Past7DaysSoldCount int64 `json:"past_7_days_sold_count"` // [Required]
}
type UpdateBuyboxModelEnrollmentonlyForBRNowRequest struct {
	ModelId             int64 `json:"model_id"`               // [Required]
	ModelToggleOnStatus bool  `json:"model_toggle_on_status"` // [Required]
}
type UpdateBuyboxModelEnrollmentonlyForBRNowResponse struct {
	BaseResponse                                                     // Common response fields
	Response     UpdateBuyboxModelEnrollmentonlyForBRNowResponseData `json:"response"` // Response data
}
type UpdateBuyboxModelEnrollmentonlyForBRNowResponseData struct {
	ModelId             int64  `json:"model_id"`               // [Required]
	ModelToggleOnStatus bool   `json:"model_toggle_on_status"` // [Required]
	LastUpdatedTime     string `json:"last_updated_time"`      // [Required]
	ShopId              int64  `json:"shop_id"`                // [Required]
}
