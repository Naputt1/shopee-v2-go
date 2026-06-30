package goshopee

import (
	"context"
)

type FollowPrizeService interface {
	// AddFollowPrize OpenAPI add Follow Prize
	// Path: /api/v2/follow_prize/add_follow_prize
	// https://open.shopee.com/documents/v2/v2.follow_prize.add_follow_prize?module=113&type=1
	AddFollowPrize(ctx context.Context, sid uint64, req AddFollowPrizeRequest, tok string) (*AddFollowPrizeResponse, error)
	// DeleteFollowPrize delete_follow_prize
	// Path: /api/v2/follow_prize/delete_follow_prize
	// https://open.shopee.com/documents/v2/v2.follow_prize.delete_follow_prize?module=113&type=1
	DeleteFollowPrize(ctx context.Context, sid uint64, req DeleteFollowPrizeRequest, tok string) (*DeleteFollowPrizeResponse, error)
	// EndFollowPrize end follow prize
	// Path: /api/v2/follow_prize/end_follow_prize
	// https://open.shopee.com/documents/v2/v2.follow_prize.end_follow_prize?module=113&type=1
	EndFollowPrize(ctx context.Context, sid uint64, req EndFollowPrizeRequest, tok string) (*EndFollowPrizeResponse, error)
	// GetFollowPrizeDetail get_follow_prize_detail
	// Path: /api/v2/follow_prize/get_follow_prize_detail
	// https://open.shopee.com/documents/v2/v2.follow_prize.get_follow_prize_detail?module=113&type=1
	GetFollowPrizeDetail(ctx context.Context, sid uint64, req GetFollowPrizeDetailRequest, tok string) (*GetFollowPrizeDetailResponse, error)
	// GetFollowPrizeList OpenAPI get_follow_prize_list
	// Path: /api/v2/follow_prize/get_follow_prize_list
	// https://open.shopee.com/documents/v2/v2.follow_prize.get_follow_prize_list?module=113&type=1
	GetFollowPrizeList(ctx context.Context, sid uint64, req GetFollowPrizeListRequest, tok string) (*GetFollowPrizeListResponse, error)
	// UpdateFollowPrize update_follow_prize
	// Path: /api/v2/follow_prize/update_follow_prize
	// https://open.shopee.com/documents/v2/v2.follow_prize.update_follow_prize?module=113&type=1
	UpdateFollowPrize(ctx context.Context, sid uint64, req UpdateFollowPrizeRequest, tok string) (*UpdateFollowPrizeResponse, error)
}

type FollowPrizeServiceOp[T any] struct {
	client *Client[T]
}

// AddFollowPrize OpenAPI add Follow Prize
// Path: /api/v2/follow_prize/add_follow_prize
// https://open.shopee.com/documents/v2/v2.follow_prize.add_follow_prize?module=113&type=1
func (s *FollowPrizeServiceOp[T]) AddFollowPrize(ctx context.Context, sid uint64, req AddFollowPrizeRequest, tok string) (*AddFollowPrizeResponse, error) {
	path := "/follow_prize/add_follow_prize"
	resp := new(AddFollowPrizeResponse)
	err := s.client.WithShop(sid, tok).Post(ctx, path, req, resp)
	return resp, err
}

// DeleteFollowPrize delete_follow_prize
// Path: /api/v2/follow_prize/delete_follow_prize
// https://open.shopee.com/documents/v2/v2.follow_prize.delete_follow_prize?module=113&type=1
func (s *FollowPrizeServiceOp[T]) DeleteFollowPrize(ctx context.Context, sid uint64, req DeleteFollowPrizeRequest, tok string) (*DeleteFollowPrizeResponse, error) {
	path := "/follow_prize/delete_follow_prize"
	resp := new(DeleteFollowPrizeResponse)
	err := s.client.WithShop(sid, tok).Post(ctx, path, req, resp)
	return resp, err
}

// EndFollowPrize end follow prize
// Path: /api/v2/follow_prize/end_follow_prize
// https://open.shopee.com/documents/v2/v2.follow_prize.end_follow_prize?module=113&type=1
func (s *FollowPrizeServiceOp[T]) EndFollowPrize(ctx context.Context, sid uint64, req EndFollowPrizeRequest, tok string) (*EndFollowPrizeResponse, error) {
	path := "/follow_prize/end_follow_prize"
	resp := new(EndFollowPrizeResponse)
	err := s.client.WithShop(sid, tok).Post(ctx, path, req, resp)
	return resp, err
}

// GetFollowPrizeDetail get_follow_prize_detail
// Path: /api/v2/follow_prize/get_follow_prize_detail
// https://open.shopee.com/documents/v2/v2.follow_prize.get_follow_prize_detail?module=113&type=1
func (s *FollowPrizeServiceOp[T]) GetFollowPrizeDetail(ctx context.Context, sid uint64, req GetFollowPrizeDetailRequest, tok string) (*GetFollowPrizeDetailResponse, error) {
	path := "/follow_prize/get_follow_prize_detail"
	resp := new(GetFollowPrizeDetailResponse)
	err := s.client.WithShop(sid, tok).Post(ctx, path, req, resp)
	return resp, err
}

// GetFollowPrizeList OpenAPI get_follow_prize_list
// Path: /api/v2/follow_prize/get_follow_prize_list
// https://open.shopee.com/documents/v2/v2.follow_prize.get_follow_prize_list?module=113&type=1
func (s *FollowPrizeServiceOp[T]) GetFollowPrizeList(ctx context.Context, sid uint64, req GetFollowPrizeListRequest, tok string) (*GetFollowPrizeListResponse, error) {
	path := "/follow_prize/get_follow_prize_list"
	resp := new(GetFollowPrizeListResponse)
	err := s.client.WithShop(sid, tok).Post(ctx, path, req, resp)
	return resp, err
}

// UpdateFollowPrize update_follow_prize
// Path: /api/v2/follow_prize/update_follow_prize
// https://open.shopee.com/documents/v2/v2.follow_prize.update_follow_prize?module=113&type=1
func (s *FollowPrizeServiceOp[T]) UpdateFollowPrize(ctx context.Context, sid uint64, req UpdateFollowPrizeRequest, tok string) (*UpdateFollowPrizeResponse, error) {
	path := "/follow_prize/update_follow_prize"
	resp := new(UpdateFollowPrizeResponse)
	err := s.client.WithShop(sid, tok).Post(ctx, path, req, resp)
	return resp, err
}
