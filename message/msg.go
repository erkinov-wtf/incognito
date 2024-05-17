package message

const (
	MaxFileSize      = 10 * 1024 * 1024 // 10MB
	MessageToUser    = "Your message has been delivered to the channel."
	PhotoToUser      = "Your photo has been delivered to the channel."
	DocumentToUser   = "Your document has been delivered to the channel."
	FileSizeExceeds  = "The file size exceeds the 10MB limit."
	AllowedFileTypes = "Only text messages, photos, and documents up to 10MB are allowed."
)
