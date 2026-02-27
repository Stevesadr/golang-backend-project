package models

type City struct{
	BaseModels
	Name string `gorm:"size: 5;type:string;not null"`
	CountryId int
	Country Country `gorm:"foreignKey:CountryId"`
}