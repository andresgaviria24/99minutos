package dto

type AddressDto struct {
	Address  string `json:"address"`
	ZipCode  string `json:"zipcode"`
	Region   string `json:"region"`
	RegionId int    `json:"region_id"`
}
