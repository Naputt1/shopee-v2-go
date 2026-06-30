package goshopee

import (
	"context"
)

type TopPicksService interface {
	// AddTopPicks add one collection
	// Path: /api/v2/top_picks/add_top_picks
	// https://open.shopee.com/documents/v2/v2.top_picks.add_top_picks?module=100&type=1
	AddTopPicks(ctx context.Context, sid uint64, req AddTopPicksRequest, tok string) (*AddTopPicksResponse, error)
	// DeleteTopPicks delete a collection
	// Path: /api/v2/top_picks/delete_top_picks
	// https://open.shopee.com/documents/v2/v2.top_picks.delete_top_picks?module=100&type=1
	DeleteTopPicks(ctx context.Context, sid uint64, req DeleteTopPicksRequest, tok string) (*DeleteTopPicksResponse, error)
	// GetTopPicksList get one TopPicks
	// Path: /api/v2/top_picks/get_top_picks_list
	// https://open.shopee.com/documents/v2/v2.top_picks.get_top_picks_list?module=100&type=1
	GetTopPicksList(ctx context.Context, sid uint64, tok string) (*GetTopPicksListResponse, error)
	// UpdateTopPicks update a collection info
	// Path: /api/v2/top_picks/update_top_picks
	// https://open.shopee.com/documents/v2/v2.top_picks.update_top_picks?module=100&type=1
	UpdateTopPicks(ctx context.Context, sid uint64, req UpdateTopPicksRequest, tok string) (*UpdateTopPicksResponse, error)
}

type TopPicksServiceOp[T any] struct {
	client *Client[T]
}

// AddTopPicks add one collection
// Path: /api/v2/top_picks/add_top_picks
// https://open.shopee.com/documents/v2/v2.top_picks.add_top_picks?module=100&type=1
func (s *TopPicksServiceOp[T]) AddTopPicks(ctx context.Context, sid uint64, req AddTopPicksRequest, tok string) (*AddTopPicksResponse, error) {
	path := "/top_picks/add_top_picks"
	resp := new(AddTopPicksResponse)
	err := s.client.Post(ctx, path, req, resp, sid, tok)
	return resp, err
}

// DeleteTopPicks delete a collection
// Path: /api/v2/top_picks/delete_top_picks
// https://open.shopee.com/documents/v2/v2.top_picks.delete_top_picks?module=100&type=1
func (s *TopPicksServiceOp[T]) DeleteTopPicks(ctx context.Context, sid uint64, req DeleteTopPicksRequest, tok string) (*DeleteTopPicksResponse, error) {
	path := "/top_picks/delete_top_picks"
	resp := new(DeleteTopPicksResponse)
	err := s.client.Post(ctx, path, req, resp, sid, tok)
	return resp, err
}

// GetTopPicksList get one TopPicks
// Path: /api/v2/top_picks/get_top_picks_list
// https://open.shopee.com/documents/v2/v2.top_picks.get_top_picks_list?module=100&type=1
func (s *TopPicksServiceOp[T]) GetTopPicksList(ctx context.Context, sid uint64, tok string) (*GetTopPicksListResponse, error) {
	path := "/top_picks/get_top_picks_list"
	resp := new(GetTopPicksListResponse)
	err := s.client.Post(ctx, path, nil, resp, sid, tok)
	return resp, err
}

// UpdateTopPicks update a collection info
// Path: /api/v2/top_picks/update_top_picks
// https://open.shopee.com/documents/v2/v2.top_picks.update_top_picks?module=100&type=1
func (s *TopPicksServiceOp[T]) UpdateTopPicks(ctx context.Context, sid uint64, req UpdateTopPicksRequest, tok string) (*UpdateTopPicksResponse, error) {
	path := "/top_picks/update_top_picks"
	resp := new(UpdateTopPicksResponse)
	err := s.client.Post(ctx, path, req, resp, sid, tok)
	return resp, err
}
