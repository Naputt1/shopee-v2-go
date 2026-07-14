package goshopee

import (
	"context"
	"io"
)

type LivestreamService interface {
	// AddItemList Add items to item bag. (For TW, ID, TH, PH, MY, SG, VN)
	// Path: /api/v2/livestream/add_item_list
	// https://open.shopee.com/documents/v2/v2.livestream.add_item_list?module=125&type=1
	AddItemList(ctx context.Context, sid uint64, mid uint64, tok string, req AddItemListRequest) (*AddItemListResponse, error)
	// ApplyItemSet Add product set to item bag. (For TW, ID, TH, PH, MY, SG, VN)
	// Path: /api/v2/livestream/apply_item_set
	// https://open.shopee.com/documents/v2/v2.livestream.apply_item_set?module=125&type=1
	ApplyItemSet(ctx context.Context, sid uint64, mid uint64, tok string, req ApplyItemSetRequest) (*ApplyItemSetResponse, error)
	// BanUserComment Ban the user from posting comments. (For TW, ID, TH, PH, MY, SG, VN)
	// Path: /api/v2/livestream/ban_user_comment
	// https://open.shopee.com/documents/v2/v2.livestream.ban_user_comment?module=125&type=1
	BanUserComment(ctx context.Context, sid uint64, mid uint64, tok string, req BanUserCommentRequest) (*BanUserCommentResponse, error)
	// CreateSession Create a new live stream, include basic information, like cover, title, description, type (test live or normal live). (For TW, ID, TH, PH, MY, SG, VN)
	// Path: /api/v2/livestream/create_session
	// https://open.shopee.com/documents/v2/v2.livestream.create_session?module=125&type=1
	CreateSession(ctx context.Context, sid uint64, mid uint64, tok string, req CreateSessionRequest) (*CreateSessionResponse, error)
	// DeleteItemList Delete items from item bag. (For TW, ID, TH, PH, MY, SG, VN)
	// Path: /api/v2/livestream/delete_item_list
	// https://open.shopee.com/documents/v2/v2.livestream.delete_item_list?module=125&type=1
	DeleteItemList(ctx context.Context, sid uint64, mid uint64, tok string, req DeleteItemListRequest) (*DeleteItemListResponse, error)
	// DeleteShowItem Unshow showing item. (For TW, ID, TH, PH, MY, SG, VN)
	// Path: /api/v2/livestream/delete_show_item
	// https://open.shopee.com/documents/v2/v2.livestream.delete_show_item?module=125&type=1
	DeleteShowItem(ctx context.Context, sid uint64, mid uint64, tok string, req DeleteShowItemRequest) (*DeleteShowItemResponse, error)
	// EndSession End Live. (For TW, ID, TH, PH, MY, SG, VN)
	// Path: /api/v2/livestream/end_session
	// https://open.shopee.com/documents/v2/v2.livestream.end_session?module=125&type=1
	EndSession(ctx context.Context, sid uint64, mid uint64, tok string, req EndSessionRequest) (*EndSessionResponse, error)
	// GetItemCount Get the number of items in the shopping bag, including the current number of items in the shopping bag, the upper limit of the number, etc. (For TW, ID, TH, PH, MY, SG, VN)
	// Path: /api/v2/livestream/get_item_count
	// https://open.shopee.com/documents/v2/v2.livestream.get_item_count?module=125&type=1
	GetItemCount(ctx context.Context, sid uint64, mid uint64, tok string, opt GetItemCountRequest) (*GetItemCountResponse, error)
	// GetItemList Get the detail information of item in item bag, including item id, item serial number, etc.(For TW, ID, TH, PH, MY, SG, VN)
	// Path: /api/v2/livestream/get_item_list
	// https://open.shopee.com/documents/v2/v2.livestream.get_item_list?module=125&type=1
	GetItemList(ctx context.Context, sid uint64, mid uint64, tok string, opt GetItemListRequest) (*GetItemListResponse, error)
	// GetItemSetItemList Get the item list of the product set, including item name, id, etc. (For TW, ID, TH, PH, MY, SG, VN)
	// Path: /api/v2/livestream/get_item_set_item_list
	// https://open.shopee.com/documents/v2/v2.livestream.get_item_set_item_list?module=125&type=1
	GetItemSetItemList(ctx context.Context, sid uint64, mid uint64, tok string, opt GetItemSetItemListRequest) (*GetItemSetItemListResponse, error)
	// GetItemSetList Get the product set of the live stream, including the product set name, id, and item number. (For TW, ID, TH, PH, MY, SG, VN)
	// Path: /api/v2/livestream/get_item_set_list
	// https://open.shopee.com/documents/v2/v2.livestream.get_item_set_list?module=125&type=1
	GetItemSetList(ctx context.Context, sid uint64, mid uint64, tok string, opt GetItemSetListRequest) (*GetItemSetListResponse, error)
	// GetLatestCommentList Get live stream room comments in the last 10 seconds, including user id, user name, comment id, comment content, and comment time. (For TW, ID, TH, PH, MY, SG, VN)
	// Path: /api/v2/livestream/get_latest_comment_list
	// https://open.shopee.com/documents/v2/v2.livestream.get_latest_comment_list?module=125&type=1
	GetLatestCommentList(ctx context.Context, sid uint64, mid uint64, tok string, opt GetLatestCommentListRequest) (*GetLatestCommentListResponse, error)
	// GetLikeItemList Get the item list of My Likes tab.(For TW, ID, TH, PH, MY, SG, VN)
	// Path: /api/v2/livestream/get_like_item_list
	// https://open.shopee.com/documents/v2/v2.livestream.get_like_item_list?module=125&type=1
	GetLikeItemList(ctx context.Context, sid uint64, mid uint64, tok string, opt GetLikeItemListRequest) (*GetLikeItemListResponse, error)
	// GetRecentItemList Get the item list of the Recently tab. (For TW, ID, TH, PH, MY, SG, VN)
	// Path: /api/v2/livestream/get_recent_item_list
	// https://open.shopee.com/documents/v2/v2.livestream.get_recent_item_list?module=125&type=1
	GetRecentItemList(ctx context.Context, sid uint64, mid uint64, tok string, opt GetRecentItemListRequest) (*GetRecentItemListResponse, error)
	// GetSessionDetail Get basic information about the live streaming room, including cover, title, description, type (test live or normal live), create time, update time, stream url, etc. (For TW, ID, TH, PH, MY, SG, VN)
	// Path: /api/v2/livestream/get_session_detail
	// https://open.shopee.com/documents/v2/v2.livestream.get_session_detail?module=125&type=1
	GetSessionDetail(ctx context.Context, sid uint64, mid uint64, tok string, opt GetSessionDetailRequest) (*GetSessionDetailResponse, error)
	// GetSessionItemMetric Get real-time indicator data of live stream products, including product clicks, add-to-cart, etc. (For TW, ID, TH, PH, MY, SG, VN)
	// Path: /api/v2/livestream/get_session_item_metric
	// https://open.shopee.com/documents/v2/v2.livestream.get_session_item_metric?module=125&type=1
	GetSessionItemMetric(ctx context.Context, sid uint64, mid uint64, tok string, opt GetSessionItemMetricRequest) (*GetSessionItemMetricResponse, error)
	// GetSessionMetric Get real-time indicator data of the live stream room, including the number of likes, comments, shares, views, etc.(For TW, ID, TH, PH, MY, SG, VN)
	// Path: /api/v2/livestream/get_session_metric
	// https://open.shopee.com/documents/v2/v2.livestream.get_session_metric?module=125&type=1
	GetSessionMetric(ctx context.Context, sid uint64, mid uint64, tok string, opt GetSessionMetricRequest) (*GetSessionMetricResponse, error)
	// GetShowItem Get the showing item. (For TW, ID, TH, PH, MY, SG, VN)
	// Path: /api/v2/livestream/get_show_item
	// https://open.shopee.com/documents/v2/v2.livestream.get_show_item?module=125&type=1
	GetShowItem(ctx context.Context, sid uint64, mid uint64, tok string, opt GetShowItemRequest) (*GetShowItemResponse, error)
	// PostComment Post comment in the live stream as streamer. (For TW, ID, TH, PH, MY, SG, VN)
	// Path: /api/v2/livestream/post_comment
	// https://open.shopee.com/documents/v2/v2.livestream.post_comment?module=125&type=1
	PostComment(ctx context.Context, sid uint64, mid uint64, tok string, req PostCommentRequest) (*PostCommentResponse, error)
	// StartSession Start Live. (For TW, ID, TH, PH, MY, SG, VN)
	// Path: /api/v2/livestream/start_session
	// https://open.shopee.com/documents/v2/v2.livestream.start_session?module=125&type=1
	StartSession(ctx context.Context, sid uint64, mid uint64, tok string, req StartSessionRequest) (*StartSessionResponse, error)
	// UnbanUserComment Unban a user from posting comments. (For TW, ID, TH, PH, MY, SG, VN)
	// Path: /api/v2/livestream/unban_user_comment
	// https://open.shopee.com/documents/v2/v2.livestream.unban_user_comment?module=125&type=1
	UnbanUserComment(ctx context.Context, sid uint64, mid uint64, tok string, req UnbanUserCommentRequest) (*UnbanUserCommentResponse, error)
	// UpdateItemList Update the order of items in item bag. (For TW, ID, TH, PH, MY, SG, VN)
	// Path: /api/v2/livestream/update_item_list
	// https://open.shopee.com/documents/v2/v2.livestream.update_item_list?module=125&type=1
	UpdateItemList(ctx context.Context, sid uint64, mid uint64, tok string, req UpdateItemListRequest) (*UpdateItemListResponse, error)
	// UpdateSession Update live stream information, including cover, title, description, and type (test live or normal live). (For TW, ID, TH, PH, MY, SG, VN)
	// Path: /api/v2/livestream/update_session
	// https://open.shopee.com/documents/v2/v2.livestream.update_session?module=125&type=1
	UpdateSession(ctx context.Context, sid uint64, mid uint64, tok string, req UpdateSessionRequest) (*UpdateSessionResponse, error)
	// UpdateShowItem Set the showing item. (For TW, ID, TH, PH, MY, SG, VN)
	// Path: /api/v2/livestream/update_show_item
	// https://open.shopee.com/documents/v2/v2.livestream.update_show_item?module=125&type=1
	UpdateShowItem(ctx context.Context, sid uint64, mid uint64, tok string, req UpdateShowItemRequest) (*UpdateShowItemResponse, error)
	// UploadImage Upload an image as the live stream cover.(For TW, ID, TH, PH, MY, SG, VN)
	// Path: /api/v2/livestream/upload_image
	// https://open.shopee.com/documents/v2/v2.livestream.upload_image?module=125&type=1
	UploadImage(ctx context.Context, sid uint64, mid uint64, tok string, filename string) (*UploadImageResponse, error)
	UploadImageFromReader(ctx context.Context, sid uint64, mid uint64, tok string, filename string, reader io.Reader) (*UploadImageResponse, error)
}

type LivestreamServiceOp[T any] struct {
	client *Client[T]
}

// AddItemList Add items to item bag. (For TW, ID, TH, PH, MY, SG, VN)
// Path: /api/v2/livestream/add_item_list
// https://open.shopee.com/documents/v2/v2.livestream.add_item_list?module=125&type=1
func (s *LivestreamServiceOp[T]) AddItemList(ctx context.Context, sid uint64, mid uint64, tok string, req AddItemListRequest) (*AddItemListResponse, error) {
	path := "/livestream/add_item_list"
	resp := new(AddItemListResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// ApplyItemSet Add product set to item bag. (For TW, ID, TH, PH, MY, SG, VN)
// Path: /api/v2/livestream/apply_item_set
// https://open.shopee.com/documents/v2/v2.livestream.apply_item_set?module=125&type=1
func (s *LivestreamServiceOp[T]) ApplyItemSet(ctx context.Context, sid uint64, mid uint64, tok string, req ApplyItemSetRequest) (*ApplyItemSetResponse, error) {
	path := "/livestream/apply_item_set"
	resp := new(ApplyItemSetResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// BanUserComment Ban the user from posting comments. (For TW, ID, TH, PH, MY, SG, VN)
// Path: /api/v2/livestream/ban_user_comment
// https://open.shopee.com/documents/v2/v2.livestream.ban_user_comment?module=125&type=1
func (s *LivestreamServiceOp[T]) BanUserComment(ctx context.Context, sid uint64, mid uint64, tok string, req BanUserCommentRequest) (*BanUserCommentResponse, error) {
	path := "/livestream/ban_user_comment"
	resp := new(BanUserCommentResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// CreateSession Create a new live stream, include basic information, like cover, title, description, type (test live or normal live). (For TW, ID, TH, PH, MY, SG, VN)
// Path: /api/v2/livestream/create_session
// https://open.shopee.com/documents/v2/v2.livestream.create_session?module=125&type=1
func (s *LivestreamServiceOp[T]) CreateSession(ctx context.Context, sid uint64, mid uint64, tok string, req CreateSessionRequest) (*CreateSessionResponse, error) {
	path := "/livestream/create_session"
	resp := new(CreateSessionResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// DeleteItemList Delete items from item bag. (For TW, ID, TH, PH, MY, SG, VN)
// Path: /api/v2/livestream/delete_item_list
// https://open.shopee.com/documents/v2/v2.livestream.delete_item_list?module=125&type=1
func (s *LivestreamServiceOp[T]) DeleteItemList(ctx context.Context, sid uint64, mid uint64, tok string, req DeleteItemListRequest) (*DeleteItemListResponse, error) {
	path := "/livestream/delete_item_list"
	resp := new(DeleteItemListResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// DeleteShowItem Unshow showing item. (For TW, ID, TH, PH, MY, SG, VN)
// Path: /api/v2/livestream/delete_show_item
// https://open.shopee.com/documents/v2/v2.livestream.delete_show_item?module=125&type=1
func (s *LivestreamServiceOp[T]) DeleteShowItem(ctx context.Context, sid uint64, mid uint64, tok string, req DeleteShowItemRequest) (*DeleteShowItemResponse, error) {
	path := "/livestream/delete_show_item"
	resp := new(DeleteShowItemResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// EndSession End Live. (For TW, ID, TH, PH, MY, SG, VN)
// Path: /api/v2/livestream/end_session
// https://open.shopee.com/documents/v2/v2.livestream.end_session?module=125&type=1
func (s *LivestreamServiceOp[T]) EndSession(ctx context.Context, sid uint64, mid uint64, tok string, req EndSessionRequest) (*EndSessionResponse, error) {
	path := "/livestream/end_session"
	resp := new(EndSessionResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// GetItemCount Get the number of items in the shopping bag, including the current number of items in the shopping bag, the upper limit of the number, etc. (For TW, ID, TH, PH, MY, SG, VN)
// Path: /api/v2/livestream/get_item_count
// https://open.shopee.com/documents/v2/v2.livestream.get_item_count?module=125&type=1
func (s *LivestreamServiceOp[T]) GetItemCount(ctx context.Context, sid uint64, mid uint64, tok string, opt GetItemCountRequest) (*GetItemCountResponse, error) {
	path := "/livestream/get_item_count"
	resp := new(GetItemCountResponse)
	err := s.client.Get(ctx, path, resp, opt, sid, mid, tok)
	return resp, err
}

// GetItemList Get the detail information of item in item bag, including item id, item serial number, etc.(For TW, ID, TH, PH, MY, SG, VN)
// Path: /api/v2/livestream/get_item_list
// https://open.shopee.com/documents/v2/v2.livestream.get_item_list?module=125&type=1
func (s *LivestreamServiceOp[T]) GetItemList(ctx context.Context, sid uint64, mid uint64, tok string, opt GetItemListRequest) (*GetItemListResponse, error) {
	path := "/livestream/get_item_list"
	resp := new(GetItemListResponse)
	err := s.client.Get(ctx, path, resp, opt, sid, mid, tok)
	return resp, err
}

// GetItemSetItemList Get the item list of the product set, including item name, id, etc. (For TW, ID, TH, PH, MY, SG, VN)
// Path: /api/v2/livestream/get_item_set_item_list
// https://open.shopee.com/documents/v2/v2.livestream.get_item_set_item_list?module=125&type=1
func (s *LivestreamServiceOp[T]) GetItemSetItemList(ctx context.Context, sid uint64, mid uint64, tok string, opt GetItemSetItemListRequest) (*GetItemSetItemListResponse, error) {
	path := "/livestream/get_item_set_item_list"
	resp := new(GetItemSetItemListResponse)
	err := s.client.Get(ctx, path, resp, opt, sid, mid, tok)
	return resp, err
}

// GetItemSetList Get the product set of the live stream, including the product set name, id, and item number. (For TW, ID, TH, PH, MY, SG, VN)
// Path: /api/v2/livestream/get_item_set_list
// https://open.shopee.com/documents/v2/v2.livestream.get_item_set_list?module=125&type=1
func (s *LivestreamServiceOp[T]) GetItemSetList(ctx context.Context, sid uint64, mid uint64, tok string, opt GetItemSetListRequest) (*GetItemSetListResponse, error) {
	path := "/livestream/get_item_set_list"
	resp := new(GetItemSetListResponse)
	err := s.client.Get(ctx, path, resp, opt, sid, mid, tok)
	return resp, err
}

// GetLatestCommentList Get live stream room comments in the last 10 seconds, including user id, user name, comment id, comment content, and comment time. (For TW, ID, TH, PH, MY, SG, VN)
// Path: /api/v2/livestream/get_latest_comment_list
// https://open.shopee.com/documents/v2/v2.livestream.get_latest_comment_list?module=125&type=1
func (s *LivestreamServiceOp[T]) GetLatestCommentList(ctx context.Context, sid uint64, mid uint64, tok string, opt GetLatestCommentListRequest) (*GetLatestCommentListResponse, error) {
	path := "/livestream/get_latest_comment_list"
	resp := new(GetLatestCommentListResponse)
	err := s.client.Get(ctx, path, resp, opt, sid, mid, tok)
	return resp, err
}

// GetLikeItemList Get the item list of My Likes tab.(For TW, ID, TH, PH, MY, SG, VN)
// Path: /api/v2/livestream/get_like_item_list
// https://open.shopee.com/documents/v2/v2.livestream.get_like_item_list?module=125&type=1
func (s *LivestreamServiceOp[T]) GetLikeItemList(ctx context.Context, sid uint64, mid uint64, tok string, opt GetLikeItemListRequest) (*GetLikeItemListResponse, error) {
	path := "/livestream/get_like_item_list"
	resp := new(GetLikeItemListResponse)
	err := s.client.Get(ctx, path, resp, opt, sid, mid, tok)
	return resp, err
}

// GetRecentItemList Get the item list of the Recently tab. (For TW, ID, TH, PH, MY, SG, VN)
// Path: /api/v2/livestream/get_recent_item_list
// https://open.shopee.com/documents/v2/v2.livestream.get_recent_item_list?module=125&type=1
func (s *LivestreamServiceOp[T]) GetRecentItemList(ctx context.Context, sid uint64, mid uint64, tok string, opt GetRecentItemListRequest) (*GetRecentItemListResponse, error) {
	path := "/livestream/get_recent_item_list"
	resp := new(GetRecentItemListResponse)
	err := s.client.Get(ctx, path, resp, opt, sid, mid, tok)
	return resp, err
}

// GetSessionDetail Get basic information about the live streaming room, including cover, title, description, type (test live or normal live), create time, update time, stream url, etc. (For TW, ID, TH, PH, MY, SG, VN)
// Path: /api/v2/livestream/get_session_detail
// https://open.shopee.com/documents/v2/v2.livestream.get_session_detail?module=125&type=1
func (s *LivestreamServiceOp[T]) GetSessionDetail(ctx context.Context, sid uint64, mid uint64, tok string, opt GetSessionDetailRequest) (*GetSessionDetailResponse, error) {
	path := "/livestream/get_session_detail"
	resp := new(GetSessionDetailResponse)
	err := s.client.Get(ctx, path, resp, opt, sid, mid, tok)
	return resp, err
}

// GetSessionItemMetric Get real-time indicator data of live stream products, including product clicks, add-to-cart, etc. (For TW, ID, TH, PH, MY, SG, VN)
// Path: /api/v2/livestream/get_session_item_metric
// https://open.shopee.com/documents/v2/v2.livestream.get_session_item_metric?module=125&type=1
func (s *LivestreamServiceOp[T]) GetSessionItemMetric(ctx context.Context, sid uint64, mid uint64, tok string, opt GetSessionItemMetricRequest) (*GetSessionItemMetricResponse, error) {
	path := "/livestream/get_session_item_metric"
	resp := new(GetSessionItemMetricResponse)
	err := s.client.Get(ctx, path, resp, opt, sid, mid, tok)
	return resp, err
}

// GetSessionMetric Get real-time indicator data of the live stream room, including the number of likes, comments, shares, views, etc.(For TW, ID, TH, PH, MY, SG, VN)
// Path: /api/v2/livestream/get_session_metric
// https://open.shopee.com/documents/v2/v2.livestream.get_session_metric?module=125&type=1
func (s *LivestreamServiceOp[T]) GetSessionMetric(ctx context.Context, sid uint64, mid uint64, tok string, opt GetSessionMetricRequest) (*GetSessionMetricResponse, error) {
	path := "/livestream/get_session_metric"
	resp := new(GetSessionMetricResponse)
	err := s.client.Get(ctx, path, resp, opt, sid, mid, tok)
	return resp, err
}

// GetShowItem Get the showing item. (For TW, ID, TH, PH, MY, SG, VN)
// Path: /api/v2/livestream/get_show_item
// https://open.shopee.com/documents/v2/v2.livestream.get_show_item?module=125&type=1
func (s *LivestreamServiceOp[T]) GetShowItem(ctx context.Context, sid uint64, mid uint64, tok string, opt GetShowItemRequest) (*GetShowItemResponse, error) {
	path := "/livestream/get_show_item"
	resp := new(GetShowItemResponse)
	err := s.client.Get(ctx, path, resp, opt, sid, mid, tok)
	return resp, err
}

// PostComment Post comment in the live stream as streamer. (For TW, ID, TH, PH, MY, SG, VN)
// Path: /api/v2/livestream/post_comment
// https://open.shopee.com/documents/v2/v2.livestream.post_comment?module=125&type=1
func (s *LivestreamServiceOp[T]) PostComment(ctx context.Context, sid uint64, mid uint64, tok string, req PostCommentRequest) (*PostCommentResponse, error) {
	path := "/livestream/post_comment"
	resp := new(PostCommentResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// StartSession Start Live. (For TW, ID, TH, PH, MY, SG, VN)
// Path: /api/v2/livestream/start_session
// https://open.shopee.com/documents/v2/v2.livestream.start_session?module=125&type=1
func (s *LivestreamServiceOp[T]) StartSession(ctx context.Context, sid uint64, mid uint64, tok string, req StartSessionRequest) (*StartSessionResponse, error) {
	path := "/livestream/start_session"
	resp := new(StartSessionResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// UnbanUserComment Unban a user from posting comments. (For TW, ID, TH, PH, MY, SG, VN)
// Path: /api/v2/livestream/unban_user_comment
// https://open.shopee.com/documents/v2/v2.livestream.unban_user_comment?module=125&type=1
func (s *LivestreamServiceOp[T]) UnbanUserComment(ctx context.Context, sid uint64, mid uint64, tok string, req UnbanUserCommentRequest) (*UnbanUserCommentResponse, error) {
	path := "/livestream/unban_user_comment"
	resp := new(UnbanUserCommentResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// UpdateItemList Update the order of items in item bag. (For TW, ID, TH, PH, MY, SG, VN)
// Path: /api/v2/livestream/update_item_list
// https://open.shopee.com/documents/v2/v2.livestream.update_item_list?module=125&type=1
func (s *LivestreamServiceOp[T]) UpdateItemList(ctx context.Context, sid uint64, mid uint64, tok string, req UpdateItemListRequest) (*UpdateItemListResponse, error) {
	path := "/livestream/update_item_list"
	resp := new(UpdateItemListResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// UpdateSession Update live stream information, including cover, title, description, and type (test live or normal live). (For TW, ID, TH, PH, MY, SG, VN)
// Path: /api/v2/livestream/update_session
// https://open.shopee.com/documents/v2/v2.livestream.update_session?module=125&type=1
func (s *LivestreamServiceOp[T]) UpdateSession(ctx context.Context, sid uint64, mid uint64, tok string, req UpdateSessionRequest) (*UpdateSessionResponse, error) {
	path := "/livestream/update_session"
	resp := new(UpdateSessionResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// UpdateShowItem Set the showing item. (For TW, ID, TH, PH, MY, SG, VN)
// Path: /api/v2/livestream/update_show_item
// https://open.shopee.com/documents/v2/v2.livestream.update_show_item?module=125&type=1
func (s *LivestreamServiceOp[T]) UpdateShowItem(ctx context.Context, sid uint64, mid uint64, tok string, req UpdateShowItemRequest) (*UpdateShowItemResponse, error) {
	path := "/livestream/update_show_item"
	resp := new(UpdateShowItemResponse)
	err := s.client.Post(ctx, path, req, resp, sid, mid, tok)
	return resp, err
}

// UploadImage Upload an image as the live stream cover.(For TW, ID, TH, PH, MY, SG, VN)
// Path: /api/v2/livestream/upload_image
// https://open.shopee.com/documents/v2/v2.livestream.upload_image?module=125&type=1
func (s *LivestreamServiceOp[T]) UploadImage(ctx context.Context, sid uint64, mid uint64, tok string, filename string) (*UploadImageResponse, error) {
	path := "/livestream/upload_image"
	resp := new(UploadImageResponse)
	err := s.client.Upload(ctx, path, "image", filename, resp, sid, mid, tok)
	return resp, err
}

func (s *LivestreamServiceOp[T]) UploadImageFromReader(ctx context.Context, sid uint64, mid uint64, tok string, filename string, reader io.Reader) (*UploadImageResponse, error) {
	path := "/livestream/upload_image"
	resp := new(UploadImageResponse)
	err := s.client.UploadFromReader(ctx, path, "image", filename, reader, resp, sid, mid, tok)
	return resp, err
}
