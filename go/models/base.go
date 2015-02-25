package models

import (
	"database/sql"
	"encoding/xml"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var (
	DbMap *sql.DB
)

// implemented by structs that need to load values from http request
type Loadable interface {
	LoadValue(prefix string, name string, value []string)
}

func LoadModel(loadable Loadable, data map[string][]string) {
	prefix := ""
	// check if there is any prefix
	prefixs := data["prefix"]
	if prefixs != nil && len(prefixs) > 0 {
		prefix = prefixs[0]
	}
	for key, value := range data {
		(loadable).LoadValue(prefix, key, value)
	}

}

///////////////////////////////////////////////////////////////////////////////////
/*
 * store the changes to the database
 */
type DbChange struct {
	XMLName  xml.Name  `xml:"db"`
	Products []Product `xml:"product"`
}

///////////////////////////////////////////////////////////////////////////////////
/*
 * for creating a function literal that set the total row count variable (in lexical scope)
 * of the list/table, based on the function parameter (i.e. queryMap)
 */
type DbCountFunc func(queryMap *map[string]string) error

/*
 * for creating a function literal that return rows from the database,
 * based on the function parameter (i.e. queryMap) and any lexical scoped variables (e.g. offset, limit)
 */
type DbRowsFunc func(queryMap *map[string]string) (*sql.Rows, error)

/*
 * for creating a function literal that create an object and populate it with values from the row,
 * before appending it to a lexical scoped list
 */
type DbListFunc func(rows *sql.Rows) error

func GetDbListTemplate(queryMap map[string]string, dbCountFunc DbCountFunc, dbRowsFunc DbRowsFunc, dbListFunc DbListFunc) {

	var rows *sql.Rows
	var err error

	err = dbCountFunc(&queryMap)

	if err != nil {
		log.Fatal(err)
	}

	// get rows
	rows, err = dbRowsFunc(&queryMap)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// create list
	for rows.Next() {
		err = dbListFunc(rows)
		if err != nil {
			log.Fatal(err)
		}
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

}

///////////////////////////////////////////////////////////////////////////////////
