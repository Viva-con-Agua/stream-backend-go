package models

type (
	Address struct {
		Uuid       string `json:"uuid" validate:"required"`
		Street     string `json:"street" validate:"required"`
		Additional string `json:"additional" validate;"required"`
		Zip        string `json:"zip" validate;"required"`
		City       string `json:"city" validate;"required"`
		Country    string `json:"country" validate;"required"`
		GoogleId   string `json:"google_id" validate;"required"`
		Updated    int    `json:"updated" validate:"required"`
		Created    int    `json:"created" validate:"required"`
	}
	AddressCreate struct {
		Street     string `json:"street" validate:"required"`
		Additional string `json:"additional" validate;"required"`
		Zip        string `json:"zip" validate;"required"`
		City       string `json:"city" validate;"required"`
		Country    string `json:"country" validate;"required"`
		GoogleId   string `json:"google_id" validate;"required"`
		Updated    int    `json:"updated" validate:"required"`
		Created    int    `json:"created" validate:"required"`
	}
)
