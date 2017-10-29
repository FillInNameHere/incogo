package dto

import "gitlab.informatik.haw-hamburg.de/abq331/incogo/businessService/domain"

type AddressDTO struct {
	Id      int
	Uuid    string
	Street  string
	Number  string
	Zip     string
	City    string
	State   string
	Country string
}

func (a AddressDTO) ConvertToDomain() *domain.Address {
	return domain.NewAddress("")
}