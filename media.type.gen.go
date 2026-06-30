package goshopee

type CancelVideoUploadRequest struct {
	VideoUploadId string `json:"video_upload_id"` // [Required] <p>The unique ID of the upload task, returned by v2.media.init_video_upload.</p>
}
type CancelVideoUploadResponse struct {
	BaseResponse // Common response fields
}
type CompleteVideoUploadRequest struct {
	VideoUploadId string `json:"video_upload_id"` // [Required] <p>The unique ID of the upload task, returned by v2.media.init_video_upload.</p>
}
type CompleteVideoUploadResponse struct {
	BaseResponse // Common response fields
}
type GetVideoUploadResultRequest struct {
	VideoUploadId string `json:"video_upload_id" url:"video_upload_id"` // [Required] <p>The unique ID of the upload task, returned by v2.media.init_video_upload.</p>
}
type GetVideoUploadResultResponse struct {
	BaseResponse                                  // Common response fields
	Response     GetVideoUploadResultResponseData `json:"response"` //
}
type GetVideoUploadResultResponseData struct {
	Status     string     `json:"status"`      // [Required] <p>Current status of the upload task. Possible values:<br /></p><p>- INITIATED: Upload task has been created (via init_video_upload) but no parts have been uploaded yet.</p><p>- UPLOADING: Video file parts are being uploaded. The upload has started but is not yet completed.</p><p>- UPLOADED: All video parts have been uploaded successfully, waiting for complete_video_upload to trigger processing.</p><p>- PROCESSING: Video is being transcoded/validated by the system (duration, format, resolution checks).</p><p>- SUCCEEDED: Video upload and transcoding completed successfully. Video URL and cover URL are available for use.</p><p>- FAILED: Upload or processing failed (e.g., invalid format, duration not within allowed range, transcoding error).</p><p>- CANCELLED: Upload task was explicitly canceled by the client (cancel_video_upload), and the video is discarded.</p>
	Reason     string     `json:"reason"`      // [Required] <p>Detailed fail or cancel reason, will be returned if status is FAILED or CANCELLED.</p>
	UpdateTime int64      `json:"update_time"` // [Required] <p>The time of video status updates.</p>
	VideoInfo  *VideoInfo `json:"video_info"`  // [Required] <p>Transcoded video info, will be returned if status is SUCCEEDED.</p>
}
type InitVideoUploadRequest struct {
	Business int64  `json:"business"`  // [Required] <p>Defines the business type of the uploaded image. Supported values:&nbsp;</p><p>3 = Video</p>
	Duration int64  `json:"duration"`  // [Required] <p>Video duration in seconds.&nbsp;Rules and restrictions by business and scene:</p><p>- If business = 3 (Video) and scene = 1 (Shopee Video): <b>1s~180s</b>.</p>
	FileName string `json:"file_name"` // [Required] <p>Original video file name.</p>
	FileSize int64  `json:"file_size"` // [Required] <p>Total video file size in bytes.&nbsp;Rules and restrictions by business and scene:&nbsp;</p><p>- If business = 3 (Video) and scene = 1 (Shopee Video):&nbsp;Maximum <b>1GB</b>.</p>
	Scene    int64  `json:"scene"`     // [Required] <p>Defines the purpose of the uploaded image under the specified business type.&nbsp;Supported values:&nbsp;</p><p>- If business = 3 (Video): 1 = Shopee Video</p>
}
type InitVideoUploadResponse struct {
	BaseResponse                             // Common response fields
	Response     InitVideoUploadResponseData `json:"response"` //
}
type InitVideoUploadResponseData struct {
	VideoUploadId string `json:"video_upload_id"` // [Required] <p>Unique upload session ID.</p>
	PartSize      int64  `json:"part_size"`       // [Required] <p>The size of each part. When uploading video chunks, the video must be split according to this part size for each upload request.</p>
}
type MediaUploadImageRequest struct {
	Business int64       `json:"business"` // [Required] <p>Defines the business type of the uploaded image. Supported values:&nbsp;</p><p>2 = Returns</p>
	Images   interface{} `json:"images"`   // [Required] <p>The image files to be uploaded. Rules and restrictions by business and scene:</p><p>- If business = 2 (Returns) and scene = 1 (Return Seller Self Arrange Pickup Proof Image): Up to&nbsp;<b>3</b>&nbsp;images can be uploaded. Each image must not exceed&nbsp;<b>10MB</b>. Supported formats:&nbsp;<b>JPG, JPEG, PNG</b>.</p>
	Scene    int64       `json:"scene"`    // [Required] <p>Defines the purpose of the uploaded image under the specified business type.&nbsp;Supported values:</p><p>- If&nbsp;business = 2 (Returns): 1 = Return Seller Self Arrange Pickup Proof Image</p>
}
type MediaUploadImageResponse struct {
	BaseResponse                              // Common response fields
	Response     MediaUploadImageResponseData `json:"response"` //
}
type MediaUploadImageResponseData struct {
	ImageList []FieldImageInfo `json:"image_list"` // [Required] <p>List of uploaded images.</p>
}
type UploadVideoPartRequest struct {
	PartContent   interface{} `json:"part_content"`    // [Required] <p>The content of this part of file. Part size should be exactly equal to part_size returned in v2.media.init_video_upload, except last part of file.</p>
	PartMd5       string      `json:"part_md5"`        // [Required] <p>MD5 checksum of this part for data integrity validation.</p>
	PartSeq       int64       `json:"part_seq"`        // [Required] <p>Sequence number of this part, starting from 0.</p>
	VideoUploadId string      `json:"video_upload_id"` // [Required] <p>The unique ID of the upload task, returned by v2.media.init_video_upload.</p>
}
type UploadVideoPartResponse struct {
	BaseResponse // Common response fields
}
type VideoInfo struct {
	VideoUrl          string `json:"video_url"`           // [Required] <p>Video playback url.</p>
	VideoThumbnailUrl string `json:"video_thumbnail_url"` // [Required] <p>Video thumbnail image url.</p>
	ThumbnailWidth    int64  `json:"thumbnail_width"`     // [Required] <p>Video thumbnail image width.</p>
	ThumbnailHeight   int64  `json:"thumbnail_height"`    // [Required] <p>Video thumbnail image width.</p>
	Duration          int64  `json:"duration"`            // [Required] <p>Video duration in seconds.</p>
	Resolution        string `json:"resolution"`          // [Required] <p>Video resolution, e.g., "1280x1280".</p>
}
