package utils

import "happy/models"

func GetResModel(status int, message string,data interface{}) models.ResModel {
	return models.ResModel{
		status,
		message,
		data,
	}
}
