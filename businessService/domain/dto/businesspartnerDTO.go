package dto

import "gitlab.informatik.haw-hamburg.de/abq331/incogo/businessService/domain"

type BusinesspartnerDTO struct {
	Id int
	Uuid string
	Name string
	Private bool
	Contactpersons []ContactpersonDTO
	Addresses []AddressDTO
}

func (b BusinesspartnerDTO) ConvertToDomain() *domain.Businesspartner {
	bp := domain.NewBusinesspartner("",true)
	return bp
}