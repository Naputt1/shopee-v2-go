package goshopee

type MediaSpaceCancelVideoUploadRequest struct {
	VideoUploadId string `json:"video_upload_id"` // [Required] The ID of this upload session, returned in init_video_upload.
}
type MediaSpaceCancelVideoUploadResponse struct {
	BaseResponse // Common response fields
}
type MediaSpaceCompleteVideoUploadRequest struct {
	PartSeqList   []int64     `json:"part_seq_list"`   // [Required] All uploaded sequence number.
	ReportData    *ReportData `json:"report_data"`     // [Required]
	VideoUploadId string      `json:"video_upload_id"` // [Required] The ID of this upload session, returned in init_video_upload.
}
type MediaSpaceCompleteVideoUploadResponse struct {
	BaseResponse // Common response fields
}
type MediaSpaceGetVideoUploadResultRequest struct {
	VideoUploadId string `json:"video_upload_id" url:"video_upload_id"` // [Required]
}
type MediaSpaceGetVideoUploadResultResponse struct {
	BaseResponse                                            // Common response fields
	Response     MediaSpaceGetVideoUploadResultResponseData `json:"response"` // Response data
}
type MediaSpaceGetVideoUploadResultResponseData struct {
	Status    string                 `json:"status"`     // [Required] Current status of this video upload session. could be: INITIATED(waiting for part uploading and/or the complete_video_upload API call), TRANSCODING(has received all video parts, and is transcoding the video file), SUCCEEDED(transcoding completed, and this upload_id can now be used for item adding/updating), FAILED(this upload failed, see the message filed for some info), CANCELLED(this upload is cancelled)
	VideoInfo *ResponseDataVideoInfo `json:"video_info"` // [Required] Transcoded video info, will be present if status is SUCCEEDED.
	Message   string                 `json:"message"`    // [Required] Detail error message if video uploading/transcoding failed.
}
type MediaSpaceInitVideoUploadRequest struct {
	FileMd5  string `json:"file_md5"`  // [Required] md5 of video file
	FileSize int64  `json:"file_size"` // [Required] size of video file, in bytes, maximum is 30MB
}
type MediaSpaceInitVideoUploadResponse struct {
	BaseResponse                                       // Common response fields
	Response     MediaSpaceInitVideoUploadResponseData `json:"response"` // Response data
}
type MediaSpaceInitVideoUploadResponseData struct {
	VideoUploadId string `json:"video_upload_id"` // [Required] The identifier of this upload session, used in following video upload request and item creating and/or updating
}
type MediaSpaceUploadImageRequest struct {
	Image interface{} `json:"image"`           // [Required] <p>image files. Max 10.0 MB each. Image format accepted: JPG, JPEG, PNG. image number should be less than 9<br /></p>
	Ratio *string     `json:"ratio,omitempty"` // [Optional] <p>only applicable to whitelisted sellers.<br />only support 1:1 and 3:4;&nbsp;</p><p>1:1 by default.<br /></p>
	Scene *string     `json:"scene,omitempty"` // [Optional] The scene where the picture is used, The value range is normal or desc; normal: we will process the image as a square image, it is recommended to use when uploading item image; desc: we will not process the image, it is recommended to use when uploading the image of extend_description, if you do not upload this field, it will be normal.
}
type MediaSpaceUploadImageResponse struct {
	BaseResponse                                   // Common response fields
	Response     MediaSpaceUploadImageResponseData `json:"response"` // Response data
}
type MediaSpaceUploadImageResponseData struct {
	ImageInfo     *ResponseDataImageInfo             `json:"image_info"`      // [Required]
	ImageInfoList []UploadImageResponseDataImageInfo `json:"image_info_list"` // [Required]
}
type MediaSpaceUploadVideoPartRequest struct {
	ContentMd5    string      `json:"content_md5"`     // [Required] md5 of this part
	PartContent   interface{} `json:"part_content"`    // [Required] The content of this part of file.  Part size should be exactly 4MB, except last part of file.
	PartSeq       int64       `json:"part_seq"`        // [Required] Sequence of the current part, starts from 0
	VideoUploadId string      `json:"video_upload_id"` // [Required] The video_upload_id in the response of initiate_video_upload
}
type MediaSpaceUploadVideoPartResponse struct {
	BaseResponse // Common response fields
}
type ReportData struct {
	UploadCost int64 `json:"upload_cost"` // [Required] Time used for uploading the video file via upload_video_part api, in milliseconds. For video upload performance tracking purpose.
}
type ResponseDataImageInfo struct {
	ImageId      string         `json:"image_id"`       // [Required] <p>Id of image</p>
	ImageUrlList []ThumbnailUrl `json:"image_url_list"` // [Required] <p>Image URL of each region<br /></p>
}
type ResponseDataVideoInfo struct {
	VideoUrlList     []VideoUrl     `json:"video_url_list"`     // [Required] Video playback URL list.
	ThumbnailUrlList []ThumbnailUrl `json:"thumbnail_url_list"` // [Required] Video thumbnail image URL list.
	Duration         int64          `json:"duration"`           // [Required] Duration of this video, in seconds.
}
type ThumbnailUrl struct {
	ImageUrlRegion string `json:"image_url_region"` // [Required] <p>Region of image url<br /></p>
	ImageUrl       string `json:"image_url"`        // [Required] <p>image url<br /></p>
}
type UploadImageResponseDataImageInfo struct {
	Id        int64                  `json:"id"`         // [Required] <p>the index of images</p>
	Error     string                 `json:"error"`      // [Required] <p>Indicate error type if this index's image upload processing hit error. Empty if no error happened for this index's image .<br /></p>
	Message   string                 `json:"message"`    // [Required] <p>Indicate error detail if this index's image upload processing hit error. Empty if no error happened for this index's image .<br /></p>
	ImageInfo *ResponseDataImageInfo `json:"image_info"` // [Required]
}
type VideoUrl struct {
	VideoUrlRegion string `json:"video_url_region"` // [Required] The region of this video URL.
	VideoUrl       string `json:"video_url"`        // [Required] Video playback URL.
}
