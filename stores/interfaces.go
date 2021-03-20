package stores

import "goApi/models"

type Product interface {
	Create(product models.Product) error
	Read() ([]models.Product, error)
	Update(price string, id int) error
	Delete(id int) error
}
