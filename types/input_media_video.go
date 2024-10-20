package types

type InputMediaVideo struct {
	Type                  string           `json:"type"`                               // Type of the result, must be video
	Media                 string           `json:"media"`                              // File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass “attach://<file_attach_name>” to upload a new one using multipart/form-data under <file_attach_name> name. More information on Sending Files »
	Thumbnail             *interface{}     `json:"thumbnail,omitempty"`                // Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More information on Sending Files »
	Caption               string           `json:"caption,omitempty"`                  // Optional. Caption of the video to be sent, 0-1024 characters after entities parsing
	ParseMode             string           `json:"parse_mode,omitempty"`               // Optional. Mode for parsing entities in the video caption. See formatting options for more details.
	CaptionEntities       []*MessageEntity `json:"caption_entities,omitempty"`         // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	ShowCaptionAboveMedia bool             `json:"show_caption_above_media,omitempty"` // Optional. Pass True, if the caption must be shown above the message media
	Width                 int32            `json:"width,omitempty"`                    // Optional. Video width
	Height                int32            `json:"height,omitempty"`                   // Optional. Video height
	Duration              int32            `json:"duration,omitempty"`                 // Optional. Video duration in seconds
	SupportsStreaming     bool             `json:"supports_streaming,omitempty"`       // Optional. Pass True if the uploaded video is suitable for streaming
	HasSpoiler            bool             `json:"has_spoiler,omitempty"`              // Optional. Pass True if the video needs to be covered with a spoiler animation
}
