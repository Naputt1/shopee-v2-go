package goshopee

import "io"

type MediaSpaceService interface {
	// CancelVideoUpload Cancel a video upload session
	// Path: /api/v2/media_space/cancel_video_upload
	// https://open.shopee.com/documents/v2/v2.media_space.cancel_video_upload?module=91&type=1
	CancelVideoUpload(sid uint64, filename string, tok string) (*MediaSpaceCancelVideoUploadResponse, error)
	CancelVideoUploadFromReader(sid uint64, filename string, reader io.Reader, tok string) (*MediaSpaceCancelVideoUploadResponse, error)
	// CompleteVideoUpload Complete the video upload and starts the transcoding process when all parts are uploaded successfully.
	// Path: /api/v2/media_space/complete_video_upload
	// https://open.shopee.com/documents/v2/v2.media_space.complete_video_upload?module=91&type=1
	CompleteVideoUpload(sid uint64, filename string, tok string) (*MediaSpaceCompleteVideoUploadResponse, error)
	CompleteVideoUploadFromReader(sid uint64, filename string, reader io.Reader, tok string) (*MediaSpaceCompleteVideoUploadResponse, error)
	// GetVideoUploadResult Query the upload status and result of video upload.
	// Path: /api/v2/media_space/get_video_upload_result
	// https://open.shopee.com/documents/v2/v2.media_space.get_video_upload_result?module=91&type=1
	GetVideoUploadResult(sid uint64, filename string, tok string) (*MediaSpaceGetVideoUploadResultResponse, error)
	GetVideoUploadResultFromReader(sid uint64, filename string, reader io.Reader, tok string) (*MediaSpaceGetVideoUploadResultResponse, error)
	// InitVideoUpload Initiate video upload session.
	// 
	// Video duration should be between 10s and 60s (inclusive).
	// Path: /api/v2/media_space/init_video_upload
	// https://open.shopee.com/documents/v2/v2.media_space.init_video_upload?module=91&type=1
	InitVideoUpload(sid uint64, filename string, tok string) (*MediaSpaceInitVideoUploadResponse, error)
	InitVideoUploadFromReader(sid uint64, filename string, reader io.Reader, tok string) (*MediaSpaceInitVideoUploadResponse, error)
	// UploadImage Use this API to upload multiple image files (less than 9 images). 
	// Path: /api/v2/media_space/upload_image
	// https://open.shopee.com/documents/v2/v2.media_space.upload_image?module=91&type=1
	UploadImage(sid uint64, filename string, tok string) (*MediaSpaceUploadImageResponse, error)
	UploadImageFromReader(sid uint64, filename string, reader io.Reader, tok string) (*MediaSpaceUploadImageResponse, error)
	// UploadVideoPart Upload video file by part using the upload_id in initiate_video_upload.
	// 
	// The request Content-Type of this API should be of multipart/form-data
	// 
	// 
	// Path: /api/v2/media_space/upload_video_part
	// https://open.shopee.com/documents/v2/v2.media_space.upload_video_part?module=91&type=1
	UploadVideoPart(sid uint64, filename string, tok string) (*MediaSpaceUploadVideoPartResponse, error)
	UploadVideoPartFromReader(sid uint64, filename string, reader io.Reader, tok string) (*MediaSpaceUploadVideoPartResponse, error)
}

type MediaSpaceServiceOp[T any] struct {
	client *Client[T]
}

// CancelVideoUpload Cancel a video upload session
// Path: /api/v2/media_space/cancel_video_upload
// https://open.shopee.com/documents/v2/v2.media_space.cancel_video_upload?module=91&type=1
func (s *MediaSpaceServiceOp[T]) CancelVideoUpload(sid uint64, filename string, tok string) (*MediaSpaceCancelVideoUploadResponse, error) {
	path := "/media_space/cancel_video_upload"
	resp := new(MediaSpaceCancelVideoUploadResponse)
	err := s.client.WithShop(sid, tok).Upload(path, "image", filename, resp)
	return resp, err
}

func (s *MediaSpaceServiceOp[T]) CancelVideoUploadFromReader(sid uint64, filename string, reader io.Reader, tok string) (*MediaSpaceCancelVideoUploadResponse, error) {
	path := "/media_space/cancel_video_upload"
	resp := new(MediaSpaceCancelVideoUploadResponse)
	err := s.client.WithShop(sid, tok).UploadFromReader(path, "image", filename, reader, resp)
	return resp, err
}

// CompleteVideoUpload Complete the video upload and starts the transcoding process when all parts are uploaded successfully.
// Path: /api/v2/media_space/complete_video_upload
// https://open.shopee.com/documents/v2/v2.media_space.complete_video_upload?module=91&type=1
func (s *MediaSpaceServiceOp[T]) CompleteVideoUpload(sid uint64, filename string, tok string) (*MediaSpaceCompleteVideoUploadResponse, error) {
	path := "/media_space/complete_video_upload"
	resp := new(MediaSpaceCompleteVideoUploadResponse)
	err := s.client.WithMerchant(sid, tok).Upload(path, "image", filename, resp)
	return resp, err
}

func (s *MediaSpaceServiceOp[T]) CompleteVideoUploadFromReader(sid uint64, filename string, reader io.Reader, tok string) (*MediaSpaceCompleteVideoUploadResponse, error) {
	path := "/media_space/complete_video_upload"
	resp := new(MediaSpaceCompleteVideoUploadResponse)
	err := s.client.WithMerchant(sid, tok).UploadFromReader(path, "image", filename, reader, resp)
	return resp, err
}

// GetVideoUploadResult Query the upload status and result of video upload.
// Path: /api/v2/media_space/get_video_upload_result
// https://open.shopee.com/documents/v2/v2.media_space.get_video_upload_result?module=91&type=1
func (s *MediaSpaceServiceOp[T]) GetVideoUploadResult(sid uint64, filename string, tok string) (*MediaSpaceGetVideoUploadResultResponse, error) {
	path := "/media_space/get_video_upload_result"
	resp := new(MediaSpaceGetVideoUploadResultResponse)
	err := s.client.WithShop(sid, tok).Upload(path, "image", filename, resp)
	return resp, err
}

func (s *MediaSpaceServiceOp[T]) GetVideoUploadResultFromReader(sid uint64, filename string, reader io.Reader, tok string) (*MediaSpaceGetVideoUploadResultResponse, error) {
	path := "/media_space/get_video_upload_result"
	resp := new(MediaSpaceGetVideoUploadResultResponse)
	err := s.client.WithShop(sid, tok).UploadFromReader(path, "image", filename, reader, resp)
	return resp, err
}

// InitVideoUpload Initiate video upload session.
// 
// Video duration should be between 10s and 60s (inclusive).
// Path: /api/v2/media_space/init_video_upload
// https://open.shopee.com/documents/v2/v2.media_space.init_video_upload?module=91&type=1
func (s *MediaSpaceServiceOp[T]) InitVideoUpload(sid uint64, filename string, tok string) (*MediaSpaceInitVideoUploadResponse, error) {
	path := "/media_space/init_video_upload"
	resp := new(MediaSpaceInitVideoUploadResponse)
	err := s.client.WithShop(sid, tok).Upload(path, "image", filename, resp)
	return resp, err
}

func (s *MediaSpaceServiceOp[T]) InitVideoUploadFromReader(sid uint64, filename string, reader io.Reader, tok string) (*MediaSpaceInitVideoUploadResponse, error) {
	path := "/media_space/init_video_upload"
	resp := new(MediaSpaceInitVideoUploadResponse)
	err := s.client.WithShop(sid, tok).UploadFromReader(path, "image", filename, reader, resp)
	return resp, err
}

// UploadImage Use this API to upload multiple image files (less than 9 images). 
// Path: /api/v2/media_space/upload_image
// https://open.shopee.com/documents/v2/v2.media_space.upload_image?module=91&type=1
func (s *MediaSpaceServiceOp[T]) UploadImage(sid uint64, filename string, tok string) (*MediaSpaceUploadImageResponse, error) {
	path := "/media_space/upload_image"
	resp := new(MediaSpaceUploadImageResponse)
	err := s.client.WithMerchant(sid, tok).Upload(path, "image", filename, resp)
	return resp, err
}

func (s *MediaSpaceServiceOp[T]) UploadImageFromReader(sid uint64, filename string, reader io.Reader, tok string) (*MediaSpaceUploadImageResponse, error) {
	path := "/media_space/upload_image"
	resp := new(MediaSpaceUploadImageResponse)
	err := s.client.WithMerchant(sid, tok).UploadFromReader(path, "image", filename, reader, resp)
	return resp, err
}

// UploadVideoPart Upload video file by part using the upload_id in initiate_video_upload.
// 
// The request Content-Type of this API should be of multipart/form-data
// 
// 
// Path: /api/v2/media_space/upload_video_part
// https://open.shopee.com/documents/v2/v2.media_space.upload_video_part?module=91&type=1
func (s *MediaSpaceServiceOp[T]) UploadVideoPart(sid uint64, filename string, tok string) (*MediaSpaceUploadVideoPartResponse, error) {
	path := "/media_space/upload_video_part"
	resp := new(MediaSpaceUploadVideoPartResponse)
	err := s.client.WithMerchant(sid, tok).Upload(path, "image", filename, resp)
	return resp, err
}

func (s *MediaSpaceServiceOp[T]) UploadVideoPartFromReader(sid uint64, filename string, reader io.Reader, tok string) (*MediaSpaceUploadVideoPartResponse, error) {
	path := "/media_space/upload_video_part"
	resp := new(MediaSpaceUploadVideoPartResponse)
	err := s.client.WithMerchant(sid, tok).UploadFromReader(path, "image", filename, reader, resp)
	return resp, err
}

