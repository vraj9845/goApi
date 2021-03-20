package stores

import (
	"context"
	"database/sql"
	"fmt"
	"goApi/models"
)

type productStore struct {
	db *sql.DB
}

func (p *productStore) Create(product models.Product) error {
	insertQuery := "INSERT INTO product (id,name,price) VALUES (?,?,?)"
	_, err := p.db.ExecContext(context.TODO(), insertQuery, product.Id, product.Name, product.Price)
	if err != nil {
		return err
	}

	fmt.Println("1 row succesfully created!")
	return nil
}

func (p *productStore) Read() ([]models.Product, error) {
	var products []models.Product
	selectQuery := "SELECT id,name,price FROM product"
	rows, err := p.db.QueryContext(context.TODO(), selectQuery)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var product models.Product
		err = rows.Scan(&product.Id, &product.Name, &product.Price)
		if err != nil {
			return nil, err
		}

		products = append(products, product)
		fmt.Printf("%v | %v | %v\n", product.Id, product.Name, product.Price)
	}

	return products, nil
}

func (p *productStore) Update(price string, id int) error {
	updateQuery := "UPDATE product SET price=? WHERE id=?"
	_, err := p.db.ExecContext(context.TODO(), updateQuery, price, id)
	if err != nil {
		return err
	}

	fmt.Println("1 row succesfully updated!")
	return nil
}

func (p *productStore) Delete(id int) error {
	deleteQuery := "DELETE FROM product WHERE id=?"
	_, err := p.db.ExecContext(context.TODO(), deleteQuery, id)
	if err != nil {
		return err
	}
	fmt.Println("1 row succesfully deleted!")

	return nil
}

func New(db *sql.DB) *productStore {
	return &productStore{db: db}
}
