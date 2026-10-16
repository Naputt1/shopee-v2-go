package goshopee

import (
	"context"
)

type BuyboxService interface {
	// GetBuyboxModelPerformanceonlyForBRNow {"content":"<p>Get Buybox model performance by model IDs for the authorized shop.</p>","raw_content":[{"name":"paragraph","children":[{"data":"Get Buybox model performance by model IDs for the authorized shop."}]}]}
	// Path: /api/v2/buybox/get_buybox_model_performance
	// https://open.shopee.com/documents/v2/v2.buybox.get_buybox_model_performance（only for BR now）?module=140&type=1
	GetBuyboxModelPerformanceonlyForBRNow(ctx context.Context, sid uint64, mid uint64, tok string, req GetBuyboxModelPerformanceonlyForBRNowRequest) (*GetBuyboxModelPerformanceonlyForBRNowResponse, error)
	// GetBuyboxModelsByModelIdonlyForBRNow {"content":"<p>Get Buybox model information by model IDs for the authorized shop.</p>","raw_content":[{"name":"paragraph","children":[{"data":"Get Buybox model information by model IDs for the authorized shop."}]}]}
	// Path: /api/v2/buybox/get_buybox_models_by_model_id
	// https://open.shopee.com/documents/v2/v2.buybox.get_buybox_models_by_model_id（only for BR now）?module=140&type=1
	GetBuyboxModelsByModelIdonlyForBRNow(ctx context.Context, sid uint64, mid uint64, tok string, req GetBuyboxModelsByModelIdonlyForBRNowRequest) (*GetBuyboxModelsByModelIdonlyForBRNowResponse, error)
	// GetBuyboxModelsByShopIdonlyForBRNow {"content":"<p>Get the paginated Buybox model list for the authorized shop.</p>","raw_content":[{"name":"paragraph","children":[{"data":"Get the paginated Buybox model list for the authorized shop."}]}]}
	// Path: /api/v2/buybox/get_buybox_models_by_shop_id
	// https://open.shopee.com/documents/v2/v2.buybox.get_buybox_models_by_shop_id（only for BR now）?module=140&type=1
	GetBuyboxModelsByShopIdonlyForBRNow(ctx context.Context, sid uint64, mid uint64, tok string, req GetBuyboxModelsByShopIdonlyForBRNowRequest) (*GetBuyboxModelsByShopIdonlyForBRNowResponse, error)
	// GetBuyboxShopPerformanceonlyForBRNow {"content":"<p>Get Buybox shop performance for the authorized shop.</p>","raw_content":[{"name":"paragraph","children":[{"data":"Get Buybox shop performance for the authorized shop."}]}]}
	// Path: /api/v2/buybox/get_buybox_shop_performance
	// https://open.shopee.com/documents/v2/v2.buybox.get_buybox_shop_performance（only for BR now）?module=140&type=1
	GetBuyboxShopPerformanceonlyForBRNow(ctx context.Context, sid uint64, mid uint64, tok string) (*GetBuyboxShopPerformanceonlyForBRNowResponse, error)
	// UpdateBuyboxModelEnrollmentonlyForBRNow {"content":"<p>Update Buybox model enrollment status for the authorized shop.</p>","raw_content":[{"name":"paragraph","children":[{"data":"Update Buybox model enrollment status for the authorized shop."}]}]}
	// Path: /api/v2/buybox/update_buybox_model_enrollment
	// https://open.shopee.com/documents/v2/v2.buybox.update_buybox_model_enrollment（only for BR now）?module=140&type=1
	UpdateBuyboxModelEnrollmentonlyForBRNow(ctx context.Context, sid uint64, mid uint64, tok string, req UpdateBuyboxModelEnrollmentonlyForBRNowRequest) (*UpdateBuyboxModelEnrollmentonlyForBRNowResponse, error)
}

type BuyboxServiceOp[T any] struct {
	client *Client[T]
}

// GetBuyboxModelPerformanceonlyForBRNow {"content":"<p>Get Buybox model performance by model IDs for the authorized shop.</p>","raw_content":[{"name":"paragraph","children":[{"data":"Get Buybox model performance by model IDs for the authorized shop."}]}]}
// Path: /api/v2/buybox/get_buybox_model_performance
// https://open.shopee.com/documents/v2/v2.buybox.get_buybox_model_performance（only for BR now）?module=140&type=1
func (s *BuyboxServiceOp[T]) GetBuyboxModelPerformanceonlyForBRNow(ctx context.Context, sid uint64, mid uint64, tok string, req GetBuyboxModelPerformanceonlyForBRNowRequest) (*GetBuyboxModelPerformanceonlyForBRNowResponse, error) {
	path := "/buybox/get_buybox_model_performance"
	resp := new(GetBuyboxModelPerformanceonlyForBRNowResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// GetBuyboxModelsByModelIdonlyForBRNow {"content":"<p>Get Buybox model information by model IDs for the authorized shop.</p>","raw_content":[{"name":"paragraph","children":[{"data":"Get Buybox model information by model IDs for the authorized shop."}]}]}
// Path: /api/v2/buybox/get_buybox_models_by_model_id
// https://open.shopee.com/documents/v2/v2.buybox.get_buybox_models_by_model_id（only for BR now）?module=140&type=1
func (s *BuyboxServiceOp[T]) GetBuyboxModelsByModelIdonlyForBRNow(ctx context.Context, sid uint64, mid uint64, tok string, req GetBuyboxModelsByModelIdonlyForBRNowRequest) (*GetBuyboxModelsByModelIdonlyForBRNowResponse, error) {
	path := "/buybox/get_buybox_models_by_model_id"
	resp := new(GetBuyboxModelsByModelIdonlyForBRNowResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// GetBuyboxModelsByShopIdonlyForBRNow {"content":"<p>Get the paginated Buybox model list for the authorized shop.</p>","raw_content":[{"name":"paragraph","children":[{"data":"Get the paginated Buybox model list for the authorized shop."}]}]}
// Path: /api/v2/buybox/get_buybox_models_by_shop_id
// https://open.shopee.com/documents/v2/v2.buybox.get_buybox_models_by_shop_id（only for BR now）?module=140&type=1
func (s *BuyboxServiceOp[T]) GetBuyboxModelsByShopIdonlyForBRNow(ctx context.Context, sid uint64, mid uint64, tok string, req GetBuyboxModelsByShopIdonlyForBRNowRequest) (*GetBuyboxModelsByShopIdonlyForBRNowResponse, error) {
	path := "/buybox/get_buybox_models_by_shop_id"
	resp := new(GetBuyboxModelsByShopIdonlyForBRNowResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// GetBuyboxShopPerformanceonlyForBRNow {"content":"<p>Get Buybox shop performance for the authorized shop.</p>","raw_content":[{"name":"paragraph","children":[{"data":"Get Buybox shop performance for the authorized shop."}]}]}
// Path: /api/v2/buybox/get_buybox_shop_performance
// https://open.shopee.com/documents/v2/v2.buybox.get_buybox_shop_performance（only for BR now）?module=140&type=1
func (s *BuyboxServiceOp[T]) GetBuyboxShopPerformanceonlyForBRNow(ctx context.Context, sid uint64, mid uint64, tok string) (*GetBuyboxShopPerformanceonlyForBRNowResponse, error) {
	path := "/buybox/get_buybox_shop_performance"
	resp := new(GetBuyboxShopPerformanceonlyForBRNowResponse)
	err := s.client.Post(ctx, path, nil, resp, sid, mid, tok)
	return resp, err
}

// UpdateBuyboxModelEnrollmentonlyForBRNow {"content":"<p>Update Buybox model enrollment status for the authorized shop.</p>","raw_content":[{"name":"paragraph","children":[{"data":"Update Buybox model enrollment status for the authorized shop."}]}]}
// Path: /api/v2/buybox/update_buybox_model_enrollment
// https://open.shopee.com/documents/v2/v2.buybox.update_buybox_model_enrollment（only for BR now）?module=140&type=1
func (s *BuyboxServiceOp[T]) UpdateBuyboxModelEnrollmentonlyForBRNow(ctx context.Context, sid uint64, mid uint64, tok string, req UpdateBuyboxModelEnrollmentonlyForBRNowRequest) (*UpdateBuyboxModelEnrollmentonlyForBRNowResponse, error) {
	path := "/buybox/update_buybox_model_enrollment"
	resp := new(UpdateBuyboxModelEnrollmentonlyForBRNowResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}
