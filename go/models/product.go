package models

import (
	"database/sql"
	"fmt"
)

type Product struct {
	ProductId   []byte `xml:"product_id"`
	ProductName []byte `xml:"product_name"`
}

/*
implements the Loadable interface
*/
func (d *Product) LoadValue(prefix string, name string, value []string) {
	fmt.Println("name:", name)
	switch name {
	case prefix + "productId":
		d.ProductId = []byte(value[0])
	case prefix + "productName":
		d.ProductName = []byte(value[0])
	}
}

func GetListProduct(queryMap map[string]string, offset int, limit int) ([]*Product, int) {
	var total int
	products := []*Product{}

	dbCountFunc := func(queryMap *map[string]string) error {
		var err error
		if (*queryMap)["query"] == "" {
			err = DbMap.QueryRow("select count(product_id) from PRODUCT").Scan(&total)
		} else {
			err = DbMap.QueryRow("select count(product_id) from PRODUCT where product_id like ?", (*queryMap)["query"]).Scan(&total)
		}
		return err
	}

	dbRowsFunc := func(queryMap *map[string]string) (*sql.Rows, error) {
		var err error
		var rows *sql.Rows
		if (*queryMap)["query"] == "" {
			rows, err = DbMap.Query("select product_id, product_name from PRODUCT limit ?, ?", offset, limit)
		} else {
			rows, err = DbMap.Query("select product_id, product_name from PRODUCT where product_id like ? limit ?, ?", (*queryMap)["query"], offset, limit)
		}
		return rows, err
	}

	dbListFunc := func(rows *sql.Rows) error {
		var err error
		var p Product
		err = rows.Scan(&p.ProductId, &p.ProductName)
		if err == nil {
			products = append(products, &p)
		}
		return err
	}

	GetDbListTemplate(queryMap, dbCountFunc, dbRowsFunc, dbListFunc)

	return products, total
}

func GetTableProduct(queryMap map[string]string, offset int, limit int) ([]*Product, int) {
	var total int
	products := []*Product{}

	dbCountFunc := func(queryMap *map[string]string) error {
		var err error
		if (*queryMap)["query"] == "" {
			err = DbMap.QueryRow("select count(product_id) from PRODUCT").Scan(&total)
		} else {
			err = DbMap.QueryRow("select count(product_id) from PRODUCT where product_id like ?", (*queryMap)["query"]).Scan(&total)
		}
		return err
	}

	dbRowsFunc := func(queryMap *map[string]string) (*sql.Rows, error) {
		var err error
		var rows *sql.Rows
		if (*queryMap)["query"] == "" {
			rows, err = DbMap.Query("select product_id, product_name from PRODUCT limit ?, ?", offset, limit)
		} else {
			rows, err = DbMap.Query("select product_id, product_name from PRODUCT where product_id like ? limit ?, ?", (*queryMap)["query"], offset, limit)
		}
		return rows, err
	}

	dbListFunc := func(rows *sql.Rows) error {
		var err error
		var p Product
		err = rows.Scan(&p.ProductId, &p.ProductName)
		if err == nil {
			products = append(products, &p)
		}
		return err
	}

	GetDbListTemplate(queryMap, dbCountFunc, dbRowsFunc, dbListFunc)

	return products, total
}
