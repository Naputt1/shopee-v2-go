package goshopee

import (
	"context"
)

type BusinessInsightsService interface {
	// GetMarketingHotListingonlyBRNow {"content":"<p>Provide all metrics currently available on the Business Insights Buybox dashboard, including metrics not explicitly requested by Local (for example, CTR), covering both:</p><ul><li>Shop-level performance</li><li>Product-level performance</li></ul>","raw_content":[{"name":"paragraph","children":[{"data":"Provide all metrics currently available on the Business Insights Buybox dashboard, including metrics not explicitly requested by Local (for example, CTR), covering both:"}]},{"attributes":{"listIndent":0,"listType":"bulleted"},"name":"listItem","children":[{"data":"Shop-level performance"}]},{"attributes":{"listIndent":0,"listType":"bulleted"},"name":"listItem","children":[{"data":"Product-level performance"}]}]}
	// Path: /api/v2/business_insights/get_marketing_hot_listing
	// https://open.shopee.com/documents/v2/v2.business_insights.get_marketing_hot_listing（only BR now）?module=141&type=1
	GetMarketingHotListingonlyBRNow(ctx context.Context, sid uint64, mid uint64, tok string, req GetMarketingHotListingonlyBRNowRequest) (*GetMarketingHotListingonlyBRNowResponse, error)
}

type BusinessInsightsServiceOp[T any] struct {
	client *Client[T]
}

// GetMarketingHotListingonlyBRNow {"content":"<p>Provide all metrics currently available on the Business Insights Buybox dashboard, including metrics not explicitly requested by Local (for example, CTR), covering both:</p><ul><li>Shop-level performance</li><li>Product-level performance</li></ul>","raw_content":[{"name":"paragraph","children":[{"data":"Provide all metrics currently available on the Business Insights Buybox dashboard, including metrics not explicitly requested by Local (for example, CTR), covering both:"}]},{"attributes":{"listIndent":0,"listType":"bulleted"},"name":"listItem","children":[{"data":"Shop-level performance"}]},{"attributes":{"listIndent":0,"listType":"bulleted"},"name":"listItem","children":[{"data":"Product-level performance"}]}]}
// Path: /api/v2/business_insights/get_marketing_hot_listing
// https://open.shopee.com/documents/v2/v2.business_insights.get_marketing_hot_listing（only BR now）?module=141&type=1
func (s *BusinessInsightsServiceOp[T]) GetMarketingHotListingonlyBRNow(ctx context.Context, sid uint64, mid uint64, tok string, req GetMarketingHotListingonlyBRNowRequest) (*GetMarketingHotListingonlyBRNowResponse, error) {
	path := "/business_insights/get_marketing_hot_listing"
	resp := new(GetMarketingHotListingonlyBRNowResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}
