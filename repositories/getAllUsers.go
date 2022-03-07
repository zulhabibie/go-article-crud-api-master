package repositories

import (
	"go-crud-article/connection"
	"go-crud-article/structs"
)

func GetAllUsers(user *structs.User, pagination *structs.Pagination) (*[]structs.User, error) {
	var users []structs.User
	offset := (pagination.Page - 1) * pagination.Limit
	queryBuilder := connection.DB.Select([]string{"userid", "name", "age"}).Limit(pagination.Limit).Offset(offset).Order(pagination.Sort)
	result := queryBuilder.Model(&structs.User{}).Where(user).Find(&users)
	if result.Error != nil {
		message := result.Error
		return nil, message
	}
	return &users, nil
}
