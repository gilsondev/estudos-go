package models

import (
	"database/sql"

	"github.com/google/uuid"
)

type Product struct {
	ID    string
	Name  string
	Price float64
}

func NewProduct(name string, price float64) *Product {
	productID := uuid.New()
	return &Product{ID: productID.String(), Name: name, Price: price}
}

func GetProducts(db *sql.DB) (products []Product, err error) {
	rows, err := db.Query("SELECT id, name, price FROM products")
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var product Product
		err = rows.Scan(&product.ID, &product.Name, &product.Price)
		if err != nil {
			return
		}

		products = append(products, product)
	}

	return products, err
}

func GetProduct(db *sql.DB, id string) (*Product, error) {
	stmt, err := db.Prepare("SELECT id, name, price FROM products WHERE id = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)
	if row.Err() != nil {
		return nil, row.Err()
	}

	var product Product
	row.Scan(&product.ID, &product.Name, &product.Price)

	return &product, nil
}

func (p *Product) Insert(db *sql.DB) error {
	stmt, err := db.Prepare("INSERT INTO products (id, name, price) VALUES ($1, $2, $3)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(p.ID, p.Name, p.Price)
	if err != nil {
		return err
	}

	return nil
}

func (p *Product) Update(db *sql.DB) error {
	stmt, err := db.Prepare("UPDATE products SET name = $1, price = $2 WHERE id = $3")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(p.Name, p.Price, p.ID)
	if err != nil {
		return err
	}

	return nil
}

func (p *Product) Delete(db *sql.DB) error {
	stmt, err := db.Prepare("DELETE FROM products WHERE id = $1")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(p.ID)
	if err != nil {
		return err
	}

	return nil
}
