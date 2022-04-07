package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name         string `gorm:"type:varchar(20);not null"`
	Telephone    string `gorm:"type:varchar(110);not null;unique"`
	Password     string `gorm:"size:255;not null"`
	Email        string `gorm:"type:varchar(110);not null;unique"`
	Img_url      string `gorm:"type:varchar(255);default:'https://folio-website-images.s3.eu-west-2.amazonaws.com/content/uploads/2021/05/26164557/Peter-Greenwood-Folio-Illustration-Seaplane-Affinity.jpg'"`
	Sex          string `gorm:"type:bool;default:true"`
	Team         string `gorm:"type:varchar(110);"`
	Introduction string `gorm:"size:2000;"`
}
