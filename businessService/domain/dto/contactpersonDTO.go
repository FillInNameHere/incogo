package dto

import "gitlab.informatik.haw-hamburg.de/abq331/incogo/businessService/util/datatype"

type ContactpersonDTO struct {
	Id int
	Uuid string
	Title string
	Prename string
	Lastname string
	Email datatype.Email
	Telephone string //vt
}
