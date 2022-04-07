package dto

import (
	"gin_demo2/model"
)

type UserDto struct {
	Name         string `json:"name"`
	Telephone    string `json:"telephone"`
	Email        string `json:"email"`
	ImgUrl       string `json:"img_url"`
	Sex          string `json:"sex"`
	Team         string `json:"team"`
	Introduction string `json:"introduction"`
	CreateTime   string `json:"created_at"`
}

func ToUserDto(user model.User) UserDto {
	return UserDto{
		Name:         user.Name,
		Telephone:    user.Telephone,
		Email:        user.Email,
		ImgUrl:       user.Img_url,
		Sex:          user.Sex,
		Team:         user.Team,
		Introduction: user.Introduction,
		CreateTime:   user.Model.CreatedAt.Format("2006-01-02"),
	}
}
