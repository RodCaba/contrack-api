package models

type Project struct {
	ID                    uint    `json:"id" gorm:"primary_key"`
	AgreedDeliveryDate    string  `json:"agreedDeliveryDate"`
	LastEstimatedIncrease float64 `json:"lastEstimatedIncrease"`
	LastEstimateDate      string  `json:"lastEstimateDate"`
	AgreedPrice           float64 `json:"agreedPrice"`
	OverallProgress       float64 `json:"overallProgress"`
	Key                   string  `json:"key" gorm:"unique, index, not null"`
	PassKey               string  `json:"passKey" gorm:"not null"`
}
