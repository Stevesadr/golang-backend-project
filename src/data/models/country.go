package models

type Country struct{
	BaseModels
	Name string `gorm:"size:15;type:string;not null;"`
	City *[]City
}