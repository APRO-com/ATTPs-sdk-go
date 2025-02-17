package consts

import "time"

// SDK related information
const (
	Version         = "0.0.1"                  // SDK version
	UserAgentFormat = "ATTPs-Go/%s (%s) GO/%s" // UserAgent's info
	APIBaseServer   = "http://10.0.54.95:8888" // base API address
)

// HTTP request message Header related constants
const (
	Accept        = "Accept"         // Header's Accept
	ContentType   = "Content-Type"   // Header's ContentType
	ContentLength = "Content-Length" // Header's ContentLength
	UserAgent     = "User-Agent"     // Header's UserAgent
	Authorization = "Authorization"  // Header's Authorization
)

// Common ContentType
const (
	ApplicationJSON = "application/json"
	ImageJPG        = "image/jpg"
	ImagePNG        = "image/png"
	VideoMP4        = "video/mp4"
)

const (
	FiveMinute     = 5 * 60
	DefaultTimeout = 30 * time.Second
)

const (
	VRFVersion = 1
)
