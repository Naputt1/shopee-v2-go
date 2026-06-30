package goshopee

import (
	"context"
	"io"
)

type MediaSpaceService interface {
	// CancelVideoUpload Cancel a video upload session
	// Path: /api/v2/media_space/cancel_video_upload
	// https://open.shopee.com/documents/v2/v2.media_space.cancel_video_upload?module=91&type=1
	CancelVideoUpload(ctx context.Context, sid uint64, filename string, tok string) (*MediaSpaceCancelVideoUploadResponse, error)
	CancelVideoUploadFromReader(ctx context.Context, sid uint64, filename string, reader io.Reader, tok string) (*MediaSpaceCancelVideoUploadResponse, error)
	// CompleteVideoUpload Complete the video upload and starts the transcoding process when all parts are uploaded successfully.
	// Path: /api/v2/media_space/complete_video_upload
	// https://open.shopee.com/documents/v2/v2.media_space.complete_video_upload?module=91&type=1
	CompleteVideoUpload(ctx context.Context, sid uint64, filename string, tok string) (*MediaSpaceCompleteVideoUploadResponse, error)
	CompleteVideoUploadFromReader(ctx context.Context, sid uint64, filename string, reader io.Reader, tok string) (*MediaSpaceCompleteVideoUploadResponse, error)
	// GetVideoUploadResult Query the upload status and result of video upload.
	// Path: /api/v2/media_space/get_video_upload_result
	// https://open.shopee.com/documents/v2/v2.media_space.get_video_upload_result?module=91&type=1
	GetVideoUploadResult(ctx context.Context, sid uint64, filename string, tok string) (*MediaSpaceGetVideoUploadResultResponse, error)
	GetVideoUploadResultFromReader(ctx context.Context, sid uint64, filename string, reader io.Reader, tok string) (*MediaSpaceGetVideoUploadResultResponse, error)
	// InitVideoUpload Initiate video upload session.
	//
	// Video duration should be between 10s and 60s (inclusive).
	// Path: /api/v2/media_space/init_video_upload
	// https://open.shopee.com/documents/v2/v2.media_space.init_video_upload?module=91&type=1
	InitVideoUpload(ctx context.Context, sid uint64, filename string, tok string) (*MediaSpaceInitVideoUploadResponse, error)
	InitVideoUploadFromReader(ctx context.Context, sid uint64, filename string, reader io.Reader, tok string) (*MediaSpaceInitVideoUploadResponse, error)
	// UploadImage Use this API to upload multiple image files (less than 9 images).
	// Path: /api/v2/media_space/upload_image
	// https://open.shopee.com/documents/v2/v2.media_space.upload_image?module=91&type=1
	UploadImage(ctx context.Context, sid uint64, filename string, tok string) (*MediaSpaceUploadImageResponse, error)
	UploadImageFromReader(ctx context.Context, sid uint64, filename string, reader io.Reader, tok string) (*MediaSpaceUploadImageResponse, error)
	// UploadVideoPart Upload video file by part using the upload_id in initiate_video_upload.
	//
	// The request Content-Type of this API should be of multipart/form-data
	//
	//
	// Path: /api/v2/media_space/upload_video_part
	// https://open.shopee.com/documents/v2/v2.media_space.upload_video_part?module=91&type=1
	UploadVideoPart(ctx context.Context, sid uint64, filename string, tok string) (*MediaSpaceUploadVideoPartResponse, error)
	UploadVideoPartFromReader(ctx context.Context, sid uint64, filename string, reader io.Reader, tok string) (*MediaSpaceUploadVideoPartResponse, error)
}

type MediaSpaceServiceOp[T any] struct {
	client *Client[T]
}

// CancelVideoUpload Cancel a video upload session
// Path: /api/v2/media_space/cancel_video_upload
// https://open.shopee.com/documents/v2/v2.media_space.cancel_video_upload?module=91&type=1
func (s *MediaSpaceServiceOp[T]) CancelVideoUpload(ctx context.Context, sid uint64, filename string, tok string) (*MediaSpaceCancelVideoUploadResponse, error) {
	path := "/media_space/cancel_video_upload"
	resp := new(MediaSpaceCancelVideoUploadResponse)
	err := s.client.WithShop(sid, tok).Upload(ctx, path, "image", filename, resp)
	return resp, err
}

func (s *MediaSpaceServiceOp[T]) CancelVideoUploadFromReader(ctx context.Context, sid uint64, filename string, reader io.Reader, tok string) (*MediaSpaceCancelVideoUploadResponse, error) {
	path := "/media_space/cancel_video_upload"
	resp := new(MediaSpaceCancelVideoUploadResponse)
	err := s.client.WithShop(sid, tok).UploadFromReader(ctx, path, "image", filename, reader, resp)
	return resp, err
}

// CompleteVideoUpload Complete the video upload and starts the transcoding process when all parts are uploaded successfully.
// Path: /api/v2/media_space/complete_video_upload
// https://open.shopee.com/documents/v2/v2.media_space.complete_video_upload?module=91&type=1
func (s *MediaSpaceServiceOp[T]) CompleteVideoUpload(ctx context.Context, sid uint64, filename string, tok string) (*MediaSpaceCompleteVideoUploadResponse, error) {
	path := "/media_space/complete_video_upload"
	resp := new(MediaSpaceCompleteVideoUploadResponse)
	err := s.client.WithMerchant(sid, tok).Upload(ctx, path, "image", filename, resp)
	return resp, err
}

func (s *MediaSpaceServiceOp[T]) CompleteVideoUploadFromReader(ctx context.Context, sid uint64, filename string, reader io.Reader, tok string) (*MediaSpaceCompleteVideoUploadResponse, error) {
	path := "/media_space/complete_video_upload"
	resp := new(MediaSpaceCompleteVideoUploadResponse)
	err := s.client.WithMerchant(sid, tok).UploadFromReader(ctx, path, "image", filename, reader, resp)
	return resp, err
}

// GetVideoUploadResult Query the upload status and result of video upload.
// Path: /api/v2/media_space/get_video_upload_result
// https://open.shopee.com/documents/v2/v2.media_space.get_video_upload_result?module=91&type=1
func (s *MediaSpaceServiceOp[T]) GetVideoUploadResult(ctx context.Context, sid uint64, filename string, tok string) (*MediaSpaceGetVideoUploadResultResponse, error) {
	path := "/media_space/get_video_upload_result"
	resp := new(MediaSpaceGetVideoUploadResultResponse)
	err := s.client.WithShop(sid, tok).Upload(ctx, path, "image", filename, resp)
	return resp, err
}

func (s *MediaSpaceServiceOp[T]) GetVideoUploadResultFromReader(ctx context.Context, sid uint64, filename string, reader io.Reader, tok string) (*MediaSpaceGetVideoUploadResultResponse, error) {
	path := "/media_space/get_video_upload_result"
	resp := new(MediaSpaceGetVideoUploadResultResponse)
	err := s.client.WithShop(sid, tok).UploadFromReader(ctx, path, "image", filename, reader, resp)
	return resp, err
}

// InitVideoUpload Initiate video upload session.
//
// Video duration should be between 10s and 60s (inclusive).
// Path: /api/v2/media_space/init_video_upload
// https://open.shopee.com/documents/v2/v2.media_space.init_video_upload?module=91&type=1
func (s *MediaSpaceServiceOp[T]) InitVideoUpload(ctx context.Context, sid uint64, filename string, tok string) (*MediaSpaceInitVideoUploadResponse, error) {
	path := "/media_space/init_video_upload"
	resp := new(MediaSpaceInitVideoUploadResponse)
	err := s.client.WithShop(sid, tok).Upload(ctx, path, "image", filename, resp)
	return resp, err
}

func (s *MediaSpaceServiceOp[T]) InitVideoUploadFromReader(ctx context.Context, sid uint64, filename string, reader io.Reader, tok string) (*MediaSpaceInitVideoUploadResponse, error) {
	path := "/media_space/init_video_upload"
	resp := new(MediaSpaceInitVideoUploadResponse)
	err := s.client.WithShop(sid, tok).UploadFromReader(ctx, path, "image", filename, reader, resp)
	return resp, err
}

// UploadImage Use this API to upload multiple image files (less than 9 images).
// Path: /api/v2/media_space/upload_image
// https://open.shopee.com/documents/v2/v2.media_space.upload_image?module=91&type=1
func (s *MediaSpaceServiceOp[T]) UploadImage(ctx context.Context, sid uint64, filename string, tok string) (*MediaSpaceUploadImageResponse, error) {
	path := "/media_space/upload_image"
	resp := new(MediaSpaceUploadImageResponse)
	err := s.client.WithMerchant(sid, tok).Upload(ctx, path, "image", filename, resp)
	return resp, err
}

func (s *MediaSpaceServiceOp[T]) UploadImageFromReader(ctx context.Context, sid uint64, filename string, reader io.Reader, tok string) (*MediaSpaceUploadImageResponse, error) {
	path := "/media_space/upload_image"
	resp := new(MediaSpaceUploadImageResponse)
	err := s.client.WithMerchant(sid, tok).UploadFromReader(ctx, path, "image", filename, reader, resp)
	return resp, err
}

// UploadVideoPart Upload video file by part using the upload_id in initiate_video_upload.
//
// The request Content-Type of this API should be of multipart/form-data
//
// Path: /api/v2/media_space/upload_video_part
// https://open.shopee.com/documents/v2/v2.media_space.upload_video_part?module=91&type=1
func (s *MediaSpaceServiceOp[T]) UploadVideoPart(ctx context.Context, sid uint64, filename string, tok string) (*MediaSpaceUploadVideoPartResponse, error) {
	path := "/media_space/upload_video_part"
	resp := new(MediaSpaceUploadVideoPartResponse)
	err := s.client.WithMerchant(sid, tok).Upload(ctx, path, "image", filename, resp)
	return resp, err
}

func (s *MediaSpaceServiceOp[T]) UploadVideoPartFromReader(ctx context.Context, sid uint64, filename string, reader io.Reader, tok string) (*MediaSpaceUploadVideoPartResponse, error) {
	path := "/media_space/upload_video_part"
	resp := new(MediaSpaceUploadVideoPartResponse)
	err := s.client.WithMerchant(sid, tok).UploadFromReader(ctx, path, "image", filename, reader, resp)
	return resp, err
}
