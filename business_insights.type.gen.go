package goshopee

type GetMarketingHotListingonlyBRNowRequest struct {
	EndTime       int64   `json:"end_time"`                  // [Required] <p>End of the requested time range as a Unix timestamp in seconds. It must be greater than or equal to start_time.</p>
	Period        string  `json:"period"`                    // [Required] <p>Time period of the requested data. Supported values are listed below.</p>
	ProductIdList *string `json:"product_id_list,omitempty"` // [Optional] <p>Product ID list, comma-separated, e.g., 111,222,333</p>
	StartTime     int64   `json:"start_time"`                // [Required] <p>Start of the requested time range as a Unix timestamp in seconds.</p>
}
type GetMarketingHotListingonlyBRNowResponse struct {
	BaseResponse                                                     // Common response fields
	Code                                               int64         `json:"code,omitempty"`                                                       // <p>Business status code. 0 indicates success.</p>
	Msg                                                string        `json:"msg,omitempty"`                                                        // <p>Business status message. The value is ok on success.</p>
	Result                                             []interface{} `json:"result,omitempty"`                                                     // <p>Metrics for every order type supported by the shop country. Non-COD shops return placed and paid; COD shops additionally return confirmed.</p>
	ResultKeyMetrics                                   interface{}   `json:"result_key_metrics,omitempty"`                                         // <p>Shop-level summary metrics and time-series data for the order type.</p>
	ResultKeyMetricsKeyMetrics                         interface{}   `json:"result_key_metrics_key_metrics,omitempty"`                             // <p>Shop-level summary metrics for the requested period.</p>
	ResultKeyMetricsKeyMetricsBuyers                   int64         `json:"result_key_metrics_key_metrics_buyers,omitempty"`                      // <p>Number of buyers.</p>
	ResultKeyMetricsKeyMetricsBuyersPctDiff            float64       `json:"result_key_metrics_key_metrics_buyers_pct_diff,omitempty"`             // <p>Buyer-count change compared with the previous period.</p>
	ResultKeyMetricsKeyMetricsClickThroughRate         float64       `json:"result_key_metrics_key_metrics_click_through_rate,omitempty"`          // <p>Product click-through rate.</p>
	ResultKeyMetricsKeyMetricsClickThroughRatePctDiff  float64       `json:"result_key_metrics_key_metrics_click_through_rate_pct_diff,omitempty"` // <p>Click-through-rate change compared with the previous period.</p>
	ResultKeyMetricsKeyMetricsConversionRate           float64       `json:"result_key_metrics_key_metrics_conversion_rate,omitempty"`             // <p>Conversion rate.</p>
	ResultKeyMetricsKeyMetricsConversionRatePctDiff    float64       `json:"result_key_metrics_key_metrics_conversion_rate_pct_diff,omitempty"`    // <p>Conversion-rate change compared with the previous period.</p>
	ResultKeyMetricsKeyMetricsOrders                   int64         `json:"result_key_metrics_key_metrics_orders,omitempty"`                      // <p>Number of orders.</p>
	ResultKeyMetricsKeyMetricsOrdersPctDiff            float64       `json:"result_key_metrics_key_metrics_orders_pct_diff,omitempty"`             // <p>Order-count change compared with the previous period.</p>
	ResultKeyMetricsKeyMetricsProductClicks            int64         `json:"result_key_metrics_key_metrics_product_clicks,omitempty"`              // <p>Number of product clicks.</p>
	ResultKeyMetricsKeyMetricsProductClicksPctDiff     float64       `json:"result_key_metrics_key_metrics_product_clicks_pct_diff,omitempty"`     // <p>Product-click change compared with the previous period.</p>
	ResultKeyMetricsKeyMetricsProductImpression        int64         `json:"result_key_metrics_key_metrics_product_impression,omitempty"`          // <p>Number of product impressions.</p>
	ResultKeyMetricsKeyMetricsProductImpressionPctDiff float64       `json:"result_key_metrics_key_metrics_product_impression_pct_diff,omitempty"` // <p>Product-impression change compared with the previous period.</p>
	ResultKeyMetricsKeyMetricsSales                    float64       `json:"result_key_metrics_key_metrics_sales,omitempty"`                       // <p>Sales amount for the order type.</p>
	ResultKeyMetricsKeyMetricsSalesPctDiff             float64       `json:"result_key_metrics_key_metrics_sales_pct_diff,omitempty"`              // <p>Sales change compared with the previous period.</p>
	ResultKeyMetricsKeyMetricsUnits                    int64         `json:"result_key_metrics_key_metrics_units,omitempty"`                       // <p>Number of units sold.</p>
	ResultKeyMetricsKeyMetricsUnitsPctDiff             float64       `json:"result_key_metrics_key_metrics_units_pct_diff,omitempty"`              // <p>Units-sold change compared with the previous period.</p>
	ResultKeyMetricsTimeSeries                         []interface{} `json:"result_key_metrics_time_series,omitempty"`                             // <p>Shop-level metrics grouped by time bucket.</p>
	ResultKeyMetricsTimeSeriesBuyers                   int64         `json:"result_key_metrics_time_series_buyers,omitempty"`                      // <p>Number of buyers in the time bucket.</p>
	ResultKeyMetricsTimeSeriesClickThroughRate         float64       `json:"result_key_metrics_time_series_click_through_rate,omitempty"`          // <p>Product click-through rate in the time bucket.</p>
	ResultKeyMetricsTimeSeriesConversionRate           float64       `json:"result_key_metrics_time_series_conversion_rate,omitempty"`             // <p>Conversion rate in the time bucket.</p>
	ResultKeyMetricsTimeSeriesOrders                   int64         `json:"result_key_metrics_time_series_orders,omitempty"`                      // <p>Number of orders in the time bucket.</p>
	ResultKeyMetricsTimeSeriesProductClicks            int64         `json:"result_key_metrics_time_series_product_clicks,omitempty"`              // <p>Number of product clicks in the time bucket.</p>
	ResultKeyMetricsTimeSeriesProductImpression        int64         `json:"result_key_metrics_time_series_product_impression,omitempty"`          // <p>Number of product impressions in the time bucket.</p>
	ResultKeyMetricsTimeSeriesSales                    float64       `json:"result_key_metrics_time_series_sales,omitempty"`                       // <p>Sales amount in the time bucket.</p>
	ResultKeyMetricsTimeSeriesT                        int64         `json:"result_key_metrics_time_series_t,omitempty"`                           // <p>Unix timestamp in seconds representing the time bucket.</p>
	ResultKeyMetricsTimeSeriesUnits                    int64         `json:"result_key_metrics_time_series_units,omitempty"`                       // <p>Number of units sold in the time bucket.</p>
	ResultOrderType                                    string        `json:"result_order_type,omitempty"`                                          // <p>Order type represented by this entry. Possible values are placed, confirmed, and paid.</p>
	ResultPerformance                                  []interface{} `json:"result_performance,omitempty"`                                         // <p>Product-level performance sorted by sales in descending order.</p>
	ResultPerformanceBuyers                            int64         `json:"result_performance_buyers,omitempty"`                                  // <p>Number of buyers.</p>
	ResultPerformanceClickThroughRate                  float64       `json:"result_performance_click_through_rate,omitempty"`                      // <p>Product click-through rate.</p>
	ResultPerformanceConversionRate                    float64       `json:"result_performance_conversion_rate,omitempty"`                         // <p>Conversion rate.</p>
	ResultPerformanceImage                             string        `json:"result_performance_image,omitempty"`                                   // <p>Product image identifier or URL.</p>
	ResultPerformanceItemId                            int64         `json:"result_performance_item_id,omitempty"`                                 // <p>Product ID.</p>
	ResultPerformanceItemName                          string        `json:"result_performance_item_name,omitempty"`                               // <p>Product name.</p>
	ResultPerformanceOrders                            int64         `json:"result_performance_orders,omitempty"`                                  // <p>Number of orders.</p>
	ResultPerformanceProductClicks                     int64         `json:"result_performance_product_clicks,omitempty"`                          // <p>Number of product clicks.</p>
	ResultPerformanceProductImpression                 int64         `json:"result_performance_product_impression,omitempty"`                      // <p>Number of product impressions.</p>
	ResultPerformanceSales                             float64       `json:"result_performance_sales,omitempty"`                                   // <p>Sales amount for the order type.</p>
	ResultPerformanceStatus                            int64         `json:"result_performance_status,omitempty"`                                  // <p>Product status code.</p>
	ResultPerformanceUnits                             int64         `json:"result_performance_units,omitempty"`                                   // <p>Number of units sold.</p>
	ResultPerformanceVariationName                     string        `json:"result_performance_variation_name,omitempty"`                          // <p>Product variation or model name.</p>
}
