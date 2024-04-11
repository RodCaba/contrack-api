package dto

type CreateProjectInput struct {
	AgreedDeliveryDate string  `json:"agreedDeliveryDate" binding:"required"`
	AgreedPrice         float64 `json:"agreedPrice" binding:"required"`
	OverallProgress    float64 `json:"overallProgress" binding:"required"`
}