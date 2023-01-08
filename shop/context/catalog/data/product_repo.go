package data

import (
	"github.com/jmoiron/sqlx"
	"log"

	"github.com/google/uuid"
)

type ProductRepo struct {
	db *sqlx.DB
}

func NewProductRepo(db *sqlx.DB) *ProductRepo {
	return &ProductRepo{
		db: db,
	}
}

func (r *ProductRepo) SelectProduct(productId uuid.UUID) (ProductEntity, error) {
	row := r.db.QueryRowx("SELECT * FROM product WHERE id=$1", productId)
	var product ProductEntity
	err := row.StructScan(&product)

	if err != nil {
		log.Printf("Could not run SQL query to read product: %v", err)
		return product, err
	}
	return product, nil
}

func (r *ProductRepo) InsertProduct(p ProductEntity) (uuid.UUID, error) {
	rows, err := r.db.Query("INSERT INTO product(sku, name, price) VALUES($1,$2,$3) RETURNING id", p.Sku, p.Name, p.Price)

	if err != nil {
		log.Printf("Could not run SQL query to insert product %v: %v", p, err)
		return uuid.Nil, err
	}
	var productId uuid.UUID
	if rows.Next() {
		err := rows.Scan(&productId)
		if err != nil {
			return [16]byte{}, err
		}
	}
	return productId, nil
}

func (r *ProductRepo) UpdateProduct(p ProductEntity) error {
	log.Printf("This is p %v", p)
	_, err := r.db.Exec("UPDATE product SET sku=$1, name=$2, price=$3 WHERE id=$4", p.Sku, p.Name, p.Price, p.ID)

	if err != nil {
		log.Printf("Could not run SQL query to update product to %v: %v", p, err)
		return err
	}
	return nil
}
