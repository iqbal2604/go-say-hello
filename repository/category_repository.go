package repository

import "github.com/iqbal2604/go-say-hello/v2/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
