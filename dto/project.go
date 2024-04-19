package dto

type CreateProjectInput struct {
	AgreedDeliveryDate string  `json:"agreedDeliveryDate" binding:"required"`
	AgreedPrice        float64 `json:"agreedPrice" binding:"required"`
	OverallProgress    float64 `json:"overallProgress" binding:"required"`
}

type CreateProjectOutput struct {
	Key     string `json:"key"`
	PassKey string `json:"passKey"`
}

type GetProjectOutput struct {
	ID 									uint    `json:"id"`
	AgreedDeliveryDate    string  `json:"agreedDeliveryDate"`
	LastEstimatedIncrease float64 `json:"lastEstimatedIncrease"`
	LastEstimateDate      string  `json:"lastEstimateDate"`
	AgreedPrice           float64 `json:"agreedPrice"`
	OverallProgress       float64 `json:"overallProgress"`
	Key                   string  `json:"key"`
}
