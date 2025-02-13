package service

// BaseResponse Base Response Details
type BaseResponse struct {
	Message      *string `json:"message"`
	Code         int64   `json:"code"`
	ResponseEnum *string `json:"responseEnum"`
}
