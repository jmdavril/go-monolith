package data

import (
	"github.com/jmoiron/sqlx"
	"log"

	"github.com/samber/lo"
)

type ProductSalesRepo struct {
	db *sqlx.DB
}

func NewProductSalesRepo(db *sqlx.DB) *ProductSalesRepo {
	return &ProductSalesRepo{
		db: db,
	}
}

func (r *ProductSalesRepo) InsertProductSales(sku string) error {
	_, err := r.db.Query("INSERT INTO product_sales(sku, quantity, total_sales) VALUES($1,0,0)", sku)

	if err != nil {
		log.Printf("Could not run SQL query to insert product sales for product %v: %v", sku, err)
		return err
	}

	return nil
}

func (r *ProductSalesRepo) UpdateAllProductSales(o OrderEntity) error {
	arg := lo.Map(o.Items, func(i LineItemEntity, idx int) map[string]interface{} {
		return map[string]interface{}{"sku": i.Sku, "spent": float64(i.Quantity) * i.UnitPrice}
	})

	_, err := r.db.NamedExec(`UPDATE product_sales set quantity=quantity+1, total_sales=total_sales+:spent WHERE sku=:sku`, arg)

	if err != nil {
		log.Printf("Could not run SQL query to update product sales for order %v: %v", o, err)
		return err
	}
	return nil
}

func (r *ProductSalesRepo) ProductSales() ([]ProductSalesEntity, error) {
	var sales []ProductSalesEntity

	err := r.db.Select(&sales, "SELECT * FROM product_sales")

	return sales, err
}
