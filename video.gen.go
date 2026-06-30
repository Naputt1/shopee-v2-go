package goshopee

type VideoService interface {
	// DeleteVideo Use this API to delete video. You can delete the video for both draft and post status.
	// Path: /api/v2/video/delete_video
	// https://open.shopee.com/documents/v2/v2.video.delete_video?module=129&type=1
	DeleteVideo(sid uint64, req DeleteVideoRequest, tok string) (*DeleteVideoResponse, error)
	// EditVideoInfo You need to call v2.media.init_video_upload, v2.media.upload_video_part, and v2.media.complete_video_upload to upload the video, and call the v2.media.get_video_upload_result to get the video_upload_id of uploaded video first, then call this API to set video post information. After submit, the video is still draft status, you need to call v2.video.post_video to post the video to Shopee Video. You can only set and update post information before the video is post.
	// Path: /api/v2/video/edit_video_info
	// https://open.shopee.com/documents/v2/v2.video.edit_video_info?module=129&type=1
	EditVideoInfo(sid uint64, req EditVideoInfoRequest, tok string) (*EditVideoInfoResponse, error)
	// GetCoverList You need to call v2.media.init_video_upload, v2.media.upload_video_part, and v2.media.complete_video_upload to upload the video, and call the v2.media.get_video_upload_result to get the video_upload_id of uploaded video. After the video is uploaded, obtain the frame-by-frame results and select a specific frame as the video cover.
	// Path: /api/v2/video/get_cover_list
	// https://open.shopee.com/documents/v2/v2.video.get_cover_list?module=129&type=1
	GetCoverList(sid uint64, opt GetCoverListRequest, tok string) (*GetCoverListResponse, error)
	// GetMetricTrend Query video data indicator trends.
	// Path: /api/v2/video/get_metric_trend
	// https://open.shopee.com/documents/v2/v2.video.get_metric_trend?module=129&type=1
	GetMetricTrend(sid uint64, opt GetMetricTrendRequest, tok string) (*GetMetricTrendResponse, error)
	// GetOverviewPerformance Get overall performance data for all post Shopee Video. There is at least a one-day delay.
	// Path: /api/v2/video/get_overview_performance
	// https://open.shopee.com/documents/v2/v2.video.get_overview_performance?module=129&type=1
	GetOverviewPerformance(sid uint64, opt GetOverviewPerformanceRequest, tok string) (*GetOverviewPerformanceResponse, error)
	// GetProdcutPerformanceList Get specific performance data for products linked with Shopee Video. There is at least a one-day delay.
	// Path: /api/v2/video/get_prodcut_performance_list
	// https://open.shopee.com/documents/v2/v2.video.get_prodcut_performance_list?module=129&type=1
	GetProdcutPerformanceList(sid uint64, opt GetProdcutPerformanceListRequest, tok string) (*GetProdcutPerformanceListResponse, error)
	// GetUserDemographics Get user demographics data to better understand the types of viewers that watch your Shopee Video.
	// Path: /api/v2/video/get_user_demographics
	// https://open.shopee.com/documents/v2/v2.video.get_user_demographics?module=129&type=1
	GetUserDemographics(sid uint64, tok string) (*GetUserDemographicsResponse, error)
	// GetVideoDetail Get the detail information of video.
	// Path: /api/v2/video/get_video_detail
	// https://open.shopee.com/documents/v2/v2.video.get_video_detail?module=129&type=1
	GetVideoDetail(sid uint64, opt GetVideoDetailRequest, tok string) (*GetVideoDetailResponse, error)
	// GetVideoDetailAudienceDistribution Get detailed audience distribution data for individual post Shopee Video. There is at least a one-day delay.
	// Path: /api/v2/video/get_video_detail_audience_distribution
	// https://open.shopee.com/documents/v2/v2.video.get_video_detail_audience_distribution?module=129&type=1
	GetVideoDetailAudienceDistribution(sid uint64, opt GetVideoDetailAudienceDistributionRequest, tok string) (*GetVideoDetailAudienceDistributionResponse, error)
	// GetVideoDetailMetricTrend Get detailed metric trend data for individual post Shopee Video. There is at least a one-day delay.
	// Path: /api/v2/video/get_video_detail_metric_trend
	// https://open.shopee.com/documents/v2/v2.video.get_video_detail_metric_trend?module=129&type=1
	GetVideoDetailMetricTrend(sid uint64, opt GetVideoDetailMetricTrendRequest, tok string) (*GetVideoDetailMetricTrendResponse, error)
	// GetVideoDetailPerformance Get detailed performance data for individual post Shopee Video. There is at least a one-day delay.
	// Path: /api/v2/video/get_video_detail_performance
	// https://open.shopee.com/documents/v2/v2.video.get_video_detail_performance?module=129&type=1
	GetVideoDetailPerformance(sid uint64, opt GetVideoDetailPerformanceRequest, tok string) (*GetVideoDetailPerformanceResponse, error)
	// GetVideoDetailProductPerformance Get performance data for products linked with individual post Shopee Video. There is at least a one-day delay.
	// Path: /api/v2/video/get_video_detail_product_performance
	// https://open.shopee.com/documents/v2/v2.video.get_video_detail_product_performance?module=129&type=1
	GetVideoDetailProductPerformance(sid uint64, opt GetVideoDetailProductPerformanceRequest, tok string) (*GetVideoDetailProductPerformanceResponse, error)
	// GetVideoList Get the list of video in draft status or video already post to Shopee Video.
	// Path: /api/v2/video/get_video_list
	// https://open.shopee.com/documents/v2/v2.video.get_video_list?module=129&type=1
	GetVideoList(sid uint64, opt GetVideoListRequest, tok string) (*GetVideoListResponse, error)
	// GetVideoPerformanceList Get specific performance data for individual post Shopee Video. There is at least a one-day delay.
	// Path: /api/v2/video/get_video_performance_list
	// https://open.shopee.com/documents/v2/v2.video.get_video_performance_list?module=129&type=1
	GetVideoPerformanceList(sid uint64, opt GetVideoPerformanceListRequest, tok string) (*GetVideoPerformanceListResponse, error)
	// PostVideo You need to call v2.media.init_video_upload, v2.media.upload_video_part, and v2.media.complete_video_upload to upload the video, then call the v2.video.edit_video_info API to set video post information, finally call this API to post the video to Shopee Video.
	// Path: /api/v2/video/post_video
	// https://open.shopee.com/documents/v2/v2.video.post_video?module=129&type=1
	PostVideo(sid uint64, req PostVideoRequest, tok string) (*PostVideoResponse, error)
}

type VideoServiceOp[T any] struct {
	client *Client[T]
}

// DeleteVideo Use this API to delete video. You can delete the video for both draft and post status.
// Path: /api/v2/video/delete_video
// https://open.shopee.com/documents/v2/v2.video.delete_video?module=129&type=1
func (s *VideoServiceOp[T]) DeleteVideo(sid uint64, req DeleteVideoRequest, tok string) (*DeleteVideoResponse, error) {
	path := "/video/delete_video"
	resp := new(DeleteVideoResponse)
	err := s.client.Post(path, req, resp)
	return resp, err
}

// EditVideoInfo You need to call v2.media.init_video_upload, v2.media.upload_video_part, and v2.media.complete_video_upload to upload the video, and call the v2.media.get_video_upload_result to get the video_upload_id of uploaded video first, then call this API to set video post information. After submit, the video is still draft status, you need to call v2.video.post_video to post the video to Shopee Video. You can only set and update post information before the video is post.
// Path: /api/v2/video/edit_video_info
// https://open.shopee.com/documents/v2/v2.video.edit_video_info?module=129&type=1
func (s *VideoServiceOp[T]) EditVideoInfo(sid uint64, req EditVideoInfoRequest, tok string) (*EditVideoInfoResponse, error) {
	path := "/video/edit_video_info"
	resp := new(EditVideoInfoResponse)
	err := s.client.Post(path, req, resp)
	return resp, err
}

// GetCoverList You need to call v2.media.init_video_upload, v2.media.upload_video_part, and v2.media.complete_video_upload to upload the video, and call the v2.media.get_video_upload_result to get the video_upload_id of uploaded video. After the video is uploaded, obtain the frame-by-frame results and select a specific frame as the video cover.
// Path: /api/v2/video/get_cover_list
// https://open.shopee.com/documents/v2/v2.video.get_cover_list?module=129&type=1
func (s *VideoServiceOp[T]) GetCoverList(sid uint64, opt GetCoverListRequest, tok string) (*GetCoverListResponse, error) {
	path := "/video/get_cover_list"
	resp := new(GetCoverListResponse)
	err := s.client.Get(path, resp, opt)
	return resp, err
}

// GetMetricTrend Query video data indicator trends.
// Path: /api/v2/video/get_metric_trend
// https://open.shopee.com/documents/v2/v2.video.get_metric_trend?module=129&type=1
func (s *VideoServiceOp[T]) GetMetricTrend(sid uint64, opt GetMetricTrendRequest, tok string) (*GetMetricTrendResponse, error) {
	path := "/video/get_metric_trend"
	resp := new(GetMetricTrendResponse)
	err := s.client.Get(path, resp, opt)
	return resp, err
}

// GetOverviewPerformance Get overall performance data for all post Shopee Video. There is at least a one-day delay.
// Path: /api/v2/video/get_overview_performance
// https://open.shopee.com/documents/v2/v2.video.get_overview_performance?module=129&type=1
func (s *VideoServiceOp[T]) GetOverviewPerformance(sid uint64, opt GetOverviewPerformanceRequest, tok string) (*GetOverviewPerformanceResponse, error) {
	path := "/video/get_overview_performance"
	resp := new(GetOverviewPerformanceResponse)
	err := s.client.Get(path, resp, opt)
	return resp, err
}

// GetProdcutPerformanceList Get specific performance data for products linked with Shopee Video. There is at least a one-day delay.
// Path: /api/v2/video/get_prodcut_performance_list
// https://open.shopee.com/documents/v2/v2.video.get_prodcut_performance_list?module=129&type=1
func (s *VideoServiceOp[T]) GetProdcutPerformanceList(sid uint64, opt GetProdcutPerformanceListRequest, tok string) (*GetProdcutPerformanceListResponse, error) {
	path := "/video/get_prodcut_performance_list"
	resp := new(GetProdcutPerformanceListResponse)
	err := s.client.Get(path, resp, opt)
	return resp, err
}

// GetUserDemographics Get user demographics data to better understand the types of viewers that watch your Shopee Video.
// Path: /api/v2/video/get_user_demographics
// https://open.shopee.com/documents/v2/v2.video.get_user_demographics?module=129&type=1
func (s *VideoServiceOp[T]) GetUserDemographics(sid uint64, tok string) (*GetUserDemographicsResponse, error) {
	path := "/video/get_user_demographics"
	resp := new(GetUserDemographicsResponse)
	err := s.client.Get(path, resp, nil)
	return resp, err
}

// GetVideoDetail Get the detail information of video.
// Path: /api/v2/video/get_video_detail
// https://open.shopee.com/documents/v2/v2.video.get_video_detail?module=129&type=1
func (s *VideoServiceOp[T]) GetVideoDetail(sid uint64, opt GetVideoDetailRequest, tok string) (*GetVideoDetailResponse, error) {
	path := "/video/get_video_detail"
	resp := new(GetVideoDetailResponse)
	err := s.client.Get(path, resp, opt)
	return resp, err
}

// GetVideoDetailAudienceDistribution Get detailed audience distribution data for individual post Shopee Video. There is at least a one-day delay.
// Path: /api/v2/video/get_video_detail_audience_distribution
// https://open.shopee.com/documents/v2/v2.video.get_video_detail_audience_distribution?module=129&type=1
func (s *VideoServiceOp[T]) GetVideoDetailAudienceDistribution(sid uint64, opt GetVideoDetailAudienceDistributionRequest, tok string) (*GetVideoDetailAudienceDistributionResponse, error) {
	path := "/video/get_video_detail_audience_distribution"
	resp := new(GetVideoDetailAudienceDistributionResponse)
	err := s.client.Get(path, resp, opt)
	return resp, err
}

// GetVideoDetailMetricTrend Get detailed metric trend data for individual post Shopee Video. There is at least a one-day delay.
// Path: /api/v2/video/get_video_detail_metric_trend
// https://open.shopee.com/documents/v2/v2.video.get_video_detail_metric_trend?module=129&type=1
func (s *VideoServiceOp[T]) GetVideoDetailMetricTrend(sid uint64, opt GetVideoDetailMetricTrendRequest, tok string) (*GetVideoDetailMetricTrendResponse, error) {
	path := "/video/get_video_detail_metric_trend"
	resp := new(GetVideoDetailMetricTrendResponse)
	err := s.client.Get(path, resp, opt)
	return resp, err
}

// GetVideoDetailPerformance Get detailed performance data for individual post Shopee Video. There is at least a one-day delay.
// Path: /api/v2/video/get_video_detail_performance
// https://open.shopee.com/documents/v2/v2.video.get_video_detail_performance?module=129&type=1
func (s *VideoServiceOp[T]) GetVideoDetailPerformance(sid uint64, opt GetVideoDetailPerformanceRequest, tok string) (*GetVideoDetailPerformanceResponse, error) {
	path := "/video/get_video_detail_performance"
	resp := new(GetVideoDetailPerformanceResponse)
	err := s.client.Get(path, resp, opt)
	return resp, err
}

// GetVideoDetailProductPerformance Get performance data for products linked with individual post Shopee Video. There is at least a one-day delay.
// Path: /api/v2/video/get_video_detail_product_performance
// https://open.shopee.com/documents/v2/v2.video.get_video_detail_product_performance?module=129&type=1
func (s *VideoServiceOp[T]) GetVideoDetailProductPerformance(sid uint64, opt GetVideoDetailProductPerformanceRequest, tok string) (*GetVideoDetailProductPerformanceResponse, error) {
	path := "/video/get_video_detail_product_performance"
	resp := new(GetVideoDetailProductPerformanceResponse)
	err := s.client.Get(path, resp, opt)
	return resp, err
}

// GetVideoList Get the list of video in draft status or video already post to Shopee Video.
// Path: /api/v2/video/get_video_list
// https://open.shopee.com/documents/v2/v2.video.get_video_list?module=129&type=1
func (s *VideoServiceOp[T]) GetVideoList(sid uint64, opt GetVideoListRequest, tok string) (*GetVideoListResponse, error) {
	path := "/video/get_video_list"
	resp := new(GetVideoListResponse)
	err := s.client.Get(path, resp, opt)
	return resp, err
}

// GetVideoPerformanceList Get specific performance data for individual post Shopee Video. There is at least a one-day delay.
// Path: /api/v2/video/get_video_performance_list
// https://open.shopee.com/documents/v2/v2.video.get_video_performance_list?module=129&type=1
func (s *VideoServiceOp[T]) GetVideoPerformanceList(sid uint64, opt GetVideoPerformanceListRequest, tok string) (*GetVideoPerformanceListResponse, error) {
	path := "/video/get_video_performance_list"
	resp := new(GetVideoPerformanceListResponse)
	err := s.client.Get(path, resp, opt)
	return resp, err
}

// PostVideo You need to call v2.media.init_video_upload, v2.media.upload_video_part, and v2.media.complete_video_upload to upload the video, then call the v2.video.edit_video_info API to set video post information, finally call this API to post the video to Shopee Video.
// Path: /api/v2/video/post_video
// https://open.shopee.com/documents/v2/v2.video.post_video?module=129&type=1
func (s *VideoServiceOp[T]) PostVideo(sid uint64, req PostVideoRequest, tok string) (*PostVideoResponse, error) {
	path := "/video/post_video"
	resp := new(PostVideoResponse)
	err := s.client.Post(path, req, resp)
	return resp, err
}
