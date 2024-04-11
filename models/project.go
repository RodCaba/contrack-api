package models

type Project struct {
	ID uint `json:"id" gorm:"primary_key"`
	AgreedDeliveryDate string `json:"agreedDeliveryDate"`
	LastEstimatedIncrease float64 `json:"lastEstimatedIncrease"`
	LastEstimateDate string `json:"lastEstimateDate"`
	AgreedPrice float64 `json:"agreedPrice"`
	OverallProgress float64 `json:"overallProgress"`
}