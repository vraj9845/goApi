package services

import (
	"goApi/models"
	"goApi/stores"
)

type productService struct {
	store stores.Product
}

func (p *productService) Create(product models.Product) error {
	err := p.store.Create(product)

	if err != nil {
		return err
	}

	return nil
}

func (p *productService) Read() ([]models.Product, error) {
	products, err := p.store.Read()
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (p *productService) Update(price string, id int) error {
	err := p.store.Update(price, id)
	if err != nil {
		return err
	}

	return nil
}

func (p *productService) Delete(id int) error {
	err := p.store.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

func New(store stores.Product) productService {
	return productService{store: store}
}
